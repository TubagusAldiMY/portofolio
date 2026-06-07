<script lang="ts">
	type Props = {
		class?: string;
		lines?: number;
		card?: boolean;
	};

	let { class: cls = '', lines = 3, card = false }: Props = $props();

	function lineIndexes(count: number): number[] {
		return [...Array(count).keys()];
	}
</script>

{#if card}
	<div class="rounded-md border border-line bg-surface p-6 {cls}">
		<div class="mb-3 flex items-center gap-2">
			<div class="h-4 w-4 animate-pulse rounded-full bg-surface-alt"></div>
			<div class="h-4 w-40 animate-pulse rounded bg-surface-alt"></div>
			<div class="ml-2 h-4 w-12 animate-pulse rounded-full bg-surface-alt"></div>
		</div>
		<div class="space-y-2">
			<div class="h-3 w-full animate-pulse rounded bg-surface-alt"></div>
			<div class="h-3 w-4/5 animate-pulse rounded bg-surface-alt"></div>
			<div class="h-3 w-3/5 animate-pulse rounded bg-surface-alt"></div>
		</div>
		<div class="mt-5 flex gap-3">
			<div class="h-3 w-16 animate-pulse rounded bg-surface-alt"></div>
			<div class="h-3 w-16 animate-pulse rounded bg-surface-alt"></div>
		</div>
	</div>
{:else}
	<div class="space-y-2 {cls}">
		{#each lineIndexes(lines) as i (i)}
			<div
				class="h-3 animate-pulse rounded bg-surface-alt"
				style="width: {i === lines - 1 ? '60' : 90 - i * 5}%"
			></div>
		{/each}
	</div>
{/if}
