export function scrollAnimation(node: Element) {
	node.classList.add('before-enter');

	const observer = new IntersectionObserver(
		([entry]) => {
			if (entry.isIntersecting) {
				node.classList.add('enter');
				observer.disconnect();
			}
		},
		{ threshold: 0.1 }
	);

	observer.observe(node);

	return {
		destroy() {
			observer.disconnect();
		}
	};
}
