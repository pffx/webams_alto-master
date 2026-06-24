import React, { useState } from 'react';
import { useTranslation } from 'react-i18next';
import { toast } from 'react-toastify';
import AXIOS from '../../../../axios';
import GLOBAL, { TOAST_CONF } from '../../../../global';
import { API_SoftwareAction } from '../../../../global/API';
import {
  buildActivePayload,
  getActiveSoftwareName,
} from '../../../../utils/softwareUpgrade';

function ActivateStep({ deviceSelection, downloadResult, onNext, onBack }) {
  const { t } = useTranslation();
  const { olt, card } = deviceSelection;
  const panel = downloadResult.panel;
  const activeName = getActiveSoftwareName(panel);
  const [status, setStatus] = useState(activeName ? 'pending' : 'skipped');
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState('');

  const handleActivate = () => {
    if (!activeName) {
      onNext({ panel, needsRebootWait: false });
      return;
    }

    setLoading(true);
    setError('');
    const payload = buildActivePayload({ ip: olt.ip, port: card.port }, activeName);

    AXIOS
      .put(API_SoftwareAction, payload)
      .then((res) => {
        setLoading(false);
        if (res.data.status === GLOBAL.ERROR_NUM.Success) {
          toast.success(t('mobile.activate_success'), TOAST_CONF);
          setStatus('success');
          onNext({
            panel,
            activatedName: activeName,
            needsRebootWait: true,
          });
        } else {
          setStatus('failed');
          setError(t('mobile.activate_failed'));
        }
      })
      .catch(() => {
        setLoading(false);
        setStatus('failed');
        setError(t('message.server_error'));
      });
  };

  const handleSkip = () => {
    onNext({ panel, needsRebootWait: false });
  };

  return (
    <div className="mobile-step">
      <h2 className="mobile-step__title">{t('mobile.step_activate')}</h2>

      {activeName ? (
        <div className="mobile-info-card">
          <p>{t('mobile.activate_confirm', { name: activeName })}</p>
          {error && <div className="mobile-feedback mobile-feedback--error">{error}</div>}
        </div>
      ) : (
        <div className="mobile-info-card">
          <p>{t('mobile.no_activate_needed')}</p>
        </div>
      )}

      <div className="mobile-actions">
        <button type="button" className="mobile-btn mobile-btn--secondary" onClick={onBack}>
          {t('mobile.back')}
        </button>
        {activeName ? (
          <button
            type="button"
            className="mobile-btn mobile-btn--primary"
            disabled={loading || status === 'success'}
            onClick={handleActivate}
          >
            {loading ? t('mobile.processing') : t('button.active')}
          </button>
        ) : (
          <button type="button" className="mobile-btn mobile-btn--primary" onClick={handleSkip}>
            {t('mobile.skip')}
          </button>
        )}
      </div>
    </div>
  );
}

export default ActivateStep;
