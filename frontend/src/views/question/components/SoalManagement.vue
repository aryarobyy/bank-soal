<template>
  <div class="w-full p-6 sm:p-8 mx-auto bg-white rounded-2xl shadow-xl">
    <div
      class="flex items-center justify-between pb-4 mb-6 border-b border-gray-200"
    >
      <h2 class="text-2xl font-bold text-dark-text">Manajemen Bank Soal</h2>
      <button
        @click="$emit('buat-soal')"
        class="flex items-center gap-2 px-4 py-2 font-semibold text-white transition-colors rounded-lg bg-blue-600 hover:bg-blue-700"
      >
        <i class="fas fa-plus-circle"></i> Buat Soal
      </button>
    </div>

    <div class="grid grid-cols-1 gap-4 mb-8 md:grid-cols-4">
      <div class="relative col-span-1 md:col-span-4">
        <input
          type="text"
          placeholder="Cari Mata Kuliah..."
          :value="searchQuery"
          @input="$emit('update:searchQuery', $event.target.value)"
          class="w-full px-3 py-2 pr-10 bg-gray-100 border border-gray-200 rounded-md focus:outline-none focus:ring-2 focus:ring-primary"
        />
        <i class="absolute text-gray-400 fas fa-search top-3 right-4"></i>
      </div>
    </div>

    <div class="space-y-4">
      <div v-if="subjects.length === 0" class="text-center py-10 text-gray-500">
        <p>Tidak ada mata kuliah yang ditemukan.</p>
      </div>

      <div
        v-for="subject in subjects"
        :key="subject.id"
        class="p-4 bg-gray-50 rounded-lg shadow-sm border border-gray-100"
      >
        <div class="flex justify-between items-start">
          <div>
            <h3 class="text-lg font-bold text-gray-800">{{ subject.title }}</h3>
          </div>

          <div class="flex flex-col items-end">
            <span class="text-xs font-medium text-gray-500 uppercase mb-2"
              >ACTIONS</span
            >

            <div class="flex items-center gap-4 text-sm">
              <button
                @click="$emit('view-details', subject)"
                class="flex items-center gap-1 font-semibold text-gray-600 hover:text-blue-600"
              >
                <i class="fas fa-eye"></i> View Details
              </button>
              <button
                @click="$emit('delete-subject', subject)"
                class="flex items-center gap-1 font-semibold text-red-500 hover:text-red-700"
              >
                <i class="fas fa-trash"></i> Hapus
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "SoalManagement",
  props: {
    subjects: {
      type: Array,
      required: true,
    },
    searchQuery: {
      type: String,
      default: "",
    },
  },
  emits: ["buat-soal", "view-details", "delete-subject", "update:searchQuery"],
};
</script>
