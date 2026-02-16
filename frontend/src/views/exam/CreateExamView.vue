<template>
  <div class="flex justify-center items-start min-h-screen bg-gray-50 py-10 px-4">
    <div class="bg-white rounded-xl shadow-lg w-full max-w-5xl overflow-hidden">
      <!-- Header Steps -->
      <div class="bg-gray-50 border-b border-gray-200 px-8 py-4 flex items-center justify-between">
        <h2 class="text-2xl font-bold text-gray-800">Buat Ujian & Soal</h2>
        <div class="flex items-center space-x-2 text-sm font-medium">
          <span :class="step === 1 ? 'text-blue-600' : 'text-gray-400'">1. Detail Ujian</span>
          <span class="text-gray-300">/</span>
          <span :class="step === 2 ? 'text-blue-600' : 'text-gray-400'">2. Tambah Soal</span>
        </div>
      </div>

      <div class="p-8">
        <!-- STEP 1: DETAIL UJIAN -->
        <div v-show="step === 1">
          <form @submit.prevent="nextStep" class="space-y-6 max-w-2xl mx-auto">
            <div>
              <label class="block text-sm font-semibold text-gray-700 mb-1">Nama Ujian</label>
              <input
                v-model="examForm.title"
                type="text"
                class="w-full border border-gray-300 rounded-lg p-2.5 focus:ring-2 focus:ring-blue-500 focus:border-blue-500 transition-all"
                placeholder="Contoh: Ujian Tengah Semester Genap"
                required
              />
            </div>

            <div>
              <label class="block text-sm font-semibold text-gray-700 mb-1">Deskripsi</label>
              <textarea
                v-model="examForm.description"
                rows="3"
                class="w-full border border-gray-300 rounded-lg p-2.5 focus:ring-2 focus:ring-blue-500 focus:border-blue-500 transition-all"
                placeholder="Deskripsi singkat mengenai ujian ini..."
              ></textarea>
            </div>

            <div>
              <label class="block text-sm font-semibold text-gray-700 mb-1">Tingkat Kesulitan</label>
              <select
                v-model="examForm.difficulty"
                class="w-full border border-gray-300 rounded-lg p-2.5 bg-white focus:ring-2 focus:ring-blue-500 focus:border-blue-500 transition-all"
                required
              >
                <option disabled value="">Pilih kesulitan</option>
                <option value="easy">Mudah</option>
                <option value="medium">Sedang</option>
                <option value="hard">Sulit</option>
              </select>
            </div>

            <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
              <div>
                <label class="block text-sm font-semibold text-gray-700 mb-1">Tanggal Mulai</label>
                <input
                  v-model="examForm.started_at"
                  type="datetime-local"
                  class="w-full border border-gray-300 rounded-lg p-2.5 focus:ring-2 focus:ring-blue-500 focus:border-blue-500 transition-all"
                  required
                />
              </div>

              <div>
                <label class="block text-sm font-semibold text-gray-700 mb-1">Tanggal Berakhir</label>
                <input
                  v-model="examForm.finished_at"
                  type="datetime-local"
                  class="w-full border border-gray-300 rounded-lg p-2.5 focus:ring-2 focus:ring-blue-500 focus:border-blue-500 transition-all"
                  required
                />
              </div>
            </div>

            <div>
              <label class="block text-sm font-semibold text-gray-700 mb-1">Durasi (menit)</label>
              <input
                v-model.number="examForm.long_time"
                type="number"
                min="1"
                class="w-full border border-gray-300 rounded-lg p-2.5 focus:ring-2 focus:ring-blue-500 focus:border-blue-500 transition-all"
                required
              />
            </div>

            <!-- Mata Kuliah -->
            <div>
              <label class="block text-sm font-semibold text-gray-700 mb-1">
                <i class="fas fa-book-open mr-1 text-indigo-500"></i>Mata Kuliah
              </label>
              <select
                v-model="examForm.subject_id"
                class="w-full border border-gray-300 rounded-lg p-2.5 bg-white focus:ring-2 focus:ring-blue-500 focus:border-blue-500 transition-all"
              >
                <option :value="null">-- Pilih Mata Kuliah (Opsional) --</option>
                <option v-for="subject in subjects" :key="subject.id" :value="subject.id">
                  {{ subject.title }} ({{ subject.code }})
                </option>
              </select>
              <p class="text-xs text-gray-400 mt-1">Kategorikan ujian berdasarkan mata kuliah</p>
            </div>

            <!-- Tombol Tambah Soal -->
            <div class="pt-4">
              <button
                type="submit"
                class="w-full px-6 py-3 bg-gradient-to-r from-blue-600 to-indigo-600 hover:from-blue-700 hover:to-indigo-700 text-white font-bold rounded-lg transition-all shadow-md hover:shadow-lg transform hover:-translate-y-0.5 flex items-center justify-center gap-3"
              >
                <i class="fas fa-plus-circle text-lg"></i>
                <span>Tambah Soal</span>
                <i class="fas fa-arrow-right ml-1"></i>
              </button>
            </div>
          </form>
        </div>

        <!-- STEP 2: TAMBAH SOAL -->
        <div v-show="step === 2">

          <!-- Import JSON Section -->
          <div class="mb-6">
            <div 
              @dragover.prevent="isDraggingJson = true" 
              @dragleave.prevent="isDraggingJson = false" 
              @drop.prevent="handleDropJson"
              :class="[
                'border-2 border-dashed rounded-xl p-6 text-center transition-all duration-300 cursor-pointer',
                isDraggingJson 
                  ? 'border-indigo-500 bg-indigo-50 scale-[1.01] shadow-lg' 
                  : 'border-gray-300 bg-gradient-to-br from-gray-50 to-white hover:border-indigo-300 hover:bg-indigo-50/30'
              ]"
              @click="triggerJsonInput"
            >
              <div v-if="!jsonPreview" class="space-y-3">
                <div class="flex justify-center">
                  <div class="w-14 h-14 rounded-full bg-indigo-100 flex items-center justify-center">
                    <i class="fas fa-file-import text-2xl text-indigo-500"></i>
                  </div>
                </div>
                <div>
                  <p class="text-sm font-semibold text-gray-700">Import Soal dari File JSON</p>
                  <p class="text-xs text-gray-400 mt-1">Seret file .json ke sini atau <span class="text-indigo-500 font-medium underline">klik untuk memilih file</span></p>
                </div>
                <div class="flex items-center justify-center gap-2 mt-2">
                  <span class="inline-flex items-center gap-1 text-[10px] font-mono bg-gray-100 text-gray-500 px-2 py-1 rounded">
                    <i class="fas fa-info-circle"></i> Format: .json
                  </span>
                </div>
              </div>

              <!-- JSON Preview -->
              <div v-else class="space-y-3" @click.stop>
                <div class="flex items-center justify-center gap-2">
                  <div class="w-10 h-10 rounded-full bg-green-100 flex items-center justify-center">
                    <i class="fas fa-check-circle text-xl text-green-500"></i>
                  </div>
                  <div class="text-left">
                    <p class="text-sm font-bold text-gray-800">{{ jsonFileName }}</p>
                    <p class="text-xs text-green-600 font-medium">{{ jsonPreview.length }} soal terdeteksi</p>
                  </div>
                </div>

                <!-- Mini preview list -->
                <div class="bg-white rounded-lg border border-gray-200 max-h-36 overflow-y-auto text-left divide-y divide-gray-100 custom-scrollbar">
                  <div v-for="(q, i) in jsonPreview.slice(0, 5)" :key="i" class="px-3 py-2 flex items-start gap-2">
                    <span class="text-[10px] font-bold text-indigo-500 bg-indigo-50 rounded px-1.5 py-0.5 mt-0.5 shrink-0">{{ i + 1 }}</span>
                    <p class="text-xs text-gray-600 line-clamp-1">{{ q.question }}</p>
                  </div>
                  <div v-if="jsonPreview.length > 5" class="px-3 py-2 text-center">
                    <span class="text-[10px] text-gray-400">...dan {{ jsonPreview.length - 5 }} soal lainnya</span>
                  </div>
                </div>

                <div class="flex justify-center gap-3 pt-1">
                  <button 
                    @click.stop="cancelJsonImport" 
                    class="px-4 py-1.5 text-xs font-medium text-gray-600 bg-gray-100 hover:bg-gray-200 rounded-lg transition"
                  >
                    <i class="fas fa-times mr-1"></i> Batal
                  </button>
                  <button 
                    @click.stop="confirmJsonImport" 
                    class="px-5 py-1.5 text-xs font-bold text-white bg-gradient-to-r from-green-500 to-emerald-600 hover:from-green-600 hover:to-emerald-700 rounded-lg shadow-sm transition-all hover:shadow-md"
                  >
                    <i class="fas fa-download mr-1"></i> Import {{ jsonPreview.length }} Soal
                  </button>
                </div>
              </div>
            </div>
            <input type="file" ref="jsonInput" @change="handleJsonSelect" accept=".json" class="hidden" />
          </div>

          <!-- Divider -->
          <div class="flex items-center gap-4 mb-6">
            <div class="flex-1 border-t border-gray-200"></div>
            <span class="text-xs font-semibold text-gray-400 uppercase tracking-wider">atau tambah manual</span>
            <div class="flex-1 border-t border-gray-200"></div>
          </div>

          <div class="flex flex-col lg:flex-row gap-8">
            <!-- Form Input Soal -->
            <div class="flex-1 space-y-6">
              <div class="bg-gray-50 p-6 rounded-lg border border-gray-200">
                <h3 class="font-bold text-gray-800 mb-4 border-b border-gray-200 pb-2">
                    {{ isEditingIndex !== null ? 'Edit Soal' : 'Form Soal Baru' }}
                </h3>

                <!-- Subjek & Level -->
                <div class="grid grid-cols-1 md:grid-cols-3 gap-4 mb-4">
                  <div>
                    <label class="block text-xs font-semibold text-gray-600 mb-1">Subjek</label>
                    <select 
                        v-model="currentQuestion.subject_id" 
                        @change="handleSubjectChange"
                        class="w-full border border-gray-300 rounded-lg p-2 text-sm"
                    >
                        <option :value="null" disabled>-- Pilih Subjek --</option>
                        <option v-for="subject in subjects" :key="subject.id" :value="subject.id">
                        {{ subject.title }} ({{ subject.code }})
                        </option>
                        <option value="NEW_SUBJECT" class="font-bold text-blue-600">+ Buat Subjek Baru</option>
                    </select>
                  </div>
                  <div>
                    <label class="block text-xs font-semibold text-gray-600 mb-1">Level</label>
                    <select v-model="currentQuestion.level" class="w-full border border-gray-300 rounded-lg p-2 text-sm">
                        <option value="easy">Easy</option>
                        <option value="medium">Medium</option>
                        <option value="hard">Hard</option>
                    </select>
                  </div>
                  <div>
                    <label class="block text-xs font-semibold text-gray-600 mb-1">Poin</label>
                    <input v-model.number="currentQuestion.mark" type="number" class="w-full border border-gray-300 rounded-lg p-2 text-sm" />
                  </div>
                </div>

                <!-- Gambar & Pertanyaan -->
                <div class="grid grid-cols-1 gap-4 mb-4">
                     <!-- Image Upload Area -->
                    <div 
                        @dragover.prevent="isDraggingImage = true" 
                        @dragleave.prevent="isDraggingImage = false" 
                        @drop.prevent="handleDropImage"
                        :class="['border-2 border-dashed rounded-lg p-4 text-center transition-colors', isDraggingImage ? 'border-blue-500 bg-blue-50' : 'border-gray-300 bg-white']"
                    >
                        <template v-if="!currentQuestion.imageUrl">
                            <p class="text-sm text-gray-600 mb-2">Seret gambar ke sini atau</p>
                            <button @click="triggerImageInput" class="px-4 py-1.5 text-xs font-medium bg-gray-200 hover:bg-gray-300 rounded text-gray-700 transition">Pilih File</button>
                        </template>
                        <template v-else>
                            <div class="relative inline-block">
                                <img :src="currentQuestion.imageUrl" alt="Preview" class="h-32 object-contain rounded border border-gray-200">
                                <button @click="removeImage" class="absolute -top-2 -right-2 bg-red-500 text-white rounded-full w-6 h-6 flex items-center justify-center shadow-sm hover:bg-red-600 text-xs">
                                    <i class="fas fa-times"></i>
                                </button>
                            </div>
                            <p class="text-xs text-gray-500 mt-1 truncate max-w-xs mx-auto">{{ uploadedImageName }}</p>
                        </template>
                        <input type="file" ref="imageInput" @change="handleImageSelect" accept="image/*" class="hidden" />
                    </div>

                    <textarea 
                        v-model="currentQuestion.question" 
                        rows="3" 
                        placeholder="Tulis pertanyaan Anda di sini..." 
                        class="w-full border border-gray-300 rounded-lg p-3 focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                    ></textarea>
                </div>

                <!-- Jawaban -->
                <div class="space-y-3 mb-6">
                    <label class="block text-xs font-semibold text-gray-600">Pilihan Jawaban (Klik lingkaran untuk memilih kunci jawaban)</label>
                    <div v-for="(answer, index) in currentQuestion.answers" :key="index" class="flex items-center gap-3">
                        <span class="w-6 h-6 flex items-center justify-center rounded-full bg-gray-200 text-xs font-bold text-gray-600">{{ String.fromCharCode(65 + index) }}</span>
                        <input 
                            v-model="answer.text" 
                            :placeholder="'Jawaban ' + (index + 1)" 
                            class="flex-1 border border-gray-300 rounded px-3 py-1.5 text-sm focus:outline-none focus:border-blue-500"
                        />
                        <button 
                            @click="toggleCorrectAnswer(index)" 
                            :class="['w-6 h-6 rounded-full border-2 flex items-center justify-center transition-all', answer.isCorrect ? 'bg-green-500 border-green-500 text-white' : 'border-gray-300 hover:border-gray-400']"
                            title="Tandai sebagai jawaban benar"
                        >
                            <i v-if="answer.isCorrect" class="fas fa-check text-xs"></i>
                        </button>
                    </div>
                </div>

                <div class="flex justify-end gap-3">
                    <button v-if="isEditingIndex !== null" @click="cancelEdit" class="px-4 py-2 text-sm text-gray-600 bg-gray-100 hover:bg-gray-200 rounded-lg transition">Batal Edit</button>
                    <button @click="param => addDataToTable()" class="px-6 py-2 text-sm font-semibold text-white bg-blue-600 hover:bg-blue-700 rounded-lg shadow-sm transition">
                        {{ isEditingIndex !== null ? 'Update Soal' : 'Tambah ke Daftar' }}
                    </button>
                </div>

              </div>
            </div>

            <!-- Daftar Soal Preview -->
            <div class="w-full lg:w-1/3">
                <div class="bg-blue-50 p-6 rounded-lg border border-blue-100 h-full">
                    <h3 class="font-bold text-blue-900 mb-4 flex items-center justify-between">
                        <span>Daftar Soal</span>
                        <span class="bg-blue-200 text-blue-800 text-xs px-2 py-1 rounded-full">{{ questionList.length }}</span>
                    </h3>

                    <div class="space-y-3 max-h-[600px] overflow-y-auto pr-2 custom-scrollbar">
                        <div v-if="questionList.length === 0" class="text-center py-10 text-gray-400 text-sm">
                            Belum ada soal ditambahkan.
                        </div>
                        <div v-else v-for="(q, idx) in questionList" :key="idx" class="bg-white p-3 rounded shadow-sm border border-gray-100 group relative transition-all hover:shadow-md hover:border-blue-200">
                            <div class="flex justify-between items-start">
                                <div class="pr-6 flex-1">
                                    <div class="flex items-center gap-2 mb-1">
                                      <span class="text-xs font-bold text-blue-600">Soal {{ idx + 1 }}</span>
                                      <span v-if="getSubjectName(q.subject_id)" class="text-[10px] bg-indigo-100 text-indigo-700 px-1.5 py-0.5 rounded-full font-medium truncate max-w-[120px]">
                                        {{ getSubjectName(q.subject_id) }}
                                      </span>
                                    </div>
                                    <p class="text-sm text-gray-700 line-clamp-2">{{ q.question }}</p>
                                </div>
                                <div class="flex flex-col gap-1">
                                    <button @click="editQuestion(idx)" class="text-gray-400 hover:text-blue-600 p-1"><i class="fas fa-edit"></i></button>
                                    <button @click="removeQuestion(idx)" class="text-gray-400 hover:text-red-600 p-1"><i class="fas fa-trash"></i></button>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
          </div>

          <!-- Bottom Action Bar -->
          <div class="mt-8 pt-6 border-t border-gray-100 flex justify-between items-center">
            <button @click="step = 1" class="text-gray-500 hover:text-gray-700 text-sm font-medium flex items-center gap-2">
                <i class="fas fa-arrow-left"></i> Kembali ke Detail
            </button>
            <button 
                @click="saveAll" 
                :disabled="loading || questionList.length === 0"
                class="px-8 py-3 bg-green-600 hover:bg-green-700 text-white font-bold rounded-lg shadow-lg transform transition hover:-translate-y-0.5 disabled:opacity-50 disabled:cursor-not-allowed"
            >
                <span v-if="loading"><i class="fas fa-spinner fa-spin mr-2"></i>Sedang Menyimpan...</span>
                <span v-else>Simpan Ujian & {{ questionList.length }} Soal</span>
            </button>
          </div>

        </div>
      </div>
    </div>

    <!-- Modal Create Subject -->
    <div v-if="showSubjectModal" class="fixed inset-0 z-50 flex items-center justify-center bg-black/50 backdrop-blur-sm p-4">
      <div class="bg-white rounded-xl shadow-2xl p-6 w-full max-w-md transform transition-all scale-100">
        <h3 class="text-lg font-bold text-gray-800 mb-4">Buat Subjek Baru</h3>
        
        <div class="space-y-4">
            <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">Judul Subjek (Title)</label>
            <input v-model="newSubject.title" type="text" class="w-full p-2.5 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500" placeholder="Contoh: Matematika Dasar">
            </div>
            
            <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">Kode Subjek (Code)</label>
            <input v-model="newSubject.code" type="text" class="w-full p-2.5 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500" placeholder="Contoh: MAT-101">
            </div>
        </div>

        <div class="flex justify-end gap-3 mt-6">
          <button @click="closeSubjectModal" class="px-4 py-2 text-gray-600 bg-gray-100 hover:bg-gray-200 rounded-lg transition">Batal</button>
          <button @click="handleCreateSubject" class="px-4 py-2 text-white bg-blue-600 hover:bg-blue-700 rounded-lg transition shadow-sm" :disabled="isCreatingSubject">
            {{ isCreatingSubject ? 'Menyimpan...' : 'Simpan' }}
          </button>
        </div>
      </div>
    </div>

  </div>
