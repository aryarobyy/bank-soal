import { ref, provide, inject, readonly } from "vue";
import { getUserById } from "../provider/user.provider";
import { useRouter } from "vue-router";

const UserSymbol = Symbol('user')

export const provideUser = () => {
  const user = ref(null)
  const loading = ref(false)
  const error = ref(null)
  const router = useRouter()
  
  const setUser = (newUser) => {
    user.value = newUser
  }
  
  const clearUser = () => {
    user.value = null
  }

  const fetchUser = async () => {
    const token = localStorage.getItem('token');
    const userId = localStorage.getItem('id');

    if (!token || !userId) {
      user.value = null;
      router.push('/login');
      return; 
    }

    loading.value = true;
    try {
      const res = await getUserById(userId);
      user.value = res.data || res; 
    } catch (err) {
      console.error("Gagal mengambil data user:", err);
      error.value = err;
      user.value = null;
      router.push('/login');
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
    fetchUser 
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
  const context = useUser();
  return context;
}