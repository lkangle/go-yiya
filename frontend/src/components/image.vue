<script setup>
import { isDev } from '@utils/platform';
import { watchEffect, ref, useAttrs } from 'vue';
defineOptions({
    inheritAttrs: false
})
const attrs = useAttrs()

// 图片的objectURL
const objSrc = ref('')

watchEffect((onCleanup) => {
    const ref = {src:''}
    const api = `/image.do?path=${attrs.src}`

    // dev环境下估计时有bug的，get请求会通过vite服务先响应个错误的数据，这里就只能post处理下
    if (isDev()) {
        fetch(api, {method: "post"}).then(resp => resp.blob()).then((b) => {
            ref.src = URL.createObjectURL(b)
            objSrc.value = ref.src
        })
        onCleanup(() => {
            if (ref.src) {
                URL.revokeObjectURL(ref.src)
            }
        })
    } else {
        objSrc.value = api
    }
})
</script>

<template>
    <img :class="attrs.class" draggable="false" :src="objSrc" alt="">
</template>