<template>
  <div
    class="min-h-screen bg-gradient-to-br from-blue-50 to-indigo-100 flex items-center justify-center p-4"
  >
    <div class="bg-white rounded-2xl shadow-xl w-full max-w-md p-8">
      <div class="text-center mb-8">
        
        <img 
          :src="logoImage" 
          alt="Logo Latih.in" 
          class="w-24 h-auto mb-4 mx-auto object-contain"
        />

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
              :required="false" 
              :id="field.id"
              :icon="field.icon"
              :type="field.type"
            />
          </div>
        </div>

        <Button
          :text="isSubmitting ? 'Masuk...' : 'Masuk'"
          :disabled="isSubmitting"
          variant="modern"
          size="medium"
          class="w-full"
          @click="handleSubmit"
        />
      </form>

    </div>

    <Toast ref="toastRef" />
  </div>
</template>

<script setup>
import { ref } from "vue";
import { User, Lock } from "lucide-vue-next"; 
import Input from '../../components/ui/Input.vue'
import Button from "../../components/ui/Button.vue";
import { login } from "../../provider/user.provider";
import { useLocalStorage } from "../../hooks/useLocalStorage";
import Toast from "../../components/utils/Toast.vue";
import { useRouter } from 'vue-router'
import { useUser } from "../../hooks/useGetCurrentUser";
import logoImage from '../../assets/login-illustration.webp';

const { setValue: setToken } = useLocalStorage("token");
const { setValue: setId } = useLocalStorage("id");
const { setUser: setGlobalUser } = useUser();

const toastRef = ref(null);
const router = useRouter()

const formData = ref({
  login_id: "",
  password: "",
});

const isSubmitting = ref(false);

const fields = [
  { 
    id: 1, 
    name: "login_id", 
    title: "Login ID", 
    type: "text", 
    placeholder: "NIM / NIP / Username", 
    icon: User, 
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


const validateForm = () => {
  const { login_id, password } = formData.value;

  if (!login_id && !password) {
    toastRef.value.showToast("error", "Validasi Gagal", "Harap isi Login ID dan Password");
    return false;
  }
 
  if (!login_id) {
    toastRef.value.showToast("error", "Validasi Gagal", "Login ID belum diisi");
    return false;
  }

  if (!password) {
    toastRef.value.showToast("error", "Validasi Gagal", "Password belum diisi");
    return false;
  }

  return true;
};

const handleSubmit = async () => {

  if (!validateForm()) {
    return; 
  }

  try {
    isSubmitting.value = true;
    
    const data = await login(formData.value);
    const userData = data.data.data;

    if (data.data.token && userData) {
      setToken(data.data.token);
      setId(userData.id);
      setGlobalUser(userData);
    }

    const userRole = userData.role;
    let redirectPath = '/'; 
    
    if (userRole === 'lecturer') {
      redirectPath = '/dosen/dashboard';
    } else if (userRole === 'admin') {
      redirectPath = '/admin/dashboard';
    } else if (userRole === 'super_admin') {
      redirectPath = '/superadmin/dashboard';
    }

    toastRef.value.showToast(
      "success",
      "Login Berhasil",
      "Selamat datang kembali!"
    );

    router.push(redirectPath);

  } catch (error) {
    console.log("Login error", error.response?.data);
    
    toastRef.value.showToast(
      "error",
      "Login Gagal",
      "Login ID atau Password yang Anda masukkan salah" 
    );
  } finally {
    isSubmitting.value = false;
  }
};
</script>