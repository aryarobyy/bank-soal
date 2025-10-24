<template>
  <div class="w-full p-6 sm:p-8 mx-auto bg-white rounded-2xl shadow-xl">
    <div class="flex justify-end mb-6">
      <button @click="saveAllSoals" class="flex items-center gap-2 px-4 py-2 text-sm font-semibold border border-gray-300 rounded-md text-medium-text hover:bg-gray-50">
        <i class="fas fa-save"></i> Simpan Semua & Selesai
      </button>
    </div>
    
    <div class="p-6 border border-gray-200 rounded-lg">
      <div class="mb-8 grid grid-cols-1 md:grid-cols-2 gap-4">
        <div class="col-span-1">
          <label for="level" class="block text-sm font-semibold text-medium-text mb-1">Level Kesulitan*</label>
          <select id="level" v-model="currentSoal.level" class="w-full px-3 py-2 bg-gray-100 border border-gray-200 rounded-md focus:outline-none focus:ring-2 focus:ring-primary">
            <option value="easy">Easy</option>
            <option value="medium">Medium</option>
            <option value="hard">Hard</option>
          </select>
        </div>
        <div class="col-span-1">
          <label for="mark" class="block text-sm font-semibold text-medium-text mb-1">Mark*</label>
          <input id="mark" v-model.number="currentSoal.mark" type="number" class="w-full px-3 py-2 bg-gray-100 border border-gray-200 rounded-md focus:outline-none focus:ring-2 focus:ring-primary">
        </div>
      </div>
      
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6 mb-8">
        <div 
          @dragover.prevent="isDraggingImage = true"
          @dragleave.prevent="isDraggingImage = false"
          @drop.prevent="handleDropImage"
          :class="isDraggingImage ? 'border-primary bg-indigo-50' : 'border-gray-300'"
          class="flex flex-col items-center justify-center p-8 text-center border-2 border-dashed rounded-lg transition-colors min-h-[200px]"
        >
          <template v-if="!currentSoal.imageUrl">
            <i class="mb-4 text-4xl text-gray-400 fas fa-cloud-upload-alt"></i>
            <p class="text-lg font-semibold text-dark-text">Masukkan Gambar</p>
            <p class="my-2 text-gray-500">or</p>
            <button @click="triggerImageInput" class="px-6 py-2 font-bold text-white transition-opacity rounded-lg bg-teal-400 hover:opacity-90">
              Select file
            </button>
          </template>
          <template v-else>
            <img :src="currentSoal.imageUrl" alt="Uploaded Image" class="max-h-48 max-w-full object-contain mb-4">
            <p class="text-sm text-gray-600 mb-2">{{ uploadedImageName }}</p>
            <button @click="removeImage" class="text-red-500 hover:underline text-sm">Remove Image</button>
          </template>
          <input 
            type="file" 
            ref="imageInput" 
            @change="handleImageSelect" 
            accept="image/*" 
            class="hidden" 
          />
        </div>

        <textarea v-model="currentSoal.question" rows="8" placeholder="Is there _____ milk in the fridge?" class="w-full p-4 bg-gray-100 border border-gray-200 rounded-lg focus:outline-none focus:ring-2 focus:ring-primary resize-none"></textarea>
      </div>
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6 mb-8">
        <div v-for="(answer, index) in currentSoal.answers" :key="index" :class="[answer.isCorrect ? 'bg-indigo-100 border-primary' : 'bg-gray-100 border-gray-200', 'flex items-center p-4 border rounded-lg transition-all duration-200']">
          <span 
            :class="[answerColors[index % answerColors.length], 'flex items-center justify-center w-8 h-8 font-bold text-white rounded-full mr-4 flex-shrink-0']"
          >
            {{ String.fromCharCode(65 + index) }}
          </span>
          <input v-model="answer.text" :placeholder="'Add answer ' + (index + 1)" class="flex-1 w-full px-2 py-1 bg-transparent focus:outline-none text-medium-text">
          <button @click="toggleCorrectAnswer(index)" class="ml-4 p-2 rounded-full transition-colors flex-shrink-0" :class="answer.isCorrect ? 'bg-green-500 text-white' : 'bg-gray-300 text-gray-600 hover:bg-gray-400'"><i class="fas fa-check"></i></button>
        </div>
      </div>
      <div class="flex justify-end">
        <button @click="addSoalToList" class="flex items-center gap-2 px-6 py-2 font-bold text-white transition-opacity rounded-lg bg-blue-600 hover:bg-blue-700">
          <i class="fas fa-plus"></i> Tambah Soal ke Daftar
        </button>
      </div>
    </div>
    <div v-if="soalList.length > 0" class="mt-12">
      <h3 class="mb-4 text-xl font-bold text-dark-text">Daftar Soal ({{ soalList.length }} soal ditambahkan)</h3>
      <div class="space-y-4">
        <div v-for="(soal, index) in soalList" :key="index" class="flex items-center justify-between p-4 bg-gray-50 rounded-lg border">
          <p class="text-medium-text">{{ index + 1 }}. {{ soal.question.substring(0, 50) }}...</p>
          <button @click="removeSoalFromList(index)" class="text-red-500 hover:text-red-700"><i class="fas fa-trash"></i></button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { createQuestionWithOptions } from '../../provider/question.provider';
