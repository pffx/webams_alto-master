import React, { useState } from 'react';
import { useTranslation } from 'react-i18next';
import { connect } from 'react-redux';
import { ToastContainer, toast } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';
import AXIOS from '../../../axios';
import GLOBAL, { TOAST_CONF } from '../../../global';
import { API_Login } from '../../../global/API';
import { loginAction, logoutAction } from '../../../actions/login';
import Redirect from '../../../router/redirect';
import LanguageSelector from '../../../components/language';
import '../../../css/mobile.css';

function MobileLoginPage({ isLogin, loginClick, logoutClick }) {
  const { t } = useTranslation();
  const [username, setUsername] = useState('admin');
  const [password, setPassword] = useState('');
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState('');

  const handleLogin = () => {
    if (!username.trim()) {
      setError(t('mobile.login_username_required'));
      return;
    }

    setLoading(true);
    setError('');

    AXIOS
      .post(API_Login, {
        account: username,
        password,
      })
      .then((res) => {
        setLoading(false);
        if (res.data.status === 200) {
          loginClick({
            account: username,
            userInfo: res.data.userInfo,
            token: res.data.token,
          });
          toast.success(res.data.message || t('mobile.login_success'), TOAST_CONF);
        } else {
          logoutClick();
          setError(res.data.message || t('mobile.login_failed'));
        }
      })
      .catch(() => {
        setLoading(false);
        logoutClick();
        setError(t('message.server_error'));
      });
  };

  if (isLogin) {
    return <Redirect to="/m/upgrade" />;
  }

  return (
    <div className="mobile-app mobile-login">
      <div className="mobile-login__card">
        <h1 className="mobile-login__title">{t('mobile.login_title')}</h1>
        <p className="mobile-login__subtitle">{t('mobile.login_subtitle')}</p>

        {error && (
          <div className="mobile-feedback mobile-feedback--error" role="alert">
            {error}
          </div>
        )}

        <label className="mobile-field">
          <span className="mobile-field__label">{t('mobile.username')}</span>
          <input
            className="mobile-field__input"
            type="text"
            value={username}
            onChange={(e) => setUsername(e.target.value)}
            autoComplete="username"
            placeholder={t('mobile.username')}
          />
        </label>

        <label className="mobile-field">
          <span className="mobile-field__label">{t('mobile.password')}</span>
          <input
            className="mobile-field__input"
            type="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            autoComplete="current-password"
            placeholder={t('mobile.password')}
          />
        </label>

        <button
          type="button"
          className="mobile-btn mobile-btn--primary mobile-btn--block"
          onClick={handleLogin}
          disabled={loading}
        >
          {loading ? t('mobile.logging_in') : t('button.login')}
        </button>

        <div className="mobile-login__footer">
          <LanguageSelector />
          <a className="mobile-link" href="/login">{t('mobile.pc_entry')}</a>
        </div>
      </div>
      <ToastContainer position="top-center" autoClose={2000} />
    </div>
  );
}

const stateToProps = (state) => ({
  isLogin: state.LoginReducer.isLogin,
});

const dispatchToProps = (dispatch) => ({
  loginClick(data) {
    dispatch(loginAction(data));
  },
  logoutClick() {
    dispatch(logoutAction());
  },
});

export default connect(stateToProps, dispatchToProps)(MobileLoginPage);
