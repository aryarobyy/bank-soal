<template>
  <router-view></router-view>
</template>

<script setup lang="ts">
import { onMounted } from 'vue';

import { provideUser } from './hooks/useGetCurrentUser';


const { fetchUser } = provideUser();

onMounted(async () => {

  await fetchUser();
});
</script>