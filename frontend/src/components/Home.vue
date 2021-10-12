<template>
    <div class="flex justify-center h-screen" style="padding-top: 200px;">
        <CreateRoom />
    </div>
</template>

<script>
import CreateRoom from './CreateRoom';

export default {
    name: 'Home',
    components: {
        CreateRoom
    }
}
</script>

<style scoped>
</style>
