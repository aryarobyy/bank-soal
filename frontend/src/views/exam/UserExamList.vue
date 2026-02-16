<template>
  <div class="min-h-screen bg-[#f4f7fc] font-sans">
    <main class="max-w-7xl mx-auto px-4 sm:px-6 py-10">
      
      <div class="mb-8 text-center sm:text-left">
        <h1 class="text-3xl font-bold text-[#2c3e50]">Pusat Ujian</h1>
        <p class="text-gray-500 mt-2">Akses ujian tersedia dan lihat riwayat pengerjaan Anda.</p>
      </div>

      <div class="flex border-b border-gray-200 mb-8">
        <button 
          @click="activeTab = 'available'" 
          :class="['px-6 py-3 font-medium text-sm transition-colors border-b-2', activeTab === 'available' ? 'border-blue-600 text-blue-600' : 'border-transparent text-gray-500 hover:text-gray-700']"
        >
          <i class="fas fa-list-ul mr-2"></i> Ujian Tersedia
        </button>
        <button 
          @click="activeTab = 'history'" 
          :class="['px-6 py-3 font-medium text-sm transition-colors border-b-2', activeTab === 'history' ? 'border-blue-600 text-blue-600' : 'border-transparent text-gray-500 hover:text-gray-700']"
        >
          <i class="fas fa-history mr-2"></i> Riwayat Saya
        </button>
      </div>

      <div v-if="activeTab === 'available'">
        <div class="bg-white p-4 rounded-xl shadow-sm border border-gray-100 mb-8">
          <div class="flex flex-col md:flex-row gap-4">
            <div class="flex-1 relative">
              <span class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                <i class="fas fa-search text-gray-400"></i>
              </span>
              <input v-model="searchQuery" type="text" placeholder="Cari judul ujian..." class="w-full pl-10 pr-4 py-2.5 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500 transition" />
            </div>
            <div class="w-full md:w-48">
              <select v-model="selectedDifficulty" class="w-full px-4 py-2.5 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 bg-white">
                <option value="">Semua Tingkat</option>
                <option value="easy">Easy</option>
                <option value="medium">Medium</option>
                <option value="hard">Hard</option>
              </select>
            </div>
          </div>
        </div>

        <div v-if="loading" class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
          <div v-for="n in 6" :key="n" class="h-48 bg-gray-200 rounded-xl animate-pulse"></div>
        </div>

        <div v-else-if="filteredExams.length === 0" class="text-center py-16 bg-white rounded-xl shadow-sm">
          <i class="fas fa-folder-open text-4xl text-gray-300 mb-4"></i>
          <p class="text-gray-500 text-lg">Tidak ada ujian yang ditemukan.</p>
          <button @click="resetFilter" class="mt-4 text-blue-600 hover:underline">Reset Filter</button>
        </div>

        <div v-else class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
          <div v-for="exam in filteredExams" :key="exam.id" class="bg-white rounded-xl shadow-sm border border-gray-100 hover:shadow-md transition-shadow duration-300 flex flex-col">
            <div class="p-6 flex-grow">
              <div class="flex justify-between items-start mb-4">
                <span :class="difficultyBadge(exam.difficulty)" class="px-3 py-1 rounded-full text-xs font-bold uppercase tracking-wide">{{ exam.difficulty }}</span>
                <span class="text-gray-400 text-xs flex items-center gap-1"><i class="far fa-clock"></i> {{ exam.long_time }} Menit</span>
              </div>
              <h3 class="text-xl font-bold text-gray-800 mb-2 line-clamp-2">{{ exam.title }}</h3>
              <p class="text-gray-600 text-sm line-clamp-3 mb-4">{{ exam.description || 'Tidak ada deskripsi ujian.' }}</p>
            </div>
            <div class="p-4 bg-gray-50 border-t border-gray-100 rounded-b-xl flex justify-between items-center">
              <div class="text-xs text-gray-500">
                <p>Mulai: {{ formatDate(exam.started_at) }}</p>
              </div>
              <button @click="goToDetail(exam.id)" class="bg-blue-600 text-white px-4 py-2 rounded-lg text-sm font-semibold hover:bg-blue-700 transition-colors flex items-center gap-2">Lihat Detail <i class="fas fa-arrow-right"></i></button>
            </div>
          </div>
        </div>
      </div>

      <div v-if="activeTab === 'history'">
        
        <div v-if="loadingHistory" class="text-center py-10 text-gray-500">Memuat riwayat...</div>
        
        <div v-else-if="historyList.length === 0" class="text-center py-16 bg-white rounded-xl shadow-sm">
          <i class="fas fa-history text-4xl text-gray-300 mb-4"></i>
          <p class="text-gray-500 text-lg">Belum ada riwayat ujian.</p>
          <button @click="activeTab = 'available'" class="mt-4 text-blue-600 hover:underline">Cari Ujian</button>
        </div>

        <div v-else class="space-y-4">
        <div 
          v-for="session in historyList" 
          :key="session.id" 
          class="bg-white p-6 rounded-xl shadow-sm border border-gray-100 flex flex-col md:flex-row justify-between items-start md:items-center gap-4 hover:shadow-md transition"
        >
          <div>
            <h3 class="text-lg font-bold text-gray-800">{{ getExamTitle(session.exam_id) }}</h3>
            <div class="text-sm text-gray-500 mt-1 space-y-1">
              <p><i class="far fa-calendar-alt mr-2"></i> {{ formatDateFull(session.finished_at || session.updated_at) }}</p>
              <p><i class="fas fa-hashtag mr-2"></i> Sesi ID: #{{ session.id }}</p>
            </div>
          </div>
          
          <div class="flex items-center gap-6 w-full md:w-auto justify-between md:justify-end">
            
            <template v-if="session.status === 'finished' || session.status === 'submitted'">
              <div class="text-right">
                <p class="text-xs text-gray-400 uppercase font-semibold">Status</p>
                <span class="px-3 py-1 rounded-full text-xs font-bold uppercase tracking-wide bg-gray-200 text-gray-700">
                  Selesai
                </span>
              </div>
            </template>
            <template v-else>
              <div class="text-right">
                <p class="text-xs text-gray-400 uppercase font-semibold">Nilai</p>
                <p class="text-lg font-bold text-gray-400">-</p>
              </div>
              <div class="text-right flex flex-col gap-2 items-end">
                <p class="text-xs text-gray-400 uppercase font-semibold">Status</p>
                <span class="px-3 py-1 rounded-full text-xs font-bold uppercase tracking-wide bg-yellow-100 text-yellow-700">
                  Sedang Dikerjakan
                </span>
                <button 
                  @click="goToDetail(session.exam_id)"
                  class="text-xs text-blue-600 font-semibold hover:underline"
                >
                  Lanjutkan Ujian →
                </button>
              </div>
            </template>

         </div>
        </div>
       </div>
      </div>

    </main>
  </div>
