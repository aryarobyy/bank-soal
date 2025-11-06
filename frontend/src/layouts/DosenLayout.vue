<template>
  <div class="font-sans bg-background min-h-screen">
    <!-- Header -->
    <header
      class="sticky top-0 z-50 flex items-center justify-between px-[5%] py-4 bg-white border-b border-gray-200 shadow-md"
    >
      <router-link
        to="/dosen/dashboard"
        class="text-3xl font-bold no-underline text-primary"
      >
        Latih.in
      </router-link>

      <nav class="items-center hidden gap-10 md:flex">
        <router-link
          to="/dosen/dashboard"
          class="font-semibold no-underline text-medium-text hover:text-primary"
        >
          Dashboard
        </router-link>
        <router-link
          to="/dosen/soal"
          class="font-semibold no-underline text-medium-text hover:text-primary"
        >
          Soal
        </router-link>
        <router-link
          to="/dosen/exam"
          class="font-semibold no-underline text-medium-text hover:text-primary"
        >
          Exam
        </router-link>
      </nav>

      <!-- Profile Dropdown -->
      <div class="relative">
        <div @click="toggleProfileDropdown" class="cursor-pointer">
          <UserCircle
            class="w-10 h-10 text-gray-600 transition-colors hover:text-primary"
          />
        </div>

        <div
          v-if="showProfileDropdown"
          class="absolute top-16 right-0 z-50 w-72 p-2 bg-white rounded-lg shadow-xl border border-gray-100"
        >
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
            :to="`/dosen/profile/${user.id}`"
            class="block w-full px-4 py-2 text-left text-gray-700 rounded hover:bg-gray-100"
          >
            Lihat Profil
          </router-link>

          <a
            href="#"
            @click.prevent="logout"
            class="block w-full px-4 py-2 text-left text-gray-700 rounded hover:bg-gray-100"
          >
            Logout
          </a>
        </div>
      </div>
    </header>

    <!-- Main Layout -->
    <main class="flex gap-6 p-6">
      <!-- Sidebar -->
      <aside class="flex-shrink-0 w-64">
        <div class="p-4 space-y-2 bg-white rounded-lg shadow-md">
          <router-link
            to="/dosen/dashboard"
            class="flex items-center gap-3 px-4 py-2 font-semibold rounded-md transition-colors"
            :class="
              isActive('/dosen/dashboard')
                ? 'text-primary bg-indigo-50'
                : 'text-gray-600 hover:bg-gray-100'
            "
          >
            <i class="fas fa-tachometer-alt"></i> Dashboard
          </router-link>

          <router-link
            to="/dosen/soal"
            class="flex items-center gap-3 px-4 py-2 font-semibold rounded-md transition-colors"
            :class="
              isActive('/dosen/soal')
                ? 'text-primary bg-indigo-50'
                : 'text-gray-600 hover:bg-gray-100'
            "
          >
            <i class="fas fa-file-alt"></i> Soal
          </router-link>

          <router-link
            to="/dosen/exam"
            class="flex items-center gap-3 px-4 py-2 font-semibold rounded-md transition-colors"
            :class="
              isActive('/dosen/exam')
                ? 'text-primary bg-indigo-50'
                : 'text-gray-600 hover:bg-gray-100'
            "
          >
            <i class="fas fa-clipboard-list"></i> Exam
          </router-link>
        </div>
      </aside>

      <!-- Page Content -->
      <div class="flex-1">
        <router-view />
      </div>
    </main>
  </div>
</template>

<script setup>
import { ref, onMounted } from "vue";
import { useRouter, useRoute } from "vue-router";
import { useLocalStorage } from "../hooks/useLocalStorage";
import { UserCircle } from "lucide-vue-next";

const showProfileDropdown = ref(false);
const user = ref(null);
const router = useRouter();
const route = useRoute();

const toggleProfileDropdown = () => {
  showProfileDropdown.value = !showProfileDropdown.value;
};

const logout = () => {
  const { removeValue: removeToken } = useLocalStorage("token");
  const { removeValue: removeUser } = useLocalStorage("user");
  removeToken();
  removeUser();
  router.push("/login");
};

const isActive = (path) => route.path.startsWith(path);

onMounted(() => {
  const { value: storedUser } = useLocalStorage("user");
  if (storedUser.value) {
    user.value = storedUser.value;
  } else {
    logout();
  }
});
</script>
