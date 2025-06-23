document.getElementById('signInForm').addEventListener('signin', async e => {
  e.preventDefault();
  const form = e.target;
  const data = {
    email: form.email.value,
    password: form.password.value
  };
  const res = await fetch('/signin', {
    method: 'POST',
    headers: {'Content-Type':'application/json'},
    body: JSON.stringify(data)
  });
  const json = await res.json();
  if (res.ok) {
    localStorage.setItem('token', json.token);
    window.location.href = '/profile.html';
  } else {
    alert(json.error);
  }
});