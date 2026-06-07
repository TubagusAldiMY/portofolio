<script lang="ts">
	import { onDestroy, onMount } from 'svelte';
	import { resolve } from '$app/paths';
	import { scrollAnimation } from '$lib/actions/scrollAnimation';
	import CardSwap from '$lib/components/CardSwap.svelte';
	import CodeWindow from '$lib/components/CodeWindow.svelte';
	import PixelBlast from '$lib/components/PixelBlast.svelte';
	import TechMarquee from '$lib/components/TechMarquee.svelte';
	import { listProjects } from '$lib/api/public';
	import type { PageData } from './$types';

	type Token = { t: string; v: string };
	type CodeLine = { tokens: Token[] };
	type CodeCard = { filename: string; lines: CodeLine[] };
	type Pill = { label: string; value: string };
	type WorkflowStep = { step: string; title: string; description: string };

	let { data }: { data: PageData } = $props();
	let projectOverride = $state<PageData['projects'] | null>(null);
	let projectErrorOverride = $state<string | null | undefined>(undefined);

	const tokenColor: Record<string, string> = {
		keyword: '#ff7b72',
		string: '#a5d6ff',
		fn: '#d2a8ff',
		num: '#79c0ff',
		comment: '#8b949e',
		attr: '#ffa657',
		tag: '#7ee787',
		text: '#c9d1d9'
	};

	const cards: CodeCard[] = [
		{
			filename: 'api.go',
			lines: [
				{
					tokens: [
						{ t: 'keyword', v: 'func' },
						{ t: 'fn', v: ' MountRoutes' },
						{ t: 'text', v: '(router *gin.Engine) {' }
					]
				},
				{ tokens: [{ t: 'comment', v: '    // Secure, typed, production-minded APIs' }] },
				{
					tokens: [
						{ t: 'text', v: '    api := router.Group(' },
						{ t: 'string', v: '"/api/v1"' },
						{ t: 'text', v: ')' }
					]
				},
				{
					tokens: [
						{ t: 'text', v: '    api.Use(' },
						{ t: 'fn', v: 'RateLimit' },
						{ t: 'text', v: '(), ' },
						{ t: 'fn', v: 'AuthGuard' },
						{ t: 'text', v: '())' }
					]
				},
				{
					tokens: [
						{ t: 'text', v: '    api.GET(' },
						{ t: 'string', v: '"/projects/:id"' },
						{ t: 'text', v: ', handler.Show)' }
					]
				},
				{ tokens: [{ t: 'text', v: '}' }] }
			]
		},
		{
			filename: 'pipeline.py',
			lines: [
				{
					tokens: [
						{ t: 'keyword', v: 'from' },
						{ t: 'text', v: ' app.ai ' },
						{ t: 'keyword', v: 'import' },
						{ t: 'text', v: ' embed, rank' }
					]
				},
				{ tokens: [] },
				{
					tokens: [
						{ t: 'keyword', v: 'def' },
						{ t: 'fn', v: ' answer_question' },
						{ t: 'text', v: '(message):' }
					]
				},
				{ tokens: [{ t: 'comment', v: '    # Retrieval first, generation second' }] },
				{
					tokens: [
						{ t: 'text', v: '    context = rank(' },
						{ t: 'fn', v: 'embed' },
						{ t: 'text', v: '(message), top_k=' },
						{ t: 'num', v: '5' },
						{ t: 'text', v: ')' }
					]
				},
				{
					tokens: [
						{ t: 'keyword', v: '    return' },
						{ t: 'text', v: ' generate(message, context)' }
					]
				}
			]
		},
		{
			filename: 'Portfolio.svelte',
			lines: [
				{
					tokens: [
						{ t: 'tag', v: '<script' },
						{ t: 'attr', v: ' lang' },
						{ t: 'text', v: '=' },
						{ t: 'string', v: '"ts"' },
						{ t: 'tag', v: '>' }
					]
				},
				{
					tokens: [
						{ t: 'keyword', v: '  let' },
						{ t: 'text', v: ' intent = ' },
						{ t: 'fn', v: '$state' },
						{ t: 'text', v: '(' },
						{ t: 'string', v: '"ship"' },
						{ t: 'text', v: ')' }
					]
				},
				{ tokens: [{ t: 'tag', v: '</' + 'script>' }] },
				{ tokens: [] },
				{
					tokens: [
						{ t: 'tag', v: '<Hero' },
						{ t: 'attr', v: ' focus' },
						{ t: 'text', v: '=' },
						{ t: 'string', v: '"reliable systems"' },
						{ t: 'tag', v: ' />' }
					]
				}
			]
		}
	];

	const heroSignals: Pill[] = [
		{ label: 'Core stack', value: 'Go, SvelteKit, TypeScript' },
		{ label: 'Product focus', value: 'API, dashboard, AI workflow' },
		{ label: 'Delivery style', value: 'Secure, typed, maintainable' }
	];

	const identityStack = ['Go', 'SvelteKit', 'TypeScript', 'PostgreSQL', 'AI'];
	const identityHighlights = ['Secure APIs', 'Typed UI', 'AI Workflow'];

	const projects = $derived(projectOverride ?? data.projects);
	const projectError = $derived(
		projectErrorOverride === undefined ? data.error : projectErrorOverride
	);
	const selectedProjects = $derived(projects.slice(0, 3));

	const workflow: WorkflowStep[] = [
		{
			step: '01',
			title: 'Clarify the product surface',
			description:
				'I start from the user flow, data shape, and failure states so the implementation does not drift from the goal.'
		},
		{
			step: '02',
			title: 'Design the contract',
			description:
				'API responses, validation, auth behavior, and frontend states are made explicit before the UI grows.'
		},
		{
			step: '03',
			title: 'Build the usable path',
			description:
				'I prioritize the real workflow first, then improve motion, density, responsiveness, and admin ergonomics.'
		},
		{
			step: '04',
			title: 'Harden and verify',
			description:
				'Build checks, linting, security constraints, and runtime smoke tests are part of the finish, not an afterthought.'
		}
	];

	let cardWidth = $state(260);
	let cardHeight = $state(340);
	let cardDistance = $state(40);
	let cardVerticalDistance = $state(40);
	let showCardSwap = $state(false);
	let showPixelBlast = $state(false);

	function projectCategory(techStack: string[]): string {
		const normalized = techStack.map((tech) => tech.toLowerCase());

		if (normalized.some((tech) => tech.includes('tauri') || tech.includes('rust'))) {
			return 'Desktop product';
		}

		if (normalized.some((tech) => tech.includes('ai') || tech.includes('gemini'))) {
			return 'AI integration';
		}

		if (normalized.some((tech) => tech.includes('go') || tech.includes('gin'))) {
			return 'Backend system';
		}

		if (normalized.some((tech) => tech.includes('svelte') || tech.includes('react'))) {
			return 'Web application';
		}

		return 'Selected project';
	}

	function projectAccent(techStack: string[]): string {
		const normalized = techStack.map((tech) => tech.toLowerCase());

		if (normalized.some((tech) => tech.includes('svelte'))) {
			return '#ff3e00';
		}

		if (normalized.some((tech) => tech.includes('go'))) {
			return '#00add8';
		}

		if (normalized.some((tech) => tech.includes('rust'))) {
			return '#ce412b';
		}

		if (normalized.some((tech) => tech.includes('ai') || tech.includes('gemini'))) {
			return '#d2a8ff';
		}

		return 'var(--theme-accent)';
	}

	async function refreshSelectedProjects(): Promise<void> {
		if (projects.length > 0 && !projectError) {
			return;
		}

		try {
			projectOverride = await listProjects();
			projectErrorOverride = null;
		} catch {
			projectErrorOverride = 'Failed to load selected work from backend.';
		}
	}

	function updateHeroMedia(): void {
		const width = window.innerWidth;
		const reducedMotion = window.matchMedia('(prefers-reduced-motion: reduce)').matches;

		showCardSwap = width >= 800;
		showPixelBlast = width >= 768 && !reducedMotion;

		if (width < 640) {
			cardWidth = 240;
			cardHeight = 320;
			cardDistance = 30;
			cardVerticalDistance = 30;
		} else if (width < 1024) {
			cardWidth = 280;
			cardHeight = 360;
			cardDistance = 40;
			cardVerticalDistance = 40;
		} else if (width < 1440) {
			cardWidth = 340;
			cardHeight = 440;
			cardDistance = 50;
			cardVerticalDistance = 50;
		} else {
			cardWidth = 400;
			cardHeight = 520;
			cardDistance = 60;
			cardVerticalDistance = 60;
		}
	}

	onMount(() => {
		updateHeroMedia();
		window.addEventListener('resize', updateHeroMedia);
		void refreshSelectedProjects();
	});

	onDestroy(() => {
		if (typeof window !== 'undefined') {
			window.removeEventListener('resize', updateHeroMedia);
		}
	});
