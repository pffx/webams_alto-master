import React, { useEffect, useRef, useState, useCallback } from 'react';
import { useTranslation } from 'react-i18next';
import { toast } from 'react-toastify';
import AXIOS from '../../../../axios';
import GLOBAL, { TOAST_CONF } from '../../../../global';
import { API_SoftwareAction } from '../../../../global/API';
import {
  buildDownloadPayload,
  pollDownloadStatus,
} from '../../../../utils/softwareUpgrade';

function DownloadStep({ deviceSelection, firmwareSelection, onNext, onBack }) {
  const { t } = useTranslation();
  const { olt, card } = deviceSelection;
  const { firmware, panel } = firmwareSelection;
  const [status, setStatus] = useState('idle');
  const [downloadResult, setDownloadResult] = useState('');
  const [panelState, setPanelState] = useState(panel);
  const stopPollRef = useRef(null);
  const startedRef = useRef(false);

  const startDownload = useCallback(() => {
    setStatus('starting');
    const payload = buildDownloadPayload(
      { ip: olt.ip, port: card.port },
      firmware.path,
      firmware.name,
    );

    AXIOS
      .put(API_SoftwareAction, payload)
      .then((res) => {
        if (res.data.status === GLOBAL.ERROR_NUM.Success) {
          setStatus('in-progress');
          toast.success(t('mobile.download_started'), TOAST_CONF);
          stopPollRef.current = pollDownloadStatus({
            oltIp: olt.ip,
            port: card.port,
            oltType: olt.type,
            onProgress: (p) => {
              setPanelState(p);
              setStatus('in-progress');
            },
            onComplete: (p) => {
              setPanelState(p);
              setStatus('success');
            },
            onError: (p) => {
              setPanelState(p);
              setDownloadResult(p.download_result || t('mobile.download_failed'));
              setStatus('failed');
            },
          });
        } else {
          setStatus('failed');
          setDownloadResult(t('mobile.download_failed'));
        }
      })
      .catch(() => {
        setStatus('failed');
        setDownloadResult(t('message.server_error'));
      });
  }, [olt, card, firmware, t]);

  useEffect(() => {
    if (panel.download_status === 'in-progress') {
      setStatus('in-progress');
      stopPollRef.current = pollDownloadStatus({
        oltIp: olt.ip,
        port: card.port,
        oltType: olt.type,
        onProgress: (p) => {
          setPanelState(p);
          setStatus('in-progress');
        },
        onComplete: (p) => {
          setPanelState(p);
          setStatus('success');
        },
        onError: (p) => {
          setPanelState(p);
          setDownloadResult(p.download_result || t('mobile.download_failed'));
          setStatus('failed');
        },
      });
    } else if (!startedRef.current) {
      startedRef.current = true;
      startDownload();
    }
    return () => {
      if (stopPollRef.current) {
        stopPollRef.current();
      }
    };
  }, [olt, card, panel.download_status, startDownload, t]);

  const handleNext = () => {
    onNext({ panel: panelState });
  };

  const handleRetry = () => {
    if (stopPollRef.current) {
      stopPollRef.current();
    }
    setDownloadResult('');
    startDownload();
  };

  return (
    <div className="mobile-step">
      <h2 className="mobile-step__title">{t('mobile.step_download')}</h2>

      <div className="mobile-info-card">
        <div className="mobile-info-row">
          <span>{t('mobile.target_firmware')}</span>
          <span>{firmware.label}</span>
        </div>
        <div className="mobile-info-row">
          <span>{t('mobile.download_status')}</span>
          <span className={'mobile-status mobile-status--' + status}>{status}</span>
        </div>
        {downloadResult && (
          <div className="mobile-feedback mobile-feedback--error">{downloadResult}</div>
        )}
      </div>

      {status === 'in-progress' && (
        <div className="mobile-progress">
          <div className="mobile-progress__bar" />
          <p>{t('mobile.downloading')}</p>
        </div>
      )}

      <div className="mobile-actions">
        <button type="button" className="mobile-btn mobile-btn--secondary" onClick={onBack}>
          {t('mobile.back')}
        </button>
        {status === 'failed' && (
          <button type="button" className="mobile-btn mobile-btn--primary" onClick={handleRetry}>
            {t('mobile.retry')}
          </button>
        )}
        {status === 'success' && (
          <button type="button" className="mobile-btn mobile-btn--primary" onClick={handleNext}>
            {t('mobile.next')}
          </button>
        )}
      </div>
    </div>
  );
}

export default DownloadStep;
