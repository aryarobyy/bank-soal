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

    <div class="grid grid-cols-1 md:grid-cols-2 gap-4 mb-6">
      <div class="flex items-center gap-2">
        <span class="text-sm text-gray-600 font-medium">Urutkan:</span>
        <select v-model="sortBy" class="border rounded-lg p-2 flex-1 bg-white">
          <option value="Last Modified">Terbaru (Last Modified)</option>
          <option value="A-Z">Judul (A-Z)</option>
          <option value="Z-A">Judul (Z-A)</option>
        </select>
      </div>

      <div class="flex items-center border rounded-lg px-3 bg-white">
        <i class="fas fa-search text-gray-400 mr-2"></i>
        <input
          type="text"
          v-model="search"
          placeholder="Cari judul ujian..."
          class="w-full p-2 outline-none"
        />
      </div>
    </div>

    <div class="bg-gray-50 rounded-lg p-4 border overflow-x-auto">
      <h3 class="font-semibold mb-3">Daftar Ujian</h3>

      <table class="w-full min-w-[600px]">
        <thead>
          <tr class="text-left border-b text-sm font-bold bg-gray-100 text-gray-700">
            <th class="p-3">Nama Ujian & Deskripsi</th>
            <th class="p-3">Durasi</th>
            <th class="p-3">Waktu Mulai</th>
            <th class="p-3 text-center">Aksi</th>
          </tr>
        </thead>

        <tbody class="text-sm">
          <tr
            v-for="exam in paginatedExams"
            :key="exam.id"
            class="border-b hover:bg-white transition"
          >
            <td class="p-3">
              <div class="font-bold text-gray-800">{{ exam.title }}</div>
              <div class="text-xs text-gray-500 mt-1 line-clamp-1">{{ exam.description || 'Tidak ada deskripsi' }}</div>
            </td>
            <td class="p-3">
              <span class="bg-blue-50 text-blue-700 px-2 py-1 rounded text-xs font-semibold">
                {{ exam.long_time }} Menit
              </span>
            </td>
            <td class="p-3 text-gray-600">
              {{ formatDate(exam.started_at) }}
            </td>
            <td class="p-3 flex items-center justify-center gap-3">
              <router-link
                :to="{ name: isAdminRoute ? 'AdminExamDetail' : 'DosenExamDetail', params: { id: exam.id } }"
                class="text-blue-600 hover:text-blue-800 hover:underline flex items-center gap-1"
              >
                <i class="fas fa-eye"></i> Detail
              </router-link>

              <button
                @click="removeExam(exam.id)"
                class="text-red-600 hover:text-red-800 flex items-center gap-1"
              >
                <i class="fas fa-trash"></i> Hapus
              </button>
            </td>
          </tr>

          <tr v-if="paginatedExams.length === 0">
            <td colspan="4" class="text-center p-8 text-gray-500">
              <div class="flex flex-col items-center">
                <i class="fas fa-inbox text-4xl mb-2 text-gray-300"></i>
                <p>Tidak ada ujian ditemukan.</p>
              </div>
            </td>
          </tr>
        </tbody>
      </table>

      <div v-if="totalPages > 1" class="flex justify-center items-center mt-6 gap-3">
        <button
          @click="prevPage"
          :disabled="page === 1"
          class="px-3 py-1 bg-white border rounded hover:bg-gray-50 disabled:opacity-50 disabled:cursor-not-allowed"
        >
          Prev
        </button>

        <span class="text-sm text-gray-600">
          Halaman <span class="font-bold">{{ page }}</span> dari {{ totalPages }}
        </span>

        <button
          @click="nextPage"
          :disabled="page === totalPages"
          class="px-3 py-1 bg-white border rounded hover:bg-gray-50 disabled:opacity-50 disabled:cursor-not-allowed"
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
import { getAllExam, deleteExam } from "../../provider/exam.provider";
import { useGetCurrentUser } from "../../hooks/useGetCurrentUser";

const route = useRoute();
const isAdminRoute = computed(() => route.path.startsWith('/admin'));

const exams = ref([]);
const search = ref("");
const sortBy = ref("Last Modified");
const page = ref(1);
const limit = 10;
const totalItems = ref(0);

const { user } = useGetCurrentUser();

// --- LOAD DATA ---
const loadExams = async () => {
  if (!user.value) return;

  try {
    const offset = (page.value - 1) * limit;
    const creatorId = isAdminRoute.value ? null : user.value.id;
    
    // Panggil Provider
    const result = await getAllExam(limit, offset, creatorId);

    if (result && Array.isArray(result.data)) {
        exams.value = result.data;
        totalItems.value = result.total || 0;
    } else if (Array.isArray(result)) {
        exams.value = result;
        totalItems.value = result.length;
    } else {
        exams.value = [];
        totalItems.value = 0;
    }

  } catch (err) {
    console.error("Gagal memuat data ujian:", err);
    exams.value = [];
    totalItems.value = 0;
  }
};

onMounted(loadExams);


watch([page, user], loadExams, { immediate: true });


const filteredExams = computed(() => {
  let data = [...exams.value];

  // 1. Search
  if (search.value) {
    data = data.filter((e) =>
      e.title.toLowerCase().includes(search.value.toLowerCase())
    );
  }

  
  if (sortBy.value === "A-Z") {
    data.sort((a, b) => a.title.localeCompare(b.title));
  } else if (sortBy.value === "Z-A") {
    data.sort((a, b) => b.title.localeCompare(a.title));
  } else if (sortBy.value === "Last Modified") {
    
    data.sort((a, b) => {
      const dateA = new Date(a.updated_at || a.created_at).getTime();
      const dateB = new Date(b.updated_at || b.created_at).getTime();
      return dateB - dateA; 
    });
  }

  return data;
});

// --- PAGINATION ---
const totalPages = computed(() => {
  if (totalItems.value === 0) return 1;
  return Math.ceil(totalItems.value / limit);
});

const paginatedExams = computed(() => {

  return filteredExams.value; 
});

const nextPage = () => {
  if (page.value < totalPages.value) {
    page.value++;
  }
};

const prevPage = () => {
  if (page.value > 1) {
    page.value--;
  }
};

// --- UTILS ---
const formatDate = (dateString) => {
  if (!dateString) return "-";
  return new Date(dateString).toLocaleDateString("id-ID", {
    day: "numeric", month: "short", year: "numeric",
    hour: "2-digit", minute: "2-digit"
  });
};

const removeExam = async (id) => {
  if (!confirm("Apakah kamu yakin ingin menghapus ujian ini? Data soal dan nilai terkait juga akan terhapus.")) return;
  
  try {

    await deleteExam(Number(id));
    
    alert("✅ Ujian berhasil dihapus!");
    
    // Jika halaman saat ini kosong setelah hapus, mundur 1 halaman
    if (exams.value.length === 1 && page.value > 1) {
      page.value--;
    } else {
      loadExams();
    }
  } catch (error) {
    console.error("Gagal menghapus ujian:", error);
    const msg = error.response?.data?.message || "Gagal menghapus ujian.";
    alert(`❌ ${msg}`);
  }
};
</script>