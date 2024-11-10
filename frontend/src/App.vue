<script setup>
import { ConfigProvider, theme } from 'ant-design-vue';
import Content from './components/content.vue';
import Toolbar from './components/toolbar.vue';

</script>

<template>
  <ConfigProvider :theme="{
      algorithm: theme.defaultAlgorithm,
      token: {
        colorPrimary: '#6366f1',
      },
    }">
    <div class="layout bg-gray-100 h-full flex flex-col justify-between">
      <Toolbar/>
      <Content/>
    </div>
  </ConfigProvider>
</template>
