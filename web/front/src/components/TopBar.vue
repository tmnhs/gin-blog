<template>
    <div>
        <v-app-bar app color="blue darken-2" flat>
                <v-avatar class="mx-12"  size="40" color="grey" >
                    <img src="../assets/bg1.png" >
                </v-avatar>
                <v-container class="py-0 fill-height" @click="$router.push('/')">
                        <v-btn text color="white">分类</v-btn>
                        <v-btn v-for="item in Catelist" :key="item.id" text color="white"@click="$router.replace(`/${item.id}`)">{{item.name}}</v-btn>
                </v-container>
                <v-spacer></v-spacer>

                <v-responsive max-width="260" color="white">
                        <v-text-field dense flag hide-details rounded solo-inverted>

                        </v-text-field>
                </v-responsive>
        </v-app-bar>
    </div>
</template>

<script>
export default {
    data(){
            return{
                Catelist:[],
            }
    },
    created(){
        this.getCateList()
    },
    methods:{
        //获取分类
          //获取分类列表
      async getCateList(){
             const{data:result} = await this.$http.get('category')
             if (result.status!=200) return this.$message.error(result.message)
             this.Catelist=result.data
      },
      goHome(){
          
      }
    }
}
</script>

<style scoped>

</style>