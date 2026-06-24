import React, { useEffect, useState } from 'react';
import { useTranslation } from 'react-i18next';
import AXIOS from '../../../../axios';
import GLOBAL from '../../../../global';
import { API_ServerSoftware } from '../../../../global/API';
import {
  listFirmwareDirsByCard,
  resolveDownloadPath,
  fetchOltSoftwareInfo,
  resolveUpgradePhase,
  getPreCommitName,
  getPendingCommitName,
} from '../../../../utils/softwareUpgrade';

function VersionPanel({ panel, t }) {
  if (!panel) {
    return null;
  }
  return (
    <div className="mobile-info-card">
      <h3 className="mobile-info-card__title">{t('mobile.current_versions')}</h3>
      {panel.version1.name && (
        <div className="mobile-info-row">
          <span>{panel.version1.name}</span>
          <span>
            {panel.version1.active ? t('mobile.active') : ''}
            {panel.version1.commit ? ' / ' + t('mobile.committed') : ''}
          </span>
        </div>
      )}
      {panel.version2.name && (
        <div className="mobile-info-row">
          <span>{panel.version2.name}</span>
          <span>
            {panel.version2.active ? t('mobile.active') : ''}
            {panel.version2.commit ? ' / ' + t('mobile.committed') : ''}
          </span>
        </div>
      )}
    </div>
  );
}

function SelectFirmware({
  deviceSelection,
  onNext,
  onBack,
  onPreCommit,
  onResumePostCommit,
  hasSession,
}) {
  const { t } = useTranslation();
  const { olt, card } = deviceSelection;
  const [firmwareDirs, setFirmwareDirs] = useState([]);
  const [panel, setPanel] = useState(null);
  const [selectedDir, setSelectedDir] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState('');
  const [phase, setPhase] = useState(null);
  const [pendingCommitBanner, setPendingCommitBanner] = useState(null);

  useEffect(() => {
    let cancelled = false;

    Promise.all([
      AXIOS.get(API_ServerSoftware),
      fetchOltSoftwareInfo(olt.ip, card.port, olt.type),
    ])
      .then(([serverRes, currentPanel]) => {
        if (cancelled) {
          return;
        }
        if (serverRes.data.status === GLOBAL.ERROR_NUM.Success) {
          const folders = JSON.parse(serverRes.data.software_list);
          setFirmwareDirs(listFirmwareDirsByCard(folders, card));
        }
        setPanel(currentPanel);
        const resolvedPhase = resolveUpgradePhase(currentPanel);
        setPhase(resolvedPhase);

        if (!hasSession) {
          const pendingName = getPendingCommitName(currentPanel, null);
          if (pendingName && resolvedPhase === 'post_activate_commit') {
            setPendingCommitBanner(pendingName);
          }
        }
        setLoading(false);
      })
      .catch(() => {
        if (!cancelled) {
          setLoading(false);
          setError(t('message.server_error'));
        }
      });

    return () => {
      cancelled = true;
    };
  }, [olt, card, t, hasSession]);

  const canProceed = selectedDir && selectedDir.valid && panel && phase === 'ready_for_download';

  const handleNext = () => {
    if (!canProceed) {
      return;
    }
    const { path, name } = resolveDownloadPath(selectedDir.cardType, selectedDir.dir);
    onNext({
      firmware: {
        dir: selectedDir.dir,
        cardType: selectedDir.cardType,
        path,
        name,
        label: path + '/' + name,
      },
      panel,
    });
  };

  const handlePreCommit = () => {
    if (!panel) {
      return;
    }
    onPreCommit({ panel, commitName: getPreCommitName(panel) });
  };

  const handleResumePostCommit = () => {
    if (!panel || !pendingCommitBanner) {
      return;
    }
    onResumePostCommit({
      panel,
      commitName: pendingCommitBanner,
      needsRebootWait: true,
    });
  };

  if (loading) {
    return <div className="mobile-step-loading">{t('mobile.loading')}</div>;
  }

  return (
    <div className="mobile-step">
      <h2 className="mobile-step__title">{t('mobile.step_version_check')}</h2>
      {error && <div className="mobile-feedback mobile-feedback--error">{error}</div>}

      {pendingCommitBanner && (
        <div className="mobile-info-card">
          <p>{t('mobile.pending_commit_banner', { name: pendingCommitBanner })}</p>
          <button
            type="button"
            className="mobile-btn mobile-btn--primary"
            onClick={handleResumePostCommit}
          >
            {t('mobile.resume_post_commit')}
          </button>
        </div>
      )}

      <VersionPanel panel={panel} t={t} />

      {phase === 'pre_commit_required' && (
        <div className="mobile-info-card">
          <p>{t('mobile.pre_commit_required_hint')}</p>
        </div>
      )}

      {phase === 'ready_for_download' && (
        <>
          <p className="mobile-step__hint">{t('mobile.select_firmware_dir')}</p>
          <div className="mobile-list mobile-list--scroll">
            {firmwareDirs.length === 0 ? (
              <div className="mobile-empty">{t('mobile.no_firmware')}</div>
            ) : (
              firmwareDirs.map((fw) => (
                <button
                  key={fw.label}
                  type="button"
                  disabled={!fw.valid}
                  className={
                    'mobile-list-item' +
                    (selectedDir && selectedDir.dir === fw.dir ? ' mobile-list-item--selected' : '') +
                    (!fw.valid ? ' mobile-list-item--disabled' : '')
                  }
                  onClick={() => fw.valid && setSelectedDir(fw)}
                >
                  <span className="mobile-list-item__title">{fw.label}</span>
                  {!fw.valid && (
                    <span className="mobile-list-item__meta">{t('mobile.firmware_file_missing')}</span>
                  )}
                </button>
              ))
            )}
          </div>
        </>
      )}

      <div className="mobile-actions">
        <button type="button" className="mobile-btn mobile-btn--secondary" onClick={onBack}>
          {t('mobile.back')}
        </button>
        {phase === 'pre_commit_required' ? (
          <button
            type="button"
            className="mobile-btn mobile-btn--primary"
            onClick={handlePreCommit}
          >
            {t('mobile.commit_current_version')}
          </button>
        ) : (
          <button
            type="button"
            className="mobile-btn mobile-btn--primary"
            disabled={!canProceed}
            onClick={handleNext}
          >
            {t('mobile.next')}
          </button>
        )}
      </div>
    </div>
  );
}

export default SelectFirmware;
