<template>
  <div class="p-6 bg-gray-50 min-h-screen">
    <div class="flex justify-between items-center mb-6">
      <h2 class="text-2xl font-bold text-gray-800">Manajemen File Excel</h2>
    </div>

    <div v-if="loading" class="text-center py-10">
      <p class="text-gray-500">Memuat data file...</p>
    </div>
    <div v-else-if="error" class="text-center py-10 bg-red-50 p-4 rounded-lg">
      <p class="text-red-600">{{ error }}</p>
    </div>

    <div v-else class="bg-white shadow rounded-lg overflow-hidden">
      <table class="min-w-full border-collapse">
        <thead class="bg-gray-100 text-gray-700 text-sm">
          <tr>
            <th class="px-4 py-3 text-left">No</th>
            <th class="px-4 py-3 text-left">Nama File</th>
            <th class="px-4 py-3 text-left">Tanggal Dibuat</th>
            <th class="px-4 py-3 text-left">Aksi</th>
          </tr>
        </thead>
        <tbody class="text-gray-800 text-sm">
          <tr
            v-for="(file, index) in fileList"
            :key="file.id"
            class="border-t hover:bg-gray-50 transition"
          >
            <td class="px-4 py-3">{{ (currentPage - 1) * itemsPerPage + index + 1 }}</td>
            <td class="px-4 py-3 font-medium">{{ getFileName(file.file_path) }}</td>
            <td class="px-4 py-3">{{ new Date(file.CreatedAt).toLocaleString("id-ID") }}</td>
            <td class="px-4 py-3">
              <button
                @click="handleDownload(file)"
                :disabled="isDownloading[file.id]"
                class="px-3 py-1 bg-blue-500 text-white rounded-md hover:bg-blue-600 transition mr-2 disabled:bg-gray-400"
              >
                {{ isDownloading[file.id] ? 'Loading...' : 'Download' }}
              </button>
              
              <button
                @click="handleDelete(file)"
                class="px-3 py-1 bg-red-500 text-white rounded-md hover:bg-red-600 transition"
              >
                Hapus
              </button>
            </td>
          </tr>
          <tr v-if="fileList.length === 0">
            <td colspan="4" class="px-4 py-4 text-center text-gray-500">
              Belum ada data file Excel.
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <div v-if="!loading && totalPages > 1" class="flex justify-between items-center mt-6">
      <span class="text-sm text-gray-700">
        Halaman <span class="font-semibold">{{ currentPage }}</span> dari <span class="font-semibold">{{ totalPages }}</span> (Total <span class="font-semibold">{{ totalItems }}</span> file)
      </span>
      <div class="flex gap-1">
        <button
          @click="prevPage"
          :disabled="currentPage === 1"
          class="px-3 py-1 bg-white border border-gray-300 rounded-md text-sm hover:bg-gray-50 disabled:opacity-50 disabled:cursor-not-allowed"
        >
          &lt; Sebelumnya
        </button>
        <button
          @click="nextPage"
          :disabled="currentPage === totalPages"
          class="px-3 py-1 bg-white border border-gray-300 rounded-md text-sm hover:bg-gray-50 disabled:opacity-50 disabled:cursor-not-allowed"
        >
          Berikutnya &gt;
        </button>
      </div>
    </div>

  </div>
</template>

<script setup>
import { ref, onMounted, computed, watch } from "vue";
// BARU: Impor fungsi download
import { getAllXlsPaths, deleteXlsPath, downloadXlsFile } from "../../provider/xlspath.provider.js"; 

const fileList = ref([]);
const loading = ref(true);
const error = ref(null);

// BARU: State untuk melacak download
const isDownloading = ref({});

const currentPage = ref(1);
const itemsPerPage = ref(10);
const totalItems = ref(0);

const totalPages = computed(() => {
  return Math.ceil(totalItems.value / itemsPerPage.value);
});

const fetchFiles = async () => {
  try {
    loading.value = true;
    const offset = (currentPage.value - 1) * itemsPerPage.value;
    const responsePayload = await getAllXlsPaths(itemsPerPage.value, offset);
    
    if (Array.isArray(responsePayload)) {
      fileList.value = responsePayload;
      totalItems.value = responsePayload.length; 
    } else if (responsePayload && typeof responsePayload === 'object' && responsePayload.data) {
      fileList.value = responsePayload.data || [];
      totalItems.value = responsePayload.total || 0;
    } else {
      fileList.value = [];
      totalItems.value = 0;
    }
    error.value = null;
  } catch (err) {
    console.error("Gagal mengambil data file:", err);
    error.value = "Tidak dapat memuat data. Silakan coba lagi nanti.";
    fileList.value = [];
    totalItems.value = 0;
  } finally {
    loading.value = false;
  }
};

onMounted(() => {
  fetchFiles();
});

watch(currentPage, (newPage, oldPage) => {
  if (newPage !== oldPage) {
    fetchFiles();
  }
});

const prevPage = () => {
  if (currentPage.value > 1) {
    currentPage.value--;
  }
};

const nextPage = () => {
  if (currentPage.value < totalPages.value) {
    currentPage.value++;
  }
};

const getFileName = (path) => {
  if (!path) return 'N/A';
  const parts = path.split('/');
  return parts[parts.length - 1];
};

// BARU: Fungsi untuk menangani download
const handleDownload = async (file) => {
  const fileId = file.id;
  isDownloading.value[fileId] = true;

  try {
    // 1. Panggil provider (mengharapkan blob)
    const response = await downloadXlsFile(fileId);

    // 2. Buat Blob dari data
    const blob = new Blob([response.data], { 
      type: response.headers['content-type'] 
    });

    // 3. Tentukan nama file
    let filename = getFileName(file.file_path); // Nama fallback
    const contentDisposition = response.headers['content-disposition'];
    if (contentDisposition) {
      const filenameMatch = contentDisposition.match(/filename="?(.+)"?/);
      if (filenameMatch && filenameMatch.length > 1) {
        filename = filenameMatch[1];
      }
    }

    // 4. Buat link sementara untuk memicu download
    const url = window.URL.createObjectURL(blob);
    const link = document.createElement('a');
    link.href = url;
    link.setAttribute('download', filename);
    document.body.appendChild(link);
    link.click();

    // 5. Bersihkan
    window.URL.revokeObjectURL(url);
    document.body.removeChild(link);

  } catch (error) {
    console.error("Gagal download file:", error);
    // Jika backend mengirim error (JSON) padahal kita minta blob
    if (error.response && error.response.data instanceof Blob) {
      try {
        const errorText = await error.response.data.text();
        const errorJson = JSON.parse(errorText);
        alert(`Gagal: ${errorJson.message || 'Error tidak diketahui'}`);
      } catch (e) {
        alert('Gagal memproses file error dari server.');
      }
    } else {
      alert(`Gagal download: ${error.response?.data?.message || 'Terjadi kesalahan.'}`);
    }
  } finally {
    isDownloading.value[fileId] = false;
  }
};

const handleDelete = async (file) => {
  const fileName = getFileName(file.file_path);
  
  if (confirm(`Yakin ingin menghapus file "${fileName}"?`)) {
    try {
      await deleteXlsPath(file.id);
      alert("File berhasil dihapus.");
      
      if (fileList.value.length === 1 && currentPage.value > 1) {
        currentPage.value--; 
      } else {
        fetchFiles(); 
      }

    } catch (err) {
      console.error("Gagal menghapus file:", err);
      alert(`Gagal menghapus: ${err.response?.data?.message || 'Error tidak diketahui'}`);
    }
  }
};
</script>