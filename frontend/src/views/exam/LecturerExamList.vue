<template>
  <div class="bg-white rounded-lg shadow-md p-6">
    <div class="flex items-center justify-between mb-6">
      <h2 class="text-xl font-bold">Manajemen Ujian</h2>
      <router-link
        to="/dosen/exam/create"
        class="bg-blue-600 hover:bg-blue-700 text-white font-medium px-4 py-2 rounded-md"
      >
        + Buat Ujian
      </router-link>
    </div>

    <!-- Filters -->
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
          placeholder="Search your file"
          class="w-full p-2 outline-none"
        />
      </div>
    </div>

    <!-- Table -->
    <div class="bg-gray-50 rounded-lg p-4 border">
      <h3 class="font-semibold mb-3">Daftar Ujian</h3>

      <table class="w-full">
        <thead>
          <tr class="text-left border-b text-sm font-bold">
            <th class="p-2">Nama Ujian</th>
            <th class="p-2">Status</th>
            <th class="p-2">Actions</th>
          </tr>
        </thead>

        <tbody>
          <tr
            v-for="exam in filteredExams"
            :key="exam.id"
            class="border-b hover:bg-gray-100 transition"
          >
            <td class="p-2">{{ exam.title }}</td>
            <td class="p-2">{{ statusText(exam.status) }}</td>
            <td class="p-2 flex items-center gap-3">
              <button class="text-blue-600 text-sm">👁️ View Details</button>
              <button class="text-gray-500 text-sm">✏️</button>
              <button @click="removeExam(exam.id)" class="text-red-600 text-sm">
                🗑️
              </button>
            </td>
          </tr>

          <tr v-if="filteredExams.length === 0">
            <td colspan="3" class="text-center p-4 text-gray-500">
              Tidak ada ujian
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from "vue";
import { getAllExam, deleteExam } from "@/provider/exam.provider";
import { useGetCurrentUser } from "@/hooks/useGetCurrentUser";

const exams = ref([]);
const search = ref("");
const sortBy = ref("Last Modified");
const statusFilter = ref("");

const { user } = useGetCurrentUser();

const loadExams = async () => {
  const res = await getAllExam();
  // Kalau backend return array full, filter untuk dosen ini saja
  exams.value = res.data.filter((e) => e.creator_id === user.value.id);
};

onMounted(loadExams);

const filteredExams = computed(() => {
  let data = [...exams.value];

  if (search.value)
    data = data.filter((e) =>
      e.title.toLowerCase().includes(search.value.toLowerCase())
    );

  if (statusFilter.value)
    data = data.filter((e) => e.status === statusFilter.value);

  return data;
});

const statusText = (status) => {
  if (status === "not_started") return "Not Started";
  if (status === "running") return "Ongoing";
  if (status === "finished") return "Finished";
  return "-";
};

const removeExam = async (id) => {
  if (!confirm("Hapus ujian ini?")) return;
  await deleteExam(id);
  loadExams();
};
</script>
