const SESSION_KEY = 'mobile_upgrade_session';
const MAX_AGE_MS = 24 * 60 * 60 * 1000;

export function saveSession(state) {
  try {
    sessionStorage.setItem(SESSION_KEY, JSON.stringify({
      ...state,
      savedAt: Date.now(),
    }));
  } catch (e) {
    // sessionStorage may be unavailable
  }
}

export function loadSession() {
  try {
    const raw = sessionStorage.getItem(SESSION_KEY);
    if (!raw) {
      return null;
    }
    const data = JSON.parse(raw);
    if (!data.savedAt || Date.now() - data.savedAt > MAX_AGE_MS) {
      clearSession();
      return null;
    }
    return data;
  } catch (e) {
    return null;
  }
}

export function clearSession() {
  try {
    sessionStorage.removeItem(SESSION_KEY);
  } catch (e) {
    // ignore
  }
}
