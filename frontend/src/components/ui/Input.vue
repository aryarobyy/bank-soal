<template>
  <div class="relative">
    <label
      v-if="title"
      :for="id"
      class="block mb-2 text-sm font-medium text-gray-900"
    >
      {{ title }}
    </label>

    <div class="relative">
      <component
        v-if="icon"
        :is="icon"
        class="absolute left-3 top-1/2 transform -translate-y-1/2 w-5 h-5 text-gray-400"
      />

      <input
        ref="inputRef"
        v-model="localValue"
        :id="id"
        class="bg-[#FAFAF5] border border-[#D6D3C2] text-[#3D3A2A] text-sm rounded-lg focus:ring-[#7A745C] focus:border-[#7A745C] block w-full p-2.5 pr-10"
        :class="icon ? 'pl-10' : 'pl-3'"
        :placeholder="placeHolder"
        :required="required"
        :type="currentType"
      />

      <button
        v-if="isPasswordType"
        type="button"
        @click="toggleShowPassword"
        class="absolute top-1/2 right-3 transform -translate-y-1/2 text-gray-500 dark:text-gray-300"
      >
        <EyeClosed v-if="showPassword" class="w-5 h-5" />
        <Eye v-else class="w-5 h-5" />
      </button>
    </div>
  </div>
</template>

<script setup>
import { Eye, EyeClosed } from 'lucide-vue-next'
import { ref, computed, watch } from 'vue'

const props = defineProps({
  modelValue: String,
  id: String,
  title: String,
  placeHolder: String,
  required: Boolean,
  icon: Object,
  type: {
    type: String,
    default: 'text'
  }
})

const emit = defineEmits(['update:modelValue'])

const localValue = ref(props.modelValue)
const showPassword = ref(false)
const isPasswordType = computed(() => props.type === 'password')
const currentType = ref(props.type)

watch(
  () => props.modelValue,
  (newVal) => {
    localValue.value = newVal
  }
)

watch(localValue, (val) => {
  emit('update:modelValue', val)
})

function toggleShowPassword() {
  showPassword.value = !showPassword.value
  if (isPasswordType.value) {
    currentType.value = showPassword.value ? 'text' : 'password'
  }
}
</script>
