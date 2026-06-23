import React from 'react';
import { useTranslation } from 'react-i18next';
import {useSelector,useDispatch} from 'react-redux'
import Hyperlink from '@nokia-csf-uxr/ccfk/Hyperlink';
import { SPACING_12, SPACING_16, SPACING_32} from '@nokia-csf-uxr/freeform-design-tokens/tokens/spacing';


const Signature = (props)=>{
  const { t } = useTranslation();
  const { variant, ...otherProps } = props;

  const dispatch = useDispatch()
  const toggleRef = React.useRef(null); 


  return (
    <div style={{display:"flex",flexDirection: "row", justifyContent:"end", width: "93%",marginLeft:"1rem",marginRight:"1rem",marginTop: SPACING_16}} >
        <Hyperlink target="_blank" href="http://www.nokia.com" >
            Nokia
        </Hyperlink>
        <div style={{width:"1rem"}} ></div>
        <Hyperlink target="_blank" href="https://confluence.ext.net.nokia.com/pages/viewpage.action?spaceKey=sdannpi&title=Alto" >
            Powered by APAC NPI
        </Hyperlink>
        <div style={{width:"1rem"}} ></div>
        <Hyperlink href="mailto:ni-apac-npi-tools@LIST.NOKIA.COM" >
            E-Mail: ni-apac-npi-tools@LIST.NOKIA.COM
        </Hyperlink>
        {/* <span>
          <a href="alden.zhou@nokia.com">E-Mail: alden.zhou@nokia.com</a>
        </span> */}
    </div>
    );

}
export default Signature