</template>

<script setup>
import { ref, reactive, computed, onMounted, watch } from "vue";
import { useRoute, useRouter } from "vue-router";
import { createExam } from "../../provider/exam.provider.js";
import { createQuestionWithOptions } from "../../provider/question.provider.js";
import { getPaginatedSubjects, createSubject } from "../../provider/subject.provider.js";
import { useGetCurrentUser } from "../../hooks/useGetCurrentUser.js";
import { usePopup } from "../../hooks/usePopup"; // Adjust if hooks path is different
import { API_BASE_URL } from '../../core/constant'; // Needed for manual constructImageUrl if used, but simple local preview is handled

const { showSuccess, showError, showConfirm } = usePopup();
const router = useRouter();
const route = useRoute();
const { user } = useGetCurrentUser();

const step = ref(1);
const loading = ref(false);

const isAdminRoute = computed(() => route.path.startsWith('/admin'));

// --- STATE: EXAM FORM ---
const examForm = reactive({
  title: "",
  description: "",
  difficulty: "",
  started_at: "",
  finished_at: "",
  long_time: "",
  subject_id: null,
});

// --- STATE: QUESTION FORM ---
const subjects = ref([]);
const questionList = ref([]);
const isEditingIndex = ref(null);

// Helper to get subject name by id
const getSubjectName = (subjectId) => {
  if (!subjectId) return null;
  const subject = subjects.value.find(s => s.id === subjectId);
  return subject ? subject.title : null;
};

