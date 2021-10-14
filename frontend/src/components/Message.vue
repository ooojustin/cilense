<template>
    <div class="msg mt-2 text-left rounded-md bg-gray-800 border border-gray-600 p-3" v-bind:class="{ 'msg-out': data.sent, 'msg-in': !data.sent }">
        {{ data.text }}
    </div>
</template>

<script>
export default {
    name: 'Message',
    props: {
        data: Object
    }
}
</script>

<style scoped>
.msg {
    width: 40%;
}
.msg-out {
    margin-left: auto;
    background-color: #2563EB;
    border-color: #60A5FA;
}
@media (max-width:500px) {
    .msg {
        width: 100%;
    }
}
</style>
