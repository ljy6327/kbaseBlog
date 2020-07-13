import Home from './components/Home'
import Welcome from './components/Welcome'
import Category from './components/Category'
import Article from './components/Article'
import Write from './components/Write'
import Login from './components/Login'
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
