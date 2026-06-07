export const AUTH_STORAGE_KEY = 'token';
export const AUTH_COOKIE_NAME = 'portfolio_admin_token';

type JwtPayload = {
	exp?: number;
	sub?: string | number;
	[key: string]: unknown;
};

function decodeBase64Url(input: string): string {
	const normalized = input.replace(/-/g, '+').replace(/_/g, '/');
	const padded = normalized.padEnd(Math.ceil(normalized.length / 4) * 4, '=');
	const decoded = globalThis.atob(padded);
	const bytes = Uint8Array.from(decoded, (char) => char.charCodeAt(0));

	return new TextDecoder().decode(bytes);
}

export function parseJwtPayload(token: string): JwtPayload | null {
	const [, payload] = token.split('.');

	if (!payload) {
		return null;
	}

	try {
		return JSON.parse(decodeBase64Url(payload)) as JwtPayload;
	} catch {
		return null;
	}
}

export function isJwtExpired(token: string, skewSeconds = 30): boolean {
	const payload = parseJwtPayload(token);

	if (typeof payload?.exp !== 'number') {
		return true;
	}

	return Date.now() / 1000 >= payload.exp - skewSeconds;
}
