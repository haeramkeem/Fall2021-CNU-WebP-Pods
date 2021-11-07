import '@babel/polyfill';
import 'mutationobserver-shim';
import './style/App.css';
import './plugins/api-mock';
import { createApp } from 'vue';
import router from './router';
import App from './App.vue';

createApp(App).use(router).mount('#app');
