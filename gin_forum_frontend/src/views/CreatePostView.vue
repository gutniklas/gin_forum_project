<template>
  <div class="home" >

  <el-form
    ref="ruleFormRef"
    :model="ruleForm"
    status-icon
    :rules="rules"
    label-width="60px"
    class="demo-ruleForm"
  
  >
    <el-form-item>
        <el-input class="el-input-header" v-model="ruleForm.title" placeholder="添加标题"></el-input>
    </el-form-item>
    <el-form-item>
        <el-input
        type="textarea"
        v-model="ruleForm.content"
        placeholder="请输入内容"
        rows=15
        class="el-input-context"
        >
        </el-input>
    </el-form-item>
    <el-form-item>
      <el-button class="createpost" type="primary" @click="submitForm(ruleFormRef)">发表</el-button>    
    </el-form-item>
    </el-form>
  </div>
</template>

<script lang="ts">
import type { FormInstance } from 'element-plus';
import {defineComponent, reactive, ref, toRefs} from "vue";
import {CreatePostData} from '../type/post';
import {createPost} from '../request/api';
import {useRouter} from 'vue-router';
import {useStore} from 'vuex'
import {ElMessage, ElMessageBox} from "element-plus";

export default defineComponent({
    setup () {
        const ruleFormRef = ref<FormInstance>();
        const router=useRouter();
        const store=useStore();
        const data=reactive( new CreatePostData());
        const rules={
        title: [
        {
          required: true,  //是否必须字段
          message: "请输入标题",   // 触发的提示信息
          trigger: "blur"   // 触发时机: 当失去焦点时（光标不显示的时候），触发此提示
        },
        {
          min: 1,       
          message: "请输入标题",  // 触发的提示信息
          trigger: "blur"
        }
        ],
        content: [
        {
          required: true,  //是否必须字段
          message: "请输入内容",   // 触发的提示信息
          trigger: "blur"   // 触发时机: 当失去焦点时（光标不显示的时候），触发此提示
        },
        {
          min: 1,       
          message: "请输入内容",  // 触发的提示信息
          trigger: "blur"
        }
        
        ]
        }
        const submitForm = (formEl: FormInstance | undefined) => {
      // 对表单内容进行验证
            if (!formEl) return
                formEl.validate((valid) => {
                if (valid) {                
                    store.dispatch('setuserinfoaction')
                    data.ruleForm.author_id=store.state.userid

                    createPost(data.ruleForm).then((res) => {
                    console.log(res)
                    // 跳转页面
                    router.push('/')
                
          })
        } else {
          console.log('error submit!')
          return false
        }
      })
    };

        return {...toRefs(data),rules,ruleFormRef,submitForm}
    }
})
</script>

<style scoped>
.home{
  display: flex;
  justify-content: center;
  background-color: #e5ffff;
  height: 100vh;
}

.el-input-header{
    width:600px;
    margin-top: 20px;
}
.el-input-context  {

  width:600px;
}

.demo-ruleForm{
    width:600px;
    background-color:#e5ffff;   
    border-radius:20px;
}
</style>