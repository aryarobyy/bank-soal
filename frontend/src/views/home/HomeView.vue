<template>
  <div class="font-sans bg-white text-dark-text">
    
    <div class="relative bg-background">
      
      <header class="absolute top-0 left-0 right-0 z-10 flex items-center justify-between px-[5%] py-6">
        <router-link to="/" class="text-3xl font-bold no-underline text-primary">Latih.in</router-link>
        <nav class="items-center hidden gap-10 md:flex">
          <router-link to="/" class="font-semibold no-underline text-primary">Home</router-link>
          <router-link to="/ujian" class="font-semibold no-underline text-medium-text hover:text-primary">Ujian</router-link>
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

      <main>
        <section 
          class="pt-32 pb-16 text-center" 
          :style="{ backgroundImage: `url('${heroBgPattern}')` }"
        >
          <h1 class="text-6xl font-bold text-dark-text">
            Test Seleksi Mahasiswa Pindahan  <br />
            <span class="text-primary">Latih.in</span>
          </h1>
          <hr class="w-1/2 mx-auto my-8 border-gray-300" />
        </section>
      </main>
    </div>

    <section class="max-w-6xl px-4 py-16 mx-auto">
      <h2 class="mb-12 text-4xl font-bold text-center">Apa itu Latih.in?</h2>
      <div class="flex flex-col items-center gap-12 md:flex-row">
        <div class="flex-1"><img :src="illustration1" alt="E-learning Illustration" class="w-full h-auto max-w-sm mx-auto" /></div>
        <div class="flex-1">
          <p class="text-lg leading-relaxed text-medium-text">Latih.in blablablablablablablablablabla</p>
        </div>
      </div>
    </section>

    <section class="max-w-6xl px-4 py-16 mx-auto">
      <div class="flex flex-col items-center gap-12 md:flex-row-reverse">
        <div class="flex-1"><img :src="illustration2" alt="Video Learning Illustration" class="w-full h-auto max-w-sm mx-auto" /></div>
        <div class="flex-1">
          <p class="text-lg leading-relaxed text-medium-text">
            Latih.in dapat diakses kapan saja dan dimana saja melalui perangkat apapun. Kami memberikan video pembelajaran yang dikemas dengan animasi dan interaksi yang menarik untuk membuat belajar jadi lebih menyenangkan.
          </p>
        </div>
      </div>
    </section>

    <div class="relative">
      <div class="absolute top-0 w-full -mt-px">
        <svg viewBox="0 0 1440 120" fill="none" xmlns="http://www.w3.org/2000/svg">
          <path d="M1440 29.6135V120H0V29.6135C240 5.86981 480 0 720 0C960 0 1200 5.86981 1440 29.6135Z" fill="#A7DDFF"/>
        </svg>
      </div>
      <footer class="relative pt-32 pb-8 text-center bg-[#A7DDFF]">
        <div class="text-4xl font-bold text-primary">Latih.in</div>
        <p class="mt-4 text-lg text-dark-text">Mari Belajar bersama Latih.in</p>
        <small class="block mt-8 text-sm opacity-75 text-medium-text">&copy; 2025 - Latih.in - All Rights Reserved</small>
      </footer>
    </div>
  </div>
</template>

<script>
// Path gambar-gambar Anda yang sudah benar
import userIcon from '../../assets/user-icon.png';
import illustration1 from '../../assets/illustration-1.jpg';
import illustration2 from '../../assets/illustration-2.jpg';

// 1. Baris 'import heroBgPattern' yang bermasalah sudah DIHAPUS dari sini.

export default {
  name: 'HomePage',
  data() {
    return {
      isLoggedIn: false,
      showProfileDropdown: false,
      user: null,
      userIcon,
      illustration1,
      illustration2,
      // 2. Data SVG dipindahkan langsung ke sini sebagai string.
      heroBgPattern: "data:image/svg+xml,%3Csvg width='60' height='60' viewBox='0 0 60 60' xmlns='http://www.w3.org/2000/svg'%3E%3Cg fill='none' fill-rule='evenodd'%3E%3Cg fill='%23d3d8fc' fill-opacity='0.4'%3E%3Cpath d='M36 34v-4h-2v4h-4v2h4v4h2v-4h4v-2h-4zm0-30V0h-2v4h-4v2h4v4h2V6h4V4h-4zM6 34v-4H4v4H0v2h4v4h2v-4h4v-2H6zM6 4V0H4v4H0v2h4v4h2V6h4V4H6z'/%3E%3C/g%3E%3C/g%3E%3C/svg%3E",
    };
  },
  methods: {
    toggleProfileDropdown() { this.showProfileDropdown = !this.showProfileDropdown; },
    simulateLogin() {
      this.isLoggedIn = true;
      this.user = { fullName: 'Hendro Paulus', studentID: 'G1A023091' };
    },
    logout() {
      localStorage.removeItem('isLoggedIn');
      this.isLoggedIn = false;
      this.user = null;
      this.showProfileDropdown = false;
    },
  },
  mounted() {
    if (localStorage.getItem('isLoggedIn') === 'true') {
      this.simulateLogin();
    }
  },
};
</script>