import { listProjects } from '$lib/api/public';
import type { PageLoad } from './$types';

export const load: PageLoad = async ({ fetch }) => {
	try {
		const projects = await listProjects(fetch);
		return { projects, error: null };
	} catch {
		return { projects: [], error: 'Failed to load selected work from backend.' };
	}
};
