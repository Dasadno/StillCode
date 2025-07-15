document.addEventListener('DOMContentLoaded', async () => {
  const pathParts = window.location.pathname.split('/');
  const taskId = pathParts[pathParts.length - 1];

  if (!taskId) {
    alert('No task ID found in URL');
    return;
  }

  try {
    const res = await fetch(`/api/task/${taskId}`);
    if (!res.ok) throw new Error('Failed to load task');

    const task = await res.json();

    document.getElementById('taskTitle').textContent = task.title;
    document.getElementById('taskDescription').textContent = task.description;

    const lang = '71';
    document.getElementById('codeEditor').value = task.templates[lang] || '';

    const testBlock = document.getElementById('testCases');
    testBlock.innerHTML = '';
    task.testCases.forEach(tc => {
      testBlock.innerHTML += `🧪 input: ${tc.input} → output: ${tc.expected}<br/>`;
    });
  } catch (err) {
    console.error(err);
    alert('Ошибка при загрузке задачи');
  }
});