<template>
  <div class="p-6 bg-[#E8EDFF] min-h-screen">
    <div class="flex flex-col md:flex-row justify-between items-start md:items-center mb-6 gap-4">
      <h1 class="text-3xl font-bold text-[#2A4DFF]">Laporan Nilai Ujian</h1>
      
      <button
        @click="downloadExcel"
        :disabled="isDownloading || !selectedExamId"
        class="flex items-center gap-2 bg-green-500 text-white px-4 py-2 rounded hover:bg-green-600 transition disabled:bg-gray-400 disabled:cursor-not-allowed shadow-sm"
      >
        <i v-if="isDownloading" class="fas fa-spinner fa-spin"></i>
        <i v-else class="fas fa-file-excel"></i>
        {{ isDownloading ? 'Mengunduh...' : 'Unduh Excel (Semua Data)' }}
      </button>
    </div>

    <div class="bg-white p-6 rounded-lg shadow-sm mb-6 border border-blue-100">
        <label class="block text-sm font-semibold text-gray-700 mb-2">Pilih Ujian:</label>
        <div class="relative">
          <select 
              v-model="selectedExamId" 
              class="w-full md:w-1/2 p-3 border border-gray-300 rounded-lg bg-gray-50 focus:outline-none focus:ring-2 focus:ring-blue-500 appearance-none"
          >
              <option :value="null" disabled>-- Pilih Ujian --</option>
              <option v-for="exam in examList" :key="exam.id" :value="exam.id">
                {{ exam.title }}
              </option>
          </select>
          <div class="pointer-events-none absolute inset-y-0 right-0 md:right-1/2 flex items-center px-3 text-gray-700">
            <i class="fas fa-chevron-down text-xs"></i>
          </div>
        </div>
    </div>

    <div v-if="loading" class="text-center py-12">
        <div class="animate-spin rounded-full h-10 w-10 border-b-2 border-blue-600 mx-auto mb-4"></div>
        <p class="text-gray-500">Memuat data laporan...</p>
    </div>

    <div v-else-if="reports.length === 0 && selectedExamId" class="text-center py-12 bg-white rounded-lg shadow-sm border border-dashed border-gray-300">
        <i class="fas fa-clipboard-list text-4xl text-gray-300 mb-3"></i>
        <p class="text-gray-500">Belum ada data nilai untuk halaman ini.</p>
    </div>

    <div v-else-if="reports.length > 0" class="bg-white rounded-lg shadow-sm overflow-hidden border border-gray-200">
      <div class="overflow-x-auto">
        <table class="w-full text-left border-collapse">
          <thead class="bg-gray-50 text-gray-600 uppercase text-xs tracking-wider">
            <tr>
              <th class="p-4 border-b font-semibold text-center w-16">No</th>
              <th class="p-4 border-b font-semibold">Nama Mahasiswa</th>
              <th class="p-4 border-b font-semibold text-center">Nilai</th>
              <th class="p-4 border-b font-semibold text-center">Status</th>
              <th class="p-4 border-b font-semibold text-center">Waktu Selesai</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-gray-100">
            <tr v-for="(item, index) in reports" :key="item.id" class="hover:bg-blue-50 transition">
              <td class="p-4 text-center text-gray-500">
                {{ (currentPage - 1) * itemsPerPage + index + 1 }}
              </td>
              <td class="p-4 font-medium text-gray-800">
                 {{ usersMap[item.user_id] || 'Memuat...' }}
                 <span class="text-xs text-gray-400 block font-normal">ID: {{ item.user_id }}</span>
              </td>
              <td class="p-4 text-center font-bold text-blue-600 text-lg">
                {{ item.score }}
              </td>
              <td class="p-4 text-center">
                <span 
                  class="px-3 py-1 rounded-full text-xs font-bold uppercase"
                  :class="item.is_passed ? 'bg-green-100 text-green-700' : 'bg-red-100 text-red-700'"
                >
                  {{ item.is_passed ? 'Lulus' : 'Tidak Lulus' }}
                </span>
              </td>
              <td class="p-4 text-center text-gray-500 text-sm">
                {{ formatDate(item.finished_at) }}
              </td>
            </tr>
          </tbody>
        </table>
      </div>

      <div class="flex justify-between items-center p-4 bg-gray-50 border-t">
        <span class="text-sm text-gray-600">
           Halaman {{ currentPage }} </span>
        <div class="flex gap-2">
            <button 
                @click="prevPage" 
                :disabled="currentPage === 1"
                class="px-3 py-1 bg-white border border-gray-300 rounded hover:bg-gray-100 disabled:opacity-50 disabled:cursor-not-allowed"
            >
                Prev
            </button>
            <button 
                @click="nextPage" 
                :disabled="reports.length < itemsPerPage"
                class="px-3 py-1 bg-white border border-gray-300 rounded hover:bg-gray-100 disabled:opacity-50 disabled:cursor-not-allowed"
            >
                Next
            </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, watch, computed } from "vue";
import { useRoute } from "vue-router";
import * as XLSX from "xlsx";

import { getAllExam, getExamsByCreator } from "../../provider/exam.provider";
import { getExamSessions } from "../../provider/examsession.provider"; 
import { getUserById } from "../../provider/user.provider";
import { useGetCurrentUser } from "../../hooks/useGetCurrentUser"; 

