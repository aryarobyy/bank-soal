<template>
  <div class="min-h-screen bg-[#e9edfc] flex items-center justify-center px-4">
    <div
      class="w-full max-w-5xl bg-white rounded-3xl shadow-lg p-6 sm:p-8 border border-gray-200"
    >
      <div v-if="loading" class="text-center text-gray-500 py-10">
        Memuat ujian...
      </div>

      <div v-else-if="error" class="text-center text-red-500 py-10">
        {{ error }}
      </div>

      <div
        v-else-if="!questions.length"
        class="text-center text-gray-600 py-10"
      >
        Soal tidak tersedia.
      </div>

      <div v-else>
        <div class="mb-6 flex justify-between items-center">
          <h1 class="text-xl sm:text-2xl font-bold text-gray-800">
            {{ exam?.title || "Ujian" }}
          </h1>

          <div class="text-right">
            <p class="text-sm text-gray-500">
              Soal {{ currentNo }} dari {{ questions.length }}
            </p>
            <p class="text-sm font-semibold text-red-600">
              ⏱ {{ formattedTime }}
            </p>
          </div>
        </div>

        <div class="bg-[#f5f7ff] rounded-2xl p-6 sm:p-8 mb-8">
          
          <div v-if="currentQuestion?.img_url" class="mb-5 flex justify-center">
            <img 
              :src="constructImageUrl(currentQuestion.img_url)" 
              alt="Gambar Soal"
              loading="lazy"
              class="max-h-[300px] max-w-full rounded-lg shadow-sm border border-gray-300 object-contain"
            />
          </div>

          <p class="text-sm sm:text-base font-semibold text-gray-800 mb-4">
            {{ currentQuestion?.question_text || "-" }}
          </p>

          <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
            <button
              v-for="opt in currentQuestion?.options || []"
              :key="opt.id"
              type="button"
              @click="selectAnswer(opt.id)"
              :class="[
                'w-full text-left px-4 py-3 rounded-xl border text-base transition',
                answers[currentNo - 1] === opt.id
                  ? 'bg-blue-600 text-white border-blue-600 shadow'
                  : 'bg-white text-gray-700 border-gray-300 hover:border-blue-400',
              ]"
            >
              <span class="font-bold mr-2">{{ opt.option_label || opt.label }}.</span>
              {{ opt.option_text || opt.text || "NO_TEXT" }}
            </button>
          </div>
        </div>

        <div class="flex flex-wrap justify-center gap-2 mb-6">
          <button
            v-for="n in questions.length"
            :key="n"
            @click="goToQuestion(n)"
            type="button"
            :class="[
              'w-8 h-8 sm:w-9 sm:h-9 text-xs sm:text-sm rounded-full border transition',
              currentNo === n
                ? 'bg-blue-600 text-white border-blue-600'
                : answers[n-1] 
                  ? 'bg-green-100 text-green-700 border-green-300'
                  : 'bg-white text-gray-700 border-gray-300 hover:border-blue-400',
            ]"
          >
            {{ n }}
          </button>
        </div>

        <div class="flex justify-between items-center">
          <button
            class="px-4 py-2 rounded-md text-sm bg-gray-200 text-gray-700 disabled:opacity-50"
            :disabled="currentNo === 1"
            @click="prevQuestion"
          >
            ← Sebelumnya
          </button>

          <button
            v-if="currentNo < questions.length"
            class="px-4 py-2 rounded-md text-sm bg-blue-600 text-white hover:bg-blue-700"
            @click="nextQuestion"
          >
            Selanjutnya →
          </button>

          <button
            v-else
            class="px-4 py-2 rounded-md text-sm bg-green-600 text-white hover:bg-green-700"
            @click="finishExam(false)"
          >
            Selesaikan Ujian
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted, onUnmounted, watch } from "vue"; // Tambahkan onUnmounted
import { useRoute, useRouter } from "vue-router";
import { useGetCurrentUser } from "../../hooks/useGetCurrentUser";
import { API_BASE_URL } from "../../core/constant"; 

// Providers
import { getExamById } from "../../provider/exam.provider";
import { getQuestionsByExam } from "../../provider/question.provider"; 
import { submitUserAnswer, getUserAnswersBySession } from "../../provider/useranswer.provider";
import { finishExamSession, updateCurrentNo, getExamSessionById } from "../../provider/examsession.provider";

const route = useRoute();
const router = useRouter();
const { user } = useGetCurrentUser();

