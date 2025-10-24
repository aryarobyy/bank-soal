<template>
  <nav class="bg-white shadow-sm border-b border-gray-200 px-6 py-3 flex justify-between items-center">
    <RouterLink to="/">
      <img 
        :src="Logo" 
        alt="Latih.in Logo"
        class="w-28 object-contain"
      />
    </RouterLink>

    <ul class="flex items-center space-x-6">
      <template v-if="user">
        <li v-if="user.role === 'user'">
          <RouterLink
            to="/ujian"
            :class="linkClass('/ujian')"
          >
            Ujian
          </RouterLink>
        </li>

        <div class="relative" ref="dropdownRef">
          <button @click="toggleDropdown" class="flex items-center space-x-2 focus:outline-none">
            <User class="w-6 h-6 text-gray-700" />
            <span class="text-sm font-medium text-gray-700">
              {{ user.name }}
            </span>
            <ChevronDown class="w-4 h-4 text-gray-600 transition-transform" :class="{ 'rotate-180': isDropdownOpen }" />
          </button>

          <transition name="fade">
            <div v-if="isDropdownOpen" class="absolute right-0 mt-2 w-48 bg-white rounded-md shadow-lg border border-gray-200 z-50">
              <ul class="py-1">
                <li>
                  <RouterLink
                    :to="`/profile/${user.id}`"
                    class="block px-4 py-2 text-sm text-gray-700 hover:bg-gray-100"
                    @click="closeDropdown"
                  >
                    Lihat Profil
                  </RouterLink>
                </li>
                <li>
                  <button
                    @click="handleLogout"
                    class="w-full text-left block px-4 py-2 text-sm text-red-600 hover:bg-gray-100"
                  >
                    Logout
                  </button>
                </li>
              </ul>
            </div>
          </transition>
        </div>
      </template>

      <template v-else>
        <li>
          <RouterLink
            to="/login"
            :class="linkClass('/login')"
          >
            Login
          </RouterLink>
        </li>
        <li>
          <RouterLink
            to="/register"
            :class="linkClass('/register')"
          >
            Register
          </RouterLink>
        </li>
      </template>
    </ul>
  </nav>
</template>

<script setup>
import { ref, onMounted, onBeforeUnmount } from 'vue'
import { useRoute, useRouter, RouterLink } from 'vue-router'
import { useGetCurrentUser } from '../../hooks/useGetCurrentUser'
import { useLocalStorage } from '../../hooks/useLocalStorage'
import { User, ChevronDown } from 'lucide-vue-next'
import Logo from '../../assets/logo.png'

const { user } = useGetCurrentUser()
const route = useRoute()
const router = useRouter()

// State untuk mengontrol visibilitas dropdown
const isDropdownOpen = ref(false)
const dropdownRef = ref(null)

// Mengambil fungsi untuk menghapus data dari Local Storage
const { removeValue: removeToken } = useLocalStorage('token');
const { removeValue: removeUser } = useLocalStorage('user');

const linkClass = (path) => {
  const isActive = route.path === path || route.path.startsWith(`${path}/`);
  return [
    'transition-colors duration-200 font-medium',
    isActive
      ? 'text-blue-600'
      : 'text-gray-600 hover:text-blue-600'
  ].join(' ');
};

// Fungsi untuk toggle dropdown
const toggleDropdown = () => {
  isDropdownOpen.value = !isDropdownOpen.value;
};

const closeDropdown = () => {
  isDropdownOpen.value = false;
}

// Fungsi untuk logout
const handleLogout = () => {
  removeToken();
  removeUser();
  closeDropdown();
  router.push('/login');
};

// Fungsi untuk menutup dropdown saat klik di luar area menu
const handleClickOutside = (event) => {
  if (dropdownRef.value && !dropdownRef.value.contains(event.target)) {
    closeDropdown();
  }
};

// Menambahkan dan menghapus event listener
onMounted(() => {
  document.addEventListener('mousedown', handleClickOutside);
});

onBeforeUnmount(() => {
  document.removeEventListener('mousedown', handleClickOutside);
});
</script>

<style scoped>
/* Transisi untuk dropdown */
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.2s ease, transform 0.2s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
  transform: translateY(-10px);
}

.rotate-180 {
  transform: rotate(180deg);
}
</style>