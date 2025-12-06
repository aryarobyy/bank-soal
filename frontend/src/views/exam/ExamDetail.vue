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
      <div class="animate-spin rounded-full h-10 w-10 border-b-2 border-blue-600 mx-auto mb-4"></div>
      Memuat data ujian...
    </div>

    <div v-else-if="error" class="text-center text-red-600 py-10 bg-red-50 rounded-lg">
      {{ error }}
    </div>

    <div v-else-if="exam" class="space-y-4">
      <div class="bg-gray-50 p-4 rounded-lg border">
        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <div class="space-y-3">
            <div>
              <h3 class="text-sm font-bold text-gray-500 uppercase">Nama Ujian</h3>
              <p class="text-gray-800 font-medium text-lg">{{ exam.title }}</p>
            </div>
            <div>
              <h3 class="text-sm font-bold text-gray-500 uppercase">Deskripsi</h3>
              <p class="text-gray-600">{{ exam.description || "-" }}</p>
            </div>
            </div>
          <div class="space-y-3">
             <div class="grid grid-cols-2 gap-4">
                <div>
                  <h3 class="text-sm font-bold text-gray-500 uppercase">Kesulitan</h3>
                  <p class="capitalize text-gray-800">{{ exam.difficulty }}</p>
                </div>
                <div>
                  <h3 class="text-sm font-bold text-gray-500 uppercase">Durasi</h3>
                  <p class="text-gray-800">{{ exam.long_time }} menit</p>
                </div>
             </div>
             <div class="grid grid-cols-2 gap-4">
                <div>
                  <h3 class="text-sm font-bold text-gray-500 uppercase">Mulai</h3>
                  <p class="text-gray-600 text-sm">{{ formatDate(exam.started_at) }}</p>
                </div>
                <div>
                  <h3 class="text-sm font-bold text-gray-500 uppercase">Selesai</h3>
                  <p class="text-gray-600 text-sm">{{ formatDate(exam.finished_at) }}</p>
                </div>
             </div>
          </div>
        </div>
      </div>

      <div class="flex gap-3 mt-4">
        <router-link
          :to="{ name: isAdminRoute ? 'AdminExamEdit' : 'DosenExamEdit', params: { id: exam.id } }"
          class="bg-blue-600 hover:bg-blue-700 text-white px-4 py-2 rounded-md font-medium text-sm flex items-center gap-2"
        >
          ✏️ Edit Ujian
        </router-link>
        <button
          @click="removeExam(exam.id)"
          class="bg-red-500 hover:bg-red-600 text-white px-4 py-2 rounded-md font-medium text-sm flex items-center gap-2"
        >
          🗑️ Hapus Ujian
        </button>
      </div>

      <div class="mt-10">
        <div class="flex justify-between items-center mb-4">
          <h3 class="text-xl font-bold text-gray-800">Daftar Soal (Total: {{ totalQuestions }})</h3>
          <button
            @click="openAddSoalModal"
            class="bg-green-600 hover:bg-green-700 text-white px-4 py-2 rounded-md shadow hover:shadow-lg transition font-medium flex items-center gap-2"
          >
            ➕ Tambah Soal
          </button>
        </div>

        <div class="border rounded-lg overflow-hidden shadow-sm">
          <table class="w-full text-left">
            <thead class="bg-gray-100 border-b text-gray-700 text-sm uppercase tracking-wider">
              <tr>
                <th class="p-4 w-16 text-center">No</th>
                <th class="p-4">Pertanyaan</th>
                <th class="p-4 w-24 text-center">Aksi</th>
              </tr>
            </thead>
            
            <tbody v-if="tableLoading">
               <tr>
                 <td colspan="3" class="p-8 text-center text-gray-500">
                   <div class="animate-spin rounded-full h-6 w-6 border-b-2 border-blue-600 mx-auto mb-2"></div>
                   Memuat soal halaman {{ mainPage }}...
                 </td>
               </tr>
            </tbody>

            <tbody v-else class="divide-y divide-gray-200">
              <tr v-for="(q, index) in examQuestions" :key="q.id" class="hover:bg-gray-50">
                <td class="p-4 text-center text-gray-500">
                  {{ (mainPage - 1) * mainLimit + index + 1 }}
                </td>
                <td class="p-4">
                  <p class="text-gray-800 line-clamp-2">{{ q.question_text || "[Konten Soal Tidak Valid]" }}</p>
                </td>
                <td class="p-4 text-center">
                  <button
                    @click="handleDeleteQuestion(q)"
                    class="text-red-500 hover:text-red-700 p-2 hover:bg-red-50 rounded transition"
                    title="Hapus dari ujian"
                  >
                    🗑️
                  </button>
                </td>
              </tr>
              <tr v-if="examQuestions.length === 0">
                <td colspan="3" class="p-8 text-center text-gray-500">
                  <p class="text-lg mb-2">Belum ada soal.</p>
                  <p class="text-sm">Klik tombol "Tambah Soal" di atas untuk memasukkan soal.</p>
                </td>
              </tr>
            </tbody>
          </table>
        </div>

        <div v-if="totalQuestions > mainLimit" class="flex justify-between items-center mt-4 px-2">
          <div class="text-sm text-gray-600">
            Halaman <span class="font-bold">{{ mainPage }}</span> dari <span class="font-bold">{{ mainTotalPages }}</span>
          </div>
          <div class="flex gap-1">
            <button 
              @click="mainPage--" 
              :disabled="mainPage === 1"
              class="px-3 py-1 border rounded bg-white hover:bg-gray-50 disabled:opacity-50 text-sm"
            >
              Prev
            </button>
            
            <button 
              v-for="p in visibleMainPages" 
              :key="p"
              @click="mainPage = p"
              :class="['px-3 py-1 border rounded text-sm', mainPage === p ? 'bg-blue-600 text-white border-blue-600' : 'bg-white hover:bg-gray-50']"
            >
              {{ p }}
            </button>
            
            <button 
              @click="mainPage++" 
              :disabled="mainPage === mainTotalPages"
              class="px-3 py-1 border rounded bg-white hover:bg-gray-50 disabled:opacity-50 text-sm"
            >
              Next
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>

  <div v-if="showAddSoalModal" class="fixed inset-0 bg-black/60 backdrop-blur-sm flex justify-center items-center p-4 z-50">
    <div class="bg-white rounded-xl shadow-2xl w-full max-w-4xl max-h-[90vh] flex flex-col">
      <div class="p-5 border-b flex justify-between items-center">
         <h3 class="text-xl font-bold text-gray-800">Bank Soal</h3>
         <button @click="closeAddSoalModal" class="text-gray-400 hover:text-gray-600">✖</button>
      </div>
      
      <div class="p-5 bg-gray-50 border-b space-y-3">
         <label class="block text-sm font-bold text-gray-700">Pilih Mata Kuliah (Subjek)</label>
         <div class="flex flex-col sm:flex-row gap-3">
            <select 
              v-model="selectedSubject" 
              class="flex-1 p-2.5 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 outline-none"
            >
              <option :value="null" disabled>-- Pilih Subjek --</option>
              <option v-for="subject in availableSubjects" :key="subject.id" :value="subject.id">
                {{ subject.title }}
              </option>
            </select>
            
            <button 
              v-if="selectedSubject"
              @click="selectAllBySubject"
              :disabled="loadingAllSubject"
              class="w-full sm:w-auto px-6 py-2.5 bg-indigo-600 hover:bg-indigo-700 text-white font-medium rounded-lg shadow-sm transition disabled:opacity-70 flex items-center justify-center gap-2 whitespace-nowrap"
            >
              <span v-if="loadingAllSubject" class="animate-spin h-4 w-4 border-2 border-white border-t-transparent rounded-full"></span>
              {{ loadingAllSubject ? 'Mengambil Data...' : 'Ambil SEMUA Soal Subjek Ini' }}
            </button>
         </div>
      </div>

      <div class="flex-1 overflow-hidden flex flex-col relative">
         <div v-if="modalLoading" class="absolute inset-0 bg-white/80 z-10 flex items-center justify-center">
            <div class="flex flex-col items-center">
               <div class="animate-spin rounded-full h-10 w-10 border-b-2 border-blue-600 mb-2"></div>
               <span class="text-gray-600 font-medium">Memuat Soal...</span>
            </div>
         </div>

         <div class="flex-1 overflow-y-auto">
            <table class="w-full text-sm text-left">
              <thead class="bg-gray-100 text-gray-700 sticky top-0 z-0 shadow-sm">
                <tr>
                  <th class="p-3 w-12 text-center">
                    <input 
                      type="checkbox" 
                      :checked="isAllPageSelected"
                      @change="toggleSelectAllPage"
                      class="w-4 h-4 rounded cursor-pointer accent-blue-600"
                    />
                  </th>
                  <th class="p-3">Pertanyaan</th>
                  <th class="p-3 w-24 text-center">Status</th>
                </tr>
              </thead>
              <tbody class="divide-y">
                <tr v-for="q in questionsForSubject" :key="q.id" class="hover:bg-blue-50 transition cursor-pointer" @click="toggleSelection(q.id)">
                  <td class="p-3 text-center align-top">
                    <input 
                      type="checkbox" 
                      :value="q.id" 
                      v-model="selectedQuestions" 
                      class="w-4 h-4 rounded cursor-pointer accent-blue-600"
                      :disabled="isQuestionAlreadyAdded(q.id)"
                      @click.stop
                    />
                  </td>
                  <td class="p-3">
                    <p class="font-medium text-gray-800">{{ q.question_text }}</p>
                  </td>
                  <td class="p-3 text-center">
                    <span v-if="isQuestionAlreadyAdded(q.id)" class="px-2 py-1 bg-green-100 text-green-700 text-xs font-bold rounded">
                      Terdaftar
                    </span>
                    <span v-else-if="selectedQuestions.includes(q.id)" class="px-2 py-1 bg-blue-100 text-blue-700 text-xs font-bold rounded">
                      Dipilih
                    </span>
                    <span v-else class="text-gray-400">-</span>
                  </td>
                </tr>
                <tr v-if="questionsForSubject.length === 0 && !modalLoading">
                  <td colspan="3" class="p-8 text-center text-gray-500">
                    {{ selectedSubject ? 'Tidak ada soal di halaman ini.' : 'Silakan pilih subjek terlebih dahulu.' }}
                  </td>
                </tr>
              </tbody>
            </table>
         </div>

         <div v-if="selectedSubject && modalTotalItems > 0" class="p-3 border-t bg-gray-50 flex justify-between items-center">
            <span class="text-sm text-gray-600">
              Menampilkan {{ (modalPage - 1) * modalLimit + 1 }} - {{ Math.min(modalPage * modalLimit, modalTotalItems) }} dari <b>{{ modalTotalItems }}</b> soal
            </span>
            <div class="flex gap-2">
               <button 
                 @click="prevModalPage" 
                 :disabled="modalPage === 1"
                 class="px-3 py-1.5 bg-white border border-gray-300 rounded hover:bg-gray-100 disabled:opacity-50 text-sm font-medium transition"
               >
                 Previous
               </button>
               <span class="px-3 py-1.5 bg-white border border-gray-300 rounded text-sm font-bold text-blue-600">
                 {{ modalPage }}
               </span>
               <button 
                 @click="nextModalPage" 
                 :disabled="modalPage >= modalTotalPages"
                 class="px-3 py-1.5 bg-white border border-gray-300 rounded hover:bg-gray-100 disabled:opacity-50 text-sm font-medium transition"
               >
                 Next
               </button>
            </div>
         </div>
      </div>

      <div class="p-5 border-t flex justify-between items-center bg-gray-50 rounded-b-xl">
        <div class="text-sm">
           <span class="font-bold text-blue-600 text-xl">{{ selectedQuestions.length }}</span>
           <span class="text-gray-600 ml-1">soal akan ditambahkan</span>
        </div>
        <div class="flex gap-3">
          <button @click="closeAddSoalModal" class="px-5 py-2.5 bg-white border border-gray-300 text-gray-700 font-medium rounded-lg hover:bg-gray-50 transition">
            Batal
          </button>
          <button 
            @click="handleAddSoal" 
            :disabled="saveLoading"
            class="px-6 py-2.5 bg-blue-600 text-white font-bold rounded-lg hover:bg-blue-700 disabled:opacity-50 disabled:cursor-not-allowed shadow-md hover:shadow-lg transition flex items-center gap-2"
          >
            <span v-if="saveLoading" class="animate-spin h-4 w-4 border-2 border-white border-t-transparent rounded-full"></span>
            {{ saveLoading ? savingText : 'Simpan Pilihan' }}
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, computed, watch } from "vue";
import { useRoute, useRouter } from "vue-router";
import { deleteExam, addQuestions, removeQuestions, getExamById } from "../../provider/exam.provider";
import { getPaginatedSubjects } from "../../provider/subject.provider";
import { getQuestionsBySubject, getQuestionsByExam } from "../../provider/question.provider";

