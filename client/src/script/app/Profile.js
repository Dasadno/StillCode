document.addEventListener('DOMContentLoaded', () => {
  (async () => {
     const token = localStorage.getItem('token'); 
    try {
      if (!token) {
        alert('Не авторизованы');
        window.location = '/signin';
        return;
      }

      const res = await fetch('/api/profile', {
        headers: { 'Authorization': `Bearer ${token}` }
      });

      if (res.status === 401 || res.status === 403) {
        localStorage.removeItem('token');
        alert('Сессия истекла или нет прав — нужно войти заново');
        window.location = '/signin';
        return;
      }

      if (!res.ok) {
        const text = await res.text().catch(() => null);
        console.error('Ошибка при получении профиля:', res.status, text);
        alert('Ошибка при получении профиля');
        window.location = '/signin';
        return;
      }

      const user = await res.json();

      const setText = (id, value) => {
        const el = document.getElementById(id);
        if (el) el.textContent = value ?? '';
        else console.warn(`Element with id="${id}" not found`);
      };
      console.log('user:', user);
      setText('userId', user.id);
      setText('userName', user.name);
      setText('userEmail', user.email);
      setText('userRating', user.rating);
      setText('userTasksSolved', user.tasksSolved);

    } catch (err) {
      console.error('Ошибка загрузки профиля:', err);
      alert('Сетевая или серверная ошибка при загрузке профиля');
    } finally {
      const logoutBtn = document.getElementById('logoutBtn');
      if (logoutBtn) {
        logoutBtn.addEventListener('click', () => {
          localStorage.removeItem('token');
          window.location = '/';
        });
      } else {
        console.warn('logoutBtn not found — кнопка выхода отсутствует в DOM');
      }
    }
  })();
});