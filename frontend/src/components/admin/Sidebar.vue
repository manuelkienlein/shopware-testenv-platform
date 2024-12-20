<script setup>
import Menu from 'primevue/menu';
import {ref} from "vue";

const items = ref([
  { label: 'Sandbox Environments', icon: 'pi pi-box', route: '/admin/sandbox-environments' },
  { label: 'Sandbox Images', icon: 'pi pi-image', route: '/admin/sandbox-images' }
]);
</script>

<template>
  <Menu :model="items">
    <template #item="{ item, props }">
      <router-link v-if="item.route" v-slot="{ href, navigate }" :to="item.route" custom>
        <a v-ripple :href="href" v-bind="props.action" @click="navigate">
          <span :class="item.icon" />
          <span class="ml-2">{{ item.label }}</span>
        </a>
      </router-link>
      <a v-else v-ripple :href="item.url" :target="item.target" v-bind="props.action">
        <span :class="item.icon" />
        <span class="ml-2">{{ item.label }}</span>
      </a>
    </template>
  </Menu>
</template>

<style scoped>

</style>