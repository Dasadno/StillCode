document.addEventListener('DOMContentLoaded', () => {
  const token = localStorage.getItem('token');
  const authButtons = document.getElementById('authButtons');
  const profileIcon  = document.getElementById('profileIcon');

  if (token) {
    authButtons.classList.add('hidden');
    profileIcon.classList.remove('hidden');
  } else {
    authButtons.classList.remove('hidden');
    profileIcon.classList.add('hidden');
  }
});


const form = document.getElementById('signInForm');
form.addEventListener('submit', async e => {
  e.preventDefault();

  const data = {
    email:    form.email.value.trim(),
    password: form.password.value.trim(),
  };

  try {
    const res = await fetch('/signin', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(data),
    });

    const result = await res.json();

    if (!res.ok) {
      alert(result.error || 'Sign‑in failed');
      return;
    }

 
    localStorage.setItem('token', result.token);
    window.location.href = '/';
  } catch (err) {
    console.error('Fetch error:', err);
    alert('Network error');
  }
});