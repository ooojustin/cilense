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

export default {
    name: 'Room',
    components: {
        Message
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

        const url = "ws" + vars.backend.slice(4) + "ws";
        let ws = new WebSocket(url);

        ws.onmessage = msg => {
            console.log(msg);
        }

        ws.onopen = () => {
            const params = {
                action: "join_room",
                data: {
                    room_id: id
                }
            };
            ws.send(JSON.stringify(params));
        }

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
