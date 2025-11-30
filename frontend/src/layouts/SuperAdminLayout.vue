<template>
  <div class="flex min-h-screen bg-[#f4f6ff] text-gray-800 relative">
    
    <div 
      v-if="isSidebarOpen" 
      @click="isSidebarOpen = false"
      class="fixed inset-0 z-20 bg-black bg-opacity-50 md:hidden"
    ></div>

    <aside 
      :class="[
        'fixed inset-y-0 left-0 z-30 w-64 bg-white shadow-lg flex flex-col transform transition-transform duration-300 ease-in-out md:relative md:translate-x-0',
        isSidebarOpen ? 'translate-x-0' : '-translate-x-full'
      ]"
    >
      <div class="flex items-center justify-between p-6">
        <div class="text-2xl font-bold text-blue-700">
          Latih.in <span class="text-xs block font-normal text-gray-500">Super Admin</span>
        </div>
        <button @click="isSidebarOpen = false" class="md:hidden text-gray-500 hover:text-red-500">
          <i class="fas fa-times text-xl"></i>
        </button>
      </div>

      <nav class="flex-1 px-4 space-y-2">
        <router-link
          to="/superadmin/dashboard"
          @click="isSidebarOpen = false"
          class="flex items-center gap-2 p-3 rounded-lg hover:bg-blue-100 transition-colors"
          active-class="bg-blue-200 font-semibold text-blue-700"
        >
          <i class="fas fa-home w-6 text-center"></i> Dashboard
        </router-link>

        <router-link
          to="/superadmin/admins"
          @click="isSidebarOpen = false"
          class="flex items-center gap-2 p-3 rounded-lg hover:bg-blue-100 transition-colors"
          active-class="bg-blue-200 font-semibold text-blue-700"
        >
          <i class="fas fa-user-shield w-6 text-center"></i> Kelola Admin
        </router-link>
      </nav>

      <div class="p-4 mt-auto">
        <button
          @click="logout"
          class="w-full flex items-center justify-center gap-2 px-4 py-2 font-semibold text-white bg-red-500 rounded-lg hover:bg-red-600 transition"
        >
          <i class="fas fa-sign-out-alt"></i> Logout
        </button>
      </div>
    </aside>

    <div class="flex-1 flex flex-col min-w-0"> <header class="flex items-center justify-between p-4 bg-white shadow-md sticky top-0 z-10">
        <div class="flex items-center gap-3">
          <button @click="isSidebarOpen = !isSidebarOpen" class="p-2 text-gray-600 rounded-md md:hidden hover:bg-gray-100">
            <i class="fas fa-bars text-xl"></i>
          </button>
          
          <h1 class="text-lg md:text-2xl font-bold text-blue-700 truncate">
            {{ $route.meta.title || 'Super Admin Panel' }}
          </h1>
        </div>

        <div v-if="user" class="font-semibold text-gray-700 text-sm md:text-base">
          Hi, {{ user.name }}
        </div>
      </header>

      <main class="flex-1 p-4 md:p-6 overflow-x-hidden">
        <router-view />
      </main>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { useRouter, useRoute } from 'vue-router';
import { useGetCurrentUser } from '../hooks/useGetCurrentUser';
import { useLocalStorage } from '../hooks/useLocalStorage';

const router = useRouter();
const route = useRoute();
const { user } = useGetCurrentUser();
const isSidebarOpen = ref(false); 

onMounted(() => {
  if (user.value) {
    if (user.value.role !== 'super_admin') { 
      alert('Akses ditolak! Halaman ini hanya untuk Super Admin.');
      logout();
    }
  }
});

const logout = () => {
  const { removeValue: removeToken } = useLocalStorage('token');
  const { removeValue: removeId } = useLocalStorage('id');
  removeToken();
  removeId();
  window.location.href = '/';
};
</script>