<!-- src/views/exam/ExamDo.vue -->
<template>
  <div class="min-h-screen bg-[#e9edfc] flex items-center justify-center px-4">
    <div
      class="w-full max-w-5xl bg-white rounded-3xl shadow-lg p-6 sm:p-8 border border-gray-200"
    >
      <!-- Loading -->
      <div v-if="loading" class="text-center text-gray-500 py-10">
        Memuat ujian...
      </div>

      <!-- Error -->
      <div v-else-if="error" class="text-center text-red-500 py-10">
        {{ error }}
      </div>

      <!-- Tidak ada soal -->
      <div
        v-else-if="!questions.length"
        class="text-center text-gray-600 py-10"
      >
        Soal tidak tersedia.
      </div>

      <!-- Konten -->
      <div v-else>
        <!-- Header -->
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

        <!-- Kartu Soal -->
        <div class="bg-[#f5f7ff] rounded-2xl p-6 sm:p-8 mb-8">
          <p class="text-sm sm:text-base font-semibold text-gray-800 mb-4">
            Question: {{ currentQuestion?.question_text || "-" }}
          </p>

          <!-- Opsi Jawaban -->
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
              {{ opt.text || opt.option_text || opt.label || "NO_TEXT" }}
            </button>
          </div>
        </div>

        <!-- Navigasi Nomor Soal -->
        <div class="flex justify-center gap-2 mb-6">
          <button
            v-for="n in questions.length"
            :key="n"
            @click="goToQuestion(n)"
            type="button"
            :class="[
              'w-8 h-8 sm:w-9 sm:h-9 text-xs sm:text-sm rounded-full border transition',
              currentNo === n
                ? 'bg-blue-600 text-white border-blue-600'
                : 'bg-white text-gray-700 border-gray-300 hover:border-blue-400',
            ]"
          >
            {{ n }}
          </button>
        </div>

        <!-- Navigasi Sebelumnya & Selanjutnya -->
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
import { ref, computed, onMounted } from "vue";
import { useRoute, useRouter } from "vue-router";
import { getExamById } from "@/provider/exam.provider";
import { getExamQuestions } from "@/provider/examquestion.provider";
import { submitUserAnswer } from "@/provider/useranswer.provider"; // ✅ WAJIB ADA

// Router
const route = useRoute();
const router = useRouter();

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

// Simpan jawaban
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
      finishExam();
    }
  }, 1000);
};

const formattedTime = computed(() => {
  const m = Math.floor(timeLeft.value / 60);
  const s = timeLeft.value % 60;
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
    const examRes = await getExamById(examId);
    exam.value = examRes?.data || examRes;

    const qRes = await getExamQuestions(examId);
    questions.value = Array.isArray(qRes) ? qRes : qRes.data || [];

    // Setup jawaban kosong
    answers.value = Array(questions.value.length).fill(null);

    // 🔥 TIMER MENGGUNAKAN duration dari DB atau default 120 menit
    const duration = exam.value?.duration || 120;
    startTimer(duration);
  } catch (e) {
    console.error(e);
    error.value = "Terjadi kesalahan saat memuat ujian.";
  } finally {
    loading.value = false;
  }
});

// Navigasi
const goToQuestion = (n) => {
  currentNo.value = n;
};
const nextQuestion = () => {
  if (currentNo.value < questions.value.length) currentNo.value++;
};
const prevQuestion = () => {
  if (currentNo.value > 1) currentNo.value--;
};

// ===================== FINISH EXAM =====================
const finishExam = async () => {
  try {
    const sessionId = Number(route.query.session_id);
    const userId = Number(localStorage.getItem("user_id"));

    console.log("SESSION ID:", sessionId);
    console.log("USER ID:", userId);

    if (!sessionId || isNaN(sessionId)) {
      alert("Session ID tidak valid.");
      return;
    }
    if (!userId || isNaN(userId)) {
      alert("User ID tidak valid.");
      return;
    }

    for (let i = 0; i < questions.value.length; i++) {
      const q = questions.value[i];
      const selectedId = answers.value[i];

      const selectedOption = q.options.find((opt) => opt.id === selectedId);

      if (!selectedOption) {
        console.warn("Tidak ada opsi yang dipilih.");
        continue;
      }

      await submitUserAnswer({
        exam_session_id: sessionId,
        user_id: userId,
        question_id: q.id,

        // 🔥 BACKEND butuh ini
        answer: selectedOption.option_label || selectedOption.label,

        is_correct: false,
      });
    }

    alert("Jawaban berhasil dikirim & ujian selesai!");
    router.push("/ujian");
  } catch (err) {
    console.error("FinishExam Error:", err);
    alert("Gagal mengirim jawaban.");
  }
};
</script>
