import './assets/main.css'
import content from "./components/content.vue"

import { createApp } from 'vue'
import App from './App.vue'

const app = createApp(App)

app.component({
    "content": content
})

app.mount('#app')
