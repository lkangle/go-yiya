<script setup>
import useList from '@/store/use-list';
import { RollbackOutlined, RetweetOutlined, CloseOutlined } from '@ant-design/icons-vue';

const current = useList(stat => stat.current)

console.log("[current]", current)

const props = defineProps(['class'])
const emit = defineEmits(["undo", "redo", "close"])
</script>

<template>
    <div :class="['flex space-x-2', props.class]">
        <a-tooltip title="还原" placement="top">
            <a-button @click="emit('undo')" type="text" shape="circle" class="flex justify-center items-center">
                <template #icon>
                    <RollbackOutlined class="text-15 text-gray-500" />
                </template>
            </a-button>
        </a-tooltip>
        <a-tooltip title="重压" placement="top">
            <a-button @click="emit('redo')" disabled type="text" shape="circle" class="flex justify-center items-center">
                <template #icon>
                    <RetweetOutlined class="text-16 text-gray-300" />
                </template>
            </a-button>
        </a-tooltip>
        <a-button @click="emit('close')" type="text" shape="circle" class="flex justify-center items-center">
            <template #icon>
                <CloseOutlined class="text-15 text-gray-500" />
            </template>
        </a-button>
    </div>
</template>