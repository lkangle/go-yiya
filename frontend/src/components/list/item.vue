<script setup>
import { ArrowRightOutlined, FolderOpenOutlined, ReadOutlined } from '@ant-design/icons-vue';
import { formatFileSize } from '@utils/index';
import Image from '../image.vue';

const props = defineProps(['item'])
const item = props.item
console.log('[item]', item, props)

</script>

<template>
    <div @click="$emit('select', item)" class="compress-item h-45 flex items-center justify-between px-15 odd:bg-gray-50">
        <div className="status-icon flex items-center justify-center">
            <i data-status="3"></i>
        </div>
        <div class="w-35 h-35 rounded-[5px] overflow-hidden mx-12 flex items-center justify-center">
          <Image class="object-contain w-full h-full" :src="item.path"/>
        </div>
        <div class="leading-14 flex-1">
            <p class="text-11 text-gray-500 max-w-180 truncate">{{ item.filename }}</p>
            <p class="status-ok text-10 text-gray-400">
                <span class="text-green-500 mr-4">90%</span>
                <span>{{ formatFileSize(item.originSize) }}</span>
                <ArrowRightOutlined class="scale-75 pt-1 mx-2"/>
                <span>{{ formatFileSize(item.size) }}</span>
            </p>
            <!-- <p class="status-ok text-10 text-gray-400">
                <span>等待压缩</span>
            </p> -->
        </div>
        <div class="flex space-x-2 scale-75 -mr-10">
          <a-button type="text" disabled shape="circle" class="flex justify-center items-center">
              <template #icon>
                  <ReadOutlined class="text-15 text-gray-500" />
              </template>
          </a-button>
          <a-button type="text" shape="circle" class="flex justify-center items-center">
              <template #icon>
                  <FolderOpenOutlined class="text-15 text-gray-500" />
              </template>
          </a-button>
        </div>
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