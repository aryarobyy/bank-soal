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
          <img
            src="/src/assets/user-icon.png"
            alt="Profile"
            class="w-8 h-8 rounded-full border border-[#3B3B98]"
          />
        </button>
      </nav>
    </header>

    <!-- Profile Card -->
    <main class="max-w-6xl mx-auto bg-white rounded-2xl shadow-md p-10 mt-10">
      <div class="flex items-center gap-5 mb-10">
        <img
          :src="superAdmin.photo || defaultAvatar"
          alt="Super Admin Avatar"
          class="w-20 h-20 rounded-full object-cover border border-gray-300"
        />
        <div>
          <h2 class="text-xl font-semibold text-gray-800">
            {{ superAdmin.fullName }}
          </h2>
          <p class="text-gray-500">{{ superAdmin.email }}</p>
        </div>
      </div>

      <!-- Two Columns -->
      <div class="grid grid-cols-2 gap-10">
        <div>
          <label class="block text-gray-600 mb-2">Full Name</label>
          <input
            v-model="superAdmin.fullName"
            type="text"
            class="w-full border border-gray-300 rounded-lg p-3"
            placeholder="Enter full name"
          />

          <label class="block text-gray-600 mt-4 mb-2">Gender</label>
          <select
            v-model="superAdmin.gender"
            class="w-full border border-gray-300 rounded-lg p-3"
          >
            <option disabled value="">Pick Your Gender</option>
            <option>Male</option>
            <option>Female</option>
          </select>

          <div class="mt-6 bg-gray-50 rounded-lg p-4">
            <h3 class="text-gray-700 font-semibold mb-2">My email address</h3>
            <div class="flex items-center gap-3">
              <img src="/src/assets/mail-icon.png" alt="Mail" class="w-5 h-5" />
              <div>
                <p class="text-gray-700">{{ superAdmin.email }}</p>
                <p class="text-gray-400 text-sm">
                  {{ superAdmin.lastUpdated }}
                </p>
              </div>
            </div>
          </div>
        </div>

        <div>
          <label class="block text-gray-600 mb-2">ID Super Admin</label>
          <input
            v-model="superAdmin.id"
            type="text"
            class="w-full border border-gray-300 rounded-lg p-3"
            placeholder="Enter ID"
          />

          <label class="block text-gray-600 mt-4 mb-2">Email</label>
          <input
            v-model="superAdmin.email"
            type="email"
            class="w-full border border-gray-300 rounded-lg p-3"
            placeholder="Enter email"
          />

          <label class="block text-gray-600 mt-4 mb-2">Riwayat Aktivitas</label>
          <textarea
            v-model="superAdmin.history"
            class="w-full border border-gray-300 rounded-lg p-3 min-h-[100px]"
            placeholder="Contoh: Menambah admin, memperbarui data ujian..."
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
import defaultAvatar from "../assets/default-avatar.png";

const superAdmin = ref({
  fullName: "Jonathan Wibowo",
  id: "SUPER001",
  email: "jonathan.wibowo@latihin.com",
  gender: "Male",
  history: "Menambah admin baru dan meninjau laporan sistem.",
  photo: "",
  lastUpdated: "2 weeks ago",
});

const saveProfile = () => {
  alert("Data profil super admin berhasil diperbarui (mock data)");
};
</script>

<style scoped>
body {
  font-family: "Poppins", sans-serif;
}
</style>
