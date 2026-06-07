<script lang="ts">
	import { tick } from 'svelte';
	import { marked } from 'marked';
	import { apiFetch } from '$lib/api/client';
	import type { ChatResponse } from '$lib/api/types';

	type Message = { id: number; text: string; isUser: boolean };

	let isOpen = $state(false);
	let userInput = $state('');
	let isLoading = $state(false);
	let messagesEl = $state<HTMLElement | null>(null);

	let messages = $state<Message[]>([
		{
			id: 1,
			text: 'Halo! Saya asisten virtual **Tubagus Aldi**. Ada yang bisa saya bantu tentang portofolio ini?',
			isUser: false
		}
	]);

	function renderMarkdown(text: string): string {
		return marked.parse(text) as string;
	}

	async function scrollToBottom() {
		await tick();
		if (messagesEl) messagesEl.scrollTop = messagesEl.scrollHeight;
	}

	function toggle() {
		isOpen = !isOpen;
		if (isOpen) scrollToBottom();
	}

	async function sendMessage(e: SubmitEvent) {
		e.preventDefault();
		const text = userInput.trim();
		if (!text || isLoading) return;

		messages.push({ id: Date.now(), text, isUser: true });
		userInput = '';
		isLoading = true;
		scrollToBottom();

		try {
			const data = await apiFetch<ChatResponse>('/api/chat', {
				auth: false,
				method: 'POST',
				body: { message: text }
			});
			messages.push({ id: Date.now() + 1, text: data.reply, isUser: false });
		} catch {
			messages.push({
				id: Date.now() + 1,
				text: 'Gagal terhubung ke server. Silakan coba lagi nanti.',
				isUser: false
			});
		} finally {
			isLoading = false;
			scrollToBottom();
		}
	}
</script>

<div class="fixed bottom-6 right-6 z-50 flex flex-col items-end">
	<!-- Chat window -->
	{#if isOpen}
		<div
			class="mb-4 flex h-[450px] w-80 flex-col overflow-hidden rounded-xl border border-line bg-surface shadow-2xl sm:w-96"
			style="animation: chatOpen 0.3s cubic-bezier(0.16,1,0.3,1)"
		>
			<!-- Header -->
			<div class="flex items-center justify-between border-b border-line bg-canvas p-4">
				<div class="flex items-center gap-3">
					<div class="relative">
						<div
							class="flex h-8 w-8 items-center justify-center rounded-full bg-gradient-to-tr from-accent-emphasis to-accent text-xs font-bold text-inverse"
						>
							AI
						</div>
						<div
							class="absolute bottom-0 right-0 h-2.5 w-2.5 animate-pulse rounded-full border-2 border-canvas bg-success"
						></div>
					</div>
					<h3 class="text-sm font-bold text-ink">Aldi's Assistant</h3>
				</div>
				<button
					onclick={toggle}
					class="rounded p-1 text-muted transition-colors hover:bg-surface-alt hover:text-ink"
					aria-label="Close chat"
				>
					<svg class="h-5 w-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
						<path
							stroke-linecap="round"
							stroke-linejoin="round"
							stroke-width="2"
							d="M6 18L18 6M6 6l12 12"
						/>
					</svg>
				</button>
			</div>

			<!-- Messages -->
			<div
				bind:this={messagesEl}
				class="scrollbar-gh flex-1 space-y-4 overflow-y-auto bg-canvas/95 p-4"
			>
				{#each messages as msg (msg.id)}
					<div class="flex {msg.isUser ? 'justify-end' : 'justify-start'}">
						<div
							class="max-w-[85%] rounded-2xl px-4 py-2.5 text-sm leading-relaxed shadow-sm
								{msg.isUser
								? 'rounded-br-none bg-accent-emphasis text-inverse'
								: 'rounded-bl-none border border-line bg-surface-alt text-body'}"
						>
							{#if msg.isUser}
								{msg.text}
							{:else}
								<!-- eslint-disable-next-line svelte/no-at-html-tags -->
								<div class="markdown-content">{@html renderMarkdown(msg.text)}</div>
							{/if}
						</div>
					</div>
				{/each}

				{#if isLoading}
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
			</div>

			<!-- Input -->
			<form class="flex gap-2 border-t border-line bg-surface p-3" onsubmit={sendMessage}>
				<input
					bind:value={userInput}
					type="text"
					placeholder="Tanya tentang pengalaman, tech stack..."
					class="flex-1 rounded-lg border border-line bg-canvas px-4 py-2.5 text-sm text-ink outline-none placeholder-muted transition focus:border-accent focus:ring-1 focus:ring-accent"
				/>
				<button
					type="submit"
					disabled={isLoading || !userInput.trim()}
					class="flex h-10 w-10 items-center justify-center rounded-lg bg-success text-inverse transition-all hover:bg-success-hover active:scale-95 disabled:cursor-not-allowed disabled:opacity-50"
					aria-label="Send message"
				>
					{#if isLoading}
						<svg class="h-4 w-4 animate-spin" fill="none" viewBox="0 0 24 24">
							<circle
								class="opacity-25"
								cx="12"
								cy="12"
								r="10"
								stroke="currentColor"
								stroke-width="4"
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
	{/if}

	<!-- Toggle button -->
	<button
		onclick={toggle}
		class="group relative z-50 flex h-14 w-14 items-center justify-center rounded-full bg-accent-emphasis text-inverse shadow-lg shadow-accent-emphasis/20 transition-all duration-300 hover:scale-110 hover:bg-accent-emphasis-hover active:scale-95"
		aria-label="Toggle AI chat"
	>
		{#if !isOpen}
			<svg
				class="h-7 w-7 transition-transform group-hover:rotate-12"
				fill="none"
				stroke="currentColor"
				viewBox="0 0 24 24"
			>
				<path
					stroke-linecap="round"
					stroke-linejoin="round"
					stroke-width="2"
					d="M8 10h.01M12 10h.01M16 10h.01M9 16H5a2 2 0 01-2-2V6a2 2 0 012-2h14a2 2 0 012 2v8a2 2 0 01-2 2h-5l-5 5v-5z"
				/>
			</svg>
			<!-- Ping indicator -->
			<span class="absolute -right-1 -top-1 flex h-3 w-3">
				<span
					class="absolute inline-flex h-full w-full animate-ping rounded-full bg-danger opacity-75"
				></span>
				<span class="relative inline-flex h-3 w-3 rounded-full bg-danger"></span>
			</span>
		{:else}
			<svg class="h-7 w-7" fill="none" stroke="currentColor" viewBox="0 0 24 24">
				<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7" />
			</svg>
		{/if}
	</button>
</div>

<style>
	@keyframes chatOpen {
		from {
			opacity: 0;
			transform: scale(0.9) translateY(20px) translateX(20px);
		}
		to {
			opacity: 1;
			transform: scale(1) translateY(0) translateX(0);
		}
	}
</style>
