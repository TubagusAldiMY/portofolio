<script lang="ts">
	import { onDestroy, onMount, tick } from 'svelte';
	import type { Snippet } from 'svelte';
	import gsap from 'gsap';

	type Props = {
		width?: number | string;
		height?: number | string;
		cardDistance?: number;
		verticalDistance?: number;
		delay?: number;
		pauseOnHover?: boolean;
		skewAmount?: number;
		easing?: 'elastic' | 'smooth';
		card0?: Snippet;
		card1?: Snippet;
		card2?: Snippet;
	};

	type Slot = {
		x: number;
		y: number;
		z: number;
		zIndex: number;
	};

	let {
		width = 300,
		height = 400,
		cardDistance = 45,
		verticalDistance = 35,
		delay = 4000,
		pauseOnHover = true,
		skewAmount = 2,
		easing = 'elastic',
		card0,
		card1,
		card2
	}: Props = $props();

	const snippets = $derived([card0, card1, card2].filter((card): card is Snippet => Boolean(card)));
	const dimensions = $derived({
		height: typeof height === 'number' ? `${height}px` : height,
		width: typeof width === 'number' ? `${width}px` : width
	});
	const config = $derived(
		easing === 'elastic'
			? {
					durDrop: 1.4,
					durMove: 1.4,
					durReturn: 1.4,
					ease: 'elastic.out(0.6,0.8)',
					promoteOverlap: 0.8,
					returnDelay: 0.1
				}
			: {
					durDrop: 0.8,
					durMove: 0.8,
					durReturn: 0.8,
					ease: 'power2.inOut',
					promoteOverlap: 0.45,
					returnDelay: 0.2
				}
	);

	let container: HTMLDivElement | undefined;
	let cardElements: Array<HTMLButtonElement | null> = [];
	let order = $state([0, 1, 2]);
	let timeline: gsap.core.Timeline | null = null;
	let interval: number | null = null;

	function makeSlot(index: number, total: number): Slot {
		return {
			x: index * cardDistance,
			y: -index * verticalDistance,
			z: -index * cardDistance * 1.5,
			zIndex: total - index
		};
	}

	function placeNow(element: HTMLButtonElement, slot: Slot): void {
		gsap.set(element, {
			force3D: true,
			skewY: skewAmount,
			transformOrigin: 'center center',
			x: slot.x,
			xPercent: -50,
			y: slot.y,
			yPercent: -50,
			z: slot.z,
			zIndex: slot.zIndex
		});
	}

	function registerCard(node: HTMLButtonElement, index: number): { destroy: () => void } {
		cardElements[index] = node;
		return {
			destroy: () => {
				cardElements[index] = null;
			}
		};
	}

	function initializeCards(): void {
		const total = snippets.length;
		cardElements.forEach((element, index) => {
			if (element) {
				placeNow(element, makeSlot(index, total));
			}
		});
	}

	function updateCardPositions(): void {
		const total = snippets.length;
		cardElements.forEach((element, index) => {
			if (!element) {
				return;
			}

			const slot = makeSlot(index, total);
			gsap.set(element, {
				skewY: skewAmount,
				x: slot.x,
				y: slot.y,
				z: slot.z
			});
		});
	}

	function swap(): void {
		if (order.length < 2) {
			return;
		}

		const [front, ...rest] = order;
		const frontElement = cardElements[front];

		if (!frontElement) {
			return;
		}

		timeline?.kill();
		const nextTimeline = gsap.timeline();
		timeline = nextTimeline;

		nextTimeline.to(frontElement, {
			duration: config.durDrop,
			ease: config.ease,
			rotation: -5,
			y: '+=350'
		});

		nextTimeline.addLabel('promote', `-=${config.durDrop * config.promoteOverlap}`);

		rest.forEach((index, restIndex) => {
			const element = cardElements[index];
			if (!element) {
				return;
			}

			const slot = makeSlot(restIndex, snippets.length);
			nextTimeline.set(element, { zIndex: slot.zIndex }, 'promote');
			nextTimeline.to(
				element,
				{
					duration: config.durMove,
					ease: config.ease,
					x: slot.x,
					y: slot.y,
					z: slot.z
				},
				`promote+=${restIndex * 0.15}`
			);
		});

		const backSlot = makeSlot(snippets.length - 1, snippets.length);
		nextTimeline.addLabel('return', `promote+=${config.durMove * config.returnDelay}`);
		nextTimeline.call(() => gsap.set(frontElement, { zIndex: backSlot.zIndex }), [], 'return');
		nextTimeline.set(frontElement, { rotation: 0, x: backSlot.x, z: backSlot.z }, 'return');
		nextTimeline.to(
			frontElement,
			{
				duration: config.durReturn,
				ease: config.ease,
				y: backSlot.y
			},
			'return'
		);

		nextTimeline.call(() => {
			order = [...rest, front];
		});
	}

	function startAnimation(): void {
		stopAnimation();
		interval = window.setInterval(swap, delay);
	}

	function stopAnimation(): void {
		timeline?.kill();
		timeline = null;

		if (interval) {
			clearInterval(interval);
			interval = null;
		}
	}

	function handleCardClick(): void {
		swap();
		startAnimation();
	}

	onMount(async () => {
		await tick();
		initializeCards();
		startAnimation();
	});

	onDestroy(() => {
		stopAnimation();
	});

	$effect(() => {
		if (typeof window !== 'undefined') {
			updateCardPositions();
		}
	});

	$effect(() => {
		if (typeof window !== 'undefined' && cardElements.some(Boolean)) {
			startAnimation();
		}
	});
</script>

<div
	bind:this={container}
	role="presentation"
	class="relative"
	style:width={dimensions.width}
	style:height={dimensions.height}
	style:perspective="1200px"
	onmouseenter={pauseOnHover ? stopAnimation : undefined}
	onmouseleave={pauseOnHover ? startAnimation : undefined}
>
	{#each snippets as card, index (index)}
		<button
			type="button"
			use:registerCard={index}
			class="absolute left-1/2 top-1/2 cursor-pointer overflow-hidden rounded-2xl border-0 bg-transparent p-0 text-left shadow-2xl [backface-visibility:hidden] [transform-style:preserve-3d] [will-change:transform]"
			style:width={dimensions.width}
			style:height={dimensions.height}
			onclick={handleCardClick}
		>
			{@render card()}
		</button>
	{/each}
</div>
