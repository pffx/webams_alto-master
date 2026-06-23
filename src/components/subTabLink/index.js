import {useDispatch} from 'react-redux'
import { useTranslation } from 'react-i18next';
//import { useNavigate, useLocation } from "react-router-dom";
import { changeSubTabIndex } from '../../actions/global'
import { CheckT } from '../../actions/login'
import { NavLink } from 'react-router-dom';

const SubTabLink = (props) => {
  //const navigate = useNavigate();
  const dispatch = useDispatch()
  // const location = useLocation();// 获取上一个页面传递进来的 state 参数
  const { t } = useTranslation();
  const parent = props.parent === "" || props.parent === undefined? "" : props.parent
  const path = '/' + parent + "/" + props.tab_id

  const handleClick = value => (action) => {
    //navigate('/' + parent + "/" + value,{ replace: true })
    dispatch(changeSubTabIndex({index:value}))
    dispatch(CheckT())
  };
  return (
    <li className="nav-item">
      <NavLink onClick={handleClick(props.tab_id)} className={props.selected ? "nav-link active" : "nav-link"} to={path} style={{ textDecoration: 'none' }} >{t(props.children)}</NavLink>
    </li>
  );
};
export default SubTabLink
