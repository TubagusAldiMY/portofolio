<script setup>
import { ref, onMounted } from 'vue';
import { apiFetch } from '../api/client.js';

const projects = ref([]);
const isLoading = ref(true);

const fetchProjects = async () => {
  try {
    const response = await apiFetch('/api/projects');
    const result = await response.json();

    // Prefer array dari backend; fallback ke CSV untuk data lama
    projects.value = (result.data || []).map(project => ({
      ...project,
      tech: Array.isArray(project.techStack)
        ? project.techStack
        : (project.techStack ? project.techStack.split(',').map(t => t.trim()).filter(Boolean) : [])
    }));
  } catch (error) {
    console.error("Gagal memuat proyek:", error);
  } finally {
    isLoading.value = false;
  }
};

onMounted(() => {
  fetchProjects();
});
</script>

<template>
  <section id="projects" class="py-20 bg-[#0d1117]" v-scroll-animation>
    <div class="container mx-auto px-4 sm:px-6 lg:px-8">

      <div class="mb-12">
        <h2 class="text-3xl md:text-4xl font-bold text-white mb-4">Top Repositories</h2>
        <p class="text-[#8b949e] text-lg">
          Koleksi proyek (repositories) yang saya kerjakan.
        </p>
      </div>

      <div v-if="isLoading" class="text-center py-12 text-[#8b949e]">
        Memuat data proyek...
      </div>

      <div v-else-if="projects.length === 0" class="text-center py-12 text-[#8b949e]">
        Belum ada proyek yang ditampilkan.
      </div>

      <div v-else class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div
            v-for="project in projects"
            :key="project.id"
            class="bg-[#0d1117] border border-[#30363d] rounded-md p-6 hover:border-[#8b949e] transition-colors duration-200 flex flex-col justify-between group"
        >
          <div>
            <div class="flex items-center justify-between mb-3">
              <div class="flex items-center gap-2 overflow-hidden">
                <svg aria-hidden="true" height="16" viewBox="0 0 16 16" version="1.1" width="16" fill="#8b949e" class="mt-1 flex-shrink-0">
                  <path d="M2 2.5A2.5 2.5 0 0 1 4.5 0h8.75a.75.75 0 0 1 .75.75v12.5a.75.75 0 0 1-.75.75h-2.5a.75.75 0 0 1 0-1.5h1.75v-2h-8a1 1 0 0 0-.714 1.7.75.75 0 1 1-1.072 1.05A2.495 2.495 0 0 1 2 11.5Zm10.5-1h-8a1 1 0 0 0-1 1v6.708A2.486 2.486 0 0 1 4.5 9h8ZM5 12.25a.25.25 0 0 1 .25-.25h3.5a.25.25 0 0 1 .25.25v3.25a.25.25 0 0 1-.4.2l-1.45-1.087a.249.249 0 0 0-.3 0L5.4 15.7a.25.25 0 0 1-.4-.2Z"></path>
                </svg>
                <a :href="project.repoUrl || '#'" target="_blank" class="text-[#58a6ff] font-bold text-lg hover:underline truncate">
                  {{ project.title }}
                </a>
                <span class="px-2 py-0.5 ml-2 text-xs font-medium border border-[#30363d] rounded-full text-[#8b949e] flex-shrink-0">Public</span>
              </div>
            </div>

            <p class="text-[#8b949e] text-sm mb-6 leading-relaxed line-clamp-3">
              {{ project.description }}
            </p>
          </div>

          <div class="flex flex-wrap items-center text-xs text-[#8b949e] gap-4 mt-auto">
            <div v-for="(t, index) in project.tech" :key="index" class="flex items-center gap-1">
              <span class="w-3 h-3 rounded-full bg-[#3572A5]"></span> <span>{{ t }}</span>
            </div>

            <div class="ml-auto flex gap-3">
              <a :href="project.liveUrl" v-if="project.liveUrl" target="_blank" class="hover:text-[#58a6ff] transition-colors flex items-center gap-1">
                View Demo ↗
              </a>
            </div>
          </div>

        </div>
      </div>
    </div>
  </section>
</template>