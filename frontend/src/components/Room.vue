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
    }
}
</script>

<style scoped>
</style>
