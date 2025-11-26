<template>
  <div class="min-h-screen bg-[#e9edfc] font-sans flex flex-col">
    <Navbar />

    <main class="flex-grow flex justify-center items-center py-10 px-4">
      <div
        class="w-full max-w-4xl bg-white rounded-3xl shadow-lg p-8 sm:p-10 border border-gray-200"
      >
        <!-- Loading -->
        <div v-if="loading" class="text-center text-gray-600 py-10">
          Memuat data ujian...
        </div>

        <!-- Error -->
        <div v-else-if="error" class="text-center text-red-500 py-10">
          {{ error }}
        </div>

        <!-- Data Ujian -->
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
          </div>

          <!-- Deskripsi -->
          <div class="px-2 sm:px-6 mb-6">
            <h3 class="font-semibold text-gray-900 mb-3 text-lg">
              Deskripsi :
            </h3>
            <p class="text-gray-700 leading-relaxed">
              {{ exam.description || "Tidak ada deskripsi." }}
            </p>
          </div>

          <!-- Aturan -->
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
            </ul>
          </div>

          <!-- Tombol -->
          <div class="flex justify-end mt-8">
            <button
              @click="startExam"
              class="bg-[#2ecc71] hover:bg-[#27ae60] text-white px-5 py-2.5 rounded-md font-semibold shadow-sm hover:shadow-md"
            >
              Mulai Ujian
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
import { ref, onMounted, computed } from "vue";
import { useRouter, useRoute } from "vue-router";
import { useGetCurrentUser } from "../../hooks/useGetCurrentUser";
import { getExamById } from "../../provider/exam.provider";
import { createExamSession } from "../../provider/examsession.provider";

const toastRef = ref(null);
const router = useRouter();
const route = useRoute();
const { user } = useGetCurrentUser();

const exam = ref(null);
const loading = ref(true);
const error = ref("");

// Format tanggal
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

// Ambil data ujian
onMounted(async () => {
  try {
    const id = route.query.id || 2; // Default ke ID 1
    const data = await getExamById(id);
    exam.value = data;
  } catch (err) {
    console.error("ERROR GET EXAM:", err);
    error.value = "Gagal mengambil data ujian!";
  } finally {
    loading.value = false;
  }
});

// Mulai ujian → halaman /exam/start?id=xx
const startExam = async () => {
  try {
    // 1️⃣ Buat session baru
    const session = await createExamSession(exam.value.id);

    // 2️⃣ Ambil sessionId dari BE
    const sessionId = session?.data?.id || session?.id;

    if (!sessionId) {
      alert("Gagal membuat sesi ujian!");
      return;
    }

    // 3️⃣ Kirim exam_id + session_id
    router.push(`/exam/start?id=${exam.value.id}&session_id=${sessionId}`);
  } catch (err) {
    console.error(err);
    alert("Gagal memulai ujian.");
  }
};
</script>
