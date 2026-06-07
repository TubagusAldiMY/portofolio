<script setup>
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import { apiFetch } from '../api/client.js';

const router = useRouter();
const username = ref('');
const password = ref('');
const errorMsg = ref('');

const handleLogin = async () => {
  try {
    const response = await apiFetch('/api/auth/login', {
      method: 'POST',
      body: JSON.stringify({ username: username.value, password: password.value })
    });

    const data = await response.json();

    if (response.ok) {
      // Simpan token di LocalStorage
      localStorage.setItem('token', data.token);
      router.push('/admin');
    } else {
      errorMsg.value = data.error || 'Login gagal';
    }
  } catch (err) {
    errorMsg.value = 'Terjadi kesalahan koneksi';
  }
};
</script>

<template>
  <div class="min-h-[80vh] flex items-center justify-center px-4">
    <div class="w-full max-w-md bg-[#161b22] border border-[#30363d] rounded-xl p-8 shadow-xl">
      <h2 class="text-2xl font-bold text-white text-center mb-6">Admin Login</h2>

      <div v-if="errorMsg" class="mb-4 p-3 bg-red-900/50 border border-red-800 text-red-200 rounded text-sm text-center">
        {{ errorMsg }}
      </div>

      <form @submit.prevent="handleLogin" class="space-y-6">
        <div>
          <label class="block text-sm font-medium text-[#c9d1d9] mb-2">Username</label>
          <input v-model="username" type="text" class="w-full bg-[#0d1117] border border-[#30363d] rounded px-3 py-2 text-white focus:border-[#58a6ff] outline-none transition">
        </div>
        <div>
          <label class="block text-sm font-medium text-[#c9d1d9] mb-2">Password</label>
          <input v-model="password" type="password" class="w-full bg-[#0d1117] border border-[#30363d] rounded px-3 py-2 text-white focus:border-[#58a6ff] outline-none transition">
        </div>
        <button type="submit" class="w-full bg-[#238636] hover:bg-[#2ea043] text-white font-bold py-2 px-4 rounded transition-colors">
          Sign In
        </button>
      </form>
    </div>
  </div>
</template>