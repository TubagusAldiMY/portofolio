<script lang="ts">
	export type Token = { t: string; v: string };
	export type CodeLine = { tokens: Token[] };
	export type CodeCard = { filename: string; lines: CodeLine[] };

	type Props = {
		card: CodeCard;
		tokenColor: Record<string, string>;
	};

	let { card, tokenColor }: Props = $props();
</script>

<div
	class="flex h-full w-full flex-col rounded-2xl border border-[var(--theme-code-line)] bg-[var(--theme-code-canvas)] p-4 font-mono text-[10px] text-[var(--theme-code-body)] shadow-[var(--theme-code-shadow)] sm:p-6 sm:text-xs"
	style="color-scheme: dark"
>
	<div class="mb-4 flex items-center gap-2 border-b border-[var(--theme-code-line)] pb-3 sm:mb-6">
		<div class="h-3 w-3 rounded-full bg-[#ff5f56]"></div>
		<div class="h-3 w-3 rounded-full bg-[#ffbd2e]"></div>
		<div class="h-3 w-3 rounded-full bg-[#27c93f]"></div>
		<span
			class="ml-auto font-sans text-[10px] text-[var(--theme-code-muted)] opacity-75 sm:text-sm"
		>
			{card.filename}
		</span>
	</div>

	<div class="space-y-2 leading-relaxed opacity-90 sm:space-y-3">
		{#each card.lines as line, lineIndex (lineIndex)}
			<p>
				{#if line.tokens.length === 0}
					&nbsp;
				{:else}
					{#each line.tokens as token, tokenIndex (`${token.t}-${token.v}-${tokenIndex}`)}
						<span style:color={tokenColor[token.t] ?? tokenColor.text ?? 'var(--theme-code-body)'}
							>{token.v}</span
						>
					{/each}
				{/if}
			</p>
		{/each}
	</div>
</div>
