import { createApp } from 'vue'
import './style.css'
import App from './App.vue'
import router from './router' // Import router
import scrollAnimation from './directives/scroll-animation'

const app = createApp(App)

app.use(router) // Gunakan router
app.directive('scroll-animation', scrollAnimation);

app.mount('#app')