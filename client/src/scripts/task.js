/**
 * Task Solving Page Script
 */

import { initHeader } from '../components/header.js';
import { auth } from '../services/auth.js';
import { api } from '../services/api.js?v=2';
import { storage } from '../services/storage.js';

// Global CodeMirror editor instance
let editor = null;
let currentTask = null;

document.addEventListener('DOMContentLoaded', () => {
  initHeader();
  initCodeEditor();
  initUI();
  initLanguageDropdown();

  const taskId = getTaskIdFromURL();
  if (taskId) loadTask(taskId);
});

function initCodeEditor() {
  const textarea = document.getElementById('codeEditor');
  if (!textarea) return;

  const taskId = getTaskIdFromURL();
  const savedLang = storage.getLanguage();
  const savedCode = taskId ? storage.getCode(taskId, savedLang) : '';

  editor = CodeMirror.fromTextArea(textarea, {
    mode: getCodeMirrorMode(savedLang),
    theme: 'dracula',
    lineNumbers: true,
    autoCloseBrackets: true,
    matchBrackets: true,
    indentUnit: 4,
    tabSize: 4,
    indentWithTabs: false,
    lineWrapping: false,
    extraKeys: {
      'Tab': (cm) => {
        if (cm.somethingSelected()) {
          cm.indentSelection('add');
        } else {
          cm.replaceSelection('    ', 'end');
        }
      },
      'Shift-Tab': (cm) => cm.indentSelection('subtract'),
      'Ctrl-Enter': () => onRunClicked(),
      'Cmd-Enter': () => onRunClicked(),
      'Ctrl-Shift-Enter': () => onSubmitClicked(),
      'Cmd-Shift-Enter': () => onSubmitClicked(),
    }
  });

  // Restore saved code
  if (savedCode) {
    editor.setValue(savedCode);
  }

  // Auto-save on change
  editor.on('change', () => {
    const taskId = getTaskIdFromURL();
    const currentLang = storage.getLanguage();
    if (taskId) {
      storage.setCode(taskId, editor.getValue(), currentLang);
    }
  });

  // Update language label
  const languageLabel = document.getElementById('languageLabel');
  if (languageLabel && savedLang) {
    languageLabel.textContent = savedLang;
  }
}

function getCodeMirrorMode(language) {
  const modes = {
    'Python': 'python',
    'C++': 'text/x-c++src',
    'Java': 'text/x-java',
    'JavaScript': 'javascript',
    'Go': 'text/x-go'
  };
  return modes[language] || 'python';
}

function setEditorLanguage(language) {
  if (editor) {
    editor.setOption('mode', getCodeMirrorMode(language));
  }
}

// Map display language name to backend key
function getLanguageKey(displayName) {
  const map = {
    'Python': 'python',
    'C++': 'cpp',
    'Java': 'java',
    'JavaScript': 'javascript',
    'Go': 'go'
  };
  return map[displayName] || 'python';
}

// Load code for the selected language (saved or starter)
function loadCodeForLanguage(displayLang) {
  if (!currentTask) return;

  const taskId = getTaskIdFromURL();
  const savedCode = taskId ? storage.getCode(taskId, displayLang) : '';

  // If saved code exists for this language, load it
  if (savedCode && savedCode.trim() !== '') {
    if (editor) editor.setValue(savedCode);
    return;
  }

  // Otherwise load starter code
  if (currentTask.starterCode) {
    const langKey = getLanguageKey(displayLang);
    const starterCode = currentTask.starterCode[langKey];
    if (starterCode && editor) {
      editor.setValue(starterCode);
    }
  }
}

// Legacy alias for backward compatibility
function loadStarterCodeForLanguage(displayLang) {
  loadCodeForLanguage(displayLang);
}

function initUI() {
  const submitBtn = document.getElementById('submitBtn');
  if (submitBtn) {
    submitBtn.addEventListener('click', onSubmitClicked);
  }

  // Create Run button
  const buttonsContainer = submitBtn?.parentNode;
  if (buttonsContainer) {
    const runBtn = document.createElement('button');
    runBtn.id = 'runBtn';
    runBtn.className = 'w-full mt-2 text-white px-4 py-2 rounded-md shadow bg-primary hover:bg-primary-dark transition-colors font-semibold';
    runBtn.textContent = 'Run';
    runBtn.addEventListener('click', onRunClicked);
    buttonsContainer.appendChild(runBtn);
  }

  // Hide results container initially
  const resultsContainer = document.getElementById('resultsContainer');
  if (resultsContainer) {
    resultsContainer.classList.add('hidden');
  }
}

