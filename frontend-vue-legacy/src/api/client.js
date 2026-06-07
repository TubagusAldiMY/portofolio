// Lightweight fetch wrapper with base URL and token injection
export const API_BASE = import.meta.env.VITE_API_URL || '';

function isFormData(body) {
  return typeof FormData !== 'undefined' && body instanceof FormData;
}

export async function apiFetch(path, options = {}) {
  const token = typeof window !== 'undefined' ? localStorage.getItem('token') : null;
  const headers = new Headers(options.headers || {});

  if (options.body && !isFormData(options.body) && !headers.has('Content-Type')) {
    headers.set('Content-Type', 'application/json');
  }
  if (token && !headers.has('Authorization')) {
    headers.set('Authorization', `Bearer ${token}`);
  }

  const res = await fetch(`${API_BASE}${path}`, { ...options, headers });
  return res;
}
