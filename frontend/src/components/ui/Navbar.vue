<template>
  <nav class="bg-white shadow-sm border-b border-gray-200 relative z-50">
    <div class="px-6 py-3 flex justify-between items-center">
      
      <RouterLink to="/dashboard" class="text-3xl font-bold no-underline text-blue-600">
        Latih.in
      </RouterLink>

      <ul class="hidden md:flex items-center space-x-6">
        <template v-if="user">
          <li>
            <RouterLink to="/dashboard" :class="linkClass('/dashboard')">Dashboard</RouterLink>
          </li>
          <li v-if="user.role === 'user'">
            <RouterLink to="/ujian" :class="linkClass('/ujian')">Ujian</RouterLink>
          </li>
   
          <div class="relative" ref="dropdownRef">
            <button @click="toggleDropdown" class="flex items-center space-x-2 focus:outline-none">
              <User class="w-6 h-6 text-gray-700" />
              <span class="text-sm font-medium text-gray-700">{{ user.name }}</span>
              <ChevronDown class="w-4 h-4 text-gray-600 transition-transform" :class="{ 'rotate-180': isDropdownOpen }" />
            </button>

            <transition name="fade">
              <div v-if="isDropdownOpen" class="absolute right-0 mt-2 w-48 bg-white rounded-md shadow-lg border border-gray-200 z-50">
                <ul class="py-1">
                  <li>
                    <RouterLink :to="`/profile/${user.id}`" class="block px-4 py-2 text-sm text-gray-700 hover:bg-gray-100" @click="closeDropdown">
                      Lihat Profil
                    </RouterLink>
                  </li>
                  <li>
                    <button @click="handleLogout" class="w-full text-left block px-4 py-2 text-sm text-red-600 hover:bg-gray-100">
                      Logout
                    </button>
                  </li>
                </ul>
              </div>
            </transition>
          </div>
        </template>

        <template v-else>
          <li><RouterLink to="/login" :class="linkClass('/login')">Login</RouterLink></li>
        </template>
      </ul>

      <button 
        @click="toggleMobileMenu" 
        class="md:hidden p-2 text-gray-600 hover:bg-gray-100 rounded-md focus:outline-none"
      >
        <Menu v-if="!isMobileMenuOpen" class="w-6 h-6" />
        <X v-else class="w-6 h-6" />
      </button>
    </div>

    <transition name="fade">
      <div 
        v-if="isMobileMenuOpen" 
        class="md:hidden absolute top-full left-0 w-full bg-white border-b border-gray-200 shadow-lg z-40"
      >
        <ul class="flex flex-col p-4 space-y-2">
          <template v-if="user">
            <li>
              <RouterLink to="/" :class="mobileLinkClass('/')" @click="closeMobileMenu">
                Dashboard
              </RouterLink>
            </li>
            <li v-if="user.role === 'user'">
              <RouterLink to="/ujian" :class="mobileLinkClass('/ujian')" @click="closeMobileMenu">
                Ujian
              </RouterLink>
            </li>

            <li class="pt-3 mt-2 border-t border-gray-100">
              <div class="flex items-center gap-2 px-3 mb-2 text-gray-500 font-medium text-xs uppercase tracking-wider">
                <User class="w-4 h-4" />
                <span>Akun Anda ({{ user.name }})</span>
              </div>
            </li>
            
            <li>
              <RouterLink :to="`/profile/${user.id}`" class="block px-3 py-2 text-sm font-medium text-gray-700 rounded-md hover:bg-blue-50 hover:text-blue-600" @click="closeMobileMenu">
                Lihat Profil
              </RouterLink>
            </li>
            <li>
              <button @click="handleLogout" class="w-full text-left block px-3 py-2 text-sm font-medium text-red-600 rounded-md hover:bg-red-50">
                Logout
              </button>
            </li>
          </template>

          <template v-else>
            <li>
              <RouterLink to="/login" :class="mobileLinkClass('/login')" @click="closeMobileMenu">
                Login
              </RouterLink>
            </li>
          </template>
        </ul>
      </div>
    </transition>
  </nav>
</template>

<script setup>
import { ref, onMounted, onBeforeUnmount } from 'vue'
import { useRoute, useRouter, RouterLink } from 'vue-router'
import { useGetCurrentUser } from '../../hooks/useGetCurrentUser'
import { useLocalStorage } from '../../hooks/useLocalStorage'
import { User, ChevronDown, Menu, X } from 'lucide-vue-next'
import { logoutUser } from '../../provider/user.provider'

const { user } = useGetCurrentUser()
const route = useRoute()
const router = useRouter()

const isDropdownOpen = ref(false)
const dropdownRef = ref(null)
const isMobileMenuOpen = ref(false)

const { removeValue: removeToken } = useLocalStorage('token');
const { removeValue: removeUser } = useLocalStorage('user');
const { removeValue: removeId } = useLocalStorage('id');

const linkClass = (path) => {
  const isActive = route.path === path;
  return [
    'transition-colors duration-200 font-medium',
    isActive ? 'text-blue-600' : 'text-gray-600 hover:text-blue-600'
  ].join(' ');
};

const mobileLinkClass = (path) => {
  const isActive = route.path === path;
  return [
    'block px-3 py-2 rounded-md text-base font-medium transition-colors',
    isActive ? 'bg-blue-50 text-blue-600' : 'text-gray-700 hover:bg-gray-50 hover:text-blue-600'
  ].join(' ');
};

const toggleDropdown = () => { isDropdownOpen.value = !isDropdownOpen.value; };
const closeDropdown = () => { isDropdownOpen.value = false; };
const toggleMobileMenu = () => { isMobileMenuOpen.value = !isMobileMenuOpen.value; };
const closeMobileMenu = () => { isMobileMenuOpen.value = false; };

const handleLogout = async () => {
  await logoutUser();
  removeToken();
  removeUser();
  removeId();
  localStorage.removeItem('user_id');
  localStorage.clear(); 

  closeDropdown();
  closeMobileMenu();
  user.value = null; 
  
  window.location.href = '/landing'; 
};

const handleClickOutside = (event) => {
  if (dropdownRef.value && !dropdownRef.value.contains(event.target)) {
    closeDropdown();
  }
};

onMounted(() => { document.addEventListener('mousedown', handleClickOutside); });
onBeforeUnmount(() => { document.removeEventListener('mousedown', handleClickOutside); });
</script>

<style scoped>
.fade-enter-active, .fade-leave-active { transition: opacity 0.2s ease, transform 0.2s ease; }
.fade-enter-from, .fade-leave-to { opacity: 0; transform: translateY(-5px); }
.rotate-180 { transform: rotate(180deg); }
</style>