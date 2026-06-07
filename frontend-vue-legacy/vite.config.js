import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import tailwindcss from '@tailwindcss/vite'

// https://vite.dev/config/
export default defineConfig({
  plugins: [vue(), tailwindcss()],
  build: {
    rollupOptions: {
      output: {
        manualChunks: {
          // Separate Three.js into its own chunk (~600KB)
          'three': ['three'],
          // Separate GSAP into its own chunk (~100KB)
          'gsap': ['gsap'],
          // Separate postprocessing library if used
          'postprocessing': ['postprocessing'],
        }
      }
    }
  }
})
