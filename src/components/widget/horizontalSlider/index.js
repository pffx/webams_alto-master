import React, { useState } from 'react';
import PropTypes from 'prop-types';
import AddIcon from '@nokia-csf-uxr/ccfk-assets/AddIcon';
import RemoveIcon from '@nokia-csf-uxr/ccfk-assets/RemoveIcon';
import IconButton from '@nokia-csf-uxr/ccfk/IconButton';
import Slider, { SliderLabel } from '@nokia-csf-uxr/ccfk/Slider';
import Label from '@nokia-csf-uxr/ccfk/Label';
import { TextInputLabelContent } from '@nokia-csf-uxr/ccfk/TextInput';
import GLOBAL, {SPACING} from "../../../global"

const HorizontalSlider = (props) => {
  const {
    style,
    max,
    min,
    current,
    horizonSliderChange,
  } = props;
  const [currentValue, setCurrentValue] = useState(current);

  const handleChange = (data) => {
    // console.log('Slider onChange: ', data);
    setCurrentValue(data.value);
    horizonSliderChange(data.value)
  };

  const handleLabelClick = (value) => {
    let normalizedValue = value;
    if (normalizedValue < min ) normalizedValue = min;
    if (normalizedValue > max) normalizedValue = max;
    handleChange({ value: normalizedValue })
  }
  const frontIcon = (<IconButton aria-label="Increase value" onClick={() => handleLabelClick(currentValue + 1)}><AddIcon /></IconButton>);
  const rearIcon = (<IconButton aria-label='Decrease value' onClick={() => handleLabelClick(currentValue - 1)}><RemoveIcon /></IconButton>);
  
  const renderMinLabel = (labelProps) => (
    <SliderLabel label={props.type === 'icon' ? rearIcon : 'Minus'} {...labelProps} />
  );
  const renderMaxLabel = (labelProps) => (
    <SliderLabel label={props.type === 'icon' ? frontIcon : 'Add'} {...labelProps} />
  );

  const getAriaLabel = () => {
    if (currentValue === min) {
      return `Reach minimum value ${min}`;
    } else if( currentValue === max) {
      return `Reach maximum value ${max}`;
    } 
    return currentValue;
  }

  return (
    <div style={style}>
      <Label
        id='rangeselection-label-id'
        verticalLayout
        maxWidth={true}
        style={{
          marginBottom: props.type === "icon" ? SPACING.SPACING_12 : SPACING.SPACING_16,
        }}
      >
        <TextInputLabelContent>{props.label}</TextInputLabelContent>
      </Label>
      <Slider
        value={currentValue}
        thumbProps={{
          minThumbProps: {
            role: "slider",
            "aria-valuemin": min,
            "aria-valuenow": currentValue,
            "aria-valuemax": max,
            "aria-valuetext": getAriaLabel(),
            "aria-labelledby": 'rangeselection-label-id',
            'aria-orientation': 'horizontal',
          },
        }}
        inputProps={{ value: currentValue }}
        onChange={handleChange}
        renderMinLabel={renderMinLabel}
        renderMaxLabel={renderMaxLabel}
        {...props}
      />
    </div>
  );
};

HorizontalSlider.propTypes = {
  style: PropTypes.object,
  horizonSliderChange:PropTypes.func.isRequired,
  max: PropTypes.number,
  min: PropTypes.number,
  current: PropTypes.number,
};

HorizontalSlider.defaultProps = {
  style:{},
  max: 10,
  min: 1,
  current: 5,
}
export default HorizontalSlider;