/**
 * Sign In Page Script
 */

import { auth } from '../services/auth.js';
import { api } from '../services/api.js';

document.addEventListener('DOMContentLoaded', () => {
  // Redirect if already logged in
  if (auth.redirectIfLoggedIn()) {
    return;
  }

  const form = document.getElementById('signInForm');
  if (!form) {
    console.error('signin.js: signInForm not found');
    return;
  }

  form.addEventListener('submit', handleSubmit);
});

async function handleSubmit(e) {
  e.preventDefault();

  const form = e.target;
  const email = form.email.value.trim();
  const password = form.password.value.trim();

  if (!email || !password) {
    showError('Please fill in all fields');
    return;
  }

  const submitBtn = form.querySelector('button[type="submit"]');
  setLoading(submitBtn, true);

  try {
    const result = await api.signin(email, password);
    auth.setToken(result.token);
    window.location.href = '/';
  } catch (err) {
    console.error('Sign in error:', err);
    showError(err.message || 'Sign in failed');
    setLoading(submitBtn, false);
  }
}

function showError(message) {
  let errorEl = document.getElementById('errorMessage');
  if (!errorEl) {
    errorEl = document.createElement('p');
    errorEl.id = 'errorMessage';
    errorEl.className = 'text-red-400 text-sm text-center mt-4';
    const form = document.getElementById('signInForm');
    if (form) form.appendChild(errorEl);
  }
  errorEl.textContent = message;
}

function setLoading(button, loading) {
  if (!button) return;
  if (loading) {
    button.disabled = true;
    button.textContent = 'Signing in...';
    button.classList.add('opacity-60', 'cursor-not-allowed');
  } else {
    button.disabled = false;
    button.textContent = 'Sign In';
    button.classList.remove('opacity-60', 'cursor-not-allowed');
  }
}
