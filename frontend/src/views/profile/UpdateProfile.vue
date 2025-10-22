<template>
  <div class="min-h-screen bg-gradient-to-br from-blue-50 via-purple-50 to-pink-50 flex items-center justify-center p-4">
    <div class="w-full max-w-2xl bg-white rounded-3xl shadow-xl p-8">
      <!-- Header with Back Button -->
      <div class="flex items-center gap-4 mb-8">
        <button
          @click="handleBack"
          class="p-2 hover:bg-gray-100 rounded-lg transition-colors"
          aria-label="Go back"
        >
          <ArrowLeft class="w-5 h-5 text-gray-600" />
        </button>
        
        <div class="relative">
          <img
            v-if="user?.profilePic"
            :src="user.profilePic"
            alt="Profile"
            class="w-16 h-16 rounded-full object-cover"
          />
          <div v-else class="w-16 h-16 rounded-full bg-gradient-to-br from-blue-400 to-purple-500 flex items-center justify-center">
            <User class="w-8 h-8 text-white" />
          </div>
          <div class="absolute -bottom-1 -right-1 w-5 h-5 bg-green-500 border-2 border-white rounded-full"></div>
        </div>
        
        <div class="flex-1">
          <h2 class="text-xl font-semibold text-gray-800">{{ formData.name || 'User' }}</h2>
          <p class="text-sm text-gray-500">{{ formData.email || 'email@example.com' }}</p>
        </div>
      </div>

      <!-- Form Fields -->
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6 mb-6">
        <div v-for="field in fields" :key="field.name">
          <label class="block text-sm font-medium text-gray-700 mb-2">
            {{ field.title }}
          </label>
          <div class="relative">
            <div class="absolute left-3 top-1/2 -translate-y-1/2 text-gray-400">
              <component :is="field.icon" class="w-5 h-5" />
            </div>
            
            <select
              v-if="field.type === 'select'"
              v-model="formData[field.name]"
              :class="[
                'w-full pl-11 pr-4 py-3 bg-gray-50 border rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent transition appearance-none',
                errors[field.name] ? 'border-red-300' : 'border-gray-200'
              ]"
            >
              <option value="">{{ field.placeholder }}</option>
              <option v-for="option in field.options" :key="option" :value="option">
                {{ option }}
              </option>
            </select>
            
            <input
              v-else
              v-model="formData[field.name]"
              :type="field.type"
              :placeholder="field.placeholder"
              :class="[
                'w-full pl-11 pr-4 py-3 bg-gray-50 border rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent transition',
                errors[field.name] ? 'border-red-300' : 'border-gray-200'
              ]"
            />
          </div>
        </div>
      </div>
<!--  //soon
      <div class="mb-6">
        <label class="block text-sm font-medium text-gray-700 mb-2">
          Biodata
        </label>
        <textarea
          v-model="formData.biodata"
          rows="4"
          :class="[
            'w-full px-4 py-3 bg-gray-50 border rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent transition resize-none',
            errors.biodata ? 'border-red-300' : 'border-gray-200'
          ]"
          placeholder="Write your biodata..."
        ></textarea>
      </div> -->

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
import { ArrowLeft, User, CreditCard, Users, Mail } from 'lucide-vue-next';
import { useGetCurrentUser } from '../../hooks/useGetCurrentUser';
import Button from '../../components/ui/Button.vue';
import { updateUser } from '../../provider/user.provider';
import { useLocalStorage } from '../../hooks/useLocalStorage';
import Toast from '../../components/utils/Toast.vue';

const router = useRouter();
const { user } = useGetCurrentUser();
const { value, setValue, removeValue } = useLocalStorage("user")

const toastRef = ref(null);
const loading = ref(false);

const formData = ref({
  name: '',
  nim: '',
  email: '',
});

const errors = ref({});

const fields = ref([
  {
    name: 'name',
    title: 'Full Name',
    placeholder: 'Enter your full name',
    icon: User,
    type: 'text'
  },
  {
    name: 'nim',
    title: 'NIM',
    placeholder: 'Enter your NIM',
    icon: CreditCard,
    type: 'text'
  },
  {
    name: 'email',
    title: 'Email',
    placeholder: 'Enter your email',
    icon: Mail,
    type: 'email'
  }
]);

onMounted(() => {
  console.log("Mounted user:", user.value)

  if (user) {
    formData.value = {
      name: user.value.name || '',
      nim: user.value.nim || '',
      email: user.value.email || '',
    };
  }
});

const handleSubmit = async () => {
  errors.value = {};

  let isValid = true;
  
  if (!formData.value.name.trim()) {
    errors.value.name = 'Full name is required';
    isValid = false;
  }
  
  if (!formData.value.nim.trim()) {
    errors.value.nim = 'NIM is required';
    isValid = false;
  }
  
  if (!formData.value.email.trim()) {
    errors.value.email = 'Email is required';
    isValid = false;
  } else if (!/^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(formData.value.email)) {
    errors.value.email = 'Invalid email format';
    isValid = false;
  }

  if (!isValid) {
    toastRef.value.showToast(
      "error",
      "Validation Error",
      "Please fill in all required fields correctly."
    );
    return;
  }

  try {
    loading.value = true;
    console.log(formData.value.email)
    
    const data = await updateUser(formData.value, user.value.id)
    console.log('Form submitted:', data.data);

    toastRef.value.showToast(
      "success",
      "Profile Updated",
      "Your profile has been updated successfully!"
    );

    removeValue()
    setValue(data.data)
    router.back();
  } catch (error) {
    console.error('Failed to update profile:', error);
    
    toastRef.value.showToast(
      "error",
      "Update Failed",
      error.response?.data?.message || "Failed to update profile. Please try again."
    );
  } finally {
    loading.value = false;
  }
};

const handleBack = () => {
  router.back();
};
</script>

<style scoped>
</style>