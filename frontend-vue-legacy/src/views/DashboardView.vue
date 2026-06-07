<script setup>
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import MessagesManager from '../components/admin/MessagesManager.vue';
import ContentManager from '../components/admin/ContentManager.vue';

const router = useRouter();
const activeTab = ref('messages');

const logout = () => {
  localStorage.removeItem('token');
  router.push('/login');
};

// KONFIGURASI FIELD UNTUK CONTENT MANAGER
const projectFields = [
  { key: 'title', label: 'Judul Proyek', type: 'text' },
  { key: 'description', label: 'Deskripsi', type: 'textarea' },
  { key: 'imageUrl', label: 'Gambar Cover', type: 'image' },
  { key: 'techStack', label: 'Teknologi (Pisahkan koma)', type: 'text' },
  { key: 'repoUrl', label: 'Link GitHub', type: 'text' },
  { key: 'liveUrl', label: 'Link Demo', type: 'text' },
];

const productFields = [
  { key: 'title', label: 'Nama Produk', type: 'text' },
  { key: 'description', label: 'Deskripsi', type: 'textarea' },
  { key: 'price', label: 'Harga (Rp)', type: 'text' },
  { key: 'imageUrl', label: 'Gambar Produk', type: 'image' },
  { key: 'features', label: 'Fitur (Pisahkan koma)', type: 'textarea' },
  { key: 'buyUrl', label: 'Link Pembelian', type: 'text' },
  { key: 'tag', label: 'Label Tag (New, Best Seller)', type: 'text' },
];

const experienceFields = [
  { key: 'role', label: 'Posisi / Role', type: 'text' },
  { key: 'company', label: 'Perusahaan / Organisasi', type: 'text' },
  { key: 'duration', label: 'Durasi (Contoh: 2023 - Sekarang)', type: 'text' },
  { key: 'description', label: 'Deskripsi Pekerjaan', type: 'textarea' },
];
</script>

<template>
  <div class="min-h-screen bg-[#0d1117] text-[#c9d1d9]">

    <header class="border-b border-[#30363d] bg-[#161b22]">
      <div class="container mx-auto px-4 h-16 flex items-center justify-between">
        <h1 class="text-lg font-bold text-white flex items-center gap-2">
          <div class="w-3 h-3 bg-[#238636] rounded-full"></div>
          Admin Panel
        </h1>
        <button @click="logout" class="text-xs font-bold text-red-400 hover:text-red-300 border border-red-900/50 px-3 py-1.5 rounded hover:bg-red-900/20 transition">
          Sign Out
        </button>
      </div>
    </header>

    <div class="container mx-auto px-4 py-8 flex flex-col md:flex-row gap-8">

      <aside class="w-full md:w-64 shrink-0">
        <nav class="space-y-1">
          <button @click="activeTab = 'messages'" :class="['w-full text-left px-4 py-2 rounded text-sm font-medium transition', activeTab === 'messages' ? 'bg-[#1f6feb] text-white' : 'hover:bg-[#21262d]']">
            Inbox Pesan
          </button>
          <div class="h-px bg-[#30363d] my-2"></div>
          <button @click="activeTab = 'projects'" :class="['w-full text-left px-4 py-2 rounded text-sm font-medium transition', activeTab === 'projects' ? 'bg-[#1f6feb] text-white' : 'hover:bg-[#21262d]']">
            Kelola Projects
          </button>
          <button @click="activeTab = 'products'" :class="['w-full text-left px-4 py-2 rounded text-sm font-medium transition', activeTab === 'products' ? 'bg-[#1f6feb] text-white' : 'hover:bg-[#21262d]']">
            Kelola Produk
          </button>
          <button @click="activeTab = 'experiences'" :class="['w-full text-left px-4 py-2 rounded text-sm font-medium transition', activeTab === 'experiences' ? 'bg-[#1f6feb] text-white' : 'hover:bg-[#21262d]']">
            Kelola Experience
          </button>
        </nav>
      </aside>

      <main class="flex-1">

        <MessagesManager v-if="activeTab === 'messages'" />

        <ContentManager
            v-if="activeTab === 'projects'"
            title="Daftar Repositori / Proyek"
            endpoint="projects"
            :fields="projectFields"
        />

        <ContentManager
            v-if="activeTab === 'products'"
            title="Daftar Produk Digital"
            endpoint="products"
            :fields="productFields"
        />

        <ContentManager
            v-if="activeTab === 'experiences'"
            title="Riwayat Pengalaman"
            endpoint="experiences"
            :fields="experienceFields"
        />

      </main>
    </div>
  </div>
</template>