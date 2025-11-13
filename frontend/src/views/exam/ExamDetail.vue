<template>
  <div class="bg-white rounded-lg shadow-md p-6">
    <div class="flex justify-between items-center mb-6">
      <h2 class="text-2xl font-bold text-gray-800">Detail Ujian</h2>
      <router-link
        :to="{ name: isAdminRoute ? 'AdminManageExam' : 'DosenManageExam' }"
        class="text-sm bg-gray-200 hover:bg-gray-300 text-gray-800 px-4 py-2 rounded-md"
      >
        ⬅️ Kembali
      </router-link>
    </div>

    <div v-if="loading" class="text-center text-gray-500 py-10">
      Memuat data ujian...
    </div>

    <div v-else-if="error" class="text-center text-red-600 py-10">
      {{ error }}
    </div>

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

      <div class="flex gap-3 mt-8">
        <router-link
          :to="{ name: isAdminRoute ? 'AdminExamEdit' : 'DosenExamEdit', params: { id: exam.id } }"
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

      <div class="mt-10">
        <div class="flex justify-between items-center mb-4">
          <h3 class="text-lg font-semibold text-gray-800">Daftar Soal Ujian</h3>
          <button
            @click="openAddSoalModal"
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
              v-for="q in examQuestions"
              :key="q.id"
              class="border-b hover:bg-gray-50"
            >
              <td class="p-3">{{ q.question.id }}</td>
              <td class="p-3">{{ q.question.question_text.substring(0, 100) }}...</td>
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

  <div v-if="showAddSoalModal" class="fixed inset-0 bg-black bg-opacity-50 flex justify-center items-center p-4 z-50">
    <div class="bg-white rounded-lg shadow-xl w-full max-w-2xl max-h-[90vh] flex flex-col">
      <h3 class="text-lg font-semibold p-4 border-b">Tambah Soal ke Ujian</h3>
      
      <div class="p-4 space-y-4 overflow-y-auto">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1">Pilih Subjek Mata Kuliah</label>
          <select v-model="selectedSubject" class="w-full p-2 border rounded-md bg-white">
            <option :value="null" disabled>-- Pilih Subjek --</option>
            <option v-for="subject in availableSubjects" :key="subject.id" :value="subject.id">
              {{ subject.title }}
            </option>
          </select>
        </div>

        <div class="border rounded-md max-h-[50vh] overflow-y-auto">
          <div v-if="modalLoading" class="text-center p-10 text-gray-500">
            Memuat soal...
          </div>
          <div v-else-if="questionsForSubject.length === 0" class="text-center p-10 text-gray-500">
            {{ selectedSubject ? 'Tidak ada soal tersedia untuk subjek ini (atau sudah ditambahkan).' : 'Silakan pilih subjek.' }}
          </div>
          <table v-else class="w-full text-sm">
            <thead>
              <tr class="bg-gray-50 border-b">
                <th class="p-2 w-10">Pilih</th>
                <th class="p-2 text-left">Soal</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="q in questionsForSubject" :key="q.id" class="border-b last:border-b-0 hover:bg-gray-50">
                <td class="p-2 text-center">
                  <input 
                    type="checkbox" 
                    :value="q.id" 
                    v-model="selectedQuestions" 
                    class="rounded"
                  />
                </td>
                <td class="p-2">
                  {{ q.question_text.substring(0, 120) }}...
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>

      <div class="p-4 border-t flex justify-end gap-3 bg-gray-50 rounded-b-lg">
        <button @click="closeAddSoalModal" class="px-4 py-2 bg-gray-200 text-gray-800 rounded-md hover:bg-gray-300">
          Batal
        </button>
        <button 
          @click="handleAddSoal" 
          :disabled="saveLoading"
          class="px-4 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700 disabled:opacity-50"
        >
          {{ saveLoading ? 'Menyimpan...' : 'Tambahkan Soal' }}
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, computed, watch } from "vue";
import { useRoute, useRouter } from "vue-router";
// Pastikan path provider Anda sudah benar
import { getExamById, deleteExam } from "/src/provider/exam.provider";
import {
  getExamQuestionsByExamId,
  deleteExamQuestion,
  addExamQuestions,
} from "/src/provider/examquestion.provider";
import { getPaginatedSubjects } from "/src/provider/subject.provider";
import { getQuestionsBySubject } from "/src/provider/question.provider";

const route = useRoute();
const router = useRouter();

const isAdminRoute = computed(() => route.path.startsWith('/admin'));

const exam = ref(null);
const examQuestions = ref([]);
const loading = ref(true);
const error = ref("");

// State (Ref) untuk Modal
const showAddSoalModal = ref(false);
const modalLoading = ref(false); 
const saveLoading = ref(false);  
const availableSubjects = ref([]);
const selectedSubject = ref(null);
const questionsForSubject = ref([]);
const selectedQuestions = ref([]); 

