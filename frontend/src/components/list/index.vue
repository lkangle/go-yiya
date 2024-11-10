<script setup>
import { range } from '@utils/index';
import Item from './item.vue';
import { PlusOutlined } from '@ant-design/icons-vue';
import Operate from './operate.vue';


</script>

<template>
    <div class="h-full">
        <Item v-for="(item, index) in range(0, 8)"></Item>
    </div>

    <div class="absolute bottom-0 h-52 w-full bg-[#f8f9fa] border-t-[1px] border-gray-300">
        <div class="flex items-center justify-between h-full px-10">
            <a-button type="primary" shape="circle" class="pt-5 flex justify-center items-center">
                <template #icon>
                    <PlusOutlined class="text-15" />
                </template>
            </a-button>
            <div class="text-12 leading-14 px-10 flex-1">
                <p class="text-gray-500">成功压缩0个图片，共10个图片</p>
                <p class="text-10 text-gray-400">
                    <span>体积降低</span>
                    <span class="text-green-500 px-2">50%</span>
                    <span>(0.00KB)</span>
                </p>
            </div>
            <Operate/>
        </div>
    </div>
</template>

<style scoped>
</style>