(async () => {
  const token = localStorage.getItem('token');
  if (!token) {
    alert('Не авторизованы');
    window.location = '/signin';
    return;
  }

  const res = await fetch('/api/profile', {
    headers: { 'Authorization': `Bearer ${token}` }
  });
  if (!res.ok) {
    alert('Ошибка авторизации');
    window.location = '/signin';
    return;
  }
  const user = await res.json();
  document.getElementById('userId').textContent = user.id;
  document.getElementById('userName').textContent = user.name;
  document.getElementById('userEmail').textContent = user.email;
  document.getElementById('userRating').textContent = user.rating;
  document.getElementById('userTasksSolved').textContent = user.tasksSolved;
})();