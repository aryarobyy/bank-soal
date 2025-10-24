<template>
  <div class="font-sans bg-background min-h-screen flex flex-col">
    <Navbar />

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
      <Toast ref="toastRef" />
    </main>
  </div>
</template>

<script setup>
// CHANGE IS HERE: Use <script setup> instead of <script>

import { useGetCurrentUser } from '../../hooks/useGetCurrentUser';
import { onMounted, ref } from 'vue';
import Toast from '../../components/utils/Toast.vue';
import { useRouter } from 'vue-router';
import Navbar from '../../components/ui/Navbar.vue';

const { user } = useGetCurrentUser()
const toastRef = ref(null);
const router = useRouter()

onMounted(() => {
  // Now this code works correctly because it's inside the setup context
  if (!user.value){
    toastRef.value.showToast(
      "error",
      "User tidak dikenali",
      "Anda harus login terlebih dahulu"
    );
    router.push("/login")
  }
})

const startExam = () => {
  // Add your logic to start the exam here
  alert("Memulai ujian!");
};
</script>