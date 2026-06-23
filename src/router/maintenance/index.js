import ManagementMaintenPage from '../../pages/maintenance/management'
import SoftManagementMaintenPage from '../../pages/maintenance/softManagement'
import ZeroTouchMaintenPage from '../../pages/maintenance/zeroTouch'
import BackupRestoreMainPage from '../../pages/maintenance/backupRestore'
import LoadCofnigMaintenPage from '../../pages/maintenance/loadConfig'
import FactoryResetMaintenPage from '../../pages/maintenance/factoryReset' 
import DeployMaintenPage from '../../pages/maintenance/deploy' 
import CommandMaintenPage from '../../pages/maintenance/command' 

import NotFoundPage from '../../pages/404';
import { useRoutes } from 'react-router-dom';

const MaintenanceRoutes = () => {
  const routes = useRoutes([
    {
      path: '',
      element: <BackupRestoreMainPage/>,
    },
    {
      path: '/management',
      element: <ManagementMaintenPage/>,
    },
    {
      path: '/s_management',
      element: <SoftManagementMaintenPage/>
    },
    {
      path: '/zero_touch',
      element: <ZeroTouchMaintenPage/>
    },
    {
      path: '/backup_restore',
      element: <BackupRestoreMainPage/>
    },
    {
      path: '/load_config',
      element: <LoadCofnigMaintenPage/>
    },
    {
      path: '/reset',
      element: <FactoryResetMaintenPage/>
    },
    {
      path: '/deploy',
      element: <DeployMaintenPage/>
    },
    {
      path: '/command',
      element: <CommandMaintenPage/>
    },
    {
      path: '/*',
      element: <NotFoundPage/>
    },
  ])
  return routes;
}

export default MaintenanceRoutes
