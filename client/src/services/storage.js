/**
 * Storage Service - LocalStorage helpers
 */

const STORAGE_KEYS = {
  TOKEN: 'token',  // Keep as 'token' for backward compatibility with backend
  LANGUAGE: 'stillcode_language',
  CODE_PREFIX: 'stillcode_code_',
  HISTORY: 'stillcode_history_v1'
};

export const storage = {
  /**
   * Get item from localStorage
   */
  get(key) {
    try {
      return localStorage.getItem(key);
    } catch (e) {
      console.warn('Storage get failed:', e);
      return null;
    }
  },

  /**
   * Set item in localStorage
   */
  set(key, value) {
    try {
      localStorage.setItem(key, value);
      return true;
    } catch (e) {
      console.warn('Storage set failed:', e);
      return false;
    }
  },

  /**
   * Remove item from localStorage
   */
  remove(key) {
    try {
      localStorage.removeItem(key);
      return true;
    } catch (e) {
      console.warn('Storage remove failed:', e);
      return false;
    }
  },

  /**
   * Get JSON parsed item
   */
  getJSON(key) {
    try {
      const raw = localStorage.getItem(key);
      return raw ? JSON.parse(raw) : null;
    } catch (e) {
      console.warn('Storage getJSON failed:', e);
      return null;
    }
  },

  /**
   * Set JSON stringified item
   */
  setJSON(key, value) {
    try {
      localStorage.setItem(key, JSON.stringify(value));
      return true;
    } catch (e) {
      console.warn('Storage setJSON failed:', e);
      return false;
    }
  },

  /**
   * Get saved code for a task
   */
  getCode(taskId) {
    return this.get(STORAGE_KEYS.CODE_PREFIX + taskId) || '';
  },

  /**
   * Save code for a task
   */
  setCode(taskId, code) {
    return this.set(STORAGE_KEYS.CODE_PREFIX + taskId, code);
  },

  /**
   * Get saved language preference
   */
  getLanguage() {
    return this.get(STORAGE_KEYS.LANGUAGE) || 'Python';
  },

  /**
   * Save language preference
   */
  setLanguage(language) {
    return this.set(STORAGE_KEYS.LANGUAGE, language);
  },

  /**
   * Get submission history
   */
  getHistory() {
    return this.getJSON(STORAGE_KEYS.HISTORY) || [];
  },

  /**
   * Add item to history (keeps last 50 items)
   */
  addToHistory(item) {
    try {
      const arr = this.getHistory();
      arr.unshift(item);
      if (arr.length > 50) arr.pop();
      return this.setJSON(STORAGE_KEYS.HISTORY, arr);
    } catch (e) {
      console.warn('History save failed:', e);
      return false;
    }
  }
};

export { STORAGE_KEYS };
export default storage;
