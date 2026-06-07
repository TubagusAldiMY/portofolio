// src/directives/scroll-animation.js
const animatedScrollObserver = new IntersectionObserver(
    (entries, observer) => {
        entries.forEach((entry) => {
            if (entry.isIntersecting) {
                entry.target.classList.add('enter');
                observer.unobserve(entry.target);
            }
        });
    }
);

export default {
    mounted(el) {
        el.classList.add('before-enter');
        animatedScrollObserver.observe(el);
    }
}