import { OLT_TYPE_PORT, OLT_PORT_NAME } from '../global';
import utils from '../global/utils';
import AXIOS from '../axios';
import { API_OltSoftwareInfo } from '../global/API';
import GLOBAL from '../global';

const emptyVersion = () => ({
  name: '',
  release: '',
  valid: false,
  active: false,
  commit: false,
  timestamp: '',
});

export const createEmptyPanelInfo = () => ({
  id: '',
  ip: '',
  hostname: '',
  port: '',
  version1: emptyVersion(),
  version2: emptyVersion(),
  download_status: '',
  download_result: '',
  config_download_status: '',
  config_download_result: '',
});

export function getHostName(oltType, port) {
  if (oltType.startsWith('DF')) {
    return oltType;
  }
  return oltType + '(' + OLT_PORT_NAME[port] + ')';
}

export function getCardPort(oltInfo, cardIndex) {
  if (oltInfo.type && oltInfo.type.startsWith('DF')) {
    return OLT_TYPE_PORT.df;
  }
  const index = 'mf_lt' + cardIndex;
  return OLT_TYPE_PORT[index];
}

function fillPanelFromHardware(tmp, hostname) {
  const info = createEmptyPanelInfo();
  info.id = tmp.Ip + tmp.Port;
  info.ip = tmp.Ip;
  info.hostname = hostname;
  info.port = tmp.Port;

  const revisions = tmp.HardwareState.Component.Software.Software.Revisions.Revision;
  info.version1.name = revisions[0].Name;
  info.version1.release = revisions[0].Version;
  info.version1.valid = revisions[0].IsValid === 'true';
  info.version1.active = revisions[0].IsActive === 'true';
  info.version1.commit = revisions[0].IsCommitted === 'true';
  info.version1.timestamp = revisions[0].DownloadTimestamp;

  if (revisions.length > 1) {
    info.version2.name = revisions[1].Name;
    info.version2.release = revisions[1].Version;
    info.version2.valid = revisions[1].IsValid === 'true';
    info.version2.active = revisions[1].IsActive === 'true';
    info.version2.commit = revisions[1].IsCommitted === 'true';
    info.version2.timestamp = revisions[1].DownloadTimestamp;
  }

  const download = tmp.HardwareState.Component.Software.Software.Download;
  info.download_status = download.CurrentState.State;
  info.download_result = download.CurrentState.State === 'idle'
    ? download.LastDownloadState.State
    : download.CurrentState.State;

  const configDownload = tmp.HardwareState.Component.Software.Software.ConfigDownload;
  if (configDownload) {
    info.config_download_status = configDownload.CurrentState.State;
    info.config_download_result = configDownload.CurrentState.State === 'idle'
      ? configDownload.LastDownloadState.State
      : configDownload.CurrentState.State;
  }

  return info;
}

export function parseOltCardSoftwareInfo(tmp, oltType) {
  const hostname = getHostName(oltType, tmp.Port);
  return fillPanelFromHardware(tmp, hostname);
}

export function parseOltSoftwareInfo(data, oltType) {
  const list = [];
  Object.keys(data).forEach((key) => {
    const tmp = utils.isTestDataUsed() ? data[key] : JSON.parse(data[key]);
    list.push(parseOltCardSoftwareInfo(tmp, oltType));
  });
  return list;
}

export function getActiveSoftwareName(panel) {
  let name = '';
  const { version1, version2 } = panel;
  if (version1.active && version1.commit && !version2.active && !version2.commit && version2.valid) {
    name = version2.name;
  }
  if (version2.active && version2.commit && !version1.active && !version1.commit && version1.valid) {
    name = version1.name;
  }
  if (!version1.active && version1.commit && version2.active && !version2.commit && version2.valid) {
    name = version1.name;
  }
  if (!version2.active && version2.commit && version1.active && !version1.commit && version1.valid) {
    name = version2.name;
  }
  return name;
}

export function getCommitSoftwareName(panel) {
  let name = '';
  const { version1, version2 } = panel;
  if (!version1.active && version1.commit && version2.active && !version2.commit && version2.valid) {
    name = version2.name;
  }
  if (!version2.active && version2.commit && version1.active && !version1.commit && version1.valid) {
    name = version1.name;
  }
  return name;
}

function findRevisionByName(panel, name) {
  if (!name) {
    return null;
  }
  if (panel.version1.name === name) {
    return panel.version1;
  }
  if (panel.version2.name === name) {
    return panel.version2;
  }
  return null;
}

export function getPendingCommitName(panel, downloadedName) {
  const fromState = getCommitSoftwareName(panel);
  if (fromState) {
    return fromState;
  }
  const revision = findRevisionByName(panel, downloadedName);
  if (revision && revision.active && !revision.commit && revision.valid) {
    return downloadedName;
  }
  return '';
}

/** DF/MF unified: index 0 → NT, index >= 1 → LT */
export function getCardFirmwareType(card) {
  return card.index >= 1 ? 'LT' : 'NT';
}

export function firmwareDirHasMatchingFile(branch, dirName) {
  if (!branch || typeof branch !== 'object') {
    return false;
  }
  const folder = branch[dirName];
  if (!folder || typeof folder !== 'object') {
    return false;
  }
  return Object.prototype.hasOwnProperty.call(folder, dirName);
}

