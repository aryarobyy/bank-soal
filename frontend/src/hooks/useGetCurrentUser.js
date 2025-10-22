import { ref, computed } from "vue";
import { useLocalStorage } from "./useLocalStorage";

export const useGetCurrentUser = () => {
  const { value, setValue } = useLocalStorage('user')
  
  const loading = ref(false)
  const error = ref(null)
  
  const user = computed(() => value.value)
  
  return {
    user,
    loading,
    error
  }
}