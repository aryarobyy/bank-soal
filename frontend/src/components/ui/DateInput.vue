<template>
  <div class="flex flex-col space-y-1">
    <label v-if="label" :for="id" class="text-sm font-medium text-gray-700">{{ label }}</label>
    <input
      type="date"
      :id="id"
      :name="name"
      v-model="internalValue"
      :class="[
        'border rounded px-3 py-2 focus:outline-none focus:ring-2 focus:ring-blue-500',
      ]"
    />
  </div>
</template>

<script setup>
import { computed } from 'vue'

const props = defineProps({
  modelValue: {
    type: String,
    default: null
  },
  label: String,
  name: String,
  id: String
})

const emit = defineEmits(['update:modelValue'])

const internalValue = computed({
  get: () => props.modelValue,
  set: (val) => emit('update:modelValue', val)
})
</script>
