<script lang="ts">
	type MarqueeItem =
		| { kind: 'tech'; name: string; color: string; bg: string; icon: string }
		| { kind: 'label'; text: string };

	// SVG icon content (viewBox="0 0 24 24")
	const icons: Record<string, string> = {
		go: `<rect width="24" height="24" rx="4" fill="#00ADD8"/><path d="M5 15.5V9.5M5 9.5L9 12.5L5 15.5M9 9.5V15.5M13 9.5H16.5C17.9 9.5 19 10.6 19 12C19 13.4 17.9 14.5 16.5 14.5H13V9.5Z" stroke="white" stroke-width="1.6" stroke-linecap="round" stroke-linejoin="round" fill="none"/>`,

		rust: `<circle cx="12" cy="12" r="9" fill="none" stroke="#CE412B" stroke-width="1.5"/><path d="M12 5L13.5 8.5H17L14.3 10.5L15.3 14L12 12L8.7 14L9.7 10.5L7 8.5H10.5Z" fill="#CE412B"/><circle cx="12" cy="12" r="2.5" fill="none" stroke="#CE412B" stroke-width="1.2"/>`,

		postgresql: `<path d="M16 5C16 3.3 14.2 2 12 2C9.8 2 8 3.3 8 5L7 17C7 19.2 9.2 21 12 21C14.8 21 17 19.2 17 17L16 5Z" fill="none" stroke="#336791" stroke-width="1.5"/><path d="M8 8H16M8 12H16" stroke="#336791" stroke-width="1.2"/><path d="M14 2C14 2 16.5 3 17 5M10 2C10 2 7.5 3 7 5" stroke="#336791" stroke-width="1.2" stroke-linecap="round"/>`,

		docker: `<rect x="2" y="11" width="4" height="3" rx="0.5" fill="#2496ED"/><rect x="7" y="11" width="4" height="3" rx="0.5" fill="#2496ED"/><rect x="12" y="11" width="4" height="3" rx="0.5" fill="#2496ED"/><rect x="7" y="7" width="4" height="3" rx="0.5" fill="#2496ED"/><rect x="12" y="7" width="4" height="3" rx="0.5" fill="#2496ED"/><path d="M17.5 12C18.5 10.5 20 11 21 11.5" stroke="#2496ED" stroke-width="1.3" stroke-linecap="round"/><path d="M2 14C3 15.5 5 16 8 15.5L18 15C19.5 15 20.5 14 20.5 12.5C20 12 19 11.5 17.5 12" fill="#2496ED" opacity="0.25"/>`,

		svelte: `<path d="M19.5 7C18.2 4.8 15.2 4 12.8 5.8L8.2 9C6.2 10.4 5.8 13.2 7 15L7.2 15.4C6.2 16.8 6 18.8 7.2 20.5C8.5 22.7 11.5 23.5 13.9 21.7L18.5 18.5C20.5 17.1 20.9 14.3 19.7 12.5L19.5 12.1C20.5 10.7 20.7 8.7 19.5 7Z" fill="#FF3E00"/><path d="M15 10.5C15.4 9.2 14.9 7.8 13.6 7.1C12.5 6.5 11.2 6.7 10.2 7.4L7.8 9.2C7.1 9.7 6.7 10.6 7 11.3C7.3 12 8.1 12.4 8.9 12.2L11.3 11.6C12 11.4 12.7 11.6 13.2 12.1C13.7 12.6 13.9 13.3 13.7 14L12.8 17C12.5 18.1 13.2 19.2 14.2 19.5C15.2 19.8 16.3 19.2 16.8 18.3L18.5 14.5C19.2 13.1 18.7 11.5 17.5 10.5C16.7 9.8 15.7 9.7 15 10Z" fill="white"/>`,

		typescript: `<rect width="24" height="24" rx="3" fill="#3178C6"/><path d="M16.5 8.5H12.5V10H14V16H16.5V8.5Z" fill="white"/><path d="M11 10V12H13V10H11ZM11 12V16H13.5V12H11Z" fill="white"/>`,

		react: `<ellipse cx="12" cy="12" rx="9.5" ry="3.5" fill="none" stroke="#61DAFB" stroke-width="1.4"/><ellipse cx="12" cy="12" rx="9.5" ry="3.5" fill="none" stroke="#61DAFB" stroke-width="1.4" transform="rotate(60 12 12)"/><ellipse cx="12" cy="12" rx="9.5" ry="3.5" fill="none" stroke="#61DAFB" stroke-width="1.4" transform="rotate(120 12 12)"/><circle cx="12" cy="12" r="2" fill="#61DAFB"/>`,

		tailwind: `<path d="M6 9C7.3 6 9.5 5.3 11.5 7C13.5 8.7 13.5 11.3 15.5 12.5C17.5 13.7 19.5 13 21 11C19.7 14 17.5 14.7 15.5 13C13.5 11.3 13.5 8.7 11.5 7.5C9.5 6.3 7.5 7 6 9Z" fill="#38BDF8"/><path d="M3 14C4.3 11 6.5 10.3 8.5 12C10.5 13.7 10.5 16.3 12.5 17.5C14.5 18.7 16.5 18 18 16C16.7 19 14.5 19.7 12.5 18C10.5 16.3 10.5 13.7 8.5 12.5C6.5 11.3 4.5 12 3 14Z" fill="#38BDF8"/>`,

		python: `<path d="M12 2C9.2 2 7 3.8 7 6V8H17V6C17 3.8 14.8 2 12 2Z" fill="#3776AB"/><circle cx="9.5" cy="5" r="1" fill="white"/><path d="M7 8H17L17 13H12.5C10 13 8 14.8 8 17V19H5C3.3 19 2 17.7 2 16V10C2 8.9 2.9 8 4 8H7Z" fill="#3776AB"/><path d="M12 22C14.8 22 17 20.2 17 18V16H7V18C7 20.2 9.2 22 12 22Z" fill="#FFD845"/><circle cx="14.5" cy="19" r="1" fill="#3776AB"/><path d="M17 16H7V11H11.5C14 11 16 9.2 16 7V5H19C20.7 5 22 6.3 22 8V14C22 15.1 21.1 16 20 16H17Z" fill="#FFD845"/>`,

		gemini: `<defs><linearGradient id="ggrad" x1="4" y1="3" x2="20" y2="21"><stop stop-color="#4285F4"/><stop offset="0.5" stop-color="#A855F7"/><stop offset="1" stop-color="#EC4899"/></linearGradient></defs><path d="M12 2.5C12 7.5 8.5 11.5 3.5 12C8.5 12.5 12 16.5 12 21.5C12 16.5 15.5 12.5 20.5 12C15.5 11.5 12 7.5 12 2.5Z" fill="url(#ggrad)"/>`,

		tensorflow: `<rect width="24" height="24" rx="3" fill="#FF6F00"/><path d="M12 4L18 7.5V14.5L12 18L6 14.5V7.5L12 4Z" fill="none" stroke="white" stroke-width="1.5"/><path d="M12 4V18M6 7.5L18 14.5M18 7.5L6 14.5" stroke="white" stroke-width="1.2" opacity="0.6"/>`,

		shield: `<path d="M12 3L4 6.5V12C4 16.7 7.5 21 12 22C16.5 21 20 16.7 20 12V6.5L12 3Z" fill="none" stroke="#238636" stroke-width="1.5"/><path d="M9 12L11 14L15.5 9.5" stroke="#238636" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/>`,

		jwt: `<circle cx="12" cy="12" r="9" fill="none" stroke="#9333EA" stroke-width="1.4"/><path d="M10 9L12 15M12 15L14 9M8 9H10M14 9H16" stroke="#9333EA" stroke-width="1.4" stroke-linecap="round"/><circle cx="12" cy="19" r="1.2" fill="#9333EA"/>`,

		testing: `<path d="M9 3H15V5L13 7H11L9 5V3Z" fill="none" stroke="#7C3AED" stroke-width="1.5"/><path d="M11 7V10L7 17H17L13 10V7" fill="none" stroke="#7C3AED" stroke-width="1.5" stroke-linejoin="round"/><path d="M9 14H15" stroke="#7C3AED" stroke-width="1.5" stroke-linecap="round"/>`,

		bun: `<circle cx="12" cy="12" r="9" fill="#FBF0DF" stroke="#F6DECE" stroke-width="0.5"/><circle cx="9" cy="10" r="1.5" fill="#14120B"/><circle cx="15" cy="10" r="1.5" fill="#14120B"/><path d="M8 14C9.5 16 14.5 16 16 14" fill="none" stroke="#14120B" stroke-width="1.4" stroke-linecap="round"/>`
	};

	const row1: MarqueeItem[] = [
		{ kind: 'label', text: 'Backend architecture' },
		{ kind: 'tech', name: 'Go', color: '#00ADD8', bg: '#001F26', icon: icons.go },
		{ kind: 'tech', name: 'Rust', color: '#CE412B', bg: '#200A06', icon: icons.rust },
		{ kind: 'tech', name: 'PostgreSQL', color: '#336791', bg: '#06141F', icon: icons.postgresql },
		{ kind: 'tech', name: 'Docker', color: '#2496ED', bg: '#031729', icon: icons.docker },
		{ kind: 'tech', name: 'Bun', color: '#FBB040', bg: '#1A1206', icon: icons.bun },
		{ kind: 'label', text: 'Typed frontend' },
		{ kind: 'tech', name: 'SvelteKit', color: '#FF3E00', bg: '#1F0700', icon: icons.svelte },
		{ kind: 'tech', name: 'TypeScript', color: '#3178C6', bg: '#030E1F', icon: icons.typescript },
		{ kind: 'tech', name: 'React', color: '#61DAFB', bg: '#011920', icon: icons.react },
		{ kind: 'tech', name: 'Tailwind CSS', color: '#38BDF8', bg: '#021520', icon: icons.tailwind }
	];

	const row2: MarqueeItem[] = [
		{ kind: 'label', text: 'AI workflow' },
		{ kind: 'tech', name: 'Python', color: '#3776AB', bg: '#030E1F', icon: icons.python },
		{ kind: 'tech', name: 'TensorFlow', color: '#FF6F00', bg: '#1A0C00', icon: icons.tensorflow },
		{ kind: 'tech', name: 'Gemini AI', color: '#A855F7', bg: '#0F0714', icon: icons.gemini },
		{ kind: 'label', text: 'Security review' },
		{ kind: 'tech', name: 'Secure APIs', color: '#238636', bg: '#071209', icon: icons.shield },
		{ kind: 'tech', name: 'JWT Auth', color: '#9333EA', bg: '#0E0714', icon: icons.jwt },
		{ kind: 'tech', name: 'Testing', color: '#7C3AED', bg: '#0B0414', icon: icons.testing },
		{ kind: 'tech', name: 'Gemini AI', color: '#A855F7', bg: '#0F0714', icon: icons.gemini },
		{ kind: 'tech', name: 'Python', color: '#3776AB', bg: '#030E1F', icon: icons.python }
	];
