import Vue from 'vue'
import axios from 'axios'

axios.defaults.baseURL = 'http://localhost/api/v1'

axios.defaults.withCredentials=true
axios.interceptors.request.use(config=>{
  config.headers.Authorization=`Bearer ${window.sessionStorage.getItem("token")}`
  return config
})
Vue.prototype.$http=axios


