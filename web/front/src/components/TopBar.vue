<template>
    <div>
        <v-app-bar app color="blue darken-2" flat>
                <v-avatar class="mx-12"  size="40" color="grey" >
                    <img :src="profileInfo.avatar" @click="$router.go(0)">
                </v-avatar>
                <v-container class="py-0 fill-height" @click="$router.go(0)">
                        <v-btn text color="white" @click="$router.push(`/home/0`)">首页</v-btn>
                        <v-btn v-for="item in Catelist" :key="item.id" text color="white"@click="$router.push(`/home/${item.id}`)">{{item.name}}</v-btn>
                </v-container>
                <v-btn color="error" @click.prevent="loginOut">退出</v-btn>
                <!-- <v-responsive max-width="260" color="white">
                        <v-text-field dense flag hide-details rounded solo-inverted>

                        </v-text-field>
                </v-responsive> -->
        </v-app-bar>
    </div>
</template>

<script>
export default {
    data(){
            return{
                Catelist:[],
                profileInfo:{},
            }
    },
    created(){
        this.getCateList()
        this.getProfile()
    },
    methods:{
        //获取分类
          //获取分类列表
      async getCateList(){
             const{data:result} = await this.$http.get('category')
             if (result.status!=200) return 
             this.Catelist=result.data
      },
        async getProfile(){
            const{data:result} = await this.$http.get(`profile`)
            if (result.status!=200) return 
            this.profileInfo=result.data
            console.log(this.profileInfo)
        },
       loginOut(){
            window.sessionStorage.clear('token')
            this.$router.push('/login')
        },
    }
}
</script>

<style scoped>

</style>