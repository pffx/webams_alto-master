import React from 'react';
import { useSelector } from 'react-redux';
import Redirect from '../../router/redirect';

function MobileLoginAuth({ children }) {
  const { isLogin } = useSelector((state) => state.LoginReducer);

  return isLogin
    ? children
    : <Redirect to="/m/login" />;
}

export default MobileLoginAuth;
