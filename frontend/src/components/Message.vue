<template>
    <div class="msg mb-4 text-left rounded-md bg-blue-600 border border-blue-400 p-2" v-bind:class="{ 'msg-out': data.sent, 'msg-in': !data.sent, 'msg-room': data.is_room_message }">
        {{ data.text }}
        <template v-if="!data.is_room_message">
            <br />
            <span class="text-xs text-gray-300">
                @{{ data.alias }}
            </span>
        </template>
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
.msg-room {
    margin-left: auto;
    margin-right: auto;
    width: 65%;
    text-align: center;
    border-width: 3px;
    background-color: #DB2777;
    border-color: #F472B6;
}
@media (max-width:500px) {
    .msg {
        width: 100% !important;
    }
}
</style>