const createEmptyQuestion = () => ({
  subject_id: null,
  level: 'easy',
  mark: 3,
  imageUrl: null,
  imageFile: null,
  question: '',
  answers: [
    { text: '', isCorrect: false },
    { text: '', isCorrect: false },
    { text: '', isCorrect: false },
    { text: '', isCorrect: false },
  ]
});

const currentQuestion = reactive(createEmptyQuestion());
const isDraggingImage = ref(false);
const uploadedImageName = ref(null);
const imageInput = ref(null);

// JSON Import State
const jsonInput = ref(null);
const isDraggingJson = ref(false);
const jsonPreview = ref(null);
const jsonFileName = ref('');

// Subject Modal State
const showSubjectModal = ref(false);
const isCreatingSubject = ref(false);
const newSubject = reactive({ title: '', code: '' });

// --- LIFECYCLE ---
onMounted(() => {
  fetchSubjects();
});

// --- METHODS: EXAM FORM ---
const nextStep = () => {
    // Validate Step 1
    if (!examForm.title || !examForm.difficulty || !examForm.started_at || !examForm.finished_at || !examForm.long_time) {
        showError("Data Tidak Lengkap", "Silakan lengkapi semua data wajib pada detail ujian.");
        return;
    }
    step.value = 2;
};

// --- METHODS: QUESTION HANDLING ---

