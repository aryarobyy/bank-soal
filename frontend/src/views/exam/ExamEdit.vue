<template>
  <div class="flex justify-center items-center min-h-screen bg-gray-50">
    <div class="bg-white rounded-xl shadow-md p-8 w-full max-w-2xl">
      <h2 class="text-2xl font-bold text-primary mb-6">Edit Ujian</h2>

      <div v-if="loading" class="text-gray-500 text-center py-10">
        Memuat data ujian...
      </div>

      <form v-else @submit.prevent="handleSubmit" class="space-y-4">
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
          <label class="font-semibold text-gray-700">Kesulitan</label>
          <select
            v-model="form.difficulty"
            class="w-full border rounded-lg p-2 bg-gray-50"
            required
          >
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
            <label class="font-semibold text-gray-700">Tanggal Selesai</label>
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
          :disabled="saving"
        >
          {{ saving ? "Menyimpan..." : "Simpan Perubahan" }}
        </button>

        <router-link
          :to="{ name: isAdminRoute ? 'AdminManageExam' : 'DosenManageExam' }"
          class="block text-center text-gray-600 hover:underline mt-3"
        >
          🔙 Kembali ke Daftar Ujian
        </router-link>
      </form>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, computed } from "vue"; // Impor computed
import { useRoute, useRouter } from "vue-router";
// Pastikan path provider Anda sudah benar
import { getExamById, updateExam } from "/src/provider/exam.provider";

const route = useRoute();
const router = useRouter();
// Tambahkan computed isAdminRoute
const isAdminRoute = computed(() => route.path.startsWith('/admin'));

const form = ref({
  title: "",
  description: "",
  difficulty: "easy",
  started_at: "",
  finished_at: "",
  long_time: 0,
});

const loading = ref(true);
const saving = ref(false);
const id = route.params.id;

// Ambil data ujian berdasarkan ID dan isi form
onMounted(async () => {
  try {
    const res = await getExamById(id);
    const exam = res?.data || res;
    if (!exam) throw new Error("Data ujian tidak ditemukan!");

    // isi form
    form.value = {
      title: exam.title,
      description: exam.description || "",
      difficulty: exam.difficulty || "easy",
      started_at: exam.started_at ? exam.started_at.slice(0, 16) : "",
      finished_at: exam.finished_at ? exam.finished_at.slice(0, 16) : "",
      long_time: exam.long_time || 0,
    };
  } catch (err) {
    alert("❌ Gagal memuat data ujian!");
    console.error(err);
    // Navigasi dinamis
    router.push({ name: isAdminRoute.value ? 'AdminManageExam' : 'DosenManageExam' });
  } finally {
    loading.value = false;
  }
});

// Submit perubahan
const handleSubmit = async () => {
  saving.value = true;
  try {
    const payload = {
      ...form.value,
      long_time: Number(form.value.long_time),
      started_at: new Date(form.value.started_at).toISOString(),
      finished_at: new Date(form.value.finished_at).toISOString(),
    };

    await updateExam(Number(id), payload);
    alert("✅ Ujian berhasil diperbarui!");
    // Navigasi dinamis
    router.push({ name: isAdminRoute.value ? 'AdminManageExam' : 'DosenManageExam' });
  } catch (err) {
    console.error("Gagal update ujian:", err);
    alert("❌ Gagal menyimpan perubahan.");
  } finally {
    saving.value = false;
  }
};
</script>