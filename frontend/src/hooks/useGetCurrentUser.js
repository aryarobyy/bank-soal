import { ref, provide, inject, readonly } from "vue";

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
  
  provide(UserSymbol, {
    user: readonly(user),
    loading: readonly(loading),
    error: readonly(error),
    setUser,
    clearUser
  })
  
  return {
    user,
    loading,
    error,
    setUser,
    clearUser
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
  const { user, loading, error } = useUser()
  
  return {
    user,
    loading,
    error
  }
}