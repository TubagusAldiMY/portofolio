<script lang="ts">
	import { scrollAnimation } from '$lib/actions/scrollAnimation';
	import { ApiError } from '$lib/api/client';
	import { sendContactMessage } from '$lib/api/public';

	let name = $state('');
	let email = $state('');
	let content = $state('');
	let isSubmitting = $state(false);
	let status = $state<string | null>(null);
	let error = $state<string | null>(null);

	async function submitContact(event: SubmitEvent): Promise<void> {
		event.preventDefault();
		isSubmitting = true;
		status = null;
		error = null;

		try {
			const res = await sendContactMessage({ name, email, content });
			status = res.message;
			name = '';
			email = '';
			content = '';
		} catch (err) {
			error = err instanceof ApiError ? err.message : 'Gagal mengirim pesan.';
		} finally {
			isSubmitting = false;
		}
	}
</script>

<svelte:head>
	<title>Contact — Tubagus Aldi</title>
	<meta
		name="description"
		content="Get in touch with Tubagus Aldi for collaboration or inquiries."
	/>
</svelte:head>

<section class="border-t border-line bg-canvas py-20" use:scrollAnimation>
	<div class="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
		<div class="mb-12 text-center">
			<h2 class="text-3xl font-extrabold text-ink">Hubungi Saya</h2>
			<p class="mt-2 text-muted">Kirim pesan, kolaborasi, atau sekadar menyapa.</p>
		</div>

		<div class="mx-auto max-w-2xl">
			<!-- Terminal-style card -->
			<div class="overflow-hidden rounded-md border border-line bg-surface shadow-lg">
				<!-- Window chrome -->
				<div class="flex items-center gap-2 border-b border-line bg-surface-alt px-4 py-3">
					<div class="h-3 w-3 rounded-full bg-[#f25d5e]"></div>
					<div class="h-3 w-3 rounded-full bg-[#fbfb8d]"></div>
					<div class="h-3 w-3 rounded-full bg-[#62c554]"></div>
					<span class="ml-2 font-mono text-xs text-muted">contact.md</span>
				</div>

				<div class="p-6">
					{#if status}
						<p
							class="mb-5 rounded-md border border-success/30 bg-success-soft px-4 py-3 text-sm text-success"
						>
							{status}
						</p>
					{/if}
					{#if error}
						<p
							class="mb-5 rounded-md border border-danger/30 bg-danger-soft px-4 py-3 text-sm text-danger"
						>
							{error}
						</p>
					{/if}

					<form class="space-y-5" onsubmit={submitContact}>
						<div class="grid gap-5 md:grid-cols-2">
							<div>
								<label
									for="c-name"
									class="mb-2 block text-xs font-semibold uppercase tracking-wide text-body"
									>Nama Lengkap</label
								>
								<input
									id="c-name"
									bind:value={name}
									required
									placeholder="Nama Anda"
									class="w-full rounded-md border border-line bg-canvas px-3 py-2 text-ink placeholder-muted outline-none transition focus:border-accent focus:ring-1 focus:ring-accent"
								/>
							</div>
							<div>
								<label
									for="c-email"
									class="mb-2 block text-xs font-semibold uppercase tracking-wide text-body"
									>Alamat Email</label
								>
								<input
									id="c-email"
									bind:value={email}
									required
									type="email"
									placeholder="nama@email.com"
									class="w-full rounded-md border border-line bg-canvas px-3 py-2 text-ink placeholder-muted outline-none transition focus:border-accent focus:ring-1 focus:ring-accent"
								/>
							</div>
						</div>

						<div>
							<label
								for="c-content"
								class="mb-2 block text-xs font-semibold uppercase tracking-wide text-body"
								>Pesan</label
							>
							<textarea
								id="c-content"
								bind:value={content}
								required
								rows="6"
								placeholder="Tulis pesan Anda di sini..."
								class="w-full rounded-md border border-line bg-canvas px-3 py-2 font-mono text-sm text-ink placeholder-muted outline-none transition focus:border-accent focus:ring-1 focus:ring-accent"
							></textarea>
							<p class="mt-1 text-right text-xs text-muted">Markdown is supported</p>
						</div>

						<div class="flex justify-end">
							<button
								type="submit"
								disabled={isSubmitting}
								class="rounded-md border border-success/40 bg-success px-6 py-2 text-sm font-semibold text-inverse shadow-sm transition-colors hover:bg-success-hover disabled:cursor-not-allowed disabled:opacity-50"
							>
								{isSubmitting ? 'Mengirim...' : 'Kirim Pesan'}
							</button>
						</div>
					</form>
				</div>
			</div>

			<div class="mt-8 text-center text-sm text-muted">
				Atau temukan saya di
				<a
					href="https://linkedin.com/in/tubagusaldi"
					target="_blank"
					rel="noreferrer external"
					class="text-accent hover:underline">LinkedIn</a
				>
				dan
				<a
					href="https://github.com/TubagusAldiMY"
					target="_blank"
					rel="noreferrer external"
					class="text-accent hover:underline">GitHub</a
				>.
			</div>
		</div>
	</div>
</section>
