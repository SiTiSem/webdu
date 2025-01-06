import { createApp } from 'vue'
import {
    Alert,
    Button,
    Card,
    Col, Flex, Form, Input,
    Layout,
    Menu,
    Progress,
    Row,
    Table,
    Tooltip,
    Tree
} from 'ant-design-vue';
import App from './App.vue'
import router from './router'
import axios from "axios";

export const HTTP = axios.create();

// HTTP.defaults.baseURL = 'http://localhost:7000';

HTTP.interceptors.request.use(config => {
    const token = localStorage.getItem('tks');
    config.headers['Authorization'] = 'Bearer ' + token;
    return config
});

HTTP.interceptors.response.use(undefined, (error) => {
    if (error.response && (error.response.status === 403 || error.response.status === 401)) {
        localStorage.removeItem('tks')
        router.push({
            name: 'Login'
        });
    }
    console.log(error)
    return Promise.reject(error.response.data);
});

const app = createApp(App)
app
    .use(router)
    .use(Layout)
    .use(Flex)
    .use(Form)
    .use(Input)
    .use(Tooltip)
    .use(Menu)
    .use(Progress)
    .use(Row)
    .use(Col)
    .use(Table)
    .use(Button)
    .use(Card)
    .use(Tree)
    .use(Alert)

app.config.globalProperties.$filters = {
    bytesToSize(bytes) {
        const sizes = ['B', 'KB', 'MB', 'GB', 'TB'];
        if (bytes === 0) return '0 B';
        const i = Math.floor(Math.log2(bytes) / 10);
        return `${(bytes / 2 ** (10 * i)).toFixed(0)} ${sizes[i]}`;
    }
}

app.mount('#app')
