import { browser } from '$app/environment';

export type ThemePreference = 'system' | 'light' | 'dark';
export type ResolvedTheme = 'light' | 'dark';

const STORAGE_KEY = 'portfolio-theme-preference';
const THEME_COLORS: Record<ResolvedTheme, string> = {
	dark: '#0d1117',
	light: '#ffffff'
};

let preference = $state<ThemePreference>('dark');
let resolvedTheme = $state<ResolvedTheme>('dark');
let initialized = false;
let mediaQuery: MediaQueryList | null = null;

function normalizePreference(value: string | null | undefined): ThemePreference {
	return value === 'light' || value === 'dark' || value === 'system' ? value : 'dark';
}

function systemTheme(): ResolvedTheme {
	if (!browser) return 'dark';
	return window.matchMedia('(prefers-color-scheme: dark)').matches ? 'dark' : 'light';
}

function resolveTheme(nextPreference: ThemePreference): ResolvedTheme {
	return nextPreference === 'system' ? systemTheme() : nextPreference;
}

function updateThemeColor(theme: ResolvedTheme): void {
	const tag = document.querySelector<HTMLMetaElement>('meta[name="theme-color"]');
	if (tag) tag.content = THEME_COLORS[theme];
}

function applyTheme(nextPreference = preference): void {
	if (!browser) return;

	preference = nextPreference;
	resolvedTheme = resolveTheme(nextPreference);

	const root = document.documentElement;
	root.dataset.theme = resolvedTheme;
	root.dataset.themePreference = preference;
	root.style.colorScheme = resolvedTheme;
	updateThemeColor(resolvedTheme);
}

export function initializeTheme(): void {
	if (!browser || initialized) return;

	initialized = true;
	preference = normalizePreference(
		localStorage.getItem(STORAGE_KEY) ?? document.documentElement.dataset.themePreference
	);
	applyTheme(preference);

	mediaQuery = window.matchMedia('(prefers-color-scheme: dark)');
	mediaQuery.addEventListener('change', () => {
		if (preference === 'system') applyTheme('system');
	});
}

export function getThemePreference(): ThemePreference {
	return preference;
}

export function getResolvedTheme(): ResolvedTheme {
	return resolvedTheme;
}

export function setThemePreference(nextPreference: ThemePreference): void {
	preference = nextPreference;

	if (browser) {
		localStorage.setItem(STORAGE_KEY, nextPreference);
	}

	applyTheme(nextPreference);
}

export function toggleTheme(): void {
	setThemePreference(resolvedTheme === 'dark' ? 'light' : 'dark');
}
