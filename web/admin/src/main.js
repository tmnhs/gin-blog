import Vue from 'vue'
import App from './App.vue'
import router from './router'
import './plugin/antui'
import './assets/css/style.css'
import store from './store'       //引入vuex
Vue.prototype.$store = store      //引入vuex



Vue.config.productionTip = false

new Vue({
  // store,
  router,
  render: h => h(App)
}).$mount('#app')
