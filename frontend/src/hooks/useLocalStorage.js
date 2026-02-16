import { ref } from 'vue'

export const useLocalStorage = (key, defaultValue = null) => {
  const stored = localStorage.getItem(key)

  let parsedValue
  try {
    parsedValue = stored ? JSON.parse(stored) : defaultValue
  } catch {
    parsedValue = stored 
  }

  const value = ref(parsedValue)

  const setValue = (newValue) => {
    value.value = newValue

    if (typeof newValue === 'object') {
      localStorage.setItem(key, JSON.stringify(newValue))
    } else {
      localStorage.setItem(key, newValue)
    }
  }

  const removeValue = () => {
    value.value = defaultValue
    localStorage.removeItem(key)
  }

  return { value, setValue, removeValue }
}
