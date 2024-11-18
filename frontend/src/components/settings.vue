<script setup>
import useOptions from '@/store/use-options';
import { Modal } from 'ant-design-vue'
import { ref, watch } from 'vue';

const quaOptions = [
    {label: "优质", value: 2},
    {label: "一般", value: 1},
]

const store = useOptions()
watch(() => store.data, (n) => store.save({...n}), {deep: true})

const isOpen = ref(false)
const toggle = () => {
    isOpen.value = !isOpen.value
}
</script>

<template>
    <div @click="toggle">
        <slot/>
        <Modal title="配置" width="280px" v-model:open="isOpen" :footer="null">
            <div class="pt-5">
                <div class="mb-12 flex items-center justify-between">
                    <div class="w-85 pr-10 whitespace-nowrap">
                        <span>覆盖原文件:</span>
                    </div>
                    <div class="flex-1 flex justify-end">
                        <a-switch v-model:checked="store.data.override" />
                    </div>
                </div>
                <div class="mb-12 flex items-center justify-between">
                    <div class="w-85 pr-10 whitespace-nowrap">
                        <span>新文件后缀:</span>
                    </div>
                    <div class="flex-1">
                        <a-input :disabled="store.data.override" autocomplete="off" v-model:value="store.data.newSuffix" />
                    </div>
                </div>
                <div class="mb-15 flex items-center justify-between">
                    <div class="w-85 pr-10 flex-nowrap flex justify-between">
                        <span>压</span>
                        <span>缩</span>
                        <span>质</span>
                        <span>量:</span>
                    </div>
                    <div class="flex-1">
                        <a-select class="w-full" :options="quaOptions" v-model:value="store.data.quality"/>
                    </div>
                </div>
                <div class="text-gray-400 text-center text-12">
                    Oh. Life Is Beautiful.
                </div>
            </div>
        </Modal>
    </div>
</template>

<style>
</style>