<template>
  <a-layout id="components-layout-demo-custom-trigger">
    <a-layout-sider v-model="collapsed" :trigger="null" collapsible>
         <div class="logo">
            <a-icon type="project" v-show="!collapsed"/> &nbsp;{{collapsed ? 'Blog' : 'My Blog'}}
        </div>
        <a-menu theme="dark" mode="inline"  @click="goToPage">
            
            <a-menu-item key="index" >
                <a-icon type="dashboard" />
                <span>仪表盘</span>
            </a-menu-item>
            
            <a-sub-menu >
                <span slot="title"><a-icon type="file"></a-icon><span>文章管理</span></span>
                <a-menu-item key="addart"><a-icon type="form"></a-icon><span>写文章</span></a-menu-item>
                <a-menu-item key="artlist"><a-icon type="snippets"></a-icon><span>文章列表</span></a-menu-item>
            </a-sub-menu>
            <a-menu-item  key="catelist">
                <a-icon type="book" />
                <span>分类列表</span>
            </a-menu-item>
            <a-menu-item key="userlist">
                <a-icon type="upload" />
                <span>用户列表</span>
            </a-menu-item>
                      <a-menu-item key="profile">
                <a-icon type="setting" />
                <span>个人设置</span>
            </a-menu-item>
        </a-menu>
    </a-layout-sider>
    <a-layout>
      <a-layout-header style="background: #fff; padding: 0">
        <a-icon
          class="trigger"
          :type="collapsed ? 'menu-unfold' : 'menu-fold'"
          @click="() => (collapsed = !collapsed)"
        />
        <a-button type="danger" @click="loginOut" >退出</a-button>
      </a-layout-header>
      <a-layout-content
        :style="{ margin: '24px 16px', padding: '24px', background: '#fff', minHeight: '280px' }"
      >
        <router-view :key="$router.path"></router-view>
      </a-layout-content >
      <a-layout-footer style="background: #fff; padding: 0">
          <Footer></Footer>
      </a-layout-footer>
    </a-layout>
  </a-layout>
</template>

<script>
import Nav from "../components/admin/Nav.vue"
import Footer from "../components/admin/Footer.vue"
export default {
  data() {
    return {
      collapsed: false,
    };
  },
  methods: {
      loginOut(){
            window.sessionStorage.clear('token')
            this.$router.push('/login')
        },
      goToPage(item){
        this.$router.push("/"+ item.key).catch(err=>err)
      }    
  },
  components:{
      Nav,
      Footer
  }
};
</script>

<style>
    #components-layout-demo-custom-trigger {
        height: 100%;
    }
    #components-layout-demo-custom-trigger .trigger {
    font-size: 18px;
    line-height: 64px;
    padding: 0 24px;
    cursor: pointer;
    transition: color 0.3s;
    }

    #components-layout-demo-custom-trigger .trigger:hover {
    color: #1890ff;
    }

    #components-layout-demo-custom-trigger .logo {
    height: 32px;
    background: rgba(255, 255, 255, 0.2);
    margin: 16px;
    }
    .logo{
        color: #fff;
        text-align: center;
        line-height: 32px;
    }
</style>