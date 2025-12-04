<template>
  <div class="p-6 bg-[#E8EDFF] min-h-screen">
    <div class="flex flex-col md:flex-row justify-between items-start md:items-center mb-6 gap-4">
      <h1 class="text-3xl font-bold text-[#2A4DFF]">Laporan Nilai Ujian</h1>
      
      <button
        @click="downloadExcel"
        :disabled="reports.length === 0"
        class="flex items-center gap-2 bg-green-500 text-white px-4 py-2 rounded hover:bg-green-600 transition disabled:bg-gray-400 disabled:cursor-not-allowed"
      >
        <i class="fas fa-file-excel"></i>
        Unduh Excel
      </button>
    </div>

    <div class="bg-white p-4 rounded-lg shadow-sm mb-6 border border-blue-100">
        <label class="block text-sm font-semibold text-gray-700 mb-2">Pilih Ujian:</label>
        <select 
            v-model="selectedExamId" 
            class="w-full md:w-1/2 p-2 border rounded-md bg-gray-50 focus:outline-none focus:ring-2 focus:ring-blue-500"
        >
            <option :value="null" disabled>-- Pilih Ujian --</option>
            <option v-for="exam in examList" :key="exam.id" :value="exam.id">
                {{ exam.title }}
            </option>
        </select>
    </div>

    <div v-if="loading" class="text-center py-10 text-gray-500">Memuat data laporan...</div>
    <div v-else-if="error" class="text-center py-10 text-red-500">{{ error }}</div>

    <div v-else class="bg-white p-6 rounded-xl shadow-md overflow-x-auto">
      <div v-if="!selectedExamId" class="text-center py-10 text-gray-500">
        Silakan pilih ujian terlebih dahulu.
      </div>
      
      <div v-else>
        <div class="mb-4 text-sm text-gray-600 flex justify-between items-center">
            <span>Halaman <strong>{{ currentPage }}</strong> dari <strong>{{ totalPages }}</strong></span>
            <span>Total Data: <strong>{{ totalItems }}</strong></span>
        </div>

        <table class="w-full border-collapse text-sm">
            <thead>
            <tr class="bg-[#f5f7ff] text-gray-600">
                <th class="p-3 text-left">No</th>
                <th class="p-3 text-left">Nama Mahasiswa</th> 
                <th class="p-3 text-left">Judul Ujian</th>
                <th class="p-3 text-left">Score</th>
                <th class="p-3 text-left">Status</th>
                <th class="p-3 text-left">Tanggal Selesai</th>
            </tr>
            </thead>
            <tbody>
            <tr
                v-for="(session, index) in reports"
                :key="session.id"
                class="border-t hover:bg-[#eef3ff]"
            >
                <td class="p-3">{{ (currentPage - 1) * limit + index + 1 }}</td>
                
                <td class="p-3 font-medium">
                  {{ usersMap[session.user_id] || `Loading ID: ${session.user_id}...` }}
                </td>
                
                <td class="p-3 text-gray-500">
                  {{ getExamTitle(session.exam_id) }}
                </td>
                
                <td class="p-3 font-bold text-blue-600">
                  {{ session.score }}
                </td>
                
                <td class="p-3">
                <span 
                    class="px-2 py-1 rounded text-xs font-semibold"
                    :class="session.is_passed ? 'bg-green-100 text-green-700' : 'bg-red-100 text-red-700'"
                >
                    {{ session.is_passed ? 'Lulus' : 'Tidak Lulus' }}
                </span>
                </td>

                <td class="p-3">
                {{ formatDate(session.finished_at || session.updated_at) }}
                </td>
            </tr>
            <tr v-if="reports.length === 0">
                <td colspan="6" class="p-8 text-center text-gray-400">
                Belum ada data untuk ujian ini.
                </td>
            </tr>
            </tbody>
        </table>

        <div v-if="totalPages > 1" class="flex justify-end items-center gap-2 mt-6">
            <button @click="prevPage" :disabled="currentPage === 1" class="px-3 py-1 bg-gray-100 border rounded hover:bg-gray-200 disabled:opacity-50">Prev</button>
            <span class="text-sm text-gray-600">{{ currentPage }} / {{ totalPages }}</span>
            <button @click="nextPage" :disabled="currentPage >= totalPages" class="px-3 py-1 bg-gray-100 border rounded hover:bg-gray-200 disabled:opacity-50">Next</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, watch, computed } from "vue";
