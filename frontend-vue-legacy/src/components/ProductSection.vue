<script setup>
import { ref, onMounted } from 'vue';
import { apiFetch } from '../api/client.js';

const products = ref([]);
const isLoading = ref(true);

const fetchProducts = async () => {
  try {
    const response = await apiFetch('/api/products');
    const result = await response.json();

    // Gunakan array dari backend; fallback ke CSV split untuk data lama
    products.value = (result.data || []).map(product => ({
      ...product,
      featuresList: Array.isArray(product.features)
        ? product.features
        : (product.features ? product.features.split(',').map(f => f.trim()).filter(Boolean) : [])
    }));
  } catch (error) {
    console.error("Gagal memuat produk:", error);
  } finally {
    isLoading.value = false;
  }
};

onMounted(() => {
  fetchProducts();
});
</script>

<template>
  <section id="products" class="py-20 bg-[#0d1117] border-t border-[#30363d]" v-scroll-animation>
    <div class="container mx-auto px-4 sm:px-6 lg:px-8">

      <div class="text-center mb-16">
        <div class="inline-block px-3 py-1 mb-4 text-xs font-semibold tracking-wide text-[#a371f7] uppercase bg-[#a371f7]/10 rounded-full border border-[#a371f7]/30">
          Marketplace
        </div>
        <h2 class="text-3xl sm:text-4xl font-extrabold text-white mb-4">
          Produk & Aplikasi Premium
        </h2>
        <p class="text-[#8b949e] max-w-2xl mx-auto text-lg">
          Solusi siap pakai yang saya kembangkan untuk meningkatkan produktivitas.
        </p>
      </div>

      <div v-if="isLoading" class="text-center text-[#8b949e]">Loading products...</div>

      <div v-else class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
        <div
            v-for="product in products"
            :key="product.id"
            class="group relative bg-[#161b22] border border-[#30363d] rounded-xl overflow-hidden transition-all duration-300 hover:border-[#58a6ff] hover:-translate-y-2 flex flex-col"
        >
          <div class="h-48 overflow-hidden relative bg-[#0d1117]">
            <div v-if="product.tag" class="absolute top-4 left-4 z-20">
              <span class="px-3 py-1 bg-[#0d1117]/80 backdrop-blur-sm border border-[#30363d] text-xs font-bold text-white rounded-full">
                {{ product.tag }}
              </span>
            </div>
            <img
                :src="product.imageUrl || 'https://placehold.co/600x400/161b22/ffffff?text=No+Image'"
                :alt="product.title"
                class="w-full h-full object-cover transition-transform duration-700 group-hover:scale-110 opacity-90 group-hover:opacity-100"
            >
          </div>

          <div class="p-6 pt-2 flex-grow flex flex-col relative z-20">
            <div class="flex justify-end -mt-10 mb-4">
              <div class="px-4 py-2 bg-[#238636] text-white font-bold rounded-lg border border-[#2ea043] shadow-lg">
                {{ product.price }}
              </div>
            </div>

            <h3 class="text-xl font-bold text-white mb-3 group-hover:text-[#58a6ff] transition-colors">
              {{ product.title }}
            </h3>

            <p class="text-[#8b949e] text-sm mb-6 leading-relaxed line-clamp-3">
              {{ product.description }}
            </p>

            <ul class="space-y-3 mb-8 flex-grow">
              <li v-for="(feature, i) in product.featuresList" :key="i" class="flex items-start text-sm text-[#c9d1d9]">
                <svg class="w-5 h-5 mr-3 text-[#238636] flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"></path></svg>
                {{ feature }}
              </li>
            </ul>

            <a :href="product.buyUrl || '#'" target="_blank" class="mt-auto w-full py-3 flex items-center justify-center gap-2 text-white font-bold bg-[#1f6feb] border border-[#388bfd] rounded-lg hover:bg-[#388bfd] transition-all duration-300">
              <svg class="w-4 h-4" fill="currentColor" viewBox="0 0 24 24"><path d="M6 2a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2H6zm0 2h12v16H6V4zm2 4v2h8V8H8zm0 4v2h8v-2H8zm0 4v2h5v-2H8z"/></svg>
              Beli Sekarang
            </a>
          </div>
        </div>
      </div>

    </div>
  </section>
</template>