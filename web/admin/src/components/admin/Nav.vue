<template>
    <div>
    <a-layout-sider v-model="collapsed" :trigger="null" collapsible>
            <div class="logo">
            <a-icon type="project" /> &nbsp;My Blog 
        </div>
        <a-menu theme="dark" mode="inline">
            
            <a-menu-item >
                <a-icon type="dashboard" />
                <span>仪表盘</span>
            </a-menu-item>
            <a-sub-menu title="文章管理">
                <span slot="title"><a-icon type="file"></a-icon><span>文章管理</span></span>
                <a-menu-item><a-icon type="form"></a-icon><span>写文章</span></a-menu-item>
                <a-menu-item><a-icon type="snippets"></a-icon><span>文章列表</span></a-menu-item>
            </a-sub-menu>
            <a-menu-item >
                <a-icon type="book" />
                <span>分类列表</span>
            </a-menu-item>
            <a-menu-item>
                <a-icon type="upload" />
                <span>用户列表</span>
            </a-menu-item>
        </a-menu>
    </a-layout-sider>
    
    </div>
</template>

<script>
export default {
    data(){
         return {
           collapsed: false,
        };
    },
}
</script>

<style scoped>
     .logo{
        color: #fff;
        text-align: center;
        line-height: 32px;
        height: 32px;
        background: rgba(255, 255, 255, 0.2);
        margin: 16px;
        }
</style>