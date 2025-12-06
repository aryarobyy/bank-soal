<template>
  <div class="w-full p-6 sm:p-8 mx-auto bg-white rounded-2xl shadow-xl">
    
    <div class="flex justify-between items-center mb-6 pb-4 border-b border-gray-200">
      <h2 class="text-2xl font-bold text-dark-text">{{ pageTitle }}</h2>

      <router-link
        v-if="isFiltered"
        :to="{ name: isAdminRoute ? 'AdminSoalHome' : 'DosenSoalHome' }"
        class="flex items-center gap-2 px-4 py-2 font-semibold text-gray-700 transition-colors rounded-lg bg-gray-200 hover:bg-gray-300"
      >
        <i class="fas fa-arrow-left"></i> Kembali ke Grup Soal
      </router-link>

      <router-link 
        v-else
        :to="{ name: isAdminRoute ? 'AdminSoalCreate' : 'DosenSoalCreate' }"
        class="flex items-center gap-2 px-4 py-2 font-semibold text-white transition-colors rounded-lg bg-blue-600 hover:bg-blue-700"
      >
        <i class="fas fa-plus-circle"></i> Tambah Soal
      </router-link>
    </div>

    <div class="overflow-x-auto">
      <div v-if="loading" class="text-center p-10">Memuat daftar soal...</div>
      <table v-else class="w-full text-left table-auto">
        <thead>
          <tr class="border-b border-gray-200">
            <th class="p-3 font-semibold text-gray-500">ID</th>
            <th class="p-3 font-semibold text-gray-500">Question</th>
            <th class="p-3 font-semibold text-gray-500">Correct Answer</th>
            <th class="p-3 font-semibold text-gray-500">Difficulty</th>
            <th class="p-3 font-semibold text-gray-500">Action</th>
          </tr>
        </thead>
        <tbody>
          <tr v-if="soalList.length === 0">
            <td colspan="5" class="text-center py-10 text-gray-500">
              {{ isFiltered ? 'Tidak ada soal untuk filter ini.' : 'Belum ada soal yang dibuat.' }}
            </td>
          </tr>
          <tr v-for="soal in soalList" :key="soal.id" class="border-b border-gray-100 last:border-b-0 hover:bg-gray-50">
            <td class="p-3 text-medium-text">{{ soal.id }}</td>
            <td class="p-3 text-medium-text">{{ soal.question_text }}</td>
            <td class="p-3 font-semibold text-medium-text">
              {{ soal.options?.find(opt => opt.is_correct)?.option_text || 'N/A' }}
            </td>
            <td class="p-3 text-medium-text">{{ soal.difficulty }}</td>
            <td class="p-3">
              <div class="flex items-center gap-4 text-lg">
                <button @click="editSoal(soal.id)" class="text-green-500 hover:text-green-700" title="Edit Soal">
                  <i class="fas fa-pencil-alt"></i>
                </button>
                <button @click="handleDeleteSoal(soal.id)" class="text-red-500 hover:text-red-700" title="Delete Soal">
                  <i class="fas fa-trash"></i>
                </button>
              </div>
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <div v-if="totalSoal > limit" class="flex justify-between items-center mt-6 pt-4 border-t">
      <button 
        @click="prevPage" 
        :disabled="currentPage === 1"
        class="px-4 py-2 text-sm font-medium text-gray-700 bg-white border rounded-md hover:bg-gray-50 disabled:opacity-50 disabled:cursor-not-allowed"
      >
        Sebelumnya
      </button>
      <span class="text-sm text-gray-600">
        Halaman {{ currentPage }} dari {{ totalPages }}
      </span>
      <button 
        @click="nextPage" 
        :disabled="currentPage === totalPages"
        class="px-4 py-2 text-sm font-medium text-gray-700 bg-white border rounded-md hover:bg-gray-50 disabled:opacity-50 disabled:cursor-not-allowed"
      >
        Berikutnya
      </button>
    </div>

  </div>
</template>

<script>
import { getmanyQuestions, deleteQuestion, getQuestionsBySubject, getQuestionsByCreator } from '../../provider/question.provider';
// Import hook usePopup
import { usePopup } from '../../hooks/usePopup';

