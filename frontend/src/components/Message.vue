<template>
    <div class="message text-left rounded-md bg-gray-800 border border-gray-600 p-3">
        {{ text }}
    </div>
</template>

<script>
export default {
    name: 'Message',
    props: {
        text: String
    }
}
</script>

<style scoped>
.message {
    width: 40%;
}
@media (max-width:500px) {
    .message {
        width: 100%;
    }
}
</style>
