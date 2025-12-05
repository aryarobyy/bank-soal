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
            class="w-full border rounded-lg p-2 bg-gray-50 focus:ring-2 focus:ring-blue-500 outline-none"
            required
          />
        </div>

        <div>
          <label class="font-semibold text-gray-700">Deskripsi</label>
          <textarea
            v-model="form.description"
            rows="3"
            class="w-full border rounded-lg p-2 bg-gray-50 focus:ring-2 focus:ring-blue-500 outline-none"
          ></textarea>
        </div>

        <div>
          <label class="font-semibold text-gray-700">Kesulitan</label>
          <select
            v-model="form.difficulty"
            class="w-full border rounded-lg p-2 bg-gray-50 focus:ring-2 focus:ring-blue-500 outline-none"
            required
          >
            <option value="easy">Mudah</option>
            <option value="medium">Sedang</option>
            <option value="hard">Sulit</option>
          </select>
        </div>

        <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
          <div>
            <label class="font-semibold text-gray-700">Tanggal Mulai</label>
            <input
              v-model="form.started_at"
              type="datetime-local"
              class="w-full border rounded-lg p-2 bg-gray-50 focus:ring-2 focus:ring-blue-500 outline-none"
              required
            />
          </div>

          <div>
            <label class="font-semibold text-gray-700">Tanggal Selesai</label>
            <input
              v-model="form.finished_at"
              type="datetime-local"
              class="w-full border rounded-lg p-2 bg-gray-50 focus:ring-2 focus:ring-blue-500 outline-none"
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
            class="w-full border rounded-lg p-2 bg-gray-50 focus:ring-2 focus:ring-blue-500 outline-none"
            required
          />
        </div>

        <div class="pt-4 flex flex-col gap-3">
          <button
            type="submit"
            class="w-full bg-blue-600 hover:bg-blue-700 text-white font-semibold rounded-lg py-2 transition disabled:opacity-50"
            :disabled="saving"
          >
            {{ saving ? "Menyimpan..." : "Simpan Perubahan" }}
          </button>

          <button
            type="button"
            @click="handleCancel"
            class="w-full text-center text-gray-600 hover:text-gray-900 hover:underline transition"
          >
            Batal & Kembali
          </button>
        </div>
      </form>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, computed } from "vue";
import { useRoute, useRouter } from "vue-router";
import { getExamById, updateExam } from "../../provider/exam.provider";

const route = useRoute();
const router = useRouter();
const isAdminRoute = computed(() => route.path.startsWith('/admin'));

const id = route.params.id;
const loading = ref(true);
const saving = ref(false);

const form = ref({
  title: "",
  description: "",
  difficulty: "easy",
  started_at: "",
  finished_at: "",
  long_time: 0,
});

const formatForInput = (dateStr) => {
  if (!dateStr) return "";
  try {
    const date = new Date(dateStr);
    const year = date.getFullYear();
    const month = String(date.getMonth() + 1).padStart(2, '0');
    const day = String(date.getDate()).padStart(2, '0');
    const hours = String(date.getHours()).padStart(2, '0');
    const minutes = String(date.getMinutes()).padStart(2, '0');
    
    
    return `${year}-${month}-${day}T${hours}:${minutes}`;
  } catch (e) {
    return "";
  }
};

onMounted(async () => {
  try {
    const res = await getExamById(id);
    const exam = res?.data || res;
    
    if (!exam) throw new Error("Data ujian tidak ditemukan!");

   
    form.value = {
      title: exam.title,
      description: exam.description || "",
      difficulty: exam.difficulty || "easy",
    
      started_at: formatForInput(exam.started_at),
      finished_at: formatForInput(exam.finished_at),
      long_time: exam.long_time || 0,
    };
  } catch (err) {
    console.error(err);
    alert("❌ Gagal memuat data ujian!");
    goBack();
  } finally {
    loading.value = false;
  }
});


const goToDetail = () => {


  const routeName = isAdminRoute.value ? 'AdminExamDetail' : 'DosenExamDetail';
  
  router.push({ 
    name: routeName, 
    params: { id: id } 
  });
};

const goBack = () => {
  const routeName = isAdminRoute.value ? 'AdminManageExam' : 'DosenManageExam';
  router.push({ name: routeName });
};

const handleCancel = () => {

  goToDetail();
};



const handleSubmit = async () => {
  saving.value = true;
  try {
    if (!form.value.started_at || !form.value.finished_at) {
      alert("Tanggal mulai dan selesai harus diisi!");
      saving.value = false;
      return;
    }

    const payload = {
      title: form.value.title,
      description: form.value.description,
      difficulty: form.value.difficulty,
      long_time: Number(form.value.long_time),
      started_at: new Date(form.value.started_at).toISOString(),
      finished_at: new Date(form.value.finished_at).toISOString(),
    };

    await updateExam(Number(id), payload);
    
    alert("✅ Ujian berhasil diperbarui!");
    
    goToDetail();

  } catch (err) {
    console.error("Gagal update ujian:", err);
    const msg = err.response?.data?.message || "Gagal menyimpan perubahan.";
    alert(`❌ ${msg}`);
  } finally {
    saving.value = false;
  }
};
</script>