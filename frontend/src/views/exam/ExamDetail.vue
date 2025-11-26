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
              <td class="p-3">{{ q.id }}</td>
              <td class="p-3">
                {{ q.question_text?.substring(0, 100) || "[Soal tidak valid]" }}...
              </td>
              <td class="p-3 text-center">
                <button
                  @click="handleDeleteQuestion(q)"
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
      
      <div class="p-4 space-y-4 overflow-y-auto flex-1">
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
            {{ selectedSubject ? 'Tidak ada soal tersedia di halaman ini.' : 'Silakan pilih subjek.' }}
          </div>
          
          <div v-else>
            <table class="w-full text-sm">
              <thead>
                <tr class="bg-gray-50 border-b sticky top-0">
                  <th class="p-2 w-10 text-center">
                    <i class="fas fa-check-square text-gray-400"></i>
                  </th>
                  <th class="p-2 text-left">Pertanyaan</th>
                </tr>
              </thead>
              <tbody>
                <tr v-for="q in questionsForSubject" :key="q.id" class="border-b last:border-b-0 hover:bg-gray-50">
                  <td class="p-2 text-center align-top pt-3">
                    <input 
                      type="checkbox" 
                      :value="q.id" 
                      v-model="selectedQuestions" 
                      class="rounded cursor-pointer w-4 h-4"
                      :disabled="isQuestionAlreadyAdded(q.id)"
                    />
                  </td>
                  <td class="p-2">
                    <div class="font-medium text-gray-800 mb-1">
                      {{ q.question_text?.substring(0, 150) }}...
                    </div>
                    <div v-if="isQuestionAlreadyAdded(q.id)" class="text-xs text-green-600 font-bold">
                      ✓ Sudah ditambahkan
                    </div>
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>

        <div v-if="selectedSubject && modalTotalPages > 1" class="flex justify-between items-center text-sm text-gray-600 pt-2">
          <span>Halaman {{ modalPage }} dari {{ modalTotalPages }}</span>
          <div class="flex gap-2">
            <button 
              @click="prevModalPage" 
              :disabled="modalPage === 1"
              class="px-3 py-1 bg-gray-200 rounded hover:bg-gray-300 disabled:opacity-50"
            >
              Prev
            </button>
            <button 
              @click="nextModalPage" 
              :disabled="modalPage === modalTotalPages"
              class="px-3 py-1 bg-gray-200 rounded hover:bg-gray-300 disabled:opacity-50"
            >
              Next
            </button>
          </div>
        </div>

      </div>

      <div class="p-4 border-t flex justify-between items-center bg-gray-50 rounded-b-lg">
        <span class="text-sm text-gray-600">
          {{ selectedQuestions.length }} soal dipilih
        </span>
        <div class="flex gap-3">
          <button @click="closeAddSoalModal" class="px-4 py-2 bg-gray-200 text-gray-800 rounded-md hover:bg-gray-300">
            Batal
          </button>
          <button 
            @click="handleAddSoal" 
            :disabled="saveLoading || selectedQuestions.length === 0"
            class="px-4 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700 disabled:opacity-50 disabled:cursor-not-allowed"
          >
            {{ saveLoading ? 'Menyimpan...' : 'Tambahkan Soal' }}
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, computed, watch } from "vue";
import { useRoute, useRouter } from "vue-router";
import { getExamById, deleteExam } from "/src/provider/exam.provider";
import {
  addExamQuestions,
  deleteExamQuestions,
} from "/src/provider/examquestion.provider";
import { getPaginatedSubjects } from "/src/provider/subject.provider";
import { getQuestionsBySubject, getQuestionsByExam } from "/src/provider/question.provider";

const route = useRoute();
const router = useRouter();
const isAdminRoute = computed(() => route.path.startsWith('/admin'));

const exam = ref(null);
const examQuestions = ref([]);
const loading = ref(true);
const error = ref("");

// State Modal & Pagination
const showAddSoalModal = ref(false);
const modalLoading = ref(false);
const saveLoading = ref(false);
const availableSubjects = ref([]);
const selectedSubject = ref(null);
const questionsForSubject = ref([]);
const selectedQuestions = ref([]);

// Pagination State Khusus Modal
const modalPage = ref(1);
const modalLimit = 10; // Ambil 10 soal per halaman agar tidak error backend
const modalTotalItems = ref(0);

// Computed Total Pages Modal
const modalTotalPages = computed(() => {
  if (modalTotalItems.value === 0) return 1;
  return Math.ceil(modalTotalItems.value / modalLimit);
});

// Load Detail Ujian
const loadExamDetails = async () => {
  try {
    const id = route.params.id;
    const res = await getExamById(id);
    exam.value = res;
  } catch (err) {
    console.error("Gagal memuat detail ujian:", err);
    error.value = "Gagal memuat data ujian.";
  }
};

