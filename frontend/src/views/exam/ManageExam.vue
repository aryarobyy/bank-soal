<template>
  <div class="p-6 bg-[#f6f8ff] min-h-screen">
    <h2 class="text-2xl font-bold text-gray-800 mb-6">Manajemen Ujian</h2>

    <div class="flex flex-wrap gap-3 mb-6 items-center justify-between">
      <div class="flex gap-3">
        <select
          v-model="sortBy"
          class="border rounded-lg px-3 py-2 bg-white text-gray-700 shadow-sm"
        >
          <option value="modified">Last Modified</option>
          <option value="created">Created Date</option>
        </select>
        <select
          v-model="levelFilter"
          class="border rounded-lg px-3 py-2 bg-white text-gray-700 shadow-sm"
        >
          <option value="all">All Levels</option>
          <option value="easy">Easy</option>
          <option value="medium">Medium</option>
          <option value="hard">Hard</option>
        </select>
      </div>

      <div class="flex items-center gap-3">
        <input
          type="text"
          v-model="searchQuery"
          placeholder="Search your test..."
          class="border rounded-lg px-3 py-2 bg-white text-gray-700 shadow-sm"
        />
        <button
          @click="goToCreateExam"
          class="bg-green-500 hover:bg-green-600 text-white font-semibold px-4 py-2 rounded-lg shadow"
        >
          + Create Test
        </button>
      </div>
    </div>

    <div
      v-for="exam in filteredExams"
      :key="exam.id"
      class="bg-white rounded-xl shadow-md p-5 mb-4"
    >
      <div
        class="flex justify-between items-center cursor-pointer"
        @click="toggleExam(exam.id)"
      >
        <h3 class="font-semibold text-lg text-gray-800">{{ exam.title }}</h3>
        <button>
          <span v-if="expandedExam === exam.id">▲</span>
          <span v-else>▼</span>
        </button>
      </div>

      <transition name="fade">
        <div v-if="expandedExam === exam.id" class="mt-4 border-t pt-4">
          </div>
      </transition>
    </div>
  </div>
</template>

<script setup>
import { ref, computed } from "vue";
import { useRouter } from "vue-router";

const router = useRouter();
const expandedExam = ref(null);
const sortBy = ref("modified");
const levelFilter = ref("all");
const searchQuery = ref("");

const exams = ref([
  {
    id: 1,
    title: "Ujian 1",
    description: "Create a User Persona Based on Research Data.",
    status: "Not Started",
    level: "easy",
  },
  {
    id: 2,
    title: "Ujian 2",
    description: "Analyze System Requirements for Case Study.",
    status: "Ongoing",
    level: "medium",
  },
  {
    id: 3,
    title: "Ujian 3",
    description: "Develop a Mini Project with Vue.js.",
    status: "Completed",
    level: "hard",
  },
]);

// ## PERBAIKAN UTAMA ADA DI SINI ##
const goToCreateExam = () => {
  // Menggunakan named route agar tetap di dalam Admin Layout
  router.push({ name: "AdminCreateExam" });
};

const toggleExam = (id) => {
  expandedExam.value = expandedExam.value === id ? null : id;
};

const filteredExams = computed(() => {
  return exams.value.filter(
    (exam) =>
      (levelFilter.value === "all" || exam.level === levelFilter.value) &&
      exam.title.toLowerCase().includes(searchQuery.value.toLowerCase())
  );
});
</script>

<style scoped>
.fade-enter-active,
.fade-leave-active {
  transition: all 0.3s ease;
}
.fade-enter-from,
.fade-leave-to {
  opacity: 0;
  transform: translateY(-5px);
}
</style>