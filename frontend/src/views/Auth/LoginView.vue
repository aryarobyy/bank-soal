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
// <-- PERUBAHAN: Impor 'User' menggantikan 'Mail'
import { User, Lock, GraduationCap } from "lucide-vue-next";
import Input from "../../components/ui/Input.vue";
import Button from "../../components/ui/Button.vue";
import { login } from "../../provider/user.provider";
import { useLocalStorage } from "../../hooks/useLocalStorage";
import Toast from "../../components/utils/Toast.vue";
import { useRouter } from "vue-router";
import { useUser } from "../../hooks/useGetCurrentUser";

const { setValue: setToken } = useLocalStorage("token");
const { setValue: setId } = useLocalStorage("id");
const { setUser: setGlobalUser } = useUser();

const toastRef = ref(null);
// <-- PERUBAHAN: 'email' diubah menjadi 'login_id'
const formData = ref({
  login_id: "",
  password: "",
});
const errors = ref({});
const isSubmitting = ref(false);
const router = useRouter();

// <-- PERUBAHAN: Field email diganti dengan login_id
const fields = [
  {
    id: 1,
    name: "login_id",

    title: "Login ID",
    type: "text",
    placeholder: "NIM / NIDN / Username",
    icon: User, // <-- Ikon diubah
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
    // Panggilan 'login' tidak perlu diubah,
    // karena 'formData.value' sekarang sudah berisi { login_id: "...", password: "..." }
    const data = await login(formData.value);

    const userData = data.data.data;

    if (data.data.token && userData) {
      setToken(data.data.token);
      setId(userData.id);
      setGlobalUser(userData);
    }

    const userRole = userData.role;
    let redirectPath = "/";

    if (userRole === "lecturer") {
      redirectPath = "/dosen/dashboard";
    } else if (userRole === "admin") {
      redirectPath = "/admin/dashboard";
    } else if (userRole === "super_admin") {
      redirectPath = "/superadmin/dashboard";
    }

    toastRef.value.showToast(
      "success",
      "Login Berhasil",
      "Selamat datang kembali!"
    );

    isSubmitting.value = false;

    // 3. Arahkan ke halaman yang sudah ditentukan

    router.push(redirectPath);
  } catch (error) {
    console.log("Something error", error.response?.data);
    toastRef.value.showToast(
      "error",
      "Login Gagal",
      "Login ID atau password salah." // <-- Pesan error diperbarui
    );
    isSubmitting.value = false;
  }
};
</script>
