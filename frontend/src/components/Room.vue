<template>
    <div class="flex items-center justify-center h-screen">
        <div class="chat-container bg-gray-700 rounded-md p-5 border border-gray-500" style="height: 60%;">
            <Message text="test message" />
        </div>
    </div>
</template>

<script>
import vars from "../variables";
import Message from "./Message";

import { 
    initWebSocket,
    joinRoom
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

        const { id } = this.$route.params;
        vars.api.get("/room/" + id)
        .then(r => {
            const { data } = r.data;
            this.room = data;
        });

        const onOpen = () => {
            console.log("Connected to websocket.");
            joinRoom(id);
        }

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