const fetchSubjects = async () => {
  try {
    const response = await getPaginatedSubjects(100, 0);
    // Handle differences in response structure
    if (response?.data?.data && Array.isArray(response.data.data)) {
        subjects.value = response.data.data;
    } else if (response?.data && Array.isArray(response.data)) {
        subjects.value = response.data;
    } else {
        subjects.value = [];
    }
    
    // Set default subject if available
    if (subjects.value.length > 0 && !currentQuestion.subject_id) {
        currentQuestion.subject_id = subjects.value[0].id;
    }
  } catch (error) {
    console.error("Failed to load subjects", error);
  }
};

const handleSubjectChange = (e) => {
    if (currentQuestion.subject_id === 'NEW_SUBJECT') {
        newSubject.title = '';
        newSubject.code = '';
        showSubjectModal.value = true;
        currentQuestion.subject_id = null;
    }
};

const closeSubjectModal = () => {
    showSubjectModal.value = false;
    // Reset selection to first subject if nothing selected
    if (!currentQuestion.subject_id && subjects.value.length > 0) {
        currentQuestion.subject_id = subjects.value[0].id;
    }
};

const handleCreateSubject = async () => {
    if (!newSubject.title || !newSubject.code) {
        showError("Validasi", "Judul dan Kode subjek wajib diisi!");
        return;
    }
    isCreatingSubject.value = true;
    try {
        const res = await createSubject(newSubject);
        const created = res.data;
        subjects.value.push(created);
        currentQuestion.subject_id = created.id;
        showSuccess("Berhasil", "Subjek baru dibuat!");
        closeSubjectModal();
    } catch (err) {
        showError("Gagal", "Gagal membuat subjek.");
    } finally {
        isCreatingSubject.value = false;
    }
};

