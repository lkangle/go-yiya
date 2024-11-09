<script setup>
import icon from '@assets/images/up-image@512.png'

const props = defineProps(['class'])
</script>

<template>
    <div :class="['w-220 h-140 text-gray-500 text-12 flex-col flex items-center justify-center', props.class]">
        <img draggable="false" :src="icon" width="64" height="52" class="w-65"/>
        <div class="mt-20 mb-8">拖入/选择图片开始压缩</div>
        <div>支持JPEG,PNG,WebP,GIF格式</div>
    </div>
</template>