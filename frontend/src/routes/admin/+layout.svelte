<script lang="ts">
	import { onMount } from 'svelte';
	import { goto } from '$app/navigation';
	import { resolve } from '$app/paths';

	import { initializeAuth, isAuthenticated } from '$lib/auth/auth.svelte';

	let { children } = $props();

	onMount(() => {
		initializeAuth();

		if (!isAuthenticated()) {
			void goto(resolve('/login'));
		}
	});
</script>

<svelte:head>
	<meta name="robots" content="noindex, nofollow" />
</svelte:head>

{@render children()}
