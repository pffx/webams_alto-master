import PropTypes from 'prop-types';
import NokiaCheckbox, { CheckboxLabelContent } from '@nokia-csf-uxr/ccfk/Checkbox';
import Label from '@nokia-csf-uxr/ccfk/Label';
// import { SPACING_24 } from '@nokia-csf-uxr/freeform-design-tokens/tokens/spacing';
import GLOBAL, {SPACING } from "../../../global"
const Checkbox = (props) => {
    const {
        onChange,
        disabled,
        title,
        ariaLabel,
        style,
        checked,
    } = props;
    return(
        <Label style={style} >
          <NokiaCheckbox
              checked={checked}
              onChange={() => {
                onChange(!checked);
            }}
              inputProps={{ 'aria-label': ariaLabel }}
              disabled={disabled}

          />
          <CheckboxLabelContent style={{ marginBottom: 0 }}>{title}</CheckboxLabelContent>
        </Label>
    )
}
Checkbox.propTypes = {
    //the callback of onChange function
    onChange:PropTypes.func.isRequired,
    //the title
    title:PropTypes.string.isRequired,
    checked: PropTypes.bool,
    indeterminate: PropTypes.bool,//??? dirty to add
    disabled: PropTypes.bool,
    ariaLabel:PropTypes.string,
    style:PropTypes.object,
};

Checkbox.defaultProps = {
    disabled: false,
    checked:false,
    indeterminate: false,
    ariaLabel: "checkbox",
    style:{ margin: `1.125rem 0 ${SPACING.SPACING_24} 0` },
}

export default Checkbox;