// Load Daftar Soal yang sudah ada di Ujian
const loadExamQuestions = async () => {
  try {
    const id = route.params.id;
    const result = await getQuestionsByExam(id);
    examQuestions.value = result.data || [];
  } catch (err) {
    console.error("Gagal memuat soal ujian:", err);
    examQuestions.value = [];
  }
};

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

const handleDeleteQuestion = async (question) => {
  if (!confirm(`Yakin ingin menghapus soal (ID: ${question.id}) dari ujian ini?`)) return;
  try {
    await deleteExamQuestions(exam.value.id, [question.id]);
    alert("Soal berhasil dihapus dari ujian!");
    loadExamQuestions();
  } catch (err) {
    console.error("Gagal menghapus soal:", err);
    alert("Gagal menghapus soal.");
  }
};

// Open Modal
const openAddSoalModal = () => {
  selectedSubject.value = null;
  questionsForSubject.value = [];
  selectedQuestions.value = [];
  modalPage.value = 1; // Reset page
  modalTotalItems.value = 0;
  
  fetchAvailableSubjects();
  showAddSoalModal.value = true;
};

const closeAddSoalModal = () => {
  showAddSoalModal.value = false;
};

// Fetch Subjects
const fetchAvailableSubjects = async () => {
  modalLoading.value = true;
  try {
    const res = await getPaginatedSubjects(100, 0, "");
    availableSubjects.value = res.data.data || [];
  } catch (err) {
    console.error("Gagal mengambil daftar subjek:", err);
  } finally {
    modalLoading.value = false;
  }
};

// Fetch Soal per Halaman (PAGINATION FIX)
const fetchQuestionsForSubject = async (subjectId) => {
  if (!subjectId) {
    questionsForSubject.value = [];
    return;
  }
  
  modalLoading.value = true;
  try {
    // Hitung offset berdasarkan halaman saat ini
    const offset = (modalPage.value - 1) * modalLimit;
    
    // Panggil API dengan limit kecil (10)
    const result = await getQuestionsBySubject(subjectId, modalLimit, offset);
    
    // Backend Anda harusnya return { data: [], total: ... }
    // Jika return array, handle fallback
    if (result.data && Array.isArray(result.data)) {
       questionsForSubject.value = result.data;
       modalTotalItems.value = result.total || result.data.length; 
    } else if (Array.isArray(result)) {
       // Jika backend lama return array langsung (tanpa total)
       questionsForSubject.value = result;
       modalTotalItems.value = result.length; // Pagination mungkin kurang akurat
    } else {
       questionsForSubject.value = [];
       modalTotalItems.value = 0;
    }

  } catch (err) {
    console.error("Gagal mengambil daftar soal:", err);
    questionsForSubject.value = [];
  } finally {
    modalLoading.value = false;
  }
};

// Navigasi Pagination Modal
const nextModalPage = () => {
  if (modalPage.value < modalTotalPages.value) {
    modalPage.value++;
    fetchQuestionsForSubject(selectedSubject.value);
  }
};

const prevModalPage = () => {
  if (modalPage.value > 1) {
    modalPage.value--;
    fetchQuestionsForSubject(selectedSubject.value);
  }
};

const handleAddSoal = async () => {
  if (selectedQuestions.value.length === 0) return;

  saveLoading.value = true;
  try {
    await addExamQuestions(exam.value.id, selectedQuestions.value);
    alert("Soal berhasil ditambahkan!");
    closeAddSoalModal();
    loadExamQuestions();
  } catch (err) {
    console.error(err);
    alert("Gagal menambahkan soal. Silakan coba lagi.");
  } finally {
    saveLoading.value = false;
  }
};

// Helper: Cek apakah soal sudah ada di ujian (untuk disable checkbox)
const isQuestionAlreadyAdded = (qId) => {
  return examQuestions.value.some(eq => eq.id === qId);
};

// Watch jika Subject berubah -> Reset page ke 1 & fetch
watch(selectedSubject, (newSubjectId) => {
  if (newSubjectId) {
    modalPage.value = 1;
    fetchQuestionsForSubject(newSubjectId);
  } else {
    questionsForSubject.value = [];
  }
});

// Formatters
const formatDate = (date) => {
  if (!date) return "-";
  return new Date(date).toLocaleString("id-ID", {
    dateStyle: "medium",
    timeStyle: "short",
  });
};
const statusText = (status) => status === "not_started" ? "Belum Dimulai" : status === "running" ? "Sedang Berlangsung" : "Selesai";
const statusClass = (status) => status === "running" ? "bg-green-100 text-green-700" : "bg-gray-100 text-gray-600";

onMounted(async () => {
  loading.value = true;
  await loadExamDetails();
  await loadExamQuestions();
  loading.value = false;
});
</script>