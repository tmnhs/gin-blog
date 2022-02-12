<template>
    <div>
        <a-card>
            <h3>{{ id ? '编辑文章':'新增文章' }}</h3>
            <a-form-model :model="articleInfo" ref="articleInfoRef" :rules="articleInfoRules" :hideRequiredMark="true">
                <a-form-model-item label="文章标题" prop="title">
                    <a-input v-model="articleInfo.title" style="width:300px"></a-input>
                </a-form-model-item>
                 <a-form-model-item  label="文章分类" prop="cid">
                    <a-select v-model="articleInfo.cid" style="width:130px" placeholder="请选择分类" @change="cateChange">
                        <a-select-option v-for="item in Catelist" :key="item.id" :value="item.id" >
                            {{item.name}}
                       </a-select-option>
                     </a-select>
                </a-form-model-item>
                 <a-form-model-item label="文章描述" prop="desc">
                    <a-input v-model="articleInfo.desc" type="textarea"></a-input>
                </a-form-model-item>
                 <a-form-model-item label="文章缩略图" prop="img" >
                     <a-upload
                         name="file"
                        :action="upUrl"
                        :headers="headers"
                        listType="picture"
                        >
                         <a-button> <a-icon type="upload" /> 点击上传 </a-button>
                         <br>
                         <template v-if="id">
                             <img :src="articleInfo.img" style="width:120px;height:90px">
                         </template>
                     </a-upload>
                </a-form-model-item>
                 <a-form-model-item label="文章内容" props="content" >
                    <Editor v-model="articleInfo.content" type="textarea"></Editor>
                </a-form-model-item>
                 <a-form-model-item >
                    <a-button type="danger" style="margin-right:15px" @click="artOk(articleInfo.id)">
                        {{articleInfo.id?'更新':'提交'}} 
                    </a-button>
                    <a-button type="primary" @click="artCancel"> 取消</a-button>
                </a-form-model-item>
            </a-form-model>
        </a-card>
    </div>
</template>

<script>
import Editor from '../editor/index.vue'
import {Url} from '../../plugin/http'
export default {
    components:{
        Editor,
        },
    props:['id'],
    data(){
        return {
            articleInfo:{
                id:0,
                title:'',
                cid:undefined,
                desc:'',
                content:'',
                img:'',
            },
            Catelist:[],
            upUrl:Url+'upload',
            headers:{},
            articleInfoRules:{
                    title:[{required:true,message:'请输入文章标题',tirgger:'blur'}],
                    // cid:[{required:true,message:'请选择文章分类',tirgger:'blur'}],
                    desc:[{required:true,message:'请输入文章描述',tirgger:'blur'},{max:120,message:'文章描述字数不能超过120',tirgger:'blur'}],
                    content:[{required:true,message:'请输入文章内容',tirgger:'blur'}],
            },
        }
    },
    created() {
        this.getCateList()
        this.headers={ Authorization:`Bearer ${window.sessionStorage.getItem("token")}`}
        if (this.id){
                this.getArtInfo(this.id)
        }
  },
    methods: {
     //查询文章信息
     async getArtInfo(id){
            const {data:result}=await this.$http.get(`article/info/${id}`)
             if (result.status!=200) return this.$message.error(result.message)
            this.articleInfo=result.data
            this.articleInfo.id=this.id
     },
     //获取分类列表
      async getCateList(){
             const{data:result} = await this.$http.get('category')
             if (result.status!=200) return this.$message.error(result.message)
             this.Catelist=result.data
      },
      //选择分类
      cateChange(value){
        this.articleInfo.cid=value
      },
      //上传
      uploadChange(info){
           if (info.file.status !== 'uploading') {
                console.log(info.file, info.fileList);
                const imgUrl = info.file.response.url
                this.articleInfo.img=imgUrl
         }
            if (info.file.status === 'done') {
                this.$message.success("图片上传成功");
            } else if (info.file.status === 'error') {
                this.$message.error("图片上传失败");
            }
        },
        //提交或者用户
        async artOk(id){
                this.$refs.articleInfoRef.validate(async (valid)=>{
                    if(!valid) return this.$message.error("输入非法内容，请重新输入")
                    if (id==0){
                        const {data:result} = await this.$http.post('article/add',this.articleInfo)
                        if (result.status!=200) return this.$message.error(result.message)
                         this.$message.success("文章添加成功")
                         this.$router.push('/artlist')

                    }else {
                        const {data:result} = await this.$http.put(`article/${id}`,this.articleInfo)
                        if (result.status!=200) return this.$message.error(result.message)
                         this.$message.success("更新文章成功")
                         this.$router.push('/artlist')
                    }
                })
        },
        artCancel(){
            this.$refs.articleInfoRef.resetFields()
            this.$router.push('/artlist')
            
        }
    },
}
</script>