// 帖子接口
export interface IPostData{

    likenum: number  
    title: string 
    authorname: string  
   
}
export interface ICreatePost{
    author_id:string
    title:string
    content:string
}

export interface IPostDetail{
    like_num:number
    postid:string
    title:string
    content:string

}

export class PostData{

    list: IPostData[]=[]    
}

export class CreatePostData{
    ruleForm: ICreatePost = {
        author_id:"",
        title: "",
        content:""
    }
}

export class PostDetailData{
    postdetail: IPostDetail={
        like_num:0,
        postid:"",
        title:"",
        content:"",
      
    }
}