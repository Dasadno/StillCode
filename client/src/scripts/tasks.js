/**
 * Tasks List Page Script
 */

import { initHeader } from '../components/header.js';
import { api } from '../services/api.js';

const PAGE_SIZE = 20;
let currentPage = 1;
let totalPages = 1;

document.addEventListener('DOMContentLoaded', () => {
  initHeader();

  const applyBtn = document.getElementById('applyBtn');
  if (applyBtn) {
    applyBtn.addEventListener('click', () => loadTasks(1));
  }

  const searchInput = document.getElementById('searchInput');
  if (searchInput) {
    searchInput.addEventListener('keypress', (e) => {
      if (e.key === 'Enter') {
        loadTasks(1);
      }
    });
  }

  loadTasks(1);
});

async function loadTasks(page = 1) {
  currentPage = page;
  const searchInput = document.getElementById('searchInput');
  const difficultyFilter = document.getElementById('difficultyFilter');
  const communityFilter = document.getElementById('communityFilter');
  const tasksList = document.getElementById('tasksList');

  const search = searchInput ? searchInput.value : '';
  const difficulty = difficultyFilter ? difficultyFilter.value : '';
  const community = communityFilter ? communityFilter.value : '';

  try {
    const { tasks, totalCount } = await api.getTasks({
      search,
      difficulty,
      community,
      page,
      pageSize: PAGE_SIZE
    });

    if (!tasksList) return;

    // Clear existing list
    while (tasksList.firstChild) {
      tasksList.removeChild(tasksList.firstChild);
    }

    if (!tasks || tasks.length === 0) {
      const li = document.createElement('li');
      li.className = 'p-4 bg-darker rounded-md text-center text-gray-400';
      li.textContent = 'No problems found';
      tasksList.appendChild(li);
      totalPages = 1;
      renderPagination();
      return;
    }

    tasks.forEach(t => {
      const li = createTaskCard(t);
      tasksList.appendChild(li);
    });

    totalPages = Math.max(1, Math.ceil(totalCount / PAGE_SIZE));
    renderPagination();
  } catch (error) {
    console.error('Failed to load tasks:', error);
    if (tasksList) {
      while (tasksList.firstChild) {
        tasksList.removeChild(tasksList.firstChild);
      }

      const li = document.createElement('li');
      li.className = 'p-4 bg-darker rounded-md text-center';

      const errorP = document.createElement('p');
      errorP.className = 'text-red-400 mb-2';
      errorP.textContent = 'Failed to load problems';

      const retryBtn = document.createElement('button');
      retryBtn.className = 'text-primary hover:underline';
      retryBtn.textContent = 'Try again';
      retryBtn.addEventListener('click', () => loadTasks(page));

      li.appendChild(errorP);
      li.appendChild(retryBtn);
      tasksList.appendChild(li);
    }
  }
}

function createTaskCard(task) {
  const li = document.createElement('li');
  li.className = 'p-4 bg-darker rounded-md shadow flex justify-between items-center task-card cursor-pointer hover:bg-secondary-dark transition-colors';

  // Left side: title and metadata
  const leftDiv = document.createElement('div');

  const title = document.createElement('h3');
  title.className = 'text-xl font-semibold';
  title.textContent = task.title;

  const meta = document.createElement('p');
  meta.className = 'text-sm text-gray-400';

  const difficultySpan = document.createElement('span');
  difficultySpan.className = 'px-2 py-0.5 rounded text-xs ' + getDifficultyClass(task.difficulty);
  difficultySpan.textContent = task.difficulty;

  const typeSpan = document.createElement('span');
  typeSpan.className = 'ml-2';
  typeSpan.textContent = task.isCommunity ? 'Community' : 'Official';

  const solvedSpan = document.createElement('span');
  solvedSpan.className = 'ml-2';
  const solvedPercent = task.solvedPercent != null ? task.solvedPercent.toFixed(1) : '0.0';
  solvedSpan.textContent = 'Solved: ' + solvedPercent + '%';

  meta.appendChild(difficultySpan);
  meta.appendChild(typeSpan);
  meta.appendChild(solvedSpan);

  leftDiv.appendChild(title);
  leftDiv.appendChild(meta);

  // Right side: solve link
  const solveLink = document.createElement('a');
  solveLink.href = '/task/' + task.id;
  solveLink.className = 'text-primary hover:underline font-medium';
  solveLink.textContent = 'Solve';

  li.appendChild(leftDiv);
  li.appendChild(solveLink);

  li.addEventListener('click', (e) => {
    if (e.target.tagName !== 'A') {
      window.location.href = '/task/' + task.id;
    }
  });

  return li;
}

function getDifficultyClass(difficulty) {
  switch (difficulty?.toLowerCase()) {
    case 'easy': return 'badge-easy bg-green-900 text-green-300';
    case 'medium': return 'badge-medium bg-yellow-900 text-yellow-300';
    case 'hard': return 'badge-hard bg-red-900 text-red-300';
    default: return 'bg-gray-700 text-gray-300';
  }
}

function renderPagination() {
  const pg = document.getElementById('pagination');
  if (!pg) return;

  // Clear existing pagination
  while (pg.firstChild) {
    pg.removeChild(pg.firstChild);
  }

  // Prev button
  const prev = document.createElement('button');
  prev.textContent = 'Prev';
  prev.disabled = currentPage === 1;
  prev.className = 'px-3 py-1 rounded pagination-btn ' + (prev.disabled ? 'bg-gray-700 cursor-not-allowed opacity-50' : 'bg-darker hover:bg-primary');
  prev.addEventListener('click', () => loadTasks(currentPage - 1));
  pg.appendChild(prev);

  // Page numbers
  const start = Math.max(1, currentPage - 2);
  const end = Math.min(totalPages, currentPage + 2);

  for (let p = start; p <= end; p++) {
    const btn = document.createElement('button');
    btn.textContent = p;
    btn.className = p === currentPage
      ? 'px-3 py-1 rounded bg-primary text-white pagination-active'
      : 'px-3 py-1 rounded bg-darker hover:bg-gray-700 pagination-btn';
    btn.addEventListener('click', () => loadTasks(p));
    pg.appendChild(btn);
  }

  // Next button
  const next = document.createElement('button');
  next.textContent = 'Next';
  next.disabled = currentPage >= totalPages;
  next.className = 'px-3 py-1 rounded pagination-btn ' + (next.disabled ? 'bg-gray-700 cursor-not-allowed opacity-50' : 'bg-darker hover:bg-primary');
  next.addEventListener('click', () => loadTasks(currentPage + 1));
  pg.appendChild(next);
}

// Export for potential external use
window.loadTasks = loadTasks;
