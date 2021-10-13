<template>
    <div>
        {{ room }}
    </div>
</template>

<script>
import vars from "../variables";
export default {
    name: 'Room',
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
            setTimeout(() => ws.send("ping"), 1000);
        }

    }
}
</script>

<style scoped>
</style>
