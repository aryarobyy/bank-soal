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
        <label class="block text-sm font-semibold text-gray-700 mb-2">Pilih Ujian untuk Melihat Nilai:</label>
        <select 
            v-model="selectedExamId" 
            @change="fetchReports"
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
        <div class="mb-4 text-sm text-gray-600">
            Menampilkan <strong>{{ reports.length }}</strong> data nilai.
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
                v-for="(report, index) in reports"
                :key="report.id"
                class="border-t hover:bg-[#eef3ff]"
            >
                <td class="p-3">{{ index + 1 }}</td>
                
                <td class="p-3 font-medium">
                  {{ usersMap[report.user_id] || `Loading ID: ${report.user_id}...` }}
                </td>
                
                <td class="p-3">
                  {{ getExamTitle(report.exam_id) }}
                </td>
                
                <td class="p-3 font-bold text-blue-600">{{ report.score }}</td>
                
                <td class="p-3">
                <span 
                    class="px-2 py-1 rounded text-xs font-semibold"
                    :class="report.status === 'finished' ? 'bg-green-100 text-green-700' : 'bg-yellow-100 text-yellow-700'"
                >
                    {{ report.status || '-' }}
                </span>
                </td>

                <td class="p-3">
                {{ formatDate(report.updated_at || report.created_at) }}
                </td>
            </tr>
            <tr v-if="reports.length === 0">
                <td colspan="6" class="p-8 text-center text-gray-400">
                Belum ada data nilai untuk ujian ini.
                </td>
            </tr>
            </tbody>
        </table>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from "vue";
import * as XLSX from "xlsx";
import { getExamScores } from "../../provider/examscore.provider"; 
import { getAllExam } from "../../provider/exam.provider"; 
// Import Provider User untuk ambil nama
import { getUserById } from "../../provider/user.provider"; 

const reports = ref([]);
const examList = ref([]);
const selectedExamId = ref(null);
const loading = ref(false);
const error = ref("");

// State untuk menyimpan Nama Mahasiswa: { 101: "Budi", 102: "Siti" }
const usersMap = ref({});

// 1. Load Daftar Ujian
const fetchExamList = async () => {
    try {
        const res = await getAllExam(100, 0);
        if (Array.isArray(res)) {
            examList.value = res;
        } else if (res.data && Array.isArray(res.data)) {
            examList.value = res.data;
        } else {
            examList.value = [];
        }
    } catch (e) {
        console.error("Gagal memuat daftar ujian:", e);
    }
};

// 2. Helper untuk mendapatkan Judul Ujian dari ID (tanpa fetch ulang)
const getExamTitle = (examId) => {
    const exam = examList.value.find(e => e.id === examId);
    return exam ? exam.title : `Exam ID: ${examId}`;
};

// 3. Fetch Reports & Fetch User Names
const fetchReports = async () => {
  if (!selectedExamId.value) return;

  try {
    loading.value = true;
    error.value = "";
    
    // A. Ambil Data Nilai
    const data = await getExamScores(100, 0, selectedExamId.value);
    reports.value = data || [];
    
    // B. Ambil Nama User satu per satu (paralel agar cepat)
    await fetchUserNames(reports.value);

  } catch (err) {
    console.error(err);
    error.value = "Gagal memuat data laporan.";
    reports.value = [];
  } finally {
    loading.value = false;
  }
};

// 4. Fungsi Mengambil Nama User
const fetchUserNames = async (reportsData) => {
    // Kumpulkan semua user_id unik agar tidak fetch double
    const userIds = [...new Set(reportsData.map(r => r.user_id))];

    for (const id of userIds) {
        // Cek apakah sudah ada di memory (cache), jika belum baru fetch
        if (!usersMap.value[id]) {
            try {
                const userRes = await getUserById(id);
                // Sesuaikan dengan struktur respon getUserById
                // Biasanya res.data atau res (tergantung provider Anda)
                const userName = userRes.data?.name || userRes.name || "Unknown";
                
                // Simpan ke map
                usersMap.value[id] = userName;
            } catch (e) {
                console.warn(`Gagal load user ${id}`, e);
                usersMap.value[id] = "User Tidak Ditemukan";
            }
        }
    }
};

const formatDate = (dateString) => {
  if (!dateString) return "-";
  return new Date(dateString).toLocaleDateString("id-ID", {
    day: "numeric", month: "short", year: "numeric", hour: "2-digit", minute: "2-digit"
  });
};

const downloadExcel = () => {
  const dataToExport = reports.value.map((item, index) => ({
    No: index + 1,
    // Ambil nama dari usersMap
    "Nama Mahasiswa": usersMap.value[item.user_id] || item.user_id,
    "Judul Ujian": getExamTitle(item.exam_id),
    "Nilai": item.score,
    "Status": item.status,
    "Tanggal": formatDate(item.updated_at)
  }));

  const ws = XLSX.utils.json_to_sheet(dataToExport);
  const wb = XLSX.utils.book_new();
  XLSX.utils.book_append_sheet(wb, ws, "Laporan Nilai");
  XLSX.writeFile(wb, "Laporan_Nilai_Ujian.xlsx");
};

onMounted(() => {
  fetchExamList();
});
</script>

<style scoped>
table th,
table td {
  border-bottom: 1px solid #ddd;
}
</style>