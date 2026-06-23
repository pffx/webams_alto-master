import i18n from 'i18next';
import { initReactI18next } from 'react-i18next'; 
import GLOBAL from '../global';
import { en_US } from './en_US';
import { zh_CN } from './zh_CN';

 
const resources = {
    en: {
        translation: en_US
    },
    cn: {
        translation: zh_CN
    },
}
 
i18n
    // .use(Backend)
    // .use(LanguageDetector)
    .use(initReactI18next) 
    .init({
        resources,
        lng: localStorage.i18nextLng ? localStorage.i18nextLng : GLOBAL.LANGUAGES[0].key,
        fallbackLng: localStorage.i18nextLng ? localStorage.i18nextLng : GLOBAL.LANGUAGES[0].key,
        interpolation: {
            escapeValue: false, 
        }
    });
 
 
export default i18n;