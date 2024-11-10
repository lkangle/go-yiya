<script setup>
import { debounce } from '@utils/index';
import { GetCompressOptions, UpdateCompressOptions } from '@wailsjs/go/services/appService';
import { Modal } from 'ant-design-vue'
import { reactive, ref, watch } from 'vue';

const quaOptions = [
    {label: "优质", value: 2},
    {label: "一般", value: 1},
]

const data = reactive({
    quality: 1,
    override: true,
    newSuffix: ''
})
GetCompressOptions().then((resp) => {
    if(resp.success) {
        Object.assign(data, resp.data)
    }
})

const saveOptions = debounce((data) => {
    UpdateCompressOptions(data)
}, 500);
watch(() => data, (n) => saveOptions({...n}), {deep: true})

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
                        <a-switch v-model:checked="data.override" />
                    </div>
                </div>
                <div class="mb-12 flex items-center justify-between">
                    <div class="w-85 pr-10 whitespace-nowrap">
                        <span>新文件后缀:</span>
                    </div>
                    <div class="flex-1">
                        <a-input :disabled="data.override" autocomplete="off" v-model:value="data.newSuffix" />
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
                        <a-select class="w-full" :options="quaOptions" v-model:value="data.quality"/>
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