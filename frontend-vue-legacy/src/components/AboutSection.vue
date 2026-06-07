<script setup>
import { ref } from 'vue';

const techStack = ref([
  'HTML5 & CSS3', 'JavaScript', 'Vue.js', 'Tailwind CSS', 'Golang', 'Git & GitHub', 'Python',
]);
</script>

<template>
  <section id="about" class="py-20 sm:py-32 bg-[#0d1117]" v-scroll-animation>
    <div class="container mx-auto px-4 sm:px-6 lg:px-8">

      <div class="text-center mb-16">
        <h2 class="text-3xl sm:text-4xl font-extrabold text-white">Tentang Saya</h2>
        <div class="w-24 h-1 bg-[#238636] mx-auto mt-4 rounded-full"></div>
        <p class="mt-4 text-lg text-[#8b949e]">
          Sedikit lebih dalam tentang perjalanan dan keahlian saya.
        </p>
      </div>

      <div class="grid grid-cols-1 md:grid-cols-2 gap-12 items-center">

        <div class="flex justify-center order-1 md:order-none">
          <div class="relative group">
            <div class=" absolute -inset-1 bg-gradient-to-r from-[#58a6ff] to-[#bc8cff] rounded-xl blur opacity-25 group-hover:opacity-50 transition duration-1000 group-hover:duration-200"></div>
            <div class=" relative w-60 h-80 sm:w-80 sm:h-96 bg-[#161b22] rounded-xl border border-[#30363d] overflow-hidden shadow-2xl">
              <img
                  src="../assets/tubsamy.webp"
                  alt="Foto Profil"
                  class="w-full h-full object-cover opacity-90 group-hover:opacity-100 transition-opacity duration-300"
              >
            </div>
          </div>
        </div>

        <div class="order-none md:order-1">
          <h3 class="text-2xl font-bold text-white mb-4">
            Perjalanan Saya di Dunia Digital
          </h3>
          <div class="space-y-4 text-[#c9d1d9] leading-relaxed">
            <p>
              Saya adalah seorang <span class="font-semibold text-[#58a6ff]">Machine Learning Engineer</span> dan <span class="font-semibold text-[#58a6ff]">Web Developer</span>. Ketertarikan saya dimulai dari rasa penasaran tentang bagaimana baris kode dapat memecahkan masalah nyata.
            </p>
            <p>
              Saat ini, fokus utama saya adalah membangun aplikasi web yang responsif (menggunakan Vue & Tailwind) serta mengembangkan model AI yang dapat diimplementasikan ke dalam sistem produksi. Saya percaya pada prinsip <em class="text-[#8b949e]">"Clean Code"</em> dan kolaborasi terbuka (Open Source).
            </p>
          </div>

          <div class="mt-8">
            <h4 class="text-sm font-semibold text-[#8b949e] uppercase tracking-wider mb-4">
              Teknologi & Tools
            </h4>
            <div class="flex flex-wrap gap-2">
              <span v-for="tech in techStack" :key="tech" class="px-3 py-1 text-sm font-medium text-[#c9d1d9] bg-[#21262d] border border-[#30363d] rounded-full hover:border-[#8b949e] hover:text-white transition-colors duration-300 cursor-default">
                {{ tech }}
              </span>
            </div>
          </div>
        </div>

      </div>
    </div>
  </section>
</template>