import React from 'react';
import PropTypes from 'prop-types';
import ToggleSwitch, { ToggleSwitchLabelContent } from '@nokia-csf-uxr/ccfk/ToggleSwitch';
import Label from '@nokia-csf-uxr/ccfk/Label';
// import { SPACING_24 } from '@nokia-csf-uxr/freeform-design-tokens/tokens/spacing';
import GLOBAL, {SPACING } from "../../../global"
const SwitchButton = (props) => {
    const {
        onChange,
        disabled,
        title,
        ariaLabel,
        style,
        checked,
        stateColor,
    } = props;
    return(
        <Label style={style}>
            <ToggleSwitch
                // {...otherProps}
                aria-checked={checked}
                ariaLabel={ariaLabel}
                checked={checked}
                stateColor={stateColor}
                onChange={() => {
                    onChange(!checked);
                    // onChange();
                }}
                disabled={disabled}
            />
            <ToggleSwitchLabelContent  style={{ marginBottom: 0 }}>{title}</ToggleSwitchLabelContent>
        </Label>
    )
}
SwitchButton.propTypes = {
    //the callback of onChange function
    onChange:PropTypes.func.isRequired,
    //the title
    title:PropTypes.string.isRequired,
    checked: PropTypes.bool,
    indeterminate: PropTypes.bool,//??? dirty to add
    disabled: PropTypes.bool,
    stateColor:PropTypes.bool,
    ariaLabel:PropTypes.string,
    style:PropTypes.object,
};

SwitchButton.defaultProps = {
    disabled: false,
    checked:false,
    indeterminate: false,
    stateColor:false,
    ariaLabel: "SwitchButton",
    style:{ margin: `1.125rem 0 ${SPACING.SPACING_24} 0` },
}

export default SwitchButton;