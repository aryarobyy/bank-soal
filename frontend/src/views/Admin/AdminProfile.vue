<template>
  <div class="min-h-screen bg-[#E8EDFF] font-[Poppins]">
    <!-- Navbar -->
    <header class="flex justify-between items-center px-12 py-6 bg-[#E8EDFF]">
      <h1 class="text-3xl font-bold text-[#3B3B98]">Latih.in</h1>
      <nav class="flex items-center gap-8">
        <a
          href="/home"
          class="text-lg font-semibold text-[#3B3B98] hover:underline"
          >Home</a
        >
        <a
          href="/exam"
          class="text-lg font-semibold text-[#3B3B98] hover:underline"
          >Dashboard</a
        >
        <button>
          <User
            class="w-8 h-8 rounded-full border border-[#3B3B98] text-[#3B3B98]"
          />
        </button>
      </nav>
    </header>

    <!-- Profile Card -->
    <main class="max-w-6xl mx-auto bg-white rounded-2xl shadow-md p-10 mt-10">
      <div class="flex items-center gap-5 mb-10">
        <img
          :src="admin.photo || defaultAvatar"
          alt="Admin Avatar"
          class="w-20 h-20 rounded-full object-cover border border-gray-300"
        />
        <div>
          <h2 class="text-xl font-semibold text-gray-800">
            {{ admin.fullName }}
          </h2>
          <p class="text-gray-500">{{ admin.email }}</p>
        </div>
      </div>

      <!-- Two Columns -->
      <div class="grid grid-cols-2 gap-10">
        <div>
          <label class="block text-gray-600 mb-2">Full Name</label>
          <input
            v-model="admin.fullName"
            type="text"
            class="w-full border border-gray-300 rounded-lg p-3"
            placeholder="Enter full name"
          />

          <label class="block text-gray-600 mt-4 mb-2">Gender</label>
          <select
            v-model="admin.gender"
            class="w-full border border-gray-300 rounded-lg p-3"
          >
            <option disabled value="">Pick Your Gender</option>
            <option>Male</option>
            <option>Female</option>
          </select>

          <div class="mt-6 bg-gray-50 rounded-lg p-4">
            <h3 class="text-gray-700 font-semibold mb-2">My email address</h3>
            <div class="flex items-center gap-3">
              <Mail class="w-5 h-5" />
              <div>
                <p class="text-gray-700">{{ admin.email }}</p>
                <p class="text-gray-400 text-sm">{{ admin.lastUpdated }}</p>
              </div>
            </div>
          </div>
        </div>

        <div>
          <label class="block text-gray-600 mb-2">NIP</label>
          <input
            v-model="admin.nip"
            type="text"
            class="w-full border border-gray-300 rounded-lg p-3"
            placeholder="Enter NIP"
          />

          <label class="block text-gray-600 mt-4 mb-2">Email</label>
          <input
            v-model="admin.email"
            type="email"
            class="w-full border border-gray-300 rounded-lg p-3"
            placeholder="Enter email"
          />

          <label class="block text-gray-600 mt-4 mb-2">Riwayat Aktivitas</label>
          <textarea
            v-model="admin.history"
            class="w-full border border-gray-300 rounded-lg p-3 min-h-[100px]"
            placeholder="Contoh: Mengelola ujian, menambahkan user..."
          ></textarea>
        </div>
      </div>

      <div class="mt-10 flex justify-center">
        <button
          @click="saveProfile"
          class="bg-[#3B3B98] text-white px-8 py-2 rounded-lg hover:bg-[#2d2d75] transition"
        >
          Edit
        </button>
      </div>
    </main>
  </div>
</template>

<script setup>
import { ref } from "vue";
import { User, Mail } from 'lucide-vue-next';


const defaultAvatar = "data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' width='100' height='100' viewBox='0 0 24 24' fill='none' stroke='%23333' stroke-width='2' stroke-linecap='round' stroke-linejoin='round'%3E%3Cpath d='M19 21v-2a4 4 0 0 0-4-4H9a4 4 0 0 0-4 4v2'%3E%3C/path%3E%3Ccircle cx='12' cy='7' r='4'%3E%3C/circle%3E%3C/svg%3E";

const admin = ref({
  fullName: "Alexa Rawles",
  nip: "ADM001",
  email: "alexarawles@gmail.com",
  gender: "Female",
  history: "Mengelola ujian dan meninjau hasil mahasiswa.",
  photo: "",
  lastUpdated: "1 month ago",
});

const saveProfile = () => {
  alert("Data profil admin berhasil diperbarui (mock data)");
};
</script>

<style scoped>
body {
  font-family: "Poppins", sans-serif;
}
</style>
