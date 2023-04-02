<template>
    <div >
        <el-header><el-row >
        <el-col :span="1"  class="gotohome"><el-button type="primary" @click="gotohome()">返回主页</el-button></el-col>
        <el-col :span="23">
            <h4>
            {{postdetail.title}}
            </h4>
        </el-col>
        
        </el-row></el-header>
    
        <div id="content">
        {{postdetail.content}}
        </div>
        <div class="icon">
        <el-button  @click="like()">
        <el-icon :size="30"><Pointer class="pointer" /></el-icon>
        点赞
        </el-button>
        </div>

    </div>
</template>

<script lang="ts">
import {defineComponent, reactive,  toRefs,onMounted} from "vue";
import {PostDetailData} from "../type/post";
import {VoteData} from "../type/vote"
import {useStore} from 'vuex';
import {vote} from '../request/api';
import {useRouter} from 'vue-router';
export default defineComponent({
    setup () {
        const data=reactive(new PostDetailData())
        const votedata=reactive(new VoteData())
        const store=useStore()
        const router=useRouter()
        const gotohome=()=>{
            router.push('/')

        }
        const like=()=>{
            votedata.ruleForm.postid=store.state.postdetail.postid
            vote(votedata.ruleForm).then((res)=>{
                console.log(res.data)
            })
            
        }
        onMounted(()=>{            
    
            data.postdetail.title=store.state.postdetail.title
            data.postdetail.content=store.state.postdetail.content
                
        })
        return {...toRefs(data),gotohome,like}
    }
})
</script>

<style scoped>
.el-row{
    
}
.el-header{
    text-align:center;
    
    height:60px;
    line-height:60px;
    background-color: #e5ffe5;

      
}
#content{
        text-align:left;
        padding-top:30px;
        padding-left:10px;

        height:calc(vh - 100px)
}
.pointer{
color:orange;
}
.gotohome{

    text-align:left;
}
.icon{
     
     text-align:right; 
     padding-right:20px;
}
</style>