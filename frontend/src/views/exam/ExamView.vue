<template>
  <div class="min-h-screen bg-[#e9edfc] font-sans flex flex-col">
    <Navbar />

    <main class="flex-grow flex justify-center items-center py-10 px-4">
      <div
        class="w-full max-w-4xl bg-white rounded-3xl shadow-lg p-8 sm:p-10 border border-gray-200"
      >
        <!-- Loading State -->
        <div v-if="loading" class="text-center text-gray-600 py-10">
          Memuat data ujian...
        </div>

        <!-- Error State -->
        <div v-else-if="error" class="text-center text-red-500 py-10">
          Gagal memuat data ujian 😢
        </div>

        <!-- Konten Ujian -->
        <div v-else>
          <!-- Header Ujian -->
          <div
            class="bg-[#f5f7ff] p-8 rounded-2xl text-center mb-8 border border-gray-100"
          >
            <h2 class="text-3xl font-bold text-gray-800 mb-2">
              {{ exam.title }}
            </h2>
            <p class="text-gray-600 text-sm sm:text-base">
              {{ formattedDate }}
            </p>
            <p class="font-medium text-gray-700 mt-2 text-sm sm:text-base">
              Durasi Ujian :
              <span class="font-semibold text-blue-700">
                {{ exam.long_time }} Menit
              </span>
            </p>
            <p class="font-medium text-gray-600 mt-2 text-sm sm:text-base">
              Kesulitan :
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
            <p class="text-gray-700 leading-relaxed text-sm sm:text-base">
              {{ exam.description }}
            </p>
          </div>

          <!-- Aturan Ujian -->
          <div class="px-2 sm:px-6">
            <h3 class="font-semibold text-gray-900 mb-4 text-lg">
              Aturan Ujian :
            </h3>
            <ul
              class="text-gray-700 space-y-3 text-sm sm:text-base leading-relaxed"
            >
              <li class="flex items-start">
                <span class="mr-2 font-bold text-blue-600">&gt;</span>
                Tidak diizinkan keluar dari ujian kecuali telah selesai / waktu
                ujian telah habis.
              </li>
              <li class="flex items-start">
                <span class="mr-2 font-bold text-blue-600">&gt;</span>
                Timer akan tetap berjalan meskipun koneksi internet Anda
                terputus.
              </li>
            </ul>
          </div>

          <!-- Tombol Mulai -->
          <div class="flex justify-end mt-8">
            <button
              @click="startExam"
              class="bg-[#2ecc71] hover:bg-[#27ae60] text-white text-sm sm:text-base px-5 py-2.5 rounded-md font-semibold transition-all shadow-sm hover:shadow-md"
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

const toastRef = ref(null);
const router = useRouter();
const route = useRoute();
const { user } = useGetCurrentUser();

// --- State
const exam = ref({});
const loading = ref(true);
const error = ref(false);

// --- Format waktu jadi lebih rapi
const formattedDate = computed(() => {
  if (!exam.value.started_at || !exam.value.finished_at) return "";
  const start = new Date(exam.value.started_at);
  const finish = new Date(exam.value.finished_at);

  return `${start.toLocaleDateString("id-ID", {
    day: "2-digit",
    month: "long",
    year: "numeric",
  })} | ${start.toLocaleTimeString("id-ID", {
    hour: "2-digit",
    minute: "2-digit",
  })} - ${finish.toLocaleTimeString("id-ID", {
    hour: "2-digit",
    minute: "2-digit",
  })} WIB`;
});

// --- Ambil data ujian berdasarkan ID
onMounted(async () => {
  if (!user) {
    toastRef.value.showToast(
      "error",
      "User tidak dikenali",
      "Anda harus login terlebih dahulu"
    );
    return router.push("/login");
  }

  try {
    const id = route.query.id || 1; // Default ke ID 1
    const data = await getExamById(id);
    exam.value = data;
  } catch (err) {
    console.error("Gagal ambil data ujian:", err);
    error.value = true;
  } finally {
    loading.value = false;
  }
});

const startExam = () => {
  router.push("/exam/start");
};
</script>