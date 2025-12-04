<template>
  <div class="font-sans bg-gray-50 text-slate-800 overflow-x-hidden">
    <Navbar />

    <section class="relative pt-40 pb-24 text-center bg-white overflow-hidden">
      <div class="absolute top-0 left-0 w-full h-full overflow-hidden z-0">
        <div class="absolute -top-24 -right-24 w-96 h-96 rounded-full bg-blue-100 blur-3xl opacity-50"></div>
        <div class="absolute top-40 -left-24 w-72 h-72 rounded-full bg-cyan-100 blur-3xl opacity-50"></div>
      </div>

      <div class="relative z-10 max-w-5xl mx-auto px-4 reveal">
        <span class="inline-block py-1 px-3 rounded-full bg-blue-50 text-blue-600 text-sm font-semibold mb-4 border border-blue-100">
          #DigitalisasiInformatika
        </span>
        <h1 class="text-5xl md:text-7xl font-extrabold text-slate-900 leading-tight mb-6">
          Seleksi Mahasiswa Pindahan <br />
          <span class="text-transparent bg-clip-text bg-gradient-to-r from-blue-600 to-cyan-500">
            Lebih Cerdas & Efisien
          </span>
        </h1>
        <p class="text-lg md:text-xl text-slate-600 max-w-2xl mx-auto mb-10 leading-relaxed">
          Platform terintegrasi untuk manajemen bank soal dan ujian online Prodi Informatika. Objektif, transparan, dan terorganisir.
        </p>
        
        <div class="flex flex-col sm:flex-row justify-center gap-4">
          <router-link to="/login" class="px-8 py-3.5 rounded-full bg-blue-600 text-white font-bold text-lg hover:bg-blue-700 transition shadow-lg hover:shadow-blue-500/30 transform hover:-translate-y-1">
            Mulai Sekarang
          </router-link>
          <a href="#about" class="px-8 py-3.5 rounded-full bg-white text-slate-700 font-bold text-lg border border-slate-200 hover:bg-gray-50 transition">
            Pelajari Lebih Lanjut
          </a>
        </div>
      </div>
    </section>

    <section id="about" class="py-20 bg-white">
      <div class="max-w-6xl mx-auto px-6">
        <div class="flex flex-col md:flex-row items-center gap-16">
          <div class="flex-1 reveal fade-left">
            <div class="relative">
              <div class="absolute inset-0 bg-blue-600 rounded-3xl rotate-3 opacity-10"></div>
              <img 
                :src="illustration1" 
                alt="Ilustrasi Edukasi" 
                loading="lazy"
                class="relative w-full max-w-md mx-auto z-10 drop-shadow-xl hover:scale-105 transition duration-500" 
              />
            </div>
          </div>
          <div class="flex-1 reveal fade-right">
            <h2 class="text-3xl md:text-4xl font-bold text-slate-900 mb-6 border-l-4 border-blue-500 pl-4">
              Apa itu Latih.in?
            </h2>
            <p class="text-lg text-slate-600 leading-loose mb-6">
              Latih.in hadir sebagai solusi digitalisasi seleksi akademik untuk Prodi Informatika. Platform ini mengintegrasikan 
              <span class="font-semibold text-blue-600">manajemen bank soal</span> dan 
              <span class="font-semibold text-blue-600">tes online</span> dalam satu sistem terpadu, 
              memastikan proses seleksi mahasiswa pindahan berjalan lebih efisien, objektif, dan terorganisir.
            </p>
            <ul class="space-y-3 text-slate-600">
              <li class="flex items-center gap-2">
                <span class="w-2 h-2 rounded-full bg-blue-500"></span> Terintegrasi Penuh
              </li>
              <li class="flex items-center gap-2">
                <span class="w-2 h-2 rounded-full bg-blue-500"></span> Objektivitas Tinggi
              </li>
            </ul>
          </div>
        </div>
      </div>
    </section>

    <section class="py-20 bg-slate-50">
      <div class="max-w-6xl mx-auto px-6">
        <div class="flex flex-col md:flex-row-reverse items-center gap-16">
          <div class="flex-1 reveal fade-right">
            <div class="relative">
              <div class="absolute inset-0 bg-cyan-500 rounded-full blur-3xl opacity-20"></div>
              <img 
                :src="illustration2" 
                alt="Ilustrasi Belajar Online" 
                loading="lazy"
                class="relative w-full max-w-md mx-auto z-10 drop-shadow-2xl hover:rotate-1 transition duration-500" 
              />
            </div>
          </div>
          <div class="flex-1 reveal fade-left">
            <h2 class="text-3xl md:text-4xl font-bold text-slate-900 mb-6 border-l-4 border-cyan-500 pl-4">
              Keunggulan & Aksesibilitas
            </h2>
            <p class="text-lg text-slate-600 leading-loose mb-8">
              Persiapkan diri Anda dengan materi uji yang terstruktur. Sistem kami mendukung penilaian otomatis dan kategorisasi tingkat kesulitan soal, memberikan umpan balik hasil ujian secara instan tanpa perlu menunggu proses manual yang lama.
            </p>
            
            <div class="bg-white p-5 rounded-xl shadow-sm border border-slate-100 flex items-start gap-4 hover:shadow-md transition">
              <div class="p-3 bg-blue-50 rounded-lg text-blue-600">
                <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z" />
                </svg>
              </div>
              <div>
                <h4 class="font-bold text-slate-800">Performa Tinggi</h4>
                <p class="text-sm text-slate-500 mt-1">
                  Antarmuka responsif dan cepat berkat teknologi Vue.js dan backend Golang, mendukung simulasi ujian tanpa hambatan (Lag-free).
                </p>
              </div>
            </div>

          </div>
        </div>
      </div>
    </section>

    <section class="py-16 bg-blue-600 text-white text-center reveal">
      <div class="max-w-4xl mx-auto px-6">
        <h2 class="text-3xl font-bold mb-6">Siap Menghadapi Seleksi?</h2>
        <p class="text-blue-100 text-lg mb-8">Bergabunglah sekarang dan rasakan pengalaman ujian yang modern.</p>
        <router-link to="/login" class="inline-block px-8 py-3 bg-white text-blue-600 font-bold rounded-lg shadow-lg hover:bg-gray-100 transition">
          Masuk ke Portal
        </router-link>
      </div>
    </section>

    <footer class="bg-slate-900 text-center py-16">
      <div class="max-w-6xl mx-auto px-6">
        <div class="text-4xl font-bold text-transparent bg-clip-text bg-gradient-to-r from-blue-400 to-cyan-400 mb-4">
          Latih.in
        </div>
        <p class="text-slate-400 text-lg mb-8">Mari belajar dan berkembang bersama Latih.in</p>
        <div class="h-px w-24 bg-slate-700 mx-auto mb-8"></div>
        <small class="block text-sm text-slate-500">
          &copy; 2025 Latih.in — All Rights Reserved. <br>
          <span class="opacity-50">Developed for Informatics Engineering Department.</span>
        </small>
      </div>
    </footer>
  </div>
</template>

<script setup>
import { onMounted } from 'vue';

import illustration1 from '../../assets/illustration-1.jpg';
import illustration2 from '../../assets/illustration-2.jpg';

onMounted(() => {
  const observerOptions = {
    threshold: 0.1, 
    rootMargin: "0px 0px -50px 0px"
  };

  const observer = new IntersectionObserver((entries) => {
    entries.forEach((entry) => {
      if (entry.isIntersecting) {
        entry.target.classList.add('active');
        
      }
    });
  }, observerOptions);

  const hiddenElements = document.querySelectorAll('.reveal');
  hiddenElements.forEach((el) => observer.observe(el));
});
</script>

<style scoped>

.reveal {
  opacity: 0;
  transform: translateY(30px);
  transition: all 0.8s ease-out;
}

.reveal.active {
  opacity: 1;
  transform: translateY(0);
}

.fade-left {
  transform: translateX(-30px);
}
.reveal.active.fade-left {
  transform: translateX(0);
}

.fade-right {
  transform: translateX(30px);
}
.reveal.active.fade-right {
  transform: translateX(0);
}
</style>