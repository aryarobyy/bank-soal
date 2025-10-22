<template>
  <div class="font-sans bg-background min-h-screen flex flex-col">
    <header class="sticky top-0 z-50 flex items-center justify-between px-[5%] py-4 bg-white">
      <router-link to="/" class="text-3xl font-bold no-underline text-primary">Latih.in</router-link>
      <nav class="items-center hidden gap-10 md:flex">
        <router-link to="/" class="font-semibold no-underline text-medium-text hover:text-primary">Home</router-link>
        <router-link to="/ujian" class="font-semibold no-underline text-primary">Ujian</router-link>
      </nav>
      <div class="relative">
        <div @click="toggleProfileDropdown" class="cursor-pointer">
          <img :src="userIcon" alt="User Profile" class="w-10 h-10 rounded-full" />
        </div>
        <div v-if="showProfileDropdown" class="absolute top-16 right-0 z-50 w-72 p-2 bg-white rounded-lg shadow-xl border border-gray-100">
          <template v-if="isLoggedIn">
            <div class="flex items-center p-2">
              <img :src="userIcon" alt="User Profile" class="w-12 h-12 rounded-full mr-4" />
              <div>
                <div class="font-semibold">{{ user.fullName }}</div>
                <div class="text-sm text-gray-500">{{ user.studentID }}</div>
              </div>
            </div>
            <hr class="my-2 border-gray-200" />
            <a href="#" class="block w-full px-4 py-2 text-left text-gray-700 rounded hover:bg-gray-100">My Exam</a>
            <router-link to="/update-profile" class="block w-full px-4 py-2 text-left text-gray-700 rounded hover:bg-gray-100">Update Profile</router-link>
            <a href="#" @click="logout" class="block w-full px-4 py-2 text-left text-gray-700 rounded hover:bg-gray-100">Logout</a>
          </template>
          <template v-else>
            <router-link to="/login" class="block w-full px-4 py-2 font-semibold text-center text-gray-700 rounded hover:bg-gray-100">Login</router-link>
          </template>
        </div>
      </div>
    </header>

    <main class="flex-grow flex items-center justify-center p-4 sm:p-8">
      <div class="w-full max-w-5xl p-6 sm:p-10 mx-auto bg-white rounded-2xl shadow-xl">
        
        <div class="p-8 bg-gray-50 border border-gray-200 rounded-lg">
          <div class="pb-6 mb-6 text-center border-b border-gray-200">
            <h3 class="mb-2 text-2xl font-bold text-dark-text">Ujian Seleksi Pindahan 2025</h3>
            <p class="text-medium-text">10 Oktober 2025 09.00 - 10.30 WIB</p>
            <p class="font-semibold text-gray-700">Durasi Ujian : 120 Menit</p>
          </div>
          
          <div>
            <h4 class="mb-4 text-lg font-bold">Aturan Ujian :</h4>
            <ul class="space-y-3 text-medium-text">
              <li class="flex items-start"><span class="mr-3 font-bold text-primary">&gt;</span>Tidak Di Perbolehkan Untuk Keluar dari Ujian Kecuali Selesai/Waktu Ujian Telah Habis</li>
              <li class="flex items-start"><span class="mr-3 font-bold text-primary">&gt;</span>Timer Akan Tetap Berjalan Internet Anda Terputus</li>
            </ul>
          </div>
        </div>

        <div class="flex justify-end mt-8">
          <button @click="startExam" class="px-5 py-2 text-sm font-semibold text-white transition-opacity rounded-md bg-exam-green hover:opacity-90">
            Mulai Ujian
          </button>
        </div>
      </div>
    </main>
  </div>
</template>

<script>
import { User } from 'lucide-vue-next';

// Placeholder image using SVG data URI as fallback
const userIcon = "data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' width='100' height='100' viewBox='0 0 24 24' fill='none' stroke='%23333' stroke-width='2' stroke-linecap='round' stroke-linejoin='round'%3E%3Cpath d='M19 21v-2a4 4 0 0 0-4-4H9a4 4 0 0 0-4 4v2'%3E%3C/path%3E%3Ccircle cx='12' cy='7' r='4'%3E%3C/circle%3E%3C/svg%3E";

export default {
  name: 'ExamPage',
  components: {
    User
  },
  data() {
    return {
      isLoggedIn: false,
      showProfileDropdown: false,
      user: null,
      userIcon,
    };
  },
  methods: {
    toggleProfileDropdown() {
      this.showProfileDropdown = !this.showProfileDropdown;
    },
    simulateLogin() {
      this.isLoggedIn = true;
      this.user = {
        fullName: 'Nama',
      };
    },
    logout() {
      localStorage.removeItem('isLoggedIn');
      this.isLoggedIn = false;
      this.user = null;
      this.showProfileDropdown = false;
      // Opsional: arahkan ke halaman utama setelah logout
      this.$router.push('/');
    },
    startExam() {
      alert('Memulai ujian...');
    }
  },
  mounted() {
    // === LOGIKA PENGECEKAN LOGIN DITAMBAHKAN DI SINI ===
    if (localStorage.getItem('isLoggedIn') !== 'true') {
      // Jika user BELUM login
      alert('Anda harus login terlebih dahulu');
      this.$router.push('/login'); // Arahkan ke halaman login
    } else {
      // Jika user SUDAH login, tampilkan data profil di navbar
      this.simulateLogin();
    }
  }
};
</script>