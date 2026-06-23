import {useDispatch} from 'react-redux'
import { useNavigate, useLocation } from "react-router-dom";
import { useTranslation } from 'react-i18next';
import HorizontalNavigation, { Tab, HorizontalNavigationLabel } from '@nokia-csf-uxr/ccfk/HorizontalNavigation';
import { changeTabIndex, changeSubTabIndex } from '../../actions/global'
import { CheckT } from '../../actions/login'


const TabLink = (props) => {
  const navigate = useNavigate();
  const location = useLocation();// 获取上一个页面传递进来的 state 参数
  const dispatch = useDispatch()
  const { t } = useTranslation();
  //console.log("TabLink   props.children:",props.children)

  const handleClick = value => (action) => {
    navigate('/' + value,{ replace: false })
    dispatch(changeTabIndex({index:value}))
    if(value === "maintenance"){
      dispatch(changeSubTabIndex({index:"backup_restore"}))
    }else if (value === "alarm_log"){
      dispatch(changeSubTabIndex({index:"alarm"}))
    }else if (value === "system"){
      dispatch(changeSubTabIndex({index:"management"}))
    }else{
      dispatch(changeSubTabIndex({index:""}))
    }
    dispatch(CheckT())
  };
  return (
    <Tab
      selected={props.selected}
      onSelect={handleClick(props.tab_id)} 
      // aria-label={this.state.selected === 0 ? `${TAB_LABELS[0]}` : `${ARIA_LABEL_TEXT_1} ${TAB_LABELS[0]} ${ARIA_LABEL_TEXT_2}`}
      role="option"
      >
      <HorizontalNavigationLabel>{t(props.label)}</HorizontalNavigationLabel>
    </Tab>
  );
};
export default TabLink