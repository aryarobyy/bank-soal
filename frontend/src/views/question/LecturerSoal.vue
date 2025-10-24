<template>
  <div v-if="loading" class="text-center p-10">
    <p>Memuat data bank soal...</p>
  </div>
  <SoalManagement
    v-else
    :soal-data="lecturerSoalData"
    @buat-soal="handleBuatSoal"
    @view-details="handleViewDetails"
    @delete-soal="handleDeleteSoal"
  />
</template>

<script>
import SoalManagement from './components/SoalManagement.vue';
import { getAllQuestions, deleteQuestion } from '../../provider/question.provider';

export default {
  name: 'LecturerSoal',
  components: { SoalManagement },
  data() {
    return {
      lecturerSoalData: [],
      loading: true,
    };
  },
  methods: {
    async fetchSoalData() {
      this.loading = true;
      try {
        const response = await getAllQuestions();
        const questions = response.data || [];
        
        // ## LOGIKA PENGELOMPOKAN BARU ADA DI SINI ##
        // Mengubah daftar soal menjadi grup berdasarkan kategori
        const groupedSoal = questions.reduce((acc, soal) => {
          // Gunakan 'category' dari soal, atau 'Lainnya' jika tidak ada kategori
          const groupTitle = soal.category || 'Soal Lainnya';

          // Jika grup belum ada di akumulator, buat grup baru
          if (!acc[groupTitle]) {
            acc[groupTitle] = {
              title: groupTitle,
              items: []
            };
          }

          // Tambahkan soal ke dalam grup yang sesuai
          acc[groupTitle].items.push({ 
            id: soal.id, 
            nama: soal.question_text, 
            status: soal.difficulty 
          });

          return acc;
        }, {});

        // Ubah objek grup menjadi array agar bisa di-loop oleh v-for
        this.lecturerSoalData = Object.values(groupedSoal);

      } catch (error) {
        console.error("Gagal mengambil data soal:", error);
        alert('Gagal memuat data bank soal.');
      } finally {
        this.loading = false;
      }
    },
    handleBuatSoal() {
      this.$router.push('/dosen/soal/create');
    },
    handleViewDetails(soalId) {
      this.$router.push('/dosen/soal/list');
    },
    async handleDeleteSoal(soalId) {
      if (confirm(`Anda yakin ingin menghapus soal dengan ID: ${soalId}?`)) {
        try {
          await deleteQuestion(soalId);
          alert('Soal berhasil dihapus!');
          this.fetchSoalData(); // Muat ulang data
        } catch (error) {
          alert('Gagal menghapus soal.');
        }
      }
    }
  },
  mounted() {
    this.fetchSoalData();
  }
};
</script>