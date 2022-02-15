<template>
        <v-col>
                <v-card class="ma-3" v-for="item in Articlelist" :key="item.id" link @click="$router.push(`/home/detail/${item.id}`)">
                        <v-row no-gutters>
                                <v-col class="d-flex justify-center align-center mx-3" cols="2" >
                                       <v-img max-height="100" max-width="100" :src="item.img"></v-img> 
                                </v-col>
                                <v-col>
                                        <v-card-title class="my-2"><v-chip color="pink" label class="mr-3 white--text">{{item.name}}</v-chip> 
                                        {{item.title}}
                                        </v-card-title>
                                        <v-card-subtitle v-text="item.desc"></v-card-subtitle>
                                        <v-divider></v-divider>
                                        <v-card-text>
                                                <v-icon>
                                                        {{'mdi-calendar-month'}}
                                                </v-icon>
                                                <span>{{item.CreatedAt | dateformat('YYYY-MM-DD HH::mm')}}</span>
                                        </v-card-text>
                                </v-col>
                        </v-row>
                </v-card>
                <div class="text-center">
                        <v-pagination
                        total-visible="7"
                         v-model="querydata.pagenum"
                         :length="Math.ceil(this.total/querydata.pagesize) "
                        @input="getCateArticle()"
                         >

                        </v-pagination>
                </div>
        </v-col>
</template>

<script>
export default { 
     props:['id'],
     data(){
        return {
            Articlelist:[],
            querydata:{
                pagesize:5,
                pagenum:1,
            },
            total:0,
        }
    },
    created(){
        this.getCateArticle()
    },
    methods:{
       async getCateArticle(){
               if (this.id){
                   this.$store.commit("saveCateId",this.id)
               }
            const {data:result}=await this.$http.get(`article/cate/${this.$store.state.cateId}`,{
                params:{
                    pagesize:this.querydata.pagesize,
                    pagenum:this.querydata.pagenum
                }
            })
             if (result.status!=200) return 
             this.Articlelist=result.data
             this.total=result.total
      }
    }
}
</script>

<style scoped>

</style>