<script setup>
import { ref, onMounted } from 'vue';
import { apiFetch } from '../../api/client.js';

const messages = ref([]);
const loading = ref(true);

const fetchMessages = async () => {
  const token = localStorage.getItem('token');
  try {
    const response = await apiFetch('/api/admin/messages', {
      headers: { 'Authorization': `Bearer ${token}` }
    });
    const result = await response.json();
    messages.value = result.data || [];
  } catch (err) {
    console.error(err);
  } finally {
    loading.value = false;
  }
};

onMounted(fetchMessages);
</script>

<template>
  <div class="bg-[#161b22] border border-[#30363d] rounded-xl overflow-hidden">
    <div class="p-4 border-b border-[#30363d] bg-[#0d1117] flex justify-between items-center">
      <h3 class="font-bold text-[#c9d1d9]">Inbox Pesan</h3>
      <span class="text-xs bg-[#1f6feb] text-white px-2 py-1 rounded-full">{{ messages.length }}</span>
    </div>

    <div v-if="loading" class="p-8 text-center text-[#8b949e]">Loading...</div>
    <div v-else-if="messages.length === 0" class="p-8 text-center text-[#8b949e]">Belum ada pesan.</div>

    <div v-else class="divide-y divide-[#30363d]">
      <div v-for="msg in messages" :key="msg.id" class="p-6 hover:bg-[#21262d]/50 transition">
        <div class="flex justify-between items-start mb-2">
          <div>
            <h4 class="text-[#58a6ff] font-bold">{{ msg.name }}</h4>
            <a :href="`mailto:${msg.email}`" class="text-xs text-[#8b949e] hover:underline">{{ msg.email }}</a>
          </div>
          <span class="text-xs text-[#8b949e]">{{ new Date(msg.created_at).toLocaleString() }}</span>
        </div>
        <p class="text-[#c9d1d9] mt-2 text-sm bg-[#0d1117] p-3 rounded border border-[#30363d]">{{ msg.content }}</p>
      </div>
    </div>
  </div>
</template>