function initLanguageDropdown() {
  const langBtn = document.getElementById('languageBtn');
  const langDropdown = document.getElementById('langDropdown');
  const languageLabel = document.getElementById('languageLabel');

  if (!langBtn || !langDropdown) return;

  langBtn.addEventListener('click', (e) => {
    e.stopPropagation();
    langDropdown.classList.toggle('hidden');
  });

  document.querySelectorAll('#langDropdown li').forEach(li => {
    li.addEventListener('click', (e) => {
      e.stopPropagation();
      const lang = li.textContent.trim();
      const currentLang = languageLabel?.textContent?.trim();

      // Skip if same language
      if (lang === currentLang) {
        langDropdown.classList.add('hidden');
        return;
      }

      // Save current code before switching
      const taskId = getTaskIdFromURL();
      if (taskId && editor && currentLang) {
        storage.setCode(taskId, editor.getValue(), currentLang);
      }

      // Switch language
      if (languageLabel) languageLabel.textContent = lang;
      langDropdown.classList.add('hidden');
      storage.setLanguage(lang);
      setEditorLanguage(lang);

      // Load code for the new language
      loadCodeForLanguage(lang);
    });
  });

  document.addEventListener('click', () => {
    langDropdown.classList.add('hidden');
  });
}

function renderTask(task) {
  const titleEl = document.getElementById('taskTitle');
  const descEl = document.getElementById('taskDescription');
  const metaEl = document.getElementById('taskMeta');
  const testCountEl = document.getElementById('testCount');
  const sideEl = document.getElementById('taskInfo');

  if (titleEl) titleEl.textContent = task.title;

  // Render description as markdown
  if (descEl) {
    if (typeof marked !== 'undefined') {
      marked.setOptions({
        highlight: function(code, lang) {
          if (typeof hljs !== 'undefined' && lang && hljs.getLanguage(lang)) {
            return hljs.highlight(code, { language: lang }).value;
          }
          return code;
        },
        breaks: true
      });
      descEl.innerHTML = marked.parse(task.description || '');
    } else {
      descEl.textContent = task.description;
    }
  }

  if (metaEl) {
    // Clear existing content
    while (metaEl.firstChild) {
      metaEl.removeChild(metaEl.firstChild);
    }

    const difficultySpan = document.createElement('span');
    difficultySpan.className = 'px-2 py-0.5 rounded text-xs ' + getDifficultyClass(task.difficulty);
    difficultySpan.textContent = task.difficulty;

    const solvedSpan = document.createElement('span');
    solvedSpan.className = 'ml-2 text-gray-400';
    solvedSpan.textContent = 'Solved: ' + Math.round(task.solvedPercent || 0) + '%';

    metaEl.appendChild(difficultySpan);
    metaEl.appendChild(solvedSpan);
  }

  if (testCountEl) {
    testCountEl.textContent = 'Tests: ' + (task.test_cases ? task.test_cases.length : 0);
  }

  if (sideEl) {
    requestAnimationFrame(() => {
      sideEl.classList.remove('opacity-0');
      sideEl.classList.add('opacity-100');
    });
  }

  // Render test cases
  const container = document.getElementById('testCases');
  if (container && task.test_cases) {
    while (container.firstChild) {
      container.removeChild(container.firstChild);
    }

    task.test_cases.forEach((tc, i) => {
      const block = document.createElement('div');
      block.className = 'mb-3 p-3 bg-secondary-darker rounded-lg border border-gray-600';

      const testTitle = document.createElement('div');
      testTitle.className = 'text-green-400 font-bold mb-2 flex items-center';
      testTitle.innerHTML = '<i class="fas fa-flask mr-2"></i>Example ' + (i + 1);

      // Parse and format input
      const inputDiv = document.createElement('div');
      inputDiv.className = 'mb-1';
      try {
        const inputData = JSON.parse(tc.input);
        Object.entries(inputData).forEach(([key, value]) => {
          const paramLine = document.createElement('div');
          const keySpan = document.createElement('span');
          keySpan.className = 'text-blue-400';
          keySpan.textContent = key + ' = ';
          const valueSpan = document.createElement('span');
          valueSpan.className = 'text-gray-300';
          valueSpan.textContent = JSON.stringify(value);
          paramLine.appendChild(keySpan);
          paramLine.appendChild(valueSpan);
          inputDiv.appendChild(paramLine);
        });
      } catch {
        const inputLabel = document.createElement('span');
        inputLabel.className = 'text-blue-400';
        inputLabel.textContent = 'Input: ';
        inputDiv.appendChild(inputLabel);
        inputDiv.appendChild(document.createTextNode(tc.input));
      }

      const expectedDiv = document.createElement('div');
      expectedDiv.className = 'mt-2 pt-2 border-t border-gray-700';
      const expectedLabel = document.createElement('span');
      expectedLabel.className = 'text-yellow-400';
      expectedLabel.textContent = 'Output: ';
      const expectedValue = document.createElement('span');
      expectedValue.className = 'text-gray-300';
      expectedValue.textContent = tc.expected;
      expectedDiv.appendChild(expectedLabel);
      expectedDiv.appendChild(expectedValue);

      block.appendChild(testTitle);
      block.appendChild(inputDiv);
      block.appendChild(expectedDiv);
      container.appendChild(block);
    });
  }
}

