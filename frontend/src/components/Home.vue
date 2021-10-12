<template>
    <div class="flex items-center justify-center h-screen" >
        <CreateRoom @room-created="onCreate" v-if="room == null" />
        <RoomShare :room="room" v-if="room != null" />
    </div>
</template>

<script>
import CreateRoom from './CreateRoom';
import RoomShare from './RoomShare';

export default {
    name: 'Home',
    components: {
        CreateRoom,
        RoomShare
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
