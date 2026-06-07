<script lang="ts">
	import { onMount } from 'svelte';
	import { goto } from '$app/navigation';
	import { resolve } from '$app/paths';
	import type { Pathname } from '$app/types';

	import { ApiError } from '$lib/api/client';
	import {
		createResource,
		deleteResource,
		listMessages,
		updateResource,
		uploadFile
	} from '$lib/api/admin';
	import { listExperiences, listProducts, listProjects } from '$lib/api/public';
	import type { AdminResource, ContactMessage, Experience, Product, Project } from '$lib/api/types';
	import { clearAuthToken } from '$lib/auth/auth.svelte';

	type Tab = 'overview' | AdminResource | 'messages';

	type ProjectForm = {
		title: string;
		description: string;
		imageUrl: string;
		techStack: string;
		repoUrl: string;
		liveUrl: string;
	};

	type ProductForm = {
		title: string;
		description: string;
		price: string;
		imageUrl: string;
		features: string;
		buyUrl: string;
		tag: string;
	};

	type ExperienceForm = {
		role: string;
		company: string;
		duration: string;
		description: string;
	};

	let activeTab = $state<Tab>('overview');
	let projects = $state<Project[]>([]);
	let products = $state<Product[]>([]);
	let experiences = $state<Experience[]>([]);
	let messages = $state<ContactMessage[]>([]);
	let isLoading = $state(true);
	let error = $state<string | null>(null);
	let notice = $state<string | null>(null);
	let savingResource = $state<AdminResource | null>(null);
	let deletingKey = $state<string | null>(null);
	let uploadingFor = $state<'project' | 'product' | null>(null);
	let editingProjectId = $state<number | null>(null);
	let editingProductId = $state<number | null>(null);
	let editingExperienceId = $state<number | null>(null);

	let projectForm = $state<ProjectForm>(emptyProjectForm());
	let productForm = $state<ProductForm>(emptyProductForm());
	let experienceForm = $state<ExperienceForm>(emptyExperienceForm());

	const tabs: { id: Tab; label: string }[] = [
		{ id: 'overview', label: 'Overview' },
		{ id: 'projects', label: 'Projects' },
		{ id: 'products', label: 'Products' },
		{ id: 'experiences', label: 'Experience' },
		{ id: 'messages', label: 'Messages' }
	];

	const resourceLinks: { href: Pathname; label: string }[] = [
		{ href: '/projects', label: 'Projects' },
		{ href: '/products', label: 'Products' },
		{ href: '/experience', label: 'Experience' }
	];

	const inputClass =
		'mt-1 w-full rounded-md border border-line bg-canvas px-3 py-2 text-sm text-ink outline-none focus:border-accent focus:ring-1 focus:ring-accent';
	const labelClass = 'block text-sm font-medium text-body';
	const secondaryButtonClass =
		'rounded-md border border-line px-3 py-2 text-sm font-semibold text-body hover:bg-surface-alt disabled:cursor-not-allowed disabled:opacity-60';

	function emptyProjectForm(): ProjectForm {
		return {
			description: '',
			imageUrl: '',
			liveUrl: '',
			repoUrl: '',
			techStack: '',
			title: ''
		};
	}

	function emptyProductForm(): ProductForm {
		return {
			buyUrl: '',
			description: '',
			features: '',
			imageUrl: '',
			price: '',
			tag: '',
			title: ''
		};
	}

	function emptyExperienceForm(): ExperienceForm {
		return {
			company: '',
			description: '',
			duration: '',
			role: ''
		};
	}

	function splitList(value: string): string[] {
		return value
			.split(',')
			.map((item) => item.trim())
			.filter((item) => item.length > 0);
	}

	function listToText(value: string[] | null | undefined): string {
		return (value ?? []).join(', ');
	}

	function setSuccess(message: string): void {
		notice = message;
		error = null;
	}

	function setError(err: unknown, fallback: string): void {
		error = err instanceof ApiError ? err.message : fallback;
		notice = null;
	}

	function replaceById<T extends { id: number }>(items: T[], nextItem: T): T[] {
		if (items.some((item) => item.id === nextItem.id)) {
			return items.map((item) => (item.id === nextItem.id ? nextItem : item));
		}

		return [nextItem, ...items];
	}

	async function loadDashboard(): Promise<void> {
		isLoading = true;
		error = null;

		try {
			const [projectList, productList, experienceList, messageList] = await Promise.all([
				listProjects(),
				listProducts(),
				listExperiences(),
				listMessages()
			]);

			projects = projectList;
			products = productList;
			experiences = experienceList;
			messages = messageList;
		} catch (err) {
			setError(err, 'Failed to load admin dashboard.');
		} finally {
			isLoading = false;
		}
	}

	async function logout(): Promise<void> {
		clearAuthToken();
		await goto(resolve('/login'));
	}

	async function handleImageUpload(event: Event, target: 'project' | 'product'): Promise<void> {
		const input = event.currentTarget as HTMLInputElement;
		const file = input.files?.[0];

		if (!file) {
			return;
		}

		uploadingFor = target;
		error = null;

		try {
			const result = await uploadFile(file);

			if (target === 'project') {
				projectForm.imageUrl = result.url;
			} else {
				productForm.imageUrl = result.url;
			}

			input.value = '';
			setSuccess('Image uploaded.');
		} catch (err) {
			setError(err, 'Failed to upload image.');
		} finally {
			uploadingFor = null;
		}
	}

	function editProject(project: Project): void {
		activeTab = 'projects';
		editingProjectId = project.id;
		projectForm = {
			description: project.description,
			imageUrl: project.imageUrl,
			liveUrl: project.liveUrl,
			repoUrl: project.repoUrl,
			techStack: listToText(project.techStack),
			title: project.title
		};
	}

	function resetProjectForm(): void {
		editingProjectId = null;
		projectForm = emptyProjectForm();
	}

	async function saveProject(): Promise<void> {
		savingResource = 'projects';
		error = null;

		const payload: Omit<Project, 'id'> = {
			description: projectForm.description.trim(),
			imageUrl: projectForm.imageUrl.trim(),
			liveUrl: projectForm.liveUrl.trim(),
			repoUrl: projectForm.repoUrl.trim(),
			techStack: splitList(projectForm.techStack),
			title: projectForm.title.trim()
		};

		try {
			const saved =
				editingProjectId == null
					? await createResource('projects', payload)
					: await updateResource('projects', editingProjectId, payload);

			projects = replaceById(projects, saved);
			resetProjectForm();
			setSuccess('Project saved.');
		} catch (err) {
			setError(err, 'Failed to save project.');
		} finally {
			savingResource = null;
		}
	}

	async function removeProject(project: Project): Promise<void> {
		if (!confirm(`Delete project "${project.title}"?`)) {
			return;
		}

		deletingKey = `projects:${project.id}`;
		error = null;

		try {
			await deleteResource('projects', project.id);
			projects = projects.filter((item) => item.id !== project.id);
			setSuccess('Project deleted.');
		} catch (err) {
			setError(err, 'Failed to delete project.');
		} finally {
			deletingKey = null;
		}
	}

	function editProduct(product: Product): void {
		activeTab = 'products';
		editingProductId = product.id;
		productForm = {
			buyUrl: product.buyUrl,
			description: product.description,
			features: listToText(product.features),
			imageUrl: product.imageUrl,
			price: product.price,
			tag: product.tag,
			title: product.title
		};
	}

	function resetProductForm(): void {
		editingProductId = null;
		productForm = emptyProductForm();
	}

	async function saveProduct(): Promise<void> {
		savingResource = 'products';
		error = null;

		const payload: Omit<Product, 'id'> = {
			buyUrl: productForm.buyUrl.trim(),
			description: productForm.description.trim(),
			features: splitList(productForm.features),
			imageUrl: productForm.imageUrl.trim(),
			price: productForm.price.trim(),
			tag: productForm.tag.trim(),
			title: productForm.title.trim()
		};

		try {
			const saved =
				editingProductId == null
					? await createResource('products', payload)
					: await updateResource('products', editingProductId, payload);

			products = replaceById(products, saved);
			resetProductForm();
			setSuccess('Product saved.');
		} catch (err) {
			setError(err, 'Failed to save product.');
		} finally {
			savingResource = null;
		}
	}

	async function removeProduct(product: Product): Promise<void> {
		if (!confirm(`Delete product "${product.title}"?`)) {
			return;
		}

		deletingKey = `products:${product.id}`;
		error = null;

		try {
			await deleteResource('products', product.id);
			products = products.filter((item) => item.id !== product.id);
			setSuccess('Product deleted.');
		} catch (err) {
			setError(err, 'Failed to delete product.');
		} finally {
			deletingKey = null;
		}
	}

	function editExperience(experience: Experience): void {
		activeTab = 'experiences';
		editingExperienceId = experience.id;
		experienceForm = {
			company: experience.company,
			description: experience.description,
			duration: experience.duration,
			role: experience.role
		};
	}

	function resetExperienceForm(): void {
		editingExperienceId = null;
		experienceForm = emptyExperienceForm();
	}

	async function saveExperience(): Promise<void> {
		savingResource = 'experiences';
		error = null;

		const payload: Omit<Experience, 'id'> = {
			company: experienceForm.company.trim(),
			description: experienceForm.description.trim(),
			duration: experienceForm.duration.trim(),
			role: experienceForm.role.trim()
		};

		try {
			const saved =
				editingExperienceId == null
					? await createResource('experiences', payload)
					: await updateResource('experiences', editingExperienceId, payload);

			experiences = replaceById(experiences, saved);
			resetExperienceForm();
			setSuccess('Experience saved.');
		} catch (err) {
			setError(err, 'Failed to save experience.');
		} finally {
			savingResource = null;
		}
	}

	async function removeExperience(experience: Experience): Promise<void> {
		if (!confirm(`Delete experience "${experience.role}"?`)) {
			return;
		}

		deletingKey = `experiences:${experience.id}`;
		error = null;

		try {
			await deleteResource('experiences', experience.id);
			experiences = experiences.filter((item) => item.id !== experience.id);
			setSuccess('Experience deleted.');
		} catch (err) {
			setError(err, 'Failed to delete experience.');
		} finally {
			deletingKey = null;
		}
	}

	onMount(() => {
		void loadDashboard();
	});
