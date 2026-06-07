<script lang="ts">
	import { getResolvedTheme, toggleTheme } from '$lib/theme/theme.svelte';

	let isDark = $derived(getResolvedTheme() === 'dark');
	let label = $derived(isDark ? 'Switch to light mode' : 'Switch to dark mode');
</script>

<button
	type="button"
	onclick={toggleTheme}
	aria-label={label}
	title={label}
	class="inline-flex h-9 w-9 shrink-0 items-center justify-center rounded-md border border-line bg-surface text-muted transition hover:border-accent hover:bg-accent-soft hover:text-accent"
>
	{#if isDark}
		<svg class="h-4 w-4" viewBox="0 0 24 24" fill="none" stroke="currentColor" aria-hidden="true">
			<path
				stroke-linecap="round"
				stroke-linejoin="round"
				stroke-width="2"
				d="M12 3v2m0 14v2m9-9h-2M5 12H3m15.364-6.364-1.414 1.414M7.05 16.95l-1.414 1.414m12.728 0-1.414-1.414M7.05 7.05 5.636 5.636"
			/>
			<circle cx="12" cy="12" r="4" stroke-width="2" />
		</svg>
	{:else}
		<svg class="h-4 w-4" viewBox="0 0 24 24" fill="none" stroke="currentColor" aria-hidden="true">
			<path
				stroke-linecap="round"
				stroke-linejoin="round"
				stroke-width="2"
				d="M20.354 15.354A9 9 0 0 1 8.646 3.646 9.003 9.003 0 1 0 20.354 15.354Z"
			/>
		</svg>
	{/if}
</button>
