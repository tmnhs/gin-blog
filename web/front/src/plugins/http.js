import Vue from 'vue'
import axios from 'axios'

axios.defaults.baseURL = 'http://localhost/api/v1'

Vue.prototype.$http=axios