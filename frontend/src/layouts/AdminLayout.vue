<template>
  <div class="font-sans bg-gray-50 min-h-screen flex flex-col">
    <header class="sticky top-0 z-30 flex items-center justify-between px-4 md:px-[5%] py-4 bg-white border-b border-gray-200 shadow-sm">
      <div class="flex items-center gap-3">
        <button @click="toggleSidebar" class="p-2 text-gray-600 rounded-md md:hidden hover:bg-gray-100 focus:outline-none">
          <i class="fas fa-bars text-xl"></i>
        </button>
        <router-link to="/admin/dashboard" class="text-2xl md:text-3xl font-bold no-underline text-red-600">
          Latih.in <span class="text-sm font-normal text-gray-500">(Admin)</span>
        </router-link>
      </div>

      <nav class="items-center hidden gap-10 md:flex">
        <router-link to="/admin/dashboard" class="font-semibold no-underline text-gray-600 hover:text-red-600">Dashboard</router-link>
      </nav>

      <div class="relative" ref="dropdownContainer">
        <div @click="toggleProfileDropdown" class="cursor-pointer flex items-center gap-2">
          <span class="hidden md:block text-sm font-medium text-gray-700 text-right">
            {{ user?.name || 'Admin' }}
          </span>
          <UserCircle class="w-10 h-10 text-gray-600 transition-colors hover:text-red-600" />
        </div>

        <transition
          enter-active-class="transition ease-out duration-100"
          enter-from-class="transform opacity-0 scale-95"
          enter-to-class="transform opacity-100 scale-100"
          leave-active-class="transition ease-in duration-75"
          leave-from-class="transform opacity-100 scale-100"
          leave-to-class="transform opacity-0 scale-95"
        >
          <div v-if="showProfileDropdown" class="absolute top-12 right-0 z-50 w-64 p-2 bg-white rounded-lg shadow-xl border border-gray-100 origin-top-right">
            <div class="flex items-center p-3 border-b border-gray-100 md:hidden">
              <div v-if="user">
                <div class="font-semibold text-gray-800">{{ user.name }}</div>
                <div class="text-xs text-gray-500">{{ user.email }}</div>
              </div>
            </div>
            
            <div class="py-1">
              <router-link 
                v-if="user"
                :to="`/admin/profile/${user.id}`" 
                @click="showProfileDropdown = false"
                class="flex items-center gap-2 px-4 py-2 text-sm text-gray-700 rounded hover:bg-red-50 hover:text-red-600"
              >
                <i class="fas fa-user w-4"></i> Lihat Profil
              </router-link>
              <a href="#" @click.prevent="logout" class="flex items-center gap-2 px-4 py-2 text-sm text-gray-700 rounded hover:bg-red-50 hover:text-red-600">
                <i class="fas fa-sign-out-alt w-4"></i> Logout
              </a>
            </div>
          </div>
        </transition>
      </div>
    </header>

    <div class="flex flex-1 relative">
      <div 
        v-if="isSidebarOpen" 
        @click="isSidebarOpen = false"
        class="fixed inset-0 z-40 bg-black bg-opacity-50 md:hidden glass-effect"
      ></div>

      <aside 
        :class="[
          'fixed inset-y-0 left-0 z-50 w-64 bg-white shadow-lg transform transition-transform duration-300 ease-in-out md:relative md:translate-x-0 md:shadow-none md:z-0',
          isSidebarOpen ? 'translate-x-0' : '-translate-x-full'
        ]"
      >
        <div class="h-full overflow-y-auto p-4 space-y-2">
          <div class="flex justify-between items-center mb-6 md:hidden">
            <span class="font-bold text-gray-700">Menu</span>
            <button @click="isSidebarOpen = false" class="text-gray-500 hover:text-red-600">
              <i class="fas fa-times text-xl"></i>
            </button>
          </div>

          <router-link
            v-for="(item, index) in menuItems"
            :key="index"
            :to="item.path"
            @click="isSidebarOpen = false"
            class="flex items-center gap-3 px-4 py-3 font-medium rounded-lg transition-all duration-200"
            :class="isActive(item.path) ? 'bg-red-50 text-red-600 shadow-sm' : 'text-gray-600 hover:bg-gray-50 hover:text-gray-900'"
          >
            <i :class="item.icon + ' w-5 text-center'"></i> {{ item.name }}
          </router-link>
        </div>
      </aside>

      <main class="flex-1 p-4 md:p-6 overflow-x-hidden w-full">
        <router-view></router-view>
      </main>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, onUnmounted } from 'vue';
import { useRouter, useRoute } from 'vue-router';
import { UserCircle } from 'lucide-vue-next';
import { useGetCurrentUser } from '../hooks/useGetCurrentUser';
import { useLocalStorage } from '../hooks/useLocalStorage';

import { logoutUser } from '../provider/user.provider';

const showProfileDropdown = ref(false);
const isSidebarOpen = ref(false); 
const dropdownContainer = ref(null); 

const router = useRouter();
const route = useRoute();
const { user } = useGetCurrentUser();

const menuItems = [
  { name: 'Dashboard', path: '/admin/dashboard', icon: 'fas fa-th-large' },
  { name: 'Mahasiswa', path: '/admin/mahasiswa', icon: 'fas fa-user-graduate' },
  { name: 'Dosen', path: '/admin/dosen', icon: 'fas fa-users' },
  { name: 'Soal', path: '/admin/soal', icon: 'fas fa-file-alt' },
  { name: 'Ujian', path: '/admin/ujian', icon: 'fas fa-chalkboard-teacher' },
  { name: 'Reports', path: '/admin/reports', icon: 'fas fa-chart-line' },
  { name: 'File Excel', path: '/admin/excel-files', icon: 'fas fa-file-excel' },
];

onMounted(() => {
  if (user.value) {
    if (user.value.role !== 'admin' && user.value.role !== 'super_admin') {
      alert('Akses ditolak! Halaman ini hanya untuk admin.');
      logout();
    }
  }
  document.addEventListener('click', handleClickOutside);
});

onUnmounted(() => {
  document.removeEventListener('click', handleClickOutside);
});

const handleClickOutside = (event) => {
  if (dropdownContainer.value && !dropdownContainer.value.contains(event.target)) {
    showProfileDropdown.value = false;
  }
};

const toggleProfileDropdown = () => {
  showProfileDropdown.value = !showProfileDropdown.value;
};

const toggleSidebar = () => {
  isSidebarOpen.value = !isSidebarOpen.value;
};

const logout = async () => {

  await logoutUser();


  const { removeValue: removeToken } = useLocalStorage('token');
  const { removeValue: removeId } = useLocalStorage('id');
  removeToken();
  removeId();
  localStorage.removeItem('user_id');
  localStorage.clear();
  
  // 3. Redirect
  user.value = null;
  window.location.href = '/landing'; 
};

const isActive = (path) => {
  return route.path.startsWith(path);
};
</script>