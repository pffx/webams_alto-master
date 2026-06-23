import HomePage from '../pages/home'
import LoginPage from '../pages/login'
import MaintenancePage from '../pages/maintenance/'
import AlarmLogPage from '../pages/alarm_log'
import SystemPage from '../pages/system'
import OntsPage from '../pages/onts'
import PortssPage from '../pages/ports'
import ServicePage from '../pages/service'
import NotFoundPage from '../pages/404';
import { Outlet, useRoutes } from 'react-router-dom';
import AppHome from '../AppHome';
import LoginAuth from '../components/loginAuth';
import MaintenanceRoutes from '../router/maintenance'
import AlarmLogRoutes from '../router/alarm_log'
import ServiceRoutes from "../router/service"
import SystemRoutes from "../router/service"
import MobileRoutes from '../router/mobile'

const RootRoutesConfig = () => {
  const routes = useRoutes([
  {
    path:'/m/*',
    element: <MobileRoutes/>
  },
  {
    path:'/',
    element: <AppHome/>,
    children:[]
  },
  {
    path:'/home',
    element: <LoginAuth><HomePage/></LoginAuth>
  },
  {
    path:'/login',
    element: <LoginPage/>
  },
  {
    path:'/ports',
    element: <LoginAuth><PortssPage/></LoginAuth>
  },
  {
    path:'/service/*',
    element: <LoginAuth><ServicePage/></LoginAuth>,
    children: [
      {
        path: '',
        element: <LoginAuth><ServiceRoutes/></LoginAuth>,
      },
    ]
  },
  {
    path:'/onts',
    element: <LoginAuth><OntsPage/></LoginAuth>
  },
  {
    path:'/system/*',
    element: <LoginAuth><SystemPage/></LoginAuth>,
    children: [
      {
        path: '',
        element: <LoginAuth><SystemRoutes/></LoginAuth>,
      },
    ]
  },
  {
    path:'/maintenance/*',
    element: <LoginAuth><MaintenancePage/></LoginAuth>,
    children: [
      {
        path: '',
        element: <LoginAuth><MaintenanceRoutes/></LoginAuth>,
      },
    ]
  },
  {
    path:'/alarm_log/*',
    element: <LoginAuth><AlarmLogPage/></LoginAuth>,
    children: [
      {
        path: '',
        element: <LoginAuth><AlarmLogRoutes/></LoginAuth>,
      },
    ]
  },
  {
    path:'*',
    element: <NotFoundPage/>
  },
  
  ]);
  return routes;
}

export default RootRoutesConfig
