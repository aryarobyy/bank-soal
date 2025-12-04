<template>
  <div class="min-h-screen bg-[#e9edfc] font-sans flex flex-col">
    <main class="flex-grow flex justify-center items-center py-10 px-4">
      <div
        class="w-full max-w-4xl bg-white rounded-3xl shadow-lg p-8 sm:p-10 border border-gray-200"
      >
        <div v-if="loading" class="text-center text-gray-600 py-10">
          Memuat data ujian...
        </div>

        <div v-else-if="error" class="text-center text-red-500 py-10">
          {{ error }}
        </div>

        <div v-else>
          <div class="bg-[#f5f7ff] p-8 rounded-2xl text-center mb-8 border">
            <h2 class="text-3xl font-bold text-gray-800 mb-2">
              {{ exam.title }}
            </h2>
            <p class="text-gray-600 text-sm sm:text-base">
              {{ formattedDate }}
            </p>
            <p class="font-medium text-gray-700 mt-2">
              Durasi:
              <span class="font-semibold text-blue-700">
                {{ exam.long_time }} menit
              </span>
            </p>
            <p class="font-medium text-gray-600 mt-2">
              Kesulitan:
              <span class="font-semibold capitalize text-blue-700">
                {{ exam.difficulty }}
              </span>
            </p>
            
            <div class="mt-4">
               <span v-if="isExamFinishedByUser" class="px-3 py-1 bg-gray-200 text-gray-700 rounded-full text-sm font-bold">Selesai Dikerjakan</span>
               <span v-else-if="isExamClosed" class="px-3 py-1 bg-red-100 text-red-700 rounded-full text-sm font-bold">Ujian Berakhir</span>
               <span v-else-if="!isExamStarted" class="px-3 py-1 bg-yellow-100 text-yellow-700 rounded-full text-sm font-bold">Belum Dimulai</span>
               <span v-else class="px-3 py-1 bg-green-100 text-green-700 rounded-full text-sm font-bold">Sedang Berlangsung</span>
            </div>
          </div>

          <div class="px-2 sm:px-6 mb-6">
            <h3 class="font-semibold text-gray-900 mb-3 text-lg">
              Deskripsi :
            </h3>
            <p class="text-gray-700 leading-relaxed">
              {{ exam.description || "Tidak ada deskripsi." }}
            </p>
          </div>

          <div class="px-2 sm:px-6">
            <h3 class="font-semibold text-gray-900 mb-4 text-lg">
              Aturan Ujian :
            </h3>
            <ul class="text-gray-700 space-y-3 leading-relaxed">
              <li class="flex">
                <span class="mr-2 font-bold text-blue-600">&gt;</span>
                Tidak boleh keluar sebelum ujian selesai.
              </li>
              <li class="flex">
                <span class="mr-2 font-bold text-blue-600">&gt;</span>
                Timer tetap berjalan walaupun internet mati.
              </li>
              <li class="flex">
                <span class="mr-2 font-bold text-blue-600">&gt;</span>
                Setiap peserta hanya diperbolehkan mengerjakan 1 kali.
              </li>
            </ul>
          </div>

          <div class="flex justify-end mt-8">
            <button
              @click="startExam"
              :disabled="isButtonDisabled"
              :class="[
                'px-5 py-2.5 rounded-md font-semibold shadow-sm transition',
                isButtonDisabled 
                  ? 'bg-gray-400 text-gray-200 cursor-not-allowed' 
                  : 'bg-[#2ecc71] hover:bg-[#27ae60] text-white hover:shadow-md'
              ]"
            >
              {{ buttonText }}
            </button>
          </div>
        </div>
      </div>
    </main>

    <Toast ref="toastRef" />
  </div>
</template>

<script setup>
import Toast from "../../components/utils/Toast.vue";
import { ref, onMounted, computed, watch } from "vue";
import { useRouter, useRoute } from "vue-router";
import { useGetCurrentUser } from "../../hooks/useGetCurrentUser";
// Providers
import { getExamById } from "../../provider/exam.provider";
import { createExamSession, getExamSessionByUser } from "../../provider/examsession.provider";


const toastRef = ref(null);
const router = useRouter();
const route = useRoute();
const { user } = useGetCurrentUser();

