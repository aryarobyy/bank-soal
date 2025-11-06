<template>
  <div>
    <div class="mb-8">
      <h2 class="text-3xl font-bold text-[#24305E]">Statistik Sistem</h2>
      <p class="text-sm text-gray-500 mt-1">
        Ringkasan data dari seluruh sistem Latih.in
      </p>
    </div>

    <div v-if="loadingStats" class="text-center py-10">
      <p class="text-gray-500">Memuat data statistik...</p>
    </div>
    <div v-else-if="errorStats" class="text-center py-10 bg-red-50 p-4 rounded-lg">
      <p class="text-red-600">{{ errorStats }}</p>
    </div>
    <div v-else class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6 mb-8">
      
      <div class="p-6 bg-white rounded-lg shadow-md flex items-center gap-5">
        <div class="bg-blue-100 p-4 rounded-full">
          <Users class="w-8 h-8 text-blue-600" />
        </div>
        <div>
          <p class="text-sm text-gray-500">Total Pengguna</p>
          <p class="text-3xl font-bold text-dark-text">{{ stats.totalUsers }}</p>
        </div>
      </div>
      
      <div class="p-6 bg-white rounded-lg shadow-md flex items-center gap-5">
        <div class="bg-red-100 p-4 rounded-full">
          <Shield class="w-8 h-8 text-red-600" />
        </div>
        <div>
          <p class="text-sm text-gray-500">Total Admin</p>
          <p class="text-3xl font-bold text-dark-text">{{ stats.totalAdmins }}</p>
        </div>
      </div>

      <div class="p-6 bg-white rounded-lg shadow-md flex items-center gap-5">
        <div class="bg-green-100 p-4 rounded-full">
          <UserCheck class="w-8 h-8 text-green-600" />
        </div>
        <div>
          <p class="text-sm text-gray-500">Total Dosen</p>
          <p class="text-3xl font-bold text-dark-text">{{ stats.totalDosen }}</p>
        </div>
      </div>
      
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
</template>

<script setup>
import { ref, onMounted } from "vue";
// ## 3. Impor 'BookOpen' dikembalikan ##
import { Users, Shield, UserCheck, FileText, BookOpen } from "lucide-vue-next";
import { getUsers } from "../../provider/user.provider";
// ## 4. Impor 'getAllExam' dikembalikan ##
import { getAllExam } from "../../provider/exam.provider";
import { getmanyQuestions } from "../../provider/question.provider";

// State untuk data statistik
const loadingStats = ref(true);
const errorStats = ref(null);
const stats = ref({
  totalUsers: 0,
  totalAdmins: 0,
  totalDosen: 0,
  totalSoal: 0,
  totalExam: 0, // <-- Dikembalikan
});

// ## 5. Fungsi fetch data dikembalikan seperti semula ##
const fetchDashboardData = async () => {
  try {
    const [userResult, examResult, questionResult] = await Promise.allSettled([
      getUsers(),
      getAllExam(), // <-- Panggilan API dikembalikan
      getmanyQuestions(1, 0),
    ]);

    // Cek hasil Panggilan User
    if (userResult.status === 'fulfilled') {
      const userList = userResult.value.data || [];
      stats.value.totalUsers = userList.length;
      stats.value.totalAdmins = userList.filter((u) => u.role === "admin").length;
      stats.value.totalDosen = userList.filter((u) => u.role === "lecturer").length;
    } else {
      console.error("Gagal mengambil data user:", userResult.reason);
      errorStats.value = "Gagal memuat data pengguna.";
    }

    // Cek hasil Panggilan Ujian
    if (examResult.status === 'fulfilled') {
      stats.value.totalExam = (examResult.value.data || []).length;
    } else {
       // Jika error 403 terjadi lagi, ini akan menampilkannya di konsol
       console.error("Gagal mengambil data ujian:", examResult.reason);
       stats.value.totalExam = "N/A";
       if (!errorStats.value) errorStats.value = "Gagal memuat data ujian.";
    }

    // Cek hasil Panggilan Soal
    if (questionResult.status === 'fulfilled') {
      stats.value.totalSoal = questionResult.value.data.total || 0;
    } else {
       console.error("Gagal mengambil data soal:", questionResult.reason);
       if (!errorStats.value) errorStats.value = "Gagal memuat data soal.";
    }

  } catch (err) {
    console.error("Gagal memuat data dasbor:", err);
    errorStats.value = "Terjadi kesalahan saat mengambil data statistik.";
  } finally {
    loadingStats.value = false;
  }
};

onMounted(() => {
  fetchDashboardData();
});
</script>

<style scoped>
/* Style tidak diperlukan lagi karena tabel sudah dihapus */
</style>