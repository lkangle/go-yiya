<script setup>
const props = defineProps(["icon", "disabled"])

</script>

<template>
  <a-button :disabled="props.disabled" type="text" shape="circle" class="flex justify-center items-center">
    <template #icon>
      <component :is="props.icon" :class="['text-15', props.disabled ? 'text-gray-300':'text-gray-500']"/>
    </template>
  </a-button>
</template>
