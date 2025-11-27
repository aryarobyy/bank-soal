<template>
  <div class="min-h-screen bg-[#f4f7fc] font-sans">
    <main class="max-w-7xl mx-auto px-4 sm:px-6 py-10">
      <div class="mb-8 text-center sm:text-left">
        <h1 class="text-3xl font-bold text-[#2c3e50]">Daftar Ujian Tersedia</h1>
        <p class="text-gray-500 mt-2">Pilih ujian yang ingin Anda kerjakan hari ini.</p>
      </div>

      <div class="bg-white p-4 rounded-xl shadow-sm border border-gray-100 mb-8">
        <div class="flex flex-col md:flex-row gap-4">
          
          <div class="flex-1 relative">
            <span class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
              <i class="fas fa-search text-gray-400"></i>
            </span>
            <input 
              v-model="searchQuery" 
              type="text" 
              placeholder="Cari judul ujian..." 
              class="w-full pl-10 pr-4 py-2.5 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500 transition"
            />
          </div>

          <div class="w-full md:w-48">
            <select 
              v-model="selectedDifficulty" 
              class="w-full px-4 py-2.5 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500 bg-white"
            >
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
        <div 
          v-for="exam in filteredExams" 
          :key="exam.id" 
          class="bg-white rounded-xl shadow-sm border border-gray-100 hover:shadow-md transition-shadow duration-300 flex flex-col"
        >
          <div class="p-6 flex-grow">
            <div class="flex justify-between items-start mb-4">
              <span :class="difficultyBadge(exam.difficulty)" class="px-3 py-1 rounded-full text-xs font-bold uppercase tracking-wide">
                {{ exam.difficulty }}
              </span>
              <span class="text-gray-400 text-xs flex items-center gap-1">
                <i class="far fa-clock"></i> {{ exam.long_time }} Menit
              </span>
            </div>

            <h3 class="text-xl font-bold text-gray-800 mb-2 line-clamp-2">{{ exam.title }}</h3>
            <p class="text-gray-600 text-sm line-clamp-3 mb-4">
              {{ exam.description || 'Tidak ada deskripsi ujian.' }}
            </p>
          </div>

          <div class="p-4 bg-gray-50 border-t border-gray-100 rounded-b-xl flex justify-between items-center">
            <div class="text-xs text-gray-500">
              <p>Mulai: {{ formatDate(exam.started_at) }}</p>
            </div>
            <button 
              @click="goToDetail(exam.id)" 
              class="bg-blue-600 text-white px-4 py-2 rounded-lg text-sm font-semibold hover:bg-blue-700 transition-colors flex items-center gap-2"
            >
              Lihat Detail <i class="fas fa-arrow-right"></i>
            </button>
          </div>
        </div>
      </div>

    </main>
  </div>
</template>

<script setup>
import { ref, onMounted, computed } from "vue";
import { useRouter } from "vue-router";
import { getAllExam } from "../../provider/exam.provider"; // Pastikan path benar


const router = useRouter();
const exams = ref([]);
const loading = ref(true);

// State untuk Filter & Search
const searchQuery = ref("");
const selectedDifficulty = ref("");

// 1. Fetch Semua Data Ujian
const fetchExams = async () => {
  try {
    loading.value = true;
    // Kita ambil 100 data (limit besar) agar search client-side berjalan mulus
    const res = await getAllExam(100, 0);
    
    // Handle struktur data dari provider (Array atau Object {data: []})
    if (Array.isArray(res)) {
      exams.value = res;
    } else if (res.data && Array.isArray(res.data)) {
      exams.value = res.data;
    } else {
      exams.value = [];
    }
  } catch (err) {
    console.error("Gagal memuat ujian:", err);
  } finally {
    loading.value = false;
  }
};

// 2. Logic Filter & Search (Client-Side)
const filteredExams = computed(() => {
  return exams.value.filter(exam => {
    // Filter berdasarkan Search Text (Case insensitive)
    const matchesSearch = exam.title.toLowerCase().includes(searchQuery.value.toLowerCase());
    
    // Filter berdasarkan Difficulty (Dropdown)
    const matchesDiff = selectedDifficulty.value === "" || exam.difficulty === selectedDifficulty.value;

    return matchesSearch && matchesDiff;
  });
});

// Helper Function
const resetFilter = () => {
  searchQuery.value = "";
  selectedDifficulty.value = "";
};

const goToDetail = (id) => {
  // Arahkan ke halaman ExamView.vue yang lama
  router.push(`/exam/view?id=${id}`); 
};

const difficultyBadge = (diff) => {
  if (diff === 'easy') return 'bg-green-100 text-green-700';
  if (diff === 'medium') return 'bg-yellow-100 text-yellow-700';
  if (diff === 'hard') return 'bg-red-100 text-red-700';
  return 'bg-gray-100 text-gray-600';
};

const formatDate = (dateString) => {
  if (!dateString) return "-";
  return new Date(dateString).toLocaleDateString("id-ID", {
    day: 'numeric', month: 'short'
  });
};

onMounted(() => {
  fetchExams();
});
</script>