import { usePopup } from "../../hooks/usePopup";

const { showSuccess, showError, showConfirm } = usePopup();

const route = useRoute();
const router = useRouter();
const isAdminRoute = computed(() => route.path.startsWith('/admin'));

const exam = ref(null);
const examQuestions = ref([]);
const totalQuestions = ref(0);
const loading = ref(true);
const tableLoading = ref(false); 
const error = ref("");

const mainPage = ref(1);
const mainLimit = 10;

const mainTotalPages = computed(() => Math.ceil(totalQuestions.value / mainLimit));

const visibleMainPages = computed(() => {
  let pages = [];
  for (let i = 1; i <= mainTotalPages.value; i++) {
     if (i === 1 || i === mainTotalPages.value || (i >= mainPage.value - 1 && i <= mainPage.value + 1)) {
        pages.push(i);
     }
  }
  return [...new Set(pages)].sort((a,b)=>a-b);
});

const loadExamDetails = async () => {
  try {
    const id = route.params.id;
    const res = await getExamById(id);
    exam.value = res;
  } catch (err) {
    error.value = "Gagal memuat data ujian.";
  }
};

const loadExamQuestions = async () => {
  tableLoading.value = true;
  try {
    const id = route.params.id;
    const offset = (mainPage.value - 1) * mainLimit;
    
    const result = await getQuestionsByExam(id, mainLimit, offset);
    
    examQuestions.value = result.data || [];
    totalQuestions.value = result.total || 0;
    
  } catch (err) {
    console.error("Gagal soal ujian:", err);
    examQuestions.value = [];
  } finally {
    tableLoading.value = false;
  }
};

