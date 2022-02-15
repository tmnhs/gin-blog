<template>
    <div>
        <h3>分类列表</h3>
        <a-card>
            <a-row  :gutter="20">
                <a-col :span="4">
                    <a-button type="primary" @click="addCateVisible=true" >新增</a-button>
                </a-col>
            </a-row>
            <a-table
                rowKey="ID"
               :columns="columns"
               :pagination='paginationOption'
               :dataSource="Catelist"
               bordered
                >
                    <template slot="action" slot-scope="data">
                        <div class="actionSlot">
                        <a-button type="primary" style="margin-right:15px" @click="editCate(data.ID)">编辑</a-button>
                        <a-button type="danger" @click="deleteCate(data.ID)">删除</a-button>
                    </div>
                    </template>
                </a-table>
        </a-card>
        <!-- 新增分类区 -->
          <a-modal
            closable
            title="新增分类"
            :visible="addCateVisible"
            @ok="addCateOk"
            @cancel="addCateCancel"
            :destroyOnClose="true"
            >
                <a-form-model
                :model="CateInfo" 
                :rules="CateRules"
                ref="addCateRef"
                :label-col="labelCol"
                :wrapper-col="wrapperCol"
                >
                    <a-form-model-item has-feedback label="分类名" prop="name">
                        <a-input  v-on:keyup.enter="addCateOk"  v-model="CateInfo.name"></a-input>
                    </a-form-model-item>
                </a-form-model>
          </a-modal>
           <!-- 编辑分类区 -->
          <a-modal
            closable
            title="编辑分类"
            :visible="editCateVisible"
            @ok="editCateOk"
            @cancel="editCateCancel"
            >
                <a-form-model
                :model="CateInfo" 
                :rules="CateRules"
                ref="editCateRef"
                :label-col="labelCol"
                :wrapper-col="wrapperCol"
                >
                    <a-form-model-item has-feedback label="分类名" prop="name">
                        <a-input  v-on:keyup.enter="editCateOk"   v-model="CateInfo.name"></a-input>
                    </a-form-model-item>
                </a-form-model>
          </a-modal>
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
    title:'分类名',
    dataIndex: 'name',
    width:'20%',
    key: 'name',
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
                    this.getCateList()

                },
                onShowSizeChange:(current,size)=>{
                    this.paginationOption.defaultCurrent=1
                    this.paginationOption.defaultPageSize=size
                    this.getCateList()
              },
            },
            Catelist:[],
            columns,
            querydata:{
                pagesize:0,
                pagenum:0,
            },
            addCateVisible:false,
            editCateVisible:false,
            CateInfo:{
                id:0,
                name:'',
            },
            CateRules:{
                name:[{validator:(rule,value,callback)=>{
                    if(this.CateInfo.name==""){
                        callback(new Error("请输入分类名"))
                    }else{
                        callback()
                    }
                },
                trigger:'blur'}],
            }
        }
    },
    created(){
            this.getCateList()
    },
    methods: {
        //分类列表
      async getCateList(){
             this.querydata.pagesize=this.paginationOption.defaultPageSize
             this.querydata.pagenum=this.paginationOption.defaultCurrent
             const{data:result} = await this.$http.get('category',this.querydata)
             if (result.status!=200) return this.$message.error(result.message)
             this.Catelist=result.data
             this.paginationOption.total=result.total
      },
      //删除分类
      deleteCate(id){
          this.$confirm({
              title:"确定要删除吗？",
              content:"一旦删除，无法恢复",
              onOk:async()=>{     
                    const {data:result}=await this.$http.delete(`category/${id}`)
                   if(result.status!=200)return this.$message.error(result.message)
                    this.$message.success('删除成功')
                    this.getCateList()
              },
              onCancel:()=>{
                  this.$message.info("已取消删除")
              }
          })
      },
      //新增分类
      addCateOk(){
          this.$refs.addCateRef.validate(async (valid)=>{
              if(!valid)return this.$message.error("输入不符合要求，请重新输入")
              const {data:result}= await this.$http.post('category/add',this.CateInfo)
              if(result.status!=200)return this.$message.error(result.message)
              this.$message.success("添加成功")
              this.addCateVisible=false
              await this.getCateList()
          })
      },
      addCateCancel(){
          this.$refs.addCateRef.resetFields()     
          this.addCateVisible=false,
          this.$message.info("新增分类已取消")
      },
      //编辑分类
       async editCate(id){
          this.editCateVisible=true
          const {data:result} = await this.$http.get(`category/${id}`)
          this.CateInfo.name=result.data.name
          this.CateInfo.id=id
      },
      editCateOk(){
            this.$refs.editCateRef.validate(async (valid)=>{
              if(!valid)return this.$message.error("输入不符合要求，请重新输入")
              const {data:result}= await this.$http.put(`category/${this.CateInfo.id}`,this.CateInfo)
              if(result.status!=200)return this.$message.error(result.message)
              this.$message.success("分类信息更新成功")
              this.editCateVisible=false
              this.CateInfo.name=''
              this.getCateList()
          })
      },
      editCateCancel(){
          this.editCateVisible=false
          this.CateInfo.name=''
          this.$message.info("编辑已取消")
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