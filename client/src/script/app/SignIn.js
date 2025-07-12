document.addEventListener('DOMContentLoaded', () => {
  const form = document.getElementById('signInForm');
  if (!form) {
    console.error('SignIn.js: form not found');
    return;
  }

  form.addEventListener('submit', async e => {
    e.preventDefault();               

    const email    = form.email.value.trim();
    const password = form.password.value.trim();

    try {
      const res = await fetch('/signin', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ email, password }),
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
});