<template>
    <div class="container">
        <div class="loginBox">
            <a-form-model
            ref="loginFormRef"  
            class="loginForm" 
            :rules="rules" 
            :model="formdata" 
            :label-col="labelCol"
            :wrapper-col="wrapperCol"
             >
                <a-form-model-item prop="username" label="用户名">
                     <a-input 
                        placeholder="请输入用户名" 
                        v-model="formdata.username"
                     >
                        <a-icon slot="prefix" type="user" style="color:rgba(0,0,0,.25)" />
                     </a-input>
                </a-form-model-item>
                 <a-form-model-item prop="password" :label="label.password">
                     <a-input 
                        v-model="formdata.password"
                        placeholder="请输入密码"
                        type="password"
                        v-on:keyup.enter="login"
                     >
                       <a-icon slot="prefix" type="lock" style="color:rgba(0,0,0,.25)" />
                     </a-input>
                </a-form-model-item>
                <a-form-model-item class="loginBtn">
                        <a-button type="primary" style="margin:10px" @click="login">登录 </a-button>
                        <a-button type="info"  style="margin-left:60px" @click="resetForm"> 取消</a-button>
                </a-form-model-item>
            </a-form-model>
        </div>
    </div>
</template>

<script>
export default {
    data(){
        return{
            labelCol: { span: 4 },
            wrapperCol: { span: 19 },
            label:{
                password:"密"+"\xa0\xa0\xa0\xa0"+"码"
            },
            formdata:{
                username:'',
                password:''
            },
            rules: {
                username: [
                    { 
                        required: true, 
                        message: '请输入用户名',
                        trigger: 'blur' 
                    },
                    {
                        min: 3,
                        max: 12, 
                        message: '用户名必须在3-12个字符之间', 
                        trigger: 'blur' 
                    },
                ],
                password: [
                    { 
                        required: true, 
                        message: '请输入密码', 
                        trigger: 'blur' 
                    },
                    { 
                        min: 6, max: 20, 
                        message: '密码必须在6-20个字符之间', 
                        trigger: 'blur' 
                    },
                ],
            }
        };
    },
    methods: {
        resetForm(){
            this.$refs.loginFormRef.resetFields()
            console.log(this)
        },
        login(){
            this.$refs.loginFormRef.validate(async valid=>{
                if(!valid) return this.$message.error("输入非法数据，请重新输入")
                else{
                    const {data:result} = await this.$http.post('login',this.formdata)
                    console.log(result)
                    if (result.code!=200)return this.$message.error(result.message)
                    window.sessionStorage.setItem('token',result.token)
                    this.$router.push('/index')
                }
            })
        }
    },
}
</script>

<style scoped>
    .container{
        height: 100%;
        background-color: #282c34;

    }
    .loginBox{
        width: 400px ;
        height: 300px;
        background-color: #fff;
        position: absolute;
        top: 50%;
        left: 50%;
        transform: translate(-50%,-50%);
        border-radius: 9px;
    }
    .loginForm{
        width: 100%;
        position: absolute;
        bottom: 10%;
        padding: 0 20px;
        box-sizing: border-box;
    }
    .loginBtn{
        display: flex;
        justify-content: flex-end;
    }
</style>