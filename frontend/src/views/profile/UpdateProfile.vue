<template>
  <div class="min-h-screen bg-gradient-to-br from-blue-50 via-purple-50 to-pink-50 flex items-center justify-center p-4">
    <div class="w-full max-w-2xl bg-white rounded-3xl shadow-xl p-8">
      <div class="flex items-center gap-4 mb-8">
        <button
          @click="handleBack"
          class="p-2 hover:bg-gray-100 rounded-lg transition-colors"
          aria-label="Go back"
        >
          <ArrowLeft class="w-5 h-5 text-gray-600" />
        </button>
        
        <div class="relative group">
          <img
            :src="avatarPreview || user?.img_url || 'https://ui-avatars.com/api/?name=' + (formData.name || 'U') + '&background=random'"
            alt="Profile"
            class="w-16 h-16 rounded-full object-cover border-2 border-gray-200"
            @error="handleImageError"
          />
          <label for="avatarInput" class="absolute inset-0 flex items-center justify-center bg-black bg-opacity-50 rounded-full cursor-pointer opacity-0 group-hover:opacity-100 transition-opacity">
            <Camera class="w-6 h-6 text-white" />
          </label>
          <input type="file" id="avatarInput" ref="avatarInputRef" @change="handleFileSelect" accept="image/*" class="hidden" />
        </div>
        
        <div class="flex-1">
          <h2 class="text-xl font-semibold text-gray-800">{{ formData.name || 'User' }}</h2>
          <p class="text-sm text-gray-500">{{ formData.email || 'email@example.com' }}</p>
        </div>
      </div>

      <div class="grid grid-cols-1 md:grid-cols-2 gap-6 mb-6">
        <div v-for="field in fields" :key="field.name">
          <label class="block text-sm font-medium text-gray-700 mb-2">
            {{ field.title }}
          </label>
          <div class="relative">
            <div class="absolute left-3 top-1/2 -translate-y-1/2 text-gray-400">
              <component :is="field.icon" class="w-5 h-5" />
            </div>
            
            <input
              v-model="formData[field.name]"
              :type="field.type"
              :placeholder="field.placeholder"
              :class="[
                'w-full pl-11 pr-4 py-3 bg-gray-50 border rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent transition',
                errors[field.name] ? 'border-red-300' : 'border-gray-200'
              ]"
            />
          </div>
          <p v-if="errors[field.name]" class="text-red-500 text-xs mt-1">{{ errors[field.name] }}</p>
        </div>
      </div>

      <div class="grid grid-cols-2 gap-3">
        <Button
          @click="handleBack"
          text="Cancel"
          variant="ghost"
          size="medium"
          class="w-full"
        />
        <Button
          @click="handleSubmit"
          text="Save Changes"
          variant="modern"
          size="medium"
          :loading="loading"
          :disabled="loading"
          class="w-full"
        />
      </div>
    </div>

    <Toast ref="toastRef" />
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { useRouter } from 'vue-router';
import { ArrowLeft, User, CreditCard, Mail, Camera, Clipboard, BookText } from 'lucide-vue-next';
import { useUser } from '../../hooks/useGetCurrentUser'; 
import Button from '../../components/ui/Button.vue';
import { updateUser, getUserById } from '../../provider/user.provider';
import { useLocalStorage } from '../../hooks/useLocalStorage';
import Toast from '../../components/utils/Toast.vue';

const router = useRouter();
const { user, setUser } = useUser(); 
const { value, setValue, removeValue } = useLocalStorage("user")

const toastRef = ref(null);
const loading = ref(false);

const avatarInputRef = ref(null);
const avatarFile = ref(null);
const avatarPreview = ref(null);

const formData = ref({
  name: '',
  nim: '',
  email: '',
  nip: '',
  nidn: '',
  major: '',
  faculty: '',
});

const errors = ref({});
const fields = ref([]);

