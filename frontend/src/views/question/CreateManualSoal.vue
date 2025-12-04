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
          <select 
            id="subject" 
            v-model="currentSoal.subject_id" 
            @change="handleSubjectChange"
            class="w-full px-3 py-2 bg-gray-50 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
          >
            <option :value="null" disabled>-- Pilih Subjek --</option>
            <option v-for="subject in subjects" :key="subject.id" :value="subject.id">
              {{ subject.title }} ({{ subject.code }})
            </option>
            <option value="NEW_SUBJECT" class="font-bold text-blue-600">+ Buat Subjek Baru</option>
          </select>
          <p v-if="subjects.length === 0" class="text-xs text-gray-400 mt-1">Memuat mata kuliah...</p>
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

    <div v-if="showSubjectModal" class="fixed inset-0 backdrop-blur-sm bg-opacity-50 flex items-center justify-center z-50">
      <div class="bg-white rounded-lg p-6 w-96 shadow-xl">
        <h3 class="text-lg font-bold mb-4">Buat Subjek Baru</h3>
        
        <div class="mb-4">
          <label class="block text-sm font-medium text-gray-700 mb-1">Judul Subjek (Title)</label>
          <input v-model="newSubject.title" type="text" class="w-full p-2 border rounded-md" placeholder="Contoh: Matematika Dasar">
        </div>
        
        <div class="mb-4">
          <label class="block text-sm font-medium text-gray-700 mb-1">Kode Subjek (Code)</label>
          <input v-model="newSubject.code" type="text" class="w-full p-2 border rounded-md" placeholder="Contoh: MAT-101">
        </div>

        <div class="flex justify-end gap-2">
          <button @click="closeSubjectModal" class="px-4 py-2 text-gray-600 bg-gray-200 rounded-md hover:bg-gray-300">Batal</button>
          <button @click="handleCreateSubject" class="px-4 py-2 text-white bg-blue-600 rounded-md hover:bg-blue-700" :disabled="isCreatingSubject">
            {{ isCreatingSubject ? 'Menyimpan...' : 'Simpan' }}
          </button>
        </div>
      </div>
    </div>

  </div>
</template>

<script>
import { createQuestionWithOptions, getQuestionById, updateQuestion } from '../../provider/question.provider';

import { getPaginatedSubjects, createSubject } from '../../provider/subject.provider';
import { useGetCurrentUser } from '../../hooks/useGetCurrentUser';
import { API_BASE_URL } from '../../core/constant'; 

const createEmptySoal = () => ({
  subject_id: null,
  level: 'easy',
  mark: 3,
  imageUrl: null,   
  imageFile: null,  
  isDeleteImage: false, 
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
      subjects: [], 
      isDraggingImage: false,
      uploadedImageName: null,
      currentSoal: createEmptySoal(),
      soalList: [],
      answerColors: ['bg-red-500', 'bg-blue-500', 'bg-yellow-500', 'bg-green-500'],
      isLoadingData: false,
      
    
      showSubjectModal: false,
      isCreatingSubject: false,
      newSubject: {
        title: '',
        code: ''
      }
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
    async fetchSubjects() {
      try {
       
        const response = await getPaginatedSubjects(100, 0);
        
        if (response && response.data && Array.isArray(response.data.data)) {
            this.subjects = response.data.data;
        } else if (response && Array.isArray(response.data)) {
            this.subjects = response.data;
        } else {
            this.subjects = [];
        }
       
       
        if (!this.isEditMode && this.subjects.length > 0 && !this.currentSoal.subject_id) {
           this.currentSoal.subject_id = this.subjects[0].id;
        }
      } catch (error) {
        console.error("Gagal memuat subject:", error);
        this.subjects = []; 
      }
    },

    
    handleSubjectChange() {
      if (this.currentSoal.subject_id === 'NEW_SUBJECT') {
        this.newSubject = { title: '', code: '' };
        this.showSubjectModal = true;
        this.currentSoal.subject_id = null; 
      }
    },

    closeSubjectModal() {
      this.showSubjectModal = false;
    
      if (!this.currentSoal.subject_id && this.subjects.length > 0) {
        this.currentSoal.subject_id = this.subjects[0].id;
      }
    },

    async handleCreateSubject() {
      if (!this.newSubject.title || !this.newSubject.code) {
        alert("Judul dan Kode subjek harus diisi!");
        return;
      }

      this.isCreatingSubject = true;
      try {
        const res = await createSubject(this.newSubject);
        
      
        const createdSubject = res.data; 
        
        alert("Subjek berhasil dibuat!");
        
     
        this.subjects.push(createdSubject);
        
    
        this.currentSoal.subject_id = createdSubject.id;
        
        this.closeSubjectModal();
      } catch (error) {
        console.error("Gagal membuat subjek:", error);
        alert("Gagal membuat subjek baru. Coba lagi.");
      } finally {
        this.isCreatingSubject = false;
      }
    },

   
    constructImageUrl(serverPath) {
      if (!serverPath) return null;
      if (serverPath.startsWith('http')) return serverPath;
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
          isDeleteImage: false, 
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
          const questionsToSave = this.soalList.length > 0 ? [...this.soalList] : [];
          
          if(this.currentSoal.question.trim()){
              if (this.currentSoal.answers.every(a => !a.text.trim())) {
                alert('Isi jawaban terlebih dahulu.'); return; 
              }
              if (!this.currentSoal.answers.some(a => a.isCorrect)) {
                alert('Pilih satu jawaban benar.'); return; 
              }
              if (!this.currentSoal.subject_id) {
                 alert('Pilih subjek terlebih dahulu.'); return;
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
    
    formatPayload(soal, creatorId, examId) {
        const payload = {
            exam_id: examId,
    
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

        if (!this.isEditMode) {
            payload.creator_id = creatorId;
        }
   
        if (soal.imageFile) {
            payload.image = soal.imageFile;
        } 
        
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
    
    addSoalToList() {
      if (!this.currentSoal.subject_id) { alert('Pilih subjek!'); return; }
      if (!this.currentSoal.question.trim()) { alert('Soal kosong!'); return; }
      if (this.currentSoal.answers.every(a => !a.text.trim())) { alert('Jawaban kosong!'); return; }
      if (!this.currentSoal.answers.some(a => a.isCorrect)) { alert('Pilih jawaban benar!'); return; }

      const newSoal = JSON.parse(JSON.stringify(this.currentSoal));
      newSoal.imageFile = this.currentSoal.imageFile;

      this.soalList.push(newSoal);
      const savedSubjectId = this.currentSoal.subject_id;
      
      this.currentSoal = createEmptySoal();
      this.currentSoal.subject_id = savedSubjectId; 
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
    this.fetchSubjects();

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