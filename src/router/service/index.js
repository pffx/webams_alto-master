import ServiceEditPage from '../../pages/service/edit';
import ServiceMainPage from "../../pages/service/main"

import NotFoundPage from '../../pages/404';
import { useRoutes } from 'react-router-dom';

const ServiceRoutes = () => {
  const routes = useRoutes([
    {
      path: '',
      element: <ServiceMainPage/>,
    },
    {
      path: '/edit',
      element: <ServiceEditPage/>,
    },
    {
      path: '/*',
      element: <NotFoundPage/>
    },
  ])
  return routes;
}

export default ServiceRoutes
