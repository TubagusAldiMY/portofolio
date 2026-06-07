<script setup>
import { ref, onMounted, onUnmounted, defineAsyncComponent } from 'vue';

// Lazy load heavy components to reduce initial bundle size
const CardSwap = defineAsyncComponent(() => import('./CardSwap.vue'));
const PixelBlast = defineAsyncComponent(() => import('./PixelBlast.vue'));

const cardWidth = ref(260);
const cardHeight = ref(340);
const cardDist = ref(40);
const cardVertDist = ref(40);

const showCardSwap = ref(true);
const showPixelBlast = ref(true);

const updateDimensions = () => {
  const w = window.innerWidth;
  showCardSwap.value = w >= 800;
  showPixelBlast.value = w >= 768;

  if (w < 640) {
    cardWidth.value = 240;
    cardHeight.value = 320;
    cardDist.value = 30;
    cardVertDist.value = 30;
  } else if (w < 1024) {
    cardWidth.value = 280;
    cardHeight.value = 360;
    cardDist.value = 40;
    cardVertDist.value = 40;
  } else if (w < 1440) {
    cardWidth.value = 340;
    cardHeight.value = 440;
    cardDist.value = 50;
    cardVertDist.value = 50;
  } else {
    cardWidth.value = 400;
    cardHeight.value = 520;
    cardDist.value = 60;
    cardVertDist.value = 60;
  }
};

onMounted(() => {
  updateDimensions();
  window.addEventListener('resize', updateDimensions);
});

onUnmounted(() => {
  window.removeEventListener('resize', updateDimensions);
});
</script>

