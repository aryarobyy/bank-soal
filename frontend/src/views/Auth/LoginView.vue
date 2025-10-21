<template>
  <div
    class="min-h-screen bg-gradient-to-br from-blue-50 to-indigo-100 flex items-center justify-center p-4"
  >
    <div class="bg-white rounded-2xl shadow-xl w-full max-w-md p-8">
      <div class="text-center mb-8">
        <div
          class="inline-flex items-center justify-center w-16 h-16 bg-indigo-600 rounded-full mb-4"
        >
          <GraduationCap class="w-8 h-8 text-white" />
        </div>
        <h1 class="text-3xl font-bold text-gray-800 mb-2">Login</h1>
        <p class="text-gray-600">Masuk ke akun Anda</p>
      </div>

      <form @submit.prevent="handleSubmit" class="space-y-5">
        <div v-for="field in fields" :key="field.name">
          <label class="block text-sm font-medium text-gray-700 mb-2">{{
            field.label
          }}</label>
          <div class="relative">
            <Input
              v-model="formData[field.name]"
              :title="field.title"
              :place-holder="field.placeholder"
              :required="true"
              :id="field.id"
              :icon="field.icon"
              :type="field.type"
            />
          </div>
          <p v-if="errors[field.name]" class="text-red-500 text-sm mt-1">
            {{ errors[field.name] }}
          </p>
        </div>

        <Button :text="isSubmitting ? 'Masuk...' : 'Masuk'" :disabled="isSubmitting" variant="modern" size="medium"
          class="w-full" @click="handleSubmit" />
      </form>

      <p class="text-center text-gray-600 mt-6">
        Belum Punya Akun?
        <router-link to="/register" class="text-indigo-600 font-semibold hover:underline cursor-pointer">
          Daftar di sini
        </router-link>
      </p>
    </div>

    <Toast ref="toastRef" />
  </div>
</template>

<script setup>
import { onMounted, ref } from 'vue'
import { Mail, Lock, GraduationCap } from 'lucide-vue-next'
import { useRouter } from 'vue-router'

const router = useRouter()
import Input from '../../components/ui/Input.vue'
import Button from '../../components/ui/Button.vue'
import { login } from '../../provider/user.provider'
import { useLocalStorage } from '../../hooks/useLocalStorage'
import Toast from '../../components/utils/Toast.vue'
import { useGetCurrentUser } from '../../hooks/useGetCurrentUser'

const { setValue: setToken, value: token } = useLocalStorage("token");
const { setValue: setUser, value } = useLocalStorage("user");
const { user } = useGetCurrentUser();

const toastRef = ref(null);
const formData = ref({
  email: "",
  password: "",
});
const errors = ref({});
const isSubmitting = ref(false);

const fields = [
  {
    id: 1,
    name: "email",
    title: "Email",
    type: "email",
    placeholder: "email@example.com",
    icon: Mail,
  },
  {
    id: 2,
    name: "password",
    title: "Password",
    type: "password",
    placeholder: "Minimal 6 karakter",
    icon: Lock,
  },
];

const handleSubmit = async () => {
  try {
    isSubmitting.value = true;
    const data = await login(formData.value);
    console.log(data.data.data);
    if (data.data.token) {
      setToken(data.data.token)
      setUser(data.data.data)
      
      // Redirect based on user role
      const userRole = data.data.data.role
      const redirectPath = userRole === 'dosen' ? '/dosen/dashboard' : '/'
      
      toastRef.value.showToast('success', 'Login Berhasil', 'Selamat datang kembali!')
      
      setTimeout(() => {
        router.push(redirectPath)
      }, 1500)
    }

    isSubmitting.value = false
  } catch (error) {
    console.log('Login error:', error.response?.data)
    
    let errorMessage = 'Email atau password salah.'
    if (error.response?.data?.message) {
      errorMessage = error.response.data.message
    }

    toastRef.value.showToast('error', 'Login Gagal', errorMessage)

    isSubmitting.value = false;
  }
};
console.log(user.value);
</script>
