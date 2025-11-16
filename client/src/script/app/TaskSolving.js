document.addEventListener("DOMContentLoaded", () => {
  initUI();
  const taskId = getTaskIdFromURL();
  if (taskId) loadTask(taskId);
});

function initUI() {
  // Добавляем кнопку Run рядом с Submit
  const submitBtn = document.getElementById("submitBtn");
  submitBtn.textContent = "Submit";

  // create Run button
  const runBtn = document.createElement("button");
  runBtn.id = "runBtn";
  runBtn.className = "ml-3 text-white px-4 py-2 rounded-md shadow bg-primary hover:bg-primary-dark transition";
  runBtn.textContent = "Run";
  submitBtn.insertAdjacentElement("afterend", runBtn);

  runBtn.addEventListener("click", onRunClicked);
  submitBtn.addEventListener("click", onSubmitClicked);

  // create result area if not present
  let outputArea = document.getElementById("outputArea");
  if (!outputArea) {
    outputArea = document.createElement("div");
    outputArea.id = "outputArea";
    outputArea.className = "mt-4 p-4 rounded-md bg-primary-darker border border-gray-700 h-64 overflow-auto font-mono text-sm";
    const editorSection = document.querySelector("section.flex-1");
    editorSection.insertBefore(outputArea, document.getElementById("testCases"));
  }

  // create results table container
  let resultsContainer = document.getElementById("resultsContainer");
  if (!resultsContainer) {
    resultsContainer = document.createElement("div");
    resultsContainer.id = "resultsContainer";
    resultsContainer.className = "mt-3 p-3 rounded-md bg-primary-darker border border-gray-700 text-sm";
    document.getElementById("testCases").parentNode.appendChild(resultsContainer);
  }
}

function renderTask(task) {
    // Title, description
    const titleEl = document.getElementById("taskTitle");
    const descEl = document.getElementById("taskDescription");
    const metaEl = document.getElementById("taskMeta");
    const testCountEl = document.getElementById("testCount");
    const sideEl = document.getElementById("taskInfo");

    titleEl.textContent = task.title;
    descEl.textContent = task.description;

    // Beautiful task meta (difficulty + solved %)
    metaEl.innerHTML = `
        <span class="text-primary">Сложность:</span> ${task.difficulty}<br>
        <span class="text-primary">Решили:</span> ${Math.round(task.solvedPercent)}%
    `;

    // Test count
    testCountEl.textContent = `Тестов: ${task.test_cases.length}`;

    // Smooth fade-in animation
    requestAnimationFrame(() => {
        sideEl.classList.remove("opacity-0");
        sideEl.classList.add("opacity-100");
    });

    // Render testcases
    const container = document.getElementById("testCases");
    container.innerHTML = "";

    task.test_cases.forEach((tc, i) => {
        const block = document.createElement("div");
        block.className = "mb-2 p-2 border-b border-gray-700";

        block.innerHTML = `
            <div class="text-green-400 font-bold">Test #${i + 1}</div>
            <div><span class="text-blue-400">Input:</span> ${tc.input}</div>
            <div><span class="text-red-400">Expected:</span> ${tc.expected}</div>
        `;

        container.appendChild(block);
    });
}

function setLoading(button, loading, text) {
  if (loading) {
    button.dataset.orig = button.textContent;
    button.disabled = true;
    button.textContent = text || "Running...";
    button.classList.add("opacity-60", "cursor-not-allowed");
  } else {
    button.disabled = false;
    button.textContent = button.dataset.orig || button.textContent;
    button.classList.remove("opacity-60", "cursor-not-allowed");
  }
}

function getEditorCode() {
  const codeEl = document.getElementById("codeEditor");
  return codeEl ? codeEl.value : "";
}

function getSelectedLanguage() {
  const langLabel = document.getElementById("languageLabel");
  const map = { "Python": "python", "C++": "cpp", "Java": "java" };
  return map[langLabel?.textContent?.trim()] || "python";
}

async function onRunClicked() {
  const runBtn = document.getElementById("runBtn");
  setLoading(runBtn, true, "Running…");

  const code = getEditorCode();
  // take first test input if present
  const testCases = window._currentTask?.test_cases || [];
  const input = testCases.length ? testCases[0].input : "";

  const payload = {
    language: getSelectedLanguage(),
    code,
    input,
  };

  try {
    const res = await fetch("/api/run", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(payload),
    });
    if (!res.ok) {
      const text = await res.text();
      showOutput(`Server error: ${text}`);
      setLoading(runBtn, false);
      return;
    }
    const json = await res.json();
    showOutput(formatRunResponse(json));
    pushHistory({ type: "run", timestamp: Date.now(), payload, result: json });
  } catch (err) {
    showOutput("Network error: " + err.message);
  } finally {
    setLoading(runBtn, false);
  }
}