</script>

<svelte:head>
	<title>Tubagus Aldi — Backend, Web, and AI Engineer</title>
	<meta
		name="description"
		content="Portfolio of Tubagus Aldi Maulana Yusuf, a Go and SvelteKit engineer building reliable web systems, AI workflows, and polished product interfaces."
	/>
	<meta property="og:title" content="Tubagus Aldi — Backend, Web, and AI Engineer" />
	<meta
		property="og:description"
		content="Go backends, SvelteKit frontends, AI integrations, and production-minded portfolio work."
	/>
</svelte:head>

<section class="relative flex min-h-[calc(100dvh-4rem)] items-center overflow-hidden bg-canvas">
	{#if showPixelBlast}
		<div class="absolute inset-0 z-0">
			<PixelBlast
				variant="circle"
				pixelSize={6}
				color="#2e527d"
				patternScale={3}
				patternDensity={1.25}
				pixelSizeJitter={0.45}
				enableRipples={true}
				rippleSpeed={0.4}
				rippleThickness={0.12}
				rippleIntensityScale={1.45}
				speed={0.55}
				edgeFade={0.25}
				transparent={true}
			/>
		</div>
	{:else}
		<div
			class="pointer-events-none absolute inset-0 opacity-[0.035]"
			style="background-image: linear-gradient(var(--theme-grid-color) 1px, transparent 1px), linear-gradient(to right, var(--theme-grid-color) 1px, transparent 1px); background-size: 64px 64px;"
		></div>
	{/if}

	<div
		class="pointer-events-none absolute inset-0 z-0"
		style="background: var(--theme-hero-overlay)"
	></div>
	<div
		class="pointer-events-none absolute inset-x-0 bottom-0 z-0 h-40 bg-gradient-to-t from-canvas to-transparent"
	></div>

	<div class="relative z-10 mx-auto w-full max-w-7xl px-4 py-16 sm:px-6 sm:py-20 lg:px-8">
		<div class="grid items-center gap-12 lg:grid-cols-[minmax(0,1.05fr)_minmax(420px,0.95fr)]">
			<div class="max-w-3xl text-center lg:text-left">
				<div
					class="mb-5 inline-flex items-center gap-2 rounded-full border border-accent/30 bg-accent-soft px-3 py-1 text-xs font-semibold uppercase tracking-wide text-accent"
				>
					<span
						class="h-2 w-2 rounded-full bg-success"
						style="box-shadow: 0 0 12px color-mix(in srgb, var(--theme-success) 80%, transparent)"
					></span>
					Available for serious product work
				</div>

				<h1
					class="text-balance text-4xl font-extrabold leading-[1.05] tracking-tight text-ink drop-shadow-md sm:text-5xl lg:text-6xl"
				>
					Membangun sistem web yang
					<span class="text-gradient-purple">cepat, aman, dan AI-ready.</span>
				</h1>

				<p class="mx-auto mt-6 max-w-2xl text-base leading-8 text-body sm:text-lg lg:mx-0">
					Saya <strong class="text-ink">Tubagus Aldi</strong>, backend-first engineer yang
					menggabungkan Go, SvelteKit, dan AI integration untuk membuat aplikasi yang rapi, terukur,
					dan enak digunakan.
				</p>

				<div
					class="mt-8 flex flex-col items-center justify-center gap-3 sm:flex-row lg:justify-start"
				>
					<a
						href={resolve('/projects')}
						class="inline-flex items-center justify-center rounded-md border border-success/40 bg-success px-6 py-3 text-sm font-bold text-inverse shadow-sm transition-all duration-200 hover:bg-success-hover"
					>
						Lihat Karya Pilihan
					</a>
					<a
						href={resolve('/contact')}
						class="inline-flex items-center justify-center rounded-md border border-line bg-surface/80 px-6 py-3 text-sm font-bold text-ink backdrop-blur transition-all duration-200 hover:border-accent hover:bg-surface-alt"
					>
						Mulai Diskusi
					</a>
				</div>

				<div
					class="mt-8 flex flex-wrap justify-center gap-x-5 gap-y-2 text-sm text-muted lg:justify-start"
				>
					<a
						href="https://github.com/TubagusAldiMY"
						target="_blank"
						rel="noreferrer external"
						class="transition hover:text-accent">GitHub ↗</a
					>
					<a
						href="https://linkedin.com/in/tubagusaldi"
						target="_blank"
						rel="noreferrer external"
						class="transition hover:text-accent">LinkedIn ↗</a
					>
					<a href="mailto:admin@tubsamy.tech" class="transition hover:text-accent"
						>admin@tubsamy.tech</a
					>
				</div>

				<div class="mt-10 grid gap-3 sm:grid-cols-3">
					{#each heroSignals as signal (signal.label)}
						<div class="rounded-md border border-line bg-canvas/72 p-4 backdrop-blur">
							<p class="text-xs font-semibold uppercase tracking-wide text-muted">
								{signal.label}
							</p>
							<p class="mt-2 text-sm font-semibold leading-6 text-ink">{signal.value}</p>
						</div>
					{/each}
				</div>
			</div>

			<div class="flex w-full justify-center">
				<div
					class="relative flex h-[360px] w-full items-center justify-center sm:h-[420px] lg:h-[520px]"
				>
					<div
						class="absolute inset-x-8 bottom-0 h-px bg-gradient-to-r from-transparent via-accent/50 to-transparent"
					></div>
					{#if showCardSwap}
						<CardSwap
							width={cardWidth}
							height={cardHeight}
							{cardDistance}
							verticalDistance={cardVerticalDistance}
							skewAmount={2}
							delay={3200}
							pauseOnHover={false}
						>
							{#snippet card0()}
								<CodeWindow card={cards[0]} {tokenColor} />
							{/snippet}
							{#snippet card1()}
								<CodeWindow card={cards[1]} {tokenColor} />
							{/snippet}
							{#snippet card2()}
								<CodeWindow card={cards[2]} {tokenColor} />
							{/snippet}
						</CardSwap>
					{:else}
						<div class="w-full max-w-xs">
							<CodeWindow card={cards[0]} {tokenColor} />
						</div>
					{/if}
				</div>
			</div>
		</div>
	</div>
</section>

<section class="border-y border-line bg-canvas-subtle py-6" use:scrollAnimation>
	<div class="mx-auto grid max-w-7xl gap-3 px-4 sm:grid-cols-2 sm:px-6 lg:grid-cols-4 lg:px-8">
		{#each ['Backend architecture', 'Typed frontend', 'AI workflow', 'Security review'] as label (label)}
			<div class="flex items-center gap-3 text-sm text-muted">
				<span class="h-2 w-2 rounded-full bg-accent"></span>
				<span>{label}</span>
			</div>
		{/each}
	</div>
</section>

<section class="bg-canvas pt-20 sm:pt-24" use:scrollAnimation>
	<div class="mx-auto max-w-7xl px-4 pb-12 sm:px-6 lg:px-8">
		<div class="max-w-3xl">
			<p class="text-sm font-semibold uppercase tracking-wide text-accent">Engineering focus</p>
			<h2 class="mt-3 text-3xl font-extrabold tracking-tight text-ink sm:text-4xl">
				Backend architecture. Typed frontend. AI workflow. Security review.
			</h2>
			<p class="mt-4 text-base leading-7 text-muted">
				Tech stack yang saya pakai sehari-hari untuk membangun sistem produksi yang bisa diandalkan.
			</p>
		</div>
	</div>
</section>

<TechMarquee />

<!-- About -->
<section class="border-t border-line bg-canvas py-20 sm:py-32" use:scrollAnimation>
	<div class="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
		<div class="mb-16 text-center">
			<h2 class="text-3xl font-extrabold text-ink sm:text-4xl">Tentang Saya</h2>
			<div class="mx-auto mt-4 h-1 w-16 rounded-full bg-success"></div>
			<p class="mt-4 text-lg text-muted">
				Sedikit lebih dalam tentang perjalanan dan keahlian saya.
			</p>
		</div>
		<div class="grid items-center gap-12 md:grid-cols-2">
			<div class="flex justify-center">
				<div class="group relative w-full max-w-md">
					<div
						class="absolute -inset-1 rounded-xl bg-gradient-to-r from-accent to-[#bc8cff] opacity-20 blur transition duration-1000 group-hover:opacity-40"
					></div>
					<div
						class="relative overflow-hidden rounded-xl border border-line bg-surface p-6 shadow-[var(--theme-shadow-float)]"
					>
						<div class="flex items-start justify-between gap-4 border-b border-line pb-5">
							<div>
								<p class="text-xs font-semibold uppercase tracking-[0.22em] text-muted">
									Engineering Profile
								</p>
								<h3 class="mt-2 text-2xl font-extrabold text-ink">Tubagus Aldi</h3>
								<p class="mt-1 text-sm font-medium text-accent">Backend-first Engineer</p>
							</div>
							<div
								class="flex h-16 w-16 shrink-0 items-center justify-center rounded-lg border border-accent/30 bg-accent-soft font-mono text-2xl font-black text-accent"
								aria-hidden="true"
							>
								TA
							</div>
						</div>

						<div class="py-6">
							<div class="grid grid-cols-[1fr_auto_1fr_auto_1fr] items-center gap-2 text-center">
								<div class="rounded-md border border-line bg-canvas px-3 py-3">
									<p class="text-[10px] font-semibold uppercase tracking-wide text-muted">
										Frontend
									</p>
									<p class="mt-1 text-sm font-bold text-ink">SvelteKit</p>
								</div>
								<div class="h-px bg-line"></div>
								<div class="rounded-md border border-line bg-canvas px-3 py-3">
									<p class="text-[10px] font-semibold uppercase tracking-wide text-muted">API</p>
									<p class="mt-1 text-sm font-bold text-ink">Go</p>
								</div>
								<div class="h-px bg-line"></div>
								<div class="rounded-md border border-line bg-canvas px-3 py-3">
									<p class="text-[10px] font-semibold uppercase tracking-wide text-muted">Data</p>
									<p class="mt-1 text-sm font-bold text-ink">AI + DB</p>
								</div>
							</div>
						</div>

						<div class="flex flex-wrap gap-2">
							{#each identityStack as tech (tech)}
								<span
									class="rounded-full border border-line bg-surface-alt px-3 py-1 text-xs font-semibold text-body"
								>
									{tech}
								</span>
							{/each}
						</div>

						<div class="mt-6 grid gap-3 sm:grid-cols-3">
							{#each identityHighlights as item (item)}
								<div class="rounded-md border border-success/20 bg-success-soft px-3 py-3">
									<span class="block h-1.5 w-1.5 rounded-full bg-success"></span>
									<p class="mt-3 text-xs font-bold text-success">{item}</p>
								</div>
							{/each}
						</div>
					</div>
				</div>
			</div>
			<div>
				<h3 class="mb-4 text-2xl font-bold text-ink">Perjalanan Saya di Dunia Digital</h3>
				<div class="space-y-4 leading-relaxed text-body">
					<p>
						Saya adalah seorang <span class="font-semibold text-accent"
							>Machine Learning Engineer</span
						>
						dan <span class="font-semibold text-accent">Web Developer</span>. Ketertarikan saya
						dimulai dari rasa penasaran tentang bagaimana baris kode dapat memecahkan masalah nyata.
					</p>
					<p>
						Fokus utama saya adalah membangun backend yang tangguh menggunakan <strong
							class="text-ink">Go (Golang)</strong
						>
						dan frontend interaktif dengan <strong class="text-ink">SvelteKit</strong>, serta
						mengembangkan sistem AI yang dapat diimplementasikan ke dalam produksi. Saya percaya
						pada prinsip <em class="text-muted">"Clean Code"</em> dan skalabilitas sistem.
					</p>
				</div>
				<div class="mt-8">
					<h4 class="mb-3 text-xs font-semibold uppercase tracking-wider text-muted">
						Fokus Utama
					</h4>
					<div class="grid gap-3 sm:grid-cols-3">
						<div class="rounded-md border border-line bg-surface-alt/60 p-4">
							<p class="text-sm font-semibold text-ink">Backend kuat</p>
							<p class="mt-2 text-xs leading-5 text-muted">API, auth, persistence, dan guard.</p>
						</div>
						<div class="rounded-md border border-line bg-surface-alt/60 p-4">
							<p class="text-sm font-semibold text-ink">Frontend rapi</p>
							<p class="mt-2 text-xs leading-5 text-muted">Typed UI, load data, dan admin flow.</p>
						</div>
						<div class="rounded-md border border-line bg-surface-alt/60 p-4">
							<p class="text-sm font-semibold text-ink">AI praktis</p>
							<p class="mt-2 text-xs leading-5 text-muted">
								Chat, retrieval, dan cost-aware guard.
							</p>
						</div>
					</div>
				</div>
			</div>
		</div>
	</div>
</section>

<section class="border-t border-line bg-canvas-subtle py-20 sm:py-24" use:scrollAnimation>
	<div class="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
		<div class="mb-12 flex flex-col gap-5 md:flex-row md:items-end md:justify-between">
			<div class="max-w-3xl">
				<p class="text-sm font-semibold uppercase tracking-wide text-success">Selected work</p>
				<h2 class="mt-3 text-3xl font-extrabold tracking-tight text-ink sm:text-4xl">
					Project yang memberi konteks pada cara saya membangun.
				</h2>
			</div>
			<a
				href={resolve('/projects')}
				class="w-fit rounded-md border border-line bg-surface px-4 py-2 text-sm font-semibold text-ink transition hover:border-accent hover:bg-surface-alt"
			>
				Semua Project
			</a>
		</div>

		{#if projectError}
			<p class="rounded-md border border-danger/30 bg-danger-soft px-4 py-3 text-sm text-danger">
				{projectError}
			</p>
		{:else if selectedProjects.length === 0}
			<div class="rounded-md border border-line bg-canvas p-6 text-sm text-muted">
				Belum ada project dari backend. Tambahkan project melalui admin dashboard atau seed script
				agar section ini terisi.
			</div>
		{:else}
			<div class="grid gap-4 lg:grid-cols-3">
				{#each selectedProjects as project (project.id)}
					<article
						class="group flex min-h-[290px] flex-col justify-between rounded-md border border-line bg-surface p-6 transition-colors duration-200 hover:border-muted"
					>
						<div>
							<div class="mb-5 flex items-start justify-between gap-4">
								<div class="min-w-0">
									<p class="text-xs font-semibold uppercase tracking-wide text-muted">
										{projectCategory(project.techStack)}
									</p>
									<a
										href={resolve(`/projects/${project.id}`)}
										class="mt-2 block truncate text-xl font-bold text-ink transition-colors group-hover:text-accent hover:underline"
									>
										{project.title}
									</a>
								</div>
								<span
									class="mt-1 h-3 w-3 flex-shrink-0 rounded-full"
									style:background-color={projectAccent(project.techStack)}
									aria-hidden="true"
								></span>
							</div>
							<p class="line-clamp-4 text-sm leading-7 text-muted">{project.description}</p>
						</div>

						<div class="mt-8">
							<div class="mb-5 flex flex-wrap gap-2">
								{#each project.techStack.slice(0, 5) as tech (tech)}
									<span class="rounded-full border border-line px-2.5 py-1 text-xs text-body">
										{tech}
									</span>
								{/each}
								{#if project.techStack.length > 5}
									<span class="rounded-full border border-line px-2.5 py-1 text-xs text-muted">
										+{project.techStack.length - 5}
									</span>
								{/if}
							</div>
							<div class="flex flex-wrap gap-4">
								<a
									href={resolve(`/projects/${project.id}`)}
									class="text-sm font-semibold text-accent hover:underline"
								>
									View details →
								</a>
								{#if project.repoUrl}
									<a
										href={project.repoUrl}
										target="_blank"
										rel="noreferrer external"
										class="text-sm font-semibold text-muted hover:text-accent hover:underline"
									>
										Repository ↗
									</a>
								{/if}
							</div>
						</div>
					</article>
				{/each}
			</div>
		{/if}
	</div>
</section>

<section class="bg-canvas py-20 sm:py-24" use:scrollAnimation>
	<div class="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
		<div class="grid gap-12 lg:grid-cols-[0.8fr_1.2fr]">
			<div>
				<p class="text-sm font-semibold uppercase tracking-wide text-accent">Working style</p>
				<h2 class="mt-3 text-3xl font-extrabold tracking-tight text-ink sm:text-4xl">
					Saya suka membangun dari kontrak yang jelas sampai UI yang siap dipakai.
				</h2>
				<p class="mt-4 text-base leading-7 text-muted">
					Portfolio yang profesional harus terasa seperti produk kecil: punya arah, struktur, dan
					perhatian ke detail.
				</p>
			</div>

			<div class="grid gap-4 sm:grid-cols-2">
				{#each workflow as item (item.step)}
					<article class="rounded-md border border-line bg-surface p-6">
						<p class="font-mono text-xs font-semibold text-accent">{item.step}</p>
						<h3 class="mt-4 text-lg font-bold text-ink">{item.title}</h3>
						<p class="mt-3 text-sm leading-7 text-muted">{item.description}</p>
					</article>
				{/each}
			</div>
		</div>
	</div>
</section>

<section class="border-t border-line bg-canvas-subtle py-16" use:scrollAnimation>
	<div
		class="mx-auto flex max-w-7xl flex-col gap-6 px-4 sm:px-6 md:flex-row md:items-center md:justify-between lg:px-8"
	>
		<div class="max-w-2xl">
			<p class="text-sm font-semibold uppercase tracking-wide text-success">Open to collaborate</p>
			<h2 class="mt-3 text-2xl font-extrabold tracking-tight text-ink sm:text-3xl">
				Punya ide produk, dashboard, atau integrasi AI yang perlu dibuat serius?
			</h2>
			<p class="mt-3 text-sm leading-7 text-muted">
				Kirim konteks singkat. Saya akan balas dengan pertanyaan teknis yang tepat sebelum masuk ke
				solusi.
			</p>
		</div>
		<div class="flex flex-col gap-3 sm:flex-row">
			<a
				href={resolve('/contact')}
				class="inline-flex justify-center rounded-md border border-success/40 bg-success px-5 py-3 text-sm font-bold text-inverse transition hover:bg-success-hover"
			>
				Hubungi Saya
			</a>
			<a
				href={resolve('/experience')}
				class="inline-flex justify-center rounded-md border border-line bg-surface px-5 py-3 text-sm font-bold text-ink transition hover:border-accent hover:bg-surface-alt"
			>
				Lihat Pengalaman
			</a>
		</div>
	</div>
</section>
