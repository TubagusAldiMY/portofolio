import { PUBLIC_API_URL } from '$env/static/public';
import { clearAuthToken, getAuthToken } from '$lib/auth/auth.svelte';

type JsonBody = Record<string, unknown> | unknown[];

type ApiFetchOptions = Omit<RequestInit, 'body'> & {
	auth?: boolean;
	body?: BodyInit | JsonBody | null;
	fetcher?: typeof fetch;
	token?: string | null;
};

export const API_BASE = PUBLIC_API_URL.replace(/\/+$/, '');

export function apiUrl(path: string): string {
	const normalizedPath = path.startsWith('/') ? path : `/${path}`;
	return `${API_BASE}${normalizedPath}`;
}

export class ApiError extends Error {
	status: number;
	payload: unknown;

	constructor(status: number, message: string, payload: unknown) {
		super(message);
		this.name = 'ApiError';
		this.status = status;
		this.payload = payload;
	}
}

function isFormData(value: unknown): value is FormData {
	return typeof FormData !== 'undefined' && value instanceof FormData;
}

function isBodyInit(value: unknown): value is BodyInit {
	return (
		typeof value === 'string' ||
		(typeof Blob !== 'undefined' && value instanceof Blob) ||
		value instanceof ArrayBuffer ||
		value instanceof URLSearchParams ||
		isFormData(value)
	);
}

function resolveBody(body: ApiFetchOptions['body'], headers: Headers): BodyInit | null | undefined {
	if (body == null) {
		return body;
	}

	if (isBodyInit(body)) {
		return body;
	}

	if (!headers.has('Content-Type')) {
		headers.set('Content-Type', 'application/json');
	}

	return JSON.stringify(body);
}

function readErrorMessage(payload: unknown, fallback: string): string {
	if (payload && typeof payload === 'object' && 'error' in payload) {
		const error = (payload as { error?: unknown }).error;
		if (typeof error === 'string' && error.length > 0) {
			return error;
		}
	}

	if (payload && typeof payload === 'object' && 'message' in payload) {
		const message = (payload as { message?: unknown }).message;
		if (typeof message === 'string' && message.length > 0) {
			return message;
		}
	}

	return fallback;
}

async function parseResponse(response: Response): Promise<unknown> {
	const contentType = response.headers.get('content-type') ?? '';

	if (contentType.includes('application/json')) {
		return response.json();
	}

	const text = await response.text();
	return text.length > 0 ? text : null;
}

export async function apiFetch<T>(path: string, options: ApiFetchOptions = {}): Promise<T> {
	const { auth = true, body, fetcher = fetch, token, ...init } = options;
	const headers = new Headers(init.headers);
	const authToken = token ?? getAuthToken();

	if (auth && authToken && !headers.has('Authorization')) {
		headers.set('Authorization', `Bearer ${authToken}`);
	}

	const response = await fetcher(apiUrl(path), {
		...init,
		body: resolveBody(body, headers),
		headers
	});
	const payload = await parseResponse(response);

	if (!response.ok) {
		if (response.status === 401) {
			clearAuthToken();
		}

		throw new ApiError(response.status, readErrorMessage(payload, response.statusText), payload);
	}

	return payload as T;
}
