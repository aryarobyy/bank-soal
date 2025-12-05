<template>
  <div v-if="user" class="min-h-screen bg-[#fafafa] font-sans pb-24">
    
    <div class="bg-white px-6 pt-8 pb-6 shadow-sm border-b border-gray-100">
      <div class="flex justify-between items-start">
        <div>
          <h1 class="text-2xl font-extrabold text-gray-800 leading-tight">
            Hai, {{ user.name.split(' ')[0] }}! 👋
          </h1>
          <p class="text-sm text-gray-500 mt-1 font-medium">{{ user.nim || 'Mahasiswa' }}</p>
        </div>
        <div @click="goToProfile" class="cursor-pointer p-1 rounded-full border-2 border-blue-100 hover:border-blue-300 transition">
          <img 
            :src="user.img_url || `https://ui-avatars.com/api/?name=${user.name}&background=random`" 
            alt="Profile" 
            class="w-12 h-12 rounded-full object-cover"
          />
        </div>
      </div>
    </div>

    <div class="p-6 max-w-3xl mx-auto">
      <router-link to="/ujian" class="group relative overflow-hidden bg-gradient-to-r from-blue-600 to-indigo-600 rounded-3xl p-6 shadow-lg shadow-blue-200/50 mb-8 flex items-center justify-between hover:shadow-xl transition-all active:scale-[0.98]">
         <div class="absolute inset-0 opacity-10" style="background-image: radial-gradient(circle, #ffffff 1px, transparent 1px); background-size: 16px 16px;"></div>
         
         <div class="relative z-10">
           <h2 class="text-white text-xl font-bold mb-2">Mulai Ujian Baru</h2>
           <p class="text-blue-100 text-sm leading-relaxed">Cari dan kerjakan ujian yang<br>tersedia untuk Anda.</p>
         </div>
         <div class="relative z-10 w-14 h-14 bg-white/20 backdrop-blur-sm rounded-full flex items-center justify-center group-hover:bg-white/30 transition">
            <svg xmlns="http://www.w3.org/2000/svg" class="w-7 h-7 text-white" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
              <path stroke-linecap="round" stroke-linejoin="round" d="M14 5l7 7m0 0l-7 7m7-7H3" />
            </svg>
         </div>
      </router-link>

      <div class="grid grid-cols-2 gap-4 mb-10">
        <button @click="goToHistory" class="group bg-white p-5 rounded-2xl shadow-sm border border-gray-100 hover:border-green-200 hover:shadow-md transition active:scale-95 flex flex-col items-start gap-3 relative overflow-hidden">
          <div class="w-12 h-12 bg-green-50 text-green-600 rounded-xl flex items-center justify-center group-hover:bg-green-600 group-hover:text-white transition">
            <svg xmlns="http://www.w3.org/2000/svg" class="w-6 h-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2m-3 7h3m-3 4h3m-6-4h.01M9 16h.01" />
            </svg>
          </div>
          <div>
            <h3 class="font-bold text-gray-800 text-base">Riwayat</h3>
            <p class="text-gray-500 text-xs mt-1">Lihat hasil ujian</p>
          </div>
        </button>

        <button @click="goToProfile" class="group bg-white p-5 rounded-2xl shadow-sm border border-gray-100 hover:border-purple-200 hover:shadow-md transition active:scale-95 flex flex-col items-start gap-3 relative overflow-hidden">
          <div class="w-12 h-12 bg-purple-50 text-purple-600 rounded-xl flex items-center justify-center group-hover:bg-purple-600 group-hover:text-white transition">
             <svg xmlns="http://www.w3.org/2000/svg" class="w-6 h-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
               <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z" />
             </svg>
          </div>
          <div>
            <h3 class="font-bold text-gray-800 text-base">Profil Saya</h3>
             <p class="text-gray-500 text-xs mt-1">Atur akun Anda</p>
          </div>
        </button>
      </div>

      <div>
        <div class="flex justify-between items-center mb-6">
          <h3 class="text-lg font-bold text-gray-900">Baru Saja Dikerjakan</h3>
          <button @click="goToHistory" class="text-xs text-gray-500 font-medium hover:text-blue-600 transition">Lihat Semua</button>
        </div>

        <div v-if="loading" class="space-y-4">
          <div class="h-20 bg-gray-100 rounded-2xl animate-pulse"></div>
          <div class="h-20 bg-gray-100 rounded-2xl animate-pulse"></div>
        </div>

        <div v-else-if="recentSessions.length === 0" class="text-center py-10 bg-white rounded-3xl border border-dashed border-gray-200">
          <p class="text-gray-400 text-sm">Belum ada aktivitas ujian.</p>
        </div>

        <div v-else class="space-y-3">
          <div 
            v-for="session in recentSessions" 
            :key="session.id" 
            class="bg-white p-4 rounded-2xl shadow-sm border border-gray-100 flex items-center justify-between hover:bg-gray-50 transition"
          >
            <div class="flex items-center gap-4 overflow-hidden w-full">
              <div class="flex-shrink-0 w-10 h-10 rounded-xl bg-blue-50 text-blue-600 flex items-center justify-center">
                <svg xmlns="http://www.w3.org/2000/svg" class="w-5 h-5" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                   <path stroke-linecap="round" stroke-linejoin="round" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" />
                </svg>
              </div>

              <div class="min-w-0 flex-1">
                <h4 class="font-bold text-gray-800 text-sm truncate">
                  {{ getExamTitle(session.exam_id) }}
                </h4>
                <p class="text-xs text-gray-500 mt-0.5">
                  {{ session.status === 'finished' ? 'Selesai: ' + formatDate(session.finished_at) : 'Sedang dikerjakan...' }}
                </p>
              </div>
            </div>
            </div>
        </div>
      </div>
    </div>

  </div>

  <HomeView v-else />