// Image Handling
const triggerImageInput = () => imageInput.value.click();
const handleImageSelect = (e) => processImage(e.target.files[0]);
const handleDropImage = (e) => {
    isDraggingImage.value = false;
    processImage(e.dataTransfer.files[0]);
};

const processImage = (file) => {
    if (file && file.type.startsWith('image/')) {
        currentQuestion.imageFile = file;
        uploadedImageName.value = file.name;
        const reader = new FileReader();
        reader.onload = (e) => { currentQuestion.imageUrl = e.target.result; };
        reader.readAsDataURL(file);
    } else {
        showError("File Salah", "Hanya file gambar yang diperbolehkan.");
    }
};

// --- JSON IMPORT METHODS ---
const triggerJsonInput = () => {
    if (!jsonPreview.value) {
        jsonInput.value.click();
    }
};

const handleJsonSelect = (e) => processJsonFile(e.target.files[0]);

const handleDropJson = (e) => {
    isDraggingJson.value = false;
    const file = e.dataTransfer.files[0];
    if (file) processJsonFile(file);
};

const processJsonFile = (file) => {
    if (!file) return;
    if (!file.name.endsWith('.json')) {
        showError('Format Salah', 'Hanya file .json yang diperbolehkan.');
        return;
    }

    jsonFileName.value = file.name;
    const reader = new FileReader();
    reader.onload = (e) => {
        try {
            const raw = JSON.parse(e.target.result);
            const questions = Array.isArray(raw) ? raw : (raw.questions || raw.soal || raw.data || []);

            if (!Array.isArray(questions) || questions.length === 0) {
                showError('Data Kosong', 'File JSON tidak berisi data soal yang valid.');
                jsonPreview.value = null;
                return;
            }

            // Normalize & validate each question
            const normalized = questions.map((q, idx) => {
                const questionText = q.question || q.question_text || q.pertanyaan || q.soal || '';
                const level = q.level || q.difficulty || 'easy';
                const mark = q.mark || q.score || q.poin || (level === 'hard' ? 18 : level === 'medium' ? 10 : 3);
                const subjectId = q.subject_id || null;

                // Parse answers/options
                const rawAnswers = q.answers || q.options || q.jawaban || q.pilihan || [];
                const answers = rawAnswers.map((a, i) => {
                    if (typeof a === 'string') {
                        return { text: a, isCorrect: false };
                    }
                    return {
                        text: a.text || a.option_text || a.jawaban || a.label || '',
                        isCorrect: a.isCorrect || a.is_correct || a.correct || a.benar || false
                    };
                });

                // Pad to 4 answers if fewer
                while (answers.length < 4) {
                    answers.push({ text: '', isCorrect: false });
                }

                // Handle correct_answer index/letter if no isCorrect flags
                if (!answers.some(a => a.isCorrect) && (q.correct_answer !== undefined || q.kunci !== undefined)) {
                    let correctIdx = q.correct_answer ?? q.kunci;
                    if (typeof correctIdx === 'string') {
                        correctIdx = correctIdx.toUpperCase().charCodeAt(0) - 65;
                    }
                    if (correctIdx >= 0 && correctIdx < answers.length) {
                        answers[correctIdx].isCorrect = true;
                    }
                }

                return {
                    subject_id: subjectId,
                    level,
                    mark,
                    imageUrl: null,
                    imageFile: null,
                    question: questionText,
                    answers: answers.slice(0, 4)
                };
            }).filter(q => q.question.trim() !== '');

            if (normalized.length === 0) {
                showError('Data Tidak Valid', 'Tidak ditemukan soal valid dalam file JSON.');
                jsonPreview.value = null;
                return;
            }

            jsonPreview.value = normalized;
        } catch (err) {
            console.error('JSON parse error:', err);
            showError('Gagal Membaca', 'File JSON tidak valid atau rusak.');
            jsonPreview.value = null;
        }
    };
    reader.readAsText(file);
};

