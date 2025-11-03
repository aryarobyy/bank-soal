<template>
  <router-view></router-view>
</template>

<script setup lang="ts">
import { onMounted } from 'vue';
import { useLocalStorage } from './hooks/useLocalStorage';
import { getUserById } from './provider/user.provider';
import { provideUser } from './hooks/useGetCurrentUser';

const { setUser, loading } = provideUser()

onMounted(async () => {
  const { value, removeValue } = useLocalStorage("id")
  const userData = await getUserById(value.value)
    if (value.value) {
    loading.value = true
      try {
        setUser(userData.data)
      } catch (err) {
        console.error(err)
        removeValue()
      } finally {
        loading.value = false
      }
    }
  })
</script>