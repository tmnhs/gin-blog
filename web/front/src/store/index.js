import Vue from "vue"
import Vuex from "vuex"
Vue.use(Vuex)
export default new Vuex.Store({
	//当做data
	state:{
        // collapsed:false
		cateId:-1
	},
	//相当于计算属性
	getters:{
	},
	//同步一些方法
	mutations:{
		saveCateId(state, cateId) {
			state.cateId= cateId;
		  },
	},
	//存放异步的方法
	actions:{

	}
})
