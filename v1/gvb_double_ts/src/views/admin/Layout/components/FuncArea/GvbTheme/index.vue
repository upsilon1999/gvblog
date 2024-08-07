<template>
  <div @click="settingStore.setTheme()" class="action-icon">
    <IconSun v-if="settingStore.getTheme()"></IconSun>
    <IconMoon v-else></IconMoon>
  </div>
</template>

<script setup lang="ts">
import {IconSun, IconMoon} from "@arco-design/web-vue/es/icon";
import useSettingStore from '@/stores/modules/settings';

const settingStore = useSettingStore()

</script>

<style lang="scss">
.action-icon {
	margin-right: 10px;
	font-size: 15px;
	cursor: pointer;
}
</style>