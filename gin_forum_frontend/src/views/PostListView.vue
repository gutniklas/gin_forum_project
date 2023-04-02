<template>

    <div> 
    <el-table :data="list" border style="width: 100%" @cell-click="gotopostdetail">
    <el-table-column prop="title" label="讨论"  />
    <el-table-column prop="author_name" label="作者" width="180" />
    <el-table-column prop="like_num" label="点赞数" width="180" :sort-method="sortByLikeNum">
 

    </el-table-column>
    </el-table>
    <el-pagination @current-change="currentChange" layout="prev, pager, next" :total="30">
  </el-pagination>
    </div>
 
</template>

<script lang="ts">
import { defineComponent,reactive,toRefs, onMounted} from 'vue';
import {getPostsList,getPostDetail} from "../request/api";
import {PostData} from "../type/post";
import {useRouter} from 'vue-router';
import {useStore} from 'vuex'
export default defineComponent({
  setup(){
    const data = reactive(new PostData()) 
    const router=useRouter()
    const store=useStore();
    const gotopostdetail=(row:any,column:any,cell:any) =>{    
      const postid=row.postid
      const url="/topic/"+ postid
      getPostDetail(url).then((res) => {
            // console.log("res",res.data)
            store.dispatch('setpostinfoaction',res.data).then(()=>{
            // console.log("postid",store.state.postdetail.postid)
            // console.log("title",store.state.postdetail.title)   
            router.push({ name: "postdetail" ,params:{
						postid:postid
					} },) })
          })
 
    }
    onMounted(()=>{
      getPostsList(1,10).then(res=>{
        console.log(res.data)
        data.list=res.data
   
      })
    })
    
    const currentChange=(page:number)=>{
       
      const size=10

      getPostsList(page,size).then(res=>{
        console.log(res.data)
        data.list=res.data
   
      })

    }

    return{...toRefs(data),gotopostdetail,router,currentChange};
  }
})
</script>

<style scoped>
.el-table {
    /deep/tbody tr:hover>td { 
        background-color:#90c0f1； 
    }	
}
</style>