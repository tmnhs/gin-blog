<template>
    <div>
        <h3>用户列表</h3>
        <a-card>
            <a-row  :gutter="20">
                <a-col :span="6">
                    <a-input-search placeholder="请输入用户名或id查找" enter-button  v-model="querydata.idorname" @search="findUser"></a-input-search>
                </a-col>
                <a-col :span="4">
                    <a-button type="primary" @click="addUserVisible=true" >新增</a-button>
                </a-col>
            </a-row>
            <a-table
                rowKey="ID"
               :columns="columns"
               :pagination='paginationOption'
               :dataSource="userlist"
               bordered
                >
                    <span slot="role" slot-scope="role">{{role==1?'管理员':'订阅者'}}</span>
                    <template slot="action" slot-scope="data">
                        <div class="actionSlot">
                        <a-button type="primary" style="margin-right:15px" @click="editUser(data.ID)">编辑</a-button>
                        <a-button type="danger" @click="deleteUser(data.ID)">删除</a-button>
                    </div>
                    </template>
                </a-table>
        </a-card>
        <!-- 新增用户区 -->
          <a-modal
            closable
            title="新增用户"
            :visible="addUserVisible"
            @ok="addUserOk"
            @cancel="addUserCancel"
            :destroyOnClose="true"
            >
                <a-form-model
                :model="userInfo" 
                :rules="userRules"
                ref="addUserRef"
                :label-col="labelCol"
                :wrapper-col="wrapperCol"
                >
                    <a-form-model-item has-feedback label="用户名" prop="username">
                        <a-input v-on:keyup.enter="addUserOk"  v-model="userInfo.username"></a-input>
                    </a-form-model-item>
                    <a-form-model-item  has-feedback label="密码" prop="password">
                        <a-input v-on:keyup.enter="addUserOk"  v-model="userInfo.password"></a-input>
                    </a-form-model-item>
                    <a-form-model-item  has-feedback label="确认密码" prop="chcekpass">
                        <a-input v-on:keyup.enter="addUserOk"  v-model="userInfo.chcekpass"></a-input>
                    </a-form-model-item>
                    <a-form-model-item label="是否为管理员" prop="role">
                        <a-select defaultValue="2" style="width:60px" @change="adminChange">
                            <a-select-option key="1" value="1">是</a-select-option>
                            <a-select-option key="2" value="2">否</a-select-option>
                        </a-select>
                    </a-form-model-item>
                    <a-form-model-item  has-feedback  label="邮箱" prop="email"> 
                            <a-input v-on:keyup.enter="addUserOk" v-model="userInfo.email"></a-input>
                    </a-form-model-item>
                </a-form-model>
          </a-modal>
           <!-- 编辑用户区 -->
          <a-modal
            closable
            title="编辑用户"
            :visible="editUserVisible"
            @ok="editUserOk"
            @cancel="editUserCancel"
            >
                <a-form-model
                :model="userInfo" 
                :rules="userRules"
                ref="editUserRef"
                :label-col="labelCol"
                :wrapper-col="wrapperCol"
                >
                    <a-form-model-item has-feedback label="用户名" prop="username">
                        <a-input   v-on:keyup.enter="editUserOk" v-model="userInfo.username"></a-input>
                    </a-form-model-item>
                    <a-form-model-item label="是否为管理员" prop="role">
                        <a-select :defaultValue="String(userInfo.role)" style="width:60px" @change="adminChange">
                            <a-select-option key="1" value="1">是</a-select-option>
                            <a-select-option key="2" value="2">否</a-select-option>
                        </a-select>
                    </a-form-model-item>
                    <a-form-model-item  has-feedback  label="邮箱" prop="email"> 
                            <a-input  v-on:keyup.enter="editUserOk" v-model="userInfo.email"></a-input>
                    </a-form-model-item>
                </a-form-model>
          </a-modal>
    </div>
</template>

