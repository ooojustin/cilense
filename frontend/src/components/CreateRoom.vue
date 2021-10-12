<template>
    <div class="create-room bg-gray-700 rounded-md">
        <div class="py-5 text-xl">
            Create Room
        </div>
        <input class="bg-gray-800 rounded-md py-1 px-2.5" type="text" />
        <br />
        <button class="bg-blue-700 rounded-md py-2 mt-2" type="button">
            Generate
        </button>
    </div>
</template>

<script>
export default {
    name: 'CreateRoom'
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