onMounted(() => {
  if (user.value) {
    formData.value = {
      name: user.value.name || '',
      email: user.value.email || '',
      nim: user.value.nim || '',
      nip: user.value.nip || '',
      nidn: user.value.nidn || '',
      major: user.value.major || '',
      faculty: user.value.faculty || '',
    };
    
    if (user.value.img_url) {
      avatarPreview.value = user.value.img_url;
    }

    const baseFields = [
      {
        name: 'name', title: 'Full Name',
        placeholder: 'Enter your full name', icon: User, type: 'text'
      },
      {
        name: 'email', title: 'Email',
        placeholder: 'Enter your email', icon: Mail, type: 'email'
      }
    ];
    
    fields.value = [...baseFields];

    if (user.value.role === 'user') {
      fields.value.push({
        name: 'nim', title: 'NIM',
        placeholder: 'Enter your NIM', icon: CreditCard, type: 'text'
      });
    } else if (user.value.role === 'lecturer') {
      fields.value.push({
        name: 'nip', title: 'NIP',
        placeholder: 'Enter your NIP', icon: Clipboard, type: 'text'
      });
      fields.value.push({
        name: 'nidn', title: 'NIDN',
        placeholder: 'Enter your NIDN', icon: Clipboard, type: 'text'
      });
    } else if (user.value.role === 'admin') {
      fields.value.push({
        name: 'major', title: 'Jurusan (Major)',
        placeholder: 'Enter your major', icon: BookText, type: 'text'
      });
      fields.value.push({
        name: 'faculty', title: 'Fakultas (Faculty)',
        placeholder: 'Enter your faculty', icon: BookText, type: 'text'
      });
    }
  }
});

const handleImageError = (e) => {
  console.error('Image failed to load:', e.target.src);
  e.target.src = 'https://ui-avatars.com/api/?name=' + encodeURIComponent(formData.value.name || 'U') + '&background=random';
};

const handleFileSelect = (event) => {
  const file = event.target.files[0];
  if (!file) return;

  if (!file.type.startsWith('image/')) {
    toastRef.value.showToast("error", "File Invalid", "Hanya file gambar yang diperbolehkan.");
    return;
  }
  
  if (file.size > 2 * 1024 * 1024) {
    toastRef.value.showToast("error", "Ukuran File", "Ukuran file maksimal 2MB.");
    return;
  }

  avatarFile.value = file;
  
  const reader = new FileReader();
  reader.onload = (e) => {
    avatarPreview.value = e.target.result;
  };
  reader.readAsDataURL(file);
};

const handleSubmit = async () => {
  errors.value = {};
  let isValid = true;
  

  if (!formData.value.name.trim()) {
    errors.value.name = 'Nama lengkap wajib diisi'; 
    isValid = false;
  } else if (!/^[a-zA-Z\s]+$/.test(formData.value.name)) {
    // Regex: Hanya a-z, A-Z, dan spasi
    errors.value.name = 'Nama hanya boleh berisi huruf'; 
    isValid = false;
  }

 
  if (user.value && user.value.role === 'user' && !formData.value.nim.trim()) {
     errors.value.nim = 'NIM wajib diisi'; 
     isValid = false;
  }


  if (!formData.value.email.trim()) {
    errors.value.email = 'Email wajib diisi'; 
    isValid = false;
  } else if (!/^[a-zA-Z0-9._-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/.test(formData.value.email)) {
   
    errors.value.email = 'Format email tidak valid (dilarang menggunakan simbol khusus selain @, titik, underscore)'; 
    isValid = false;
  }
  
  if (!isValid) {
    toastRef.value.showToast("error", "Validasi Gagal", "Mohon periksa kembali inputan Anda.");
    return;
  }

  try {
    loading.value = true;
    
    const dataToUpdate = {
      name: formData.value.name,
      email: formData.value.email,
    };

    if (user.value.role === 'user') {
      dataToUpdate.nim = formData.value.nim;
    } else if (user.value.role === 'lecturer') {
      dataToUpdate.nip = formData.value.nip;
      dataToUpdate.nidn = formData.value.nidn;
    } else if (user.value.role === 'admin') {
      dataToUpdate.major = formData.value.major;
      dataToUpdate.faculty = formData.value.faculty;
    }

    if (avatarFile.value) {
      dataToUpdate.image = avatarFile.value; 
    }
    
    const response = await updateUser(dataToUpdate, user.value.id);
    
    const freshUserResponse = await getUserById(user.value.id);
    const freshUserData = freshUserResponse.data;
    
    setUser(freshUserData); 
    removeValue();
    setValue(freshUserData);
    
    toastRef.value.showToast("success", "Berhasil", "Profil berhasil diperbarui!");
    
    setTimeout(() => {
      router.back();
    }, 1500);
    
  } catch (error) {
    console.error('Failed to update profile:', error);
    const errorMessage = error.response?.data?.message || "Gagal memperbarui profil.";
    toastRef.value.showToast("error", "Gagal", errorMessage);
  } finally {
    loading.value = false;
  }
};

const handleBack = () => {
  router.back();
};
</script>

<style scoped>
.group:hover .opacity-0 {
  opacity: 1;
}
</style>