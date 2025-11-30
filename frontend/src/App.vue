<template>
  <router-view></router-view>
</template>

<script setup lang="ts">
import { onMounted } from 'vue';
// Hapus import useLocalStorage dan getUserById karena sudah dipindah ke hook
import { provideUser } from './hooks/useGetCurrentUser';

// 1. Ambil fungsi 'fetchUser' dari provider yang baru diperbarui
const { fetchUser } = provideUser();

onMounted(async () => {
  // 2. Panggil fetchUser()
  // Fungsi ini sudah punya pengaman "if (!token) return" di dalamnya.
  // Jadi jika user sudah logout, dia tidak akan memanggil API sama sekali.
  await fetchUser();
});
</script>