import AlarmPage from '../../pages/alarm_log/alarm'
import LogPage from '../../pages/alarm_log/log'

import NotFoundPage from '../../pages/404';
import { useRoutes } from 'react-router-dom';

const AlarmLogRoutes = () => {
  const routes = useRoutes([
    {
      path: '',
      element: <AlarmPage/>,
    },
    {
      path: '/alarm',
      element: <AlarmPage/>,
    },
    {
      path: '/log',
      element: <LogPage/>
    },
    {
      path: '/*',
      element: <NotFoundPage/>
    },
  ])
  return routes;
}

export default AlarmLogRoutes
