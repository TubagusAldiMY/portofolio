<script setup>
import { ref, onMounted } from 'vue';
import { apiFetch } from '../api/client.js';

const experiences = ref([]);
const isLoading = ref(true);

const fetchExperiences = async () => {
  try {
    const response = await apiFetch('/api/experiences');
    const result = await response.json();

    // Transform deskripsi (String multiline) menjadi Array list
    experiences.value = (result.data || []).map(exp => ({
      ...exp,
      descList: exp.description ? exp.description.split('\n') : []
    }));
  } catch (error) {
    console.error("Gagal memuat pengalaman:", error);
  } finally {
    isLoading.value = false;
  }
};

onMounted(() => {
  fetchExperiences();
});
</script>

<template>
  <section id="experience" class="py-20 sm:py-32 bg-[#0d1117]" v-scroll-animation>
    <div class="container mx-auto px-4 sm:px-6 lg:px-8">

      <div class="text-center mb-16">
        <h2 class="text-3xl sm:text-4xl font-extrabold text-white">Contribution Activity</h2>
        <p class="mt-4 text-lg text-[#8b949e]">
          Riwayat pengalaman dan kontribusi profesional saya.
        </p>
      </div>

      <div class="max-w-3xl mx-auto">

        <div v-if="isLoading" class="text-center text-[#8b949e]">Loading activity...</div>

        <div v-else class="relative border-l border-[#30363d] ml-3 md:ml-6 space-y-12">

          <div v-for="exp in experiences" :key="exp.id" class="relative pl-8 md:pl-12 group">
            <span class="absolute -left-[5px] top-0 bg-[#0d1117] border-2 border-[#58a6ff] w-3 h-3 rounded-full mt-1.5 ring-4 ring-[#0d1117] group-hover:bg-[#58a6ff] transition-colors"></span>

            <div class="flex flex-col sm:flex-row sm:items-center sm:justify-between mb-2">
              <h3 class="text-xl font-bold text-white group-hover:text-[#58a6ff] transition-colors">
                {{ exp.role }}
              </h3>
              <span class="text-xs font-mono text-[#8b949e] bg-[#161b22] px-2 py-1 rounded border border-[#30363d] mt-1 sm:mt-0 w-fit">
                {{ exp.duration }}
              </span>
            </div>

            <p class="text-base font-medium text-[#c9d1d9] mb-4">
              {{ exp.company }}
            </p>

            <div class="bg-[#161b22] border border-[#30363d] rounded-md p-4 relative hover:border-[#8b949e] transition-colors">
              <div class="absolute top-4 -left-1.5 w-3 h-3 bg-[#161b22] border-l border-t border-[#30363d] transform -rotate-45"></div>

              <ul class="space-y-2">
                <li v-for="(desc, i) in exp.descList" :key="i" class="flex items-start text-[#8b949e] text-sm">
                  <svg class="w-4 h-4 mr-2 mt-0.5 text-[#238636] flex-shrink-0" fill="currentColor" viewBox="0 0 20 20">
                    <path fill-rule="evenodd" d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z" clip-rule="evenodd"/>
                  </svg>
                  <span>{{ desc }}</span>
                </li>
              </ul>
            </div>

          </div>

        </div>
      </div>
    </div>
  </section>
</template>