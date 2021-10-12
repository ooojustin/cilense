<template>
    <div class="flex items-center justify-center h-screen" >
        <CreateRoom @room-created="onCreate" v-if="room == null" />
    </div>
</template>

<script>
import CreateRoom from './CreateRoom';

export default {
    name: 'Home',
    components: {
        CreateRoom
    },
    data() {
        return {
            room: null
        };
    },
    methods: {
        onCreate(room) {
            this.room = room;
        }
    }
}
</script>

<style scoped>
</style>
