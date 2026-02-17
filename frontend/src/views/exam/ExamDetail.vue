<template>
  <div class="bg-white rounded-lg shadow-md p-6">
    <div class="flex justify-between items-center mb-6">
      <h2 class="text-2xl font-bold text-gray-800">Detail Ujian</h2>
      <Button
        :to="{ name: isAdminRoute ? 'AdminManageExam' : 'DosenManageExam' }"
        variant="secondary"
        icon="fas fa-arrow-left"
      >
        Kembali
      </Button>
    </div>

    <div v-if="loading" class="space-y-6 py-4">
      <div class="bg-gray-50 p-4 rounded-lg border">
        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <div class="space-y-4">
            <div class="space-y-2">
              <div class="h-3 bg-gray-200 rounded w-20 animate-pulse"></div>
              <div class="h-5 bg-gray-200 rounded w-48 animate-pulse"></div>
            </div>
            <div class="space-y-2">
              <div class="h-3 bg-gray-200 rounded w-16 animate-pulse"></div>
              <div class="h-4 bg-gray-200 rounded w-full animate-pulse"></div>
            </div>
          </div>
          <div class="space-y-4">
            <div class="grid grid-cols-2 gap-4">
              <div class="space-y-2">
                <div class="h-3 bg-gray-200 rounded w-16 animate-pulse"></div>
                <div class="h-4 bg-gray-200 rounded w-20 animate-pulse"></div>
              </div>
              <div class="space-y-2">
                <div class="h-3 bg-gray-200 rounded w-12 animate-pulse"></div>
                <div class="h-4 bg-gray-200 rounded w-24 animate-pulse"></div>
              </div>
            </div>
          </div>
        </div>
      </div>
      <div class="space-y-3">
        <div class="h-5 bg-gray-200 rounded w-40 animate-pulse"></div>
        <div v-for="n in 3" :key="n" class="flex items-center gap-4 py-3 border-b border-gray-100">
          <div class="h-4 bg-gray-200 rounded w-8 animate-pulse"></div>
          <div class="flex-1 h-4 bg-gray-200 rounded animate-pulse"></div>
          <div class="h-4 bg-gray-200 rounded w-8 animate-pulse"></div>
        </div>
      </div>
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
        <Button
          :to="{ name: isAdminRoute ? 'AdminExamEdit' : 'DosenExamEdit', params: { id: exam.id } }"
          variant="primary"
          icon="fas fa-pencil-alt"
        >
          Edit Ujian
        </Button>
        <Button
          @click="removeExam(exam.id)"
          variant="danger"
          icon="fas fa-trash"
        >
          Hapus Ujian
        </Button>
      </div>

      <div class="mt-10">
        <div class="flex justify-between items-center mb-4">
          <h3 class="text-xl font-bold text-gray-800">Daftar Soal (Total: {{ totalQuestions }})</h3>
          <Button
            @click="openAddSoalModal"
            variant="success"
            icon="fas fa-plus"
          >
            Tambah Soal
          </Button>
        </div>

        <Table
          :headers="questionHeaders"
          :items="examQuestions"
          :loading="tableLoading"
          empty-text="Belum ada soal. Klik tombol 'Tambah Soal' di atas untuk memasukkan soal."
        >
          <template #loading>
            <div class="space-y-3 py-2">
              <div v-for="n in 5" :key="n" class="flex items-center gap-4 py-2 border-b border-gray-100">
                <div class="h-4 bg-gray-200 rounded w-8 animate-pulse"></div>
                <div class="flex-1 h-4 bg-gray-200 rounded animate-pulse"></div>
                <div class="h-4 bg-gray-200 rounded w-8 animate-pulse"></div>
              </div>
            </div>
          </template>

          <template #cell-no="{ index }">
            {{ (mainPage - 1) * mainLimit + index + 1 }}
          </template>

          <template #cell-question="{ item }">
            <p class="text-gray-800 line-clamp-2">{{ item.question_text || "[Konten Soal Tidak Valid]" }}</p>
          </template>

          <template #cell-actions="{ item }">
            <Button
              @click="handleDeleteQuestion(item)"
              variant="ghost"
              size="sm"
              class="text-red-500 hover:text-red-700 hover:bg-red-50"
              title="Hapus dari ujian"
            >
              🗑️
            </Button>
          </template>

          <template #footer>
            <div v-if="totalQuestions > mainLimit" class="flex justify-between items-center px-2">
              <div class="text-sm text-gray-600">
                Halaman <span class="font-bold">{{ mainPage }}</span> dari <span class="font-bold">{{ mainTotalPages }}</span>
              </div>
              <div class="flex gap-1">
                <Button 
                  @click="mainPage--" 
                  :disabled="mainPage === 1"
                  variant="outline"
                  size="sm"
                >
                  Prev
                </Button>
                
                <Button 
                  v-for="p in visibleMainPages" 
                  :key="p"
                  @click="mainPage = p"
                  size="sm"
                  :variant="mainPage === p ? 'primary' : 'outline'"
                >
                  {{ p }}
                </Button>
                
                <Button 
                  @click="mainPage++" 
                  :disabled="mainPage === mainTotalPages"
                  variant="outline"
                  size="sm"
                >
                  Next
                </Button>
              </div>
            </div>
          </template>
        </Table>
      </div>
    </div>
  </div>

  <div v-if="showAddSoalModal" class="fixed inset-0 bg-black/60 backdrop-blur-sm flex justify-center items-center p-4 z-50">
    <div class="bg-white rounded-xl shadow-2xl w-full max-w-4xl max-h-[90vh] flex flex-col">
      <div class="p-5 border-b flex justify-between items-center">
         <h3 class="text-xl font-bold text-gray-800">Bank Soal</h3>
         <button @click="closeAddSoalModal" class="text-gray-400 hover:text-gray-600">✖</button>
      </div>
      
      <div class="p-5 bg-gray-50 border-b space-y-4">
         
         <div>
            <label class="block text-sm font-bold text-gray-700 mb-2">Pilih Mata Kuliah (Subjek)</label>
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
                
                <Button 
                  v-if="selectedSubject"
                  @click="selectAllBySubject"
                  :loading="loadingAllSubject"
                  variant="secondary"
                >
                  Ambil SEMUA Soal
                </Button>
            </div>
         </div>

         <div v-if="selectedSubject" class="p-4 bg-blue-50 rounded-lg border border-blue-100 flex flex-col sm:flex-row items-center justify-between gap-4">
            <div class="flex items-center gap-3 w-full sm:w-auto">
                <div class="bg-blue-600 text-white p-2 rounded-full">
                    <i class="fas fa-dice"></i>
                </div>
                <div>
                    <p class="text-sm font-bold text-blue-900">Ambil Soal Acak</p>
                    <p class="text-xs text-blue-600">Pilih soal secara random dari subjek ini</p>
                </div>
            </div>

            <div class="flex items-center gap-2 w-full sm:w-auto">
                <input 
                    v-model.number="randomCount" 
                    type="number" 
                    min="1" 
                    class="w-20 p-2 border border-blue-300 rounded-lg text-center focus:ring-2 focus:ring-blue-500 outline-none"
                    placeholder="Jml"
                >
                <Button 
                    @click="handleGetRandom"
                    :loading="loadingRandom"
                    variant="primary"
                >
                    Generate
                </Button>
            </div>
         </div>

      </div>

      <div class="flex-1 overflow-hidden flex flex-col relative">
         <div v-if="modalLoading" class="absolute inset-0 bg-white/80 z-10 flex items-center justify-center">
            <div class="w-full max-w-lg space-y-3 px-8">
              <div v-for="n in 4" :key="n" class="flex items-center gap-3 py-2">
                <div class="w-4 h-4 bg-gray-300 rounded animate-pulse"></div>
                <div class="flex-1 h-4 bg-gray-300 rounded animate-pulse"></div>
                <div class="h-4 bg-gray-300 rounded w-16 animate-pulse"></div>
              </div>
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
               <Button 
                 @click="prevModalPage" 
                 :disabled="modalPage === 1"
                 variant="outline"
                 size="sm"
               >
                 Previous
               </Button>
               <span class="px-3 py-1.5 bg-white border border-gray-300 rounded text-sm font-bold text-blue-600">
                 {{ modalPage }}
               </span>
               <Button 
                 @click="nextModalPage" 
                 :disabled="modalPage >= modalTotalPages"
                 variant="outline"
                 size="sm"
               >
                 Next
               </Button>
            </div>
         </div>
      </div>

      <div class="p-5 border-t flex justify-between items-center bg-gray-50 rounded-b-xl">
        <div class="text-sm">
           <span class="font-bold text-blue-600 text-xl">{{ selectedQuestions.length }}</span>
           <span class="text-gray-600 ml-1">soal akan ditambahkan</span>
        </div>
        <div class="flex gap-3">
          <Button @click="closeAddSoalModal" variant="outline">
            Batal
          </Button>
          <Button 
            @click="handleAddSoal" 
            :loading="saveLoading"
            variant="primary"
            class="shadow-md"
          >
            {{ saveLoading ? savingText : 'Simpan Pilihan' }}
          </Button>
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
import { getQuestionsBySubject, getQuestionsByExam, getRandomQuestions,getQuestionsByCreatorAndSubject } from "../../provider/question.provider";
import { usePopup } from "../../hooks/usePopup";
import { useGetCurrentUser } from "../../hooks/useGetCurrentUser";
import Button from "../../components/ui/button/Button.vue";
import Table from "../../components/ui/table/Table.vue";

