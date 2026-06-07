import { error } from '@sveltejs/kit';
import { getProject } from '$lib/api/public';
import type { PageLoad } from './$types';

export const load: PageLoad = async ({ params, fetch }) => {
	const id = Number(params.id);
	if (!Number.isFinite(id) || id <= 0) {
		error(404, 'Project not found');
	}

	try {
		const project = await getProject(id, fetch);
		return { project };
	} catch {
		error(404, 'Project not found');
	}
};
