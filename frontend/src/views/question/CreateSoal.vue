<template>
  <div class="w-full p-6 sm:p-8 mx-auto bg-white rounded-2xl shadow-xl">
    <h2 class="text-2xl font-bold text-dark-text mb-8">Manajemen Soal Ujian</h2>
    
    <div class="p-8 border border-gray-200 rounded-lg">
      <p class="mb-6 font-semibold text-medium-text">Pilih Salah Satu Untuk Membuat Soal Ujian</p>
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <button 
          @click="uploadJson" 
          class="w-full h-24 text-lg font-bold text-white transition-colors rounded-lg bg-blue-600 hover:bg-blue-700">
          UPLOAD FILE JSON
        </button>
        <button 
          @click="createManual" 
          class="w-full h-24 text-lg font-bold text-white transition-colors rounded-lg bg-blue-600 hover:bg-blue-700">
          Buat Manual
        </button>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'CreateSoal',
  
 
  computed: {
    isAdminRoute() {
      return this.$route.path.startsWith('/admin/soal');
    }
  },

  methods: {

    uploadJson() {
   
      const routeName = this.isAdminRoute ? 'AdminSoalUploadJson' : 'DosenSoalUploadJson';
      this.$router.push({ name: routeName });
    },
    

    createManual() {
  
      const routeName = this.isAdminRoute ? 'AdminSoalCreateManual' : 'DosenSoalCreateManual';
      this.$router.push({ name: routeName });
    }
  },
};
</script>