import service from ".";
import {ILoginData} from "../type/login";
import {ISignupData} from "../type/signup";
import {ICreatePost} from "../type/post";
import {IVoteData} from "../type/vote"
export function login(data:ILoginData){
    return service({
        url:"/login",
        method:"post",
        data
    })
}
//注册接口
export function signup(data:ISignupData){
    return service({
        url:"/signup",
        method:"post",
        data
    })
}

//帖子列表接口
export function getPostsList(page:number,size:number){
    return service({
        url:"/posts?page="+String(page)+"&size="+String(size),
        method:"get",
        
    })
}

//发表帖子接口
export function createPost(data:ICreatePost){
    return service({
        url:"/createpost",
        method:"post",
        data
    })
}

//查看帖子内容接口

export function getPostDetail(url :string){
    return service({
        url:url,
        method:"get",
        
    })
}

//点赞接口
export function vote(data:IVoteData){
    return service({
        url:"/vote",
        method:"post",
        data
        
    })
}
