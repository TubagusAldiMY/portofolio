<script lang="ts">
	import { scrollAnimation } from '$lib/actions/scrollAnimation';
	import type { PageData } from './$types';

	let { data }: { data: PageData } = $props();
</script>

<svelte:head>
	<title>Products — Tubagus Aldi</title>
	<meta name="description" content="Digital products and tools by Tubagus Aldi." />
</svelte:head>

<section class="bg-canvas py-20" use:scrollAnimation>
	<div class="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
		<div class="mb-12">
			<h2 class="text-3xl font-bold text-ink md:text-4xl">Digital Products</h2>
			<p class="mt-4 text-lg text-muted">Tools, templates, dan solusi siap pakai.</p>
		</div>

		{#if data.error}
			<p class="rounded-md border border-danger/30 bg-danger-soft px-4 py-3 text-sm text-danger">
				{data.error}
			</p>
		{:else if data.products.length === 0}
			<p class="text-center text-muted">Belum ada produk yang tersedia.</p>
		{:else}
			<div class="grid grid-cols-1 gap-6 md:grid-cols-2">
				{#each data.products as product (product.id)}
					<article
						class="group flex flex-col justify-between overflow-hidden rounded-md border border-line bg-surface transition-colors hover:border-muted"
					>
						<!-- Card header -->
						<div class="p-6">
							<div class="mb-4 flex items-start justify-between gap-3">
								<div class="min-w-0">
									<div class="flex items-center gap-2">
										<svg
											aria-hidden="true"
											height="16"
											viewBox="0 0 16 16"
											width="16"
											fill="currentColor"
											class="flex-shrink-0 text-muted"
										>
											<path
												d="M2 2.5A2.5 2.5 0 0 1 4.5 0h8.75a.75.75 0 0 1 .75.75v12.5a.75.75 0 0 1-.75.75h-2.5a.75.75 0 0 1 0-1.5h1.75v-2h-8a1 1 0 0 0-.714 1.7.75.75 0 1 1-1.072 1.05A2.495 2.495 0 0 1 2 11.5Zm10.5-1h-8a1 1 0 0 0-1 1v6.708A2.486 2.486 0 0 1 4.5 9h8Z"
											/>
										</svg>
										<h3 class="text-lg font-bold text-accent">{product.title}</h3>
									</div>
								</div>
								{#if product.tag}
									<span
										class="flex-shrink-0 rounded-full border border-accent/40 bg-accent-soft px-2.5 py-0.5 text-xs font-semibold text-accent"
									>
										{product.tag}
									</span>
								{/if}
							</div>

							<p class="mb-4 text-sm leading-relaxed text-muted">{product.description}</p>

							{#if product.features.length > 0}
								<ul class="mb-4 space-y-1.5">
									{#each product.features as feature (feature)}
										<li class="flex items-start gap-2 text-sm text-body">
											<svg
												class="mt-0.5 h-3.5 w-3.5 flex-shrink-0 text-success"
												fill="currentColor"
												viewBox="0 0 20 20"
											>
												<path
													fill-rule="evenodd"
													d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z"
													clip-rule="evenodd"
												/>
											</svg>
											{feature}
										</li>
									{/each}
								</ul>
							{/if}
						</div>

						<!-- Card footer: price + CTA -->
						<div class="flex items-center justify-between border-t border-line bg-canvas px-6 py-4">
							{#if product.price}
								<span class="text-xl font-bold text-ink">{product.price}</span>
							{:else}
								<span class="text-sm font-medium text-success">Free</span>
							{/if}

							{#if product.buyUrl}
								<a
									href={product.buyUrl}
									target="_blank"
									rel="noreferrer external"
									class="rounded-md border border-success/40 bg-success px-4 py-1.5 text-sm font-semibold text-inverse shadow-sm transition hover:bg-success-hover"
								>
									Get it ↗
								</a>
							{/if}
						</div>
					</article>
				{/each}
			</div>
		{/if}
	</div>
</section>
