<script lang="ts">
	import { resolve } from '$app/paths';
	import { scrollAnimation } from '$lib/actions/scrollAnimation';
	import type { PageData } from './$types';

	let { data }: { data: PageData } = $props();
	const project = $derived(data.project);
</script>

<svelte:head>
	<title>{project.title} — Tubagus Aldi</title>
	<meta name="description" content={project.description} />
	<meta property="og:title" content="{project.title} — Tubagus Aldi" />
	<meta property="og:description" content={project.description} />
</svelte:head>

<section class="bg-canvas py-12" use:scrollAnimation>
	<div class="mx-auto max-w-4xl px-4 sm:px-6 lg:px-8">
		<!-- Breadcrumb -->
		<nav class="mb-8 font-mono text-sm text-muted">
			<a href={resolve('/projects')} class="transition hover:text-accent">TubagusAldiMY</a>
			<span class="mx-2 opacity-50">/</span>
			<span class="font-semibold text-accent">{project.title}</span>
			<span class="ml-3 rounded-full border border-line px-2 py-0.5 text-xs text-muted">Public</span
			>
		</nav>

		<!-- Header -->
		<header class="mb-10 border-b border-line pb-6">
			<h1 class="mb-3 text-3xl font-bold text-ink">{project.title}</h1>
			<p class="max-w-2xl text-base leading-relaxed text-muted">{project.description}</p>

			<div class="mt-5 flex flex-wrap gap-3">
				{#if project.repoUrl}
					<a
						href={project.repoUrl}
						target="_blank"
						rel="noreferrer external"
						class="flex items-center gap-2 rounded-md border border-line bg-surface-alt px-4 py-1.5 text-sm font-semibold text-body transition hover:border-accent hover:bg-accent-soft hover:text-accent"
					>
						<svg aria-hidden="true" height="16" viewBox="0 0 16 16" width="16" fill="currentColor">
							<path
								d="M8 0c4.42 0 8 3.58 8 8a8.013 8.013 0 0 1-5.45 7.59c-.4.08-.55-.17-.55-.38 0-.27.01-1.13.01-2.2 0-.75-.25-1.23-.54-1.48 1.78-.2 3.65-.88 3.65-3.95 0-.88-.31-1.59-.82-2.15.08-.2.36-1.02-.08-2.12 0 0-.67-.22-2.2.82-.64-.18-1.32-.27-2-.27-.68 0-1.36.09-2 .27-1.53-1.03-2.2-.82-2.2-.82-.44 1.1-.16 1.92-.08 2.12-.51.56-.82 1.28-.82 2.15 0 3.06 1.86 3.75 3.64 3.95-.23.2-.44.55-.51 1.07-.46.21-1.61.55-2.33-.66-.15-.24-.6-.83-1.23-.82-.67.01-.27.38.01.53.34.19.73.9.82 1.13.16.45.68 1.31 2.69.94 0 .67.01 1.3.01 1.49 0 .21-.15.45-.55.38A7.995 7.995 0 0 1 0 8c0-4.42 3.58-8 8-8Z"
							/>
						</svg>
						View on GitHub ↗
					</a>
				{/if}
				{#if project.liveUrl}
					<a
						href={project.liveUrl}
						target="_blank"
						rel="noreferrer external"
						class="rounded-md border border-success/40 bg-success px-4 py-1.5 text-sm font-semibold text-inverse transition hover:bg-success-hover"
					>
						Live Demo ↗
					</a>
				{/if}
			</div>
		</header>

		<!-- Tech stack -->
		{#if project.techStack.length > 0}
			<div class="mb-8">
				<h2 class="mb-3 text-xs font-semibold uppercase tracking-widest text-muted">
					Languages & Tools
				</h2>
				<div class="flex flex-wrap gap-2">
					{#each project.techStack as tech (tech)}
						<span
							class="flex items-center gap-1.5 rounded-full border border-line bg-surface px-3 py-1 text-sm text-body"
						>
							<span class="h-3 w-3 rounded-full bg-[#3572A5]"></span>
							{tech}
						</span>
					{/each}
				</div>
			</div>
		{/if}

		<!-- Image -->
		{#if project.imageUrl}
			<div class="mb-8">
				<img
					src={project.imageUrl}
					alt="{project.title} screenshot"
					class="w-full rounded-md border border-line shadow-lg"
					loading="lazy"
				/>
			</div>
		{/if}
	</div>
</section>
