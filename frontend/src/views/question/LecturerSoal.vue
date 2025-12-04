<template>
  <div v-if="loading" class="text-center p-10">
    <p>Memuat data bank soal...</p>
  </div>
  
  <SoalManagement
    v-else
    :subjects="paginatedSubjects" :search-query="searchQuery"
    @update:searchQuery="searchQuery = $event"
    @buat-soal="handleBuatSoal"
    @view-details="handleViewDetails"
    @delete-subject="handleDeleteSubject" 
  />

  <div v-if="!loading && totalPages > 1" class="flex justify-center items-center space-x-2 mt-8">
    <button
      @click="prevPage"
      :disabled="currentPage === 1"
      class="px-3 py-1 rounded-md"
      :class="currentPage === 1 ? 'text-gray-400 cursor-not-allowed' : 'text-gray-700 hover:bg-gray-100'"
    >
      <i class="fas fa-chevron-left"></i>
    </button>
    
    <button
      v-for="page in totalPages"
      :key="page"
      @click="goToPage(page)"
      class="w-8 h-8 text-sm font-medium rounded-md"
      :class="page === currentPage ? 'bg-blue-600 text-white' : 'bg-white text-gray-700 hover:bg-gray-100 border'"
    >
      {{ page }}
    </button>
    
    <button
      @click="nextPage"
      :disabled="currentPage === totalPages"
      class="px-3 py-1 rounded-md"
       :class="currentPage === totalPages ? 'text-gray-400 cursor-not-allowed' : 'text-gray-700 hover:bg-gray-100'"
    >
      <i class="fas fa-chevron-right"></i>
    </button>
  </div>
</template>

<script>
import SoalManagement from './components/SoalManagement.vue';
import { getPaginatedSubjects, deleteSubject } from '../../provider/subject.provider';

export default {
  name: 'LecturerSoal',
  components: { SoalManagement },
  data() {
    return {
      allSubjects: [], 
      totalSubjects: 0, 
      loading: true,
      currentPage: 1,
      itemsPerPage: 10,
      searchQuery: '',
    
    };
  },
  watch: {
    searchQuery() {
      this.currentPage = 1;
      
    },
  },
  computed: {
    isAdminRoute() {
      return this.$route.path.startsWith('/admin/soal');
    },
    
 
    filteredSubjects() {
      if (!this.searchQuery) {
        return this.allSubjects; 
      }
      const searchLower = this.searchQuery.toLowerCase();
      return this.allSubjects.filter(subject => 
        subject.title.toLowerCase().includes(searchLower) 
      );
    },

    totalPages() {
     
      return Math.ceil(this.filteredSubjects.length / this.itemsPerPage);
    },
    
 
    paginatedSubjects() {
    
      const start = (this.currentPage - 1) * this.itemsPerPage;
      const end = start + this.itemsPerPage;
      return this.filteredSubjects.slice(start, end);
    }
  },
  methods: {
  
    async fetchSoalData() {
      this.loading = true;
      try {
      
        const response = await getPaginatedSubjects(0, 0, ''); 
        
       
        this.allSubjects = response.data.data || []; 
        this.totalSubjects = response.total || 0; 

      } catch (error) {
        console.error("Gagal mengambil data subjek:", error);
        alert('Gagal memuat data bank soal.');
      } finally {
        this.loading = false;
      }
    },
    
    handleBuatSoal() {
      const routeName = this.isAdminRoute ? 'AdminSoalCreate' : 'DosenSoalCreate';
      this.$router.push({ name: routeName });
    },

    handleViewDetails(subject) {
      const routeName = this.isAdminRoute ? 'AdminSoalList' : 'DosenSoalList';
      this.$router.push({ name: routeName, query: { subject_id: subject.id } });
    },

   
    async handleDeleteSubject(subject) {
     
      if (!confirm(`Anda yakin ingin menghapus mata kuliah "${subject.title}"?`)) {
        return;
      }
      
      try {
        await deleteSubject(subject.id); 
        alert(`Mata kuliah "${subject.title}" berhasil dihapus.`);
        this.fetchSoalData(); 
        
      } catch (error) {
        console.error("Gagal menghapus subjek:", error);
        alert('Terjadi kesalahan saat proses penghapusan.');
      }
    },

    nextPage() {
      if (this.currentPage < this.totalPages) {
        this.currentPage++;
      }
    },
    prevPage() {
      if (this.currentPage > 1) {
        this.currentPage--;
      }
    },
    goToPage(page) {
      this.currentPage = page;
    }
  },
  mounted() {
    this.fetchSoalData();
  }
};
</script>