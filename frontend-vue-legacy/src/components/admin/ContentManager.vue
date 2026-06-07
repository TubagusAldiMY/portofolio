<script setup>
import { ref, onMounted } from 'vue';
import { apiFetch } from '../../api/client.js';

// Props untuk konfigurasi dinamis
const props = defineProps({
  title: String,
  endpoint: String, // Contoh: 'projects' atau 'products'
  fields: Array     // Konfigurasi form input
});

const items = ref([]);
const showModal = ref(false);
const isEditing = ref(false);
const isLoading = ref(false);
const editId = ref(null);

// State untuk Form Data
const formData = ref({});

// Inisialisasi form kosong
const resetForm = () => {
  const initialData = {};
  props.fields.forEach(field => initialData[field.key] = '');
  formData.value = initialData;
};

// --- API ACTIONS ---

const fetchItems = async () => {
  const response = await apiFetch(`/api/${props.endpoint}`);
  const result = await response.json();
  items.value = result.data || [];
};

const handleFileUpload = async (event, key) => {
  const file = event.target.files[0];
  if (!file) return;

  const formDataUpload = new FormData();
  formDataUpload.append('file', file);

  try {
    const token = localStorage.getItem('token');
    const res = await apiFetch('/api/admin/upload', {
      method: 'POST',
      headers: { 'Authorization': `Bearer ${token}` },
      body: formDataUpload
    });
    const data = await res.json();

    if (res.ok) {
      formData.value[key] = data.url; // Simpan URL gambar ke form
      alert('Upload Berhasil!');
    } else {
      alert('Upload Gagal: ' + data.error);
    }
  } catch (err) {
    console.error(err);
    alert('Error upload file');
  }
};

const saveItem = async () => {
  isLoading.value = true;
  const token = localStorage.getItem('token');
  const url = isEditing.value
      ? `/api/admin/${props.endpoint}/${editId.value}`
      : `/api/admin/${props.endpoint}`;

  const method = isEditing.value ? 'PUT' : 'POST';

  try {
    const payload = { ...formData.value };
    if (typeof payload.techStack === 'string') {
      payload.techStack = payload.techStack.split(',').map(t => t.trim()).filter(Boolean);
    }
    if (typeof payload.features === 'string') {
      payload.features = payload.features.split(',').map(f => f.trim()).filter(Boolean);
    }

    const res = await apiFetch(url, {
      method: method,
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${token}`
      },
      body: JSON.stringify(payload)
    });

    if (res.ok) {
      closeModal();
      fetchItems();
    } else {
      alert('Gagal menyimpan data');
    }
  } catch (err) {
    console.error(err);
  } finally {
    isLoading.value = false;
  }
};

const deleteItem = async (id) => {
  if (!confirm('Yakin ingin menghapus item ini?')) return;
  const token = localStorage.getItem('token');
  await apiFetch(`/api/admin/${props.endpoint}/${id}`, {
    method: 'DELETE',
    headers: { 'Authorization': `Bearer ${token}` }
  });
  fetchItems();
};

// --- UI HELPERS ---

const openAddModal = () => {
  isEditing.value = false;
  resetForm();
  showModal.value = true;
};

const openEditModal = (item) => {
  isEditing.value = true;
  editId.value = item.id;
  const copy = { ...item };
  if (Array.isArray(copy.techStack)) {
    copy.techStack = copy.techStack.join(', ');
  }
  if (Array.isArray(copy.features)) {
    copy.features = copy.features.join(', ');
  }
  formData.value = copy;
  showModal.value = true;
};

const closeModal = () => {
  showModal.value = false;
};

onMounted(() => {
  resetForm();
  fetchItems();
});
</script>

<template>
  <div>
    <div class="flex justify-between items-center mb-6">
      <h2 class="text-xl font-bold text-white">{{ title }}</h2>
      <button @click="openAddModal" class="bg-[#238636] hover:bg-[#2ea043] text-white px-4 py-2 rounded text-sm font-bold transition">
        + Tambah Baru
      </button>
    </div>

    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
      <div v-for="item in items" :key="item.id" class="bg-[#161b22] border border-[#30363d] rounded-lg p-4 flex flex-col">
        <img v-if="item.imageUrl" :src="item.imageUrl" class="w-full h-32 object-cover rounded mb-3 bg-[#0d1117]" alt="Preview">

        <h3 class="font-bold text-[#c9d1d9] text-lg mb-1 truncate">{{ item.title || item.role || item.name }}</h3>
        <p class="text-[#8b949e] text-sm mb-4 line-clamp-2">{{ item.description || item.company }}</p>

        <div class="mt-auto flex gap-2">
          <button @click="openEditModal(item)" class="flex-1 bg-[#1f6feb] hover:bg-[#388bfd] text-white py-1.5 rounded text-xs font-bold">Edit</button>
          <button @click="deleteItem(item.id)" class="flex-1 bg-red-600 hover:bg-red-700 text-white py-1.5 rounded text-xs font-bold">Hapus</button>
        </div>
      </div>
    </div>

    <div v-if="showModal" class="fixed inset-0 bg-black/80 backdrop-blur-sm flex items-center justify-center z-50 p-4">
      <div class="bg-[#161b22] border border-[#30363d] rounded-xl w-full max-w-lg max-h-[90vh] overflow-y-auto shadow-2xl">
        <div class="p-6">
          <h3 class="text-xl font-bold text-white mb-4">{{ isEditing ? 'Edit Data' : 'Tambah Data' }}</h3>

          <form @submit.prevent="saveItem" class="space-y-4">

            <div v-for="field in fields" :key="field.key">
              <label class="block text-xs font-bold text-[#8b949e] uppercase mb-1">{{ field.label }}</label>

              <input v-if="field.type === 'text'" v-model="formData[field.key]" type="text" class="w-full bg-[#0d1117] border border-[#30363d] rounded px-3 py-2 text-white focus:border-[#58a6ff] outline-none text-sm">

              <textarea v-else-if="field.type === 'textarea'" v-model="formData[field.key]" rows="3" class="w-full bg-[#0d1117] border border-[#30363d] rounded px-3 py-2 text-white focus:border-[#58a6ff] outline-none text-sm"></textarea>

              <div v-else-if="field.type === 'image'">
                <div class="flex gap-2 items-center">
                  <input type="text" v-model="formData[field.key]" placeholder="URL Gambar" class="flex-1 bg-[#0d1117] border border-[#30363d] rounded px-3 py-2 text-white text-sm">
                  <label class="cursor-pointer bg-[#21262d] border border-[#30363d] px-3 py-2 rounded text-sm text-white hover:bg-[#30363d]">
                    Upload
                    <input type="file" class="hidden" @change="(e) => handleFileUpload(e, field.key)">
                  </label>
                </div>
                <img v-if="formData[field.key]" :src="formData[field.key]" class="mt-2 h-20 rounded border border-[#30363d]">
              </div>

            </div>

            <div class="flex justify-end gap-3 mt-6">
              <button type="button" @click="closeModal" class="px-4 py-2 text-[#c9d1d9] hover:text-white text-sm font-bold">Batal</button>
              <button type="submit" :disabled="isLoading" class="bg-[#238636] hover:bg-[#2ea043] text-white px-6 py-2 rounded text-sm font-bold">
                {{ isLoading ? 'Menyimpan...' : 'Simpan' }}
              </button>
            </div>

          </form>
        </div>
      </div>
    </div>

  </div>
</template>