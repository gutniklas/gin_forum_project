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
  <h2>新用户注册</h2>
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
    <el-form-item label="重复密码:" prop="password">
    <el-input
        v-model="ruleForm.re_password"
        type="password"
        autocomplete="off"
      />
    </el-form-item>
    <el-form-item>
      <el-button class="signupBtn" type="primary" @click="submitForm(ruleFormRef)"
        >注册</el-button
      >
      <el-button class="signupBtn" @click="resetForm(ruleFormRef)">重置</el-button>
     
    </el-form-item>

    
  </el-form>
    </div>

</template>

<script lang="ts">
import type { FormInstance } from 'element-plus';
import {defineComponent, reactive, ref, toRefs} from "vue";
import {SignupData} from '../type/signup';
import {signup} from '../request/api';
import {useRouter} from 'vue-router';
import {ElMessage, ElMessageBox} from "element-plus";
export default defineComponent({
    setup () {
        const data=reactive( new SignupData());

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
        ],
        re_password:[
            {
                required: true,
                message: "请再次输入密码", 
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
          signup(data.ruleForm).then((res) => { 
          ElMessage({
              type: 'success',
              message: '注册成功，返回登陆页面',
          })
          // 跳转页面
          router.push('/login')
          })
        } else {
          console.log('error submit!')
          return false
        }
      })
    };
          // 重置
    const resetForm = () => {
      data.ruleForm.username = "";
      data.ruleForm.password = "";
    }


        return {...toRefs(data),rules,ruleFormRef,submitForm,resetForm};
    }
})
</script>

<style lang='scss' scoped>


.login-box{
    width:100%;
    height: 100vh;
    background-color: #A9C9FF;
    background-image: linear-gradient(135deg, #A9C9FF 0%, #FFBBEC 100%);

    padding: 1px;
    text-align:center;
    background-repeat:no-repeat;

}
.demo-ruleForm{
    width:500px;
    margin:200px auto;
    background-color:#fff;
    padding: 40px;
    border-radius:20px;
}
.signupBtn{
    width:48%;
}

h2{ margin-left:30px;
    margin-bottom:15px;
    color:black;
}
</style>