import { useLocalStorage } from '../../hooks/useLocalStorage';

const createEmptySoal = () => ({
  level: 'easy', mark: 10, imageUrl: null, question: '',
  answers: [
    { text: '', isCorrect: false }, { text: '', isCorrect: false },
    { text: '', isCorrect: false }, { text: '', isCorrect: false },
  ]
});

export default {
  name: 'CreateManualSoal',
  data() {
    return {
      // Data untuk 'exams' dan 'selectedExamId' dihapus
      isDraggingImage: false,
      uploadedImageName: null,
      currentSoal: createEmptySoal(),
      soalList: [],
      answerColors: ['bg-red-500', 'bg-blue-500', 'bg-yellow-500', 'bg-green-500']
    };
  },
  methods: {
    // Method fetchExams() dihapus
    triggerImageInput() { this.$refs.imageInput.click(); },
    handleImageSelect(event) { this.processImage(event.target.files[0]); },
    handleDropImage(event) { this.isDraggingImage = false; this.processImage(event.dataTransfer.files[0]); },
    processImage(file) {
      if (file && file.type.startsWith('image/')) {
        const reader = new FileReader();
        reader.onload = (e) => { this.currentSoal.imageUrl = e.target.result; this.uploadedImageName = file.name; };
        reader.readAsDataURL(file);
      } else {
        alert('Hanya file gambar yang diperbolehkan!'); this.currentSoal.imageUrl = null; this.uploadedImageName = null;
      }
    },
    removeImage() {
      this.currentSoal.imageUrl = null; this.uploadedImageName = null; this.$refs.imageInput.value = null;
    },
    
    toggleCorrectAnswer(selectedIndex) {
      this.currentSoal.answers.forEach((answer, index) => {
        answer.isCorrect = (index === selectedIndex) ? !answer.isCorrect : false;
      });
    },

    addSoalToList() {
      if (!this.currentSoal.question.trim()) { alert('Soal tidak boleh kosong!'); return; }
      if (this.currentSoal.answers.every(a => !a.text.trim())) { alert('Setidaknya satu jawaban harus diisi!'); return; }
      if (!this.currentSoal.answers.some(a => a.isCorrect)) { alert('Pilih setidaknya satu jawaban yang benar!'); return; }

      this.soalList.push(JSON.parse(JSON.stringify(this.currentSoal)));
      this.currentSoal = createEmptySoal();
      this.uploadedImageName = null;
      if (this.$refs.imageInput) { this.$refs.imageInput.value = null; }
      alert('Soal berhasil ditambahkan ke daftar!');
    },
    
    removeSoalFromList(index) {
      if (confirm('Anda yakin ingin menghapus soal ini dari daftar?')) {
        this.soalList.splice(index, 1);
      }
    },

    // ## FUNGSI SAVE DIUBAH UNTUK MENGGUNAKAN EXAM_ID STATIS ##
    async saveAllSoals() {
      if (this.soalList.length === 0) {
        alert('Daftar soal masih kosong. Silakan tambah soal terlebih dahulu.');
        return;
      }

      // Ambil ID pembuat soal dari Local Storage
      const { value: user } = useLocalStorage('user');
      const creatorId = user.value?.id;
      if (!creatorId) {
        alert('Gagal mendapatkan ID pengguna. Silakan login ulang.');
        return;
      }
      
      // Tentukan exam_id secara statis/hardcoded di sini
      const hardcodedExamId = 2; // Ganti dengan ID ujian yang sesuai

      try {
        for (const soal of this.soalList) {
          const payload = {
            exam_id: hardcodedExamId,
            creator_id: creatorId,
            question_text: soal.question,
            difficulty: soal.level,
            options: soal.answers
              .filter(a => a.text.trim() !== '')
              .map((a, index) => ({
                option_label: String.fromCharCode(65 + index),
                option_text: a.text,
                is_correct: a.isCorrect,
              })),
          };
          
          await createQuestionWithOptions(payload);
        }

        alert(`${this.soalList.length} soal berhasil disimpan!`);
        this.$router.push('/dosen/soal/list');

      } catch (error) {
        console.error("Gagal menyimpan soal:", error);
        alert('Terjadi kesalahan saat menyimpan soal. Periksa konsol untuk detail.');
      }
    }
  },
  // Hook mounted() tidak lagi diperlukan karena tidak ada data yang perlu di-fetch
};
</script>