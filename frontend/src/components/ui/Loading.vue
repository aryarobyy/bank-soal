<template>
  <div v-if="props.isLoading" class="py-6 space-y-4">
    <div v-for="i in rows" :key="i" class="animate-pulse flex items-center gap-4">
      <div v-if="showAvatar" class="rounded-full bg-gray-200 h-10 w-10 shrink-0"></div>
      <div class="flex-1 space-y-2">
        <div class="h-4 bg-gray-200 rounded w-3/4"></div>
        <div class="h-3 bg-gray-200 rounded w-1/2"></div>
      </div>
    </div>
  </div>
</template>

<script setup>
const props = defineProps({
  isLoading: {
    type: Boolean,
    default: false
  },
  rows: {
    type: Number,
    default: 3
  },
  showAvatar: {
    type: Boolean,
    default: false
  }
})
</script>
