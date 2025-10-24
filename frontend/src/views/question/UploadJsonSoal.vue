<template>
  <div class="w-full p-6 sm:p-8 mx-auto bg-white rounded-2xl shadow-xl">
    <div class="flex justify-end mb-6">
      <button 
        @click="saveFile" 
        :disabled="!selectedFileName"
        class="flex items-center gap-2 px-6 py-2 font-semibold text-white transition-colors rounded-lg bg-blue-600 hover:bg-blue-700 disabled:bg-gray-400 disabled:cursor-not-allowed"
      >
        <i class="fas fa-save"></i> Simpan File
      </button>
    </div>
    
    <div 
      @dragover.prevent="isDragging = true"
      @dragleave.prevent="isDragging = false"
      @drop.prevent="handleDrop"
      :class="isDragging ? 'border-blue-500 bg-blue-50' : 'border-gray-300'"
      class="flex flex-col items-center justify-center p-12 text-center border-2 border-dashed rounded-lg transition-colors"
    >
      <i class="mb-4 text-4xl text-gray-400 fas fa-cloud-upload-alt"></i>
      <p v-if="!selectedFileName" class="text-lg font-semibold text-dark-text">Drag and Drop here</p>
      <p v-if="!selectedFileName" class="my-2 text-gray-500">or</p>
      
      <p v-if="selectedFileName" class="text-lg font-semibold text-green-600">File Selected:</p>
      <p v-if="selectedFileName" class="mb-4 text-medium-text">{{ selectedFileName }}</p>

      <button @click="triggerFileInput" class="px-8 py-3 font-bold text-white transition-opacity rounded-lg bg-teal-400 hover:opacity-90">
        Select file
      </button>
      <input 
        type="file" 
        ref="fileInput" 
        @change="handleFileSelect" 
        accept=".json" 
        class="hidden" 
      />
    </div>
  </div>
</template>

<script>
// I've also updated the API call logic to be functional
import { createQuestionFromJson } from '../../provider/question.provider';

export default {
  name: 'UploadJsonSoal',
  data() {
    return {
      isDragging: false,
      selectedFile: null, // Store the whole file object
      selectedFileName: null,
    };
  },
  methods: {
    triggerFileInput() { this.$refs.fileInput.click(); },
    handleFileSelect(event) { this.processFile(event.target.files[0]); },
    handleDrop(event) { this.isDragging = false; this.processFile(event.dataTransfer.files[0]); },
    processFile(file) {
      if (file && file.type === 'application/json') {
        this.selectedFile = file; // Save the file object
        this.selectedFileName = file.name;
      } else {
        alert('Hanya file dengan format .json yang diperbolehkan!'); 
        this.selectedFile = null;
        this.selectedFileName = null;
      }
    },
    
    // ## SAVE FUNCTION IS UPDATED HERE ##
    async saveFile() {
      if (!this.selectedFile) {
        alert('Silakan pilih file terlebih dahulu.'); 
        return;
      }

      const formData = new FormData();
      formData.append('file', this.selectedFile);

      try {
        await createQuestionFromJson(formData);
        alert(`File ${this.selectedFileName} berhasil diunggah dan soal telah dibuat!`);
        this.$router.push('/dosen/soal/list');
      } catch (error) {
        console.error("Gagal mengunggah file JSON:", error);
        alert('Terjadi kesalahan saat menyimpan file. Silakan periksa format JSON Anda.');
      }
    }
  }
};
</script>