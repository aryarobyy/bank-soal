<template>
  <Transition name="fade-slide">
    <div
      v-if="toast.show"
      class="fixed bottom-4 right-4 z-50 max-w-sm w-full bg-white rounded-lg shadow-lg border border-gray-200 transform transition-all duration-300 ease-in-out"
      :class="toast.type === 'success' ? 'border-l-4 border-l-green-500' : 'border-l-4 border-l-red-500'"
    >
      <div class="p-4">
        <div class="flex items-start">
          <div class="flex-shrink-0">
            <CheckCircle v-if="toast.type === 'success'" class="w-5 h-5 text-green-500" />
            <XCircle v-else class="w-5 h-5 text-red-500" />
          </div>
          <div class="ml-3 w-0 flex-1 pt-0.5">
            <p class="text-sm font-medium text-gray-900">{{ toast.title }}</p>
            <p class="mt-1 text-sm text-gray-500">{{ toast.message }}</p>
          </div>
          <div class="ml-4 flex-shrink-0 flex">
            <button @click="hideToast" class="inline-flex text-gray-400 hover:text-gray-500 focus:outline-none">
              <X class="h-5 w-5" />
            </button>
          </div>
        </div>
      </div>
    </div>
  </Transition>
</template>

<script setup lang="ts">
import { ref, reactive } from 'vue'
import { CheckCircle, XCircle, X } from 'lucide-vue-next'

const toast = reactive({
  show: false,
  type: 'success', // 'success' | 'error'
  title: '',
  message: '',
})

let timeout: ReturnType<typeof setTimeout>

function showToast(type: 'success' | 'error', title: string, message: string, duration = 3000) {
  toast.type = type
  toast.title = title
  toast.message = message
  toast.show = true

  clearTimeout(timeout)
  timeout = setTimeout(() => {
    toast.show = false
  }, duration)
}

function hideToast() {
  toast.show = false
}

defineExpose({ showToast }) 
</script>

<style scoped>
.fade-slide-enter-active,
.fade-slide-leave-active {
  transition: all 0.3s ease;
}
.fade-slide-enter-from,
.fade-slide-leave-to {
  opacity: 0;
  transform: translateY(20px);
}
</style>