watch(mainPage, () => {
    loadExamQuestions();
});

const removeExam = async (id) => {
  const isConfirmed = await showConfirm('Konfirmasi Hapus', "Hapus ujian ini permanen?");
  if (!isConfirmed) return;

  try {
    await deleteExam(id);
    await showSuccess('Berhasil', 'Ujian berhasil dihapus');
    router.push({ name: isAdminRoute.value ? 'AdminManageExam' : 'DosenManageExam' });
  } catch (err) {
    showError('Gagal', "Gagal menghapus.");
  }
};

const handleDeleteQuestion = async (question) => {
  const isConfirmed = await showConfirm('Konfirmasi Hapus', "Hapus soal ini dari ujian?");
  if (!isConfirmed) return;
  
  try {
    await removeQuestions(exam.value.id, { "question_ids": [question.id] }); 
    await showSuccess('Berhasil', "✅ Soal berhasil dihapus dari ujian!");
    
    if (examQuestions.value.length === 1 && mainPage.value > 1) {
        mainPage.value--;
    } else {
        await loadExamQuestions();
    }
  } catch (err) {
    showError('Gagal', "Gagal menghapus soal.");
  }
};

const showAddSoalModal = ref(false);
const modalLoading = ref(false);
const saveLoading = ref(false);
const savingText = ref("Menyimpan...");
const loadingAllSubject = ref(false);
const availableSubjects = ref([]);
const selectedSubject = ref(null);
const questionsForSubject = ref([]);
const selectedQuestions = ref([]);
const modalPage = ref(1);
const modalLimit = 10;
const modalTotalItems = ref(0);

