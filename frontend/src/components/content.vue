<script setup>
import { ref } from 'vue';
import Empty from './empty.vue';
import List from './list.vue';

let isEmpty = ref(true)
</script>

<template>
    <div class="flex-1 min-h-300 bg-[#f8f9fb]">
        <template v-if="isEmpty">
            <Empty class="mx-auto mt-115"/>
            <div class="px-40 mt-100">
                <button class="btn-indigo w-full">
                    选择图片  
                </button>
            </div>
            <footer class="text-11 text-center text-gray-400 mt-50">
                <p>Oh. Work-Life Balance</p>
            </footer>
        </template>
        <List v-if="!isEmpty"/>
    </div>
</template>

<style scoped>
.btn-indigo {
    @apply bg-indigo-500 hover:bg-indigo-600 active:bg-indigo-700 focus:outline-none text-white font-bold py-8 px-12 rounded-xl;
}
</style>