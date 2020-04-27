import Vue from 'vue';
import Antd from 'ant-design-vue';
import App from './App';
import axios from 'axios';
import 'ant-design-vue/dist/antd.css';
import router from './router'

Vue.config.productionTip = false;

Vue.use(Antd);

export const HTTP = axios.create();

HTTP.interceptors.request.use(config => {
  const token = localStorage.getItem('tks');
  config.headers.common['Authorization'] = 'Bearer ' + token;
  return config
});


HTTP.interceptors.response.use(undefined, (error) => {
  if (error.response && (error.response.status === 403 || error.response.status === 401)) {
    localStorage.removeItem('tks')
    router.push({
      name: 'Login'
    });
  }
  return Promise.reject(error.response.data);
});

new Vue({
  render: h => h(App),
  router,
}).$mount('#app')
