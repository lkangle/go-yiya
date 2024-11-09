<script setup>
import List from './components/list.vue';
import Toolbar from './components/toolbar.vue';

</script>

<template>
  <div class="layout h-full flex flex-col justify-between">
    <Toolbar/>
    <div class="flex-1 h-12 bg-[#f8f9fb]">
      <List/>
    </div>
    <div class="footer bg-[#fff] border-t-2 border-dotted border-[#9f9f9f] w-full h-42"></div>
  </div>
</template>

<style>
</style>
