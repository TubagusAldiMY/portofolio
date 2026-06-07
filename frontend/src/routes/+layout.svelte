<script lang="ts">
	import { onMount } from 'svelte';
	import { page } from '$app/state';

	import '../app.css';
	import favicon from '$lib/assets/favicon.svg';
	import ChatWidget from '$lib/components/ChatWidget.svelte';
	import Footer from '$lib/components/Footer.svelte';
	import Navbar from '$lib/components/Navbar.svelte';
	import { initializeAuth } from '$lib/auth/auth.svelte';
	import { initializeTheme } from '$lib/theme/theme.svelte';

	let { children } = $props();

	onMount(() => {
		initializeTheme();
		initializeAuth();
	});
</script>

<svelte:head>
	<link rel="icon" href={favicon} />
	<meta name="author" content="Tubagus Aldi Maulana Yusuf" />
	<meta property="og:site_name" content="Tubagus Aldi" />
	<meta property="og:type" content="website" />
	<meta property="og:image" content="/tubsamy.webp" />
	<meta name="twitter:card" content="summary_large_image" />
	<meta name="twitter:image" content="/tubsamy.webp" />
	<meta name="theme-color" content="#0d1117" />
</svelte:head>

<div class="flex min-h-dvh flex-col bg-canvas text-body">
	<Navbar pathname={page.url.pathname} />
	<main class="flex-1">
		{@render children()}
	</main>
	<Footer />
</div>

<ChatWidget />
