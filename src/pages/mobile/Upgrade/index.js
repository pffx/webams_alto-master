import React, { useState } from 'react';
import { useTranslation } from 'react-i18next';
import SelectDevice from './steps/SelectDevice';
import SelectFirmware from './steps/SelectFirmware';
import DownloadStep from './steps/DownloadStep';
import ActivateStep from './steps/ActivateStep';
import CommitStep from './steps/CommitStep';
import DoneStep from './steps/DoneStep';

const STEPS = ['device', 'firmware', 'download', 'activate', 'commit', 'done'];

const STEP_LABEL_KEYS = {
  device: 'mobile.step_device',
  firmware: 'mobile.step_firmware',
  download: 'mobile.step_download',
  activate: 'mobile.step_activate',
  commit: 'mobile.step_commit',
};

function MobileUpgradePage() {
  const { t } = useTranslation();
  const [stepIndex, setStepIndex] = useState(0);
  const [deviceSelection, setDeviceSelection] = useState(null);
  const [firmwareSelection, setFirmwareSelection] = useState(null);
  const [downloadResult, setDownloadResult] = useState(null);
  const [activateResult, setActivateResult] = useState(null);
  const [commitResult, setCommitResult] = useState(null);

  const currentStep = STEPS[stepIndex];

  const handleRestart = () => {
    setStepIndex(0);
    setDeviceSelection(null);
    setFirmwareSelection(null);
    setDownloadResult(null);
    setActivateResult(null);
    setCommitResult(null);
  };

  const renderStepIndicator = () => (
    <div className="mobile-step-indicator">
      {STEPS.slice(0, -1).map((step, index) => (
        <div
          key={step}
          className={
            'mobile-step-indicator__item' +
            (index < stepIndex ? ' mobile-step-indicator__item--done' : '') +
            (index === stepIndex ? ' mobile-step-indicator__item--active' : '')
          }
        >
          <div className="mobile-step-indicator__dot" />
          <span className="mobile-step-indicator__label">{t(STEP_LABEL_KEYS[step])}</span>
        </div>
      ))}
    </div>
  );

  const renderStep = () => {
    switch (currentStep) {
      case 'device':
        return (
          <SelectDevice
            onNext={(data) => {
              setDeviceSelection(data);
              setStepIndex(1);
            }}
          />
        );
      case 'firmware':
        return (
          <SelectFirmware
            deviceSelection={deviceSelection}
            onBack={() => setStepIndex(0)}
            onNext={(data) => {
              setFirmwareSelection(data);
              setStepIndex(2);
            }}
          />
        );
      case 'download':
        return (
          <DownloadStep
            deviceSelection={deviceSelection}
            firmwareSelection={firmwareSelection}
            onBack={() => setStepIndex(1)}
            onNext={(data) => {
              setDownloadResult(data);
              setStepIndex(3);
            }}
          />
        );
      case 'activate':
        return (
          <ActivateStep
            deviceSelection={deviceSelection}
            downloadResult={downloadResult}
            onBack={() => setStepIndex(2)}
            onNext={(data) => {
              setActivateResult(data);
              setStepIndex(4);
            }}
          />
        );
      case 'commit':
        return (
          <CommitStep
            deviceSelection={deviceSelection}
            activateResult={activateResult}
            downloadedName={firmwareSelection && firmwareSelection.firmware ? firmwareSelection.firmware.name : ''}
            onBack={() => setStepIndex(3)}
            onNext={(data) => {
              setCommitResult(data);
              setStepIndex(5);
            }}
          />
        );
      case 'done':
        return (
          <DoneStep
            deviceSelection={deviceSelection}
            commitResult={commitResult}
            onRestart={handleRestart}
          />
        );
      default:
        return null;
    }
  };

  return (
    <div className="mobile-upgrade">
      {currentStep !== 'done' && renderStepIndicator()}
      {renderStep()}
    </div>
  );
}

export default MobileUpgradePage;