const modalTotalPages = computed(() => {
  return modalTotalItems.value > 0 ? Math.ceil(modalTotalItems.value / modalLimit) : 1;
});

const openAddSoalModal = () => {
  selectedSubject.value = null;
  questionsForSubject.value = [];
  selectedQuestions.value = [];
  modalPage.value = 1;
  modalTotalItems.value = 0;
  fetchAvailableSubjects();
  showAddSoalModal.value = true;
};

const closeAddSoalModal = () => {
  showAddSoalModal.value = false;
};

const fetchAvailableSubjects = async () => {
  try {
    const res = await getPaginatedSubjects(100, 0, "");
    availableSubjects.value = res.data.data || [];
  } catch (err) { /* ignore */ }
};

const fetchQuestionsForSubject = async (subjectId) => {
  if (!subjectId) return;
  modalLoading.value = true;
  try {
    const offset = (modalPage.value - 1) * modalLimit;
    const result = await getQuestionsBySubject(subjectId, modalLimit, offset);
    if (result.data && Array.isArray(result.data)) {
       questionsForSubject.value = result.data;
       modalTotalItems.value = result.total || 9999;
    } else {
       questionsForSubject.value = [];
       modalTotalItems.value = 0;
    }
  } catch (err) {
    questionsForSubject.value = [];
  } finally {
    modalLoading.value = false;
  }
};

