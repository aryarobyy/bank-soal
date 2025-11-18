<template>
  <div class="p-6">
    <div v-if="loading" class="text-center py-10 text-gray-500">
      Memuat ujian...
    </div>

    <div v-else>
      <h2 class="text-2xl font-bold mb-4">{{ exam?.title }}</h2>

      <!-- Soal -->
      <div class="p-4 bg-white shadow rounded mb-6">
        <p class="font-semibold mb-3">Soal {{ currentNo }}:</p>
        <p class="text-gray-700">{{ currentQuestion?.question }}</p>

        <div class="mt-4 space-y-2">
          <label
            v-for="opt in currentQuestion?.options"
            :key="opt.id"
            class="flex items-center gap-2"
          >
            <input type="radio" :value="opt.id" v-model="selectedOption" />
            {{ opt.text }}
          </label>
        </div>
      </div>

      <!-- Navigasi -->
      <div class="flex justify-between">
        <button
          class="px-4 py-2 bg-gray-300 rounded"
          @click="prevQuestion"
          :disabled="currentNo === 1"
        >
          Sebelumnya
        </button>

        <button
          class="px-4 py-2 bg-blue-600 text-white rounded"
          @click="nextQuestion"
          v-if="currentNo < questions.length"
        >
          Selanjutnya
        </button>

        <button
          class="px-4 py-2 bg-green-600 text-white rounded"
          @click="finishExam"
          v-else
        >
          Selesaikan Ujian
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from "vue";
import { useRoute, useRouter } from "vue-router";
import {
  getExamSessionById,
  updateCurrentNo,
  finishExamSession,
} from "@/provider/examsession.provider";
import { getExamById } from "@/provider/exam.provider";
import { getExamQuestions } from "@/provider/examquestion.provider";

const route = useRoute();
const router = useRouter();

const loading = ref(true);
const session = ref(null);
const exam = ref(null);
const questions = ref([]);
const currentNo = ref(1);
const selectedOption = ref(null);

const currentQuestion = computed(() => questions.value[currentNo.value - 1]);

onMounted(async () => {
  try {
    const sessionId = route.params.id;

    // Ambil session
    const sessionRes = await getExamSessionById(sessionId);
    session.value = sessionRes.data;
    currentNo.value = session.value.current_no;

    // Ambil exam
    const examRes = await getExamById(session.value.exam_id);
    exam.value = examRes.data || examRes;

    // Ambil soal
    const qRes = await getExamQuestions(session.value.exam_id);
    questions.value = qRes.data;
  } catch (err) {
    console.error(err);
    alert("Gagal memuat ujian");
  } finally {
    loading.value = false;
  }
});

const nextQuestion = async () => {
  currentNo.value++;
  await updateCurrentNo(session.value.id, currentNo.value);
};

const prevQuestion = async () => {
  currentNo.value--;
  await updateCurrentNo(session.value.id, currentNo.value);
};

const finishExam = async () => {
  await finishExamSession();
  alert("Ujian selesai!");
  router.push("/ujian");
};
</script>
