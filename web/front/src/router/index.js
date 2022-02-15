import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/Home.vue'
import ArticleList from '../components/ArticleList.vue'
import Detail from '../components/Detail.vue'
import Login from '../views/Login.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/login',
    name: 'login',
    component: Login,
    meta:{title:"请先登录",},

  },
  {
    path: '/home',
    name: 'home',
    component: Home,
    children:[
      {
        path:'',
        component:ArticleList,
        meta:{title:"欢迎来到GinBlog",},
      },
   
      {
        path:`detail/:id`,
        component:Detail,
        props:true,
        meta:{title:"文章详情",},
      },
      {
        path:`:id`,
        component:ArticleList,
        meta:{title:"欢迎来到GinBlog",},
        props:true
      },
    ]
  },
  {
    path: "*",
    redirect: "login"
}

]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})
router.beforeEach((to,from,next)=>{
  if (to.meta.title){
    document.title=to.meta.title
  }
  const token = window.sessionStorage.getItem('token')
  console.log("token",token)
  if (to.path=="/")next("/login")
  if(to.path==='/login')return next()
  if(!token &&to.path=='/home'){
    // Vue.prototype.$message.error("请先登录")
    next("/login")
  }else{
    next()
  }
})

export default router
