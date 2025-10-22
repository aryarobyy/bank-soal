<template>
  <div class="lecturer-layout">
    <header>
      <!-- Navigasi untuk dosen -->
      <nav>
        <router-link to="/dosen/dashboard">Dashboard</router-link>
        <router-link to="/dosen/soal">Soal</router-link>
        <router-link to="/dosen/soal/list">Daftar Soal</router-link>
        <router-link to="/dosen/soal/create">Buat Soal</router-link>
        <router-link to="/dosen/soal/create-manual">Buat Soal Manual</router-link>
        <router-link to="/dosen/soal/upload-json">Upload JSON</router-link>
        <router-link to="/dosen/update-profile">Profil</router-link>
      </nav>
    </header>
    
    <main>
      <!-- Slot untuk konten anak-anak rute -->
      <slot />
    </main>
    
    <footer>
      <!-- Footer untuk layout dosen -->
    </footer>
  </div>
</template>

<script>
export default {
  name: 'DosenLayout'
}
</script>

<style scoped>
.lecturer-layout {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
}

header {
  background-color: #f8f9fa;
  padding: 1rem;
  border-bottom: 1px solid #dee2e6;
}

nav {
  display: flex;
  gap: 1rem;
}

nav a {
  text-decoration: none;
  color: #007bff;
  padding: 0.5rem;
}

nav a.router-link-exact-active {
  font-weight: bold;
  color: #0056b3;
}

main {
  flex: 1;
  padding: 2rem;
}
</style>