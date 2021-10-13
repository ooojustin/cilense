<template>
    <div class="share-room bg-gray-700 rounded-md border border-gray-500">
        <div class="py-5 text-xl">
            Share Room
        </div>
        <div class="bg-green-600 mx-auto rounded-md py-2 border-2 border-green-300 mb-3 font-semibold copy-alert" v-if="copied">
            URL copied to clipboard.
        </div>
        <input class="text-center bg-gray-800 rounded-md py-2 px-2.5" id="room-url" type="text" v-model="roomURL" readOnly />
        <br />
        <button class="bg-blue-700 hover:bg-blue-500 rounded-md py-2 my-3" type="button" @click="onCopy">
            Copy URL
        </button>
        <br />
        <button class="bg-blue-700 hover:bg-blue-500 rounded-md py-2 mb-6" type="button" @click="onOpen">
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
            roomURL: null,
            copied: false
        };
    },
    created() {
        const url = `${vars.frontend}c/${this.room.id}`;
        this.roomURL = url;
    },
    methods: {
        onCopy() {
            const input = document.getElementById("room-url");
            input.focus();
            input.select();
            document.execCommand("copy");
            this.copied = true;
        },
        onOpen() {
            this.$router.push(`/c/${this.room.id}`);
        }
    }
}
</script>

<style scoped>

.share-room {
    width: 600px;
    margin-bottom: 240px;
}

input, 
button,
.copy-alert { 
    width: 450px; 
}

@media (max-width:500px) {
    input, 
    button,
    .share-room,
    .copy-alert { 
        width: 90% !important; 
    }
}

</style>

