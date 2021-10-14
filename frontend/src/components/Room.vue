<template>
    <div class="flex items-center justify-center h-screen">
        <div class="chat-container bg-gray-700 p-5 rounded-md border border-gray-500" style="height: 60%;">
            <Message text="test message" />
        </div>
    </div>
</template>

<script>
import vars from "../variables";
import Message from "./Message";

import { 
    initWebSocket,
    restoreSession,
    joinRoom,
    sendMessage,
    authenticate
} from '../websocket';

export default {
    name: 'Room',
    components: {
        Message
    },
    methods: {
        handleMessage(msg) {
            console.log("receive", msg);

            // handle messages from the websocket
            // check if response has a 'type', for processing
            let type = null;
            if (Object.prototype.hasOwnProperty.call(msg, "type"))
                type = msg.type;

            // skip message if no type is specified
            if (!type)
                return;

            switch (type) {
                case "room_joined": {
                    // store session inside local storage
                    const sessions = JSON.parse(localStorage.getItem("sessions")) || {};
                    sessions[this.room.id] = msg.session_id;
                    localStorage.setItem("sessions", JSON.stringify(sessions));
                    break;
                }
            }

        }
    },
    data() {
        return {
            room: null,
            token: null,
            session: null
        };
    },
    created() {


        // load room data into local state
        const { id } = this.$route.params;
        vars.api.get("/room/" + id)
        .then(r => {
            const { data } = r.data;
            this.room = data;
        });

        // retrieve room token from local storage
        const tokens = JSON.parse(localStorage.getItem("tokens")) || {};
        if (Object.prototype.hasOwnProperty.call(tokens, id))
            this.token = tokens[id];

        // retrieve stored session from local storage
        const sessions = JSON.parse(localStorage.getItem("sessions")) || {};
        if (Object.prototype.hasOwnProperty.call(sessions, id))
            this.session = sessions[id];

        // func will when websocket connect is opened
        const onOpen = () => {
            console.log("Connected to websocket.");
            this.session && restoreSession(this.session);
            this.session || joinRoom(id);
            this.token && setTimeout(() => authenticate(this.token), 3000);
            sendMessage("hello");
        }

        // initializie web socket connection
        initWebSocket(onOpen, this.handleMessage);

    }
}
</script>

<style scoped>

.chat-container {
    width: 60%;
}

@media (max-width:500px) {
    .chat-container {
        width: 90% !important;
    }
}

</style>
