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
            <td colspan="5" class="text-center py-10 text-gray-500">Belum ada soal yang dibuat.</td>
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

    <div v-if="!isFiltered && totalSoal > limit" class="flex justify-between items-center mt-6 pt-4 border-t">
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
// ## 3. Impor 'getmanyQuestions' dan 'getQuestionsByCreator' ##
import { getmanyQuestions, deleteQuestion, getQuestionsBySubject, getQuestionsByCreator } from '../../provider/question.provider';

export default {
  name: 'LecturerSoalList',
  data() {
    return {
      soalList: [],
      loading: true,
      subjectId: null, 
      creatorId: null, // <-- State untuk ID kreator
      currentPage: 1,
      limit: 10, // Tampilkan 10 soal per halaman
      totalSoal: 0,
    };
  },
  computed: {
    isAdminRoute() {
      return this.$route.path.startsWith('/admin/soal');
    },
    // ## 4. 'isFiltered' diperbarui ##
    isFiltered() {
      // Cek apakah ada query param subject_id ATAU creator_id
      return !!this.$route.query.subject_id || !!this.$route.query.creator_id;
    },
    pageTitle() {
      // Judul halaman dinamis
      if (this.$route.query.subject_id) return 'Daftar Soal per Subjek';
      if (this.$route.query.creator_id) return 'Soal yang Baru Dibuat';
      return 'Daftar Semua Soal';
    },
    totalPages() {
      return Math.ceil(this.totalSoal / this.limit);
    }
  },
  methods: {
    // ## 5. 'fetchSoalList' diperbarui total ##
    async fetchSoalList() {
      this.loading = true;
      this.subjectId = this.$route.query.subject_id;
      this.creatorId = this.$route.query.creator_id; // <-- Ambil creator_id

      try {
        let response;
        if (this.subjectId) {
          // MODE 1: Filter berdasarkan Subjek (permintaan "view detail")
          response = await getQuestionsBySubject(this.subjectId);
          this.soalList = response.data || [];
          this.totalSoal = response.data.length;
        } else if (this.creatorId) {
          // MODE 2: Filter berdasarkan Kreator (permintaan "simpan selesai")
          response = await getQuestionsByCreator(this.creatorId);
          this.soalList = response.data || [];
          this.totalSoal = response.data.length;
        } else {
          // MODE 3: Paginasi (default)
          const offset = (this.currentPage - 1) * this.limit;
          response = await getmanyQuestions(this.limit, offset);
          this.soalList = response.data.data || [];
          this.totalSoal = response.data.total || 0;
        }
      } catch (error) {
        console.error("Gagal mengambil daftar soal:", error);
        alert('Gagal memuat daftar soal.');
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
      this.$router.push({ name: routeName, params: { id: id } });
    },

    async handleDeleteSoal(id) {
      if (confirm(`Apakah Anda yakin ingin menghapus soal ID: ${id}?`)) {
        try {
          await deleteQuestion(id);
          alert('Soal berhasil dihapus dari database!');
          this.fetchSoalList(); // Muat ulang data
        } catch (error) {
          console.error("Gagal menghapus soal:", error);
          alert('Gagal menghapus soal.');
        }
      }
    }
  },
  created() {
    this.fetchSoalList();
  },
  // ## 6. 'watch' diperbarui ##
  // untuk mendeteksi perubahan filter APAPUN
  watch: {
    '$route.query'() {
      this.currentPage = 1; // Reset halaman saat filter berubah
      this.fetchSoalList();
    }
  }
};
</script>