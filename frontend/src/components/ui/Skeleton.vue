<template>
  <div class="animate-pulse bg-gray-200 rounded" :class="props.class"></div>
</template>

<script setup>
const props = defineProps({
  class: {
    type: String,
    default: "",
  },
});
</script>
