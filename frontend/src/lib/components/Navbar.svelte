<script lang="ts">
	import { goto } from '$app/navigation';
	import { resolve } from '$app/paths';
	import type { Pathname } from '$app/types';

	import { clearAuthToken, isAuthenticated } from '$lib/auth/auth.svelte';
	import ThemeToggle from '$lib/components/ThemeToggle.svelte';

	type NavLink = { href: Pathname; label: string };

	const navLinks: NavLink[] = [
		{ href: '/', label: 'Overview' },
		{ href: '/projects', label: 'Work' },
		{ href: '/products', label: 'Products' },
		{ href: '/experience', label: 'Experience' },
		{ href: '/chat', label: 'AI Chat' }
	];

	let { pathname }: { pathname: string } = $props();
	let isOpen = $state(false);

	function isActive(href: string): boolean {
		return href === '/' ? pathname === '/' : pathname.startsWith(href);
	}

	function closeMenu() {
		isOpen = false;
	}

	async function logout() {
		clearAuthToken();
		closeMenu();
		await goto(resolve('/'));
	}
</script>

<nav class="sticky top-0 z-50 border-b border-line bg-canvas/90 backdrop-blur-md">
	<div class="mx-auto flex h-16 max-w-7xl items-center justify-between px-4 sm:px-6 lg:px-8">
		<a
			href={resolve('/')}
			onclick={closeMenu}
			class="flex flex-shrink-0 items-center gap-2 text-xl font-bold tracking-tight text-ink transition-colors hover:text-body"
		>
			<span
				class="h-2.5 w-2.5 rounded-full bg-success shadow-[0_0_12px_color-mix(in_srgb,var(--theme-success)_70%,transparent)]"
			></span>
			TubsAMY
		</a>

		<div class="hidden items-center gap-1 md:flex">
			{#each navLinks as link (link.href)}
				<a
					href={resolve(link.href)}
					class="rounded-md px-3 py-2 text-sm font-medium transition-all duration-200 {isActive(
						link.href
					)
						? 'text-ink'
						: 'text-body opacity-80 hover:text-ink hover:opacity-100'}"
				>
					{link.label}
				</a>
			{/each}
		</div>

		<div class="hidden items-center gap-2 md:flex">
			{#if isAuthenticated()}
				<a
					href={resolve('/admin')}
					class="px-3 py-1.5 text-sm font-medium text-body opacity-80 transition-all hover:text-ink hover:opacity-100"
				>
					Admin
				</a>
				<button
					onclick={logout}
					class="rounded-md border border-line bg-surface-alt px-4 py-1.5 text-sm font-semibold text-ink transition-all hover:border-accent hover:bg-accent-soft hover:text-accent"
				>
					Sign out
				</button>
			{:else}
				<a
					href={resolve('/contact')}
					class="rounded-md border border-success/40 bg-success px-4 py-1.5 text-sm font-semibold text-inverse shadow-sm transition-all duration-200 hover:bg-success-hover"
				>
					Mulai Diskusi
				</a>
			{/if}
			<ThemeToggle />
		</div>

		<button
			type="button"
			aria-expanded={isOpen}
			aria-controls="mobile-nav"
			onclick={() => (isOpen = !isOpen)}
			class="rounded-md p-2 text-body hover:text-ink md:hidden"
		>
			<svg class="h-5 w-5" stroke="currentColor" fill="none" viewBox="0 0 24 24">
				{#if isOpen}
					<path
						stroke-linecap="round"
						stroke-linejoin="round"
						stroke-width="2"
						d="M6 18L18 6M6 6l12 12"
					/>
				{:else}
					<path
						stroke-linecap="round"
						stroke-linejoin="round"
						stroke-width="2"
						d="M4 6h16M4 12h16M4 18h16"
					/>
				{/if}
			</svg>
		</button>
	</div>

	{#if isOpen}
		<div id="mobile-nav" class="border-t border-line bg-surface px-4 pb-4 pt-2 md:hidden">
			<div class="space-y-1">
				{#each navLinks as link (link.href)}
					<a
						href={resolve(link.href)}
						onclick={closeMenu}
						class="block rounded-md px-3 py-2 text-sm font-medium transition-colors {isActive(
							link.href
						)
							? 'bg-accent-soft text-accent'
							: 'text-body hover:bg-surface-alt hover:text-ink'}"
					>
						{link.label}
					</a>
				{/each}
				<div class="flex items-center justify-between px-3 py-2">
					<span class="text-sm font-medium text-body">Theme</span>
					<ThemeToggle />
				</div>
				{#if isAuthenticated()}
					<a
						href={resolve('/admin')}
						onclick={closeMenu}
						class="block rounded-md px-3 py-2 text-sm font-medium text-body hover:bg-surface-alt hover:text-ink"
					>
						Admin
					</a>
					<button
						onclick={logout}
						class="block w-full rounded-md px-3 py-2 text-left text-sm font-medium text-body hover:bg-surface-alt hover:text-ink"
					>
						Sign out
					</button>
				{:else}
					<a
						href={resolve('/contact')}
						onclick={closeMenu}
						class="block rounded-md bg-success px-3 py-2 text-sm font-semibold text-inverse hover:bg-success-hover"
					>
						Mulai Diskusi
					</a>
				{/if}
			</div>
		</div>
	{/if}
</nav>
