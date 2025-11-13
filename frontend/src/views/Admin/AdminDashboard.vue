<template>
  <div>
    <div v-if="loading" class="text-center py-20">
      <p class="text-gray-600 text-lg">Memuat data dasbor admin...</p>
    </div>

    <div v-else-if="error" class="text-center py-20 bg-red-50 p-8 rounded-lg">
      <h3 class="text-xl font-semibold text-red-700">Gagal memuat data</h3>
      <p class="text-red-600 mt-2">{{ error }}</p>
    </div>

    <div v-else class="grid grid-cols-1 gap-6">
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
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
          <div class="bg-green-100 p-4 rounded-full">
            <UserCheck class="w-8 h-8 text-green-600" />
          </div>
          <div>
            <p class="text-sm text-gray-500">Total Dosen</p>
            <p class="text-3xl font-bold text-dark-text">{{ stats.totalLecturers }}</p>
          </div>
        </div>
        
        <div class="p-6 bg-white rounded-lg shadow-md flex items-center gap-5">
          <div class="bg-indigo-100 p-4 rounded-full">
            <BookOpen class="w-8 h-8 text-indigo-600" />
          </div>
          <div>
            <p class="text-sm text-gray-500">Total Ujian</p>
            <p class="text-3xl font-bold text-dark-text">{{ stats.totalExams }}</p>
          </div>
        </div>

        <div class="p-6 bg-white rounded-lg shadow-md flex items-center gap-5">
          <div class="bg-yellow-100 p-4 rounded-full">
            <FileText class="w-8 h-8 text-yellow-600" />
          </div>
          <div>
            <p class="text-sm text-gray-500">Total Soal</p>
            <p class="text-3xl font-bold text-dark-text">{{ stats.totalQuestions }}</p>
          </div>
        </div>
      </div>

      <div class="p-6 bg-white rounded-lg shadow-md">
        <h3 class="mb-4 text-lg font-semibold text-dark-text">Pengguna Baru Terdaftar</h3>
        <div v-if="recentUsers.length > 0" class="space-y-3">
          <div v-for="user in recentUsers" :key="user.id" class="flex justify-between items-center p-3 bg-gray-50 rounded-md border">
            <div>
              <p class="text-sm font-semibold text-gray-800">{{ user.name }}</p>
              <p class="text-xs text-gray-500">{{ user.email }}</p>
            </div>
            <span :class="roleClass(user.role)" class="text-xs font-semibold px-2 py-1 rounded-full capitalize">
              {{ user.role === 'user' ? 'Mahasiswa' : user.role }}
            </span>
          </div>
        </div>
        <div v-else class="flex items-center justify-center h-40 text-gray-500">
          Belum ada pengguna baru yang terdaftar.
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
// ## 3. Impor 'BookOpen' dikembalikan ##
import { Users, UserCheck, BookOpen, FileText } from 'lucide-vue-next';
import { getUsers } from '../../provider/user.provider';
// ## 4. Impor 'getAllExam' dikembalikan ##
import { getAllExam } from '../../provider/exam.provider';
import { getmanyQuestions } from '../../provider/question.provider';

const loading = ref(true);
const error = ref(null);
const stats = ref({
  totalUsers: 0,
  totalLecturers: 0,
  totalExams: 0, // <-- Dikembalikan
  totalQuestions: 0,
});
const recentUsers = ref([]);

const fetchDashboardData = async () => {
  try {
    // ## 5. Panggilan API 'getAllExam' dikembalikan ##
    const [userResponse, examResponse, questionResponse] = await Promise.all([
      getUsers(),
      getAllExam(),
      getmanyQuestions(1, 0)
    ]);

    const userList = userResponse.data || [];
    const examList = examResponse.data || []; // <-- Dikembalikan
    const totalQuestions = questionResponse.total || 0;

    // Menghitung statistik
    stats.value.totalUsers = userList.length;
    stats.value.totalLecturers = userList.filter(u => u.role === 'lecturer').length;
    stats.value.totalExams = examList.length; // <-- Dikembalikan
    stats.value.totalQuestions = totalQuestions;

    recentUsers.value = userList.slice(-5).reverse();

  } catch (err) {
    console.error("Gagal memuat data dasbor admin:", err);
    error.value = "Terjadi kesalahan saat mengambil data dari server.";
  } finally {
    loading.value = false;
  }
};

const roleClass = (role) => {
  if (role === 'lecturer') return 'bg-green-100 text-green-800';
  if (role === 'user') return 'bg-blue-100 text-blue-800';
  if (role === 'admin') return 'bg-red-100 text-red-800';
  return 'bg-gray-100 text-gray-800';
};

onMounted(() => {
  fetchDashboardData();
});
</script>