<template>
   <div>
       
        <a-button type="danger" class="headerBtn" @click="loginOut">退出</a-button>
   </div>
</template>

<script>
export default {
    data(){
         return {
           collapsed: false,
        };
    },
    methods: {
        loginOut(){
            window.sessionStorage.clear('token')
            this.$router.push('/login')
        }    
    },
}
</script>

<style scoped>
    
 
</style>