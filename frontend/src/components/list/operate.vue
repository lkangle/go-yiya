<script setup>
import useCurrent from '@/store/use-current';
import { RollbackOutlined, RetweetOutlined, CloseOutlined } from '@ant-design/icons-vue';
import OpButton from "@/components/op-button.vue";
import {RestoreImage} from "@wailsjs/go/services/fileService";
import {debounce} from "@utils/index.js";
import useCompress from "@/store/use-compress.js";

const emitCompress = useCompress()

const props = defineProps(['class'])
const emit = defineEmits(["close"])

const store = useCurrent()

const onUndo = debounce(async () => {
  await RestoreImage(store.current.result)
  Object.assign(store.current.result, {
    status: 5
  })
}, 500, true)

const onRedo= debounce(() => {
  emitCompress(store.current, (it) => {
    Object.assign(store.current.result, it)
  })
}, 500, true)
</script>

<template>
    <div :class="['flex space-x-2', props.class]">
        <template v-if="store.current?.id">
            <a-tooltip title="还原" placement="top">
                <op-button @click="onUndo" :icon="RollbackOutlined" :disabled="!store.canUndo"/>
            </a-tooltip>
            <a-tooltip title="重压" placement="top">
                <op-button @click="onRedo" :icon="RetweetOutlined" :disabled="!store.canRedo"/>
            </a-tooltip>
        </template>
        <a-button @click="emit('close')" type="text" shape="circle" class="flex justify-center items-center">
            <template #icon>
                <CloseOutlined class="text-15 text-gray-500" />
            </template>
        </a-button>
    </div>
</template>