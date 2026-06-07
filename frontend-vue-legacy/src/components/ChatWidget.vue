<script setup>
import { ref, nextTick } from 'vue';
import { apiFetch } from '../api/client.js';
import { marked } from 'marked'; // Import library markdown parser

const isOpen = ref(false);
const messages = ref([
  {
    id: 1,
    text: "Halo! Saya asisten virtual **Tubagus Aldi**. Ada yang bisa saya bantu tentang portofolio ini?",
    isUser: false
  }
]);
const userInput = ref('');
const isLoading = ref(false);
const messagesContainer = ref(null);

// Fungsi helper untuk merender markdown menjadi HTML
const renderMessage = (text) => {
  // marked.parse akan mengubah string markdown (seperti **teks**) menjadi HTML (<strong>teks</strong>)
  return marked.parse(text);
};

const toggleChat = () => {
  isOpen.value = !isOpen.value;
  if (isOpen.value) scrollToBottom();
};

const scrollToBottom = () => {
  nextTick(() => {
    if (messagesContainer.value) {
      messagesContainer.value.scrollTop = messagesContainer.value.scrollHeight;
    }
  });
};

const sendMessage = async () => {
  if (!userInput.value.trim() || isLoading.value) return;

  const text = userInput.value;
  // Tambahkan pesan user ke list
  messages.value.push({ id: Date.now(), text, isUser: true });

  userInput.value = '';
  isLoading.value = true;
  scrollToBottom();

  try {
    const res = await apiFetch('/api/chat', {
      method: 'POST',
      body: JSON.stringify({ message: text })
    });

    const data = await res.json();

    // Tambahkan balasan AI ke list
    messages.value.push({
      id: Date.now() + 1,
      text: data.reply || "Maaf, terjadi kesalahan pada server.",
      isUser: false
    });

  } catch (err) {
    console.error(err);
    messages.value.push({
      id: Date.now() + 1,
      text: "Gagal terhubung ke server. Silakan coba lagi nanti.",
      isUser: false
    });
  } finally {
    isLoading.value = false;
    scrollToBottom();
  }
};
</script>

