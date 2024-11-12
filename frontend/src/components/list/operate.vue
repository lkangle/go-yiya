<script setup>
import { UndoOutlined, RedoOutlined, CloseOutlined, FolderOpenOutlined } from '@ant-design/icons-vue';
import { ref } from 'vue';

const props = defineProps(['class', 'isItem'])
const emit = defineEmits(["undo", "redo", "openfolder", "close"])

let isOverride = ref(true)
</script>

<template>
    <div :class="['flex space-x-2', props.class]">
        <template v-if="isOverride">
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
        </template>
        <a-button  v-if="!isOverride" @click="emit('openfolder')" type="text" shape="circle" class="flex justify-center items-center">
            <template #icon>
                <FolderOpenOutlined class="text-15 text-gray-500" />
            </template>
        </a-button>

        <a-button v-if="!props.isItem" @click="emit('close')" type="text" shape="circle" class="flex justify-center items-center">
            <template #icon>
                <CloseOutlined class="text-15 text-gray-500" />
            </template>
        </a-button>
    </div>
</template>