function getDifficultyClass(difficulty) {
  switch (difficulty?.toLowerCase()) {
    case 'easy': return 'bg-green-900 text-green-300';
    case 'medium': return 'bg-yellow-900 text-yellow-300';
    case 'hard': return 'bg-red-900 text-red-300';
    default: return 'bg-gray-700 text-gray-300';
  }
}

function setLoading(button, loading, text) {
  if (!button) return;
  if (loading) {
    button.dataset.orig = button.textContent;
    button.disabled = true;
    button.textContent = text || 'Loading...';
    button.classList.add('opacity-60', 'cursor-not-allowed');
  } else {
    button.disabled = false;
    button.textContent = button.dataset.orig || button.textContent;
    button.classList.remove('opacity-60', 'cursor-not-allowed');
  }
}

function getEditorCode() {
  return editor ? editor.getValue() : '';
}

function getSelectedLanguage() {
  const langLabel = document.getElementById('languageLabel');
  const map = {
    'Python': 'python',
    'C++': 'cpp',
    'Java': 'java',
    'JavaScript': 'javascript',
    'Go': 'go'
  };
  return map[langLabel?.textContent?.trim()] || 'python';
}

function checkAuth() {
  if (!auth.isLoggedIn()) {
    showOutput('Please sign in to run code.');
    return false;
  }
  return true;
}

async function onRunClicked() {
  if (!checkAuth()) return;

  const runBtn = document.getElementById('runBtn');
  setLoading(runBtn, true, 'Running...');

  const code = getEditorCode();
  const taskId = getTaskIdFromURL();
  const testCases = currentTask?.test_cases || [];
  const input = testCases.length ? testCases[0].input : '';

  try {
    const result = await api.runCode({
      language: getSelectedLanguage(),
      code,
      input,
      taskId: taskId ? Number(taskId) : undefined
    });

    showOutput(formatRunResponse(result));
    hideResults();
    storage.addToHistory({
      type: 'run',
      timestamp: Date.now(),
      language: getSelectedLanguage(),
      result
    });
  } catch (err) {
    if (err.message === 'Unauthorized') {
      showOutput('Session expired. Please sign in again.');
    } else {
      showOutput('Error: ' + err.message);
    }
  } finally {
    setLoading(runBtn, false);
  }
}

async function onSubmitClicked() {
  if (!checkAuth()) return;

  const submitBtn = document.getElementById('submitBtn');
  setLoading(submitBtn, true, 'Submitting...');

  const code = getEditorCode();
  const taskId = getTaskIdFromURL();

  if (!taskId) {
    showOutput('Task ID not found.');
    setLoading(submitBtn, false);
    return;
  }

  try {
    const result = await api.submitCode(taskId, {
      language: getSelectedLanguage(),
      code
    });

    showOutput('');
    renderSubmitResults(result);
    storage.addToHistory({
      type: 'submit',
      taskId,
      timestamp: Date.now(),
      language: getSelectedLanguage(),
      result
    });
  } catch (err) {
    if (err.message === 'Unauthorized') {
      showOutput('Session expired. Please sign in again.');
    } else {
      showOutput('Error: ' + err.message);
    }
  } finally {
    setLoading(submitBtn, false);
  }
}

function showOutput(text) {
  const out = document.getElementById('outputArea');
  if (out) out.textContent = text;
}

function hideResults() {
  const container = document.getElementById('resultsContainer');
  if (container) {
    container.classList.add('hidden');
    while (container.firstChild) {
      container.removeChild(container.firstChild);
    }
  }
}

function formatRunResponse(r) {
  let output = 'Status: ' + r.status + '\nTime: ' + r.time_ms + ' ms\n';
  if (r.stdout) output += '\nOutput:\n' + r.stdout;
  if (r.stderr) output += '\nErrors:\n' + r.stderr;
  return output;
}

