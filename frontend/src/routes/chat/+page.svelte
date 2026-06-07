<script lang="ts">
	import { tick } from 'svelte';
	import { marked } from 'marked';
	import { ApiError } from '$lib/api/client';
	import { sendChatMessage } from '$lib/api/public';

	type ChatEntry = { id: number; role: 'assistant' | 'user'; text: string };

	let input = $state('');
	let isSending = $state(false);
	let error = $state<string | null>(null);
	let messagesEl = $state<HTMLElement | null>(null);

	let messages = $state<ChatEntry[]>([
		{
			id: 1,
			role: 'assistant',
			text: 'Halo! Saya asisten virtual **Tubagus Aldi**. Tanyakan apa saja tentang proyek, pengalaman, atau tech stack.'
		}
	]);

	const suggestions = [
		'Apa tech stack yang kamu pakai?',
		'Ceritakan tentang KasKu project.',
		'Apa pengalaman kerja kamu?',
		'Bagaimana cara menghubungi Tubagus Aldi?'
	];

	async function scrollToBottom() {
		await tick();
		if (messagesEl) messagesEl.scrollTop = messagesEl.scrollHeight;
	}

	async function submitChat(e: SubmitEvent) {
		e.preventDefault();
		const message = input.trim();
		if (!message || isSending) return;

		messages.push({ id: Date.now(), role: 'user', text: message });
		input = '';
		isSending = true;
		error = null;
		scrollToBottom();

		try {
			const res = await sendChatMessage({ message });
			messages.push({ id: Date.now() + 1, role: 'assistant', text: res.reply });
		} catch (err) {
			error = err instanceof ApiError ? err.message : 'Gagal menghubungi AI service.';
		} finally {
			isSending = false;
			scrollToBottom();
		}
	}

	async function useSuggestion(s: string) {
		input = s;
		await tick();
		const form = document.querySelector<HTMLFormElement>('#chat-form');
		form?.requestSubmit();
	}
</script>

<svelte:head>
	<title>AI Assistant — Tubagus Aldi</title>
	<meta
		name="description"
		content="Chat with Tubagus Aldi's AI assistant about his projects, skills, and experience."
	/>
</svelte:head>

<section
	class="mx-auto flex max-w-3xl flex-col px-4 py-10 sm:px-6 lg:px-8"
	style="min-height: calc(100dvh - 8rem)"
>
	<!-- Header -->
	<div class="mb-6">
		<div class="flex items-center gap-3 mb-1">
			<div class="relative flex-shrink-0">
				<div
					class="flex h-10 w-10 items-center justify-center rounded-full bg-gradient-to-tr from-accent-emphasis to-accent text-sm font-bold text-inverse"
				>
					AI
				</div>
				<div
					class="absolute bottom-0 right-0 h-3 w-3 animate-pulse rounded-full border-2 border-canvas bg-success"
				></div>
			</div>
			<div>
				<h1 class="text-xl font-bold text-ink">Aldi's Assistant</h1>
				<p class="text-xs text-muted">Online · Powered by Gemini</p>
			</div>
		</div>
	</div>

	<!-- Chat window -->
	<div class="flex flex-1 flex-col overflow-hidden rounded-md border border-line bg-surface">
		<!-- Messages -->
		<div bind:this={messagesEl} class="scrollbar-gh flex-1 space-y-4 overflow-y-auto p-5">
			{#each messages as msg (msg.id)}
				<div class="flex {msg.role === 'user' ? 'justify-end' : 'justify-start'}">
					<div
						class="max-w-[85%] rounded-2xl px-4 py-2.5 text-sm leading-relaxed
						{msg.role === 'user'
							? 'rounded-br-none bg-accent-emphasis text-inverse'
							: 'rounded-bl-none border border-line bg-surface-alt text-body'}"
					>
						{#if msg.role === 'assistant'}
							<!-- eslint-disable-next-line svelte/no-at-html-tags -->
							<div class="markdown-content">{@html marked.parse(msg.text)}</div>
						{:else}
							{msg.text}
						{/if}
					</div>
				</div>
			{/each}

			{#if isSending}
				<div class="flex justify-start">
					<div
						class="flex items-center gap-1 rounded-2xl rounded-bl-none border border-line bg-surface-alt px-4 py-3"
					>
						<div class="h-1.5 w-1.5 animate-bounce rounded-full bg-muted"></div>
						<div class="h-1.5 w-1.5 animate-bounce rounded-full bg-muted delay-75"></div>
						<div class="h-1.5 w-1.5 animate-bounce rounded-full bg-muted delay-150"></div>
					</div>
				</div>
			{/if}

			{#if error}
				<p class="text-center text-xs text-danger">{error}</p>
			{/if}
		</div>

		<!-- Suggestions (only show initially) -->
		{#if messages.length <= 1}
			<div class="border-t border-line px-4 py-3">
				<p class="mb-2 text-xs text-muted">Coba tanyakan:</p>
				<div class="flex flex-wrap gap-2">
					{#each suggestions as s (s)}
						<button
							onclick={() => useSuggestion(s)}
							class="rounded-full border border-line bg-surface-alt px-3 py-1.5 text-xs text-body transition hover:border-accent hover:text-accent"
						>
							{s}
						</button>
					{/each}
				</div>
			</div>
		{/if}

		<!-- Input -->
		<form
			id="chat-form"
			class="flex gap-2 border-t border-line bg-canvas p-3"
			onsubmit={submitChat}
		>
			<input
				bind:value={input}
				placeholder="Tanya tentang proyek, pengalaman, atau tech stack..."
				class="flex-1 rounded-lg border border-line bg-surface px-4 py-2.5 text-sm text-ink outline-none placeholder-muted transition focus:border-accent focus:ring-1 focus:ring-accent"
			/>
			<button
				type="submit"
				disabled={isSending || !input.trim()}
				class="flex h-10 w-10 items-center justify-center rounded-lg bg-success text-inverse transition hover:bg-success-hover active:scale-95 disabled:cursor-not-allowed disabled:opacity-50"
				aria-label="Send"
			>
				{#if isSending}
					<svg class="h-4 w-4 animate-spin" fill="none" viewBox="0 0 24 24">
						<circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"
						></circle>
						<path
							class="opacity-75"
							fill="currentColor"
							d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"
						></path>
					</svg>
				{:else}
					<svg class="ml-0.5 h-4 w-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
						<path
							stroke-linecap="round"
							stroke-linejoin="round"
							stroke-width="2"
							d="M12 19l9 2-9-18-9 18 9-2zm0 0v-8"
						/>
					</svg>
				{/if}
			</button>
		</form>
	</div>
</section>
