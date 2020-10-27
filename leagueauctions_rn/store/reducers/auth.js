import { ActionSheetIOS } from 'react-native';
import { LOGIN, SIGNUP } from '../actions/auth';

const initialState = {
  token: null,
  userId: null,
  uuid: null
};

export default (state = initialState, action) => {
  switch (action.type) {
    case LOGIN:
      return {
        token: action.token,
        userId: action.userId,
        uuid: action.uuid
      };
    case SIGNUP:
      return {
        token: action.token,
        userId: action.userId,
        uuid: null
      };
    default:
      return state;
  }
};