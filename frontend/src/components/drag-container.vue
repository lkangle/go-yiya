<script setup>
import { ref } from 'vue';

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