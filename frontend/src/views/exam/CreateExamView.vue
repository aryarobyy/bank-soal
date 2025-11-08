<template>
  <div
    class="min-h-screen bg-[#e8edff] flex flex-col items-center justify-center"
  >
    <div
      class="absolute top-0 left-0 right-0 flex justify-between items-center px-8 py-4"
    >
      <h1 class="text-3xl font-bold text-[#2a4dff]">Latih.in</h1>
      <div class="flex space-x-6 text-[#2a4dff] font-semibold">
        <router-link to="/admin/dashboard" class="hover:underline">Home</router-link>
        <router-link to="/admin/dashboard" class="hover:underline"
          >Dashboard</router-link
        >
        <i class="fas fa-user text-xl"></i>
      </div>
    </div>

    <div class="bg-white rounded-2xl shadow-lg p-8 w-[400px] text-center">
      <h2 class="text-lg font-bold text-[#2a4dff] mb-6">Name Exam</h2>

      <div class="space-y-4">
        <div>
          <label class="block text-left font-semibold mb-1">Pilih Soal</label>
          <select
            v-model="selectedQuestion"
            class="w-full border rounded-lg p-2"
          >
            <option disabled value="">-- Pilih Soal --</option>
            <option
              v-for="question in questions"
              :key="question.id"
              :value="question.id"
            >
              {{ question.name }}
            </option>
          </select>
        </div>

        <div>
          <label class="block text-left font-semibold mb-1"
            >Waktu Pengerjaan (menit)</label
          >
          <input
            v-model="duration"
            type="number"
            placeholder="Masukkan waktu"
            class="w-full border rounded-lg p-2"
          />
        </div>
      </div>

      <div class="mt-6 flex justify-center gap-4">
        <button
          @click="handleCancel"
          type="button" 
          class="bg-gray-300 text-gray-700 px-4 py-2 rounded hover:bg-gray-400 transition"
        >
          Cancel
        </button>
        <button
          @click="createExam"
          type="button"
          class="bg-teal-500 text-white px-4 py-2 rounded hover:bg-teal-600 transition"
        >
          Create
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from "vue";
import { useRouter } from "vue-router";

const router = useRouter();
const selectedQuestion = ref("");
const duration = ref("");

// Data dummy daftar soal
const questions = ref([
  { id: 1, name: "Soal Logika Dasar" },
  { id: 2, name: "Soal Algoritma" },
  { id: 3, name: "Soal Database" },
]);

const createExam = () => {
  if (!selectedQuestion.value || !duration.value) {
    alert("Harap isi semua field!");
    return;
  }

  // Simulasi simpan data ujian
  console.log("Ujian dibuat:", {
    soal_id: selectedQuestion.value,
    waktu: duration.value,
  });

  alert("Ujian berhasil dibuat!");
  
  // Kembali ke halaman daftar ujian admin
  router.push({ name: "AdminManageExam" }); 
};

// ## TAMBAHKAN FUNGSI INI ##
// Fungsi untuk tombol Cancel
const handleCancel = () => {
  // Kembali ke halaman daftar ujian admin
  router.push({ name: "AdminManageExam" });
};
</script>

<style scoped>
option[disabled] {
  color: #aaa;
}
</style>