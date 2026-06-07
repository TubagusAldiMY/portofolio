import { listProducts } from '$lib/api/public';
import type { PageLoad } from './$types';

export const load: PageLoad = async ({ fetch }) => {
	try {
		const products = await listProducts(fetch);
		return { products, error: null };
	} catch {
		return { products: [], error: 'Failed to load products.' };
	}
};
