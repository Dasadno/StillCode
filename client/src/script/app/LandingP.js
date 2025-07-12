document.addEventListener('DOMContentLoaded', async () => {
  const token = localStorage.getItem('token');
  const authButtons = document.getElementById('authButtons');
  const userWidget  = document.getElementById('userWidget');
  const profileName = document.getElementById('profileName');

  if (!token) {
    authButtons.style.display = 'flex';
    userWidget.style.display  = 'none';
    return;
  }


  authButtons.style.display = 'none';
  userWidget.style.display  = 'flex';



  try {
    const res = await fetch('/api/profile', {
      headers: { 'Authorization': `Bearer ${token}` }
    });
    if (!res.ok) throw new Error('bad profile fetch');
    const user = await res.json();

    profileName.textContent = user.name;
  } catch (err) {
    console.warn('Не удалось загрузить профиль:', err);
    profileName.textContent = 'User';
  }
});


document.addEventListener('ProblemsButton', async () => {
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
  else {
    window.location = '/problems';
  }
});