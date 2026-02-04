/**
 * Auth Service - Token management and authentication state
 */

import { storage, STORAGE_KEYS } from './storage.js';

const TOKEN_KEY = STORAGE_KEYS.TOKEN;

export const auth = {
  /**
   * Get the current auth token
   */
  getToken() {
    return storage.get(TOKEN_KEY);
  },

  /**
   * Set the auth token
   */
  setToken(token) {
    return storage.set(TOKEN_KEY, token);
  },

  /**
   * Remove the auth token (logout)
   */
  removeToken() {
    return storage.remove(TOKEN_KEY);
  },

  /**
   * Check if user is logged in
   */
  isLoggedIn() {
    return !!this.getToken();
  },

  /**
   * Logout and optionally redirect
   */
  logout(redirectTo = '/') {
    this.removeToken();
    if (redirectTo) {
      window.location.href = redirectTo;
    }
  },

  /**
   * Get authorization headers for API calls
   */
  getAuthHeaders() {
    const token = this.getToken();
    const headers = { 'Content-Type': 'application/json' };
    if (token) {
      headers['Authorization'] = `Bearer ${token}`;
    }
    return headers;
  },

  /**
   * Handle 401 unauthorized response
   */
  handleUnauthorized() {
    this.removeToken();
    window.location.href = '/signin';
  },

  /**
   * Require authentication - redirect to signin if not logged in
   */
  requireAuth(redirectTo = '/signin') {
    if (!this.isLoggedIn()) {
      window.location.href = redirectTo;
      return false;
    }
    return true;
  },

  /**
   * Redirect if already logged in (for signin/signup pages)
   */
  redirectIfLoggedIn(redirectTo = '/') {
    if (this.isLoggedIn()) {
      window.location.href = redirectTo;
      return true;
    }
    return false;
  }
};

export default auth;
