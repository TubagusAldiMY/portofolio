<template>
  <nav class="sticky top-0 z-50 bg-[#0d1117]/90 backdrop-blur-md border-b border-[#30363d]">
    <div class="container mx-auto px-4 sm:px-6 lg:px-8">
      <div class="flex items-center justify-between h-16">

        <div class="flex-shrink-0 flex items-center gap-2">
          <a href="#" class="text-xl font-bold text-white tracking-tight">
            TubsAMY
          </a>
        </div>

        <div class="hidden md:flex md:items-center md:space-x-8">
          <a v-for="link in navLinks" :key="link.name" :href="link.href" class="text-[#c9d1d9] hover:text-white hover:opacity-100 opacity-80 transition-all duration-300 font-medium text-sm">
            {{ link.name }}
          </a>
        </div>

        <div class="hidden md:block">
          <a href="#contact" class="px-4 py-2 text-sm font-semibold text-white bg-[#238636] border border-[rgba(240,246,252,0.1)] rounded-md hover:bg-[#2ea043] transition-all duration-300 shadow-sm">
            Hubungi Saya
          </a>
        </div>

        <div class="md:hidden flex items-center">
          <button @click="toggleMenu" class="text-[#c9d1d9] hover:text-white focus:outline-none">
            <svg class="h-6 w-6" stroke="currentColor" fill="none" viewBox="0 0 24 24">
              <path v-if="isMenuOpen" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
              <path v-else stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16m-7 6h7" />
            </svg>
          </button>
        </div>

      </div>
    </div>

    <transition name="slide-fade">
      <div v-if="isMenuOpen" class="md:hidden bg-[#161b22] border-t border-[#30363d]">
        <div class="px-2 pt-2 pb-3 space-y-1 sm:px-3">
          <a v-for="link in navLinks" :key="link.name" :href="link.href" class="block px-3 py-2 rounded-md text-base font-medium text-[#c9d1d9] hover:text-white hover:bg-[#1f6feb]/20" @click="toggleMenu">
            {{ link.name }}
          </a>
        </div>
      </div>
    </transition>
  </nav>
</template>

<script setup>
import { ref } from 'vue';

const isMenuOpen = ref(false);
const toggleMenu = () => {
  isMenuOpen.value = !isMenuOpen.value;
};

const navLinks = ref([
  { name: 'Overview', href: '#home' },
  { name: 'Repositories', href: '#projects' },
  { name: 'Products', href: '#products' },
  { name: 'Contribution', href: '#experience' },
  { name: 'About', href: '#about' },
]);
</script>

<style scoped>
.slide-fade-enter-active {
  transition: all 0.3s ease-out;
}
.slide-fade-leave-active {
  transition: all 0.3s cubic-bezier(1, 0.5, 0.8, 1);
}
.slide-fade-enter-from,
.slide-fade-leave-to {
  transform: translateY(-20px);
  opacity: 0;
}
</style>