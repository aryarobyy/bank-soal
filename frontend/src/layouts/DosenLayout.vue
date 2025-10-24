<template>
  <div class="font-sans bg-background min-h-screen">
    
    <header class="sticky top-0 z-50 flex items-center justify-between px-[5%] py-4 bg-white border-b border-gray-200 shadow-md">
      <router-link to="/dosen/dashboard" class="text-3xl font-bold no-underline text-primary">Latih.in</router-link>
      <nav class="items-center hidden gap-10 md:flex">
        <router-link to="/dosen/dashboard" class="font-semibold no-underline text-medium-text hover:text-primary">Dashboard</router-link>
        <router-link to="/dosen/soal" class="font-semibold no-underline text-medium-text hover:text-primary">Soal</router-link>
      </nav>
      <div class="relative">
        <div @click="toggleProfileDropdown" class="cursor-pointer">
          <img :src="userIcon" alt="User Profile" class="w-10 h-10 rounded-full" />
        </div>
        <div v-if="showProfileDropdown" class="absolute top-16 right-0 z-50 w-72 p-2 bg-white rounded-lg shadow-xl border border-gray-100">
            <div class="flex items-center p-2">
              <img :src="userIcon" alt="User Profile" class="w-12 h-12 rounded-full mr-4" />
              <div v-if="user">
                <div class="font-semibold">{{ user.name }}</div>
                <div class="text-sm text-gray-500">{{ user.email }}</div>
              </div>
            </div>
            <hr class="my-2 border-gray-200" />
            <router-link to="/dosen/update-profile" class="block w-full px-4 py-2 text-left text-gray-700 rounded hover:bg-gray-100">Update Profile</router-link>
            <a href="#" @click.prevent="logout" class="block w-full px-4 py-2 text-left text-gray-700 rounded hover:bg-gray-100">Logout</a>
        </div>
      </div>
    </header>

    <main class="flex gap-6 p-6">
      <aside class="flex-shrink-0 w-64">
        <div class="p-4 space-y-2 bg-white rounded-lg shadow-md">
          <router-link 
            to="/dosen/dashboard" 
            class="flex items-center gap-3 px-4 py-2 font-semibold rounded-md transition-colors" 
            :class="isActive('/dosen/dashboard') ? 'text-primary bg-indigo-50' : 'text-gray-600 hover:bg-gray-100'">
            <i class="fas fa-tachometer-alt"></i> Dashboard
          </router-link>
          <router-link 
            to="/dosen/soal" 
            class="flex items-center gap-3 px-4 py-2 font-semibold rounded-md transition-colors" 
            :class="isActive('/dosen/soal') ? 'text-primary bg-indigo-50' : 'text-gray-600 hover:bg-gray-100'">
            <i class="fas fa-file-alt"></i> Soal
          </router-link>
        </div>
      </aside>

      <div class="flex-1">
        <router-view></router-view>
      </div>
    </main>
  </div>
</template>

<script>
import userIcon from '../assets/user-icon.png';
import { useLocalStorage } from '../hooks/useLocalStorage';

export default {
  name: 'DosenLayout',
  data() {
    return {
      showProfileDropdown: false,
      user: null,
      userIcon,
    };
  },
  methods: {
    toggleProfileDropdown() {
      this.showProfileDropdown = !this.showProfileDropdown;
    },
    logout() {
      const { removeValue: removeToken } = useLocalStorage('token');
      const { removeValue: removeUser } = useLocalStorage('user');
      removeToken();
      removeUser();
      this.$router.push('/login');
    },
    isActive(path) {
      return this.$route.path.startsWith(path);
    }
  },
  created() {
    const { value: storedUser } = useLocalStorage('user');
    if (storedUser.value) {
      this.user = storedUser.value;
    } else {
      this.logout();
    }
  }
};
</script>