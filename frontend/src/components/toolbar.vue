<script setup>
import { PushpinOutlined, SettingFilled, PushpinFilled, SettingOutlined, MinusOutlined, CloseOutlined} from '@ant-design/icons-vue'
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
    <div v-if="isWindows()" class="h-full flex items-center justify-between px-8">
        <!-- 置顶和设置按钮 -->
        <div class="flex space-x-3">
            <div @click="toggleAllowsTop" :class="['bar-icon', {
                actived: allowsTop
            }]">
                <PushpinOutlined :class="{ '-rotate-45': allowsTop}"/>
            </div>
            <div class="bar-icon">
                <SettingOutlined />
            </div>
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
    @apply text-center text-12 w-22 h-22 flex items-center justify-center cursor-pointer rounded-[5px] hover:bg-gray-100;
    transition: all linear 100ms;

    &.actived {
        @apply bg-indigo-500 text-gray-100;
    }
}
</style>