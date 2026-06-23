import React, { useState, useRef,useEffect } from 'react';
import PropTypes from 'prop-types';
import _isEmpty from 'lodash/isEmpty'
import _toLower from 'lodash/toLower';
import _replace from 'lodash/replace';
import _findIndex from 'lodash/findIndex';
import Label from '@nokia-csf-uxr/ccfk/Label';
import SelectItem, {
  SelectItemLabelContent,
  SelectListItem,
  SelectItemInput,
  SelectItemText,
  SelectItemButton,
  SelectItemClearButton,
  SelectGroupHeading,
  SelectItemBaseText,
} from '@nokia-csf-uxr/ccfk/SelectItem';
import {ARROW_DOWN,
    ENTER_KEY,
    SPACE_KEY,
    ESCAPE,
    TAB,
} from '../../../global/keybaord';

const Select = (props) => {
  const {
    onChange,
    dataItems,
    selectedItem,
    disabled,
    required, 
    variant,
    hasClearButton,
    error,
    errorMessage,
    searchable,
    caseInsentiveMatching,
    placeholder,
    maxWidth,
    stickyHeader,
    title,
    style,
  } = props;

//  const title = props.children? props.children:undefined

  const TEXT_TRUNCATE_STYLE = {
    whiteSpace: 'nowrap',
    overflow: 'hidden',
    textOverflow: 'ellipsis',
  };

  const HIGHLIGHT_MARK_START = '<span style="font-family:NokiaPureBold">';
  const HIGHLIGHT_MARK_END ='</span>';
  const MARK_LENGTH = 47; // length of marking strings

  const getInputItem = (value) => {
    const item = {key:"",label:""}
    //number 0 will be treated as true in _isEmpty
    if(_isEmpty(value) && value!==0){
      return item
    }
    // if(dataItems.length > 0){
    //   dataItems.map((ob, x) => {
    //     if(ob.value == value){
    //       item = ob
    //       return
    //     }
    //   })
    // }
    const foundItem = dataItems.find((ob) => ob.value === value);
    if (foundItem) {
      return foundItem;
    }
  
    return item;
  }
  //number 0 will be judged as ture in _isEmpty
  const [selectValue, setSelectValue] = useState((_isEmpty(selectedItem) && selectedItem!==0) ? undefined : selectedItem);// the value of selected option
  //const [selected, setSelected] = useState(selectedItem?selectedItem.label:'');// the index of selected option
  const [isOpen, setIsOpen] = useState(false);

  ///////////////// SEARCHABLE //////////////////////////
  const [newSearch, setNewSearch] = useState(false);
  const [shouldMarkItems, setShouldMarkItems] = useState(false);
  const [selectItems, setSelectItems] = useState(dataItems);//all the options
  const [selecting, setSelecting] = useState(false);
  const [inputText, setInputText] = useState(_isEmpty(selectedItem) ? "" : getInputItem(selectedItem).label);//the label of selected option
  const selectItemRef = useRef(null);
  const selectInputRef = useRef(null);
  const isSelectionKeyPressed = key => key && (key === ENTER_KEY || key === SPACE_KEY);

  useEffect(() => {
    if(!_isEmpty(selectedItem) || selectedItem === 0){
      setInputText(getInputItem(selectedItem).label)
      setSelectValue(selectedItem)
    }else{
      setSelectValue(undefined)
      setInputText("")
    }
  }, [selectedItem]);

  const matchHeadersAndItems = (data, filterValue, caseInsentiveMatching) => {
    let idx;
    let header;
    let usedHeader = false;
    const filteredValues = [];
    for (idx in data) {
      const item = data[idx];
      if (item.isHeader) {
        header = item;
        usedHeader = false;
      } else {
        const isValueMatched = caseInsentiveMatching ? isMatched(item.label, filterValue) : isMatchedExactly(item.label, filterValue)
        if (isValueMatched) {
          if (!usedHeader) {
            usedHeader = true;
            filteredValues.push(header);
          }
          filteredValues.push(item);
        }
      }
    }
    return filteredValues;
  }

  // case insensitive matching all occurrences
  const isMatched = (item, searchPattern) => {
    const itemIsIncluded = _toLower(item).indexOf(_toLower(searchPattern)) >= 0;
    return _isEmpty(searchPattern) || itemIsIncluded;
  };
  const markItem = (item, searchPattern, truncateListText=true) => {
    if (!item) {
      return null;
    }
    if (!searchPattern || searchPattern.length === 0) {
      return item;
    }
    let endingIndex = 0;
    const searchPatternLength = searchPattern.length;
    let markedText = item;
    while (true) {
      const startIndex = _toLower(markedText).indexOf(_toLower(searchPattern), endingIndex);
      if (startIndex === -1) {
         break;
      }
      const actualText = markedText.substr(startIndex, searchPatternLength);
      const replacementText = `${HIGHLIGHT_MARK_START}${actualText}${HIGHLIGHT_MARK_END}`;
      const beginningText = startIndex > 0 ?  markedText.slice(0,startIndex) : '';
      const endingText = markedText.slice(startIndex + searchPatternLength);
      markedText = beginningText + replacementText + endingText;
      endingIndex = startIndex + searchPatternLength + MARK_LENGTH;
    }
    return <div style={truncateListText ? TEXT_TRUNCATE_STYLE : {}} dangerouslySetInnerHTML={{ __html: markedText }}></div>;
  };
  // exact matching
  const isMatchedExactly = (item, searchPattern) => {
    const itemIsIncluded = item && item.indexOf(searchPattern) >= 0;
    return _isEmpty(searchPattern) || itemIsIncluded;
  };
  const markItemExactly = (item, searchPattern, truncateListText=true) => {
    if (!item) {
      return null;
    }
    const markedText = _replace(item, new RegExp(searchPattern, 'g'), `${HIGHLIGHT_MARK_START}${searchPattern}${HIGHLIGHT_MARK_END}`);
    return <div style={truncateListText ? TEXT_TRUNCATE_STYLE : {}} dangerouslySetInnerHTML={{ __html: markedText }}></div>;
  };

  const stopEvents = (event) => { 
    event.preventDefault();
    event.stopPropagation();
  };

  const haveSelectedItems = () => selectItems && selectItems.length > 0;
  
  const resetSelectItems = () => {
    setSelectItems(dataItems);
  };

  const updateSelection = (item) => {
    console.log("updateSelection   item:= ",item)
    // remote it temporaty
    //setSelected(item);
    setInputText(item);
    setIsOpen(false);
    resetSelectItems();
    setSelecting(false);
  };

  const shouldUpdateSelection = (event) => {
    // no action now, if find it is useful, add it back
    const { key, type} = event;
    console.log("shouldUpdateSelection   event:  ",event)

    // if ESCAPE or TAB or click (outside) event closed the menu keep current value of `selected`
    if ((key === ESCAPE) || (key === TAB) || (type === 'mousedown')) {
      //updateSelection(selected);
    }
  }

  // on SelectListItem, handle Menu item selection and close dropdown Menu after item is selected
  const handleSearchableEvent = (item) => (event) => {
    const { type, key } = event;
    switch (type) {
      case 'keydown':
        if (isSelectionKeyPressed(event.key)) {
          updateSelection(item);
        }
      break;
      case 'click':
        updateSelection(item);
      break;
      default:
    }
  };

  const handleInputChange = (event) => {
    let newText = event.target.value;
    if (newSearch) {
      // remove the unnecessary text
      newText = newText.replace(inputText,'');
    }
    setNewSearch(false);
    if (newText === ' ') {
      // space entered to open the List
      resetSelectItems();
      newText = '';
    } else {
      // filter the list of values
      setSelectItems(matchHeadersAndItems(dataItems, newText, caseInsentiveMatching));
      setShouldMarkItems(true);
    }

    setInputText(newText);
    setIsOpen(true);
  };

  const handleInputSubmit = (event) => {
    if (!haveSelectedItems()) {
      // no item to select, leave menu open
      setIsOpen(true);
      return;
    }
    //What's the usage?
    /*if (inputText === '') {
      // if inputText is empty, keep latest selected value
      setInputText(selected);
      setSelected(selected);
    } else if (inputText !== selected) {
      // select the 1st item from the filtered list
      // selectItems[0] is always a header for this dataset
      setInputText(selectItems[1].label);
      setSelected(selectItems[1].label);
    }*/
    setIsOpen(false);
    resetSelectItems();
  }

  const handleInputKeyDown = (event) => {
    const { key } = event;
    console.log("handleInputKeyDown   inputText= ",inputText)
    if (key === ARROW_DOWN) {
      setSelecting(true);
    } else if (key === TAB) {
      // prevent tabbing off SelectItem when not done selecting (e.g. no matching items in search)
      if (!haveSelectedItems()) {
        stopEvents(event);
      } else {
        const isTypedTextInData = _findIndex(dataItems, item => item.label === inputText) > 0;
        // prevent tabbing when the search text entered is not in the list
        if (!isTypedTextInData && inputText.length > 0) {
          stopEvents(event);
        }
        // delete this part temparory, do not know it's usage
        /* else {
          // if user pressed TAB w/o entering data, keep the original selected data
          if (inputText.length > 0) {
            setSelected(inputText);
          } else {
            setInputText(selected);
          }
          setIsOpen(false);
          resetSelectItems();
        }*/
      }
    }
  };

  const renderClearButton1 = (props) => {
    return (
      <SelectItemClearButton
        aria-label="clear input"
        onClick={() => { updateSelection(''); selectItemRef.current.selectItemRef.focus();}}
        {...props}
      />
    );
  };

  const renderSelectItemBase = (props) => {
    // Always display ClearButton when searchable when an item is selected 
    console.log("renderSelectItemBase  inputText= ",inputText)
    const clearButton = hasClearButton && !disabled && (inputText !== '' || selecting) ?  renderClearButton1 : undefined;
    return (
      <SelectItemInput
        id="selectitem-component-input"
        autoComplete="off"
        ref={selectInputRef}
        placeholder={inputText ? undefined : placeholder}
        onChange={handleInputChange}
        onSubmit={handleInputSubmit}
        onFocus={() => { setNewSearch(true); }}
        onKeyDown={handleInputKeyDown}
        value={inputText}
        {...props}
        renderClearButton={clearButton}
      />
    )
  };

   ///////////////// NON SEARCHABLE //////////////////////////
  // handle list item selection and close dropdown list after item is selected
  const handleEvent = (index) => (event) => {
    const { type, key } = event;
    // console.log("hadle event   index=",index)
    switch (type) {
      case 'keydown':
        if (isSelectionKeyPressed(key)) {
          //setSelected(index);
          setSelectValue(dataItems[index].value);
          setInputText(dataItems[index].label)
          onChange(dataItems[index])
          setIsOpen(false);
        }
      break;
      case 'click':
        //setSelected(index);
        setSelectValue(dataItems[index].value);
        setInputText(dataItems[index].label)
        onChange(dataItems[index])
        setIsOpen(false);
      break;
      default:
    }
  };

  const renderClearButton = (props) => {
    return (
      <SelectItemClearButton
        aria-label="clear input"
        onClick={() => { setSelectValue(undefined); /*setSelected(undefined);*/ selectItemRef.current.selectItemRef.focus();}}
        {...props}
      />
    );
  };

  // Render the necessary selected values and desired dropdown icon into the
  // SelectItem base area.
  const renderSelectItemValues = (props)  => {
    // Only display ClearButton when not searchable when an item is selected and the dropdown is open
    const clearButton = hasClearButton && !disabled && selectValue && isOpen ?  renderClearButton : undefined;
    //console.log("renderSelectItemValues    inputText: ",inputText)
    //console.log("renderSelectItemValues    selected: ",selected)
    return (
      <SelectItemButton
        placeholder={placeholder}
        inputProps={{ value: selectValue ? selectValue : placeholder }}
        renderClearButton={clearButton}
        {...props}
      >
        {inputText && <SelectItemBaseText>{inputText}</SelectItemBaseText>}
      </SelectItemButton>
    )
  };

  const nonSearchableSelectItem = (
    <>
      <Label id="selectitem-component-label" verticalLayout>
        <SelectItemLabelContent required={required}>{title}</SelectItemLabelContent>
      </Label>
      <SelectItem
        aria-labelledby="selectitem-component-label"
        ref={selectItemRef}
        aria-label={selectValue !== '' ? undefined : placeholder}
        renderSelectItemBase={renderSelectItemValues}
        variant={variant}
        disabled={disabled}
        isOpen={isOpen}
        onOpen={() => { setIsOpen(true); }}
        onClose={() => { setIsOpen(false); }}
        error={error}
        errorMessage={errorMessage}
        maxWidth={maxWidth}
        listProps={{
          stickyHeader,
        }}
      >
      {dataItems.length > 0 && dataItems.map((item, x) => {
        const selectitem = item.isHeader ?
          <SelectGroupHeading key={`head-${item.value}`}>{item.label}</SelectGroupHeading>
        :
          <SelectListItem
            key={`${item.value}-${x}`}
            selected={selectValue === item.value}
            onClick={handleEvent(x)}
            onKeyDown={handleEvent(x)}
          >
            {item.label}
          </SelectListItem>
          ;
          return selectitem;
      })}
      </SelectItem>
    </>
  );

  const searchableSelectItem = (
    <Label htmlFor="selectitem-component-input" verticalLayout>
      {title&&<SelectItemLabelContent required={required}>{title}</SelectItemLabelContent>}
      <SelectItem
        ref={selectItemRef}
        aria-label={inputText ? undefined : placeholder}
        renderSelectItemBase={renderSelectItemBase}
        variant={variant}
        disabled={disabled}
        searchable
        isOpen={isOpen}
        onOpen={() => { setIsOpen(true); setShouldMarkItems(false); }}
        onClose={(event) => { if (haveSelectedItems()) {setIsOpen(false)} ; shouldUpdateSelection(event); }}
        error={error}
        errorMessage={errorMessage}
        maxWidth={maxWidth}
      >
      {haveSelectedItems() && selectItems.map((item, x) => {
        const selectitem = item.isHeader ?
        <SelectGroupHeading key={`head-${item.value}`}>{item.label}</SelectGroupHeading>
        :
        <SelectListItem
            key={`${item.value}-${x}`}
            selected={selectValue === item.value}
            searchSelection={shouldMarkItems && !selecting && (x === 0)}
            disabled={item.disabled}
            onClick={handleSearchableEvent(item.value)}
            onKeyDown={handleSearchableEvent(item.value)}
          >
              {shouldMarkItems && caseInsentiveMatching && <SelectItemText>{markItem(item.label, inputText)}</SelectItemText>}
              {shouldMarkItems && !caseInsentiveMatching && <SelectItemText>{markItemExactly(item.label, inputText)}</SelectItemText>}
              {!shouldMarkItems && item.label}
          </SelectListItem>
          ;
          return selectitem;
      })}
      {!haveSelectedItems() &&  <SelectListItem disabled>No matching items</SelectListItem>}
      </SelectItem>
    </Label>
  );

  return (
    <div style={style}>
      {!searchable && nonSearchableSelectItem}
      {searchable && searchableSelectItem}
    </div>
  );
};

  Select.propTypes = {
    //the callback of changed function
    onChange:PropTypes.func.isRequired,
    //the selected value
    selectedItem:PropTypes.oneOfType([ 
      PropTypes.string,
      PropTypes.number
    ]),
    // the previous data, the base form like this
    //const data = [
    //   {value:"card",label: 'Card', isHeader: true },
    //   {value:"nt",label:"NT"},
    //   {value:"lt",label:"LT"},
    //   {value:"shelf",label:"Shelf"},
    // ]
    dataItems:PropTypes.array.isRequired,
    disabled: PropTypes.bool,
    // show this input is must option or not
    required: PropTypes.bool,
    //oneOf 'outlined' | 'underlined'
    variant: PropTypes.string,
    hasClearButton: PropTypes.bool,
    searchable: PropTypes.bool,
    error: PropTypes.bool,
    errorMessage: PropTypes.string,
    caseInsentiveMatching: PropTypes.bool,
    placeholder: PropTypes.string,
    maxWidth: PropTypes.bool,
    stickyHeader: PropTypes.bool,
    style: PropTypes.object,
  };
  
  Select.defaultProps = {
    selectedItem:undefined,
    dataItems:[],
    disabled: false,
    required: false,
    variant: 'outlined',
    hasClearButton: false,
    searchable: false,
    error: false,
    errorMessage: undefined,
    caseInsentiveMatching: true,
    placeholder: 'Select',
    maxWidth: true,
    stickyHeader: false,
    style:{},
  }
  
export default Select;
