import React, { useEffect, useState } from 'react';
import { useTranslation } from 'react-i18next';
import { toast } from 'react-toastify';
import AXIOS from '../../../../axios';
import GLOBAL, { TOAST_CONF } from '../../../../global';
import { API_SoftwareAction } from '../../../../global/API';
import {
  buildCommitPayload,
  getPendingCommitName,
  fetchOltSoftwareInfo,
} from '../../../../utils/softwareUpgrade';

function CommitStep({ deviceSelection, activateResult, downloadedName, onNext, onBack }) {
  const { t } = useTranslation();
  const { olt, card } = deviceSelection;
  const [panel, setPanel] = useState(activateResult.panel);
  const [refreshing, setRefreshing] = useState(true);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState('');

  useEffect(() => {
    let cancelled = false;
    fetchOltSoftwareInfo(olt.ip, card.port, olt.type)
      .then((freshPanel) => {
        if (!cancelled) {
          setPanel(freshPanel);
          setRefreshing(false);
        }
      })
      .catch(() => {
        if (!cancelled) {
          setRefreshing(false);
        }
      });
    return () => {
      cancelled = true;
    };
  }, [olt, card]);

  const commitName = refreshing ? '' : getPendingCommitName(panel, downloadedName);

  const handleCommit = () => {
    if (!commitName) {
      onNext({ panel });
      return;
    }

    setLoading(true);
    setError('');
    const payload = buildCommitPayload({ ip: olt.ip, port: card.port }, commitName);

    AXIOS
      .put(API_SoftwareAction, payload)
      .then((res) => {
        setLoading(false);
        if (res.data.status === GLOBAL.ERROR_NUM.Success) {
          toast.success(t('mobile.commit_success'), TOAST_CONF);
          fetchOltSoftwareInfo(olt.ip, card.port, olt.type)
            .then((updatedPanel) => onNext({ panel: updatedPanel }))
            .catch(() => onNext({ panel }));
        } else {
          setError(t('mobile.commit_failed'));
        }
      })
      .catch(() => {
        setLoading(false);
        setError(t('message.server_error'));
      });
  };

  const handleContinue = () => {
    onNext({ panel });
  };

  if (refreshing) {
    return <div className="mobile-step-loading">{t('mobile.loading')}</div>;
  }

  return (
    <div className="mobile-step">
      <h2 className="mobile-step__title">{t('mobile.step_commit')}</h2>

      {commitName ? (
        <div className="mobile-info-card">
          <p>{t('mobile.commit_confirm', { name: commitName })}</p>
          {error && <div className="mobile-feedback mobile-feedback--error">{error}</div>}
        </div>
      ) : (
        <div className="mobile-info-card">
          <p>{t('mobile.no_commit_needed')}</p>
        </div>
      )}

      <div className="mobile-actions">
        <button type="button" className="mobile-btn mobile-btn--secondary" onClick={onBack}>
          {t('mobile.back')}
        </button>
        {commitName ? (
          <button
            type="button"
            className="mobile-btn mobile-btn--primary"
            disabled={loading}
            onClick={handleCommit}
          >
            {loading ? t('mobile.processing') : t('button.commit')}
          </button>
        ) : (
          <button type="button" className="mobile-btn mobile-btn--primary" onClick={handleContinue}>
            {t('mobile.continue_no_commit')}
          </button>
        )}
      </div>
    </div>
  );
}

export default CommitStep;
