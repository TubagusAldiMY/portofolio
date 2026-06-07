import { redirect } from '@sveltejs/kit';

import { AUTH_COOKIE_NAME, isJwtExpired } from '$lib/auth/token';
import type { LayoutServerLoad } from './$types';

export const load: LayoutServerLoad = ({ cookies, url }) => {
	const token = cookies.get(AUTH_COOKIE_NAME);

	if (!token || isJwtExpired(token)) {
		throw redirect(303, `/login?redirectTo=${encodeURIComponent(url.pathname + url.search)}`);
	}

	return {};
};
