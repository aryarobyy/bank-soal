<template>
  <div class="font-sans bg-background min-h-screen">
    <header class="sticky top-0 z-50 flex items-center justify-between px-[5%] py-4 bg-white border-b border-gray-200 shadow-md">
      <router-link to="/admin/dashboard" class="text-3xl font-bold no-underline text-red-600">Latih.in (Admin)</router-link>
      <nav class="items-center hidden gap-10 md:flex">
        <router-link to="/admin/dashboard" class="font-semibold no-underline text-medium-text hover:text-red-600">Dashboard</router-link>
        </nav>
      <div class="relative">
        <div @click="toggleProfileDropdown" class="cursor-pointer">
          <UserCircle class="w-10 h-10 text-gray-600 transition-colors hover:text-red-600" />
        </div>
        <div v-if="showProfileDropdown" class="absolute top-16 right-0 z-50 w-72 p-2 bg-white rounded-lg shadow-xl border border-gray-100">
            <div class="flex items-center p-2">
              <UserCircle class="w-12 h-12 text-gray-600 mr-4" />
              <div v-if="user">
                <div class="font-semibold">{{ user.name }}</div>
                <div class="text-sm text-gray-500">{{ user.email }}</div>
              </div>
            </div>
            <hr class="my-2 border-gray-200" />
            
            <router-link 
              v-if="user"
              :to="`/admin/profile/${user.id}`" 
              class="block w-full px-4 py-2 text-left text-gray-700 rounded hover:bg-gray-100"
            >
              Lihat Profil
            </router-link>
            <a href="#" @click.prevent="logout" class="block w-full px-4 py-2 text-left text-gray-700 rounded hover:bg-gray-100">Logout</a>
        </div>
      </div>
    </header>

    <main class="flex gap-6 p-6">
      <aside class="flex-shrink-0 w-64">
        <div class="p-4 space-y-2 bg-white rounded-lg shadow-md">
          <router-link
            to="/admin/dashboard"
            class="flex items-center gap-3 px-4 py-2 font-semibold rounded-md transition-colors"
            :class="isActive('/admin/dashboard') ? 'text-primary bg-indigo-50' : 'text-gray-600 hover:bg-gray-100'"
          >
            <i class="fas fa-th-large"></i> Dashboard
          </router-link>
          <router-link
            to="/admin/mahasiswa"
            class="flex items-center gap-3 px-4 py-2 font-semibold rounded-md transition-colors"
            :class="isActive('/admin/mahasiswa') ? 'text-primary bg-indigo-50' : 'text-gray-600 hover:bg-gray-100'"
          >
            <i class="fas fa-user-graduate"></i> Mahasiswa
          </router-link>
          <router-link
            to="/admin/dosen"
            class="flex items-center gap-3 px-4 py-2 font-semibold rounded-md transition-colors"
            :class="isActive('/admin/dosen') ? 'text-primary bg-indigo-50' : 'text-gray-600 hover:bg-gray-100'"
          >
            <i class="fas fa-users"></i> Dosen
          </router-link>
          <router-link
            to="/admin/soal"
            class="flex items-center gap-3 px-4 py-2 font-semibold rounded-md transition-colors"
            :class="isActive('/admin/soal') ? 'text-primary bg-indigo-50' : 'text-gray-600 hover:bg-gray-100'"
          >
            <i class="fas fa-file-alt"></i> Soal
          </router-link>
          <router-link
            to="/admin/ujian"
            class="flex items-center gap-3 px-4 py-2 font-semibold rounded-md transition-colors"
            :class="isActive('/admin/ujian') ? 'text-primary bg-indigo-50' : 'text-gray-600 hover:bg-gray-100'"
          >
            <i class="fas fa-chalkboard-teacher"></i> Ujian
          </router-link>
          <router-link
            to="/admin/reports"
            class="flex items-center gap-3 px-4 py-2 font-semibold rounded-md transition-colors"
            :class="isActive('/admin/reports') ? 'text-primary bg-indigo-50' : 'text-gray-600 hover:bg-gray-100'"
          >
            <i class="fas fa-chart-line"></i> Reports
          </router-link>
        </div>
      </aside>

      <div class="flex-1">
        <router-view></router-view>
      </div>
    </main>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { useRouter, useRoute, RouterLink } from 'vue-router';
import { UserCircle } from 'lucide-vue-next';
import { useGetCurrentUser } from '../hooks/useGetCurrentUser';
import { useLocalStorage } from '../hooks/useLocalStorage';

const showProfileDropdown = ref(false);
const router = useRouter();
const route = useRoute();

const { user } = useGetCurrentUser(); // user sudah tersedia di sini

onMounted(() => {
  if (user.value) {
    if (user.value.role !== 'admin' && user.value.role !== 'super_admin') {
      alert('Akses ditolak! Halaman ini hanya untuk admin.');
      logout();
    }
  } else {
    // App.vue akan menangani redirect
  }
});

const toggleProfileDropdown = () => {
  showProfileDropdown.value = !showProfileDropdown.value;
};

const logout = () => {
  const { removeValue: removeToken } = useLocalStorage('token');
  const { removeValue: removeId } = useLocalStorage('id');
  removeToken();
  removeId();
  window.location.href = '/login'; 
};

const isActive = (path) => {
  return route.path.startsWith(path);
};
</script>