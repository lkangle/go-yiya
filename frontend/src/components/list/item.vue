<script setup>
import { ArrowRightOutlined } from '@ant-design/icons-vue';
import Operate from './operate.vue';

const props = defineProps(['item'])

console.log('props', props)
</script>

<template>
    <div class="compress-item h-45 flex items-center justify-between px-15 odd:bg-gray-50">
        <div className="status-icon flex items-center justify-center">
            <i data-status="3"></i>
        </div>
        <div class="w-38 h-38 mx-12 flex items-center justify-center">
            <img class="object-contain" draggable="false" src="https://cdn.lkangle.cn/zimage/fa7t9r3n8u.png" alt=""/>
        </div>
        <div class="leading-14 flex-1">
            <p class="text-11 text-gray-500">这是文件名字.png</p>
            <p class="status-ok text-10 text-gray-400">
                <span class="text-green-500 mr-4">90%</span>
                <span>200KB</span>
                <ArrowRightOutlined class="scale-75 pt-1 mx-2"/>
                <span>10KB</span>
            </p>
            <!-- <p class="status-ok text-10 text-gray-400">
                <span>等待压缩</span>
            </p> -->
        </div>
        <Operate class="scale-75 -mr-10" :is-item="true"/>
    </div>
</template>

<style scoped>
@keyframes spinning {
  0% {
    transform: rotate(0deg);
  }
  100% {
    transform: rotate(360deg);
  }
}

.status-icon {
    width: 8px;

    i {
      display: block;
      box-sizing: border-box;
      width: 6px;
      height: 6px;
      border-radius: 50%;
      @apply bg-gray-400;

      /* 压缩中 */
      &[data-status='2'] {
        width: 8px;
        height: 8px;
        background-color: transparent;
        border: solid 1px;
        border-bottom-color: transparent !important;
        animation: spinning 800ms linear infinite;
        
        @apply border-blue-600;
      }
      
      /* 成功 */
      &[data-status='3'] {
        @apply bg-green-500;
      }
      
      /* 失败 */
      &[data-status='4'] {
        @apply bg-red-600;
      }
    }
}
</style>