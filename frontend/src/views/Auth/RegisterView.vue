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
        <h1 class="text-3xl font-bold text-gray-800 mb-2">Daftar Akun</h1>
        <p class="text-gray-600">Lengkapi data diri Anda untuk mendaftar</p>
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
          :text="isSubmitting ? 'Mendaftar...' : 'Daftar'"
          :disabled="isSubmitting"
          variant="modern"
          size="medium"
          class="w-full"
          @click="handleSubmit"
        />
      </form>

      <p class="text-center text-gray-600 mt-6">
        Sudah punya akun?
        <router-link
          to="/login"
          class="text-indigo-600 font-semibold hover:underline cursor-pointer"
        >
          Login di sini
        </router-link>
      </p>
    </div>

    <Toast ref="toastRef" />
  </div>
</template>

<script setup>
import { ref } from "vue";
import {
  Mail,
  Lock,
  User,
  BookOpen,
  Building2,
  GraduationCap,
} from "lucide-vue-next";
import { useRouter } from "vue-router";

const router = useRouter();
import Input from "../../components/ui/Input.vue";
import Button from "../../components/ui/Button.vue";
import { register } from "../../provider/user.provider";
import Toast from "../../components/utils/Toast.vue";

const toastRef = ref(null);
const formData = ref({
  email: "",
  password: "",
  name: "",
  nim: "",
  major: "",
  faculty: "",
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
  {
    id: 3,
    name: "name",
    title: "Nama Lengkap",
    type: "text",
    placeholder: "Nama lengkap Anda",
    icon: User,
  },
  {
    id: 4,
    name: "nim",
    title: "NIM",
    type: "text",
    placeholder: "Nomor Induk Mahasiswa",
    icon: BookOpen,
  },
  {
    id: 5,
    name: "major",
    title: "Jurusan",
    type: "text",
    placeholder: "Contoh: Teknik Informatika",
    icon: GraduationCap,
  },
  {
    id: 6,
    name: "faculty",
    title: "Fakultas",
    type: "text",
    placeholder: "Contoh: Fakultas Teknik",
    icon: Building2,
  },
];

const handleSubmit = async () => {
  try {
    isSubmitting.value = true;
    const data = await register(formData.value);
    console.log("Response:", data.data);

    toastRef.value.showToast(
      "success",
      "Registrasi Berhasil",
      "Akun Anda berhasil dibuat!"
    );

    // Redirect to login page after successful registration
    setTimeout(() => {
      router.push("/login");
    }, 1500);

    isSubmitting.value = false;
  } catch (error) {
    console.log("Something error", error.response);
    errors.value = error.response?.error || {};

    toastRef.value.showToast(
      "error",
      "Registrasi Gagal",
      "Periksa kembali data yang Anda isi."
    );

    isSubmitting.value = false;
  }
};
</script>
