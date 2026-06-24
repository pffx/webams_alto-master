import React, { useEffect, useState } from 'react';
import { useTranslation } from 'react-i18next';
import AXIOS from '../../../../axios';
import GLOBAL from '../../../../global';
import { API_ServerSoftware } from '../../../../global/API';
import {
  listTopLevelSoftwareDirs,
  resolveDownloadPath,
  fetchOltSoftwareInfo,
} from '../../../../utils/softwareUpgrade';

function SelectFirmware({ deviceSelection, onNext, onBack }) {
  const { t } = useTranslation();
  const { olt, card } = deviceSelection;
  const [firmwareDirs, setFirmwareDirs] = useState([]);
  const [panel, setPanel] = useState(null);
  const [selectedDir, setSelectedDir] = useState(null);
  const [firmwareInput, setFirmwareInput] = useState('');
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState('');

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
          setFirmwareDirs(listTopLevelSoftwareDirs(folders));
        }
        setPanel(currentPanel);
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
  }, [olt, card, t]);

  const canProceed = selectedDir && firmwareInput.trim() && panel;

  const handleNext = () => {
    if (!canProceed) {
      return;
    }
    const trimmed = firmwareInput.trim();
    const { path, name } = resolveDownloadPath(selectedDir.dir, trimmed);
    onNext({
      firmware: {
        dir: selectedDir.dir,
        userInput: trimmed,
        path,
        name,
        label: path + '/' + name,
      },
      panel,
    });
  };

  if (loading) {
    return <div className="mobile-step-loading">{t('mobile.loading')}</div>;
  }

  return (
    <div className="mobile-step">
      <h2 className="mobile-step__title">{t('mobile.step_firmware')}</h2>
      {error && <div className="mobile-feedback mobile-feedback--error">{error}</div>}

      {panel && (
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
      )}

      <p className="mobile-step__hint">{t('mobile.select_firmware_dir')}</p>
      <div className="mobile-list mobile-list--scroll">
        {firmwareDirs.length === 0 ? (
          <div className="mobile-empty">{t('mobile.no_firmware')}</div>
        ) : (
          firmwareDirs.map((fw) => (
            <button
              key={fw.dir}
              type="button"
              className={'mobile-list-item' + (selectedDir && selectedDir.dir === fw.dir ? ' mobile-list-item--selected' : '')}
              onClick={() => setSelectedDir(fw)}
            >
              <span className="mobile-list-item__title">{fw.dir}</span>
            </button>
          ))
        )}
      </div>

      <label className="mobile-field">
        <span className="mobile-field__label">{t('mobile.firmware_name_input')}</span>
        <input
          className="mobile-field__input"
          type="text"
          value={firmwareInput}
          onChange={(e) => setFirmwareInput(e.target.value)}
          placeholder={t('mobile.firmware_name_placeholder')}
        />
      </label>

      <div className="mobile-actions">
        <button type="button" className="mobile-btn mobile-btn--secondary" onClick={onBack}>
          {t('mobile.back')}
        </button>
        <button
          type="button"
          className="mobile-btn mobile-btn--primary"
          disabled={!canProceed}
          onClick={handleNext}
        >
          {t('mobile.next')}
        </button>
      </div>
    </div>
  );
}

export default SelectFirmware;
