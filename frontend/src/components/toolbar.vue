<script setup>
import { PushpinOutlined, SettingFilled, PushpinFilled, SettingOutlined, MinusOutlined, CloseOutlined} from '@ant-design/icons-vue'
import { getAppName, isAlwaysOnTop, isWindows } from '@utils/platform';
import { SetIsAlwaysOnTop } from '@wailsjs/go/services/appService';
import { WindowMinimise, Quit } from '@wailsjs/runtime/runtime';
import { ref } from 'vue';
import Settings from './settings.vue';

const title = ref(getAppName())
const alwaysTop = ref(isAlwaysOnTop())

const toggleAlwaysTop = () => {
    if (alwaysTop.value) {
        alwaysTop.value = false
    } else{
        alwaysTop.value = true
    }

    SetIsAlwaysOnTop(alwaysTop.value)
}
</script>

<template>
<div
    id="app-toolbar"
    class="h-38 w-full text-gray-600 bg-white"
    style="--wails-draggable: drag">
    <div v-if="isWindows()" class="h-full flex items-center justify-between px-8">
        <!-- 置顶和设置按钮 -->
        <div class="flex space-x-3">
            <div @click="toggleAlwaysTop" :class="['bar-icon', {
                actived: alwaysTop
            }]">
                <PushpinOutlined :class="{ '-rotate-45': alwaysTop}"/>
            </div>
            <Settings class="bar-icon">
                <SettingOutlined />
            </Settings>
        </div>
        <h3>{{ title }}</h3>
        <!-- window的关闭按钮 -->
        <div class="flex space-x-3">
            <div @click="WindowMinimise" class="bar-icon hover:text-yellow-600">
                <MinusOutlined />
            </div>
            <div @click="Quit" class="bar-icon hover:bg-red-600 hover:text-gray-100">
                <CloseOutlined />
            </div>
        </div>
    </div>

    <div v-if="!isWindows()" class="h-full flex items-center justify-center relative">
        <h3>{{ title }}</h3>
        <!-- 置顶和设置按钮 -->
        <div class="flex space-x-3 absolute right-8">
            <div @click="toggleAlwaysTop" :class="['bar-icon', {
                actived: alwaysTop
            }]">
                <PushpinFilled :class="{ '-rotate-45': alwaysTop}"/>
            </div>
            <Settings class="bar-icon">
                <SettingFilled />
            </Settings>
        </div>
    </div>
</div>
</template>

<style>
.bar-icon {
    @apply text-center text-12 w-22 h-22 flex items-center justify-center cursor-pointer rounded-[5px] hover:bg-gray-100;
    transition: all linear 100ms;

    &.actived {
        @apply bg-indigo-500 text-gray-100;
    }
}
</style>