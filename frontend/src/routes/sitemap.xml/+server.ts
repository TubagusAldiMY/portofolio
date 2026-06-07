import type { RequestHandler } from './$types';

const SITE = 'https://tubsamy.tech';

const staticRoutes = [
	{ path: '/', priority: '1.0', changefreq: 'weekly' },
	{ path: '/projects', priority: '0.9', changefreq: 'weekly' },
	{ path: '/products', priority: '0.8', changefreq: 'monthly' },
	{ path: '/experience', priority: '0.8', changefreq: 'monthly' },
	{ path: '/contact', priority: '0.7', changefreq: 'yearly' },
	{ path: '/chat', priority: '0.6', changefreq: 'monthly' }
];

export const GET: RequestHandler = () => {
	const now = new Date().toISOString().split('T')[0];

	const urls = staticRoutes
		.map(
			(r) => `
  <url>
    <loc>${SITE}${r.path}</loc>
    <lastmod>${now}</lastmod>
    <changefreq>${r.changefreq}</changefreq>
    <priority>${r.priority}</priority>
  </url>`
		)
		.join('');

	const xml = `<?xml version="1.0" encoding="UTF-8"?>
<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">${urls}
</urlset>`;

	return new Response(xml, {
		headers: {
			'Content-Type': 'application/xml',
			'Cache-Control': 'public, max-age=3600'
		}
	});
};
