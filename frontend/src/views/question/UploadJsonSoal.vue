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

<script setup>
import { ref, computed } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { createQuestionFromJson } from '../../provider/question.provider';

import { usePopup } from '../../hooks/usePopup';

const route = useRoute();
const router = useRouter();


const { showSuccess, showError } = usePopup();

const isDragging = ref(false);
const selectedFile = ref(null);
const selectedFileName = ref(null);

const isAdminRoute = computed(() =>
  route.path.startsWith('/admin/soal')
);

const listRouteName = computed(() =>
  isAdminRoute.value ? 'AdminSoalList' : 'DosenSoalList'
);

const fileInput = ref(null);

const triggerFileInput = () => {
  fileInput.value.click();
}

const handleFileSelect = (event) => {
  processFile(event.target.files[0]);
}

const handleDrop = (event) => {
  isDragging.value = false;
  processFile(event.dataTransfer.files[0]);
}

const processFile = (file) =>  {
  if (file && file.type === 'application/json') {
    selectedFile.value = file;
    selectedFileName.value = file.name;
  } else {
    
    showError('Format Tidak Sesuai', 'Hanya file dengan format .json yang diperbolehkan!');
    selectedFile.value = null;
    selectedFileName.value = null;
  }
}

const saveFile = async () => {
  if (!selectedFile.value) {
    
    showError('File Belum Dipilih', 'Silakan pilih file JSON terlebih dahulu.');
    return;
  }

  const formData = new FormData();
  formData.append('file', selectedFile.value);

  try {
    await createQuestionFromJson(formData);
    
    
    await showSuccess('Berhasil', `File ${selectedFileName.value} berhasil diunggah dan soal telah dibuat!`);

  
    router.push({ 
      name: listRouteName.value,
      query: { show_last_page: 'true' }
    });

  } catch (error) {
    console.error("Gagal mengunggah file JSON:", error);
 
    showError(
      'Gagal Upload', 
      'Terjadi kesalahan saat menyimpan file.', 
      'Silakan periksa format JSON Anda atau koneksi internet.'
    );
  }
};
</script>