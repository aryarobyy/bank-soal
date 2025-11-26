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
          <p class="text-sm sm:text-base font-semibold text-gray-800 mb-4">
            Question: {{ currentQuestion?.question_text || "-" }}
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
                  ? 'bg-green-100 text-green-700 border-green-300' // Hijau jika sudah dijawab
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
            @click="finishExam"
          >
            Selesaikan Ujian
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted, watch } from "vue";
import { useRoute, useRouter } from "vue-router";
import { useGetCurrentUser } from "../../hooks/useGetCurrentUser";
// Providers
import { getExamById } from "../../provider/exam.provider";
import { getExamQuestions } from "../../provider/examquestion.provider";
import { submitUserAnswer } from "../../provider/useranswer.provider";
import { finishExamSession, updateCurrentNo } from "../../provider/examsession.provider";

// Router & User
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

// Soal aktif
const currentQuestion = computed(() => {
  return questions.value[currentNo.value - 1] || null;
});

// Simpan jawaban lokal (state)
const selectAnswer = (optionId) => {
  answers.value[currentNo.value - 1] = optionId;
};

// ===================== TIMER =====================
const timeLeft = ref(0);
let timer = null;

const startTimer = (minutes) => {
  timeLeft.value = minutes * 60;

  timer = setInterval(() => {
    if (timeLeft.value > 0) {
      timeLeft.value--;
    } else {
      clearInterval(timer);
      alert("Waktu Habis! Ujian akan diselesaikan otomatis.");
      finishExam(); // Auto finish saat waktu habis
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
// =================================================

// Load data ujian
onMounted(async () => {
  const examId = Number(route.query.id);

  if (!examId) {
    error.value = "ID ujian tidak ditemukan.";
    loading.value = false;
    return;
  }

  try {
    // 1. Get Exam Info
    const examRes = await getExamById(examId);
    exam.value = examRes?.data || examRes;

    // 2. Get Questions
    const qRes = await getExamQuestions(examId);
    questions.value = Array.isArray(qRes) ? qRes : qRes.data || [];

    // 3. Setup array jawaban kosong
    answers.value = Array(questions.value.length).fill(null);

    // 4. Jalankan Timer
    const duration = exam.value?.long_time || 120; 
    startTimer(duration);
    
  } catch (e) {
    console.error(e);
    error.value = "Terjadi kesalahan saat memuat ujian.";
  } finally {
    loading.value = false;
  }
});

// ===================== NAVIGASI & AUTO SAVE POSISI =====================

const goToQuestion = (n) => {
  currentNo.value = n;
};
const nextQuestion = () => {
  if (currentNo.value < questions.value.length) currentNo.value++;
};
const prevQuestion = () => {
  if (currentNo.value > 1) currentNo.value--;
};

// Watcher: Simpan posisi nomor soal ke backend setiap pindah soal
watch(currentNo, async (newNo) => {
  const sessionId = Number(route.query.session_id);
  if (sessionId) {
    try {
      await updateCurrentNo(sessionId, newNo);
    } catch (e) {
      console.error("Gagal menyimpan posisi soal", e);
    }
  }
});

// ===================== FINISH EXAM =====================
const finishExam = async () => {
  // 1. Cek Timer & Konfirmasi
  if (timeLeft.value > 0 && !confirm("Apakah Anda yakin ingin menyelesaikan ujian?")) {
    return;
  }

  try {
    // 2. Ambil Data & Validasi
    // Gunakan Number() untuk memaksa konversi ke angka
    const sessionId = Number(route.query.session_id);
    const examId = Number(route.query.id); 
    
    // Ambil User ID (Prioritas: dari State > LocalStorage)
    let userId = user.value?.id;
    if (!userId) {
        userId = Number(localStorage.getItem("id"));
    }

    // --- DEBUGGING (Cek di Console Browser) ---
    console.log("DEBUG FINISH EXAM:", {
        sessionId,
        examId,
        userId
    });

    // 3. Validasi Ketat
    if (!sessionId || isNaN(sessionId)) {
      alert("Error: Session ID hilang. Silakan refresh halaman atau login ulang.");
      return;
    }
    if (!userId || isNaN(userId)) {
      alert("Error: User ID tidak valid. Silakan login ulang.");
      return;
    }
    if (!examId || isNaN(examId)) {
      alert("Error: Exam ID tidak valid.");
      return;
    }

    // 4. Kirim Jawaban (Looping)
    // (Bagian ini biarkan seperti kode sebelumnya...)
    for (let i = 0; i < questions.value.length; i++) {
      const q = questions.value[i];
      const selectedId = answers.value[i];

      if (selectedId) {
        const selectedOption = q.options.find((opt) => opt.id === selectedId);
        if (selectedOption) {
          await submitUserAnswer({
            exam_session_id: sessionId,
            user_id: userId,
            question_id: q.id,
            answer: selectedOption.option_label || selectedOption.label, 
          });
        }
      }
    }

    // 5. FINISH SESSION (Payload yang Benar)
    const payload = {
      user_id: userId,  // Pastikan ini Number
      exam_id: examId   // Pastikan ini Number
    };

    console.log("Mengirim Payload ke Backend:", payload);

    await finishExamSession(payload);

    alert("Ujian selesai! Nilai Anda telah disimpan.");
    router.push("/ujian"); 

  } catch (err) {
    console.error("FinishExam Error Full:", err);
    
    // Tampilkan pesan error yang lebih jelas
    const backendMsg = err.response?.data?.message;
    if (backendMsg === "failed to find session: record not found") {
        alert("Gagal: Sesi ujian tidak ditemukan di server. Kemungkinan sesi sudah selesai atau belum dimulai dengan benar.");
    } else {
        alert(`Gagal menyelesaikan ujian: ${backendMsg || "Terjadi kesalahan server."}`);
    }
  }
};

</script>