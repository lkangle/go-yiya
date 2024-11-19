<script setup>
import Item from './item.vue';
import { PlusOutlined } from '@ant-design/icons-vue';
import Operate from './operate.vue';
import useList from '@/store/use-list';
import { onUnmounted } from 'vue';
import useCurrent from '@/store/use-current';

const store = useList()

const curStore = useCurrent()
onUnmounted(() => {
    curStore.$reset()
})
</script>

<template>
<div class="h-full flex flex-col">
    <div class="scrollbar h-full overflow-auto">
        <Item v-for="(item) in store.dataList" :key="item.id" :item="item"/>
    </div>

    <div class="min-h-50 h-50 w-full bg-[#f8f9fa] border-t-[1px] border-gray-300">
        <div class="flex items-center justify-between h-full px-10">
            <a-button @click="store.openSelectImages" type="primary" shape="circle" class="pt-5 flex justify-center items-center">
                <template #icon>
                    <PlusOutlined class="text-15" />
                </template>
            </a-button>
            <div class="text-12 leading-14 px-10 flex-1">
                <p class="text-gray-500">
                    成功压缩{{ store.stat.successCount }}个图片，共{{ store.stat.total }}个图片
                </p>
                <p class="text-10 text-gray-400">
                    <span>体积节省</span>
                    <span :class="['text-green-500 px-2', {
                        '!text-yellow-600': store.stat.rate<=0
                    }]">{{ store.stat.rateText }}</span>
                    <span>({{ store.stat.dtext }})</span>
                </p>
            </div>
            <Operate @close="store.clear"/>
        </div>
    </div>
</div>
</template>

<style scoped>
/* 定义滚动条样式 */
.scrollbar::-webkit-scrollbar {
    width: 8px;
    height: 1px;
}

.scrollbar::-webkit-scrollbar-thumb {
    width: 6px;
    border-radius: 5px;
    border: 2px solid #f1f1f1;
    @apply bg-gray-400 hover:bg-gray-500;
}
</style>