<script setup>
import Empty from './empty.vue';
import List from './list/index.vue';
import DragContainer from './drag-container.vue';
import { getAppVersion } from '@utils/platform';
import useList from '@/store/use-list';

const store = useList()
</script>

<template>
    <DragContainer class="flex-1 min-h-300">
        <template v-if="store.isEmpty">
            <Empty class="mx-auto mt-115"/>
            <div class="px-40 mt-100">
                <button @click="store.openSelectImages" class="btn-indigo w-full">
                    选择图片  
                </button>
            </div>
            <footer class="text-11 text-center text-gray-400 mt-50">
                <p>Oh. Work-Life Balance</p>
                <p class="text-10">© 2024. v{{ getAppVersion() }} </p>
            </footer>
        </template>
        <List v-if="!store.isEmpty"/>
    </DragContainer>
</template>
<style scoped>
.btn-indigo {
    @apply bg-indigo-500 hover:bg-indigo-600 active:bg-indigo-700 focus:outline-none text-white font-bold py-8 px-12 rounded-xl;
}
</style>