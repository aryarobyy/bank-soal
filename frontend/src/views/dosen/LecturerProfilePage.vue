<template>
  <div class="font-sans bg-background min-h-screen">
    <header class="sticky top-0 z-50 flex items-center justify-between px-[5%] py-4 bg-white">
       <router-link to="/" class="text-3xl font-bold no-underline text-primary">Latih.in</router-link>
    </header>

    <main class="p-4 sm:p-8">
       <UserProfileForm 
        v-if="lecturerData"
        :profile-data="lecturerData" 
        user-role="dosen"
        @save="handleProfileSave"
      />
    </main>
  </div>
</template>

<script>

import UserProfileForm from '../../components/UserProfileForm.vue';

export default {
  name: 'LecturerProfilePage',
  components: { UserProfileForm },
  data() {
    return {
      lecturerData: null,
    };
  },
  methods: {
    fetchLecturerData() {
     
      this.lecturerData = {
        fullName: 'Dr. Budi Santoso',
        email: 'budi.s@latih.in',
        idNumber: '198503152010011005', 
        gender: 'Male',
        history: 'Spesialisasi Machine Learning',
        imageUrl: null,
      };
    },
    handleProfileSave(updatedData) {
      console.log('Menyimpan data dosen:', updatedData);
     
    }
  },
  created() {
    this.fetchLecturerData();
  }
};
</script>