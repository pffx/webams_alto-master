import PropTypes from 'prop-types';
import NokiaButton, { ButtonText } from '@nokia-csf-uxr/ccfk/Button';
import {NOKIA_BLUE_500,} from '@nokia-csf-uxr/freeform-design-tokens/tokens/colors';

const Button = ({
        onClick,
        disabled,
        variant,
        title,
        ariaLabel,
        backgroundColor,
        style,
        fullWidth,
    }) => {

    return (
        <div style={{...style,}}>
            <NokiaButton
                variant={variant}
                aria-label={ariaLabel}
                onClick={() => {
                    onClick();
                }}
                //change the button to nokia blue
                style={{backgroundColor:backgroundColor,border:"1px",borderRadius:"8px"}}
                fullWidth={fullWidth}
                disabled={disabled}
            >
                <ButtonText>{title}</ButtonText>
            </NokiaButton>
        </div>
    );
}
Button.propTypes = {
    //the callback of click function
    onClick:PropTypes.func.isRequired,
    //the title
    title:PropTypes.string.isRequired,
    disabled: PropTypes.bool,
    variant: PropTypes.string,
    ariaLabel:PropTypes.string,
    style:PropTypes.object,
    backgroundColor:PropTypes.string,
    fullWidth:PropTypes.bool,
};

Button.defaultProps = {
    disabled: false,
    variant: 'call-to-action',
    ariaLabel: "button",
    backgroundColor:NOKIA_BLUE_500,
    fullWidth: false,
    style:{ width: "100%",display:"flex",justifyContent:"center",alignItems:"center", marginTop:"0.5rem",marginBottom:"0.5rem"},
}

export default Button;