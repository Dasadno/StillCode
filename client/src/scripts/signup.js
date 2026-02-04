/**
 * Sign Up Page Script
 */

import { auth } from '../services/auth.js';
import { api } from '../services/api.js';

document.addEventListener('DOMContentLoaded', () => {
  // Redirect if already logged in
  if (auth.redirectIfLoggedIn()) {
    return;
  }

  const form = document.getElementById('registerForm');
  if (!form) {
    console.error('signup.js: registerForm not found');
    return;
  }

  form.addEventListener('submit', handleSubmit);
});

async function handleSubmit(e) {
  e.preventDefault();

  const form = e.target;
  const name = form.name.value.trim();
  const email = form.email.value.trim();
  const password = form.password.value.trim();

  if (!name || !email || !password) {
    showError('Please fill in all fields');
    return;
  }

  if (password.length < 8) {
    showError('Password must be at least 8 characters');
    return;
  }

  const submitBtn = form.querySelector('button[type="submit"]');
  setLoading(submitBtn, true);

  try {
    await api.signup(name, email, password);
    showSuccess('Registration successful! Redirecting to sign in...');
    setTimeout(() => {
      window.location.href = '/signin';
    }, 1500);
  } catch (err) {
    console.error('Sign up error:', err);
    showError(err.message || 'Registration failed');
    setLoading(submitBtn, false);
  }
}

function showError(message) {
  clearMessages();
  let errorEl = document.getElementById('errorMessage');
  if (!errorEl) {
    errorEl = document.createElement('p');
    errorEl.id = 'errorMessage';
    errorEl.className = 'text-red-400 text-sm text-center mt-4';
    const form = document.getElementById('registerForm');
    if (form) form.appendChild(errorEl);
  }
  errorEl.textContent = message;
}

function showSuccess(message) {
  clearMessages();
  let successEl = document.getElementById('successMessage');
  if (!successEl) {
    successEl = document.createElement('p');
    successEl.id = 'successMessage';
    successEl.className = 'text-green-400 text-sm text-center mt-4';
    const form = document.getElementById('registerForm');
    if (form) form.appendChild(successEl);
  }
  successEl.textContent = message;
}

function clearMessages() {
  const errorEl = document.getElementById('errorMessage');
  const successEl = document.getElementById('successMessage');
  if (errorEl) errorEl.textContent = '';
  if (successEl) successEl.textContent = '';
}

function setLoading(button, loading) {
  if (!button) return;
  if (loading) {
    button.disabled = true;
    button.textContent = 'Signing up...';
    button.classList.add('opacity-60', 'cursor-not-allowed');
  } else {
    button.disabled = false;
    button.textContent = 'Sign Up';
    button.classList.remove('opacity-60', 'cursor-not-allowed');
  }
}
