import {createApp} from 'vue'
import App from './App.vue'
import { loadEnvironment } from './utils/platform';
import { createPinia } from 'pinia';
import './style.css';

async function setup() {
    await loadEnvironment()
    
    let pinia = createPinia()
    let app = createApp(App)

    app.use(pinia)
    app.mount('#app')
}

setup()