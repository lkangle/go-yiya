import {createApp} from 'vue'
import App from './App.vue'
import { loadEnvironment } from './utils/platform';
import './style.css';

async function setup() {
    await loadEnvironment()

    let app = createApp(App)
    app.mount('#app')
}

setup()