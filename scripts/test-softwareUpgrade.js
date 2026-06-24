/**
 * Regression tests for software upgrade pure logic (self-contained).
 * Run: node scripts/test-softwareUpgrade.js
 */

const assert = require('assert');

function getActiveSoftwareName(panel) {
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

function getCommitSoftwareName(panel) {
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

function flattenServerSoftware(tree, basePath = '') {
  const result = [];
  if (!tree || typeof tree !== 'object') return result;
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
          result.push({ label: currentPath + '/' + name, path: currentPath, name });
        });
      }
    }
  });
  return result;
}

function getPendingCommitName(panel, downloadedName) {
  const fromState = getCommitSoftwareName(panel);
  if (fromState) return fromState;
  const rev = panel.version1.name === downloadedName ? panel.version1
    : panel.version2.name === downloadedName ? panel.version2 : null;
  if (rev && rev.active && !rev.commit && rev.valid) return downloadedName;
  return '';
}

function getCardFirmwareType(card) {
  return card.index >= 1 ? 'LT' : 'NT';
}

function firmwareDirHasMatchingFile(branch, dirName) {
  if (!branch || typeof branch !== 'object') return false;
  const folder = branch[dirName];
  if (!folder || typeof folder !== 'object') return false;
  return Object.prototype.hasOwnProperty.call(folder, dirName);
}

function listFirmwareDirsByCard(tree, card) {
  const cardType = getCardFirmwareType(card);
  const branch = tree[cardType];
  if (!branch || typeof branch !== 'object') return [];
  return Object.keys(branch)
    .filter((key) => typeof branch[key] === 'object' && branch[key] !== null)
    .map((dir) => ({
      dir,
      label: cardType + '/' + dir,
      cardType,
      valid: firmwareDirHasMatchingFile(branch, dir),
    }));
}

function resolveDownloadPath(cardType, selectedDir) {
  return { path: cardType + '/' + selectedDir, name: selectedDir };
}

const OLT_TYPE_PORT = {
  df: 830,
  mf_lt0: 832,
  mf_lt1: 833,
  mf_lt3: 835,
};

function getAvailableCards(olt) {
  const cards = [];
  if (olt.type && olt.type.startsWith('DF')) {
    cards.push({ index: 0, label: olt.type, port: OLT_TYPE_PORT.df });
  } else {
    cards.push({ index: 0, label: 'NT', port: OLT_TYPE_PORT.mf_lt0 });
    if (olt.ltCardStatus) {
      for (let i = 1; i <= 14; i++) {
        if (olt.ltCardStatus[i - 1] === 1) {
          cards.push({ index: i, label: 'LT' + i, port: OLT_TYPE_PORT['mf_lt' + i] });
        }
      }
    }
  }
  return cards;
}

function testGetActiveSoftwareName() {
  const panel = {
    version1: { active: true, commit: true, valid: true, name: 'v1' },
    version2: { active: false, commit: false, valid: true, name: 'v2' },
  };
  assert.strictEqual(getActiveSoftwareName(panel), 'v2');
}

function testGetCommitSoftwareName() {
  const panel = {
    version1: { active: true, commit: false, valid: true, name: 'v1' },
    version2: { active: false, commit: true, valid: true, name: 'v2' },
  };
  assert.strictEqual(getCommitSoftwareName(panel), 'v1');
}

function testFlattenServerSoftware() {
  const tree = {
    lightspan_2203: {
      L6GQAG: { L6GQAG2203: {}, L6GQAG2204: {} },
    },
  };
  const flat = flattenServerSoftware(tree);
  assert.strictEqual(flat.length, 2);
  assert.ok(flat.some((f) => f.name === 'L6GQAG2203'));
}

function testGetPendingCommitName() {
  const panel = {
    version1: { active: true, commit: true, valid: true, name: 'v1' },
    version2: { active: true, commit: false, valid: true, name: 'v2' },
  };
  assert.strictEqual(getPendingCommitName(panel, 'v2'), 'v2');
}

function testListFirmwareDirsByCard() {
  const tree = {
    NT: {
      L6GQAG2209: { L6GQAG2209: 'L6GQAG2209' },
      broken: { other: 'x' },
    },
    LT: {
      L6GQAG2210: { L6GQAG2210: 'L6GQAG2210' },
    },
  };
  const ntDirs = listFirmwareDirsByCard(tree, { index: 0, label: 'NT' });
  assert.strictEqual(ntDirs.length, 2);
  assert.ok(ntDirs.find((d) => d.dir === 'L6GQAG2209').valid);
  assert.ok(!ntDirs.find((d) => d.dir === 'broken').valid);

  const ltDirs = listFirmwareDirsByCard(tree, { index: 1, label: 'LT1' });
  assert.strictEqual(ltDirs.length, 1);
  assert.strictEqual(ltDirs[0].label, 'LT/L6GQAG2210');

  const dfDirs = listFirmwareDirsByCard(tree, { index: 0, label: 'DF-16' });
  assert.strictEqual(dfDirs.length, 2);
  assert.strictEqual(getCardFirmwareType({ index: 0 }), 'NT');
}

function testResolveDownloadPath() {
  const nt = resolveDownloadPath('NT', 'L6GQAG2209.421');
  assert.strictEqual(nt.path, 'NT/L6GQAG2209.421');
  assert.strictEqual(nt.name, 'L6GQAG2209.421');

  const lt = resolveDownloadPath('LT', 'L6GQAG2209.421');
  assert.strictEqual(lt.path, 'LT/L6GQAG2209.421');
  assert.strictEqual(lt.name, 'L6GQAG2209.421');
}

function testGetAvailableCards() {
  assert.strictEqual(getAvailableCards({ type: 'DF-16' }).length, 1);
  const mfOlt = { type: 'MF-14', ltCardStatus: [1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0] };
  const cards = getAvailableCards(mfOlt);
  assert.ok(cards.some((c) => c.label === 'NT'));
  assert.ok(cards.some((c) => c.label === 'LT1'));
  assert.ok(cards.some((c) => c.label === 'LT3'));
}

const tests = [
  testGetActiveSoftwareName,
  testGetCommitSoftwareName,
  testGetPendingCommitName,
  testFlattenServerSoftware,
  testListFirmwareDirsByCard,
  testResolveDownloadPath,
  testGetAvailableCards,
];
tests.forEach((fn) => { fn(); console.log('OK', fn.name); });
console.log(`\n${tests.length}/${tests.length} tests passed`);
