<template>
  <div class="lecturer-layout">
    <header>
      <!-- Navigasi untuk dosen -->
      <nav>
        <router-link to="/dosen/dashboard">Dashboard</router-link>
        <router-link to="/dosen/soal">Soal</router-link>
        <router-link to="/dosen/soal/list">Daftar Soal</router-link>
        <router-link to="/dosen/soal/create">Buat Soal</router-link>
        <router-link to="/dosen/soal/create-manual"
          >Buat Soal Manual</router-link
        >
        <router-link to="/dosen/soal/upload-json">Upload JSON</router-link>
        <router-link to="/dosen/update-profile">Profil</router-link>
      </nav>
      <div class="relative">
        <div @click="toggleProfileDropdown" class="cursor-pointer">
          <img
            :src="userIcon"
            alt="User Profile"
            class="w-10 h-10 rounded-full"
          />
        </div>
        <div
          v-if="showProfileDropdown"
          class="absolute top-16 right-0 z-50 w-72 p-2 bg-white rounded-lg shadow-xl border border-gray-100"
        >
          <div class="flex items-center p-2">
            <img
              :src="userIcon"
              alt="User Profile"
              class="w-12 h-12 rounded-full mr-4"
            />
            <div v-if="user">
              <div class="font-semibold">{{ user.name }}</div>
              <div class="text-sm text-gray-500">{{ user.email }}</div>
            </div>
          </div>
          <hr class="my-2 border-gray-200" />
          <router-link
            to="/dosen/update-profile"
            class="block w-full px-4 py-2 text-left text-gray-700 rounded hover:bg-gray-100"
            >Update Profile</router-link
          >
          <a
            href="#"
            @click.prevent="logout"
            class="block w-full px-4 py-2 text-left text-gray-700 rounded hover:bg-gray-100"
            >Logout</a
          >
        </div>
      </div>
    </header>

    <main>
      <!-- Slot untuk konten anak-anak rute -->
      <slot />
    </main>

    <footer>
      <!-- Footer untuk layout dosen -->
    </footer>
  </div>
</template>

<script>
import userIcon from "../assets/user-icon.png";
import { useLocalStorage } from "../hooks/useLocalStorage";

export default {
  name: "DosenLayout",
};
</script>

<style scoped>
.lecturer-layout {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
}

header {
  background-color: #f8f9fa;
  padding: 1rem;
  border-bottom: 1px solid #dee2e6;
}

nav {
  display: flex;
  gap: 1rem;
}

nav a {
  text-decoration: none;
  color: #007bff;
  padding: 0.5rem;
}

nav a.router-link-exact-active {
  font-weight: bold;
  color: #0056b3;
}

main {
  flex: 1;
  padding: 2rem;
}
</style>
