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
// ## 1. HAPUS 'deleteQuestion' (sudah ada di provider) dan impor 'getAllQuestions' ##
import { getAllQuestions, deleteQuestion } from '../../provider/question.provider';

// ## 2. BUAT MAP UNTUK SUBJECT (SAMA SEPERTI DI HALAMAN CREATE) ##
// (Anda harus melengkapi ini agar sesuai dengan data Anda)
const subjects = [
  { id: 1, title: 'Kalkulus' },
  { id: 2, title: 'Matematika Diskrit' },
  { id: 3, title: 'Teori Bahasa dan Automata' },
  { id: 4, title: 'Basis Data Lanjut' },
  { id: 5, title: 'Metode Numerik' },
];
const subjectsMap = new Map(subjects.map(s => [s.id, s.title]));


export default {
  name: 'LecturerSoal',
  components: { SoalManagement },
  data() {
    return {
      lecturerSoalData: [],
      loading: true,
    };
  },
  computed: {
    isAdminRoute() {
      return this.$route.path.startsWith('/admin/soal');
    }
  },
  methods: {
    async fetchSoalData() {
      this.loading = true;
      try {
        const response = await getAllQuestions();
        const questions = response.data || [];
        
        // ## 3. UBAH LOGIKA PENGELOMPOKAN (GROUPING) ##
        // Mengelompokkan berdasarkan subject_id, bukan category
        const groupedSoal = questions.reduce((acc, soal) => {
          // Mengambil nama subjek dari map, atau 'Lainnya' jika tidak ditemukan
          const groupTitle = subjectsMap.get(soal.subject_id) || 'Soal Lainnya';
          
          if (!acc[groupTitle]) {
            acc[groupTitle] = { title: groupTitle, items: [] };
          }
          
          // Simpan subject_id bersama data soal
          acc[groupTitle].items.push({ 
            id: soal.id, 
            nama: soal.question_text, 
            status: soal.difficulty,
            subject_id: soal.subject_id // <-- PENTING: Simpan subject_id
          });
          return acc;
        }, {});

        this.lecturerSoalData = Object.values(groupedSoal);

      } catch (error) {
        console.error("Gagal mengambil data soal:", error);
        alert('Gagal memuat data bank soal.');
      } finally {
        this.loading = false;
      }
    },
    handleBuatSoal() {
      const routeName = this.isAdminRoute ? 'AdminSoalCreate' : 'DosenSoalCreate';
      this.$router.push({ name: routeName });
    },

    // ## 4. UBAH LOGIKA "VIEW DETAILS" ##
    handleViewDetails(soalId) {
      let foundSubjectId = null;
      
      // Cari soal di dalam data yang sudah dikelompokkan
      for (const group of this.lecturerSoalData) {
        const foundSoal = group.items.find(soal => soal.id === soalId);
        if (foundSoal) {
          foundSubjectId = foundSoal.subject_id;
          break; 
        }
      }

      const routeName = this.isAdminRoute ? 'AdminSoalList' : 'DosenSoalList';
      
      if (foundSubjectId) {
        // Navigasi ke halaman list DENGAN query parameter subject_id
        this.$router.push({ name: routeName, query: { subject_id: foundSubjectId } });
      } else {
        // Fallback jika soal tidak ditemukan (seharusnya tidak terjadi)
        console.warn(`Soal dengan ID ${soalId} tidak ditemukan.`);
        this.$router.push({ name: routeName });
      }
    },

    async handleDeleteSoal(soalId) {
      if (confirm(`Anda yakin ingin menghapus soal dengan ID: ${soalId}?`)) {
        try {
          await deleteQuestion(soalId);
          alert('Soal berhasil dihapus!');
          this.fetchSoalData(); 
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