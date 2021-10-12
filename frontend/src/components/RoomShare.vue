<template>
    <div class="share-room bg-gray-700 rounded-md">
        <div class="py-5 text-xl">
            Share Room
        </div>
        <input class="text-center bg-gray-800 rounded-md py-2 px-2.5" type="text" v-model="roomURL" readOnly />
        <br />
        <button class="bg-blue-700 rounded-md py-2 my-3" type="button" @click="onSubmit">
            Copy URL
        </button>
        <br />
        <button class="bg-blue-700 rounded-md py-2 mb-6" type="button" @click="onSubmit">
            Open Room
        </button>
    </div>
</template>

<script>
import vars from '../variables';
export default {
    name: 'RoomShare',
    props: {
        room: Object
    },
    data() {
        return {
            roomURL: null
        };
    },
    created() {
        const url = `${vars.frontend}c/${this.room.id}`;
        this.roomURL = url;
    }
}
</script>

<style scoped>

.share-room {
    width: 600px;
    margin-bottom: 240px;
}

input, 
button { 
    width: 450px; 
}

@media (max-width:500px) {
    input, 
    button,
    .share-room { 
        width: 90% !important; 
    }
}

</style>

