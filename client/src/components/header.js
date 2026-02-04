/**
 * Header Component - Reusable site header with navigation
 */

import { auth } from '../services/auth.js';
import { api } from '../services/api.js';

/**
 * Create header element using DOM methods
 * @param {Object} options - Header options
 * @param {boolean} options.showNav - Show navigation links
 * @param {boolean} options.minimal - Minimal header (just logo)
 */
export function createHeader(options = {}) {
  const { showNav = true, minimal = false } = options;

  const header = document.createElement('header');
  header.className = 'bg-darker py-4 sticky top-0 z-50 shadow-lg';

  const container = document.createElement('div');
  container.className = 'container mx-auto px-4 flex justify-between items-center';

  // Logo section
  const logoSection = document.createElement('div');
  logoSection.className = 'flex items-center';

  const logoIcon = document.createElement('div');
  logoIcon.className = 'text-primary text-3xl font-bold mr-2';
  const icon = document.createElement('i');
  icon.className = 'fas fa-code';
  logoIcon.appendChild(icon);

  const logoLink = document.createElement('a');
  logoLink.href = '/';
  const logoH1 = document.createElement('h1');
  logoH1.className = 'text-2xl font-bold';
  logoH1.textContent = 'Still';
  const logoSpan = document.createElement('span');
  logoSpan.className = 'text-primary';
  logoSpan.textContent = 'Code';
  logoH1.appendChild(logoSpan);
  logoLink.appendChild(logoH1);

  logoSection.appendChild(logoIcon);
  logoSection.appendChild(logoLink);
  container.appendChild(logoSection);

  if (!minimal) {
    // Navigation
    if (showNav) {
      const nav = document.createElement('nav');
      nav.className = 'hidden md:flex space-x-8';

      const navLinks = [
        { href: '/problems', id: 'problemsButton', text: 'Problems' },
        { href: '/features', id: 'featuresButton', text: 'Features' },
        { href: '/community', id: 'communityButton', text: 'Community' }
      ];

      navLinks.forEach(link => {
        const a = document.createElement('a');
        a.href = link.href;
        a.id = link.id;
        a.className = 'hover:text-primary transition-colors';
        a.textContent = link.text;
        nav.appendChild(a);
      });

      container.appendChild(nav);
    }

    // User widget (hidden by default, shown when logged in)
    const userWidget = document.createElement('div');
    userWidget.id = 'userWidget';
    userWidget.className = 'flex items-center space-x-2 hidden cursor-pointer';
    userWidget.onclick = () => { window.location.href = '/profile'; };

    const avatarSvg = document.createElementNS('http://www.w3.org/2000/svg', 'svg');
    avatarSvg.id = 'profileAvatar';
    avatarSvg.setAttribute('viewBox', '0 0 24 24');
    avatarSvg.setAttribute('fill', 'currentColor');
    avatarSvg.className.baseVal = 'w-10 h-10 text-primary';
    const avatarPath = document.createElementNS('http://www.w3.org/2000/svg', 'path');
    avatarPath.setAttribute('d', 'M12 12c2.21 0 4-1.79 4-4s-1.79-4-4-4-4 1.79-4 4 1.79 4 4 4zm0 2c-2.67 0-8 1.34-8 4v2h16v-2c0-2.66-5.33-4-8-4z');
    avatarSvg.appendChild(avatarPath);

    const profileName = document.createElement('span');
    profileName.id = 'profileName';
    profileName.className = 'text-primary font-medium text-lg';

    userWidget.appendChild(avatarSvg);
    userWidget.appendChild(profileName);
    container.appendChild(userWidget);

    // Auth buttons
    const authButtons = document.createElement('div');
    authButtons.id = 'authButtons';
    authButtons.className = 'flex items-center space-x-4';

    const signInLink = document.createElement('a');
    signInLink.href = '/signin';
    signInLink.className = 'hover:text-primary transition-colors';
    signInLink.textContent = 'Sign In';

    const signUpLink = document.createElement('a');
    signUpLink.href = '/signup';
    signUpLink.className = 'bg-primary hover:bg-primary-dark text-white px-4 py-2 rounded-md transition-colors';
    signUpLink.textContent = 'Sign Up';

    authButtons.appendChild(signInLink);
    authButtons.appendChild(signUpLink);
    container.appendChild(authButtons);

    // Mobile menu button
    const mobileBtn = document.createElement('button');
    mobileBtn.className = 'md:hidden text-2xl';
    mobileBtn.id = 'mobileMenuBtn';
    const menuIcon = document.createElement('i');
    menuIcon.className = 'fas fa-bars';
    mobileBtn.appendChild(menuIcon);
    container.appendChild(mobileBtn);
  }

  header.appendChild(container);
  return header;
}

/**
 * Initialize header functionality (auth state, profile loading)
 */
export async function initHeader() {
  const authButtons = document.getElementById('authButtons');
  const userWidget = document.getElementById('userWidget');
  const profileName = document.getElementById('profileName');
  const featuresButton = document.getElementById('featuresButton');
  const problemsButton = document.getElementById('problemsButton');
  const communityButton = document.getElementById('communityButton');

  if (!auth.isLoggedIn()) {
    // Not logged in
    if (authButtons) authButtons.style.display = 'flex';
    if (userWidget) userWidget.style.display = 'none';

    // Redirect nav buttons to signin for non-authenticated users
    if (featuresButton) featuresButton.href = '/signin';
    if (problemsButton) problemsButton.href = '/signin';
    if (communityButton) communityButton.href = '/signin';
    return;
  }

  // Logged in
  if (authButtons) authButtons.style.display = 'none';
  if (userWidget) userWidget.style.display = 'flex';

  // Restore nav links for authenticated users
  if (featuresButton) featuresButton.href = '/features';
  if (problemsButton) problemsButton.href = '/problems';
  if (communityButton) communityButton.href = '/community';

  // Load profile name
  try {
    const user = await api.getProfile();
    if (profileName) {
      profileName.textContent = user.name || 'User';
    }
  } catch (err) {
    console.warn('Failed to load profile:', err);
    if (profileName) {
      profileName.textContent = 'User';
    }
  }
}

/**
 * Inject header into a container element
 */
export function injectHeader(containerId = 'header-container', options = {}) {
  const container = document.getElementById(containerId);
  if (container) {
    container.appendChild(createHeader(options));
    initHeader();
  }
}

export default { createHeader, initHeader, injectHeader };