const toggleSelection = (qId) => {
    if (isQuestionAlreadyAdded(qId)) return;
    const idx = selectedQuestions.value.indexOf(qId);
    if (idx > -1) selectedQuestions.value.splice(idx, 1);
    else selectedQuestions.value.push(qId);
};

const isAllPageSelected = computed(() => {
  if (questionsForSubject.value.length === 0) return false;
  return questionsForSubject.value.every(q => 
    selectedQuestions.value.includes(q.id) || isQuestionAlreadyAdded(q.id)
  );
});

const toggleSelectAllPage = () => {
  if (isAllPageSelected.value) {
    questionsForSubject.value.forEach(q => {
      if (!isQuestionAlreadyAdded(q.id)) {
        const index = selectedQuestions.value.indexOf(q.id);
        if (index > -1) selectedQuestions.value.splice(index, 1);
      }
    });
  } else {
    questionsForSubject.value.forEach(q => {
      if (!isQuestionAlreadyAdded(q.id) && !selectedQuestions.value.includes(q.id)) {
        selectedQuestions.value.push(q.id);
      }
    });
  }
};

const selectAllBySubject = async () => {
  if (!selectedSubject.value) return;
  loadingAllSubject.value = true;
  try {
    let allFetchedIds = [];
    let offset = 0;
    const BATCH_LIMIT = 20; 
    let hasMoreData = true;

    while (hasMoreData) {
       const result = await getQuestionsBySubject(selectedSubject.value, BATCH_LIMIT, offset);
       let chunk = [];
       if (Array.isArray(result)) chunk = result;
       else if (result.data && Array.isArray(result.data)) chunk = result.data;
       else if (result.data && result.data.data && Array.isArray(result.data.data)) chunk = result.data.data;

       if (chunk.length > 0) {
          chunk.forEach(q => allFetchedIds.push(q.id));
          offset += BATCH_LIMIT;
       } 
       if (chunk.length < BATCH_LIMIT) hasMoreData = false;
    }

    if (allFetchedIds.length === 0) {
      showError('Kosong', "Tidak ada soal ditemukan di subjek ini.");
      loadingAllSubject.value = false;
      return;
    }

    let addedCount = 0;
    allFetchedIds.forEach(id => {
      const qId = Number(id);
      const alreadySelected = selectedQuestions.value.includes(qId);
      const existsInLoaded = examQuestions.value.some(eq => eq.id === qId); 

      if (!alreadySelected && !existsInLoaded) {
        selectedQuestions.value.push(qId);
        addedCount++;
      }
    });

    if (addedCount > 0) {
       showSuccess('Berhasil', `Berhasil memilih ${addedCount} soal baru!`);
    } else {
       showSuccess('Info', "Semua soal sudah terpilih.");
    }
  } catch (err) {
    showError('Gagal', "Gagal mengambil semua soal.");
  } finally {
    loadingAllSubject.value = false;
  }
};

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
  if (selectedQuestions.value.length === 0) {
      showError('Validasi', "Pilih soal terlebih dahulu.");
      return;
  }
  
  saveLoading.value = true;
  savingText.value = "Menyiapkan...";

  try {
    
    const BATCH_SIZE = 50; 
    const total = selectedQuestions.value.length;
    let processed = 0;

    for (let i = 0; i < total; i += BATCH_SIZE) {
      const chunk = selectedQuestions.value.slice(i, i + BATCH_SIZE);
      savingText.value = `Menyimpan ${Math.min(i + BATCH_SIZE, total)} / ${total}`;
      
      await addQuestions(exam.value.id, { "question_ids": chunk });
      processed += chunk.length;
    }
    
    await showSuccess('Berhasil', `${processed} soal berhasil ditambahkan!`);
    
    closeAddSoalModal();
    mainPage.value = 1;
    await loadExamQuestions();
  } catch (err) {
    console.error(err);
    showError('Gagal', "Gagal menyimpan sebagian soal. Silakan cek koneksi dan coba lagi.");
  } finally {
    saveLoading.value = false;
    savingText.value = "Simpan Pilihan";
  }
};

const isQuestionAlreadyAdded = (qId) => {
  return examQuestions.value.some(eq => eq.id === qId);
};

watch(selectedSubject, (newId) => {
  if (newId) {
    modalPage.value = 1;
    fetchQuestionsForSubject(newId);
  } else {
    questionsForSubject.value = [];
  }
});

const formatDate = (d) => d ? new Date(d).toLocaleString("id-ID") : "-";

onMounted(async () => {
  loading.value = true;
  await loadExamDetails();
  await loadExamQuestions();
  loading.value = false;
});
</script>