<template>
  <section
      id="home"
      class="relative min-h-screen flex items-center pt-20 sm:pt-24 pb-16 overflow-hidden bg-[#0d1117]"
      v-scroll-animation
  >
    <div class="absolute inset-0 z-0">
      <Suspense v-if="showPixelBlast">
        <PixelBlast
            variant="circle"
            :pixel-size="6"
            color="#2e527d"
            :pattern-scale="3"
            :pattern-density="1.2"
            :pixel-size-jitter="0.5"
            :enable-ripples="true"
            :ripple-speed="0.4"
            :ripple-thickness="0.12"
            :ripple-intensity-scale="1.5"
            :liquid="true"
            :liquid-strength="0.12"
            :liquid-radius="1.2"
            :liquid-wobble-speed="5"
            :speed="0.6"
            :edge-fade="0.25"
            :transparent="true"
        />
        <template #fallback>
          <div class="absolute inset-0 bg-[#0d1117] animate-pulse"></div>
        </template>
      </Suspense>
    </div>

    <div class="absolute inset-0 z-0 bg-gradient-to-r from-[#0d1117] via-[#0d1117]/90 to-transparent pointer-events-none"></div>

    <div class="absolute inset-0 z-0 bg-gradient-to-t from-[#0d1117] via-transparent to-transparent pointer-events-none"></div>

    <div class="container mx-auto px-4 sm:px-6 lg:px-8 relative z-10 pointer-events-none">
      <div class="flex flex-col lg:flex-row items-center justify-between gap-10 lg:gap-12">

        <div class="w-full lg:w-1/2 text-center lg:text-left order-2 lg:order-1 pointer-events-auto">
          <div class="inline-block px-3 py-1 mb-4 sm:mb-5 text-xs font-semibold tracking-wide text-[#58a6ff] uppercase bg-[#388bfd]/10 rounded-full border border-[#388bfd]/30 backdrop-blur-sm">
            Available for hire
          </div>

          <h1 class="text-3xl sm:text-4xl lg:text-5xl font-extrabold text-white leading-tight mb-4 sm:mb-6 tracking-tight drop-shadow-md">
            Membangun masa depan <br />
            <span class="text-gradient-purple">Machine Learning</span> & Web.
          </h1>

          <p class="text-[#c9d1d9] text-base sm:text-lg lg:text-xl mb-6 sm:mb-8 lg:mb-10 max-w-2xl mx-auto lg:mx-0 leading-relaxed drop-shadow-sm font-medium">
            Saya Tubagus Aldi. Saya menggabungkan kecerdasan buatan dengan
            pengembangan web modern untuk menciptakan solusi yang efisien.
          </p>

          <div class="flex flex-col sm:flex-row justify-center lg:justify-start gap-3 sm:gap-4">
            <a href="#projects" class="px-7 py-3 sm:px-8 sm:py-4 text-sm sm:text-base font-bold text-white bg-white/10 border border-white/20 rounded-lg hover:bg-white/20 hover:border-white/40 transition-all duration-300 backdrop-blur-md">
              Lihat Proyek
            </a>
          </div>
        </div>

        <div class="w-full lg:w-1/2 flex justify-center items-center order-1 lg:order-2 mb-8 lg:mb-0 pointer-events-auto">
          <div v-if="showCardSwap" class="w-full relative flex justify-center items-center h-[420px] lg:h-[500px]">
            <Suspense>
              <CardSwap
                  :width="cardWidth"
                  :height="cardHeight"
                  :cardDistance="cardDist"
                  :verticalDistance="cardVertDist"
                  :skewAmount="2"
                  :delay="3000"
                  :pauseOnHover="false"
              >
                <template #card-0>
                  <div class="w-full h-full bg-[#0d1117] p-4 sm:p-6 font-mono text-[10px] sm:text-xs text-gray-300 flex flex-col border border-[#30363d] rounded-2xl shadow-[0_0_30px_rgba(88,166,255,0.1)]">
                    <div class="flex items-center gap-2 mb-4 sm:mb-6 border-b border-[#30363d] pb-2 sm:pb-3">
                      <div class="w-3 h-3 rounded-full bg-[#ff5f56]"></div>
                      <div class="w-3 h-3 rounded-full bg-[#ffbd2e]"></div>
                      <div class="w-3 h-3 rounded-full bg-[#27c93f]"></div>
                      <span class="ml-auto text-gray-500 opacity-60 font-sans text-[10px] sm:text-sm">main.go</span>
                    </div>

                    <div class="space-y-2 sm:space-y-3 leading-relaxed opacity-90">
                      <p><span class="text-[#ff7b72]">package</span> main</p>

                      <p class="mt-2"><span class="text-[#ff7b72]">import</span> <span class="text-[#a5d6ff]">"fmt"</span></p>

                      <p class="mt-2 sm:mt-4"><span class="text-[#ff7b72]">func</span> <span class="text-[#d2a8ff]">main</span>() {</p>

                      <p class="pl-4 text-[#8b949e]">// Print greeting</p>
                      <p class="pl-4">
                        fmt.Println(<span class="text-[#a5d6ff]">"Hello World"</span>)
                      </p>
                      <p class="pl-4">
                        fmt.Println(<span class="text-[#a5d6ff]">"from Tubagus Aldi"</span>)
                      </p>

                      <p>}</p>

                      <br class="hidden sm:block"/>
                      <p class="text-[#8b949e]"># Go is awesome! 🐹</p>
                    </div>
                  </div>
                </template>

                <template #card-1>
                  <div class="w-full h-full bg-[#0d1117] p-4 sm:p-6 font-mono text-[10px] sm:text-xs text-gray-300 flex flex-col border border-[#30363d] rounded-2xl shadow-[0_0_30px_rgba(88,166,255,0.1)]">
                    <div class="flex items-center gap-2 mb-4 sm:mb-6 border-b border-[#30363d] pb-2 sm:pb-3">
                      <div class="w-3 h-3 rounded-full bg-[#ff5f56]"></div>
                      <div class="w-3 h-3 rounded-full bg-[#ffbd2e]"></div>
                      <div class="w-3 h-3 rounded-full bg-[#27c93f]"></div>
                      <span class="ml-auto text-gray-500 opacity-60 font-sans text-[10px] sm:text-sm">train.py</span>
                    </div>
                    <div class="space-y-2 sm:space-y-3 leading-relaxed opacity-90">
                      <p><span class="text-[#ff7b72]">import</span> tensorflow <span class="text-[#ff7b72]">as</span> tf</p>
                      <p class="mt-1 sm:mt-2"><span class="text-[#ff7b72]">def</span> <span class="text-[#d2a8ff]">build_model</span>():</p>
                      <p class="pl-4 text-[#8b949e]"># Define architecture</p>
                      <p class="pl-4">model = tf.keras.Sequential([</p>
                      <p class="pl-8">tf.keras.layers.Dense(<span class="text-[#79c0ff]">256</span>, activation=<span class="text-[#a5d6ff]">'relu'</span>),</p>
                      <p class="pl-8">tf.keras.layers.Dropout(<span class="text-[#79c0ff]">0.3</span>),</p>
                      <p class="pl-8">tf.keras.layers.Dense(<span class="text-[#79c0ff]">1</span>, activation=<span class="text-[#a5d6ff]">'sigmoid'</span>)</p>
                      <p class="pl-4">])</p>
                      <p class="pl-4"><span class="text-[#ff7b72]">return</span> model</p>
                      <br />
                      <p class="text-[#8b949e]"># Let's go! 🚀</p>
                    </div>
                  </div>
                </template>

                <template #card-2>
                  <div class="w-full h-full bg-gradient-to-br from-[#161b22] to-[#0d1117] flex flex-col items-center justify-center p-6 sm:p-8 border border-[#30363d] rounded-2xl">
                    <div class="grid grid-cols-2 gap-4 sm:gap-5 w-full mb-4 sm:mb-6">
                      <div class="aspect-square bg-[#21262d]/50 rounded-xl border border-[#30363d] flex items-center justify-center backdrop-blur-sm hover:border-[#41b883] transition-colors">
                        <span class="text-2xl sm:text-3xl font-bold text-[#41b883]">Vue</span>
                      </div>
                      <div class="aspect-square bg-[#21262d]/50 rounded-xl border border-[#30363d] flex items-center justify-center backdrop-blur-sm hover:border-[#3572A5] transition-colors">
                        <span class="text-2xl sm:text-3xl font-bold text-[#3572A5]">Py</span>
                      </div>
                      <div class="aspect-square bg-[#21262d]/50 rounded-xl border border-[#30363d] flex items-center justify-center backdrop-blur-sm hover:border-[#00ADD8] transition-colors">
                        <span class="text-2xl sm:text-3xl font-bold text-[#00ADD8]">Go</span>
                      </div>
                      <div class="aspect-square bg-[#21262d]/50 rounded-xl border border-[#30363d] flex items-center justify-center backdrop-blur-sm hover:border-[#F7DF1E] transition-colors">
                        <span class="text-xl sm:text-2xl font-bold text-white">JS</span>
                      </div>
                    </div>
                    <span class="text-[#8b949e] font-medium uppercase tracking-widest text-xs sm:text-sm">Core Stack</span>
                  </div>
                </template>
              </CardSwap>
              <template #fallback>
                <div class="w-full h-full flex items-center justify-center">
                  <div class="w-[260px] h-[340px] bg-[#161b22]/50 border border-[#30363d] rounded-2xl animate-pulse"></div>
                </div>
              </template>
            </Suspense>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>