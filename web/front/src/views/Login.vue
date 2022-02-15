<template>
        <v-app>
            <div class="loginBox">
                <v-form
                    class="loginForm"
                    ref="loginFormRef"
                    :model="formdata"
                    lazy-validation
                >
                    <v-text-field
                    v-model="formdata.username"
                    :counter="10"
                    :rules="nameRules"
                    label="用户名"
                    required
                    ></v-text-field>

                    <v-text-field
                    v-model="formdata.password"
                    :rules="passwordRules"
                    label="密码"
                    required
                    ></v-text-field>

            

                    <v-btn
                    color="success"
                    class="mr-4"
                    type="submit"
                     @click.prevent="login"
                    >
                    登录
                    </v-btn>

                    <v-btn
                    color="error"
                    class="mr-4"
                    @click="resetForm"
                    >
                    取消
                    </v-btn>
                </v-form>
            </div>
        </v-app>
</template>

<script>
export default {
    data(){
        return{
            formdata:{
                username:'',
                password:''
            },
            valid: true,
            nameRules: [
                v => !!v || '用户名不能为空',
                v => (v&&v.length>=4&& v.length <= 12) || '用户名长度必须在4-12之间',
            ],
            passwordRules: [
                  v => !!v || '密码不能为空',
                v => (v&&v.length >=6&& v.length <= 20) || '密码长度必须在6-20之间',
            ],
        };
    },
    methods: {
        resetForm(){
            this.$refs.loginFormRef.reset()
        },
        async login(){
            const valid =this.$refs.loginFormRef.validate()
            // alert("111")
            if(!valid) return //this.$message.error("输入非法数据，请重新输入")
            else{
                const {data:result} = await this.$http.post('login',this.formdata)
                if (result.code!=200)return 
                window.sessionStorage.setItem('token',result.token)
                // console.log(result.data)
                this.$router.push('/home')
            }
        }
    },
}
</script>

<style scoped>
    .container{
        height: 100%;
        background-color: #282c34;
        width: 100%;
    }
    .loginBox{
        width: 400px ;
        height: 300px;
        background-color: #fff;
        position: absolute;
        top: 50%;
        left: 50%;
        transform: translate(-50%,-50%);
        border: 1px solid #000;
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
