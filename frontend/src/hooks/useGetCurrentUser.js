import { ref, provide, inject, readonly } from "vue";
import { getProfile } from "../provider/user.provider";
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

    if (!token) {
      user.value = null;
      return; 
    }

    try {
      const userStr = localStorage.getItem('user');
      if (userStr && userStr !== 'undefined' && userStr !== 'null') {
        const parsed = JSON.parse(userStr);
        if (parsed && parsed.id) {
          user.value = parsed;
          return; 
        }
      }
    } catch (e) {
      console.warn("Gagal parse user dari localStorage:", e);
    }

    loading.value = true;
    try {
      const res = await getProfile();
      const userData = res.data?.data || res.data || res;
      user.value = userData;
      localStorage.setItem('user', JSON.stringify(userData));
    } catch (err) {
      console.error("Gagal mengambil data user:", err);
      error.value = err;
      user.value = null;
      localStorage.clear();
      router.push('/landing');
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