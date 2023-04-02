<template>

    <div class='login-box'>
    <el-form
    ref="ruleFormRef"
    :model="ruleForm"
    status-icon
    :rules="rules"
    label-width="60px"
    class="demo-ruleForm"
  >
  <h2>用户登录</h2>
    <el-form-item label="账号:" prop="username">
      <el-input v-model="ruleForm.username"  autocomplete="off" />
    </el-form-item>

    <el-form-item label="密码:" prop="password">
      <el-input
        v-model="ruleForm.password"
        type="password"
        autocomplete="off"
      />
    </el-form-item>
    <el-form-item>
      <el-button class="loginBtn" type="primary" @click="submitForm(ruleFormRef)"
        >登录</el-button
      >
      
     
    </el-form-item>
    <el-form-item>

      <el-button class="signupBtn" @click="gotosignup()">注册</el-button>
     
    </el-form-item>
    
  </el-form>
    </div>

</template>

<script lang="ts">
import type { FormInstance } from 'element-plus';
import {defineComponent, reactive, ref, toRefs} from "vue";
import {LoginData} from '../type/login';
import {login} from '../request/api';
import {useRouter} from 'vue-router';
export default defineComponent({
    setup () {
        const data=reactive( new LoginData());
        const rules={
        username: [
        {
          required: true,  //是否必须字段
          message: "请输入用户名",   // 触发的提示信息
          trigger: "blur"   // 触发时机: 当失去焦点时（光标不显示的时候），触发此提示
        },
        {
          min: 3,   // 最小字符书
      // 最大字符数
          message: "用户名长度至少需要3个字符",  // 触发的提示信息
          trigger: "blur"
        }
        ],
        password: [
        {
          required: true,  //是否必须字段
          message: "请输入密码",   // 触发的提示信息
          trigger: "blur"   // 触发时机: 当失去焦点时（光标不显示的时候），触发此提示
        },
        {
          min: 6,   // 最小字符书

          message: "密码长度至少需要6个字符",  // 触发的提示信息
          trigger: "blur"
        }
        ]
        }
  const ruleFormRef = ref<FormInstance>();
  const router=useRouter();
  const submitForm = (formEl: FormInstance | undefined) => {
      // 对表单内容进行验证
      if (!formEl) return
      formEl.validate((valid) => {
        if (valid) {
          login(data.ruleForm).then((res) => {
     
            // 将token进行保存
            console.log(res.data)
            localStorage.setItem("Authorization", "Bearer "+res.data.token)
            sessionStorage.setItem("userid",res.data.user_id)
            sessionStorage.setItem("username",res.data.user_name)
            // 跳转页面
            router.push('/')
          })
        } else {
          console.log('error submit!')
          return false
        }
      })
    };


    const gotosignup=()=>{
       router.push('/signup')
    }
        return {...toRefs(data),rules,ruleFormRef,submitForm, gotosignup};
    }
})
</script>

<style lang='scss' scoped>


.login-box{

    width: 100vw;
    height: 100vh;
    background-color: #0093E9;
    background-image: linear-gradient(90deg, #0093E9 0%, #80D0C7 100%);

    padding: 1px;
    
 

}
.demo-ruleForm{
    width:500px;
    margin:200px auto;
    background-color:#fff;
    padding: 20px;
    border-radius:20px;
}
.loginBtn{
    width:98%;
}
.signupBtn{
    width:98%;
    background-color:#fff;
    color:black;
}
h2{
    text-align:center;
    margin-left:20px;
    margin-bottom:15px;
    color:black;
}
</style>