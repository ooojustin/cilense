<template>
    <div class="flex items-center justify-center h-screen">

        <JoinRoom @password-submit="onSubmitPassword" :wrong_password="this.wrong_password" v-if="!session" />

        <div class="chat-container bg-gray-700 p-5 rounded-md border border-gray-500 flex flex-col justify-between" style="height: 60%;" v-if="session">
            <div class="message-container">
                <Message :data="msg" :key="index" v-for="(msg, index) in messages" />
            </div>
            <div>
                <input class="bg-gray-800 rounded-md py-2 px-2.5 float-left msg-input" style="width: 88%;" name="password" v-model="message" type="text" placeholder="..." />
                <button class="bg-blue-700 hover:bg-blue-500 rounded-md py-2 float-right send-btn" style="width: 10%;" type="button" @click="onSendMessage">
                    Send
                </button>
            </div>
        </div>

    </div>
</template>

<script>
import vars from "../variables";
import JoinRoom from "./JoinRoom";
import Message from "./Message";

import { 
    initWebSocket,
    restoreSession,
    joinRoom,
    sendMessage
} from '../websocket';

export default {
    name: 'Room',
    components: {
        JoinRoom,
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
                    // store session inside state & local storage
                    this.session = msg.session_id;
                    const sessions = JSON.parse(localStorage.getItem("sessions")) || {};
                    sessions[this.room.id] = msg.session_id;
                    localStorage.setItem("sessions", JSON.stringify(sessions));
                    break;
                }
                case "wrong_password": {
                    this.wrong_password = true;
                    break;
                }
                case "message_sent":
                case "new_message": {
                    this.messages = this.messages.concat(msg.message);
                    break;
                }
            }

        },
        onSubmitPassword(password) {
            // join via websocket
            console.log("onSubmitPassword");
            joinRoom(this.room.id, password);
        },
        onSendMessage() {
            sendMessage(this.message);
            this.message = "";
        }
    },
    data() {
        return {
            room: null,
            session: null,
            password: "",
            wrong_password: false,
            message: "",
            messages: []
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

        // retrieve stored session from local storage
        const sessions = JSON.parse(localStorage.getItem("sessions")) || {};
        if (Object.prototype.hasOwnProperty.call(sessions, id))
            this.session = sessions[id];

        // func will when websocket connect is opened
        const onOpen = () => {
            console.log("Connected to websocket.");
            this.session && restoreSession(this.session);
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
.message-container {
    max-height: 90%;
    overflow-x: hidden;
    overflow-y: scroll;
}

@media (max-width:500px) {
    input,
    button,
    .chat-container {
        width: 90% !important;
    }
    .send-btn,
    .msg-input {
        margin-top: 10px;
        width: 100% !important;
    }
    .send-btn {
        margin-top: 10px;
    }
}

</style>