<script>
const columns=[
 {
    title:'ID',
    dataIndex: 'ID',
    width:'10%',
    key: 'id',
    align:'center',
  },
   {
    title:'用户名',
    dataIndex: 'username',
    width:'20%',
    key: 'username',
    align:'center',

  },
   {
    title:'角色',
    dataIndex: 'role',
    width:'20%',
    key: 'role',
    align:'center',
    scopedSlots:{customRender:'role'}
  },
   {
    title:'邮箱',
    dataIndex: 'email',
    width:'20%',
    key: 'email',
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
                    this.getUserList()

                },
                onShowSizeChange:(current,size)=>{
                    this.paginationOption.defaultCurrent=1
                    this.paginationOption.defaultPageSize=size
                    this.getUserList()
              },
            },
            userlist:[],
            columns,
            querydata:{
                pagesize:0,
                pagenum:0,
                idorname:'',
            },
            addUserVisible:false,
            editUserVisible:false,
            userInfo:{
                id:0,
                username:'',
                password:'',
                chcekpass:'',
                role:2,
                email:'',
            },
            userRules:{
                username:[{validator:(rule,value,callback)=>{
                    if(this.userInfo.username==""){
                        callback(new Error("请输入用户名"))
                    }
                    if([...this.userInfo.username].length<4||[...this.userInfo.password].length>12){
                        callback(new Error("用户名长度应该在4-12位之间"))
                    }else{
                        callback()
                    }
                },
                trigger:'blur'}],

                password:[{validator:(rule,value,callback)=>{
                    if(this.userInfo.password==""){
                        callback(new Error("请输入密码"))
                    }
                    if([...this.userInfo.password].length<6||[...this.userInfo.password].length>20){
                        callback(new Error("密码长度应该在6-20位之间"))
                    }else{
                        callback()
                    }
                },
                
                trigger:'blur'}],
                chcekpass:[{validator:(rule,value,callback)=>{
                    if(this.userInfo.password==""){
                        callback(new Error("请输入密码"))
                    }
                    if(this.userInfo.password!=this.userInfo.chcekpass){
                        callback(new Error("密码不一致，请重新输入"))
                    }else{
                        callback()
                    }
                   },
                 trigger:'blur'}],
                email:[{
                    pattern:'^[a-zA-Z0-9_.-]+@[a-zA-Z0-9-]+(\.[a-zA-Z0-9-]+)*\.[a-zA-Z0-9]{2,6}$',
                    message:"请输入正确的邮箱",
                    trigger:"blur",
                }]
            }
        }
    },
    created(){
            this.getUserList()
    },
    methods: {
        //用户列表
      async getUserList(){
             this.querydata.pagesize=this.paginationOption.defaultPageSize
             this.querydata.pagenum=this.paginationOption.defaultCurrent
             const{data:result} = await this.$http.post('users',this.querydata)
             if (result.status!=200) return this.$message.error(result.message)
             this.userlist=result.data
             this.paginationOption.total=result.total
      },
      async findUser(){
             this.querydata.pagesize=this.paginationOption.defaultPageSize
             this.querydata.pagenum=1
             const{data:result} = await this.$http.post('users',this.querydata)
             if (result.status!=200) return this.$message.error(result.message)
             this.userlist=result.data
             this.paginationOption.total=result.total
      },
      //删除用户
      deleteUser(id){
          this.$confirm({
              title:"确定要删除吗？",
              content:"一旦删除，无法恢复",
              onOk:async()=>{     
                    const {data:result}=await this.$http.delete(`user/${id}`)
                   if(result.status!=200)return this.$message.error(result.message)
                    this.$message.success('删除成功')
                    this.getUserList()
              },
              onCancel:()=>{
                  this.$message.info("已取消删除")
              }
          })
      },
      //新增用户
      addUserOk(){
          this.$refs.addUserRef.validate(async (valid)=>{
              if(!valid)return this.$message.error("输入不符合要求，请重新输入")
              const {data:result}= await this.$http.post('user/add',this.userInfo)
              console.log(result)
              if(result.status!=200)return this.$message.error(result.message)
              this.$message.success("添加成功")
              this.addUserVisible=false
              this.getUserList()
          })
      },
      addUserCancel(){
          this.$refs.addUserRef.resetFields()     
          this.addUserVisible=false,
          this.$message.info("新增用户已取消")
      },
      adminChange(value){
          this.userInfo.role=Number(value)
      },
      //编辑用户
       async editUser(id){
          const {data:result} = await this.$http.get(`user/${id}`)
          this.userInfo.username=result.data.username
          this.userInfo.role=Number(result.data.role) 
          this.userInfo.email=result.data.email
          this.userInfo.id=id
          this.editUserVisible=true

      },
      editUserOk(){
            this.$refs.editUserRef.validate(async (valid)=>{
              if(!valid)return this.$message.error("输入不符合要求，请重新输入")
              console.log(this.userInfo.role)
              const {data:result}= await this.$http.post("user/update",this.userInfo)
              if(result.status!=200)return this.$message.error(result.message)
              this.$message.success("用户信息更新成功")
              this.editUserVisible=false
              this.userInfo.email=''
              this.userInfo.username=''
              this.getUserList()
          })
      },
      editUserCancel(){
          this.editUserVisible=false
          this.userInfo.email=''
          this.userInfo.username=''
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