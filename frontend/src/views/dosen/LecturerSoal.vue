<template>
  <SoalManagement
    :soal-data="lecturerSoalData"
    @buat-soal="handleBuatSoal"
    @view-details="handleViewDetails"
    @delete-soal="handleDeleteSoal"
  />
</template>

<script>
import SoalManagement from './SoalManagement.vue';

export default {
  name: 'LecturerSoal',
  components: { SoalManagement },
  // Script sekarang hanya berisi logika untuk halaman ini saja
  data() {
    return {
      lecturerSoalData: [
        {
          title: 'Soal 1',
          items: [ { id: 101, nama: 'Create a User Persona Based on Research Data.', status: 'Not Started' } ]
        },
        {
          title: 'Soal 2',
          items: [ { id: 102, nama: 'Design a Landing Page for a New App.', status: 'In Progress' } ]
        },
        {
          title: 'Module 1',
          items: [ { id: 103, nama: 'Fundamentals of UI/UX.', status: 'Completed' } ]
        }
      ]
    };
  },
  methods: {
    handleBuatSoal() {
      this.$router.push('/dosen/soal/create');
    },
    handleViewDetails(soalId) {
      // Mengarahkan ke halaman daftar detail soal
      this.$router.push('/dosen/soal/list');
    },
    handleDeleteSoal(soalId) {
      if (confirm(`Yakin ingin menghapus soal ID: ${soalId}?`)) {
        alert('Soal dihapus!');
      }
    }
  }
};
</script>