<template>
  <div class="fixed bottom-6 right-6 z-50 flex flex-col items-end font-sans">

    <transition name="scale">
      <div v-if="isOpen" class="mb-4 w-80 sm:w-96 bg-[#161b22] border border-[#30363d] rounded-xl shadow-2xl overflow-hidden flex flex-col h-[450px]">

        <div class="bg-[#0d1117] p-4 border-b border-[#30363d] flex justify-between items-center">
          <div class="flex items-center gap-3">
            <div class="relative">
              <div class="w-8 h-8 rounded-full bg-gradient-to-tr from-[#1f6feb] to-[#58a6ff] flex items-center justify-center text-white font-bold text-xs">
                AI
              </div>
              <div class="absolute bottom-0 right-0 w-2.5 h-2.5 bg-green-500 border-2 border-[#0d1117] rounded-full animate-pulse"></div>
            </div>
            <div>
              <h3 class="font-bold text-white text-sm">Aldi's Assistant</h3>
            </div>
          </div>
          <button @click="toggleChat" class="text-[#8b949e] hover:text-white transition-colors p-1 rounded hover:bg-[#30363d]">
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path></svg>
          </button>
        </div>

        <div ref="messagesContainer" class="flex-1 p-4 overflow-y-auto space-y-4 bg-[#0d1117]/95 scrollbar-thin scrollbar-thumb-[#30363d] scrollbar-track-transparent">

          <div v-for="msg in messages" :key="msg.id" :class="['flex', msg.isUser ? 'justify-end' : 'justify-start']">

            <div
                :class="[
                'max-w-[85%] rounded-2xl px-4 py-2.5 text-sm leading-relaxed shadow-sm',
                msg.isUser
                  ? 'bg-[#1f6feb] text-white rounded-br-none'
                  : 'bg-[#21262d] text-[#c9d1d9] border border-[#30363d] rounded-bl-none'
              ]"
            >
              <div
                  v-if="!msg.isUser"
                  class="markdown-content prose prose-invert prose-sm max-w-none prose-p:my-1 prose-ul:my-1 prose-li:my-0"
                  v-html="renderMessage(msg.text)"
              ></div>
              <div v-else>
                {{ msg.text }}
              </div>
            </div>

          </div>

          <div v-if="isLoading" class="flex justify-start animate-pulse">
            <div class="bg-[#21262d] border border-[#30363d] rounded-2xl rounded-bl-none px-4 py-3 flex gap-1 items-center">
              <div class="w-1.5 h-1.5 bg-[#8b949e] rounded-full animate-bounce"></div>
              <div class="w-1.5 h-1.5 bg-[#8b949e] rounded-full animate-bounce delay-75"></div>
              <div class="w-1.5 h-1.5 bg-[#8b949e] rounded-full animate-bounce delay-150"></div>
            </div>
          </div>

        </div>

        <form @submit.prevent="sendMessage" class="p-3 bg-[#161b22] border-t border-[#30363d] flex gap-2">
          <input
              v-model="userInput"
              type="text"
              placeholder="Tanya tentang pengalaman, tech stack..."
              class="flex-1 bg-[#0d1117] border border-[#30363d] rounded-lg px-4 py-2.5 text-sm text-white focus:border-[#58a6ff] focus:ring-1 focus:ring-[#58a6ff] outline-none transition placeholder-[#484f58]"
          >
          <button
              type="submit"
              :disabled="isLoading || !userInput.trim()"
              class="bg-[#238636] hover:bg-[#2ea043] text-white px-3 py-2 rounded-lg disabled:opacity-50 disabled:cursor-not-allowed transition-all active:scale-95 flex items-center justify-center w-10"
          >
            <svg v-if="!isLoading" class="w-4 h-4 ml-0.5" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 19l9 2-9-18-9 18 9-2zm0 0v-8"></path></svg>
            <svg v-else class="animate-spin w-4 h-4" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24"><circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle><path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path></svg>
          </button>
        </form>

      </div>
    </transition>

    <button
        @click="toggleChat"
        class="bg-[#1f6feb] hover:bg-[#388bfd] text-white w-14 h-14 rounded-full shadow-lg shadow-[#1f6feb]/20 flex items-center justify-center transition-all duration-300 hover:scale-110 active:scale-95 group z-50"
    >
      <div class="relative">
        <svg v-if="!isOpen" class="w-7 h-7 transition-transform group-hover:rotate-12" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 10h.01M12 10h.01M16 10h.01M9 16H5a2 2 0 01-2-2V6a2 2 0 012-2h14a2 2 0 012 2v8a2 2 0 01-2 2h-5l-5 5v-5z"></path></svg>
        <svg v-else class="w-7 h-7" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7"></path></svg>

        <span v-if="!isOpen" class="absolute -top-1 -right-1 flex h-3 w-3">
          <span class="animate-ping absolute inline-flex h-full w-full rounded-full bg-red-400 opacity-75"></span>
          <span class="relative inline-flex rounded-full h-3 w-3 bg-red-500"></span>
        </span>
      </div>
    </button>

  </div>
</template>

<style scoped>
.scale-enter-active, .scale-leave-active {
  transition: all 0.3s cubic-bezier(0.16, 1, 0.3, 1);
}
.scale-enter-from, .scale-leave-to {
  opacity: 0;
  transform: scale(0.9) translateY(20px) translateX(20px);
}

/* Custom Scrollbar for Chat Window */
.scrollbar-thin::-webkit-scrollbar {
  width: 6px;
}
.scrollbar-thin::-webkit-scrollbar-track {
  background: transparent;
}
.scrollbar-thin::-webkit-scrollbar-thumb {
  background-color: #30363d;
  border-radius: 3px;
}

/* Markdown Styles Override (Agar sesuai tema GitHub Dark) */
:deep(.markdown-content ul) {
  list-style-type: disc;
  padding-left: 1.2em;
}
:deep(.markdown-content ol) {
  list-style-type: decimal;
  padding-left: 1.2em;
}
:deep(.markdown-content strong) {
  color: #fff; /* Teks tebal lebih putih */
  font-weight: 700;
}
:deep(.markdown-content a) {
  color: #58a6ff;
  text-decoration: underline;
}
:deep(.markdown-content p) {
  margin-bottom: 0.5em;
}
:deep(.markdown-content p:last-child) {
  margin-bottom: 0;
}
</style>
