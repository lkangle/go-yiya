<script setup>
import { ref } from 'vue';
import empty from './empty.vue';

let isEmpty = ref(true)
</script>

<template>
    <div :if="isEmpty" class="w-full h-full flex justify-center items-center pb-50">
        <empty/>
    </div>
</template>