</template>

<script setup>
import { ref, onMounted, watch } from "vue";
import { useRouter } from "vue-router";
import { useGetCurrentUser } from "../../hooks/useGetCurrentUser";
import { getExamSessionByUser } from "../../provider/examsession.provider";
import { getAllExam } from "../../provider/exam.provider"; 
import HomeView from './HomeView.vue'; 

const router = useRouter();
const { user } = useGetCurrentUser();

const loading = ref(true);
const recentSessions = ref([]);
const examMap = ref({}); 

const goToHistory = () => { router.push('/ujian'); };
const goToProfile = () => {
  if (user.value) {
    router.push({ name: 'Profile', params: { id: user.value.id } });
  }
};

const formatDate = (dateString) => {
  if (!dateString || dateString.startsWith('0001')) return "-";
  const date = new Date(dateString);
  return date.toLocaleDateString("id-ID", { day: 'numeric', month: 'short', hour: '2-digit', minute:'2-digit' });
};

const getExamTitle = (examId) => {
  return examMap.value[examId] || `Ujian #${examId}`;
};

const fetchDashboardData = async () => {
  if (!user.value) return;
  
  try {
    loading.value = true;
    
  
    const resSession = await getExamSessionByUser(user.value.id, 50, 0);
    
    let sessionData = [];
    if (Array.isArray(resSession)) {
        sessionData = resSession;
    } else if (resSession.data && Array.isArray(resSession.data)) {
        sessionData = resSession.data; 
    } else if (resSession.data) {
        sessionData = resSession.data; 
    }

    sessionData.sort((a, b) => {
        return new Date(b.started_at) - new Date(a.started_at);
    });

    recentSessions.value = sessionData.slice(0, 3);

    
    const resExam = await getAllExam(100, 0); 
    let examData = [];
    if (Array.isArray(resExam)) examData = resExam;
    else if (resExam.data) examData = resExam.data;

    examData.forEach(e => {
      examMap.value[e.id] = e.title;
    });

  } catch (err) {
    console.error("Gagal load dashboard:", err);
  } finally {
    loading.value = false;
  }
};

onMounted(() => {
  if (user.value) {
    fetchDashboardData();
  } else {
    const unwatch = watch(user, (val) => {
      if (val) {
        fetchDashboardData();
        unwatch();
      }
    });
  }
});
</script>