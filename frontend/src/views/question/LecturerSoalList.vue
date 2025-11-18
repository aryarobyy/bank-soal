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
// Pastikan path provider Anda sudah benar
import { getmanyQuestions, deleteQuestion, getQuestionsBySubject, getQuestionsByCreator } from '../../provider/question.provider';

export default {
  name: 'LecturerSoalList',
  data() {
    return {
      soalList: [],
      loading: true,
      subjectId: null, 
      creatorId: null,
      currentPage: 1, // Akan di-set oleh 'handleRouteChange'
      limit: 10,
      totalSoal: 0,
    };
  },
  computed: {
    isAdminRoute() {
      return this.$route.path.startsWith('/admin/soal');
    },
    isFiltered() {
      // 'show_last_page' dan 'page' tidak dihitung sebagai filter visual
      return !!this.$route.query.subject_id || !!this.$route.query.creator_id;
    },
    pageTitle() {
      if (this.$route.query.subject_id) return 'Daftar Soal per Subjek';
      if (this.$route.query.creator_id) return 'Soal yang Dibuat Pengguna';
      return 'Daftar Semua Soal'; // Judul default
    },
    totalPages() {
      const pages = Math.ceil(this.totalSoal / this.limit);
      return pages > 0 ? pages : 1;
    }
  },
  methods: {
    // ## METHOD UTAMA UNTUK MENGATUR NAVIGASI & DATA ##
    async handleRouteChange(isFirstLoad = false) {
      const query = this.$route.query;
      
      // PRIORITAS 1: Datang dari 'Edit Soal'
      if (query.page) {
        this.loading = true;
        this.currentPage = parseInt(query.page, 10) || 1;
        
        // Bersihkan query 'page' dari URL, tapi pertahankan query filter
        const newQuery = { ...query };
        delete newQuery.page;
        this.$router.replace({ query: newQuery });
        
        this.fetchSoalList(); // Ambil data untuk halaman yang dituju

      // PRIORITAS 2: Datang dari 'Buat Soal'
      } else if (query.show_last_page === 'true') {
        this.loading = true;
        this.$router.replace({ query: {} }); // Hapus semua query
        try {
          // Ambil total untuk hitung halaman terakhir
          const response = await getmanyQuestions(1, 0);
          this.totalSoal = response.total || 0;
          this.currentPage = this.totalPages; // Set ke halaman terakhir
          this.fetchSoalList(); // Ambil data halaman terakhir
        } catch (error) {
          console.error("Gagal pre-fetch total soal:", error);
          this.currentPage = 1;
          this.fetchSoalList();
        }
        
      // PRIORITAS 3: Alur Normal (Filter atau Load Pertama Kali)
      } else {
        // Jika ini BUKAN load pertama (artinya user klik filter), reset ke hal 1
        if (!isFirstLoad) {
          this.currentPage = 1;
        }
        // Jika ini load pertama, 'currentPage' tetap 1 (default)
        this.fetchSoalList();
      }
    },

    // 'fetchSoalList' hanya mengambil data berdasarkan state saat ini
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
        alert('Gagal memuat daftar soal.');
        this.soalList = [];
        this.totalSoal = 0;
      } finally {
        this.loading = false;
      }
    },
    
    // Tombol paginasi
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

    // ## editSoal DIPERBARUI ##
    editSoal(id) {
      const routeName = this.isAdminRoute ? 'AdminSoalEdit' : 'DosenSoalEdit';
      const query = {};
      
      // Kirim filter subjek (jika ada)
      if (this.subjectId) {
        query.return_subject_id = this.subjectId;
      }
      
      // Kirim halaman saat ini
      query.return_page = this.currentPage;
      
      this.$router.push({ 
        name: routeName, 
        params: { id: id },
        query: query // query berisi return_page
      });
    },

    async handleDeleteSoal(id) {
      if (confirm(`Apakah Anda yakin ingin menghapus soal ID: ${id}?`)) {
        try {
          await deleteQuestion(id);
          alert('Soal berhasil dihapus dari database!');

          // Cek jika ini item terakhir di halaman terakhir
          if (this.soalList.length === 1 && this.currentPage > 1) {
            this.currentPage--;
          }
          
          this.fetchSoalList(); // Muat ulang data halaman saat ini
        } catch (error) {
          console.error("Gagal menghapus soal:", error);
          alert('Gagal menghapus soal.');
        }
      }
    }
  },
  
  // ## created DIPERBARUI ##
  created() {
    // Panggil controller utama, tandai sebagai 'load pertama'
    this.handleRouteChange(true);
  },
  
  // ## watch DIPERBARUI ##
  watch: {
    '$route.query'(newQuery, oldQuery) {
      // Cek apakah perubahan query ini dari 'replace' programatik kita
      const isProgrammaticChange = (oldQuery.show_last_page && !newQuery.show_last_page) || 
                                   (oldQuery.page && !newQuery.page);
      
      // Jika BUKAN perubahan programatik, jalankan handle (ini filter)
      if (!isProgrammaticChange) {
        // Kirim 'false' (bukan load pertama)
        this.handleRouteChange(false); 
      }
    }
  }
};
</script>