const exam = ref(null);
const userSession = ref(null); 
const loading = ref(true);
const error = ref("");




const isExamClosed = computed(() => {
    if (!exam.value?.finished_at) return false;
    return new Date() > new Date(exam.value.finished_at);
});


const isExamStarted = computed(() => {
    if (!exam.value?.started_at) return false;
    return new Date() >= new Date(exam.value.started_at);
});


const isExamFinishedByUser = computed(() => {
    return userSession.value?.status === 'finished'; 
});


const isButtonDisabled = computed(() => {
    
    return isExamClosed.value || !isExamStarted.value || isExamFinishedByUser.value;
});


const buttonText = computed(() => {
    if (isExamFinishedByUser.value) return "Anda Sudah Mengerjakan";
    if (isExamClosed.value) return "Ujian Telah Ditutup";
    if (!isExamStarted.value) return "Ujian Belum Dimulai";
   
    if (userSession.value && userSession.value.status === 'in_progress') return "Lanjutkan Ujian"; 
    return "Mulai Ujian";
});


const formattedDate = computed(() => {
  if (!exam.value?.started_at || !exam.value?.finished_at) return "";
  const s = new Date(exam.value.started_at);
  const f = new Date(exam.value.finished_at);

  return `${s.toLocaleDateString("id-ID", {
    day: "2-digit",
    month: "long",
    year: "numeric",
  })} | ${s.toLocaleTimeString("id-ID", {
    hour: "2-digit",
    minute: "2-digit",
  })} - ${f.toLocaleTimeString("id-ID", {
    hour: "2-digit",
    minute: "2-digit",
  })} WIB`;
});


const checkUserSession = async () => {
    if (!user.value || !exam.value) return;
    
    try {
    
        const sessionsRes = await getExamSessionByUser(user.value.id, 100, 0);
        const sessions = Array.isArray(sessionsRes) ? sessionsRes : (sessionsRes.data || []);
        
        const mySession = sessions.find(s => s.exam_id === Number(exam.value.id));
        
        if (mySession) {
            userSession.value = mySession;
        }
    } catch (err) {
        console.warn("Gagal mengecek status sesi user:", err);
    }
};

onMounted(async () => {
  try {
    const id = route.query.id; 
    if(!id) throw new Error("ID Ujian tidak ditemukan");


    const data = await getExamById(id);
    exam.value = data?.data || data; 

  
    if (user.value) {
        await checkUserSession();
    } else {
       
        const unwatch = watch(user, async (val) => {
            if (val) {
                await checkUserSession();
                unwatch();
            }
        });
    }

  } catch (err) {
    console.error("ERROR GET EXAM:", err);
    error.value = "Gagal mengambil data ujian!";
  } finally {
    loading.value = false;
  }
});


const startExam = async () => {

  if (isExamFinishedByUser.value) {
      alert("Anda sudah menyelesaikan ujian ini.");
      return;
  }
  if (isExamClosed.value) {
      alert("Maaf, waktu ujian sudah berakhir.");
      return;
  }
  if (!isExamStarted.value) {
      alert("Maaf, ujian belum dimulai.");
      return;
  }

  try {
  
    const sessionRes = await createExamSession(exam.value.id);
    
    
    const sessionData = sessionRes?.data || sessionRes;
    const sessionId = sessionData?.id;

    if (!sessionId) {
      throw new Error("Gagal mendapatkan ID Sesi Ujian.");
    }

    router.push(`/exam/start?id=${exam.value.id}&session_id=${sessionId}`);

  } catch (err) {
    console.error("Start Exam Error:", err);
    
    const msg = err.response?.data?.message || "";
    

    if (msg.includes("exam is already closed")) {
        alert("Gagal: Ujian ini sudah ditutup oleh sistem.");
    } else if (msg.includes("exam has not started")) {
        alert("Gagal: Ujian belum dimulai.");
    } else if (msg.includes("already finished") || msg.includes("session already finished")) {
        alert("Gagal: Anda sudah menyelesaikan ujian ini.");
        // Refresh status di UI
        checkUserSession(); 
    } else {
        alert(`Tidak dapat memulai ujian: ${msg || "Terjadi kesalahan sistem."}`);
    }
  }
};
</script>