// State
const loading = ref(true);
const error = ref("");
const exam = ref(null);
const questions = ref([]);
const currentNo = ref(1);
const answers = ref([]); 

// Timer & Interval State
const timeLeft = ref(0);
let timer = null;
let statusCheckInterval = null; // Variable untuk Polling Interval

const currentQuestion = computed(() => {
  return questions.value[currentNo.value - 1] || null;
});

const selectAnswer = async (optionId) => {
  answers.value[currentNo.value - 1] = optionId;

  const sessionId = Number(route.query.session_id);
  const examId = Number(route.query.id);
  let userId = user.value?.id || Number(localStorage.getItem("id"));
  const questionId = currentQuestion.value.id;
  
  const selectedOption = currentQuestion.value.options.find(opt => opt.id === optionId);
  const answerLabel = selectedOption?.option_label || selectedOption?.label;

  if (!sessionId || !userId || !answerLabel) return;

  try {
    await submitUserAnswer({
      exam_session_id: sessionId,
      user_id: userId,
      question_id: questionId,
      answer: answerLabel,
      exam_id: examId 
    });
  } catch (err) {
    console.error("Gagal menyimpan jawaban:", err);
  }
};

const constructImageUrl = (path) => {
  if (!path) return null;
  if (path.startsWith("http")) return path;
  const cleanPath = path.startsWith("./") ? path.substring(2) : path.startsWith("/") ? path.substring(1) : path;
  return `${API_BASE_URL}/${cleanPath}`;
};


const startTimer = (seconds) => {
  timeLeft.value = seconds;
  if (timer) clearInterval(timer);

  timer = setInterval(() => {
    if (timeLeft.value > 0) {
      timeLeft.value--;
    } else {
      clearInterval(timer);
      // Panggil finishExam dengan mode Auto (true) saat waktu habis
      finishExam(true); 
    }
  }, 1000);
};

const formattedTime = computed(() => {
  const h = Math.floor(timeLeft.value / 3600);
  const m = Math.floor((timeLeft.value % 3600) / 60);
  const s = timeLeft.value % 60;
  
  if (h > 0) {
      return `${h}:${m < 10 ? "0" + m : m}:${s < 10 ? "0" + s : s}`;
  }
  return `${m}:${s < 10 ? "0" + s : s}`;
});


onMounted(async () => {
  const examId = Number(route.query.id);
  const sessionId = Number(route.query.session_id);

  if (!examId || !sessionId) {
    error.value = "ID ujian atau sesi tidak ditemukan.";
    loading.value = false;
    return;
  }

  try {
    const examRes = await getExamById(examId);
    exam.value = examRes?.data || examRes;

    const sessionRes = await getExamSessionById(sessionId);
    const sessionData = sessionRes.data || sessionRes; 

    const qRes = await getQuestionsByExam(examId);
    questions.value = Array.isArray(qRes) ? qRes : qRes.data || [];

    answers.value = Array(questions.value.length).fill(null);

    // 1. Restore Jawaban
    try {
        const savedAnswers = await getUserAnswersBySession(sessionId);
        if (savedAnswers && savedAnswers.length > 0) {
            savedAnswers.forEach(ans => {
                const qIndex = questions.value.findIndex(q => q.id === ans.question_id);
                if (qIndex !== -1) {
                    const question = questions.value[qIndex];
                    const selectedOption = question.options.find(opt => 
                        (opt.option_label || opt.label) === ans.answer
                    );
                    if (selectedOption) {
                        answers.value[qIndex] = selectedOption.id;
                    }
                }
            });
        }
    } catch (err) {
        console.warn("Gagal merestore jawaban:", err);
    }

    // 2. Restore Posisi Soal
    if (sessionData && sessionData.current_no && sessionData.current_no > 0) {
        currentNo.value = sessionData.current_no;
    }

    // 3. Restore Timer Logic
    if (sessionData && sessionData.started_at) {
        const durationMinutes = exam.value?.long_time || 120;
        const startTime = new Date(sessionData.started_at).getTime();
        const now = new Date().getTime();
        const endTime = startTime + (durationMinutes * 60 * 1000);
        const remainingSeconds = Math.floor((endTime - now) / 1000);

        if (remainingSeconds > 0) {
            startTimer(remainingSeconds);
        } else {
            // Jika saat dimuat waktu sudah habis, langsung finish auto
            finishExam(true);
        }
    } else {
        startTimer((exam.value?.long_time || 120) * 60);
    }

    statusCheckInterval = setInterval(async () => {
        try {
            const res = await getExamSessionById(sessionId);
            const currentSession = res.data || res;

            // Jika Backend bilang status sudah 'finished' (karena waktu server habis/admin close)
            // TAPI di frontend kita masih mengerjakan
            if (currentSession.status === 'finished' || currentSession.status === 'submitted') {
                clearInterval(statusCheckInterval);
                clearInterval(timer);
                alert("Sesi ujian telah berakhir menurut Server (Waktu Habis/Ditutup Admin).");
                router.replace("/ujian");
            }
        } catch (err) {
            console.warn("Gagal cek status berkala:", err);
        }
    }, 60000); // Cek setiap 60 detik (1 menit)
    
  } catch (e) {
    console.error(e);
    if (e.response && e.response.status === 500) {
       console.warn("Backend error 500 saat load sesi, menggunakan mode fallback.");
       if (!questions.value.length) {
          const qRes = await getQuestionsByExam(examId);
          questions.value = Array.isArray(qRes) ? qRes : qRes.data || [];
          answers.value = Array(questions.value.length).fill(null);
       }
       startTimer((exam.value?.long_time || 120) * 60);
    } else {
       error.value = "Terjadi kesalahan saat memuat ujian.";
    }
  } finally {
    loading.value = false;
  }
});

