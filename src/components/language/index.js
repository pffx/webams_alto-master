import React from 'react';
import i18n from 'i18next';
import PropTypes from 'prop-types';
import { useNavigate, useLocation } from "react-router-dom";
import {useDispatch} from 'react-redux'
import SelectItem, { SelectItemButton, SelectListItem } from '@nokia-csf-uxr/ccfk/SelectItem';
import GLOBAL from '../../global'
import { clearAllTabIndex } from '../../actions/global'
import { KEYDOWN, CLICK, ENTER_KEY, SPACE_KEY } from '../../global/keybaord';

const LanguageSelector = (props) => {
  const navigate = useNavigate();
  const location = useLocation();
  const dispatch = useDispatch()
  const getLanguageItem = (value) => {
    let item = undefined
    if(value && GLOBAL.LANGUAGES.length > 0){
      GLOBAL.LANGUAGES.map((ob, x) => {
        if(ob.key === value){
          item = ob
          return
        }
      })
    }
    return item
  }
  const [value, setValue] = React.useState(localStorage.i18nextLng? localStorage.i18nextLng : GLOBAL.LANGUAGES[0].key);
  const [label, setValueLabel] = React.useState(localStorage.i18nextLng ?getLanguageItem(localStorage.i18nextLng).label:GLOBAL.LANGUAGES[0].label);
  const [isOpen, setIsOpen] = React.useState(false);

  const isSelectionKeyPressed = key => key && (key === ENTER_KEY || key === SPACE_KEY);
  const handleEvent = lang => (event) => {
    const { type } = event;
    if (type === KEYDOWN) {
      if (isSelectionKeyPressed(event.key)) {
        setValue(lang.key);
        setValueLabel(lang.label);
        setIsOpen(false);
      }
    } else if (type === CLICK) {
      setValue(lang.key);
      setValueLabel(lang.label);
      setIsOpen(false);
    }
    i18n.changeLanguage(lang.key);
    localStorage.setItem('i18nextLng', lang.key);
    //navigate(location.pathname,{ replace: true });
    navigate("/")
    dispatch(clearAllTabIndex())

  };
  const renderSelectItemValues = props => (
    <SelectItemButton placeholder="Select an item" inputProps={{ value: value || 'Select an item' }} {...props}>
      {value && <SelectListItem>{label}</SelectListItem>}
    </SelectItemButton>
  );
  return (
    <SelectItem
      data-id="language-selector"
      renderSelectItemBase={renderSelectItemValues}
      variant="underlined"
      value={value}
      aria-label={`Language selector, currently selected: ${value}`}
      isOpen={isOpen}
      onOpen={() => {
        setIsOpen(true);
      }}
      onClose={() => {
        setIsOpen(false);
      }}
      style={{ width: 100, display: 'inline-flex', boxShadow: 'none',background: props.light? GLOBAL.COLOR.Background : "" }}
    >
      {GLOBAL.LANGUAGES.map(lang => (
        <SelectListItem key={lang.key} selected={value == lang.key} onClick={handleEvent(lang)} onKeyDown={handleEvent(lang)}>
          {lang.label}
        </SelectListItem>
      ))}
    </SelectItem>
  );
};
LanguageSelector.propTypes = {
  light:PropTypes.bool,
};

LanguageSelector.defaultProps = {
  light: false,
}

export default LanguageSelector