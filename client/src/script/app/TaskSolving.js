document.addEventListener('DOMContentLoaded', async () => {
  const pathParts = window.location.pathname.split('/');
  const taskId = pathParts[pathParts.length - 1]; 

  if (!taskId) {
    alert('No task ID found in URL');
    return;
  }

  const setText = (id, value) => {
    const el = document.getElementById(id);
    if (el) el.textContent = value ?? '';
    else console.warn(`Element with id="${id}" not found`);
  };

  try {
    const res = await fetch(`/api/task/${taskId}`);
    if (!res.ok) throw new Error('Failed to load task');

    const task = await res.json();

    setText('taskTitle', task.title);
    setText('taskDescription', task.description); 

    const lang = '71';
    const editor = document.getElementById('codeEditor');
    if (editor) editor.value = task.templates?.[lang] || '';

    const testBlock = document.getElementById('testCases');
    if (testBlock) {
      testBlock.innerHTML = '';
      task.testCases.forEach(tc => {
        testBlock.innerHTML += `input: ${tc.input} → output: ${tc.expected}<br/>`;
      });
    }
  } catch (err) {
    console.error(err);
    alert('Ошибка при загрузке задачи');
  }
});