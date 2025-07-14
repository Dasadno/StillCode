const pageSize = 20;
    let currentPage = 1;
    let totalPages  = 1; 

    async function loadTasks(page = 1) {
      currentPage = page;
      const search     = document.getElementById('searchInput').value;
      const difficulty = document.getElementById('difficultyFilter').value;
      const community  = document.getElementById('communityFilter').value;

      const params = new URLSearchParams({search, difficulty, community, page, pageSize});
      const res    = await fetch('/api/tasks?' + params);
      const tasks  = await res.json();

      const ul = document.getElementById('tasksList');
      ul.innerHTML = '';
      tasks.forEach(t => {
        const li = document.createElement('li');
        li.className = 'p-4 bg-darker rounded-md shadow flex justify-between items-center';
        li.innerHTML = `
          <div>
            <h3 class="text-xl font-semibold">${t.title}</h3>
            <p class="text-sm text-gray-400">
              Difficulty: ${t.difficulty} •
              ${t.isCommunity ? 'Community' : 'Official'} •
              Solved: ${t.solvedPercent.toFixed(1)}%
            </p>
          </div>
          <a href="/task/${t.id}" class="text-primary hover:underline">View</a>
        `;
        ul.appendChild(li);
      });


      const totalCount = parseInt(res.headers.get('X-Total-Count') || (tasks.length + (page - 1) * pageSize));
      totalPages = Math.ceil(totalCount / pageSize);

      renderPagination();
    }

    function renderPagination() {
      const pg = document.getElementById('pagination');
      pg.innerHTML = '';


      const prev = document.createElement('button');
      prev.textContent = '← Prev';
      prev.disabled = currentPage === 1;
      prev.className = `px-3 py-1 rounded ${prev.disabled ? 'bg-gray-700 cursor-not-allowed' : 'bg-primary hover:bg-primary-dark'}`;
      prev.onclick = () => loadTasks(currentPage - 1);
      pg.appendChild(prev);


      const start = Math.max(1, currentPage - 2);
      const end   = Math.min(totalPages, currentPage + 2);
      for (let p = start; p <= end; p++) {
        const btn = document.createElement('button');
        btn.textContent = p;
        btn.className = p === currentPage
          ? 'px-3 py-1 rounded bg-primary'
          : 'px-3 py-1 rounded bg-darker hover:bg-gray-700';
        btn.onclick = () => loadTasks(p);
        pg.appendChild(btn);
      }

 
      const next = document.createElement('button');
      next.textContent = 'Next →';
      next.disabled = currentPage === totalPages;
      next.className = `px-3 py-1 rounded ${next.disabled ? 'bg-gray-700 cursor-not-allowed' : 'bg-primary hover:bg-primary-dark'}`;
      next.onclick = () => loadTasks(currentPage + 1);
      pg.appendChild(next);
    }

    document.getElementById('applyBtn').addEventListener('click', () => loadTasks(1));
    window.addEventListener('DOMContentLoaded', () => loadTasks(1));