import { createStore, combineReducers, applyMiddleware } from 'redux'
import { persistStore, persistReducer } from 'redux-persist'
import storage from 'redux-persist/lib/storage'
import hardSet from 'redux-persist/lib/stateReconciler/hardSet'
//import { configureStore } from '@reduxjs/toolkit'
import {createLogger} from 'redux-logger';
import thunkMiddleware from 'redux-thunk';
import LoginReducer from '../reducers/login'
import FeatureListReducer from '../reducers/feature'
import GlobalReducer from '../reducers/global'

const persistConfig = {
  key: 'root',
  storage,
  //blacklist: ["page404"]
  stateReconciler: hardSet,
}
const rootReducer = combineReducers({
    LoginReducer,
    FeatureListReducer,
    GlobalReducer,
})
const myPersistReducer = persistReducer(persistConfig, rootReducer)

const loggerMiddleware = createLogger()
const store = createStore(
  myPersistReducer,
  applyMiddleware(
    thunkMiddleware,
    loggerMiddleware
  )
)

const persistor = persistStore(store)

// const store = configureStore({
//   reducer: {
//     login: LoginReducer,
//   },
//   applyMiddleware:{loggerMiddleware}
// })

export default store
export { persistor }