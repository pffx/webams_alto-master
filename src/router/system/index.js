import ManagementSystemPage from '../../pages/system/management'
import SoftManagementSystemPage from '../../pages/system/softManagement'
import HardwareSystemPage from '../../pages/system/hardware'


import NotFoundPage from '../../pages/404';
import { useRoutes } from 'react-router-dom';

const SystemRoutes = () => {
  const routes = useRoutes([
    {
      path: '',
      element: <ManagementSystemPage/>,
    },
    {
      path: '/management',
      element: <ManagementSystemPage/>,
    },
    {
      path: '/s_management',
      element: <SoftManagementSystemPage/>
    },
    
    {
      path: '/hardware',
      element: <HardwareSystemPage/>
    },
    {
      path: '/*',
      element: <NotFoundPage/>
    },
  ])
  return routes;
}

export default SystemRoutes
