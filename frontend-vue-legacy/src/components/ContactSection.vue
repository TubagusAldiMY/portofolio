<script setup>
import { ref } from 'vue';
import { apiFetch } from '../api/client.js';

const form = ref({
  name: '',
  email: '',
  message: '', // Di backend kita mapping ini ke 'content'
});

const isSubmitting = ref(false);

const handleSubmit = async () => {
  isSubmitting.value = true;

  try {
    const response = await apiFetch('/api/contact', {
      method: 'POST',
      body: JSON.stringify({
        name: form.value.name,
        email: form.value.email,
        content: form.value.message,
      }),
    });

    const result = await response.json();

    if (response.ok) {
      alert('Pesan berhasil dikirim! Terima kasih.');
      form.value = { name: '', email: '', message: '' };
    } else {
      alert('Gagal mengirim pesan: ' + result.error);
    }
  } catch (error) {
    console.error(error);
    alert('Terjadi kesalahan koneksi ke server.');
  } finally {
    isSubmitting.value = false;
  }
};
</script>

<template>
  <section id="contact" class="py-20 bg-[#0d1117] border-t border-[#30363d]">
    <div class="container mx-auto px-4 sm:px-6 lg:px-8">

      <div class="text-center mb-12">
        <h2 class="text-3xl font-extrabold text-white">Hubungi Saya</h2>
        <p class="mt-2 text-[#8b949e]">
          Kirim pesan, kolaborasi, atau sekadar menyapa.
        </p>
      </div>

      <div class="max-w-2xl mx-auto">

        <div class="bg-[#161b22] border border-[#30363d] rounded-md overflow-hidden shadow-lg">

          <div class="bg-[#21262d] border-b border-[#30363d] px-4 py-3 flex items-center gap-2">
            <div class="w-3 h-3 rounded-full bg-[#f25d5e]"></div> <div class="w-3 h-3 rounded-full bg-[#fbfb8d]"></div> <div class="w-3 h-3 rounded-full bg-[#62c554]"></div> <span class="ml-2 text-xs text-[#8b949e] font-mono">contact.md</span>
          </div>

          <div class="p-6">
            <form @submit.prevent="handleSubmit" class="space-y-5">

              <div class="grid grid-cols-1 md:grid-cols-2 gap-5">
                <div>
                  <label class="block text-xs font-semibold text-[#c9d1d9] mb-2 uppercase tracking-wide">Nama Lengkap</label>
                  <input v-model="form.name" type="text" required
                         class="w-full bg-[#0d1117] border border-[#30363d] rounded-md px-3 py-2 text-white placeholder-[#484f58] focus:border-[#58a6ff] focus:ring-1 focus:ring-[#58a6ff] outline-none transition"
                         placeholder="Nama Anda"
                  >
                </div>
                <div>
                  <label class="block text-xs font-semibold text-[#c9d1d9] mb-2 uppercase tracking-wide">Alamat Email</label>
                  <input v-model="form.email" type="email" required
                         class="w-full bg-[#0d1117] border border-[#30363d] rounded-md px-3 py-2 text-white placeholder-[#484f58] focus:border-[#58a6ff] focus:ring-1 focus:ring-[#58a6ff] outline-none transition"
                         placeholder="nama@email.com"
                  >
                </div>
              </div>

              <div>
                <label class="block text-xs font-semibold text-[#c9d1d9] mb-2 uppercase tracking-wide">Pesan</label>
                <textarea v-model="form.message" rows="6" required
                          class="w-full bg-[#0d1117] border border-[#30363d] rounded-md px-3 py-2 text-white font-mono text-sm placeholder-[#484f58] focus:border-[#58a6ff] focus:ring-1 focus:ring-[#58a6ff] outline-none transition"
                          placeholder="Tulis pesan Anda di sini..."
                ></textarea>
                <p class="text-xs text-[#8b949e] mt-1 text-right">Markdown is supported</p>
              </div>

              <div class="flex justify-end">
                <button type="submit" :disabled="isSubmitting"
                        class="bg-[#238636] text-white font-semibold px-6 py-2 rounded-md hover:bg-[#2ea043] border border-[rgba(240,246,252,0.1)] transition-colors shadow-sm disabled:opacity-50 disabled:cursor-not-allowed"
                >
                  {{ isSubmitting ? 'Mengirim...' : 'Kirim Pesan' }}
                </button>
              </div>

            </form>
          </div>
        </div>

        <div class="mt-8 text-center">
          <p class="text-[#8b949e] text-sm">
            Atau temukan saya di
            <a href="https://linkedin.com/in/tubagusaldi" target="_blank" class="text-[#58a6ff] hover:underline">LinkedIn</a>
            dan
            <a href="https://github.com/TubagusAldiMY" target="_blank" class="text-[#58a6ff] hover:underline">GitHub</a>.
          </p>
        </div>

      </div>
    </div>
  </section>
</template>