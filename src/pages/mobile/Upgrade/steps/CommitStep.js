import React, { useCallback, useEffect, useRef, useState } from 'react';
import { useTranslation } from 'react-i18next';
import { toast } from 'react-toastify';
import AXIOS from '../../../../axios';
import GLOBAL, { TOAST_CONF } from '../../../../global';
import { API_SoftwareAction } from '../../../../global/API';
import {
  buildCommitPayload,
  getPendingCommitName,
  fetchOltSoftwareInfo,
  pollDeviceReadyForCommit,
} from '../../../../utils/softwareUpgrade';

function formatElapsed(seconds) {
  const m = Math.floor(seconds / 60);
  const s = seconds % 60;
  if (m > 0) {
    return m + ':' + String(s).padStart(2, '0');
  }
  return String(s) + 's';
}

function CommitStep({ deviceSelection, activateResult, downloadedName, onNext, onBack }) {
  const { t } = useTranslation();
  const { olt, card } = deviceSelection;
  const needsRebootWait = activateResult.needsRebootWait;
  const [panel, setPanel] = useState(activateResult.panel);
  const [phase, setPhase] = useState(needsRebootWait ? 'waiting' : 'loading');
  const [commitName, setCommitName] = useState('');
  const [waitConnected, setWaitConnected] = useState(false);
  const [elapsedSeconds, setElapsedSeconds] = useState(0);
  const [loading, setLoading] = useState(false);
  const [checking, setChecking] = useState(false);
  const [error, setError] = useState('');
  const stopPollRef = useRef(null);
  const startTimeRef = useRef(Date.now());

  const applyPanelResult = useCallback((freshPanel) => {
    const name = getPendingCommitName(freshPanel, downloadedName);
    setPanel(freshPanel);
    setCommitName(name);
    setPhase(name ? 'ready' : 'no_commit');
  }, [downloadedName]);

  useEffect(() => {
    if (phase !== 'waiting') {
      return undefined;
    }
    const id = setInterval(() => {
      setElapsedSeconds(Math.floor((Date.now() - startTimeRef.current) / 1000));
    }, 1000);
    return () => clearInterval(id);
  }, [phase]);

  useEffect(() => {
    if (needsRebootWait) {
      startTimeRef.current = Date.now();
      stopPollRef.current = pollDeviceReadyForCommit({
        oltIp: olt.ip,
        port: card.port,
        oltType: olt.type,
        downloadedName,
        onWaiting: ({ connected, panel: freshPanel }) => {
          setWaitConnected(connected);
          if (freshPanel) {
            setPanel(freshPanel);
          }
        },
        onReady: ({ panel: freshPanel, commitName: name }) => {
          setPanel(freshPanel);
          setCommitName(name);
          setPhase('ready');
        },
      });
    } else {
      let cancelled = false;
      fetchOltSoftwareInfo(olt.ip, card.port, olt.type)
        .then((freshPanel) => {
          if (!cancelled) {
            applyPanelResult(freshPanel);
          }
        })
        .catch(() => {
          if (!cancelled) {
            setPhase('no_commit');
          }
        });
      return () => {
        cancelled = true;
      };
    }

    return () => {
      if (stopPollRef.current) {
        stopPollRef.current();
      }
    };
  }, [needsRebootWait, olt, card, downloadedName, applyPanelResult]);

  const handleCheckNow = () => {
    if (checking) {
      return;
    }
    setChecking(true);
    fetchOltSoftwareInfo(olt.ip, card.port, olt.type)
      .then((freshPanel) => {
        setChecking(false);
        const name = getPendingCommitName(freshPanel, downloadedName);
        setPanel(freshPanel);
        setWaitConnected(true);
        if (name) {
          if (stopPollRef.current) {
            stopPollRef.current();
            stopPollRef.current = null;
          }
          setCommitName(name);
          setPhase('ready');
        }
      })
      .catch(() => {
        setChecking(false);
        setWaitConnected(false);
      });
  };

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

  if (phase === 'loading') {
    return <div className="mobile-step-loading">{t('mobile.loading')}</div>;
  }

  if (phase === 'waiting') {
    return (
      <div className="mobile-step">
        <h2 className="mobile-step__title">{t('mobile.step_commit')}</h2>

        <div className="mobile-progress">
          <div className="mobile-progress__bar" />
          <p>
            {waitConnected
              ? t('mobile.waiting_commit_ready')
              : t('mobile.waiting_reboot')}
          </p>
          <p className="mobile-wait-elapsed">
            {t('mobile.elapsed_wait', { time: formatElapsed(elapsedSeconds) })}
          </p>
        </div>

        <div className="mobile-actions">
          <button type="button" className="mobile-btn mobile-btn--secondary" onClick={onBack}>
            {t('mobile.back')}
          </button>
          <button
            type="button"
            className="mobile-btn mobile-btn--primary"
            disabled={checking}
            onClick={handleCheckNow}
          >
            {checking ? t('mobile.processing') : t('mobile.check_now')}
          </button>
        </div>
      </div>
    );
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
