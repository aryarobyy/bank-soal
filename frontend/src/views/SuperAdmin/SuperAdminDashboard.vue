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
import { Users, Shield, UserCheck, FileText, BookOpen } from "lucide-vue-next";
// ## PERBAIKAN 1: Tambahkan 'getUsersByRole' ##
import { getUsers, getUsersByRole } from "../../provider/user.provider";
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
  totalExam: 0,
});

const fetchDashboardData = async () => {
  try {
    // ## PERBAIKAN 2: Gunakan limit 1 dan panggil API spesifik ##
    const [userTotalRes, adminRes, dosenRes, examRes, soalRes] = await Promise.all([
      getUsers(1, 0),                    // Ambil Total User
      getUsersByRole('admin', 1, 0),     // Ambil Total Admin
      getUsersByRole('lecturer', 1, 0),  // Ambil Total Dosen
      getAllExam(1, 0),                  // Ambil Total Ujian (sudah diperbaiki providernya)
      getmanyQuestions(1, 0),            // Ambil Total Soal
    ]);

    // ## PERBAIKAN 3: Ambil nilai .total langsung dari response ##
    
    // Total Users
    stats.value.totalUsers = userTotalRes?.total || 0;
    
    // Total Admins
    stats.value.totalAdmins = adminRes?.total || 0;
    
    // Total Dosen
    stats.value.totalDosen = dosenRes?.total || 0;

    // Total Ujian (Menggunakan .total dari object {data, total})
    stats.value.totalExam = examRes?.total || 0;

    // Total Soal
    stats.value.totalSoal = soalRes?.total || 0;

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