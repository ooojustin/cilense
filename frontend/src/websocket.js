import vars from './variables';

let ws = null;

export const send = data => {
    console.log("send", data);
    const params = JSON.stringify(data);
    ws.send(params);
}

export const initWebSocket = (onopen, onmessage) => {
    
    const url = "ws" + vars.backend.slice(4) + "ws";
    ws = new WebSocket(url);

    ws.onopen = onopen;
    ws.onmessage = msg => {
        // convert message to object before passing to handler
        const datastr = unescape(msg.data);
        const data = JSON.parse(datastr);
        data && onmessage(data);
    };

}

export const restoreSession = session_id => {
    send({
        action: "restore_session",
        data: { session_id }
    });
}

export const joinRoom = (room_id, password) => {
    send({
        action: "join_room",
        data: { room_id, password }
    });
}

export const sendMessage = message => {
    send({
        action: "send_message",
        data: { message }
    });
}
