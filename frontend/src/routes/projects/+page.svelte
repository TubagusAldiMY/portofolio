<script lang="ts">
	import { resolve } from '$app/paths';
	import { scrollAnimation } from '$lib/actions/scrollAnimation';
	import type { PageData } from './$types';

	let { data }: { data: PageData } = $props();
</script>

<svelte:head>
	<title>Repositories — Tubagus Aldi</title>
	<meta
		name="description"
		content="Projects by Tubagus Aldi — Go backends, SvelteKit frontends, Rust desktop apps, and AI integrations."
	/>
</svelte:head>

<section class="bg-canvas py-20" use:scrollAnimation>
	<div class="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
		<div class="mb-12">
			<h2 class="text-3xl font-bold text-ink md:text-4xl">Top Repositories</h2>
			<p class="mt-4 text-lg text-muted">Koleksi project yang saya kerjakan.</p>
		</div>

		{#if data.error}
			<p class="rounded-md border border-danger/30 bg-danger-soft px-4 py-3 text-sm text-danger">
				{data.error}
			</p>
		{:else if data.projects.length === 0}
			<p class="text-center text-muted">Belum ada proyek yang ditampilkan.</p>
		{:else}
			<div class="grid grid-cols-1 gap-4 md:grid-cols-2">
				{#each data.projects as project (project.id)}
					<div
						class="group flex flex-col justify-between rounded-md border border-line bg-surface p-6 transition-colors duration-200 hover:border-muted"
					>
						<div>
							<div class="mb-3 flex items-center justify-between gap-2">
								<div class="flex min-w-0 items-center gap-2">
									<svg
										aria-hidden="true"
										height="16"
										viewBox="0 0 16 16"
										width="16"
										fill="currentColor"
										class="mt-0.5 flex-shrink-0 text-muted"
									>
										<path
											d="M2 2.5A2.5 2.5 0 0 1 4.5 0h8.75a.75.75 0 0 1 .75.75v12.5a.75.75 0 0 1-.75.75h-2.5a.75.75 0 0 1 0-1.5h1.75v-2h-8a1 1 0 0 0-.714 1.7.75.75 0 1 1-1.072 1.05A2.495 2.495 0 0 1 2 11.5Zm10.5-1h-8a1 1 0 0 0-1 1v6.708A2.486 2.486 0 0 1 4.5 9h8Z"
										/>
									</svg>
									<a
										href={resolve(`/projects/${project.id}`)}
										class="truncate text-lg font-bold text-accent hover:underline"
									>
										{project.title}
									</a>
									<span
										class="flex-shrink-0 rounded-full border border-line px-2 py-0.5 text-xs font-medium text-muted"
									>
										Public
									</span>
								</div>
							</div>

							<p class="mb-5 line-clamp-3 text-sm leading-relaxed text-muted">
								{project.description}
							</p>
						</div>

						<div class="mt-auto flex flex-wrap items-center gap-x-4 gap-y-2 text-xs text-muted">
							{#each project.techStack.slice(0, 4) as tech (tech)}
								<div class="flex items-center gap-1">
									<span class="h-3 w-3 rounded-full bg-[#3572A5]"></span>
									<span>{tech}</span>
								</div>
							{/each}
							{#if project.techStack.length > 4}
								<span class="text-muted">+{project.techStack.length - 4} more</span>
							{/if}
							{#if project.liveUrl}
								<a
									href={project.liveUrl}
									target="_blank"
									rel="noreferrer external"
									class="ml-auto transition hover:text-accent"
								>
									View Demo ↗
								</a>
							{/if}
						</div>
					</div>
				{/each}
			</div>
		{/if}
	</div>
</section>
