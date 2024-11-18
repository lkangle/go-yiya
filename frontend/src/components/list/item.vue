<script setup>
import { ArrowRightOutlined, FolderOpenOutlined, ReadOutlined } from '@ant-design/icons-vue';
import { OpenFile } from '@wailsjs/go/services/fileService';
import { formatFileSize, lessRate } from '@utils/index';
import Image from '../image.vue';
import useCompress from '@/store/use-compress';
import useCurrent from '@/store/use-current';
import { computed } from 'vue';

const { item } = defineProps(['item'])
const status = computed(() => item.result.status)

// 执行压缩
{
  const emitCompress = useCompress()
  emitCompress(item, (it) => {
    Object.assign(item.result, it)
  })
}

const curStore = useCurrent()
const select = () => {
  // 未压缩完成不能选中
  if (item.result.status<3) {
    return
  }
  curStore.updateCurrent(item)
}

// 压缩后是不是变大了
const isBig =  computed(() => {
  return item.result.status === 3 && item.result.size != null && item.result.size > item.size
})
// 是不是被选中的
const isSelected = computed(() => {
  let sid = curStore.current.id
  return sid === item.id
})

const open = async () => {
  await OpenFile(item.path)
}
</script>

<template>
    <div @click="select" class="img-item cursor-pointer h-45 flex items-center justify-between px-15 odd:bg-gray-50" :data-selected="isSelected">
        <div className="status-icon flex items-center justify-center">
            <i :data-status="status" :data-big="isBig"/>
        </div>
        <div class="w-35 h-35 rounded-[5px] overflow-hidden mx-12 flex items-center justify-center">
          <Image class="object-contain w-full h-full" :src="item.path"/>
        </div>
        <div class="leading-14 flex-1">
            <p class="text-11 text-gray-500 max-w-180 truncate">{{ item.filename }}</p>
            <p v-if="status<2" class="text-10 text-gray-400">
                <span>等待压缩</span>
            </p>
            <p v-if="status===2"  class="text-10 text-gray-400">
                <span>压缩中...</span>
            </p>
            <p v-if="status===3" class="text-10 text-gray-400">
                <span :class="['mr-4', {
                  'text-green-500': !isBig,
                  'text-yellow-600': isBig
                }]">{{ lessRate(item, item.result) }}</span>
                <span>{{ formatFileSize(item.size) }}</span>
                <ArrowRightOutlined class="scale-75 pt-1 mx-2"/>
                <span>{{ formatFileSize(item.result.size) }}</span>
            </p>
            <p v-if="status===4" class="text-10 text-red-600">
                <span>压缩失败</span>
            </p>
        </div>
        <div class="flex space-x-2 scale-75 -mr-10">
          <a-button type="text" disabled shape="circle" class="flex justify-center items-center">
              <template #icon>
                  <ReadOutlined class="text-15 text-gray-500" />
              </template>
          </a-button>
          <a-button @click="open" type="text" shape="circle" class="flex justify-center items-center">
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

.img-item {
  &[data-selected='true'] {
    @apply !bg-indigo-100;
  }
}

/* 前面的状态小icon */
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

        &[data-big='true'] {
          @apply !bg-yellow-600;
        }
      }
      
      /* 失败 */
      &[data-status='4'] {
        @apply bg-red-600;
      }
    }
}
</style>