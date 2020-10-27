import React, { useEffect } from 'react';
import { useSelector, useDispatch } from 'react-redux';
import { View, ScrollView, StyleSheet, Text } from 'react-native';
import AuctionGroupCards from '../../components/AuctionGroupCard';
import { HeaderButtons, Item } from 'react-navigation-header-buttons';
import HeaderButton from '../../components/HeaderButton';
import * as messages from '../../protobuf-generated-code/command_pb';
import * as socketActions from '../../store/actions/websockets';

const HomeScreen = (props) => {
    const dispatch = useDispatch();

    const token = useSelector(state => state.auth.token);
    const email = useSelector(state => state.auth.userId);
    const uuid = useSelector(state => state.auth.uuid);

    const WS_URL = "ws://192.168.1.22:8282/connect?user=" + email + ";token=" + token;
    console.log("Connecting to " + WS_URL);
    const ws = new WebSocket(WS_URL);


    const fetchPlayerInfo = new messages.FetchPlayerInfoByUserUUIDRequest();
    fetchPlayerInfo.setUserUuid(uuid);

    const auctionRequest = new messages.AuctionRequest();
    auctionRequest.setRequestType(messages.AuctionRequest.RequestType.FETCH_PLAYER_INFO_BY_USER_UUID);
    auctionRequest.setFetchPlayerInfoByUserUuidRequest(fetchPlayerInfo);

    const serializeAuctionRequest = auctionRequest.serializeBinary();

    console.log("Sending msg = " + serializeAuctionRequest);
    
    ws.onopen = function () {
        ws.send(serializeAuctionRequest);
    }
    dispatch(socketActions.sendMsg(serializeAuctionRequest));

    useEffect(() => {
        ws.addEventListener('open', event => {
            console.log("ws onOpen message received", event);
            //this.props.recieveMsg(event.data);
        });

        ws.addEventListener('message', event => {
            console.log("ws message received", event);
            //this.props.recieveMsg(event.data);
        });
    }, [])



    return (
        <View style={styles.container}>
            <ScrollView>
                <AuctionGroupCards name='Blake' />
                <AuctionGroupCards name='FCL' />
            </ScrollView>
        </View>
    );
}

HomeScreen.navigationOptions = navigationData => {
    return {
        headerRight: () =>
            <HeaderButtons HeaderButtonComponent={HeaderButton}>
                <Item
                    title="Favorite"
                    iconName="ios-star"
                    onPress={() => {
                        console.log('Mark as favorite!');
                    }}
                />
            </HeaderButtons>
    }
}

const styles = StyleSheet.create({
    container: {
        flex: 1,
        alignItems: 'center',
        justifyContent: 'center',
    },
});


export default HomeScreen