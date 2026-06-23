import React from 'react';
import { useTranslation } from 'react-i18next';
import { useDispatch, useSelector } from 'react-redux';
import { ToastContainer } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';
import { logoutAction } from '../../actions/login';
import { clearAllTabIndex } from '../../actions/global';
import Redirect from '../../router/redirect';
import '../../css/mobile.css';

function MobileLayout({ children, title }) {
  const { t } = useTranslation();
  const dispatch = useDispatch();
  const { isLogin } = useSelector((state) => state.LoginReducer);

  const handleLogout = () => {
    dispatch(logoutAction());
    dispatch(clearAllTabIndex());
  };

  if (!isLogin) {
    return <Redirect to="/m/login" />;
  }

  return (
    <div className="mobile-app">
      <header className="mobile-header">
        <h1 className="mobile-header__title">{title || t('mobile.upgrade_title')}</h1>
        <button type="button" className="mobile-header__logout" onClick={handleLogout}>
          {t('button.logout')}
        </button>
      </header>
      <main className="mobile-main">
        {children}
      </main>
      <ToastContainer position="top-center" autoClose={2000} />
    </div>
  );
}

export default MobileLayout;
