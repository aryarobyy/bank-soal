<template>
  <div class="w-full p-6 sm:p-8 mx-auto bg-white rounded-2xl shadow-xl">
    
    <div class="flex items-center justify-between pb-4 mb-6 border-b border-gray-200">
      <h2 class="text-2xl font-bold text-dark-text">Manajemen Bank Soal</h2>
      
      <button 
        @click="$emit('buat-soal')" 
        class="flex items-center gap-2 px-4 py-2 font-semibold text-white transition-colors rounded-lg bg-blue-600 hover:bg-blue-700"
      >
        <i class="fas fa-plus-circle"></i> Buat Soal
      </button>
      </div>
    <div class="grid grid-cols-1 gap-4 mb-8 md:grid-cols-4">
      <select class="w-full px-3 py-2 bg-gray-100 border border-gray-200 rounded-md focus:outline-none focus:ring-2 focus:ring-primary">
        <option>Last Modified</option>
      </select>
      
      <select class="w-full px-3 py-2 bg-gray-100 border border-gray-200 rounded-md focus:outline-none focus:ring-2 focus:ring-primary">
        <option value="">Filter by Status</option>
        <option value="Not Started">Not Started</option>
        <option value="In Progress">In Progress</option>
        <option value="Completed">Completed</option>
      </select>
      
      <div class="relative col-span-1 md:col-span-2">
        <input type="text" placeholder="Search your file" class="w-full px-3 py-2 pr-10 bg-gray-100 border border-gray-200 rounded-md focus:outline-none focus:ring-2 focus:ring-primary" />
        <i class="absolute text-gray-400 fas fa-search top-3 right-4"></i>
      </div>
    </div>

    <div class="space-y-4">
      <div v-if="soalGroups.length === 0" class="text-center py-10 text-gray-500">
        <p>Belum ada soal di bank soal.</p>
      </div>
      <div v-for="(group, index) in soalGroups" :key="index" class="bg-indigo-50/50 rounded-lg">
        <button @click="toggleAccordion(group)" class="flex items-center justify-between w-full p-4 font-bold text-left text-dark-text">
          <span>{{ group.title }}</span>
          <i :class="group.isOpen ? 'fa-chevron-up' : 'fa-chevron-down'" class="fas"></i>
        </button>
        <div v-if="group.isOpen" class="p-4 border-t border-indigo-100">
          <table class="w-full text-left">
            <thead>
              <tr class="border-b border-gray-200">
                <th class="py-2 font-semibold text-gray-500">NAMA SOAL</th>
                <th class="py-2 font-semibold text-gray-500">STATUS</th>
                <th class="py-2 font-semibold text-gray-500">ACTIONS</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="soal in group.items" :key="soal.id" class="border-b border-gray-100 last:border-b-0">
                <td class="py-4 text-medium-text">{{ soal.nama }}</td>
                <td class="py-4 text-medium-text">{{ soal.status }}</td>
                <td class="py-4">
                  <div class="flex items-center gap-4 text-sm">
                    <button @click="$emit('view-details', soal.id)" class="flex items-center gap-1 font-semibold text-gray-500 hover:text-primary">
                      <i class="fas fa-eye"></i> View Details
                    </button>
                    <button @click="$emit('delete-soal', soal.id)" class="text-red-500 hover:text-red-700">
                      <i class="fas fa-trash"></i>
                    </button>
                  </div>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'SoalManagement',
  props: {
    soalData: {
      type: Array,
      required: true,
    }
  },
  data() {
    return {
      soalGroups: this.soalData.map(group => ({ ...group, isOpen: true }))
    };
  },
  watch: {
    // Menambahkan watcher agar data diupdate jika ada perubahan dari parent
    soalData(newData) {
      this.soalGroups = newData.map(group => ({ ...group, isOpen: true }));
    }
  },
  methods: {
    toggleAccordion(group) {
      group.isOpen = !group.isOpen;
    }
  }
};
</script>