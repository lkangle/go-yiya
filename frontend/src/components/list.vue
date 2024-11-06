<script setup>
import { ref } from 'vue';

let isEmpty = ref(true)
</script>

<template>
    <div :if="isEmpty" class="w-full h-full flex justify-center items-center pb-50">
        <Empty/>
    </div>
</template>