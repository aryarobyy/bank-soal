<template>
  <div class="font-sans bg-background min-h-screen">
    <header class="sticky top-0 z-50 flex items-center justify-between px-[5%] py-4 bg-white">
      <router-link to="/" class="text-3xl font-bold no-underline text-primary">Latih.in</router-link>
    </header>

    <main class="p-4 sm:p-8">
      <UserProfileForm 
        v-if="studentData"
        :profile-data="studentData" 
        user-role="mahasiswa"
        @save="handleProfileSave"
      />
    </main>
  </div>
</template>

<script>
// Path yang benar adalah ../
import UserProfileForm from '../components/UserProfileForm.vue';

export default {
  name: 'StudentProfilePage',
  components: { UserProfileForm },
  data() {
    return {
      studentData: null,
    };
  },
  methods: {
    fetchStudentData() {
      this.studentData = {
        fullName: 'Alexa Rawles',
        email: 'alexarawles@gmail.com',
        idNumber: 'G1A023091',
        gender: 'Female',
        history: 'Lulusan terbaik 2023',
        imageUrl: null,
      };
    },
    handleProfileSave(updatedData) {
      console.log('Menyimpan data mahasiswa:', updatedData);
    }
  },
  created() {
    this.fetchStudentData();
  }
};
</script>