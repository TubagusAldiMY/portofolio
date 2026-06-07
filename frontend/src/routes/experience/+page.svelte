<script lang="ts">
	import { scrollAnimation } from '$lib/actions/scrollAnimation';
	import type { PageData } from './$types';

	let { data }: { data: PageData } = $props();
</script>

<svelte:head>
	<title>Contributions — Tubagus Aldi</title>
	<meta name="description" content="Professional experience and contributions by Tubagus Aldi." />
</svelte:head>

<section class="bg-canvas py-20 sm:py-32" use:scrollAnimation>
	<div class="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
		<div class="mb-16 text-center">
			<h2 class="text-3xl font-extrabold text-ink sm:text-4xl">Contribution Activity</h2>
			<p class="mt-4 text-lg text-muted">Riwayat pengalaman dan kontribusi profesional saya.</p>
		</div>

		{#if data.error}
			<p
				class="rounded-md border border-danger/30 bg-danger-soft px-4 py-3 text-center text-sm text-danger"
			>
				{data.error}
			</p>
		{:else if data.experiences.length === 0}
			<p class="text-center text-muted">Belum ada data pengalaman.</p>
		{:else}
			<div class="relative mx-auto ml-3 max-w-3xl space-y-12 border-l border-line md:ml-6">
				{#each data.experiences as exp (exp.id)}
					<div class="group relative pl-8 md:pl-12">
						<span
							class="absolute -left-[5px] top-0 mt-1.5 h-3 w-3 rounded-full border-2 border-accent bg-canvas ring-4 ring-canvas transition-colors group-hover:bg-accent"
						></span>

						<div class="mb-2 flex flex-col gap-1 sm:flex-row sm:items-center sm:justify-between">
							<h3 class="text-xl font-bold text-ink transition-colors group-hover:text-accent">
								{exp.role}
							</h3>
							<span
								class="w-fit rounded border border-line bg-surface px-2 py-1 font-mono text-xs text-muted"
							>
								{exp.duration}
							</span>
						</div>

						<p class="mb-4 text-base font-medium text-body">{exp.company}</p>

						<div
							class="relative rounded-md border border-line bg-surface p-4 transition-colors hover:border-muted"
						>
							<div
								class="absolute -left-1.5 top-4 h-3 w-3 rotate-45 border-b border-l border-line bg-surface"
							></div>
							<ul class="space-y-2">
								{#each exp.description.split('\n').filter(Boolean) as line (line)}
									<li class="flex items-start gap-2 text-sm text-muted">
										<svg
											class="mt-0.5 h-4 w-4 flex-shrink-0 text-success"
											fill="currentColor"
											viewBox="0 0 20 20"
										>
											<path
												fill-rule="evenodd"
												d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z"
												clip-rule="evenodd"
											/>
										</svg>
										{line}
									</li>
								{/each}
							</ul>
						</div>
					</div>
				{/each}
			</div>
		{/if}
	</div>
</section>
