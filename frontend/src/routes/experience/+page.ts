import { listExperiences } from '$lib/api/public';
import type { PageLoad } from './$types';

export const load: PageLoad = async ({ fetch }) => {
	try {
		const experiences = await listExperiences(fetch);
		return { experiences, error: null };
	} catch {
		return { experiences: [], error: 'Failed to load experience.' };
	}
};
