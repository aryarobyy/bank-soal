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
            <p class="text-3xl font-bold text-dark-text">
              {{ stats.totalExam }}
            </p>
          </div>
        </div>

        <div class="p-6 bg-white rounded-lg shadow-md flex items-center gap-5">
          <div class="bg-yellow-100 p-4 rounded-full">
            <FileText class="w-8 h-8 text-yellow-600" />
          </div>
          <div>
            <p class="text-sm text-gray-500">Total Soal</p>
            <p class="text-3xl font-bold text-dark-text">
              {{ stats.totalSoal }}
            </p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, watch } from "vue";
import { BookOpen, FileText } from "lucide-vue-next";
import { getAllExam } from "../../provider/exam.provider";
import { getmanyQuestions } from "../../provider/question.provider";
import { useGetCurrentUser } from "../../hooks/useGetCurrentUser";

const { user } = useGetCurrentUser();

const stats = ref({
  totalExam: 0,
  totalSoal: 0,
});
const loading = ref(true);
const error = ref(null);

const fetchDashboardData = async () => {
  // Tunggu sampai user.value ada
  if (!user.value) {
    return;
  }

  try {
    const [examResult, questionResult] = await Promise.allSettled([
      // ## PERBAIKAN 1: Gunakan limit 1 saja karena kita cuma butuh field 'total' ##
      // Parameter ke-3 adalah creator_id (user.value.id) untuk filter dosen ybs
      getAllExam(1, 0, user.value.id),

      getmanyQuestions(1, 0),
    ]);

    // ## PERBAIKAN 2: Ambil field .total dari response provider ##
    if (examResult.status === "fulfilled") {
      // Karena provider me-return { data: [...], total: ... }
      // Kita ambil .total-nya langsung.
      stats.value.totalExam = examResult.value?.total || 0;
    } else {
      console.error("Gagal mengambil data ujian:", examResult.reason);
      if (!error.value) error.value = "Gagal memuat data ujian.";
      stats.value.totalExam = 0;
    }

    // Cek hasil Panggilan Soal
    if (questionResult.status === "fulfilled") {
      stats.value.totalSoal = questionResult.value.total || 0;
    } else {
      console.error("Gagal mengambil data soal:", questionResult.reason);
      if (!error.value) error.value = "Gagal memuat data soal.";
      stats.value.totalSoal = 0;
    }
  } catch (err) {
    console.error("Gagal memuat data dasbor:", err);
    error.value = "Terjadi kesalahan saat mengambil data.";
  } finally {
    loading.value = false;
  }
};

// Watcher untuk memuat data saat user sudah login
watch(
  user,
  (currentUser) => {
    if (currentUser) {
      fetchDashboardData();
    }
  },
  { immediate: true }
);
</script>
