//懒加载
const Home = () => import(/* webpackChunkName: "group-Home" */  '@/components/Home')
const Welcome = () => import(/* webpackChunkName: "group-Welcome" */  '@/components/Welcome')
const Category = () => import(/* webpackChunkName: "group-Category" */  '@/components/Category')
const Article = () => import(/* webpackChunkName: "group-Article" */  '@/components/Article')
const Write = () => import(/* webpackChunkName: "group-Write" */  '@/components/Write')
const Login = () => import(/* webpackChunkName: "group-Login" */  '@/components/Login')

const routers = [
    {
        path: '/login',
        component: Login
    },
    {
        path: '/',
        component: Home,
        children:[
            {
                path: '/',
                component: Welcome
            },
            {
                path: 'article',
                component: Article
            },
        ]
    },
    {
        path: '/category',
        component: Category,
        children:[
            {
                path: '',
                component: Welcome
            },
            {
                path: 'article',
                component: Article
            },
            {
                path: 'write',
                component: Write
            }
        ]
    }
]
export default routers