// Load detail ujian
const loadExamDetails = async () => {
  try {
    const id = route.params.id;
    const res = await getExamById(id);
    // Provider getExamById mengembalikan data langsung
    exam.value = res; 
  } catch (err) {
    console.error("Gagal memuat detail ujian:", err);
    error.value = "Gagal memuat data ujian.";
  }
};

// Load semua soal dalam ujian
const loadExamQuestions = async () => {
  try {
    const id = route.params.id;
    // Provider getExamQuestionsByExamId mengembalikan array
    examQuestions.value = await getExamQuestionsByExamId(id);
  } catch (err) {
    console.error("Gagal memuat soal ujian:", err);
  }
};

// Hapus ujian
const removeExam = async (id) => {
  if (!confirm("Apakah kamu yakin ingin menghapus ujian ini?")) return;
  try {
    await deleteExam(id);
    alert("✅ Ujian berhasil dihapus!");
    router.push({ name: isAdminRoute.value ? 'AdminManageExam' : 'DosenManageExam' });
  } catch (err) {
    console.error("Gagal menghapus ujian:", err);
    alert("❌ Gagal menghapus ujian!");
  }
};

// Hapus soal dari ujian
const handleDeleteQuestion = async (id) => {
  if (!confirm("Apakah kamu yakin ingin menghapus soal ini dari ujian?"))
    return;
  try {
    // 'id' di sini adalah examQuestionId
    await deleteExamQuestion(id);
    alert("Soal berhasil dihapus dari ujian!");
    loadExamQuestions(); // Muat ulang daftar soal
  } catch (err) {
    console.error("Gagal menghapus soal:", err);
    alert("Gagal menghapus soal.");
  }
};

// --- Logika Modal Tambah Soal ---

// Membuka modal
const openAddSoalModal = () => {
  selectedSubject.value = null;
  questionsForSubject.value = [];
  selectedQuestions.value = [];
  
  fetchAvailableSubjects(); 
  showAddSoalModal.value = true;
};

// Menutup modal
const closeAddSoalModal = () => {
  showAddSoalModal.value = false;
};

// Mengambil daftar subjek dari database
const fetchAvailableSubjects = async () => {
  modalLoading.value = true;
  try {
    const res = await getPaginatedSubjects(100, 0, "");
    availableSubjects.value = res.data || [];
  } catch (err) {
    console.error("Gagal mengambil daftar subjek:", err);
  } finally {
    modalLoading.value = false;
  }
};

// Mengambil soal berdasarkan subjek yang dipilih
const fetchQuestionsForSubject = async (subjectId) => {
  if (!subjectId) {
    questionsForSubject.value = [];
    return;
  }
  modalLoading.value = true;
  try {
    // Ambil semua soal (asumsi < 500)
    const res = await getQuestionsBySubject(subjectId, 500, 0);
    
    // Buat Set (daftar) ID soal yang SUDAH ADA di ujian ini
    const existingQuestionIds = new Set(examQuestions.value.map(q => q.question.id));
    
    // Filter soal: hanya tampilkan soal yang BELUM ADA di ujian
    questionsForSubject.value = (res.data || []).filter(q => !existingQuestionIds.has(q.id));

  } catch (err) {
    console.error("Gagal mengambil daftar soal:", err);
    questionsForSubject.value = [];
  } finally {
    modalLoading.value = false;
  }
};

// Kirim soal yang dipilih ke backend
const handleAddSoal = async () => {
  if (selectedQuestions.value.length === 0) {
    alert("Pilih setidaknya satu soal untuk ditambahkan.");
    return;
  }

  saveLoading.value = true;
  try {
    const examId = exam.value.id;
    await addExamQuestions(examId, selectedQuestions.value);
    
    alert("Soal berhasil ditambahkan!");
    closeAddSoalModal();
    loadExamQuestions(); // Muat ulang daftar soal di halaman detail

  } catch (err) {
    alert("Gagal menambahkan soal. Silakan coba lagi.");
  } finally {
    saveLoading.value = false;
  }
};

// Watcher untuk memantau dropdown subjek
watch(selectedSubject, (newSubjectId) => {
  fetchQuestionsForSubject(newSubjectId);
});

// --- Fungsi Helper (Format Tampilan) ---

const formatDate = (date) => {
  if (!date) return "-";
  return new Date(date).toLocaleString("id-ID", {
    dateStyle: "medium",
    timeStyle: "short",
  });
};
const statusText = (status) => {
  if (status === "not_started") return "Belum Dimulai";
  if (status === "running") return "Sedang Berlangsung";
  if (status === "finished") return "Selesai";
  return "-";
};
const statusClass = (status) => {
  if (status === "not_started")
    return "bg-yellow-100 text-yellow-700 border border-yellow-300";
  if (status === "running")
    return "bg-green-100 text-green-700 border border-green-300";
  if (status === "finished")
    return "bg-gray-100 text-gray-700 border border-gray-300";
  return "bg-gray-100 text-gray-600 border border-gray-300";
};

// onMounted
onMounted(async () => {
  await loadExamDetails();
  await loadExamQuestions();
  loading.value = false;
});
</script>