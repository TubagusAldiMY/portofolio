import { browser } from '$app/environment';

import { AUTH_COOKIE_NAME, AUTH_STORAGE_KEY, isJwtExpired } from './token';

const COOKIE_MAX_AGE_SECONDS = 60 * 60 * 24;

let token = $state<string | null>(null);
let initialized = $state(false);

function secureCookieAttribute(): string {
	return browser && window.location.protocol === 'https:' ? '; secure' : '';
}

function setCookie(nextToken: string): void {
	document.cookie = `${AUTH_COOKIE_NAME}=${encodeURIComponent(nextToken)}; path=/; max-age=${COOKIE_MAX_AGE_SECONDS}; samesite=lax${secureCookieAttribute()}`;
}

function clearCookie(): void {
	document.cookie = `${AUTH_COOKIE_NAME}=; path=/; max-age=0; samesite=lax${secureCookieAttribute()}`;
}

export function initializeAuth(): void {
	if (!browser || initialized) {
		return;
	}

	const storedToken = localStorage.getItem(AUTH_STORAGE_KEY);
	initialized = true;

	if (!storedToken || isJwtExpired(storedToken)) {
		clearAuthToken();
		return;
	}

	token = storedToken;
	setCookie(storedToken);
}

export function getAuthToken(): string | null {
	return token;
}

export function isAuthenticated(): boolean {
	return Boolean(token && !isJwtExpired(token));
}

export function setAuthToken(nextToken: string): void {
	token = nextToken;
	initialized = true;

	if (!browser) {
		return;
	}

	localStorage.setItem(AUTH_STORAGE_KEY, nextToken);
	setCookie(nextToken);
}

export function clearAuthToken(): void {
	token = null;
	initialized = true;

	if (!browser) {
		return;
	}

	localStorage.removeItem(AUTH_STORAGE_KEY);
	clearCookie();
}
