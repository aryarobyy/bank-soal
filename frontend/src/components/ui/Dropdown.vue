<template>
  <div class="flex flex-col space-y-1 relative" ref="dropdownRef">
    <label v-if="label" :for="id" class="text-sm font-medium text-gray-700">{{ label }}</label>

    <div
      class="border rounded px-3 py-2 bg-white focus:outline-none focus:ring-2 focus:ring-blue-500 cursor-pointer border-gray-300"
      @click="toggleDropdown"
    >
      <div class="flex flex-row justify-between items-center">
        <span class="text-gray-900">
          {{ displayText }}
        </span>
        <ChevronDown class="w-4 h-4 text-gray-500" :class="{ 'rotate-180': isOpen }" />
      </div>
    </div>

    <div
      v-if="isOpen"
      class="absolute top-full left-0 right-0 z-50 mt-1 bg-white border border-gray-300 rounded-md shadow-lg max-h-60 overflow-auto"
      style="min-width: 100%;"
    >
      <div
        v-for="(option, index) in options"
        :key="index"
        class="flex items-center px-3 py-2 hover:bg-gray-100 cursor-pointer text-gray-900"
        @click.stop="toggleSelect(option)"
      >
        <input
          v-if="multiple"
          type="checkbox"
          class="mr-3 w-4 h-4"
          :checked="isSelected(option)"
          @change.prevent
        />
        <input
          v-else
          type="radio"
          class="mr-3 w-4 h-4"
          :checked="isSelected(option)"
          @change.prevent
        />
        <span class="flex-1">{{ option[labelKey] }}</span>
      </div>
    </div>

    <p v-if="error" class="text-xs text-red-500 mt-1">{{ error }}</p>
  </div>
</template>

<script setup>
import { ChevronDown } from 'lucide-vue-next'
import { ref, computed, onMounted, onBeforeUnmount } from 'vue'

const props = defineProps({
  modelValue: { type: [Array, Object, String, Number, Boolean], default: null },
  options: { type: Array, required: true },
  label: String,
  name: String,
  id: String,
  error: String,
  valueKey: { type: String, required: true },
  labelKey: { type: String, required: true },
  multiple: { type: Boolean, default: false }
})

const emit = defineEmits(['update:modelValue'])

const isOpen = ref(false)
const dropdownRef = ref(null)

const selectedItems = computed(() => {
  if (props.multiple) {
    return Array.isArray(props.modelValue) ? props.modelValue : []
  } else {
    return props.modelValue ? [props.modelValue] : []
  }
})

const displayText = computed(() => {
  if (props.multiple) {
    const validItems = selectedItems.value.filter(item =>
      props.options.some(option => option[props.valueKey] === item[props.valueKey])
    )
    return validItems.length > 0
      ? validItems.map(item => item[props.labelKey]).join(', ')
      : '-- Pilih --'
  } else {
    const selected = selectedItems.value[0]
    const match = props.options.find(option => option[props.valueKey] === selected?.[props.valueKey])
    return match ? match[props.labelKey] : '-- Pilih --'
  }
})

const toggleDropdown = () => {
  isOpen.value = !isOpen.value
}

const isSelected = (item) => {
  return selectedItems.value.some(selected => selected[props.valueKey] === item[props.valueKey])
}

const toggleSelect = (item) => {
  if (props.multiple) {
    const exists = isSelected(item)
    let newSelected

    if (exists) {
      newSelected = selectedItems.value.filter(
        (i) => i[props.valueKey] !== item[props.valueKey]
      )
    } else {
      newSelected = [...selectedItems.value, item]
    }

    emit('update:modelValue', newSelected)
  } else {
    emit('update:modelValue', item)
    isOpen.value = false
  }
}

const handleClickOutside = (event) => {
  if (dropdownRef.value && !dropdownRef.value.contains(event.target)) {
    isOpen.value = false
  }
}

onMounted(() => {
  document.addEventListener('click', handleClickOutside)
})

onBeforeUnmount(() => {
  document.removeEventListener('click', handleClickOutside)
})
</script>

<style scoped>
.rotate-180 {
  transform: rotate(180deg);
}
</style>
