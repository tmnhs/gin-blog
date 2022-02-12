import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/Home.vue'
import ArticleList from '../components/ArticleList.vue'
import Detail from '../components/Detail.vue'


Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home,
    children:[
      {
        path:'/',
        component:ArticleList,
        meta:{title:"欢迎来到GinBlog",},
      },
   
      {
        path:`detail/:id`,
        component:Detail,
        props:true,
        meta:{title:"文章详情",},
      },
    ]
  },
  {
    path:`/:id`,
    component:ArticleList,
    meta:{title:"欢迎来到GinBlog",},
    props:true
  },

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
  next()
})
export default router