</script>

<div class="tech-marquee relative overflow-hidden border-y border-line bg-canvas-subtle py-6">
	<!-- Fade masks -->
	<div
		class="pointer-events-none absolute inset-y-0 left-0 z-10 w-24 bg-gradient-to-r from-canvas-subtle to-transparent"
	></div>
	<div
		class="pointer-events-none absolute inset-y-0 right-0 z-10 w-24 bg-gradient-to-l from-canvas-subtle to-transparent"
	></div>

	<!-- Row 1: scroll left -->
	<div class="marquee-track marquee-track-left mb-3 flex w-max gap-3">
		{#each [...row1, ...row1] as item, i (i)}
			{#if item.kind === 'label'}
				<div class="flex flex-shrink-0 items-center gap-1.5 px-2">
					<span class="text-[10px] font-semibold uppercase tracking-[0.2em] text-muted">◆</span>
					<span class="text-xs font-semibold uppercase tracking-widest text-muted">{item.text}</span
					>
					<span class="text-[10px] font-semibold uppercase tracking-[0.2em] text-muted">◆</span>
				</div>
			{:else}
				<div
					class="flex flex-shrink-0 items-center gap-2.5 rounded-full border border-line px-3 py-1.5 transition-colors hover:border-muted"
					style="background-color: color-mix(in srgb, {item.color} 12%, var(--theme-surface))"
				>
					<svg viewBox="0 0 24 24" width="18" height="18" class="flex-shrink-0">
						<!-- eslint-disable-next-line svelte/no-at-html-tags -->
						{@html item.icon}
					</svg>
					<span class="text-sm font-medium text-body">{item.name}</span>
				</div>
			{/if}
		{/each}
	</div>

	<!-- Row 2: scroll right -->
	<div class="marquee-track marquee-track-right flex w-max gap-3">
		{#each [...row2, ...row2] as item, i (i)}
			{#if item.kind === 'label'}
				<div class="flex flex-shrink-0 items-center gap-1.5 px-2">
					<span class="text-[10px] font-semibold uppercase tracking-[0.2em] text-muted">◆</span>
					<span class="text-xs font-semibold uppercase tracking-widest text-muted">{item.text}</span
					>
					<span class="text-[10px] font-semibold uppercase tracking-[0.2em] text-muted">◆</span>
				</div>
			{:else}
				<div
					class="flex flex-shrink-0 items-center gap-2.5 rounded-full border border-line px-3 py-1.5 transition-colors hover:border-muted"
					style="background-color: color-mix(in srgb, {item.color} 12%, var(--theme-surface))"
				>
					<svg viewBox="0 0 24 24" width="18" height="18" class="flex-shrink-0">
						<!-- eslint-disable-next-line svelte/no-at-html-tags -->
						{@html item.icon}
					</svg>
					<span class="text-sm font-medium text-body">{item.name}</span>
				</div>
			{/if}
		{/each}
	</div>
</div>

<style>
	.marquee-track {
		will-change: transform;
	}

	.marquee-track-left {
		animation: marquee-left 28s linear infinite;
	}

	.marquee-track-right {
		animation: marquee-right 34s linear infinite;
	}

	@keyframes marquee-left {
		from {
			transform: translate3d(0, 0, 0);
		}
		to {
			transform: translate3d(-50%, 0, 0);
		}
	}

	@keyframes marquee-right {
		from {
			transform: translate3d(-50%, 0, 0);
		}
		to {
			transform: translate3d(0, 0, 0);
		}
	}

	.tech-marquee:hover .marquee-track {
		animation-play-state: paused;
	}

	@media (prefers-reduced-motion: reduce) {
		.marquee-track {
			animation-duration: 1ms;
			animation-iteration-count: 1;
			transform: translate3d(0, 0, 0);
		}
	}
</style>
