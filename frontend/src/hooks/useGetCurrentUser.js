import { ref, provide, inject, readonly } from "vue";
import { getUserById } from "../provider/user.provider"; // Pastikan path ini benar

const UserSymbol = Symbol('user')

export const provideUser = () => {
  const user = ref(null)
  const loading = ref(false)
  const error = ref(null)
  
  const setUser = (newUser) => {
    user.value = newUser
  }
  
  const clearUser = () => {
    user.value = null
  }

  // --- FUNGSI BARU: Fetch User dengan Validasi Token ---
  const fetchUser = async () => {
    const token = localStorage.getItem('token');
    const userId = localStorage.getItem('id');

    // 1. CEK TOKEN & ID (SOLUSI MASALAH ANDA)
    // Jika tidak ada token, JANGAN panggil API. Stop di sini.
    if (!token || !userId) {
       user.value = null;
       return; 
    }

    loading.value = true;
    try {
       // Panggil API hanya jika token ada
       const res = await getUserById(userId);
       // Sesuaikan dengan struktur response (res.data atau res)
       user.value = res.data || res; 
    } catch (err) {
       console.error("Gagal mengambil data user:", err);
       error.value = err;
       // Jika error 401, biasanya api.handler sudah handle logout
       user.value = null;
    } finally {
       loading.value = false;
    }
  }
  
  provide(UserSymbol, {
    user: readonly(user),
    loading: readonly(loading),
    error: readonly(error),
    setUser,
    clearUser,
    fetchUser // <-- Expose fungsi ini agar bisa dipanggil di App.vue
  })
  
  return {
    user,
    loading,
    error,
    setUser,
    clearUser,
    fetchUser
  }
}

export const useUser = () => {
  const context = inject(UserSymbol)
  
  if (!context) {
    throw new Error('useUser must be used within a component where provideUser has been called')
  }
  
  return context
}

export const useGetCurrentUser = () => {
  // Return semua yang dibutuhkan komponen
  const context = useUser();
  return context;
}