// --- STATE ---
const examList = ref([]);
const selectedExamId = ref(null);
const reports = ref([]);
const usersMap = ref({});
const loading = ref(false);
const isDownloading = ref(false);


const currentPage = ref(1);
const itemsPerPage = 10; 


const { user } = useGetCurrentUser();
const route = useRoute();

const isAdminRoute = computed(() => route.path.startsWith('/admin') );


const fetchExamList = async () => {
    try {
        let res;
        if (isAdminRoute.value) {
            const response = await getAllExam(100, 0); 
            res = response.data || response;
        } else {
            if (!user.value?.id) return;
            res = await getExamsByCreator(user.value.id, 100, 0);
        }
        
        if (Array.isArray(res)) {
            examList.value = res;
        } else if (res && res.data) {
            examList.value = res.data;
        } else {
            examList.value = [];
        }
    } catch (e) {
        console.error("Gagal mengambil daftar ujian:", e);
    }
};


const fetchReports = async () => {
    if (!selectedExamId.value) return;
    loading.value = true;
    
    try {
        const offset = (currentPage.value - 1) * itemsPerPage;
        
      
        const res = await getExamSessions(itemsPerPage, offset, selectedExamId.value);
        
        let data = [];
        
        if (Array.isArray(res)) {
            data = res;
        } else if (res && res.data && Array.isArray(res.data)) {
            data = res.data; 
        } else if (res && Array.isArray(res.data)) {
            data = res.data;
        }

        reports.value = data;
        
     
        if (data.length > 0) {
            await fetchUserNames(data);
        }

    } catch (e) {
        console.error("Gagal mengambil laporan:", e);
        reports.value = [];
    } finally {
        loading.value = false;
    }
};


const nextPage = () => {

    if (reports.value.length < itemsPerPage) return;
    currentPage.value++;
    fetchReports();
};

const prevPage = () => {
    if (currentPage.value > 1) {
        currentPage.value--;
        fetchReports();
    }
};

const fetchUserNames = async (reportsData) => {
    const userIds = [...new Set(reportsData.map(r => r.user_id))];
    for (const id of userIds) {
        if (!usersMap.value[id]) { 
            try {
                const userRes = await getUserById(id);
                const userName = userRes.data?.name || userRes.name || "Unknown User";
                usersMap.value[id] = userName;
            } catch (e) { 
                usersMap.value[id] = `User ID ${id}`; 
            }
        }
    }
};

const getExamTitle = (examId) => {
    const ex = examList.value.find(e => e.id === examId);
    return ex ? ex.title : `Exam #${examId}`;
};

const formatDate = (d) => {
    if (!d) return "-";
    return new Date(d).toLocaleDateString("id-ID", { 
        day: "numeric", month: "short", year: "numeric", 
        hour: "2-digit", minute: "2-digit" 
    });
};


const downloadExcel = async () => {
  if (!selectedExamId.value) return;
  isDownloading.value = true;

  try {
    let allData = [];
    let offset = 0;
    const BATCH_LIMIT = 100; 
    let keepFetching = true;

   
    while (keepFetching) {
        const res = await getExamSessions(BATCH_LIMIT, offset, selectedExamId.value);
        
        let chunk = [];
        if (Array.isArray(res)) chunk = res;
        else if (res && res.data && Array.isArray(res.data)) chunk = res.data;
        else if (res && Array.isArray(res.data)) chunk = res.data;

        if (chunk.length > 0) {
            allData = [...allData, ...chunk];
            offset += BATCH_LIMIT;
            if (chunk.length < BATCH_LIMIT) keepFetching = false;
        } else {
            keepFetching = false;
        }
    }


    const finishedData = allData.filter(s => s.status === 'finished' || s.status === 'submitted');
    
  
    await fetchUserNames(finishedData);


    const dataToExport = finishedData.map((item, index) => ({
      No: index + 1,
      "Nama Mahasiswa": usersMap.value[item.user_id] || `ID: ${item.user_id}`,
      "Judul Ujian": getExamTitle(item.exam_id),
      "Nilai": item.score,
      "Status": item.is_passed ? 'Lulus' : 'Tidak Lulus',
      "Tanggal Selesai": formatDate(item.finished_at)
    }));


    const ws = XLSX.utils.json_to_sheet(dataToExport);
    const wb = XLSX.utils.book_new();
    const wscols = [{wch: 5}, {wch: 30}, {wch: 30}, {wch: 10}, {wch: 15}, {wch: 20}];
    ws['!cols'] = wscols;

    XLSX.utils.book_append_sheet(wb, ws, "Laporan Nilai");
    const fileName = `Laporan_Ujian_${getExamTitle(selectedExamId.value)}.xlsx`.replace(/ /g, "_");
    XLSX.writeFile(wb, fileName);
    
  } catch (e) {
    console.error("Gagal download excel:", e);
    alert("Gagal mengunduh Excel: " + e.message);
  } finally {
    isDownloading.value = false;
  }
};


watch(selectedExamId, () => {
    currentPage.value = 1; 
    reports.value = [];
    fetchReports();
});

onMounted(() => {
    fetchExamList();
});
</script>