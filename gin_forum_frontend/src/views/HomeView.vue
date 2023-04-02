<template>
  <div class="home">
    <el-container>
    <el-header><el-row>
    <el-col :span="4" class="post"><el-button  @click="creatpost()">发帖</el-button></el-col>
    <el-col :span="16"><h2>论坛</h2></el-col>
    <el-col :span="4" class="log-out" ><el-button @click="logout()">退出登录</el-button></el-col>

    </el-row></el-header>
      
        <el-main>  
        <PostListView></PostListView>
        <router-view> </router-view>
        </el-main>
      
    </el-container>
  </div>
</template>

<script lang="ts">
import { defineComponent,onMounted } from 'vue';
import {useRouter} from 'vue-router';
import {useStore} from 'vuex';
import PostListView from "./PostListView.vue"
export default defineComponent({
  name: 'HomeView',
  setup(){
      const router=useRouter()
      // const list=router.getRoutes().filter(v=>v.meta.isShow)   
      const store=useStore()
      const creatpost=()=> {
        router.push('./createpost')
      }
      const logout=()=>{
      localStorage.removeItem("Authorization")
      sessionStorage.removeItem("username")
      sessionStorage.removeItem("userid")
      router.push('/login')

      }

      return {logout,creatpost}
  },
  components: {
   PostListView
  },
});
</script>
<style lang="scss" scoped>


.el-header{
height:80px;
background-color: #F4D03F;
background-image: linear-gradient(121deg, #F4D03F 0%, #16A085 100%);


.el-button{
  color:#ff8000;
  width:80px;
  background-color:#fff2cc;
}
h2,h1{
  color:#0040ff;
  text-align: center;
  height: 80px;
  line-height:80px;
  font-weight:bolder

}

}
.log-out{
  padding-left:100px;
  height: 80px;
  line-height:80px;

}


.post{
  
  color:blue;
  height: 80px;
  line-height:80px;
}

.el-main {
  background-color: #ccf2ff;

  height: calc(100vh - 80px);

}

</style>