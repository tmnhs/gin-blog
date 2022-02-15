<template>
    <div>
        <h3>个人设置</h3>
        <a-card>
            <a-form-model :model="profileInfo" :hideRequiredMark="true">
                <a-form-model-item label="作者名称" >
                    <a-input v-model="profileInfo.name"  style="width:300px"></a-input>
                </a-form-model-item>
                    <a-form-model-item label="个人简介" >
                    <a-input v-model="profileInfo.desc" style="width:300px"></a-input>
                </a-form-model-item>
                    <a-form-model-item label="QQ" >
                    <a-input v-model="profileInfo.qq_chat" style="width:300px"></a-input>
                </a-form-model-item>
                    <a-form-model-item label="微信" >
                    <a-input v-model="profileInfo.wechat" style="width:300px"></a-input>
                </a-form-model-item>
                    <a-form-model-item label="微博" >
                    <a-input v-model="profileInfo.weibo" style="width:300px"></a-input>
                </a-form-model-item>
                    <a-form-model-item label="Email" >
                    <a-input v-model="profileInfo.email" style="width:300px"></a-input>
                </a-form-model-item>
                 <a-form-model-item label="头像">
                     <a-upload 
                         name="file"
                         :multiple="false"
                        :action="upUrl"
                        :headers="headers"
                        listType="picture"
                        @change="avatarChange"
                        >
                         <a-button> <a-icon type="upload" /> 点击上传 </a-button>
                         <br>
                         <template v-if="profileInfo.id">
                             <img :src="profileInfo.avatar" style="width:120px;height:90px">
                         </template>
                     </a-upload>
                </a-form-model-item>
                   <a-form-model-item label="头像背景">
                     <a-upload
                         name="file"
                         :multiple="false"
                        :action="upUrl"
                        :headers="headers"
                        listType="picture"
                        @change="imgChange"
                        >
                         <a-button> <a-icon type="upload" /> 点击上传 </a-button>
                         <br>
                         <template v-if="profileInfo.id">
                             <img :src="profileInfo.img" style="width:120px;height:90px">
                         </template>
                     </a-upload>
                </a-form-model-item>
                 <a-form-model-item >
                    <a-button type="danger" style="margin-right:15px" @click="updateProfile()">
                      更新
                    </a-button>
                </a-form-model-item>
            </a-form-model>
        </a-card>
    </div>
</template>


<script>
import {Url} from '../../plugin/http'
export default {
    data(){
        return{
            profileInfo:{
                id:0,
                name:'',
                desc:'',
                qq_chat:'',
                wechat:'',
                weibo:'',
                email:'',
                img:'',
                avatar:'',
            },
            upUrl:Url+'/upload',
            headers:{},
        }
    },
    created() {
        this.headers={ Authorization:`Bearer ${window.sessionStorage.getItem("token")}`}
        this.getProfile()
    },
    methods:{
          async getProfile(){
             const{data:result} = await this.$http.get(`profile`)
             if (result.status!=200) return this.$message.error(result.message)
             this.profileInfo=result.data
         },
         avatarChange(info){
                if (info.file.status !== 'uploading') {
                 console.log(info.file, info.fileList);
                  }
                if (info.file.status === 'done') {
                     const imgUrl = info.file.response.url
                      this.profileInfo.avatar=imgUrl
                    this.$message.success("图片上传成功");
                } else if (info.file.status === 'error') {
                    this.$message.error("图片上传失败");
                }
            },
            imgChange(info){
                if (info.file.status !== 'uploading') {
                console.log(info.file, info.fileList);
 }
                if (info.file.status === 'done') {
                    const imgUrl = info.file.response.url
                     this.profileInfo.img=imgUrl
                    this.$message.success("图片上传成功");
                } else if (info.file.status === 'error') {
                    this.$message.error("图片上传失败");
                }
            },
            async updateProfile(){
                const{data:result} = await this.$http.put(`profile`,this.profileInfo)
                if (result.status!=200) return this.$message.error(result.message)
                this.profileInfo=result.data
                this.$router.push('/index')
            }
    }
}
</script>

<style scoped>
    
</style>