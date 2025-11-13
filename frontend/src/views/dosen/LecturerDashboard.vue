<template>
  <div>
    <div v-if="loading" class="text-center py-20">
      <p class="text-gray-600">Memuat data dasbor...</p>
    </div>

    <div v-else-if="error" class="text-center py-20 bg-red-50 p-8 rounded-lg">
      <h3 class="text-xl font-semibold text-red-700">Gagal memuat data</h3>
      <p class="text-red-600 mt-2">{{ error }}</p>
    </div>

    <div v-else class="grid grid-cols-1 gap-6">
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        
        <div class="p-6 bg-white rounded-lg shadow-md flex items-center gap-5">
          <div class="bg-indigo-100 p-4 rounded-full">
            <BookOpen class="w-8 h-8 text-indigo-600" />
          </div>
          <div>
            <p class="text-sm text-gray-500">Total Ujian</p>
            <p class="text-3xl font-bold text-dark-text">{{ stats.totalExam }}</p>
          </div>
        </div>
        
        <div class="p-6 bg-white rounded-lg shadow-md flex items-center gap-5">
          <div class="bg-yellow-100 p-4 rounded-full">
            <FileText class="w-8 h-8 text-yellow-600" />
          </div>
          <div>
            <p class="text-sm text-gray-500">Total Soal</p>
            <p class="text-3xl font-bold text-dark-text">{{ stats.totalSoal }}</p>
          </div>
        </div>
      </div>

      </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { BookOpen, FileText } from 'lucide-vue-next';
import { getAllExam } from '../../provider/exam.provider';
import { getmanyQuestions } from '../../provider/question.provider';

const stats = ref({
  totalExam: 0,
  totalSoal: 0,
});
const loading = ref(true);
const error = ref(null);

const fetchDashboardData = async () => {
  try {
    const [examResult, questionResult] = await Promise.allSettled([
      // ## PERUBAHAN 1: Panggil TANPA parameter ##
      // agar mengikuti limit default (misal: 10)
      getAllExam(), 
      getmanyQuestions(1, 0) // Biarkan ini, karena total soal sudah benar
    ]);

    // Cek hasil Panggilan Ujian
    if (examResult.status === 'fulfilled') {
      // ## PERUBAHAN 2: LOGIKA IDENTIK DENGAN ADMIN DASHBOARD ##
      // 'examResult.value' adalah { data: [...], total: 48 }
      // Kita ambil array 'data' di dalamnya
      const examList = examResult.value.data || [];
      
      // Kita hitung panjang array (length), BUKAN total
      stats.value.totalExam = examList.length;
      
    } else {
      console.error("Gagal mengambil data ujian:", examResult.reason);
      if (!error.value) error.value = "Gagal memuat data ujian.";
      stats.value.totalExam = 'N/A';
    }

    // Cek hasil Panggilan Soal (Ini tetap menggunakan logika .total)
    if (questionResult.status === 'fulfilled') {
      stats.value.totalSoal = questionResult.value.total || 0;
    } else {
      console.error("Gagal mengambil data soal:", questionResult.reason);
      if (!error.value) error.value = "Gagal memuat data soal.";
      stats.value.totalSoal = 'N/A';
    }

  } catch (err) {
    console.error("Gagal memuat data dasbor:", err);
    error.value = "Terjadi kesalahan saat mengambil data.";
  } finally {
    loading.value = false;
  }
};

onMounted(() => {
  fetchDashboardData();
});
</script>