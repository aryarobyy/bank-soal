<template>
  <div class="bg-white rounded-lg shadow-md p-6">
    <div class="flex items-center justify-between mb-6">
      <h2 class="text-xl font-bold">Manajemen Ujian</h2>
      <router-link
        :to="{ name: isAdminRoute ? 'AdminCreateExam' : 'DosenCreateExam' }"
        class="bg-blue-600 hover:bg-blue-700 text-white font-medium px-4 py-2 rounded-md"
      >
        + Buat Ujian
      </router-link>
    </div>

    <div class="grid grid-cols-3 gap-4 mb-6">
      <select v-model="sortBy" class="border rounded-lg p-2">
        <option>Last Modified</option>
        <option>A-Z</option>
        <option>Z-A</option>
      </select>

      <select v-model="statusFilter" class="border rounded-lg p-2">
        <option value="">Filter by Status</option>
        <option value="not_started">Not Started</option>
        <option value="running">Running</option>
        <option value="finished">Finished</option>
      </select>

      <div class="flex items-center border rounded-lg px-3">
        <input
          type="text"
          v-model="search"
          placeholder="Search your exam"
          class="w-full p-2 outline-none"
        />
      </div>
    </div>

    <div class="bg-gray-50 rounded-lg p-4 border">
      <h3 class="font-semibold mb-3">Daftar Ujian</h3>

      <table class="w-full">
        <thead>
          <tr class="text-left border-b text-sm font-bold">
            <th class="p-2">Nama Ujian</th>
            <th class="p-2">Status</th>
            <th class="p-2 text-center">Actions</th>
          </tr>
        </thead>

        <tbody>
          <tr
            v-for="exam in paginatedExams"
            :key="exam.id"
            class="border-b hover:bg-gray-100 transition"
          >
            <td class="p-2">{{ exam.title }}</td>
            <td class="p-2">{{ statusText(exam.status) }}</td>
            <td class="p-2 flex items-center justify-center gap-3">
              <router-link
                :to="{ name: isAdminRoute ? 'AdminExamDetail' : 'DosenExamDetail', params: { id: exam.id } }"
                class="text-blue-600 hover:underline text-sm"
              >
                👁️ View Details
              </router-link>

              <button
                @click="removeExam(exam.id)"
                class="text-red-600 hover:text-red-800 text-sm"
              >
                🗑️ Delete
              </button>
            </td>
          </tr>

          <tr v-if="paginatedExams.length === 0">
            <td colspan="3" class="text-center p-4 text-gray-500">
              Tidak ada ujian
            </td>
          </tr>
        </tbody>
      </table>

      <div class="flex justify-center items-center mt-6 gap-3">
        <button
          @click="prevPage"
          :disabled="page === 1"
          class="px-3 py-1 bg-gray-200 rounded disabled:opacity-50"
        >
          Prev
        </button>

        <span>Halaman {{ page }} dari {{ totalPages }}</span>

        <button
          @click="nextPage"
          :disabled="page === totalPages"
          class="px-3 py-1 bg-gray-200 rounded disabled:opacity-50"
        >
          Next
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted, watch } from "vue";
import { useRoute } from "vue-router";
import { getAllExam, deleteExam } from "/src/provider/exam.provider";
import { useGetCurrentUser } from "/src/hooks/useGetCurrentUser";

const route = useRoute();
const isAdminRoute = computed(() => route.path.startsWith('/admin'));

const exams = ref([]);
const search = ref("");
const sortBy = ref("Last Modified");
const statusFilter = ref("");
const page = ref(1);
const limit = 10;
const totalItems = ref(0); // Tambahkan totalItems

const { user } = useGetCurrentUser();

// ## FUNGSI DIPERBARUI ##
const loadExams = async () => {
  if (!user.value) return; // Jangan lakukan apa-apa jika user belum dimuat

  try {
    const offset = (page.value - 1) * limit;
    
    // Tentukan apakah kita perlu memfilter berdasarkan creator
    const creatorId = isAdminRoute.value ? null : user.value.id;
    
    // Panggil provider baru dengan creatorId
    // 'result' sekarang adalah { data: [...], total: ... }
    // ATAU provider Anda mengembalikan array (mari kita asumsikan array)
    const result = await getAllExam(limit, offset, creatorId);

    // Provider Anda mengembalikan array, bukan objek { data, total }
    exams.value = Array.isArray(result) ? result : [];
    
    // (Paginasi sederhana berbasis data yang diambil, BUKAN total)
    totalItems.value = exams.value.length; 

  } catch (err) {
    console.error("Gagal memuat data ujian:", err);
    exams.value = [];
  }
};
// ## AKHIR PERUBAHAN ##

onMounted(loadExams);

// PERBAIKAN: Tonton juga 'user' untuk memastikan data dimuat setelah user ada
watch([page, user], loadExams, { immediate: true });

const filteredExams = computed(() => {
  let data = [...exams.value];
  if (search.value)
    data = data.filter((e) =>
      e.title.toLowerCase().includes(search.value.toLowerCase())
    );
  if (statusFilter.value)
    data = data.filter((e) => e.status === statusFilter.value);
  if (sortBy.value === "A-Z")
    data.sort((a, b) => a.title.localeCompare(b.title));
  else if (sortBy.value === "Z-A")
    data.sort((a, b) => b.title.localeCompare(a.title));
  return data;
});

// Paginasi ini masih berbasis data yang di-filter, bukan total dari DB
const totalPages = computed(() => {
  const total = filteredExams.value.length;
  if (total === 0) return 1;
  return Math.ceil(total / limit);
});


const paginatedExams = computed(() => {
  const start = (page.value - 1) * limit;
  const end = start + limit;
  return filteredExams.value.slice(start, end);
});

const nextPage = () => {
  if (page.value < totalPages.value) page.value++;
};

const prevPage = () => {
  if (page.value > 1) page.value--;
};

const statusText = (status) => {
  if (status === "not_started") return "Not Started";
  if (status === "running") return "Ongoing";
  if (status === "finished") return "Finished";
  return "-";
};

const removeExam = async (id) => {
  if (!confirm("Apakah kamu yakin ingin menghapus ujian ini?")) return;
  try {
    await deleteExam(id);
    alert("✅ Ujian berhasil dihapus!");
    loadExams();
  } catch (error) {
    console.error("Gagal menghapus ujian:", error);
    alert("❌ Gagal menghapus ujian. Coba lagi nanti.");
  }
};
</script>