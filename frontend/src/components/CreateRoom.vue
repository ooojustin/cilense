<template>
    <div class="create-room bg-gray-700 rounded-md">
        <div class="py-5 text-xl">
            Create Room
        </div>
        <input class="bg-gray-800 rounded-md py-2 px-2.5" name="password" v-model="password" type="password" placeholder="Password" />
        <br />
        <button class="bg-blue-700 rounded-md py-2 mt-3" type="button" @click="onSubmit">
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
            .then(r => console.log(r.data));
        }
    }
}
</script>

<style scoped>
.create-room {
    width: 600px;
    height: 180px;
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
