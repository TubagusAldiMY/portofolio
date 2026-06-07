import { sveltekit } from '@sveltejs/kit/vite';
import tailwindcss from '@tailwindcss/vite';
import { defineConfig, loadEnv } from 'vite';

export default defineConfig(({ mode }) => {
	const env = loadEnv(mode, process.cwd(), '');
	const apiBase = (env.PUBLIC_API_URL || 'http://localhost:8000').replace(/\/+$/, '');

	return {
		plugins: [sveltekit(), tailwindcss()],
		server: {
			proxy: {
				'/api': {
					changeOrigin: true,
					target: apiBase
				},
				'/uploads': {
					changeOrigin: true,
					target: apiBase
				}
			}
		}
	};
});
