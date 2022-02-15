<template>
   <v-card class="mx-auto" max-width="500">
       <v-img :src="profileInfo.img">
           <v-card-title>
               <v-col align="center">
                   <v-avatar size="90" color="grey">
                       <img :src="profileInfo.avatar" >
                   </v-avatar>
                   <div class="ma-4 white--text">{{profileInfo.name}}</div>
               </v-col>
           </v-card-title>
           <v-divider> </v-divider>
       </v-img>
           <v-col>
               <div class="ma-3">About Me :</div>
                <div class="ma-3">{{profileInfo.desc}}</div>
           </v-col>
           <v-divider color="indigo"> </v-divider>
            
            <v-list nav dense>
                    <v-list-item>
                        <v-list-item-icon class="ma-3">
                            <v-icon color="blue darkern-2">{{`mdi-qqchat`}}</v-icon>
                        </v-list-item-icon>
                        <v-list-item-content class="grey-text">{{profileInfo.qq_chat}} </v-list-item-content>
                    </v-list-item>
                     <v-list-item>
                        <v-list-item-icon class="ma-3">
                            <v-icon color="green darken-2">{{`mdi-wechat`}}</v-icon>
                        </v-list-item-icon>
                        <v-list-item-content class="grey-text">{{profileInfo.wechat}}</v-list-item-content>
                    </v-list-item>
                    <v-list-item>
                        <v-list-item-icon class="ma-3">
                            <v-icon color="orange darken-2">{{`mdi-sina-weibo`}}</v-icon>
                        </v-list-item-icon>
                        <v-list-item-content class="grey-text">{{profileInfo.weibo}}</v-list-item-content>
                    </v-list-item>
                     <v-list-item>
                        <v-list-item-icon class="ma-3">
                            <v-icon color="yellow dorken-2">{{`mdi-email`}}</v-icon>
                        </v-list-item-icon>
                        <v-list-item-content class="grey-text">{{profileInfo.email}}
                        </v-list-item-content>
                    </v-list-item>
            </v-list>
   </v-card>
</template>

<script>
export default {
    data(){
        return{
           profileInfo:{
                id:1,
                name:'',
                desc:'',
                qq_chat:'',
                wechat:'',
                weibo:'',
                email:'',
                img:'',
                avatar:'',
            },
        }
    },
    created(){
      this.getProfile()  
    },
    methods:{
        async getProfile(){
             const{data:result} = await this.$http.get(`profile`)
             if (result.status!=200) return 
             this.profileInfo=result.data
            //  console.log(this.profileInfo)
         },
    }
}
</script>

<style scoped>

</style>