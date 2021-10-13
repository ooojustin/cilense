<template>
    <div class="create-room bg-gray-700 rounded-md border border-gray-500">
        <div class="py-5 text-xl">
            Create Room
        </div>
        <input class="bg-gray-800 rounded-md py-2 px-2.5" name="password" v-model="password" type="password" placeholder="Password" />
        <br />
        <button class="bg-blue-700 hover:bg-blue-500 rounded-md py-2 mt-3 mb-6" type="button" @click="onSubmit">
            Generate
        </button>
    </div>
</template>

<script>
import vars from '../variables';

export default {
    name: 'CreateRoom',
    data() {
        return {
            password: ""
        }
    },
    methods: {
        onSubmit() {
            const { password } = this;
            vars.api.post("create_room", { password })
            .then(r => {
                const { data } = r.data;
                this.$emit("room-created", data);
                this.storeToken(data);
            });
        },
        storeToken(room) {
            const tokens = JSON.parse(localStorage.getItem("tokens")) || {};
            tokens[room.id] = room.token;
            localStorage.setItem("tokens", JSON.stringify(tokens));
        }
    }
}
</script>

<style scoped>

.create-room {
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
    .create-room { 
        width: 90% !important; 
    }
}

</style>
