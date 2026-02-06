/**
 * API Service - Centralized API calls to backend
 */

import { auth } from './auth.js';

const API_BASE = '/api';

/**
 * Make an API request
 */
async function request(endpoint, options = {}) {
  const url = endpoint.startsWith('/') ? endpoint : `${API_BASE}/${endpoint}`;

  const config = {
    headers: options.authenticated !== false ? auth.getAuthHeaders() : { 'Content-Type': 'application/json' },
    ...options
  };

  const response = await fetch(url, config);

  // Handle unauthorized
  if (response.status === 401) {
    auth.handleUnauthorized();
    throw new Error('Unauthorized');
  }

  return response;
}

/**
 * Parse JSON response with error handling
 */
async function parseJSON(response) {
  const text = await response.text();
  try {
    return JSON.parse(text);
  } catch {
    return text;
  }
}

export const api = {
  // ==================== AUTH ====================

  /**
   * Sign in user
   */
  async signin(email, password) {
    const response = await request('/api/auth/signin', {
      method: 'POST',
      authenticated: false,
      body: JSON.stringify({ email, password })
    });

    const result = await parseJSON(response);

    if (!response.ok) {
      throw new Error(result.error || 'Sign in failed');
    }

    return result;
  },

  /**
   * Sign up new user
   */
  async signup(name, email, password) {
    const response = await request('/api/auth/signup', {
      method: 'POST',
      authenticated: false,
      body: JSON.stringify({ name, email, password })
    });

    const result = await parseJSON(response);

    if (!response.ok) {
      throw new Error(result.error || 'Registration failed');
    }

    return result;
  },

  // ==================== PROFILE ====================

  /**
   * Get current user profile
   */
  async getProfile() {
    const response = await request('/api/profile');

    if (!response.ok) {
      throw new Error('Failed to fetch profile');
    }

    return response.json();
  },

  // ==================== TASKS ====================

  /**
   * Get tasks list with filters
   */
  async getTasks({ search = '', difficulty = '', community = '', page = 1, pageSize = 20 } = {}) {
    const params = new URLSearchParams({ search, difficulty, community, page, pageSize });
    const response = await request(`/api/tasks?${params}`, {
      authenticated: false
    });

    if (!response.ok) {
      throw new Error(`Server error: ${response.status}`);
    }

    const tasks = await response.json();
    const totalCount = parseInt(response.headers.get('X-Total-Count') || tasks.length);

    return { tasks, totalCount };
  },

  /**
   * Get single task by ID
   */
  async getTask(taskId) {
    const response = await request(`/api/tasks/${taskId}`, {
      authenticated: false
    });

    if (!response.ok) {
      throw new Error('Failed to load task');
    }

    return response.json();
  },

  // ==================== CODE EXECUTION ====================

  /**
   * Run code with input
   */
  async runCode({ language, code, input, taskId }) {
    const response = await request('/api/run', {
      method: 'POST',
      body: JSON.stringify({ language, code, input, taskId })
    });

    if (response.status === 429) {
      throw new Error('Rate limit exceeded. Please wait before making more requests.');
    }

    if (!response.ok) {
      const text = await response.text();
      throw new Error(text || 'Run failed');
    }

    return response.json();
  },

  /**
   * Submit code for a task
   */
  async submitCode(taskId, { language, code }) {
    const response = await request(`/api/submit/${taskId}`, {
      method: 'POST',
      body: JSON.stringify({ language, code })
    });

    if (response.status === 429) {
      throw new Error('Rate limit exceeded. Please wait before making more requests.');
    }

    if (!response.ok) {
      const text = await response.text();
      throw new Error(text || 'Submit failed');
    }

    return response.json();
  }
};

export default api;
