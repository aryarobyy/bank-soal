<template>
  <div class="w-full p-6 sm:p-8 mx-auto bg-white rounded-2xl shadow-xl">
    
    <div class="flex justify-between items-center mb-6 pb-4 border-b border-gray-200">
      <h2 class="text-2xl font-bold text-dark-text">Daftar Semua Soal</h2>
      <router-link 
        to="/dosen/soal/create" 
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
  </div>
</template>

<script>
import { getAllQuestions, deleteQuestion } from '../../provider/question.provider';

export default {
  name: 'LecturerSoalList',
  data() {
    return {
      soalList: [],
      loading: true,
    };
  },
  methods: {
    async fetchSoalList() {
      this.loading = true;
      try {
        const response = await getAllQuestions();
        this.soalList = response.data || [];
      } catch (error) {
        console.error("Gagal mengambil daftar soal:", error);
        alert('Gagal memuat daftar soal.');
      } finally {
        this.loading = false;
      }
    },
    
    // ## LOGIKA TOMBOL EDIT DIPERBARUI DI SINI ##
    editSoal(id) {
      // Mengarahkan ke halaman edit dengan membawa ID soal
      this.$router.push(`/dosen/soal/edit/${id}`);
    },

    // ## LOGIKA HAPUS SUDAH BENAR DAN TERINTEGRASI ##
    async handleDeleteSoal(id) {
      if (confirm(`Apakah Anda yakin ingin menghapus soal ID: ${id}?`)) {
        try {
          await deleteQuestion(id);
          alert('Soal berhasil dihapus dari database!');
          this.fetchSoalList(); // Muat ulang data untuk memperbarui tampilan
        } catch (error) {
          console.error("Gagal menghapus soal:", error);
          alert('Gagal menghapus soal.');
        }
      }
    }
  },
  mounted() {
    this.fetchSoalList();
  }
};
</script>