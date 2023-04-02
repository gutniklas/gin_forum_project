import { createStore } from "vuex";

const store=createStore({
    state(){
        return{
            userid:'',
            postdetail:{
                like_num:0,
                postid:'',
                title:'',
                content:'',
            }

        }

    },
    mutations:{
        setuserinfo(state:any,userid){
            state.userid=userid;
        },
        setpostinfo(state:any,postdetail){
            state.postdetail=postdetail;

        },
    },
    actions:{
        setuserinfoaction:function(context){
            context.commit('setuserinfo',sessionStorage.getItem('userid'))
        },
        setpostinfoaction:function(context,payload){
            context.commit('setpostinfo',payload)
        }
    }

})

export default store