const confirmJsonImport = () => {
    if (!jsonPreview.value) return;

    // If no subject_id set, use the first available subject
    const defaultSubjectId = subjects.value.length > 0 ? subjects.value[0].id : null;

    jsonPreview.value.forEach(q => {
        if (!q.subject_id) {
            q.subject_id = defaultSubjectId;
        }
        questionList.value.push(q);
    });

    showSuccess('Import Berhasil', `${jsonPreview.value.length} soal berhasil ditambahkan ke daftar.`);
    cancelJsonImport();
};

const cancelJsonImport = () => {
    jsonPreview.value = null;
    jsonFileName.value = '';
    if (jsonInput.value) jsonInput.value.value = null;
};

const removeImage = () => {
    currentQuestion.imageUrl = null;
    currentQuestion.imageFile = null;
    uploadedImageName.value = null;
    if (imageInput.value) imageInput.value.value = null;
};

// Answer Handling
const toggleCorrectAnswer = (idx) => {
    currentQuestion.answers.forEach((ans, i) => {
        ans.isCorrect = (i === idx);
    });
};

// Add/Update List
const addDataToTable = () => {
    // Validation
    if (!currentQuestion.subject_id) { showError("Validasi", "Pilih subjek dulu."); return; }
    if (!currentQuestion.question.trim()) { showError("Validasi", "Pertanyaan wajib diisi."); return; }
    if (currentQuestion.answers.every(a => !a.text.trim())) { showError("Validasi", "Minimal satu jawaban diisi."); return; }
    if (!currentQuestion.answers.some(a => a.isCorrect)) { showError("Validasi", "Pilih kunci jawaban yang benar."); return; }

    // Clone data
    const questionObj = JSON.parse(JSON.stringify(currentQuestion));
    questionObj.imageFile = currentQuestion.imageFile; // Copy file reference manually

    if (isEditingIndex.value !== null) {
        // Update existing
        questionList.value[isEditingIndex.value] = questionObj;
        isEditingIndex.value = null;
        showSuccess("Updated", "Soal diperbarui.");
    } else {
        // Add new
        questionList.value.push(questionObj);
    }

    // Reset Form
    resetQuestionForm();
};

