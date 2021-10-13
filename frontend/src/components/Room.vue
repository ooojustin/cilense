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
    joinRoom,
    authenticate
} from '../websocket';

export default {
    name: 'Room',
    components: {
        Message
    },
    methods: {
        handleMessage(msg) {
            // handle message from websocket
            console.log(msg);
        }
    },
    data() {
        return {
            room: null
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

        // retrieve token for room from local storage, if it exists
        let token = null;
        const tokens = JSON.parse(localStorage.getItem("tokens")) || {};
        if (Object.prototype.hasOwnProperty.call(tokens, id))
            token = tokens[id];

        // func will when websocket connect is opened
        const onOpen = () => {
            console.log("Connected to websocket.");
            joinRoom(id);
            token && authenticate(token);
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
