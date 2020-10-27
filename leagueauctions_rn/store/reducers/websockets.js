import { MESSAGE_SENT, MESSAGE_RECEIVED } from '../../constants/socket';

INITIAL_STATE = {
    receivedMsg: "None"
}
 export default (state = INITIAL_STATE, action) => {
    switch (action.type) {
        case MESSAGE_SENT:
            return { ...state, receivedMsg: action.payload }
        case MESSAGE_RECEIVED:
            return { ...state, receivedMsg: action.payload }
        default:
            return state;
    }
}