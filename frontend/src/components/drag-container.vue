<script setup>
import { EventsOn, OnFileDrop } from '@wailsjs/runtime/runtime';
import { ref } from 'vue';

// 这里只做拖拽的样式处理，文件交给wails处理

OnFileDrop((x,y,paths) => {
    console.log(x, y, paths)
})

EventsOn("file_handler_finish", data => {
    console.log("...data", data)
})

const over = ref(false)

const onDrop = event => {
    event.preventDefault()
    over.value = false
}

const onDragEnter = event => {
    event.preventDefault()
    over.value = true
}

const onDragOver = (event) => {
    event.preventDefault()
}

const onDragCancel = event => {
    event.preventDefault()
    over.value = false
}
</script>

<template>
    <div
        :data-dragging-over="over"
        style="--wails-drop-target: drop"
        v-on:dragenter="onDragEnter"
        v-on:dragover="onDragOver"
        v-on:dragleave="onDragCancel"
        v-on:drop="onDrop"
    >
        <slot></slot>
    </div>
</template>