</template>

<script setup>
import { ref, onMounted, computed, watch } from "vue";
import { useRouter } from "vue-router";
import { getAllExam } from "../../provider/exam.provider";
import { getExamSessionByUser } from "../../provider/examsession.provider";
import { useGetCurrentUser } from "../../hooks/useGetCurrentUser";

import { usePopup } from "../../hooks/usePopup";


const { showError } = usePopup();

const router = useRouter();
const { user } = useGetCurrentUser();

const activeTab = ref('available'); 

const exams = ref([]);
const historyList = ref([]);
const loading = ref(true);
const loadingHistory = ref(false);

const searchQuery = ref("");
const selectedDifficulty = ref("");

const fetchExams = async () => {
  try {
    loading.value = true;
    const res = await getAllExam(100, 0);
    
    if (Array.isArray(res)) exams.value = res;
    else if (res.data && Array.isArray(res.data)) exams.value = res.data;
    else exams.value = [];
    
  } catch (err) {
    console.error("Gagal memuat ujian:", err);
    
    showError("Gagal", "Gagal memuat daftar ujian. Coba lagi nanti.");
  } finally {
    loading.value = false;
  }
};


const fetchHistory = async () => {
  if (!user.value) return;
  try {
    loadingHistory.value = true;
    
    const res = await getExamSessionByUser(user.value.id, 50, 0); 
    
    if (Array.isArray(res)) {
        historyList.value = res;
    } else {
        historyList.value = [];
    }
    
    console.log("Riwayat User:", historyList.value);

  } catch (err) {
    console.error("Gagal memuat riwayat:", err);
    
    showError("Gagal", "Gagal memuat riwayat ujian.");
  } finally {
    loadingHistory.value = false;
  }
};


const filteredExams = computed(() => {
  return exams.value.filter(exam => {
    const matchesSearch = exam.title.toLowerCase().includes(searchQuery.value.toLowerCase());
    const matchesDiff = selectedDifficulty.value === "" || exam.difficulty === selectedDifficulty.value;
    return matchesSearch && matchesDiff;
  });
});

const getExamTitle = (examId) => {
  const exam = exams.value.find(e => e.id === examId);
  return exam ? exam.title : `Ujian ID: ${examId}`;
};

const resetFilter = () => { searchQuery.value = ""; selectedDifficulty.value = ""; };
const goToDetail = (id) => { router.push(`/exam/view?id=${id}`); };

const difficultyBadge = (diff) => {
  if (diff === 'easy') return 'bg-green-100 text-green-700';
  if (diff === 'medium') return 'bg-yellow-100 text-yellow-700';
  if (diff === 'hard') return 'bg-red-100 text-red-700';
  return 'bg-gray-100 text-gray-600';
};

const formatDate = (d) => d ? new Date(d).toLocaleDateString("id-ID", { day: 'numeric', month: 'short' }) : "-";
const formatDateFull = (d) => {
  if (!d || d.startsWith("0001")) return "Sedang Berjalan";
  return new Date(d).toLocaleDateString("id-ID", { weekday: 'long', day: 'numeric', month: 'short', year: 'numeric', hour: '2-digit', minute: '2-digit' });
};


onMounted(() => {
  fetchExams();
  
  if (user.value) fetchHistory();
  else {
    const unwatch = watch(user, (val) => {
      if (val) { fetchHistory(); unwatch(); }
    });
  }
});
</script>