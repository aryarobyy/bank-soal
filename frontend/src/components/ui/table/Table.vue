<template>
  <div class="bg-white rounded-lg shadow-sm border border-gray-200 overflow-hidden">
    <div class="overflow-x-auto">
      <table class="min-w-full divide-y divide-gray-200">
        <thead class="bg-gray-50">
          <tr>
            <th
              v-for="header in headers"
              :key="header.key"
              scope="col"
              class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"
              :class="header.class"
            >
              {{ header.label }}
            </th>
          </tr>
        </thead>
        <tbody class="bg-white divide-y divide-gray-200">
          <!-- Loading State (Skeleton) -->
          <tr v-if="loading" v-for="i in 5" :key="'skeleton-' + i">
            <td
              v-for="header in headers"
              :key="'skeleton-' + i + '-' + header.key"
              class="px-6 py-4"
            >
              <div class="animate-pulse">
                <div class="h-4 bg-gray-200 rounded" :style="{ width: Math.floor(Math.random() * 40 + 50) + '%' }"></div>
              </div>
            </td>
          </tr>

          <!-- Empty State -->
          <tr v-else-if="items.length === 0">
            <td :colspan="headers.length" class="px-6 py-10 text-center text-gray-500">
              <slot name="empty">
                <p class="text-lg">{{ emptyText }}</p>
              </slot>
            </td>
          </tr>

          <!-- Data Rows -->
          <tr
            v-else
            v-for="(item, index) in items"
            :key="item.id || index"
            class="hover:bg-gray-50 transition-colors"
          >
            <td
              v-for="header in headers"
              :key="`${index}-${header.key}`"
              class="px-6 py-4 whitespace-nowrap text-sm text-gray-900"
              :class="header.tdClass"
            >
              <slot
                :name="`cell-${header.key}`"
                :item="item"
                :index="index"
                :value="item[header.key]"
              >
                {{ item[header.key] }}
              </slot>
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <!-- Pagination / Footer Slot -->
    <div v-if="$slots.footer" class="bg-gray-50 px-6 py-3 border-t border-gray-200">
      <slot name="footer" />
    </div>
  </div>
</template>

<script setup>
defineProps({
  headers: {
    type: Array, // Array<{ key: string, label: string, class?: string, tdClass?: string }>
    required: true,
  },
  items: {
    type: Array,
    default: () => [],
  },
  loading: {
    type: Boolean,
    default: false,
  },
  emptyText: {
    type: String,
    default: 'Tidak ada data',
  },
});
</script>
