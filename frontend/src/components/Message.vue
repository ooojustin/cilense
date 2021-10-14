<template>
    <div class="msg mt-2 text-left rounded-md bg-blue-600 border border-blue-400 p-2" v-bind:class="{ 'msg-out': data.sent, 'msg-in': !data.sent }">
        {{ data.text }}
        <br />
        <span class="text-xs text-gray-300">
            @{{ data.alias }}
        </span>
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
    background-color: #1F2937;
    border-color: #4B5563;
    margin-right: 10px;
}
@media (max-width:500px) {
    .msg {
        width: 100%;
    }
}
</style>