</script>

<section class="mx-auto max-w-6xl px-4 py-10 sm:px-6 lg:px-8">
	<header
		class="mb-8 flex flex-col gap-4 border-b border-line pb-6 md:flex-row md:items-center md:justify-between"
	>
		<div>
			<p class="text-sm font-semibold uppercase tracking-[0.2em] text-success">Admin</p>
			<h1 class="mt-2 text-3xl font-semibold tracking-tight text-ink">Dashboard</h1>
		</div>
		<div class="flex flex-wrap gap-2">
			<button type="button" class={secondaryButtonClass} onclick={() => void loadDashboard()}>
				Refresh
			</button>
			<button
				type="button"
				class="rounded-md border border-danger/30 px-4 py-2 text-sm font-semibold text-danger hover:bg-danger-soft"
				onclick={logout}
			>
				Logout
			</button>
		</div>
	</header>

	{#if error}
		<p class="mb-6 rounded-md border border-danger/30 bg-danger-soft px-4 py-3 text-sm text-danger">
			{error}
		</p>
	{/if}

	{#if notice}
		<p
			class="mb-6 rounded-md border border-success/30 bg-success-soft px-4 py-3 text-sm text-success"
		>
			{notice}
		</p>
	{/if}

	<div class="mb-6 flex flex-wrap gap-2">
		{#each tabs as tab (tab.id)}
			<button
				type="button"
				class={[
					'rounded-md px-3 py-2 text-sm font-semibold',
					activeTab === tab.id ? 'bg-success text-inverse' : 'border border-line text-body'
				]}
				onclick={() => (activeTab = tab.id)}
			>
				{tab.label}
			</button>
		{/each}
	</div>

	{#if isLoading}
		<p class="text-muted">Loading admin data...</p>
	{:else if activeTab === 'overview'}
		<div class="grid gap-4 md:grid-cols-4">
			<button
				type="button"
				class="rounded-lg border border-line bg-surface p-5 text-left hover:border-success"
				onclick={() => (activeTab = 'projects')}
			>
				<p class="text-sm text-muted">Projects</p>
				<p class="mt-2 text-3xl font-semibold text-ink">{projects.length}</p>
			</button>
			<button
				type="button"
				class="rounded-lg border border-line bg-surface p-5 text-left hover:border-success"
				onclick={() => (activeTab = 'products')}
			>
				<p class="text-sm text-muted">Products</p>
				<p class="mt-2 text-3xl font-semibold text-ink">{products.length}</p>
			</button>
			<button
				type="button"
				class="rounded-lg border border-line bg-surface p-5 text-left hover:border-success"
				onclick={() => (activeTab = 'experiences')}
			>
				<p class="text-sm text-muted">Experience</p>
				<p class="mt-2 text-3xl font-semibold text-ink">{experiences.length}</p>
			</button>
			<button
				type="button"
				class="rounded-lg border border-line bg-surface p-5 text-left hover:border-success"
				onclick={() => (activeTab = 'messages')}
			>
				<p class="text-sm text-muted">Messages</p>
				<p class="mt-2 text-3xl font-semibold text-ink">{messages.length}</p>
			</button>
		</div>

		<div class="mt-6 rounded-lg border border-line bg-surface p-5">
			<h2 class="text-lg font-semibold text-ink">Public routes</h2>
			<div class="mt-4 flex flex-wrap gap-3">
				{#each resourceLinks as link (link.href)}
					<a
						href={resolve(link.href)}
						class="rounded-md border border-line px-3 py-2 text-sm font-semibold text-body hover:bg-surface-alt"
					>
						{link.label}
					</a>
				{/each}
			</div>
		</div>
	{:else if activeTab === 'projects'}
		<div class="grid gap-6 lg:grid-cols-[minmax(0,0.85fr)_minmax(0,1.15fr)]">
			<form
				class="rounded-lg border border-line bg-surface p-5"
				onsubmit={(event) => {
					event.preventDefault();
					void saveProject();
				}}
			>
				<h2 class="text-lg font-semibold text-ink">
					{editingProjectId == null ? 'New project' : 'Edit project'}
				</h2>
				<div class="mt-4 space-y-4">
					<label class={labelClass}>
						Title
						<input class={inputClass} required bind:value={projectForm.title} />
					</label>
					<label class={labelClass}>
						Description
						<textarea
							class={`${inputClass} min-h-28 resize-y`}
							required
							bind:value={projectForm.description}
						></textarea>
					</label>
					<label class={labelClass}>
						Image URL
						<input class={inputClass} bind:value={projectForm.imageUrl} />
					</label>
					<label class={labelClass}>
						Upload image
						<input
							class={`${inputClass} file:mr-3 file:rounded-md file:border-0 file:bg-surface-alt file:px-3 file:py-1.5 file:text-sm file:font-semibold file:text-body`}
							type="file"
							accept="image/jpeg,image/png,image/webp,image/gif"
							disabled={uploadingFor !== null}
							onchange={(event) => void handleImageUpload(event, 'project')}
						/>
					</label>
					<label class={labelClass}>
						Tech stack
						<input class={inputClass} bind:value={projectForm.techStack} />
					</label>
					<label class={labelClass}>
						Repository URL
						<input class={inputClass} type="url" bind:value={projectForm.repoUrl} />
					</label>
					<label class={labelClass}>
						Live URL
						<input class={inputClass} type="url" bind:value={projectForm.liveUrl} />
					</label>
				</div>
				<div class="mt-5 flex flex-wrap gap-2">
					<button
						type="submit"
						class="rounded-md bg-success px-4 py-2 text-sm font-semibold text-inverse hover:bg-success-hover disabled:cursor-not-allowed disabled:opacity-60"
						disabled={savingResource !== null || uploadingFor !== null}
					>
						{savingResource === 'projects' ? 'Saving...' : 'Save project'}
					</button>
					<button type="button" class={secondaryButtonClass} onclick={resetProjectForm}>
						Cancel
					</button>
				</div>
			</form>

			<div class="space-y-4">
				{#if projects.length === 0}
					<p class="rounded-md border border-line bg-surface-alt px-4 py-3 text-sm text-body">
						No projects are published yet.
					</p>
				{:else}
					{#each projects as project (project.id)}
						<article class="rounded-lg border border-line bg-surface p-5">
							<div class="flex flex-col gap-4 sm:flex-row sm:items-start sm:justify-between">
								<div>
									<h3 class="text-lg font-semibold text-ink">{project.title}</h3>
									<p class="mt-2 text-sm leading-6 text-body">{project.description}</p>
									{#if project.techStack?.length > 0}
										<div class="mt-3 flex flex-wrap gap-2">
											{#each project.techStack as tech (tech)}
												<span
													class="rounded-full bg-surface-alt px-2.5 py-1 text-xs font-medium text-body"
												>
													{tech}
												</span>
											{/each}
										</div>
									{/if}
								</div>
								<div class="flex shrink-0 gap-2">
									<button
										type="button"
										class={secondaryButtonClass}
										onclick={() => editProject(project)}
									>
										Edit
									</button>
									<button
										type="button"
										class="rounded-md border border-danger/30 px-3 py-2 text-sm font-semibold text-danger hover:bg-danger-soft disabled:cursor-not-allowed disabled:opacity-60"
										disabled={deletingKey === `projects:${project.id}`}
										onclick={() => void removeProject(project)}
									>
										Delete
									</button>
								</div>
							</div>
						</article>
					{/each}
				{/if}
			</div>
		</div>
	{:else if activeTab === 'products'}
		<div class="grid gap-6 lg:grid-cols-[minmax(0,0.85fr)_minmax(0,1.15fr)]">
			<form
				class="rounded-lg border border-line bg-surface p-5"
				onsubmit={(event) => {
					event.preventDefault();
					void saveProduct();
				}}
			>
				<h2 class="text-lg font-semibold text-ink">
					{editingProductId == null ? 'New product' : 'Edit product'}
				</h2>
				<div class="mt-4 space-y-4">
					<label class={labelClass}>
						Title
						<input class={inputClass} required bind:value={productForm.title} />
					</label>
					<label class={labelClass}>
						Description
						<textarea
							class={`${inputClass} min-h-28 resize-y`}
							required
							bind:value={productForm.description}
						></textarea>
					</label>
					<div class="grid gap-4 sm:grid-cols-2">
						<label class={labelClass}>
							Price
							<input class={inputClass} bind:value={productForm.price} />
						</label>
						<label class={labelClass}>
							Tag
							<input class={inputClass} bind:value={productForm.tag} />
						</label>
					</div>
					<label class={labelClass}>
						Image URL
						<input class={inputClass} bind:value={productForm.imageUrl} />
					</label>
					<label class={labelClass}>
						Upload image
						<input
							class={`${inputClass} file:mr-3 file:rounded-md file:border-0 file:bg-surface-alt file:px-3 file:py-1.5 file:text-sm file:font-semibold file:text-body`}
							type="file"
							accept="image/jpeg,image/png,image/webp,image/gif"
							disabled={uploadingFor !== null}
							onchange={(event) => void handleImageUpload(event, 'product')}
						/>
					</label>
					<label class={labelClass}>
						Features
						<input class={inputClass} bind:value={productForm.features} />
					</label>
					<label class={labelClass}>
						Buy URL
						<input class={inputClass} type="url" bind:value={productForm.buyUrl} />
					</label>
				</div>
				<div class="mt-5 flex flex-wrap gap-2">
					<button
						type="submit"
						class="rounded-md bg-success px-4 py-2 text-sm font-semibold text-inverse hover:bg-success-hover disabled:cursor-not-allowed disabled:opacity-60"
						disabled={savingResource !== null || uploadingFor !== null}
					>
						{savingResource === 'products' ? 'Saving...' : 'Save product'}
					</button>
					<button type="button" class={secondaryButtonClass} onclick={resetProductForm}>
						Cancel
					</button>
				</div>
			</form>

			<div class="space-y-4">
				{#if products.length === 0}
					<p class="rounded-md border border-line bg-surface-alt px-4 py-3 text-sm text-body">
						No products are published yet.
					</p>
				{:else}
					{#each products as product (product.id)}
						<article class="rounded-lg border border-line bg-surface p-5">
							<div class="flex flex-col gap-4 sm:flex-row sm:items-start sm:justify-between">
								<div>
									<div class="flex flex-wrap items-center gap-2">
										<h3 class="text-lg font-semibold text-ink">{product.title}</h3>
										{#if product.tag}
											<span
												class="rounded-full bg-success-soft px-2.5 py-1 text-xs font-semibold text-success"
											>
												{product.tag}
											</span>
										{/if}
									</div>
									<p class="mt-2 text-sm leading-6 text-body">{product.description}</p>
									{#if product.price}
										<p class="mt-3 text-sm font-semibold text-ink">{product.price}</p>
									{/if}
									{#if product.features?.length > 0}
										<p class="mt-3 text-sm text-muted">{product.features.join(', ')}</p>
									{/if}
								</div>
								<div class="flex shrink-0 gap-2">
									<button
										type="button"
										class={secondaryButtonClass}
										onclick={() => editProduct(product)}
									>
										Edit
									</button>
									<button
										type="button"
										class="rounded-md border border-danger/30 px-3 py-2 text-sm font-semibold text-danger hover:bg-danger-soft disabled:cursor-not-allowed disabled:opacity-60"
										disabled={deletingKey === `products:${product.id}`}
										onclick={() => void removeProduct(product)}
									>
										Delete
									</button>
								</div>
							</div>
						</article>
					{/each}
				{/if}
			</div>
		</div>
	{:else if activeTab === 'experiences'}
		<div class="grid gap-6 lg:grid-cols-[minmax(0,0.85fr)_minmax(0,1.15fr)]">
			<form
				class="rounded-lg border border-line bg-surface p-5"
				onsubmit={(event) => {
					event.preventDefault();
					void saveExperience();
				}}
			>
				<h2 class="text-lg font-semibold text-ink">
					{editingExperienceId == null ? 'New experience' : 'Edit experience'}
				</h2>
				<div class="mt-4 space-y-4">
					<label class={labelClass}>
						Role
						<input class={inputClass} required bind:value={experienceForm.role} />
					</label>
					<label class={labelClass}>
						Company
						<input class={inputClass} required bind:value={experienceForm.company} />
					</label>
					<label class={labelClass}>
						Duration
						<input class={inputClass} required bind:value={experienceForm.duration} />
					</label>
					<label class={labelClass}>
						Description
						<textarea
							class={`${inputClass} min-h-28 resize-y`}
							required
							bind:value={experienceForm.description}
						></textarea>
					</label>
				</div>
				<div class="mt-5 flex flex-wrap gap-2">
					<button
						type="submit"
						class="rounded-md bg-success px-4 py-2 text-sm font-semibold text-inverse hover:bg-success-hover disabled:cursor-not-allowed disabled:opacity-60"
						disabled={savingResource !== null}
					>
						{savingResource === 'experiences' ? 'Saving...' : 'Save experience'}
					</button>
					<button type="button" class={secondaryButtonClass} onclick={resetExperienceForm}>
						Cancel
					</button>
				</div>
			</form>

			<div class="space-y-4">
				{#if experiences.length === 0}
					<p class="rounded-md border border-line bg-surface-alt px-4 py-3 text-sm text-body">
						No experience entries are published yet.
					</p>
				{:else}
					{#each experiences as experience (experience.id)}
						<article class="rounded-lg border border-line bg-surface p-5">
							<div class="flex flex-col gap-4 sm:flex-row sm:items-start sm:justify-between">
								<div>
									<p class="text-sm font-medium text-success">{experience.duration}</p>
									<h3 class="mt-2 text-lg font-semibold text-ink">{experience.role}</h3>
									<p class="text-sm text-muted">{experience.company}</p>
									<p class="mt-3 text-sm leading-6 text-body">{experience.description}</p>
								</div>
								<div class="flex shrink-0 gap-2">
									<button
										type="button"
										class={secondaryButtonClass}
										onclick={() => editExperience(experience)}
									>
										Edit
									</button>
									<button
										type="button"
										class="rounded-md border border-danger/30 px-3 py-2 text-sm font-semibold text-danger hover:bg-danger-soft disabled:cursor-not-allowed disabled:opacity-60"
										disabled={deletingKey === `experiences:${experience.id}`}
										onclick={() => void removeExperience(experience)}
									>
										Delete
									</button>
								</div>
							</div>
						</article>
					{/each}
				{/if}
			</div>
		</div>
	{:else}
		<div class="space-y-4">
			{#if messages.length === 0}
				<p class="rounded-md border border-line bg-surface-alt px-4 py-3 text-sm text-body">
					No contact messages yet.
				</p>
			{:else}
				{#each messages as message (message.id)}
					<article class="rounded-lg border border-line bg-surface p-5">
						<div class="flex flex-col gap-1 sm:flex-row sm:items-center sm:justify-between">
							<h2 class="font-semibold text-ink">{message.name}</h2>
							<p class="text-sm text-muted">{message.email}</p>
						</div>
						<p class="mt-3 text-sm leading-6 text-body">{message.content}</p>
						<p class="mt-4 text-xs text-subtle">{message.created_at}</p>
					</article>
				{/each}
			{/if}
		</div>
	{/if}
</section>
