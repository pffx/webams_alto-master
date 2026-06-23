import React, { useState } from 'react';
import SelectDevice from './steps/SelectDevice';
import SelectFirmware from './steps/SelectFirmware';
import DownloadStep from './steps/DownloadStep';
import ActivateStep from './steps/ActivateStep';
import CommitStep from './steps/CommitStep';
import DoneStep from './steps/DoneStep';

const STEPS = ['device', 'firmware', 'download', 'activate', 'commit', 'done'];

function MobileUpgradePage() {
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
            'mobile-step-indicator__dot' +
            (index < stepIndex ? ' mobile-step-indicator__dot--done' : '') +
            (index === stepIndex ? ' mobile-step-indicator__dot--active' : '')
          }
        />
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
