import React, { useEffect, useState } from 'react';
import { useTranslation } from 'react-i18next';
import AXIOS from '../../../../axios';
import { API_OltList } from '../../../../global/API';
import utils from '../../../../global/utils';
import { getAvailableCards } from '../../../../utils/softwareUpgrade';

function SelectDevice({ onNext }) {
  const { t } = useTranslation();
  const [olts, setOlts] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState('');
  const [selectedOlt, setSelectedOlt] = useState(null);
  const [selectedCard, setSelectedCard] = useState(null);

  useEffect(() => {
    AXIOS
      .get(API_OltList)
      .then((res) => {
        setLoading(false);
        if (res.data.status === 200) {
          const list = (res.data.olt_list || []).map((item) => ({
            ip: item.IP,
            hostname: item.HostName,
            status: item.Status === 'Connected' ? 'UP' : 'Down',
            type: item.OltType,
            ltCardStatus: utils.generateLTCardPlanned(item),
          }));
          setOlts(list);
        } else {
          setError(t('mobile.load_olt_failed'));
        }
      })
      .catch(() => {
        setLoading(false);
        setError(t('message.server_error'));
      });
  }, [t]);

  const cards = selectedOlt ? getAvailableCards(selectedOlt) : [];

  const handleNext = () => {
    if (!selectedOlt || !selectedCard) {
      return;
    }
    onNext({
      olt: selectedOlt,
      card: selectedCard,
    });
  };

  if (loading) {
    return <div className="mobile-step-loading">{t('mobile.loading')}</div>;
  }

  return (
    <div className="mobile-step">
      <h2 className="mobile-step__title">{t('mobile.step_device')}</h2>
      {error && <div className="mobile-feedback mobile-feedback--error">{error}</div>}

      <p className="mobile-step__hint">{t('mobile.select_olt')}</p>
      <div className="mobile-list">
        {olts.length === 0 ? (
          <div className="mobile-empty">{t('mobile.no_olt')}</div>
        ) : (
          olts.map((olt) => (
            <button
              key={olt.ip + olt.hostname}
              type="button"
              className={'mobile-list-item' + (selectedOlt && selectedOlt.ip === olt.ip ? ' mobile-list-item--selected' : '')}
              onClick={() => {
                setSelectedOlt(olt);
                setSelectedCard(null);
              }}
            >
              <span className="mobile-list-item__title">{olt.ip}</span>
              <span className="mobile-list-item__meta">{olt.hostname} · {olt.type}</span>
              <span className={'mobile-badge mobile-badge--' + (olt.status === 'UP' ? 'up' : 'down')}>
                {olt.status}
              </span>
            </button>
          ))
        )}
      </div>

      {selectedOlt && (
        <>
          <p className="mobile-step__hint">{t('mobile.select_card')}</p>
          <div className="mobile-list">
            {cards.map((card) => (
              <button
                key={card.port}
                type="button"
                className={'mobile-list-item' + (selectedCard && selectedCard.port === card.port ? ' mobile-list-item--selected' : '')}
                onClick={() => setSelectedCard(card)}
              >
                <span className="mobile-list-item__title">{card.label}</span>
                <span className="mobile-list-item__meta">{t('mobile.port')}: {card.port}</span>
              </button>
            ))}
          </div>
        </>
      )}

      <div className="mobile-actions">
        <button
          type="button"
          className="mobile-btn mobile-btn--primary"
          disabled={!selectedOlt || !selectedCard}
          onClick={handleNext}
        >
          {t('mobile.next')}
        </button>
      </div>
    </div>
  );
}

export default SelectDevice;
