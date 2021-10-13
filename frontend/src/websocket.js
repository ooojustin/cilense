import vars from './variables';

let ws = null;

export const send = data => {
    const params = JSON.stringify(data);
    ws.send(params);
}

export const initWebSocket = (onopen, onmessage) => {
    
    const url = "ws" + vars.backend.slice(4) + "ws";
    ws = new WebSocket(url);

    ws.onmessage = onmessage;
    ws.onopen = onopen;

}

export const joinRoom = room_id => {
    send({
        action: "join_room",
        data: { room_id }
    });
}

export const sendMessage = message => {
    send({
        action: "send_message",
        data: { message }
    });
}

export const authenticate = token => {
    send({
        action: "authenticate",
        data: { token }
    });
}
