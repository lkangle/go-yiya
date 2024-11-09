<script setup>
import { PushpinFilled, SettingFilled, MinusOutlined, CloseOutlined} from '@ant-design/icons-vue'
import { getAppName, isAllowsOnTop, isWindows } from '@utils/platform';
import { SetIsAllowsOnTop } from '@wailsjs/go/services/systimeService';
import { WindowMinimise, Quit } from '@wailsjs/runtime/runtime';
import { ref } from 'vue';

const title = ref(getAppName())
const allowsTop = ref(isAllowsOnTop())

const toggleAllowsTop = () => {
    if (allowsTop.value) {
        allowsTop.value = false
    } else{
        allowsTop.value = true
    }

    SetIsAllowsOnTop(allowsTop.value)
}

</script>

<template>
<div
    id="app-toolbar"
    class="h-38 w-full text-gray-600 bg-white"
    style="--wails-draggable: drag">
    <div class="h-full flex items-center justify-between px-6">
        <!-- window的关闭按钮 -->
        <div :class="['flex space-x-2', {
            invisible: !isWindows()
        }]">
            <div @click="WindowMinimise" class="bar-icon hover:text-yellow-600 hover:bg-gray-100">
                <MinusOutlined />
            </div>
            <div @click="Quit" class="bar-icon hover:text-red-800 hover:bg-gray-100">
                <CloseOutlined />
            </div>
        </div>
        <h3>{{ title }}</h3>
        <!-- 指定和设置按钮 -->
        <div class="flex space-x-2">
            <div @click="toggleAllowsTop" :class="['bar-icon', {
                actived: allowsTop
            }]">
                <PushpinFilled :class="{ '-rotate-45': allowsTop}"/>
            </div>
            <div class="bar-icon">
                <SettingFilled />
            </div>
        </div>
    </div>
</div>
</template>

<style>
.bar-icon {
    @apply text-12 w-21 h-21 flex items-center justify-center cursor-pointer rounded-[5px];

    &.actived {
        @apply bg-indigo-500 text-gray-100;
    }
}
</style>