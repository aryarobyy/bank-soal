<template>
  <div class="flex justify-center items-center min-h-screen bg-gray-50">
    <div class="bg-white rounded-xl shadow-md p-8 w-full max-w-2xl">
      <h2 class="text-2xl font-bold text-primary mb-6">Buat Ujian Baru</h2>

      <form @submit.prevent="handleSubmit" class="space-y-4">
        <div>
          <label class="font-semibold text-gray-700">Nama Ujian</label>
          <input
            v-model="form.title"
            type="text"
            class="w-full border rounded-lg p-2 bg-gray-50"
            required
          />
        </div>

        <div>
          <label class="font-semibold text-gray-700">Deskripsi</label>
          <textarea
            v-model="form.description"
            rows="3"
            class="w-full border rounded-lg p-2 bg-gray-50"
          ></textarea>
        </div>

        <div>
          <label class="font-semibold text-gray-700">Tingkat Kesulitan</label>
          <select
            v-model="form.difficulty"
            class="w-full border rounded-lg p-2 bg-gray-50"
            required
          >
            <option disabled value="">Pilih kesulitan</option>
            <option value="easy">Mudah</option>
            <option value="medium">Sedang</option>
            <option value="hard">Sulit</option>
          </select>
        </div>

        <div class="grid grid-cols-2 gap-4">
          <div>
            <label class="font-semibold text-gray-700">Tanggal Mulai</label>
            <input
              v-model="form.started_at"
              type="datetime-local"
              class="w-full border rounded-lg p-2 bg-gray-50"
              required
            />
          </div>

          <div>
            <label class="font-semibold text-gray-700">Tanggal Berakhir</label>
            <input
              v-model="form.finished_at"
              type="datetime-local"
              class="w-full border rounded-lg p-2 bg-gray-50"
              required
            />
          </div>
        </div>

        <div>
          <label class="font-semibold text-gray-700">Durasi (menit)</label>
          <input
            v-model.number="form.long_time"
            type="number"
            min="1"
            class="w-full border rounded-lg p-2 bg-gray-50"
            required
          />
        </div>

        <button
          type="submit"
          class="w-full bg-blue-600 hover:bg-blue-700 text-white font-semibold rounded-lg py-2 transition"
          :disabled="loading"
        >
          {{ loading ? "Menyimpan..." : "Simpan Ujian" }}
        </button>
      </form>

      <div
        v-if="showPopup"
        class="mt-6 text-center bg-green-100 text-green-800 border border-green-300 rounded-lg p-3"
      >
        {{ popupMessage }}
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed } from "vue"; // Impor computed
import { useRoute, useRouter } from "vue-router"; // Impor useRoute
// Pastikan path provider Anda sudah benar
import { createExam } from "../../provider/exam.provider.js";
import { useGetCurrentUser } from "../../hooks/useGetCurrentUser.js";

const router = useRouter();
const route = useRoute(); // Tambahkan route
const { user } = useGetCurrentUser();

// Tambahkan computed isAdminRoute
const isAdminRoute = computed(() => route.path.startsWith('/admin'));

const form = ref({
  title: "",
  description: "",
  difficulty: "",
  started_at: "",
  finished_at: "",
  long_time: "",
});

const loading = ref(false);
const showPopup = ref(false);
const popupMessage = ref("");

const handleSubmit = async () => {
  if (!user?.value) {
    alert("User belum login!");
    return;
  }

  loading.value = true;

  const payload = {
    title: form.value.title.trim(),
    description: form.value.description.trim(),
    difficulty: form.value.difficulty.toLowerCase(),
    started_at: new Date(form.value.started_at).toISOString(),
    finished_at: new Date(form.value.finished_at).toISOString(),
    long_time: Number(form.value.long_time),
    creator_id: user.value.id,
  };

  try {
    await createExam(payload);
    popupMessage.value = "✅ Ujian berhasil dibuat!";
    showPopup.value = true;

    // Navigasi dinamis
    const returnRouteName = isAdminRoute.value ? 'AdminManageExam' : 'DosenManageExam';
    
    setTimeout(() => {
      router.push({ name: returnRouteName });
    }, 800);
  } catch (err) {
    console.error(err);
    // Coba akses pesan error yang lebih spesifik jika ada
    const errorMsg = err.response?.data?.message || "Terjadi kesalahan";
    popupMessage.value = `❌ Gagal membuat ujian. ${errorMsg} `;
    showPopup.value = true;
  } finally {
    loading.value = false;
  }
};
</script>