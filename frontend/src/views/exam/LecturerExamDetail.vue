<template>
  <div class="bg-white rounded-lg shadow-md p-6">
    <!-- Header -->
    <div class="flex justify-between items-center mb-6">
      <h2 class="text-2xl font-bold text-gray-800">Detail Ujian</h2>
      <router-link
        to="/dosen/exam"
        class="text-sm bg-gray-200 hover:bg-gray-300 text-gray-800 px-4 py-2 rounded-md"
      >
        ⬅️ Kembali
      </router-link>
    </div>

    <!-- Loading -->
    <div v-if="loading" class="text-center text-gray-500 py-10">
      Memuat data ujian...
    </div>

    <!-- Error -->
    <div v-else-if="error" class="text-center text-red-600 py-10">
      {{ error }}
    </div>

    <!-- Data Ujian -->
    <div v-else-if="exam" class="space-y-4">
      <div>
        <h3 class="text-lg font-semibold text-gray-700">Nama Ujian</h3>
        <p class="text-gray-600">{{ exam.title }}</p>
      </div>

      <div>
        <h3 class="text-lg font-semibold text-gray-700">Deskripsi</h3>
        <p class="text-gray-600">
          {{ exam.description || "Tidak ada deskripsi" }}
        </p>
      </div>

      <div class="grid grid-cols-2 gap-4">
        <div>
          <h3 class="text-lg font-semibold text-gray-700">Kesulitan</h3>
          <p class="capitalize text-gray-600">{{ exam.difficulty }}</p>
        </div>
        <div>
          <h3 class="text-lg font-semibold text-gray-700">Durasi</h3>
          <p class="text-gray-600">{{ exam.long_time }} menit</p>
        </div>
      </div>

      <div class="grid grid-cols-2 gap-4">
        <div>
          <h3 class="text-lg font-semibold text-gray-700">Waktu Mulai</h3>
          <p class="text-gray-600">{{ formatDate(exam.started_at) }}</p>
        </div>
        <div>
          <h3 class="text-lg font-semibold text-gray-700">Waktu Selesai</h3>
          <p class="text-gray-600">{{ formatDate(exam.finished_at) }}</p>
        </div>
      </div>

      <div>
        <h3 class="text-lg font-semibold text-gray-700">Status</h3>
        <span
          :class="statusClass(exam.status)"
          class="inline-block px-3 py-1 rounded-full text-sm font-medium"
        >
          {{ statusText(exam.status) }}
        </span>
      </div>

      <!-- Tombol Aksi -->
      <div class="flex gap-3 mt-8">
        <router-link
          :to="`/dosen/exam/edit/${exam.id}`"
          class="bg-blue-600 hover:bg-blue-700 text-white px-4 py-2 rounded-md"
        >
          ✏️ Edit Ujian
        </router-link>

        <button
          @click="removeExam(exam.id)"
          class="bg-red-500 hover:bg-red-600 text-white px-4 py-2 rounded-md"
        >
          🗑️ Hapus Ujian
        </button>
      </div>

      <!-- Soal -->
      <div class="mt-10">
        <div class="flex justify-between items-center mb-4">
          <h3 class="text-lg font-semibold text-gray-800">Daftar Soal Ujian</h3>
          <button
            class="bg-green-500 hover:bg-green-600 text-white px-4 py-2 rounded-md"
          >
            ➕ Tambah Soal
          </button>
        </div>

        <table class="w-full border border-gray-200 rounded-lg text-left">
          <thead>
            <tr class="bg-gray-100 border-b text-gray-700">
              <th class="p-3">ID Soal</th>
              <th class="p-3">Judul Soal</th>
              <th class="p-3 text-center">Aksi</th>
            </tr>
          </thead>
          <tbody>
            <tr
              v-for="(q, index) in examQuestions"
              :key="index"
              class="border-b hover:bg-gray-50"
            >
              <td class="p-3">{{ q.id }}</td>
              <td class="p-3">{{ q.title || "Untitled Question" }}</td>
              <td class="p-3 text-center">
                <button
                  @click="handleDeleteQuestion(q.id)"
                  class="text-red-600 hover:text-red-800"
                >
                  🗑️ Hapus
                </button>
              </td>
            </tr>
            <tr v-if="examQuestions.length === 0">
              <td colspan="3" class="p-4 text-center text-gray-500">
                Belum ada soal ditambahkan.
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from "vue";
import { useRoute, useRouter } from "vue-router";
import { getExamById, deleteExam } from "@/provider/exam.provider";
import {
  getExamQuestionsByExamId,
  deleteExamQuestion,
} from "@/provider/examquestion.provider";

const route = useRoute();
const router = useRouter();

const exam = ref(null);
const examQuestions = ref([]);
const loading = ref(true);
const error = ref("");

// ✅ Load detail ujian
const loadExamDetails = async () => {
  try {
    const id = route.params.id;
    const res = await getExamById(id);
    exam.value = res?.data || res;
  } catch (err) {
    console.error("Gagal memuat detail ujian:", err);
    error.value = "Gagal memuat data ujian.";
  }
};

// ✅ Load semua soal dalam ujian
const loadExamQuestions = async () => {
  try {
    const id = route.params.id;
    const res = await getExamQuestionsByExamId(id);
    examQuestions.value = Array.isArray(res) ? res : res?.data || [];
  } catch (err) {
    console.error("Gagal memuat soal ujian:", err);
  }
};

// ✅ Hapus ujian
const removeExam = async (id) => {
  if (!confirm("Apakah kamu yakin ingin menghapus ujian ini?")) return;
  try {
    await deleteExam(id);
    alert("✅ Ujian berhasil dihapus!");
    router.push("/dosen/exam");
  } catch (err) {
    console.error("Gagal menghapus ujian:", err);
    alert("❌ Gagal menghapus ujian!");
  }
};

// ✅ Hapus soal dari ujian
const handleDeleteQuestion = async (id) => {
  if (!confirm("Apakah kamu yakin ingin menghapus soal ini dari ujian?"))
    return;
  try {
    await deleteExamQuestion(id);
    alert("Soal berhasil dihapus dari ujian!");
    loadExamQuestions();
  } catch (err) {
    console.error("Gagal menghapus soal:", err);
    alert("Gagal menghapus soal.");
  }
};

// ✅ Format tanggal
const formatDate = (date) => {
  if (!date) return "-";
  return new Date(date).toLocaleString("id-ID", {
    dateStyle: "medium",
    timeStyle: "short",
  });
};

// ✅ Status text
const statusText = (status) => {
  if (status === "not_started") return "Belum Dimulai";
  if (status === "running") return "Sedang Berlangsung";
  if (status === "finished") return "Selesai";
  return "-";
};

// ✅ Warna label status
const statusClass = (status) => {
  if (status === "not_started")
    return "bg-yellow-100 text-yellow-700 border border-yellow-300";
  if (status === "running")
    return "bg-green-100 text-green-700 border border-green-300";
  if (status === "finished")
    return "bg-gray-100 text-gray-700 border border-gray-300";
  return "bg-gray-100 text-gray-600 border border-gray-300";
};

onMounted(async () => {
  await loadExamDetails();
  await loadExamQuestions();
  loading.value = false;
});
</script>
