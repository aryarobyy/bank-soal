<template>
  <div class="w-full p-6 sm:p-8 mx-auto bg-white rounded-2xl shadow-xl">
    <div class="flex justify-between items-center mb-6">
      <h2 class="text-2xl font-bold text-dark-text">
        {{ isEditMode ? `Edit Soal (ID: ${questionId})` : 'Buat Soal Manual' }}
      </h2>
      <button @click="saveSoal" class="px-5 py-2 text-sm font-semibold border border-gray-300 rounded-md text-gray-700 hover:bg-gray-100 transition-colors">
        {{ isEditMode ? 'Update Soal' : 'Simpan Semua & Selesai' }}
      </button>
    </div>
    
    <div class="p-8 border border-gray-200 rounded-lg">
      <div class="mb-8 grid grid-cols-1 md:grid-cols-3 gap-6">
        <div>
          <label for="subject" class="block text-sm font-medium text-gray-700 mb-1">Subjek Mata Kuliah*</label>
          <select id="subject" v-model="currentSoal.subject_id" class="w-full px-3 py-2 bg-gray-50 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500">
            <option :value="null" disabled>-- Pilih Subjek --</option>
            <option v-for="subject in subjects" :key="subject.id" :value="subject.id">
              {{ subject.title }}
            </option>
          </select>
        </div>
        <div>
          <label for="level" class="block text-sm font-medium text-gray-700 mb-1">Level Kesulitan*</label>
          <select id="level" v-model="currentSoal.level" class="w-full px-3 py-2 bg-gray-50 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500">
            <option value="easy">Easy</option>
            <option value="medium">Medium</option>
            <option value="hard">Hard</option>
          </select>
        </div>
        <div>
          <label for="mark" class="block text-sm font-medium text-gray-700 mb-1">Mark*</label>
          <input id="mark" v-model.number="currentSoal.mark" type="number" class="w-full px-3 py-2 bg-gray-50 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500">
        </div>
      </div>
      
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6 mb-8">
        <div 
          @dragover.prevent="isDraggingImage = true" @dragleave.prevent="isDraggingImage = false" @drop.prevent="handleDropImage"
          :class="isDraggingImage ? 'border-blue-500 bg-blue-50' : 'border-gray-300'"
          class="flex flex-col items-center justify-center p-8 text-center border-2 border-dashed rounded-lg transition-colors min-h-[200px]"
        >
          <template v-if="!currentSoal.imageUrl">
            <p class="text-lg font-semibold text-gray-700">Masukkan Gambar</p>
            <p class="my-2 text-gray-500">or</p>
            <button @click="triggerImageInput" class="px-6 py-2 font-bold text-white transition-opacity rounded-lg bg-teal-500 hover:bg-teal-600">Select file</button>
          </template>
          <template v-else>
            <img :src="currentSoal.imageUrl" alt="Uploaded Image" class="max-h-48 max-w-full object-contain mb-4">
            <p class="text-sm text-gray-600 mb-2">{{ uploadedImageName }}</p>
            <button @click="removeImage" class="text-red-500 hover:underline text-sm">Remove Image</button>
          </template>
          <input type="file" ref="imageInput" @change="handleImageSelect" accept="image/*" class="hidden" />
        </div>
        <textarea v-model="currentSoal.question" rows="8" placeholder="Is there _____ milk in the fridge?" class="w-full p-4 bg-gray-50 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 resize-none"></textarea>
      </div>

      <div class="grid grid-cols-1 md:grid-cols-2 gap-6 mb-8">
        <div v-for="(answer, index) in currentSoal.answers" :key="index" class="flex items-center p-4 bg-gray-100 border border-gray-200 rounded-lg">
          <span :class="[answerColors[index % answerColors.length], 'flex items-center justify-center w-8 h-8 font-bold text-white rounded-full mr-4 flex-shrink-0']">{{ String.fromCharCode(65 + index) }}</span>
          <input v-model="answer.text" :placeholder="'Add answer ' + (index + 1)" class="flex-1 w-full bg-transparent focus:outline-none text-gray-800" />
          <button @click="toggleCorrectAnswer(index)" class="ml-4 w-6 h-6 rounded-full border-2 transition-colors flex-shrink-0" :class="answer.isCorrect ? 'bg-blue-500 border-blue-500' : 'bg-white border-gray-300 hover:border-gray-400'"></button>
        </div>
      </div>
      
      <div class="flex justify-end" v-if="!isEditMode">
        <button @click="addSoalToList" class="px-6 py-3 font-bold text-white transition-opacity rounded-lg bg-blue-600 hover:bg-blue-700">
          Tambah Soal ke Daftar
        </button>
      </div>
    </div>
    
    <div v-if="soalList.length > 0 && !isEditMode" class="mt-12">
      <h3 class="mb-4 text-xl font-bold text-gray-800">Daftar Soal ({{ soalList.length }} soal ditambahkan)</h3>
      <div class="space-y-4">
        <div v-for="(soal, index) in soalList" :key="index" class="flex items-center justify-between p-4 bg-gray-50 rounded-lg border">
          <p class="text-gray-700">{{ index + 1 }}. {{ soal.question.substring(0, 70) }}...</p>
          <button @click="removeSoalFromList(index)" class="text-red-500 hover:text-red-700"><i class="fas fa-trash"></i></button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { createQuestionWithOptions, getQuestionById, updateQuestion } from '../../provider/question.provider';
