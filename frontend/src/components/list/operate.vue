<script setup>
import { UndoOutlined, RedoOutlined, CloseOutlined } from '@ant-design/icons-vue';

const props = defineProps(['class'])
const emit = defineEmits(["undo", "redo", "close"])
</script>

<template>
    <div :class="['flex space-x-2', props.class]">
        <a-button @click="emit('undo')" type="text" shape="circle" class="flex justify-center items-center">
            <template #icon>
                <UndoOutlined class="text-15 text-gray-500" />
            </template>
        </a-button>
        <a-button @click="emit('redo')" disabled type="text" shape="circle" class="flex justify-center items-center">
            <template #icon>
                <RedoOutlined class="text-15 text-gray-300" />
            </template>
        </a-button>
        <a-button @click="emit('close')" type="text" shape="circle" class="flex justify-center items-center">
            <template #icon>
                <CloseOutlined class="text-15 text-gray-500" />
            </template>
        </a-button>
    </div>
</template>