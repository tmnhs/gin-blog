<template>
    <div>
        <h3>文章列表</h3>
        <a-card>
            <a-row  :gutter="20">
                <a-col :span="6">
                    <a-input-search placeholder="请输入文章标题查找" enter-button  v-model="querydata.title" @search="findArticle"></a-input-search>
                </a-col>
                <a-col :span="4">
                    <a-button type="primary" @click="$router.push(`/addart`)">新增</a-button>
                </a-col>
                 <a-col :span="4" :offset="10">
                     <a-select defaultValue="请选择分类" style="width:130px" @change="CateChange">
                        <a-select-option v-for="item in Catelist" :key="item.id" :value="item.id" >
                            {{item.name}}
                       </a-select-option>
                     </a-select>

                </a-col>
            </a-row>
            <a-table
               :rowKey="(record,index)=>{return index}"
               :columns="columns"
               :pagination='paginationOption'
               :dataSource="Articlelist"
               bordered
                >
                    <span slot="img" slot-scope="img">
                        <img :src="img" style="width:120px;height:90px">
                    </span>
                    <template slot="action" slot-scope="data">
                        <div class="actionSlot">
                        <a-button type="primary" style="margin-right:15px" @click="$router.push(`/addart/${data.id}`)">编辑</a-button>
                        <a-button type="danger" @click="deleteArticle(data.id)">删除</a-button>
                    </div>
                    </template>
                </a-table>
        </a-card>

    </div>
</template>

<script>
const columns=[
 {
    title:'ID',
    dataIndex: 'id',
    width:'10%',
    key: 'id',
    align:'center',
  },
   {
    title:'分类',
    dataIndex: 'name',
    width:'20%',
    key: 'name',
    align:'center',
  },
   {
    title:'文章标题',
    dataIndex: 'title',
    width:'20%',
    key: 'title',
    align:'center',

  },

   {
    title:'文章描述',
    dataIndex: 'desc',
    width:'20%',
    key: 'desc',
    align:'center',
  },
   {
    title:'缩略图',
    dataIndex: 'img',
    width:'20%',
    key: 'img',
    scopedSlots:{customRender:'img'},
    align:'center',
  },
   {
    title:'操作',
    width:'20%',
    key: 'action',
    scopedSlots:{customRender:'action'},
    align:'center',
  },

]
export default {
    data(){
        return {
            labelCol: { span: 5 },
            wrapperCol: { span: 16 },
            paginationOption:{
                pageSizeOptions:['5','10','20'],
                defaultCurrent:1,
                defaultPageSize:5,
                total :0,
                showSizeChanger:true,
                showTotal:(total)=>`共${total}条`,
                onChange:(current,size)=>{
                    this.paginationOption.defaultCurrent=current
                    this.paginationOption.defaultPageSize=size
                    this.getArticleList()

                },
                onShowSizeChange:(current,size)=>{
                    this.paginationOption.defaultCurrent=1
                    this.paginationOption.defaultPageSize=size
                    this.getArticleList()
              },
            },
            Articlelist:[],
            Catelist:[],
            columns,
            querydata:{
                pagesize:0,
                pagenum:0,
                title:'',
            },

        }
    },
    created(){
            this.getArticleList()
            this.getCateList()
    },
    methods: {
        gotoEditArt(id){
            alert(id)
            this.$router.push(`/addart/`+id)
        },
        //文章列表
        async getArticleList(){
             this.querydata.pagesize=this.paginationOption.defaultPageSize
             this.querydata.pagenum=this.paginationOption.defaultCurrent
             const{data:result} = await this.$http.post('articles',this.querydata)
             if (result.status!=200) return this.$message.error(result.message)
             this.Articlelist=result.data

             this.paginationOption.total=result.total
      },
      async findArticle(){
             this.querydata.pagesize=this.paginationOption.defaultPageSize
             this.querydata.pagenum=1
             console.log(this.querydata)
             const{data:result} = await this.$http.post('articles',this.querydata)
             if (result.status!=200) return this.$message.error(result.message)
             this.Articlelist=result.data
             this.paginationOption.total=result.total
      },
      //删除文章
      deleteArticle(id){
          this.$confirm({
              title:"确定要删除吗？",
              content:"一旦删除，无法恢复",
              onOk:async()=>{
                    const {data:result}=await this.$http.delete(`article/${id}`)
                    if(result.status!=200)return this.$message.error(result.message)
                    this.$message.success('删除成功')
                    this.getArticleList()
              },
              onCancel:()=>{
                  this.$message.info("已取消删除")
              }
          })
      },
     //获取分类列表
      async getCateList(){
             const{data:result} = await this.$http.get('category')
             if (result.status!=200) return this.$message.error(result.message)
             this.Catelist=result.data
             this.paginationOption.total=result.total
      },
      //查询分类下的文章
      CateChange(value){
            this.getCateArticle(value)
      },
      async getCateArticle(id){
            const {data:result}=await this.$http.get(`article/cate/${id}`,{
                params:{
                    pagesize:this.paginationOption.defaultPageSize,
                    pagenum:this.paginationOption.defaultCurrent
                }
            })
             if (result.status!=200) return this.$message.error(result.message)
             this.Articlelist=result.data
             this.paginationOption.total=result.total
      }
    },
}
</script>

<style scoped>
    .actionSlot{
        display: flex;
        justify-content: center;
    }
</style>