import { useGetCurrentUser } from '../../hooks/useGetCurrentUser';
// Pastikan path constant ini benar
import { API_BASE_URL } from '../../core/constant'; 

// Data subjek ini HANYA CONTOH (hardcoded)
const subjects = [
  { id: 1, title: 'Kalkulus', code: 'MFG-101' },
  { id: 2, title: 'Matematika Diskrit', code: 'TIF-1203' },
  { id: 3, title: 3, title: 'Teori Bahasa dan Automata', code: 'TIF-2204' },
  { id: 4, title: 4, title: 'Basis Data Lanjut', code: 'TIF-2206' },
  { id: 5, title: 5, title: 'Metode Numerik', code: 'TIF-3107' },
];

const createEmptySoal = () => ({
  subject_id: subjects.length > 0 ? subjects[0].id : null,
  level: 'easy',
  mark: 3,
  imageUrl: null,   // URL untuk Preview
  imageFile: null,  // File Object mentah untuk diupload
  isDeleteImage: false, // <-- Flag untuk menghapus gambar lama di backend
  question: '',
  answers: [
    { text: '', isCorrect: false }, { text: '', isCorrect: false },
    { text: '', isCorrect: false }, { text: '', isCorrect: false },
  ]
});

export default {
  name: 'CreateManualSoal',

  setup() {
    const { user } = useGetCurrentUser();
    return { user };
  },

  data() {
    return {
      isEditMode: false,
      questionId: null,
      returnSubjectId: null, 
      returnPage: null, 
      subjects: subjects, 
      isDraggingImage: false,
      uploadedImageName: null,
      currentSoal: createEmptySoal(),
      soalList: [],
      answerColors: ['bg-red-500', 'bg-blue-500', 'bg-yellow-500', 'bg-green-500'],
      isLoadingData: false,
    };
  },
  
  computed: {
    isAdminRoute() {
      return this.$route.path.startsWith('/admin/soal');
    },
    listRouteName() {
      return this.isAdminRoute ? 'AdminSoalList' : 'DosenSoalList';
    }
  },
  
  methods: {
    constructImageUrl(serverPath) {
      if (!serverPath) return null;
      
      if (serverPath.startsWith('http')) {
        return serverPath;
      }

      const cleanPath = serverPath.startsWith('.') ? serverPath.substring(1) : serverPath;
      return `${API_BASE_URL}${cleanPath}`;
    },

    async fetchQuestionData(id) {
      try {
        const response = await getQuestionById(id);
        const questionData = response.data;
        if (!questionData) throw new Error("Data soal tidak ditemukan");

        this.currentSoal = {
          subject_id: questionData.subject_id,
          level: questionData.difficulty,
          mark: questionData.score,
          question: questionData.question_text,
          answers: this.prepareAnswers(questionData.options),
          imageUrl: this.constructImageUrl(questionData.img_url), 
          imageFile: null, 
          isDeleteImage: false, // Reset flag saat load
        };
        
        if (this.currentSoal.imageUrl) {
          this.uploadedImageName = "Gambar tersimpan di server";
        }

      } catch (error) {
        console.error("Gagal mengambil data soal:", error);
        alert('Gagal memuat data soal untuk diedit.');
        this.$router.push({ name: this.listRouteName });
      }
    },
    
    prepareAnswers(options = []) {
      const answers = options.map(opt => ({ text: opt.option_text, isCorrect: opt.is_correct }));
      while (answers.length < 4) {
        answers.push({ text: '', isCorrect: false });
      }
      return answers.slice(0, 4);
    },

    async saveSoal() {
      const creatorId = this.user?.id; 
      if (!creatorId) {
        alert('Gagal mendapatkan ID pengguna. Silakan login ulang.');
        return;
      }
      const hardcodedExamId = 2; 

      try {
        if (this.isEditMode) {
          // --- ALUR EDIT ---
          if (!this.currentSoal.question.trim() || !this.currentSoal.subject_id || !this.currentSoal.answers.some(a => a.isCorrect)) {
            alert('Harap lengkapi semua field yang wajib diisi.');
            return;
          }
          
          const payload = this.formatPayload(this.currentSoal, creatorId, hardcodedExamId);
          await updateQuestion(this.questionId, payload);
          alert('Soal berhasil diperbarui!');
          
          const query = {};
          if (this.returnSubjectId) query.subject_id = this.returnSubjectId;
          if (this.returnPage) query.page = this.returnPage; 
          this.$router.push({ name: this.listRouteName, query: query });

        } else {
          // --- ALUR BUAT BARU ---
          const questionsToSave = this.soalList.length > 0 ? [...this.soalList] : [];
          if(this.currentSoal.question.trim()){
              if (this.currentSoal.answers.every(a => !a.text.trim())) {
                alert('Isi jawaban terlebih dahulu.'); return; 
              }
              if (!this.currentSoal.answers.some(a => a.isCorrect)) {
                alert('Pilih satu jawaban benar.'); return; 
              }
              questionsToSave.push(this.currentSoal);
          }
          if (questionsToSave.length === 0) {
            alert('Tidak ada soal untuk disimpan.'); return;
          }

          for (const soal of questionsToSave) {
            if (!soal.question.trim()) continue; 
            const payload = this.formatPayload(soal, creatorId, hardcodedExamId);
            await createQuestionWithOptions(payload);
          }
          
          alert(`${questionsToSave.length} soal berhasil disimpan!`);
          this.$router.push({ name: this.listRouteName, query: { show_last_page: 'true' } });
        }
      } catch (error) {
        console.error("Gagal menyimpan:", error.response?.data || error);
        const translatedMessage = this.translateBackendError(error);
        alert(translatedMessage);
      }
    },

    translateBackendError(error) {
      const rawMessage = error?.response?.data?.message || 'Terjadi kesalahan.';
      const message = rawMessage.toLowerCase();
      if (message.includes("easy difficulty must be between")) return "Skor Easy harus 3-8.";
      if (message.includes("medium difficulty must be between")) return "Skor Medium harus 10-15.";
      if (message.includes("hard difficulty must be between")) return "Skor Hard harus 18-23.";
      return rawMessage;
    },
    
    // ## FUNGSI FORMAT PAYLOAD DENGAN LOGIC img_delete ##
    formatPayload(soal, creatorId, examId) {
        const payload = {
            exam_id: examId,
            creator_id: creatorId,
            subject_id: soal.subject_id,
            question_text: soal.question,
            difficulty: soal.level,
            score: soal.mark,
            options: soal.answers
                .filter(a => a.text.trim() !== '') 
                .map((a, index) => ({
                    option_label: String.fromCharCode(65 + index),
                    option_text: a.text,
                    is_correct: a.isCorrect,
                })),
        };

        // Kirim file mentah di bawah key 'image' (diperlukan untuk upload)
        if (soal.imageFile) {
            payload.image = soal.imageFile;
        } 
        
        // Kirim flag img_delete: true HANYA di mode Edit jika user menekan 'Remove' 
        // DAN TIDAK ada file baru yang di-upload.
        if (this.isEditMode && soal.isDeleteImage && !soal.imageFile) {
          payload.img_delete = true; 
        }

        return payload;
    },
    
    triggerImageInput() { this.$refs.imageInput.click(); },
    handleImageSelect(event) { this.processImage(event.target.files[0]); },
    handleDropImage(event) { this.isDraggingImage = false; this.processImage(event.dataTransfer.files[0]); },
    
    processImage(file) {
      if (file && file.type.startsWith('image/')) {
        this.currentSoal.imageFile = file; 
        this.uploadedImageName = file.name;
        
        // Penting: Reset isDeleteImage jika user memilih file baru
        this.currentSoal.isDeleteImage = false; 

        const reader = new FileReader();
        reader.onload = (e) => { this.currentSoal.imageUrl = e.target.result; };
        reader.readAsDataURL(file);
      } else {
        alert('Hanya file gambar yang diperbolehkan!');
      }
    },
    
    removeImage() {
      this.currentSoal.imageUrl = null;
      this.currentSoal.imageFile = null; 
      this.uploadedImageName = null;
      
      // ## PERBAIKAN: Set flag isDeleteImage saat di mode Edit ##
      if (this.isEditMode) {
        this.currentSoal.isDeleteImage = true;
      }
      
      if (this.$refs.imageInput) {
        this.$refs.imageInput.value = null; 
      }
    },
    
    toggleCorrectAnswer(selectedIndex) {
      this.currentSoal.answers.forEach((answer, index) => {
        answer.isCorrect = (index === selectedIndex);
      });
    },
    
    // ## FUNGSI addSoalToList YANG MENYEBABKAN BUG BATCH SAVE IMAGE ##
    addSoalToList() {
      if (!this.currentSoal.subject_id) { alert('Pilih subjek!'); return; }
      if (!this.currentSoal.question.trim()) { alert('Soal kosong!'); return; }
      if (this.currentSoal.answers.every(a => !a.text.trim())) { alert('Jawaban kosong!'); return; }
      if (!this.currentSoal.answers.some(a => a.isCorrect)) { alert('Pilih jawaban benar!'); return; }

      // 1. Buat salinan JSON data (yang menghilangkan File object)
      const newSoal = JSON.parse(JSON.stringify(this.currentSoal));
      
      // 2. KOREKSI KRITIS: Salin kembali referensi File object yang hilang
      // Ini memastikan File object yang asli (binary data) masuk ke item list.
      newSoal.imageFile = this.currentSoal.imageFile;

      this.soalList.push(newSoal);
      
      const savedSubjectId = this.currentSoal.subject_id;
      
      // Clear form state
      this.currentSoal = createEmptySoal();
      this.currentSoal.subject_id = savedSubjectId; 
      
      // Clear image input field di UI (input file)
      this.removeImage(); 
      alert('Soal ditambahkan ke daftar!');
    },
    
    removeSoalFromList(index) {
        if (confirm('Hapus soal dari daftar?')) {
            this.soalList.splice(index, 1);
        }
    },
  },

  watch: {
    'currentSoal.level'(newLevel) {
      if (this.isLoadingData) return;
      switch (newLevel) {
        case 'easy': this.currentSoal.mark = 3; break;
        case 'medium': this.currentSoal.mark = 10; break;
        case 'hard': this.currentSoal.mark = 18; break;
        default: this.currentSoal.mark = 3;
      }
    }
  },
  
  created() {
    const id = this.$route.params.id;
    const returnId = this.$route.query.return_subject_id; 
    const returnPg = this.$route.query.return_page; 
    
    if (returnId) this.returnSubjectId = returnId; 
    if (returnPg) this.returnPage = parseInt(returnPg, 10);
    
    if (id) {
      this.isEditMode = true;
      this.questionId = id;
      this.isLoadingData = true;
      this.fetchQuestionData(id).finally(() => {
        this.isLoadingData = false;
      });
    }
  }
};
</script>