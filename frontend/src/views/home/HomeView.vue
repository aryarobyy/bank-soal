<template>
  <div class="font-sans bg-white text-dark-text">
    <Navbar />

    <section class="pt-32 pb-16 text-center bg-background">
      <h1 class="text-5xl md:text-6xl font-bold text-dark-text">
        Test Seleksi Mahasiswa Pindahan <br />
        <span class="text-primary">Latih.in</span>
      </h1>
      <hr class="w-1/2 mx-auto my-8 border-gray-300" />
    </section>

    <section class="max-w-6xl mx-auto px-4 py-16">
      <h2 class="text-4xl font-bold text-center mb-12">Apa itu Latih.in?</h2>
      <div class="flex flex-col md:flex-row items-center gap-12">
        <div class="flex-1">
          <img 
            :src="illustration1" 
            alt="Ilustrasi Edukasi" 
            loading="lazy"
            class="w-full h-auto max-w-sm mx-auto" 
          />
        </div>
        <div class="flex-1">
          <p class="text-lg leading-relaxed text-medium-text">
            Latih.in adalah platform latihan soal dan ujian yang dirancang untuk membantu mahasiswa mempersiapkan diri menghadapi tes seleksi dengan cara yang lebih efisien dan menyenangkan.
          </p>
        </div>
      </div>
    </section>

    <section class="max-w-6xl mx-auto px-4 py-16">
      <div class="flex flex-col md:flex-row-reverse items-center gap-12">
        <div class="flex-1">
          <img 
            :src="illustration2" 
            alt="Ilustrasi Belajar Online" 
            loading="lazy"
            class="w-full h-auto max-w-sm mx-auto" 
          />
        </div>
        <div class="flex-1">
          <p class="text-lg leading-relaxed text-medium-text">
            Kamu bisa belajar kapan saja dan di mana saja, karena Latih.in dapat diakses melalui perangkat apa pun. Nikmati pengalaman belajar interaktif dengan video dan animasi menarik.
          </p>
        </div>
      </div>
    </section>

    <footer class="bg-[#A7DDFF] text-center py-16 mt-16">
      <div class="text-4xl font-bold text-primary">Latih.in</div>
      <p class="mt-4 text-lg text-dark-text">Mari belajar bersama Latih.in</p>
      <small class="block mt-8 text-sm text-medium-text opacity-75">
        &copy; 2025 Latih.in — All Rights Reserved
      </small>
    </footer>
  </div>
</template>

<script setup>

import illustration1 from '../../assets/illustration-1.jpg';
import illustration2 from '../../assets/illustration-2.jpg';


</script>