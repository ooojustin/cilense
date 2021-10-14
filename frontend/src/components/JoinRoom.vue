<template>
    <div class="pw-container bg-gray-700 rounded-md border border-gray-500" v-if="!session">
        <div class="py-5 text-xl">
            Join Room
        </div>
        <div class="bg-red-600 mx-auto rounded-md py-2 border-2 border-red-300 mb-3 font-semibold pw-alert" v-if="wrong_password">
            Incorrect password.
        </div>
        <input class="bg-gray-800 rounded-md py-2 px-2.5" name="password" v-model="password" type="password" placeholder="Password" />
        <br />
        <button class="bg-blue-700 hover:bg-blue-500 rounded-md py-2 mt-3 mb-6" type="button" @click="onSubmit">
            Submit
        </button>
    </div>
</template>

<script>
export default {
    name: 'JoinRoom',
    props: {
        wrong_password: Boolean
    },
    data() {
        return {
            password: ""
        };
    },
    methods: {
        onSubmit() {
            this.$emit("password-submit", this.password);
        }
    }
}
</script>

<style scoped>
.pw-container {
    width: 600px;
    margin-bottom: 240px;
}
input, 
button,
.pw-alert { 
    width: 450px; 
}
@media (max-width:500px) {
    input,
    button,
    .pw-container,
    .pw-alert {
        width: 90% !important;
    }
}
</style>
