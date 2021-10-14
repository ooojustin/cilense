<template>
    <div class="flex items-center justify-center h-screen">
        <div class="pw-container bg-gray-700 rounded-md border border-gray-500" v-if="!session">
            <div class="py-5 text-xl">
                Join Room
            </div>
            <div class="bg-red-600 mx-auto rounded-md py-2 border-2 border-red-300 mb-3 font-semibold pw-alert" v-if="wrong_password">
                Incorrect password.
            </div>
            <input class="bg-gray-800 rounded-md py-2 px-2.5" name="password" v-model="password" type="password" placeholder="Password" />
            <br />
            <button class="bg-blue-700 hover:bg-blue-500 rounded-md py-2 mt-3 mb-6" type="button" @click="onSubmitPassword">
                Submit
            </button>
        </div>
        <div class="chat-container bg-gray-700 p-5 rounded-md border border-gray-500" style="height: 60%;" v-if="session">
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
    sendMessage
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
            }

        },
        onSubmitPassword() {
            // join via websocket
            joinRoom(this.room.id, this.password);
        }
    },
    data() {
        return {
            room: null,
            session: null,
            password: "",
            wrong_password: false
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
            // this.session || joinRoom(id);
            // setTimeout(() => sendMessage("hello"), 3000);
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
.pw-container {
    width: 600px;
    margin-bottom: 240px;
}
input, 
button,
.pw-alert { 
    width: 450px; 
}

@media (max-width:500px) {
    input,
    button,
    .chat-container,
    .pw-container,
    .pw-alert {
        width: 90% !important;
    }
}

</style>
