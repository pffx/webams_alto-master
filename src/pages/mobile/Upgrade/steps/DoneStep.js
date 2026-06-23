import React from 'react';
import { useTranslation } from 'react-i18next';

function DoneStep({ deviceSelection, commitResult, onRestart }) {
  const { t } = useTranslation();
  const { olt, card } = deviceSelection;
  const panel = commitResult.panel;

  return (
    <div className="mobile-step mobile-step--done">
      <div className="mobile-done-icon">✓</div>
      <h2 className="mobile-step__title">{t('mobile.step_done')}</h2>
      <p className="mobile-step__hint">{t('mobile.upgrade_complete')}</p>

      <div className="mobile-info-card">
        <div className="mobile-info-row">
          <span>{t('proper_noun.ip')}</span>
          <span>{olt.ip}</span>
        </div>
        <div className="mobile-info-row">
          <span>{t('mobile.port')}</span>
          <span>{card.port} ({card.label})</span>
        </div>
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

      <div className="mobile-actions">
        <button type="button" className="mobile-btn mobile-btn--primary mobile-btn--block" onClick={onRestart}>
          {t('mobile.upgrade_another')}
        </button>
      </div>
    </div>
  );
}

export default DoneStep;
