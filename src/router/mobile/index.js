import React from 'react';
import { useRoutes, Navigate } from 'react-router-dom';
import MobileLoginAuth from '../../components/mobile/MobileLoginAuth';
import MobileLayout from '../../components/mobile/MobileLayout';
import MobileLoginPage from '../../pages/mobile/Login';
import MobileUpgradePage from '../../pages/mobile/Upgrade';

const MobileRoutes = () => {
  return useRoutes([
    {
      path: 'login',
      element: <MobileLoginPage />,
    },
    {
      path: 'upgrade',
      element: (
        <MobileLoginAuth>
          <MobileLayout>
            <MobileUpgradePage />
          </MobileLayout>
        </MobileLoginAuth>
      ),
    },
    {
      path: '',
      element: <Navigate to="/m/upgrade" replace />,
    },
    {
      path: '*',
      element: <Navigate to="/m/upgrade" replace />,
    },
  ]);
};

export default MobileRoutes;