// CLEANUP SAAT KELUAR HALAMAN
onUnmounted(() => {
  if (timer) clearInterval(timer);
  if (statusCheckInterval) clearInterval(statusCheckInterval);
});

const goToQuestion = (n) => { currentNo.value = n; };
const nextQuestion = () => { if (currentNo.value < questions.value.length) currentNo.value++; };
const prevQuestion = () => { if (currentNo.value > 1) currentNo.value--; };

watch(currentNo, async (newNo) => {
  const sessionId = Number(route.query.session_id);
  if (sessionId) {
    try { await updateCurrentNo(sessionId, newNo); } catch (e) { console.error(e); }
  }
});



const finishExam = async (isAuto = false) => {
  // 1. Konfirmasi manual (skip jika waktu habis/auto)
  if (!isAuto && timeLeft.value > 0) {
     if (!confirm("Apakah Anda yakin ingin menyelesaikan ujian?")) return;
  }

  // Set loading agar user tahu proses sedang berjalan
  loading.value = true;

  try {
    const sessionId = Number(route.query.session_id);
    const examId = Number(route.query.id); 
    let userId = user.value?.id || Number(localStorage.getItem("id"));

   
    const savePromises = questions.value.map((q, index) => {
        const selectedId = answers.value[index];
        if (selectedId) {
            const selectedOption = q.options.find((opt) => opt.id === selectedId);
            if (selectedOption) {
              // Return promise
                return submitUserAnswer({
                    exam_session_id: sessionId,
                    user_id: userId,
                    question_id: q.id,
                    answer: selectedOption.option_label || selectedOption.label, 
                    exam_id: Number(examId) 
                }).catch(e => console.warn(`Gagal save soal ${q.id}:`, e)); 
            }
        }
        return Promise.resolve();
    });

    // Tunggu semua proses simpan selesai
    await Promise.allSettled(savePromises);

    // ===============================================
    // EKSEKUSI FINISH
    // ===============================================
    const payload = {
      session_id: sessionId,
      exam_id: examId,
      user_id: userId
    };
    
    await finishExamSession(payload);

    // Feedback User
    if (isAuto) {
        alert("Waktu Habis! Jawaban Anda telah disimpan otomatis.");
    } else {
        alert("Ujian selesai! Nilai Anda telah disimpan.");
    }

    // Redirect Keluar (Replace agar tidak bisa Back)
    router.replace("/ujian");

  } catch (err) {
    console.error("FinishExam Error:", err);
    const backendMsg = err.response?.data?.message;
    
    // Handle error khusus
    if (backendMsg && (backendMsg.includes("not found") || backendMsg.includes("record not found"))) {
        router.replace("/ujian");
    } else {
        // Jika mode Auto (waktu habis), tetap paksa keluar walau error
        if (isAuto) {
             router.replace("/ujian");
        } else {
            alert(`Gagal menyelesaikan ujian: ${backendMsg || "Terjadi kesalahan server."}`);
        }
    }
  } finally {
    loading.value = false;
  }
};
</script>