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

    <div class="grid grid-cols-1 md:grid-cols-3 gap-4 mb-6">
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

    <div class="bg-gray-50 rounded-lg p-4 border overflow-x-auto">
      <h3 class="font-semibold mb-3">Daftar Ujian</h3>

      <table class="w-full min-w-[600px]">
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
            <td class="p-2">
              <div class="font-medium">{{ exam.title }}</div>
              <div class="text-xs text-gray-500 line-clamp-1">{{ exam.description }}</div>
            </td>
            <td class="p-2">
              <span :class="statusBadgeClass(exam.status)">
                {{ statusText(exam.status) }}
              </span>
            </td>
            <td class="p-2 flex items-center justify-center gap-3">
              <router-link
                :to="{ name: isAdminRoute ? 'AdminExamDetail' : 'DosenExamDetail', params: { id: exam.id } }"
                class="text-blue-600 hover:underline text-sm"
              >
                👁️ View
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
            <td colspan="3" class="text-center p-8 text-gray-500">
              <div class="flex flex-col items-center">
                <i class="fas fa-inbox text-4xl mb-2 text-gray-300"></i>
                <p>Tidak ada ujian ditemukan</p>
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
// Pastikan path import ini sesuai dengan struktur folder Anda
import { getAllExam, deleteExam } from "../../provider/exam.provider";
import { useGetCurrentUser } from "../../hooks/useGetCurrentUser";

const route = useRoute();
const isAdminRoute = computed(() => route.path.startsWith('/admin'));

const exams = ref([]);
const search = ref("");
const sortBy = ref("Last Modified");
const statusFilter = ref("");
const page = ref(1);
const limit = 10;
const totalItems = ref(0);

const { user } = useGetCurrentUser();

// --- LOGIC LOAD DATA (DIPERBAIKI) ---
const loadExams = async () => {
  if (!user.value) return;

  try {
    const offset = (page.value - 1) * limit;
    const creatorId = isAdminRoute.value ? null : user.value.id;
    
    // Panggil Provider
    const result = await getAllExam(limit, offset, creatorId);

    // FIX: Cek apakah result memiliki properti .data (Objek) atau result itu sendiri adalah Array
    if (result && Array.isArray(result.data)) {
        // Format Baru: { data: [...], total: ... }
        exams.value = result.data;
        totalItems.value = result.total || 0;
    } else if (Array.isArray(result)) {
        // Format Lama (Fallback): [...]
        exams.value = result;
        totalItems.value = result.length;
    } else {
        // Data kosong atau format salah
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

// Watcher: Reload saat page berubah atau user baru login
watch([page, user], loadExams, { immediate: true });

// --- FILTERING DI SISI CLIENT (Hanya untuk data halaman ini) ---
const filteredExams = computed(() => {
  let data = [...exams.value];

  // 1. Search
  if (search.value) {
    data = data.filter((e) =>
      e.title.toLowerCase().includes(search.value.toLowerCase())
    );
  }

  // 2. Filter Status
  if (statusFilter.value) {
    data = data.filter((e) => e.status === statusFilter.value);
  }

  // 3. Sorting
  if (sortBy.value === "A-Z") {
    data.sort((a, b) => a.title.localeCompare(b.title));
  } else if (sortBy.value === "Z-A") {
    data.sort((a, b) => b.title.localeCompare(a.title));
  }
  // 'Last Modified' biasanya default dari server (created_at desc), jadi tidak perlu sort manual di sini

  return data;
});

// --- LOGIC PAGINATION (DIPERBAIKI) ---

// Total halaman dihitung dari totalItems server, bukan filteredExams
const totalPages = computed(() => {
  if (totalItems.value === 0) return 1;
  return Math.ceil(totalItems.value / limit);
});

// FIX: Jangan slice lagi! Data dari server sudah terpotong (paginated)
// Kita langsung tampilkan hasil filter dari data yang ada.
const paginatedExams = computed(() => {
  return filteredExams.value; 
});

const nextPage = () => {
  if (page.value < totalPages.value) {
    page.value++;
    // Watcher 'page' akan otomatis memanggil loadExams()
  }
};

const prevPage = () => {
  if (page.value > 1) {
    page.value--;
    // Watcher 'page' akan otomatis memanggil loadExams()
  }
};

// --- UTILS ---
const statusText = (status) => {
  if (status === "not_started") return "Not Started";
  if (status === "running") return "Ongoing";
  if (status === "finished") return "Finished";
  return status || "-";
};

const statusBadgeClass = (status) => {
  const base = "px-2 py-1 rounded text-xs font-semibold";
  if (status === "not_started") return `${base} bg-gray-100 text-gray-600`;
  if (status === "running") return `${base} bg-blue-100 text-blue-600`;
  if (status === "finished") return `${base} bg-green-100 text-green-600`;
  return `${base} bg-gray-100 text-gray-600`;
};

const removeExam = async (id) => {
  if (!confirm("Apakah kamu yakin ingin menghapus ujian ini?")) return;
  try {
    await deleteExam(id);
    alert("✅ Ujian berhasil dihapus!");
    // Jika halaman saat ini kosong setelah hapus, mundur 1 halaman
    if (exams.value.length === 1 && page.value > 1) {
      page.value--;
    } else {
      loadExams();
    }
  } catch (error) {
    console.error("Gagal menghapus ujian:", error);
    alert("❌ Gagal menghapus ujian. Coba lagi nanti.");
  }
};
</script>