export function listFirmwareDirsByCard(tree, card) {
  const cardType = getCardFirmwareType(card);
  const branch = tree[cardType];
  if (!branch || typeof branch !== 'object') {
    return [];
  }
  return Object.keys(branch)
    .filter((key) => typeof branch[key] === 'object' && branch[key] !== null)
    .map((dir) => ({
      dir,
      label: cardType + '/' + dir,
      cardType,
      valid: firmwareDirHasMatchingFile(branch, dir),
    }));
}

export function flattenServerSoftware(tree, basePath = '') {
  const result = [];
  if (!tree || typeof tree !== 'object') {
    return result;
  }
  Object.keys(tree).forEach((key) => {
    const value = tree[key];
    const currentPath = basePath ? basePath + '/' + key : key;
    if (typeof value === 'object' && value !== null) {
      const childKeys = Object.keys(value);
      const hasNestedDirs = childKeys.some((k) => {
        const child = value[k];
        return typeof child === 'object' && child !== null && Object.keys(child).length > 0;
      });
      if (hasNestedDirs) {
        result.push(...flattenServerSoftware(value, currentPath));
      } else {
        childKeys.forEach((name) => {
          result.push({
            label: currentPath + '/' + name,
            path: currentPath,
            name,
          });
        });
      }
    }
  });
  return result;
}

export function listTopLevelSoftwareDirs(tree) {
  const result = [];
  if (!tree || typeof tree !== 'object') {
    return result;
  }
  Object.keys(tree).forEach((key) => {
    const value = tree[key];
    if (typeof value === 'object' && value !== null) {
      result.push({ label: key, dir: key });
    }
  });
  return result;
}

export function resolveDownloadPath(cardType, selectedDir) {
  return {
    path: cardType + '/' + selectedDir,
    name: selectedDir,
  };
}

export function buildDownloadPayload(panel, path, name) {
  return {
    oltId: panel.ip,
    dstPort: panel.port,
    action: 'download',
    url: utils.slashCompatibly(path) + '/' + name,
    name,
  };
}

export function buildActivePayload(panel, name) {
  return {
    oltId: panel.ip,
    dstPort: panel.port,
    action: 'active',
    name,
  };
}

export function buildCommitPayload(panel, name) {
  return {
    oltId: panel.ip,
    dstPort: panel.port,
    action: 'commit',
    name,
  };
}

export function fetchOltSoftwareInfo(oltIp, port, oltType) {
  return AXIOS
    .get(API_OltSoftwareInfo, { oltId: oltIp, dstPort: port })
    .then((res) => {
      if (res.data.status === GLOBAL.ERROR_NUM.Success) {
        const tmp = JSON.parse(res.data.olt_software_info);
        return parseOltCardSoftwareInfo(tmp, oltType);
      }
      return Promise.reject(res.data);
    });
}

export function pollDownloadStatus({ oltIp, port, oltType, onProgress, onComplete, onError, interval = 5000 }) {
  let timerId = null;
  let cancelled = false;

  const poll = () => {
    if (cancelled) {
      return;
    }
    fetchOltSoftwareInfo(oltIp, port, oltType)
      .then((panel) => {
        if (cancelled) {
          return;
        }
        if (panel.download_status === 'in-progress') {
          onProgress && onProgress(panel);
          timerId = setTimeout(poll, interval);
        } else if (panel.download_status === 'idle' && panel.download_result === 'successful') {
          onComplete && onComplete(panel);
        } else if (panel.download_status === 'failed' || (panel.download_status === 'idle' && panel.download_result !== 'successful')) {
          onError && onError(panel);
        } else {
          timerId = setTimeout(poll, interval);
        }
      })
      .catch((err) => {
        if (!cancelled) {
          onError && onError(err);
        }
      });
  };

  poll();

  return () => {
    cancelled = true;
    if (timerId) {
      clearTimeout(timerId);
    }
  };
}

export function pollDeviceReadyForCommit({
  oltIp,
  port,
  oltType,
  downloadedName,
  onWaiting,
  onReady,
  onTimeout,
  interval = 15000,
  timeout = null,
}) {
  let timerId = null;
  let cancelled = false;
  const startTime = Date.now();

  const poll = () => {
    if (cancelled) {
      return;
    }
    if (timeout != null && Date.now() - startTime >= timeout) {
      onTimeout && onTimeout();
      return;
    }
    fetchOltSoftwareInfo(oltIp, port, oltType)
      .then((panel) => {
        if (cancelled) {
          return;
        }
        const commitName = getPendingCommitName(panel, downloadedName);
        if (commitName) {
          onReady && onReady({ panel, commitName });
        } else {
          onWaiting && onWaiting({ connected: true, panel });
          timerId = setTimeout(poll, interval);
        }
      })
      .catch(() => {
        if (cancelled) {
          return;
        }
        onWaiting && onWaiting({ connected: false, panel: null });
        timerId = setTimeout(poll, interval);
      });
  };

  poll();

  return () => {
    cancelled = true;
    if (timerId) {
      clearTimeout(timerId);
    }
  };
}

export function getAvailableCards(olt) {
  const cards = [];
  if (olt.type && olt.type.startsWith('DF')) {
    cards.push({ index: 0, label: olt.type, port: OLT_TYPE_PORT.df });
  } else {
    cards.push({ index: 0, label: 'NT', port: OLT_TYPE_PORT.mf_lt0 });
    if (olt.ltCardStatus) {
      for (let i = 1; i <= 14; i++) {
        if (olt.ltCardStatus[i - 1] === 1) {
          const portKey = 'mf_lt' + i;
          cards.push({
            index: i,
            label: 'LT' + i,
            port: OLT_TYPE_PORT[portKey],
          });
        }
      }
    }
  }
  return cards;
}
