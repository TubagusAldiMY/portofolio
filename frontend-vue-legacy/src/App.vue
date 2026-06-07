<script setup>
import Navbar from './components/Navbar.vue';
import Footer from './components/Footer.vue';
import ChatWidget from './components/ChatWidget.vue';
</script>

<template>
  <div id="app" class="min-h-screen text-[#c9d1d9] bg-[#0d1117]">
    <Navbar />
    <router-view />
    <Footer />
    <ChatWidget />
  </div>
</template>