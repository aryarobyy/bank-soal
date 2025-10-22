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

<script>
import { User } from 'lucide-vue-next';
import { useGetCurrentUser } from '../../hooks/useGetCurrentUser';
import { onMounted, ref } from 'vue';
import Toast from '../../components/utils/Toast.vue';
import { useRouter } from 'vue-router';
import Navbar from '../../components/ui/Navbar.vue';

const userIcon = "data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' width='100' height='100' viewBox='0 0 24 24' fill='none' stroke='%23333' stroke-width='2' stroke-linecap='round' stroke-linejoin='round'%3E%3Cpath d='M19 21v-2a4 4 0 0 0-4-4H9a4 4 0 0 0-4 4v2'%3E%3C/path%3E%3Ccircle cx='12' cy='7' r='4'%3E%3C/circle%3E%3C/svg%3E";

const { user, loading, error } = useGetCurrentUser()

const toastRef = ref(null);
const router = useRouter()

onMounted(() => {
  if (!user){
    toastRef.value.showToast(
      "error",
      "User tidak dikenali",
      "Anda harus login terlebih dahulu"
    );
    router.push("/login")
  }
})

</script>