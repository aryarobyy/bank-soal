<template>
  <div class="flex min-h-screen bg-[#f4f6ff] text-gray-800">
    <aside class="w-64 bg-white shadow-lg flex flex-col">
      <div class="p-6 text-2xl font-bold text-blue-700">
        Latih.in Super Admin
      </div>

      <nav class="flex-1 px-4 space-y-2">
        <router-link
          to="/superadmin/dashboard"
          class="flex items-center gap-2 p-3 rounded-lg hover:bg-blue-100"
          active-class="bg-blue-200 font-semibold text-blue-700"
        >
          <i class="fa fa-home"></i> Dashboard
        </router-link>

        <router-link
          to="/superadmin/admins"
          class="flex items-center gap-2 p-3 rounded-lg hover:bg-blue-100"
          active-class="bg-blue-200 font-semibold text-blue-700"
        >
          <i class="fa fa-user-shield"></i> Kelola Admin
        </router-link>
      </nav>

      <div class="p-4 mt-auto">
        <button
          @click="logout"
          class="w-full flex items-center justify-center gap-2 px-4 py-2 font-semibold text-white bg-red-500 rounded-lg hover:bg-red-600 transition"
        >
          <i class="fa fa-sign-out-alt"></i> Logout
        </button>
      </div>
    </aside>

    <div class="flex-1 flex flex-col">
      <header class="flex items-center justify-between p-4 bg-white shadow-md">
        <h1 class="text-2xl font-bold text-blue-700">{{ $route.meta.title || 'Super Admin Panel' }}</h1>
        <div v-if="user" class="font-semibold text-gray-700">
          Selamat datang, {{ user.name }}
        </div>
      </header>

      <main class="flex-1 p-6">
        <router-view />
      </main>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { useRouter, useRoute } from 'vue-router';
// ## 1. Impor hook state global yang baru ##
import { useGetCurrentUser } from '../hooks/useGetCurrentUser';
import { useLocalStorage } from '../hooks/useLocalStorage';

const router = useRouter();
const route = useRoute(); // Untuk $route.meta.title

// ## 2. Gunakan hook baru untuk mendapatkan data user ##
const { user } = useGetCurrentUser();

// ## 3. Perbarui 'onMounted' untuk menggunakan 'user' dari global state ##
onMounted(() => {
  if (user.value) {
    // Periksa apakah role pengguna adalah 'super_admin'
    if (user.value.role !== 'super_admin') { 
      alert('Akses ditolak! Halaman ini hanya untuk Super Admin.');
      logout();
    }
  } else {
    // Jika tidak ada data pengguna, App.vue akan menangani
    // Jika App.vue gagal (token tidak valid), ia akan menghapus token
    // dan reload ini akan gagal lagi, yang akan ditangani oleh router guard (jika ada)
    // atau App.vue akan mengarahkan ke login.
  }
});

// ## 4. Perbarui fungsi logout ##
const logout = () => {
  const { removeValue: removeToken } = useLocalStorage('token');
  const { removeValue: removeId } = useLocalStorage('id'); // Ganti 'user' menjadi 'id'
  removeToken();
  removeId();
  
  // Arahkan ke /login dan lakukan hard reload
  window.location.href = '/login';
};
</script>