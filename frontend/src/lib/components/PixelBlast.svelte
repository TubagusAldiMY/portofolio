<script lang="ts">
	import { onMount } from 'svelte';

	type Variant = 'square' | 'circle' | 'triangle' | 'diamond';

	type Props = {
		variant?: Variant;
		pixelSize?: number;
		color?: string;
		className?: string;
		antialias?: boolean;
		patternScale?: number;
		patternDensity?: number;
		pixelSizeJitter?: number;
		enableRipples?: boolean;
		rippleIntensityScale?: number;
		rippleThickness?: number;
		rippleSpeed?: number;
		autoPauseOffscreen?: boolean;
		speed?: number;
		transparent?: boolean;
		edgeFade?: number;
	};

	const SHAPE_MAP: Record<Variant, number> = {
		circle: 1,
		diamond: 3,
		square: 0,
		triangle: 2
	};
	const MAX_CLICKS = 10;

	const VERTEX_SRC = `
void main() {
	gl_Position = vec4(position, 1.0);
}
`;

	const FRAGMENT_SRC = `
precision highp float;

uniform vec3  uColor;
uniform vec2  uResolution;
uniform float uTime;
uniform float uPixelSize;
uniform float uScale;
uniform float uDensity;
uniform float uPixelJitter;
uniform int   uEnableRipples;
uniform float uRippleSpeed;
uniform float uRippleThickness;
uniform float uRippleIntensity;
uniform float uEdgeFade;
uniform int   uShapeType;
uniform vec2  uClickPos  [${MAX_CLICKS}];
uniform float uClickTimes[${MAX_CLICKS}];

out vec4 fragColor;

float Bayer2(vec2 a) {
	a = floor(a);
	return fract(a.x / 2. + a.y * a.y * .75);
}
#define Bayer4(a) (Bayer2(.5*(a))*0.25 + Bayer2(a))
#define Bayer8(a) (Bayer4(.5*(a))*0.25 + Bayer2(a))

float hash11(float n){ return fract(sin(n)*43758.5453); }

float vnoise(vec3 p){
	vec3 ip = floor(p);
	vec3 fp = fract(p);
	float n000 = hash11(dot(ip + vec3(0.0,0.0,0.0), vec3(1.0,57.0,113.0)));
	float n100 = hash11(dot(ip + vec3(1.0,0.0,0.0), vec3(1.0,57.0,113.0)));
	float n010 = hash11(dot(ip + vec3(0.0,1.0,0.0), vec3(1.0,57.0,113.0)));
	float n110 = hash11(dot(ip + vec3(1.0,1.0,0.0), vec3(1.0,57.0,113.0)));
	float n001 = hash11(dot(ip + vec3(0.0,0.0,1.0), vec3(1.0,57.0,113.0)));
	float n101 = hash11(dot(ip + vec3(1.0,0.0,1.0), vec3(1.0,57.0,113.0)));
	float n011 = hash11(dot(ip + vec3(0.0,1.0,1.0), vec3(1.0,57.0,113.0)));
	float n111 = hash11(dot(ip + vec3(1.0,1.0,1.0), vec3(1.0,57.0,113.0)));
	vec3 w = fp*fp*fp*(fp*(fp*6.0-15.0)+10.0);
	float x00 = mix(n000, n100, w.x);
	float x10 = mix(n010, n110, w.x);
	float x01 = mix(n001, n101, w.x);
	float x11 = mix(n011, n111, w.x);
	float y0  = mix(x00, x10, w.y);
	float y1  = mix(x01, x11, w.y);
	return mix(y0, y1, w.z) * 2.0 - 1.0;
}

float fbm2(vec2 uv, float t){
	vec3 p = vec3(uv * uScale, t);
	float amp = 1.0;
	float freq = 1.0;
	float sum = 1.0;
	for (int i = 0; i < 5; ++i){
		sum  += amp * vnoise(p * freq);
		freq *= 1.25;
	}
	return sum * 0.5 + 0.5;
}

float maskCircle(vec2 p, float cov){
	float r = sqrt(cov) * .25;
	float d = length(p - 0.5) - r;
	float aa = 0.5 * fwidth(d);
	return cov * (1.0 - smoothstep(-aa, aa, d * 2.0));
}

float maskTriangle(vec2 p, vec2 id, float cov){
	bool flip = mod(id.x + id.y, 2.0) > 0.5;
	if (flip) p.x = 1.0 - p.x;
	float r = sqrt(cov);
	float d  = p.y - r*(1.0 - p.x);
	float aa = fwidth(d);
	return cov * clamp(0.5 - d/aa, 0.0, 1.0);
}

float maskDiamond(vec2 p, float cov){
	float r = sqrt(cov) * 0.564;
	return step(abs(p.x - 0.49) + abs(p.y - 0.49), r);
}

void main(){
	float pixelSize = uPixelSize;
	vec2 fragCoord = gl_FragCoord.xy - uResolution * .5;
	float aspectRatio = uResolution.x / uResolution.y;
	vec2 pixelId = floor(fragCoord / pixelSize);
	vec2 pixelUV = fract(fragCoord / pixelSize);
	float cellPixelSize = 8.0 * pixelSize;
	vec2 cellId = floor(fragCoord / cellPixelSize);
	vec2 cellCoord = cellId * cellPixelSize;
	vec2 uv = cellCoord / uResolution * vec2(aspectRatio, 1.0);
	float base = fbm2(uv, uTime * 0.05);
	base = base * 0.5 - 0.65;
	float feed = base + (uDensity - 0.5) * 0.3;

	if (uEnableRipples == 1) {
		for (int i = 0; i < ${MAX_CLICKS}; ++i){
			vec2 pos = uClickPos[i];
			if (pos.x < 0.0) continue;
			vec2 cuv = (((pos - uResolution * .5 - cellPixelSize * .5) / (uResolution))) * vec2(aspectRatio, 1.0);
			float t = max(uTime - uClickTimes[i], 0.0);
			float r = distance(uv, cuv);
			float waveR = uRippleSpeed * t;
			float ring  = exp(-pow((r - waveR) / uRippleThickness, 2.0));
			float atten = exp(-1.0 * t) * exp(-10.0 * r);
			feed = max(feed, ring * atten * uRippleIntensity);
		}
	}

	float bayer = Bayer8(fragCoord / uPixelSize) - 0.5;
	float bw = step(0.5, feed + bayer);
	float h = fract(sin(dot(floor(fragCoord / uPixelSize), vec2(127.1, 311.7))) * 43758.5453);
	float jitterScale = 1.0 + (h - 0.5) * uPixelJitter;
	float coverage = bw * jitterScale;
	float mask;
	if      (uShapeType == 1) mask = maskCircle(pixelUV, coverage);
	else if (uShapeType == 2) mask = maskTriangle(pixelUV, pixelId, coverage);
	else if (uShapeType == 3) mask = maskDiamond(pixelUV, coverage);
	else                      mask = coverage;

	if (uEdgeFade > 0.0) {
		vec2 norm = gl_FragCoord.xy / uResolution;
		float edge = min(min(norm.x, norm.y), min(1.0 - norm.x, 1.0 - norm.y));
		float fade = smoothstep(0.0, uEdgeFade, edge);
		mask *= fade;
	}

	fragColor = vec4(uColor, mask);
}
`;

	let {
		antialias = true,
		autoPauseOffscreen = true,
		className = '',
		color = '#2e527d',
		edgeFade = 0.25,
		enableRipples = true,
		patternDensity = 1.2,
		patternScale = 3,
		pixelSize = 6,
		pixelSizeJitter = 0.5,
		rippleIntensityScale = 1.5,
		rippleSpeed = 0.4,
		rippleThickness = 0.12,
		speed = 0.6,
		transparent = true,
		variant = 'circle'
	}: Props = $props();

	let container: HTMLDivElement | undefined;
	let cleanup: (() => void) | undefined;

	onMount(() => {
		let cancelled = false;

		const init = async () => {
			const [THREE, { EffectComposer, RenderPass }] = await Promise.all([
				import('three'),
				import('postprocessing')
			]);

			if (cancelled || !container) {
				return;
			}

			const mount = container;
			const reducedMotion = window.matchMedia('(prefers-reduced-motion: reduce)').matches;
			const canvas = document.createElement('canvas');
			const context = canvas.getContext('webgl2', { alpha: true, antialias });

			if (!context) {
				return;
			}

			const renderer = new THREE.WebGLRenderer({
				alpha: true,
				antialias,
				canvas,
				context
			});
			renderer.setPixelRatio(Math.min(window.devicePixelRatio || 1, 2));
			renderer.domElement.className = 'h-full w-full';
			renderer.setClearAlpha(transparent ? 0 : 1);
			mount.appendChild(renderer.domElement);

			const uniforms = {
				uClickPos: { value: Array.from({ length: MAX_CLICKS }, () => new THREE.Vector2(-1, -1)) },
				uClickTimes: { value: new Float32Array(MAX_CLICKS) },
				uColor: { value: new THREE.Color(color) },
				uDensity: { value: patternDensity },
				uEdgeFade: { value: edgeFade },
				uEnableRipples: { value: enableRipples && !reducedMotion ? 1 : 0 },
				uPixelJitter: { value: pixelSizeJitter },
				uPixelSize: { value: pixelSize * renderer.getPixelRatio() },
				uResolution: { value: new THREE.Vector2(0, 0) },
				uRippleIntensity: { value: rippleIntensityScale },
				uRippleSpeed: { value: rippleSpeed },
				uRippleThickness: { value: rippleThickness },
				uScale: { value: patternScale },
				uShapeType: { value: SHAPE_MAP[variant] ?? 0 },
				uTime: { value: 0 }
			};

			const scene = new THREE.Scene();
			const camera = new THREE.OrthographicCamera(-1, 1, 1, -1, 0, 1);
			const material = new THREE.ShaderMaterial({
				depthTest: false,
				depthWrite: false,
				fragmentShader: FRAGMENT_SRC,
				glslVersion: THREE.GLSL3,
				transparent: true,
				uniforms,
				vertexShader: VERTEX_SRC
			});
			const quad = new THREE.Mesh(new THREE.PlaneGeometry(2, 2), material);
			scene.add(quad);

			const composer = new EffectComposer(renderer);
			composer.addPass(new RenderPass(scene, camera));

			const resize = () => {
				const width = mount.clientWidth || 1;
				const height = mount.clientHeight || 1;
				renderer.setSize(width, height, false);
				composer.setSize(renderer.domElement.width, renderer.domElement.height);
				uniforms.uResolution.value.set(renderer.domElement.width, renderer.domElement.height);
				uniforms.uPixelSize.value = pixelSize * renderer.getPixelRatio();
			};
			resize();

			const resizeObserver = new ResizeObserver(resize);
			resizeObserver.observe(mount);

			let visible = true;
			const intersectionObserver = autoPauseOffscreen
				? new IntersectionObserver(([entry]) => {
						visible = Boolean(entry?.isIntersecting);
					})
				: undefined;
			intersectionObserver?.observe(mount);

			let clickIndex = 0;
			const mapPointer = (event: PointerEvent) => {
				const rect = renderer.domElement.getBoundingClientRect();
				const scaleX = renderer.domElement.width / rect.width;
				const scaleY = renderer.domElement.height / rect.height;
				return {
					x: (event.clientX - rect.left) * scaleX,
					y: (rect.height - (event.clientY - rect.top)) * scaleY
				};
			};
			const onPointerDown = (event: PointerEvent) => {
				const point = mapPointer(event);
				uniforms.uClickPos.value[clickIndex].set(point.x, point.y);
				uniforms.uClickTimes.value[clickIndex] = uniforms.uTime.value;
				clickIndex = (clickIndex + 1) % MAX_CLICKS;
			};
			renderer.domElement.addEventListener('pointerdown', onPointerDown, { passive: true });

			const offset = Math.random() * 1000;
			const startedAt = performance.now();
			let frame = 0;

			const animate = () => {
				if (!autoPauseOffscreen || visible) {
					const elapsed = (performance.now() - startedAt) / 1000;
					uniforms.uTime.value = offset + elapsed * (reducedMotion ? speed * 0.2 : speed);
					composer.render();
				}

				frame = requestAnimationFrame(animate);
			};
			frame = requestAnimationFrame(animate);

			cleanup = () => {
				cancelAnimationFrame(frame);
				renderer.domElement.removeEventListener('pointerdown', onPointerDown);
				intersectionObserver?.disconnect();
				resizeObserver.disconnect();
				composer.dispose();
				quad.geometry.dispose();
				material.dispose();
				renderer.dispose();

				if (renderer.domElement.parentElement === mount) {
					mount.removeChild(renderer.domElement);
				}
			};
		};

		void init();

		return () => {
			cancelled = true;
			cleanup?.();
			cleanup = undefined;
		};
	});
</script>

<div
	bind:this={container}
	class={['relative h-full w-full overflow-hidden', className]}
	aria-label="PixelBlast interactive background"
></div>
