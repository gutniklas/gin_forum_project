import axios from 'axios'
//创建axios实例
const service=axios.create({
    baseURL:"/api/",
    timeout:5000,
    headers:{
        "Content-Type":"application/json;chartset=utf-8"
    }
})
//请求拦截
service.interceptors.request.use((config)=>{
   config.headers=config.headers||{}
   if(localStorage.getItem("Authorization")){
      
       config.headers.Authorization=localStorage.getItem("Authorization")
   }
   return config
})
//响应 拦截
service.interceptors.response.use((res)=>{
        const code:number=res.data.code 
        if (code!=1000){
            return Promise.reject(res.data)
        }
        return res.data
},(err)=>{
    console.log(err)
})


export default service