async function onSubmitClicked() {
  const submitBtn = document.getElementById("submitBtn");
  setLoading(submitBtn, true, "Submitting…");

  const code = getEditorCode();
  const payload = { language: getSelectedLanguage(), code };

  const taskId = getTaskIdFromURL();
  if (!taskId) {
    showOutput("Task ID not found.");
    setLoading(submitBtn, false);
    return;
  }

  try {
    const res = await fetch(`/api/submit/${taskId}`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(payload),
    });
    if (!res.ok) {
      const t = await res.text();
      showOutput("Server error: " + t);
      setLoading(submitBtn, false);
      return;
    }
    const json = await res.json();
    renderSubmitResults(json);
    pushHistory({ type: "submit", taskId, timestamp: Date.now(), payload, result: json });
  } catch (err) {
    showOutput("Network error: " + err.message);
  } finally {
    setLoading(submitBtn, false);
  }
}

// simple output area writer
function showOutput(text) {
  const out = document.getElementById("outputArea");
  out.textContent = text;
}

// format /api/run response
function formatRunResponse(r) {
  return `Status: ${r.status}\nTime: ${r.time_ms} ms\n\nStdout:\n${r.stdout}\n\nStderr:\n${r.stderr}`;
}

// render submit results table and summary
function renderSubmitResults(resp) {
  window._lastSubmit = resp;
  const container = document.getElementById("resultsContainer");
  container.innerHTML = "";

  const summary = document.createElement("div");
  summary.className = "mb-3";
  summary.innerHTML = `<strong>Passed ${resp.summary.passed}/${resp.summary.total}</strong> — total ${resp.summary.time_ms} ms`;
  container.appendChild(summary);

  // table-like list
  resp.results.forEach(r => {
    const row = document.createElement("div");
    row.className = "p-3 mb-2 rounded border";
    // status color
    const color = r.status === "passed" ? "border-green-500" : (r.status === "timeout" ? "border-yellow-400" : "border-red-500");
    row.classList.add(color, "bg-primary-darker");
    row.innerHTML = `
      <div class="flex justify-between items-center">
        <div><strong>Test #${r.test_index}</strong> — ${r.status.toUpperCase()}</div>
        <div class="text-xs">time ${r.time_ms} ms</div>
      </div>
      <div class="mt-2 font-mono text-xs">
        <div><span class="text-blue-300">Input:</span> ${escapeHtml(r.input)}</div>
        <div><span class="text-green-300">Expected:</span> ${escapeHtml(r.expected)}</div>
        <div><span class="text-yellow-300">Output:</span> ${escapeHtml(r.output)}</div>
      </div>
    `;
    container.appendChild(row);
  });
}

// simple escaping
function escapeHtml(s) {
  if (s == null) return "";
  return s.replace(/&/g, "&amp;").replace(/</g,"&lt;").replace(/>/g,"&gt;");
}

// local history in localStorage
function pushHistory(item) {
  try {
    const key = "stillcode_history_v1";
    const raw = localStorage.getItem(key);
    const arr = raw ? JSON.parse(raw) : [];
    arr.unshift(item);
    if (arr.length > 50) arr.pop();
    localStorage.setItem(key, JSON.stringify(arr));
  } catch (e) {
    console.warn("history save failed", e);
  }
}

// helper functions from previous messages
function getTaskIdFromURL() {
  const path = window.location.pathname.split("/");
  const last = path[path.length - 1];
  if (!isNaN(last)) return Number(last);
  const params = new URLSearchParams(window.location.search);
  return params.get("id");
}

async function loadTask(id) {
  try {
    const res = await fetch(`/api/tasks/${id}`, {headers: {"Content-Type":"application/json"}});
    if (!res.ok) {
      showOutput("Failed to load task: " + await res.text());
      return;
    }
    const task = await res.json();
    window._currentTask = task;
    renderTask(task);  // from earlier code that sets title/description/testcases
    // show test count in side
    const testCountEl = document.getElementById("testCount");
    if (testCountEl) testCountEl.textContent = `Тестов: ${task.test_cases.length}`;
  } catch (err) {
    showOutput("Network error when loading task: " + err.message);
  }
}