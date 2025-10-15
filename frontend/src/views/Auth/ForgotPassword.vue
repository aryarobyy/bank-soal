
<template>
  <div class="page-container">
    <div class="forgot-password-card">
      <router-link to="/login" class="back-arrow">&#8592;</router-link>

      <h2 class="title">FORGOT PASSWORD</h2>

      <form @submit.prevent="handleConfirm" class="form-wrapper">
        <div class="form-group">
          <label for="email">Email</label>
          <input
            id="email"
            v-model="email"
            type="email"
            placeholder="Email"
            required
          />
        </div>

        <button type="submit" class="btn-confirm">Confirm</button>
      </form>
    </div>
  </div>
</template>
<script>
export default {
  name: 'ForgotPassword',
  data() {
    return {
      email: '',
    };
  },
  methods: {
    handleConfirm() {
      // Logika untuk mengirim email reset password
      // Untuk sekarang, kita hanya menampilkan alert
      if (this.email) {
        alert(`Tautan reset password telah dikirim ke: ${this.email}`);
        // Di aplikasi nyata, di sini Anda akan memanggil API backend
      } else {
        alert('Silakan masukkan email Anda.');
      }
    },
  },
};
</script>

<style scoped>
/* Menggunakan font dari Google Fonts untuk tampilan yang lebih modern */
@import url('https://fonts.googleapis.com/css2?family=Poppins:wght@400;600;700&display=swap');

.page-container {
  font-family: 'Poppins', sans-serif;
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 100vh;
  background-color: #e3e8ff;
  width: 100%;
}

/* KODE DI BAWAH INI SUDAH DISAMAKAN DENGAN HALAMAN LOGIN */
.forgot-password-card {
  position: relative;
  background: #ffffff;
  padding: 40px 50px; /* <-- Diubah agar sama dengan login */
  border-radius: 15px; /* <-- Diubah agar sama dengan login */
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.1); /* <-- Diubah agar sama dengan login */
  width: 400px; /* <-- Diubah agar sama dengan login */
  text-align: center;
  box-sizing: border-box;
}

.back-arrow {
  position: absolute;
  top: 25px;
  left: 25px;
  font-size: 24px;
  color: #333;
  text-decoration: none;
}

.title {
  font-size: 24px;
  font-weight: 700;
  color: #1a1a1a;
  margin-top: 20px;
  margin-bottom: 35px;
  letter-spacing: 1px;
}

.form-wrapper {
  width: 100%;
}

.form-group {
  text-align: left;
  margin-bottom: 30px;
}

label {
  display: block;
  font-size: 14px;
  color: #555;
  margin-bottom: 8px;
  font-weight: 600;
}

input {
  width: 100%;
  padding: 12px 15px;
  border: 1px solid #ddd;
  border-radius: 8px;
  font-size: 16px;
  box-sizing: border-box;
  transition: border-color 0.3s;
}

input:focus {
  outline: none;
  border-color: #7b68ee;
}

.btn-confirm {
  width: 100%;
  padding: 14px;
  background-color: #7b68ee;
  color: white;
  border: none;
  border-radius: 8px;
  font-size: 16px;
  font-weight: 600;
  cursor: pointer;
  transition: background-color 0.3s ease;
}

.btn-confirm:hover {
  background-color: #6a5acd;
}
</style>