export default {
  name: 'LecturerSoalList',
 
  setup() {
    const { showSuccess, showError, showConfirm } = usePopup();
    return { showSuccess, showError, showConfirm };
  },
  data() {
    return {
      soalList: [],
      loading: true,
      subjectId: null, 
      creatorId: null,
      currentPage: 1, 
      limit: 10,
      totalSoal: 0,
    };
  },
  computed: {
    isAdminRoute() {
      return this.$route.path.startsWith('/admin/soal');
    },
    isFiltered() {
      return !!this.$route.query.subject_id || !!this.$route.query.creator_id;
    },
    pageTitle() {
      if (this.$route.query.subject_id) return 'Daftar Soal per Subjek';
      if (this.$route.query.creator_id) return 'Soal yang Dibuat Pengguna';
      return 'Daftar Semua Soal'; 
    },
    totalPages() {
      const pages = Math.ceil(this.totalSoal / this.limit);
      return pages > 0 ? pages : 1;
    }
  },
  methods: {
  
    async handleRouteChange(isFirstLoad = false) {
      const query = this.$route.query;
      
      if (query.page) {
        this.loading = true;
        this.currentPage = parseInt(query.page, 10) || 1;
        
        const newQuery = { ...query };
        delete newQuery.page;
        this.$router.replace({ query: newQuery });
        
        this.fetchSoalList(); 
    
      } else if (query.show_last_page === 'true') {
        this.loading = true;
        this.$router.replace({ query: {} }); 
        try {
          const response = await getmanyQuestions(1, 0);
          this.totalSoal = response.total || 0;
          this.currentPage = this.totalPages; 
          this.fetchSoalList();
        } catch (error) {
          console.error("Gagal pre-fetch total soal:", error);
          this.currentPage = 1;
          this.fetchSoalList();
        }
        
      } else {
        if (!isFirstLoad) {
          this.currentPage = 1;
        }
        this.fetchSoalList();
      }
    },

    async fetchSoalList() {
      this.loading = true;
      this.subjectId = this.$route.query.subject_id;
      this.creatorId = this.$route.query.creator_id;

      try {
        let response;
        const offset = (this.currentPage - 1) * this.limit; 

        if (this.subjectId) {
          response = await getQuestionsBySubject(this.subjectId, this.limit, offset);
        } else if (this.creatorId) {
          response = await getQuestionsByCreator(this.creatorId, this.limit, offset);
        } else {
          response = await getmanyQuestions(this.limit, offset);
        }

        this.soalList = response.data || [];
        this.totalSoal = response.total || 0;

      } catch (error) {
        console.error("Gagal mengambil daftar soal:", error);
        this.showError('Gagal', 'Gagal memuat daftar soal.');
        this.soalList = [];
        this.totalSoal = 0;
      } finally {
        this.loading = false;
      }
    },
    
    nextPage() {
      if (this.currentPage < this.totalPages) {
        this.currentPage++;
        this.fetchSoalList();
      }
    },
    prevPage() {
      if (this.currentPage > 1) {
        this.currentPage--;
        this.fetchSoalList();
      }
    },

    editSoal(id) {
      const routeName = this.isAdminRoute ? 'AdminSoalEdit' : 'DosenSoalEdit';
      const query = {};
      
      if (this.subjectId) {
        query.return_subject_id = this.subjectId;
      }
      
      query.return_page = this.currentPage;
      
      this.$router.push({ 
        name: routeName, 
        params: { id: id },
        query: query 
      });
    },

    async handleDeleteSoal(id) {

      const isConfirmed = await this.showConfirm(
        'Konfirmasi Hapus',
        `Apakah Anda yakin ingin menghapus soal ID: ${id}?`,
        'Ya, Hapus'
      );

      if (isConfirmed) {
        try {
          await deleteQuestion(id);
          
  
          this.showSuccess('Berhasil', 'Soal berhasil dihapus dari database!');

          if (this.soalList.length === 1 && this.currentPage > 1) {
            this.currentPage--;
          }
          
          this.fetchSoalList(); 
        } catch (error) {
          console.error("Gagal menghapus soal:", error);
        
          const backendMsg = error.response?.data?.message || "";
          
          if (backendMsg.includes("foreign key constraint fails") || backendMsg.includes("Cannot delete or update")) {
              this.showError(
                "Gagal Menghapus",
                "Soal ini tidak bisa dihapus karena sedang digunakan di dalam Ujian.",
                "Silakan hapus soal ini dari Ujian terlebih dahulu."
              );
          } else {
              this.showError("Gagal Menghapus", backendMsg || "Terjadi kesalahan server");
          }
        }
      }
    }
  },

  created() {
    this.handleRouteChange(true);
  },
  
  watch: {
    '$route.query'(newQuery, oldQuery) {
      const isProgrammaticChange = (oldQuery.show_last_page && !newQuery.show_last_page) || 
                                   (oldQuery.page && !newQuery.page);
      
      if (!isProgrammaticChange) {
        this.handleRouteChange(false); 
      }
    }
  }
};
</script>