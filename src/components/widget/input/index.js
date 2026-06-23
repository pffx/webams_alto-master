import PropTypes from 'prop-types';
import Label from '@nokia-csf-uxr/ccfk/Label';
import TextInput, { TextInputLabelContent } from '@nokia-csf-uxr/ccfk/TextInput';
const Input = (props) => {
    const {
        onChange,
        disabled,
        variant,
        value,
        title,
        style,
        maxWidth,
        required,
        placeholder,
    } = props;

    return (
        <div style={style}>
            <Label maxWidth={maxWidth} >
              <TextInputLabelContent required={required} >{title}</TextInputLabelContent>
            </Label>
            <TextInput
              variant={variant}
              maxWidth={maxWidth}
              type="text"
              disabled={disabled}
              value={value}
              onChange={(event) => {onChange(event)}}
              inputProps={{ autoComplete: 'on' }}
              placeholder={placeholder}
            />
        </div>
    );
}
Input.propTypes = {
    //the callback of click function
    onChange:PropTypes.func.isRequired,
    value:PropTypes.string.isRequired,
    //the title
    title:PropTypes.string.isRequired,
    disabled: PropTypes.bool,
    placeholder:PropTypes.string,
    required: PropTypes.bool,
    variant: PropTypes.string,
    style:PropTypes.object,
    maxWidth:PropTypes.bool,
};

Input.defaultProps = {
    disabled: false,
    variant: 'outlined',
    maxWidth: false,
    required:false,
    placeholder:"",
    style:{ },
}

export default Input;