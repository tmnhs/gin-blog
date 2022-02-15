<template>
    <div>
        <div class="d-flex justify-center pa-3 text-h3 font-weight-blod">{{articleInfo.title}}</div>
        <div class="d-flex justify-center">
            <v-icon>
                    {{'mdi-calendar-month'}}
            </v-icon>
        <span>{{articleInfo.CreatedAt | dateformat('YYYY-MM-DD HH::mm')}}</span>
        </div>
      
        <v-divider class="pa-3 ma-3"></v-divider>
        <v-alert elevation="1" class="ma-4"  color="indigo" dark border="left" outlined> {{articleInfo.desc}}</v-alert>
        <div class="ma-5 pa-3 text-justify" v-html="articleInfo.content"></div>
   </div>
</template>

<script>
export default {
    props:['id'],
    data(){
        return {
            articleInfo:{}
        }
    },
    created(){
        this.getArtInfo()
    },
    methods:{
        //查询文章基本信息
        async getArtInfo(){
            const {data:result}=await this.$http.get(`article/info/${this.id}`)
             if (result.status!=200) return 
            this.articleInfo=result.data
        }
    }
}
</script>

<style scoped>
    
</style>