const editQuestion = (idx) => {
    const q = questionList.value[idx];
    Object.assign(currentQuestion, JSON.parse(JSON.stringify(q)));
    currentQuestion.imageFile = q.imageFile; // Restore file ref
    isEditingIndex.value = idx;
    uploadedImageName.value = q.imageFile ? q.imageFile.name : null;
};

const removeQuestion = async (idx) => {
   // const confirmed = await showConfirm("Hapus?", "Yakin hapus soal ini dari daftar?");
   // if (confirmed) {
        questionList.value.splice(idx, 1);
   // }
};

const cancelEdit = () => {
    isEditingIndex.value = null;
    resetQuestionForm();
};

const resetQuestionForm = () => {
    const savedSubject = currentQuestion.subject_id; // Keep last subject selection
    Object.assign(currentQuestion, createEmptyQuestion());
    currentQuestion.subject_id = savedSubject;
    isDraggingImage.value = false;
    uploadedImageName.value = null;
    // reset answers explicitly if needed
};

// Watcher for Difficulty Points
watch(() => currentQuestion.level, (val) => {
    if (val === 'easy') currentQuestion.mark = 3;
    if (val === 'medium') currentQuestion.mark = 10;
    if (val === 'hard') currentQuestion.mark = 18;
});

// --- METHODS: SAVE ALL ---
const saveAll = async () => {
    if (!user?.value) { showError("Akses Ditolak", "Login dlu gan."); return; }

    loading.value = true;
    try {
        // 1. Create Exam
        const examPayload = {
            title: examForm.title.trim(),
            description: examForm.description.trim(),
            difficulty: examForm.difficulty,
            started_at: new Date(examForm.started_at).toISOString(),
            finished_at: new Date(examForm.finished_at).toISOString(),
            long_time: Number(examForm.long_time),
            creator_id: user.value.id,
            subject_id: examForm.subject_id || null,
        };

        const examRes = await createExam(examPayload);
        console.log('createExam response:', JSON.stringify(examRes));
        const newExamId = examRes?.id || examRes?.data?.id || examRes?.data?.data?.id;

        if (!newExamId) throw new Error("Gagal mendapatkan ID ujian baru.");

        // 2. Create Questions
        let successCount = 0;
        for (const q of questionList.value) {
            const qPayload = {
                exam_id: newExamId,
                subject_id: q.subject_id,
                question_text: q.question,
                difficulty: q.level,
                score: q.mark,
                creator_id: user.value.id,
                options: q.answers.filter(a => a.text.trim()).map((a, i) => ({
                    option_label: String.fromCharCode(65 + i),
                    option_text: a.text,
                    is_correct: a.isCorrect
                }))
            };
             // Add image if exists
            if (q.imageFile) {
                qPayload.image = q.imageFile;
            }

            try {
                await createQuestionWithOptions(qPayload);
                successCount++;
            } catch (qErr) {
                console.error("Gagal save soal:", qErr);
                // Continue saving others
            }
        }

        await showSuccess("Selesai!", `Ujian dibuat dengan ${successCount} soal berhasil disimpan.`);
        
        // Redirect
        const returnRouteName = isAdminRoute.value ? 'AdminManageExam' : 'DosenManageExam';
        router.push({ name: returnRouteName });

    } catch (err) {
        console.error(err);
        showError("Gagal", "Terjadi kesalahan saat menyimpan data.");
    } finally {
        loading.value = false;
    }
};

</script>

<style scoped>
/* Custom scrollbar for question list */
.custom-scrollbar::-webkit-scrollbar {
  width: 6px;
}
.custom-scrollbar::-webkit-scrollbar-track {
  background: #f1f1f1; 
}
.custom-scrollbar::-webkit-scrollbar-thumb {
  background: #cbd5e1; 
  border-radius: 3px;
}
.custom-scrollbar::-webkit-scrollbar-thumb:hover {
  background: #94a3b8; 
}
</style>