const { showSuccess, showError, showConfirm } = usePopup();
const { user } = useGetCurrentUser();
const route = useRoute();
const router = useRouter();
const isAdminRoute = computed(() => route.path.startsWith('/admin'));

const exam = ref(null);
const examQuestions = ref([]);
const totalQuestions = ref(0);
const loading = ref(true);
const tableLoading = ref(false); 
const error = ref("");

const questionHeaders = [
  { key: 'no', label: 'No', class: 'w-16 text-center', tdClass: 'text-center text-gray-500' },
  { key: 'question', label: 'Pertanyaan' },
  { key: 'actions', label: 'Aksi', class: 'w-24 text-center', tdClass: 'text-center' },
];

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


const randomCount = ref(10);
const loadingRandom = ref(false);

const modalTotalPages = computed(() => {
  return modalTotalItems.value > 0 ? Math.ceil(modalTotalItems.value / modalLimit) : 1;
});

const openAddSoalModal = () => {
  selectedSubject.value = null;
  questionsForSubject.value = [];
  selectedQuestions.value = [];
  randomCount.value = 10;
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
    let result;


    if (isAdminRoute.value) {

       result = await getQuestionsBySubject(subjectId, modalLimit, offset);
    } else {

       result = await getQuestionsByCreatorAndSubject(user.value.id, subjectId, modalLimit, offset);
    }


    if (result && (result.data || result.items)) {
       const dataArr = result.data || result.items || [];
       questionsForSubject.value = Array.isArray(dataArr) ? dataArr : [];
       modalTotalItems.value = result.total || questionsForSubject.value.length;
    } else if (Array.isArray(result)) {
       questionsForSubject.value = result;
       modalTotalItems.value = result.length;
    } else {
       questionsForSubject.value = [];
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


const handleGetRandom = async () => {

  if (!selectedSubject.value) {
    showError("Validasi", "Pilih subjek terlebih dahulu.");
    return;
  }
  if (randomCount.value <= 0) {
    showError("Validasi", "Jumlah soal harus lebih dari 0.");
    return;
  }

  loadingRandom.value = true;
  try {
   
    const creatorFilter = isAdminRoute.value ? null : user.value?.id;

 
    const res = await getRandomQuestions(
        randomCount.value, 
        selectedSubject.value, 
        creatorFilter
    );
    

    const list = Array.isArray(res) ? res : (res?.data || []);

    if (!list || list.length === 0) {
      showError("Kosong", "Tidak ada soal tersedia di subjek ini.");
      return;
    }


    let addedCount = 0;
    
    list.forEach(q => {
      const qId = Number(q.id);

      
    
      if (!isQuestionAlreadyAdded(qId) && !selectedQuestions.value.includes(qId)) {
        selectedQuestions.value.push(qId);
        addedCount++;
      }
    });

   
    if (addedCount > 0) {
      showSuccess("Berhasil", `Berhasil memilih ${addedCount} soal secara acak! Jangan lupa klik 'Simpan Pilihan'.`);
    } else {
      showSuccess("Info", "Soal-soal yang terambil random ternyata sudah ada di ujian ini.");
    }

  } catch (err) {
    console.error("Error Random:", err);
    showError("Gagal", "Gagal mengambil soal random. Pastikan koneksi aman.");
  } finally {
    loadingRandom.value = false;
  }
};


const selectAllBySubject = async () => {
  if (!selectedSubject.value) return;
  
  loadingAllSubject.value = true;
  try {
    let allFetchedIds = [];
    let offset = 0;
    const BATCH_LIMIT = 50; 
    let hasMoreData = true;

    while (hasMoreData) {
       let result;

   
       if (isAdminRoute.value) {
   
          result = await getQuestionsBySubject(selectedSubject.value, BATCH_LIMIT, offset);
       } else {
        
          result = await getQuestionsByCreatorAndSubject(user.value.id, selectedSubject.value, BATCH_LIMIT, offset);
       }
      

      
       let chunk = [];
       
       if (result.items) {
          chunk = result.items; 
       } else if (Array.isArray(result)) {
          chunk = result;       
       } else if (result.data && Array.isArray(result.data)) {
          chunk = result.data;  
       } else if (result.data && result.data.data && Array.isArray(result.data.data)) {
          chunk = result.data.data; 
       }

       if (chunk.length > 0) {
          chunk.forEach(q => allFetchedIds.push(q.id));
          offset += BATCH_LIMIT;
       } 
       
     
       if (chunk.length < BATCH_LIMIT) {
          hasMoreData = false;
       }
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
    console.error(err);
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