function renderSubmitResults(resp) {
  const container = document.getElementById('resultsContainer');
  if (!container) return;

  container.classList.remove('hidden');
  while (container.firstChild) {
    container.removeChild(container.firstChild);
  }

  const allPassed = resp.summary.passed === resp.summary.total;
  const summaryClass = allPassed ? 'text-green-400' : 'text-yellow-400';
  const summaryIcon = allPassed ? 'fa-check-circle' : 'fa-exclamation-circle';

  const summary = document.createElement('div');
  summary.className = 'mb-3 ' + summaryClass + ' font-bold text-lg';

  const icon = document.createElement('i');
  icon.className = 'fas ' + summaryIcon + ' mr-2';

  summary.appendChild(icon);
  summary.appendChild(document.createTextNode(
    'Passed ' + resp.summary.passed + '/' + resp.summary.total + ' - Total time: ' + resp.summary.time_ms + ' ms'
  ));
  container.appendChild(summary);

  resp.results.forEach(r => {
    const row = document.createElement('div');
    row.className = 'p-3 mb-2 rounded border bg-primary-darker';

    let statusClass, statusColor, statusIcon;
    if (r.status === 'passed') {
      statusClass = 'test-passed';
      statusColor = 'border-green-500 text-green-400';
      statusIcon = 'fa-check';
    } else if (r.status === 'timeout') {
      statusClass = 'test-timeout';
      statusColor = 'border-yellow-500 text-yellow-400';
      statusIcon = 'fa-clock';
    } else {
      statusClass = 'test-failed';
      statusColor = 'border-red-500 text-red-400';
      statusIcon = 'fa-times';
    }

    row.classList.add(statusClass);

    // Header row
    const headerDiv = document.createElement('div');
    headerDiv.className = 'flex justify-between items-center ' + statusColor;

    const leftHeader = document.createElement('div');
    const headerIcon = document.createElement('i');
    headerIcon.className = 'fas ' + statusIcon + ' mr-2';
    const strong = document.createElement('strong');
    strong.textContent = 'Test #' + r.test_index;
    leftHeader.appendChild(headerIcon);
    leftHeader.appendChild(strong);
    leftHeader.appendChild(document.createTextNode(' - ' + r.status.toUpperCase()));

    const timeDiv = document.createElement('div');
    timeDiv.className = 'text-xs text-gray-400';
    timeDiv.textContent = r.time_ms + ' ms';

    headerDiv.appendChild(leftHeader);
    headerDiv.appendChild(timeDiv);
    row.appendChild(headerDiv);

    // Details
    const detailsDiv = document.createElement('div');
    detailsDiv.className = 'mt-2 font-mono text-xs';

    const inputLine = document.createElement('div');
    const inputSpan = document.createElement('span');
    inputSpan.className = 'text-blue-300';
    inputSpan.textContent = 'Input: ';
    inputLine.appendChild(inputSpan);
    inputLine.appendChild(document.createTextNode(r.input || ''));

    const expectedLine = document.createElement('div');
    const expectedSpan = document.createElement('span');
    expectedSpan.className = 'text-green-300';
    expectedSpan.textContent = 'Expected: ';
    expectedLine.appendChild(expectedSpan);
    expectedLine.appendChild(document.createTextNode(r.expected || ''));

    const outputLine = document.createElement('div');
    const outputSpan = document.createElement('span');
    outputSpan.className = 'text-yellow-300';
    outputSpan.textContent = 'Output: ';
    outputLine.appendChild(outputSpan);
    outputLine.appendChild(document.createTextNode(r.output || ''));

    detailsDiv.appendChild(inputLine);
    detailsDiv.appendChild(expectedLine);
    detailsDiv.appendChild(outputLine);
    row.appendChild(detailsDiv);

    container.appendChild(row);
  });
}

function getTaskIdFromURL() {
  const path = window.location.pathname.split('/');
  const last = path[path.length - 1];
  if (!isNaN(last) && last !== '') return Number(last);
  const params = new URLSearchParams(window.location.search);
  return params.get('id');
}

async function loadTask(id) {
  try {
    const task = await api.getTask(id);
    currentTask = task;
    renderTask(task);
    // Load code for the current language (saved or starter)
    const currentLang = storage.getLanguage() || 'Python';
    loadCodeForLanguage(currentLang);
  } catch (err) {
    showOutput('Failed to load task: ' + err.message);
  }
}
