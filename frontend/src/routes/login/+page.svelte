<script lang="ts">
	import { goto } from '$app/navigation';
	import { resolve } from '$app/paths';
	import { page } from '$app/state';
	import type { Pathname } from '$app/types';

	import { ApiError } from '$lib/api/client';
	import { login } from '$lib/api/public';
	import { setAuthToken } from '$lib/auth/auth.svelte';

	let username = $state('');
	let password = $state('');
	let isSubmitting = $state(false);
	let error = $state<string | null>(null);

	function getRedirectTarget(): Pathname {
		const redirectTo = page.url.searchParams.get('redirectTo');
		// prefix check: hanya izinkan path internal di bawah /admin
		// mencegah open-redirect ke domain eksternal (https://, //)
		if (redirectTo && redirectTo.startsWith('/admin')) {
			return redirectTo as Pathname;
		}
		return '/admin';
	}

	async function submitLogin(event: SubmitEvent): Promise<void> {
		event.preventDefault();
		isSubmitting = true;
		error = null;

		try {
			const response = await login({ username, password });
			setAuthToken(response.token);
			await goto(resolve(getRedirectTarget()));
		} catch (err) {
			error = err instanceof ApiError ? err.message : 'Login failed.';
		} finally {
			isSubmitting = false;
		}
	}
</script>

<svelte:head>
	<meta name="robots" content="noindex, nofollow" />
	<title>Admin Login — Tubagus Aldi</title>
</svelte:head>

<section class="flex min-h-[calc(100vh-8rem)] items-center justify-center px-4 py-12">
	<form
		class="w-full max-w-md rounded-lg border border-line bg-surface p-6 shadow-[var(--theme-shadow-soft)]"
		onsubmit={submitLogin}
	>
		<div class="mb-6">
			<p class="text-sm font-semibold uppercase tracking-[0.2em] text-success">Admin</p>
			<h1 class="mt-2 text-2xl font-semibold text-ink">Sign in</h1>
		</div>

		{#if error}
			<p
				class="mb-4 rounded-md border border-danger/30 bg-danger-soft px-4 py-3 text-sm text-danger"
			>
				{error}
			</p>
		{/if}

		<div class="grid gap-4">
			<label class="grid gap-2 text-sm font-medium text-body">
				Username
				<input
					bind:value={username}
					required
					autocomplete="username"
					class="rounded-md border border-line bg-canvas px-3 py-2 text-ink outline-none transition focus:border-accent focus:ring-1 focus:ring-accent"
				/>
			</label>
			<label class="grid gap-2 text-sm font-medium text-body">
				Password
				<input
					bind:value={password}
					required
					type="password"
					autocomplete="current-password"
					class="rounded-md border border-line bg-canvas px-3 py-2 text-ink outline-none transition focus:border-accent focus:ring-1 focus:ring-accent"
				/>
			</label>
			<button
				type="submit"
				disabled={isSubmitting}
				class="rounded-md bg-success px-4 py-2 text-sm font-semibold text-inverse hover:bg-success-hover disabled:cursor-not-allowed disabled:opacity-60"
			>
				{isSubmitting ? 'Signing in...' : 'Sign In'}
			</button>
		</div>
	</form>
</section>
