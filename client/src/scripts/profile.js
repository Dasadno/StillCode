/**
 * Profile Page Script
 */

import { initHeader } from '../components/header.js';
import { auth } from '../services/auth.js';
import { api } from '../services/api.js';

document.addEventListener('DOMContentLoaded', () => {
  // Require authentication
  if (!auth.requireAuth()) {
    return;
  }

  initHeader();
  loadProfile();
  setupLogout();
});

async function loadProfile() {
  try {
    const user = await api.getProfile();

    setText('userName', user.name || 'Unknown');
    setText('userEmail', user.email || '-');
    setText('userRating', user.rating != null ? user.rating : '0');
    setText('userTasksSolved', user.tasksSolved != null ? user.tasksSolved : '0');
  } catch (err) {
    console.error('Failed to load profile:', err);

    if (err.message === 'Unauthorized') {
      // Token is invalid, redirect to signin
      auth.logout('/signin');
      return;
    }

    showError('Failed to load profile. Please try again.');
  }
}

function setText(id, value) {
  const el = document.getElementById(id);
  if (el) {
    el.textContent = value;
  }
}

function showError(message) {
  const main = document.querySelector('main');
  if (main) {
    while (main.firstChild) {
      main.removeChild(main.firstChild);
    }

    const container = document.createElement('div');
    container.className = 'max-w-lg mx-auto bg-darker p-8 rounded-xl shadow-lg text-center';

    const errorP = document.createElement('p');
    errorP.className = 'text-red-400 mb-4';
    errorP.textContent = message;

    const retryBtn = document.createElement('button');
    retryBtn.className = 'text-primary hover:underline';
    retryBtn.textContent = 'Try again';
    retryBtn.addEventListener('click', () => window.location.reload());

    container.appendChild(errorP);
    container.appendChild(retryBtn);
    main.appendChild(container);
  }
}

function setupLogout() {
  const logoutBtn = document.getElementById('logoutBtn');
  if (logoutBtn) {
    logoutBtn.addEventListener('click', () => {
      auth.logout('/');
    });
  }
}
