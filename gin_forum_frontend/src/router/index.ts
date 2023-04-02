import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router'
import HomeView from '../views/HomeView.vue'

const routes: Array<RouteRecordRaw> = [


  {
    path: '/',
    name: 'home',
    component: HomeView,
    beforeEnter:(to,from)=>{
      console.log(to);
      console.log(from);

    },
    children:[
      {
        path:"posts",
        name:"posts",
        meta:{
          isShow:true,
          title:"帖子"
        },
        component: () => import(/* webpackChunkName: "goods" */ '../views/PostListView.vue')
      },
    ]
  },
  {
    path: '/about',
    name: 'about',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/AboutView.vue')
  },
  {
    path: '/login',
    name: 'login',
    component: () => import(/* webpackChunkName: "login" */ '../views/LoginView.vue')
  },
  {
    path: '/signup',
    name: 'signup',
    component: () => import(/* webpackChunkName: "signup" */ '../views/SignUpView.vue')
  },
  {
    path: '/createpost',
    name: 'createpost',
    component: () => import(/* webpackChunkName: "createpos" */ '../views/CreatePostView.vue')
  },
  {
    path:"/topic/:postid",
    name:"postdetail",
    meta:{
      isShow:true,
      title:"帖子详情"
    },
    component: () => import(/* webpackChunkName: "postdetail" */ '../views/PostDetailView.vue')
  },

]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

router.beforeEach((to,from,next)=>{
  const token: string | null = localStorage.getItem("Authorization")
  if(to.path == '/login' || to.path == '/signup'||token){
    next();
  }else{
    next('/login');
  }
})


export default router