import * as XLSX from "xlsx";
import { getAllExam } from "../../provider/exam.provider"; 
import { getUserById } from "../../provider/user.provider"; 
import { getExamSessions } from "../../provider/examsession.provider"; 

const reports = ref([]);
const examList = ref([]);
const usersMap = ref({});
const selectedExamId = ref(null);
const loading = ref(false);
const error = ref("");


const currentPage = ref(1);
const limit = 10;
const totalItems = ref(0);

const totalPages = computed(() => {
    if (totalItems.value === 0) return 1;
    return Math.ceil(totalItems.value / limit);
});


const fetchExamList = async () => {
    try {
        const res = await getAllExam(100, 0);
        if (Array.isArray(res)) examList.value = res;
        else if (res.data) examList.value = res.data;
        else examList.value = [];
    } catch (e) { console.error(e); }
};

const getExamTitle = (examId) => {
    const exam = examList.value.find(e => e.id === examId);
    return exam ? exam.title : `Exam ID: ${examId}`;
};


const fetchReports = async () => {
  if (!selectedExamId.value) return;

  try {
    loading.value = true;
    error.value = "";
    
    const offset = (currentPage.value - 1) * limit;

    
    const result = await getExamSessions(limit, offset, selectedExamId.value);
    
    if (result && Array.isArray(result.data)) {
        reports.value = result.data;
        totalItems.value = result.total || 0;
    } else if (Array.isArray(result)) {
        
        reports.value = result;
        totalItems.value = result.length;
    } else {
        reports.value = [];
        totalItems.value = 0;
    }
    

    if (reports.value.length > 0) {
        await fetchUserNames(reports.value);
    }

  } catch (err) {
    console.error(err);
    error.value = "Gagal memuat data laporan.";
    reports.value = [];
  } finally {
    loading.value = false;
  }
};

const prevPage = () => { if (currentPage.value > 1) { currentPage.value--; fetchReports(); } };
const nextPage = () => { if (currentPage.value < totalPages.value) { currentPage.value++; fetchReports(); } };

watch(selectedExamId, () => {
    currentPage.value = 1;
    reports.value = [];
    fetchReports();
});

const fetchUserNames = async (reportsData) => {
    const userIds = [...new Set(reportsData.map(r => r.user_id))];
    for (const id of userIds) {
        if (!usersMap.value[id]) {
            try {
                const userRes = await getUserById(id);
                const userName = userRes.data?.name || userRes.name || "Unknown";
                usersMap.value[id] = userName;
            } catch (e) { usersMap.value[id] = `User ${id}`; }
        }
    }
};

const formatDate = (d) => d ? new Date(d).toLocaleDateString("id-ID", { day: "numeric", month: "short", year: "numeric", hour: "2-digit", minute: "2-digit" }) : "-";

const downloadExcel = () => {
  const dataToExport = reports.value.map((item, index) => ({
    No: index + 1,
    "Nama Mahasiswa": usersMap.value[item.user_id] || item.user_id,
    "Judul Ujian": getExamTitle(item.exam_id),
    "Nilai": item.score,
    "Status": item.is_passed ? 'Lulus' : 'Tidak Lulus',
    "Tanggal Selesai": formatDate(item.finished_at)
  }));

  const ws = XLSX.utils.json_to_sheet(dataToExport);
  const wb = XLSX.utils.book_new();
  XLSX.utils.book_append_sheet(wb, ws, "Laporan Nilai");
  XLSX.writeFile(wb, `Laporan_Nilai_Page_${currentPage.value}.xlsx`);
};

onMounted(() => { fetchExamList(); });
</script>