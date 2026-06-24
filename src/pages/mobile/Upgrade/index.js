import React, { useEffect, useMemo, useState } from 'react';
import { useTranslation } from 'react-i18next';
import SelectDevice from './steps/SelectDevice';
import SelectFirmware from './steps/SelectFirmware';
import DownloadStep from './steps/DownloadStep';
import ActivateStep from './steps/ActivateStep';
import CommitStep from './steps/CommitStep';
import DoneStep from './steps/DoneStep';
import { saveSession, loadSession, clearSession } from '../../../utils/mobileUpgradeSession';

const STEP_LABEL_KEYS = {
  device: 'mobile.step_device',
  version_check: 'mobile.step_version_check',
  pre_commit: 'mobile.step_pre_commit',
  download: 'mobile.step_download',
  activate: 'mobile.step_activate',
  post_commit: 'mobile.step_post_commit',
};

const BASE_INDICATOR_STEPS = ['device', 'version_check', 'download', 'activate', 'post_commit'];

function MobileUpgradePage() {
  const { t } = useTranslation();
  const [currentStep, setCurrentStep] = useState('device');
  const [showPreCommitInIndicator, setShowPreCommitInIndicator] = useState(false);
  const [flowMode, setFlowMode] = useState('normal');
  const [deviceSelection, setDeviceSelection] = useState(null);
  const [firmwareSelection, setFirmwareSelection] = useState(null);
  const [downloadResult, setDownloadResult] = useState(null);
  const [activateResult, setActivateResult] = useState(null);
  const [commitResult, setCommitResult] = useState(null);
  const [sessionRestored, setSessionRestored] = useState(false);
  const [hasSession, setHasSession] = useState(false);

  useEffect(() => {
    const session = loadSession();
    if (session) {
      setHasSession(true);
      if (session.deviceSelection) {
        setDeviceSelection(session.deviceSelection);
      }
      if (session.firmwareSelection) {
        setFirmwareSelection(session.firmwareSelection);
      }
      if (session.downloadResult) {
        setDownloadResult(session.downloadResult);
      }
      if (session.activateResult) {
        setActivateResult(session.activateResult);
      }
      if (session.commitResult) {
        setCommitResult(session.commitResult);
      }
      if (session.flowMode) {
        setFlowMode(session.flowMode);
      }
      if (session.showPreCommitInIndicator) {
        setShowPreCommitInIndicator(session.showPreCommitInIndicator);
      }

      const step = session.currentStep
        || (session.stepIndex != null ? ['device', 'version_check', 'download', 'activate', 'post_commit', 'done'][session.stepIndex] : null)
        || (session.activateResult && session.activateResult.needsRebootWait ? 'post_commit' : 'device');

      if (step === 'firmware') {
        setCurrentStep('version_check');
      } else if (step === 'commit') {
        setCurrentStep('post_commit');
      } else {
        setCurrentStep(step);
      }
    }
    setSessionRestored(true);
  }, []);

  useEffect(() => {
    if (!sessionRestored) {
      return;
    }
    if (currentStep === 'device' && !deviceSelection) {
      return;
    }
    saveSession({
      currentStep,
      deviceSelection,
      firmwareSelection,
      downloadResult,
      activateResult,
      commitResult,
      flowMode,
      showPreCommitInIndicator,
    });
  }, [
    sessionRestored,
    currentStep,
    deviceSelection,
    firmwareSelection,
    downloadResult,
    activateResult,
    commitResult,
    flowMode,
    showPreCommitInIndicator,
  ]);

  const indicatorSteps = useMemo(() => {
    if (showPreCommitInIndicator) {
      return ['device', 'version_check', 'pre_commit', 'download', 'activate', 'post_commit'];
    }
    return BASE_INDICATOR_STEPS;
  }, [showPreCommitInIndicator]);

  const stepIndex = indicatorSteps.indexOf(
    currentStep === 'pre_commit' ? 'pre_commit'
      : currentStep === 'post_commit' ? 'post_commit'
        : currentStep,
  );

  const handleRestart = () => {
    clearSession();
    setHasSession(false);
    setCurrentStep('device');
    setShowPreCommitInIndicator(false);
    setFlowMode('normal');
    setDeviceSelection(null);
    setFirmwareSelection(null);
    setDownloadResult(null);
    setActivateResult(null);
    setCommitResult(null);
  };

  const downloadedName = firmwareSelection && firmwareSelection.firmware
    ? firmwareSelection.firmware.name
    : '';

  const renderStepIndicator = () => (
    <div className="mobile-step-indicator">
      {indicatorSteps.map((step, index) => (
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
              setCurrentStep('version_check');
            }}
          />
        );
      case 'version_check':
        return (
          <SelectFirmware
            deviceSelection={deviceSelection}
            hasSession={hasSession}
            onBack={() => setCurrentStep('device')}
            onPreCommit={({ panel }) => {
              setShowPreCommitInIndicator(true);
              setFlowMode('pre_commit');
              setActivateResult({ panel, needsRebootWait: false });
              setCurrentStep('pre_commit');
            }}
            onResumePostCommit={(data) => {
              setFlowMode('post_activate');
              setActivateResult(data);
              setCurrentStep('post_commit');
            }}
            onNext={(data) => {
              setFirmwareSelection(data);
              setCurrentStep('download');
            }}
          />
        );
      case 'pre_commit':
        return (
          <CommitStep
            deviceSelection={deviceSelection}
            activateResult={activateResult}
            downloadedName=""
            flowMode="pre_commit"
            onBack={() => setCurrentStep('version_check')}
            onNext={() => {
              setFlowMode('normal');
              setActivateResult(null);
              setCurrentStep('version_check');
            }}
          />
        );
      case 'download':
        return (
          <DownloadStep
            deviceSelection={deviceSelection}
            firmwareSelection={firmwareSelection}
            onBack={() => setCurrentStep('version_check')}
            onNext={(data) => {
              setDownloadResult(data);
              setCurrentStep('activate');
            }}
          />
        );
      case 'activate':
        return (
          <ActivateStep
            deviceSelection={deviceSelection}
            downloadResult={downloadResult}
            onBack={() => setCurrentStep('download')}
            onNext={(data) => {
              setActivateResult(data);
              setFlowMode('post_activate');
              setCurrentStep('post_commit');
            }}
          />
        );
      case 'post_commit':
        return (
          <CommitStep
            deviceSelection={deviceSelection}
            activateResult={activateResult}
            downloadedName={downloadedName}
            flowMode="post_activate"
            onBack={() => setCurrentStep('activate')}
            onNext={(data) => {
              setCommitResult(data);
              setFlowMode('normal');
              clearSession();
              setHasSession(false);
              setCurrentStep('done');
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
