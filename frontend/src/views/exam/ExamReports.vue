<template>
  <div class="p-6 bg-[#E8EDFF] min-h-screen">
    <!-- Header -->
    <div class="flex justify-between items-center mb-6">
      <h1 class="text-3xl font-bold text-[#2A4DFF]">Laporan Nilai Ujian</h1>
      <button
        @click="downloadExcel"
        class="flex items-center gap-2 bg-green-500 text-white px-4 py-2 rounded hover:bg-green-600 transition"
      >
        <i class="fas fa-file-excel"></i>
        Unduh
      </button>
    </div>

    <!-- Table -->
    <div class="bg-white p-6 rounded-xl shadow-md">
      <table class="w-full border-collapse text-sm">
        <thead>
          <tr class="bg-[#f5f7ff] text-gray-600">
            <th class="p-3 text-left">ID</th>
            <th class="p-3 text-left">Username</th>
            <th class="p-3 text-left">Fullname</th>
            <th class="p-3 text-left">Score</th>
            <th class="p-3 text-left">Keterangan</th>
          </tr>
        </thead>
        <tbody>
          <tr
            v-for="(report, index) in reports"
            :key="index"
            class="border-t hover:bg-[#eef3ff]"
          >
            <td class="p-3">{{ report.id }}</td>
            <td class="p-3">{{ report.username }}</td>
            <td class="p-3">{{ report.fullname }}</td>
            <td class="p-3">{{ report.score }}</td>
            <td
              class="p-3 font-semibold"
              :class="{
                'text-green-600': report.score >= 70,
                'text-red-500': report.score < 70,
              }"
            >
              {{ report.score >= 70 ? "Lulus" : "Tidak Lulus" }}
            </td>
          </tr>
          <tr v-if="reports.length === 0">
            <td colspan="5" class="p-4 text-center text-gray-400">
              Belum ada data laporan.
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script setup>
import * as XLSX from "xlsx";

// Data dummy sementara
const reports = [
  { id: 1, username: "mryan93", fullname: "Muhammad Ryan", score: 85 },
  { id: 2, username: "aisyah02", fullname: "Aisyah Rahma", score: 67 },
  { id: 3, username: "dimas11", fullname: "Dimas Ardiansyah", score: 90 },
];

// Fungsi export Excel
const downloadExcel = () => {
  const ws = XLSX.utils.json_to_sheet(reports);
  const wb = XLSX.utils.book_new();
  XLSX.utils.book_append_sheet(wb, ws, "Laporan Nilai");
  XLSX.writeFile(wb, "laporan_nilai.xlsx");
};
</script>

<style scoped>
table th,
table td {
  border-bottom: 1px solid #ddd;
}
</style>
