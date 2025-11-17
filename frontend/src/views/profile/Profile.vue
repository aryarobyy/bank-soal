<template>
  <div class="min-h-screen bg-gradient-to-br from-blue-50 via-purple-50 to-pink-50 flex items-center justify-center p-4">
    <div v-if="loading || isLoading" class="w-full max-w-2xl bg-white rounded-3xl shadow-xl p-8">
      <div class="flex flex-col items-center justify-center py-12">
        <div class="relative w-16 h-16 mb-4">
          <div class="absolute top-0 left-0 w-full h-full border-4 border-blue-200 rounded-full"></div>
          <div class="absolute top-0 left-0 w-full h-full border-4 border-blue-500 rounded-full border-t-transparent animate-spin"></div>
        </div>
        <p class="text-gray-600 font-medium">Loading profile...</p>
      </div>
    </div>

    <div v-else-if="error || isError" class="w-full max-w-2xl bg-white rounded-3xl shadow-xl p-8">
      <div class="flex flex-col items-center justify-center py-12">
        <div class="w-16 h-16 bg-red-100 rounded-full flex items-center justify-center mb-4">
          <User class="w-8 h-8 text-red-500"/>
        </div>
        <h3 class="text-xl font-semibold text-gray-800 mb-2">Failed to load profile</h3>
        <p class="text-gray-600 text-center mb-6">{{ error }}</p>
        <button 
          @click="handleRetry"
          class="bg-blue-500 hover:bg-blue-600 text-white font-medium px-6 py-2.5 rounded-lg transition shadow-sm hover:shadow-md"
        >
          Try Again
        </button>
      </div>
    </div>

    <div v-else class="w-full max-w-2xl bg-white rounded-3xl shadow-xl p-8">
      <div class="flex items-center justify-between mb-8">
        <div class="flex items-center gap-4">
          <div class="relative">
            <img
              :key="imageKey"
              :src="getImageUrl()"
              :alt="user?.name || 'Profile'"
              class="w-16 h-16 rounded-full object-cover border-2 border-gray-200"
              @error="handleImageError"
            />
          </div>
          <div>
            <h2 class="text-xl font-semibold text-gray-800">{{ user?.name || 'User' }}</h2>
            <p class="text-sm text-gray-500">{{ user?.email || '-' }}</p>
          </div>
        </div>
        
        <button 
          @click="goToEditProfile" 
          class="flex items-center gap-2 bg-blue-500 hover:bg-blue-600 text-white font-medium px-6 py-2.5 rounded-lg transition shadow-sm hover:shadow-md"
        >
          <svg xmlns="http://www.w3.org/2000/svg" class="w-4 h-4" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
            <path d="M17 3a2.828 2.828 0 1 1 4 4L7.5 20.5 2 22l1.5-5.5L17 3z"></path>
          </svg>
          Edit Profile
        </button>
      </div>

      <div>
        <div class="grid grid-cols-1 md:grid-cols-2 gap-6 mb-6">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">
              Nama
            </label>
            <div class="w-full px-4 py-3 bg-gray-50 border border-gray-200 rounded-lg text-gray-800">
              {{ user?.name || '-' }}
            </div>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">
              Email
            </label>
            <div class="w-full px-4 py-3 bg-gray-50 border border-gray-200 rounded-lg text-gray-800">
              {{ user?.email || '-' }}
            </div>
          </div>
          
          <template v-if="user?.role === 'user'">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">
                NIM
              </label>
              <div class="w-full px-4 py-3 bg-gray-50 border border-gray-200 rounded-lg text-gray-800">
                {{ user?.nim || '-' }}
              </div>
            </div>
          </template>
          
          <template v-if="user?.role === 'lecturer'">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">
                NIP
              </label>
              <div class="w-full px-4 py-3 bg-gray-50 border border-gray-200 rounded-lg text-gray-800">
                {{ user?.nip || '-' }}
              </div>
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">
                NIDN
              </label>
              <div class="w-full px-4 py-3 bg-gray-50 border border-gray-200 rounded-lg text-gray-800">
                {{ user?.nidn || '-' }}
              </div>
            </div>
          </template>
          
          </div>

        <div class="mb-6">
          <label class="block text-sm font-medium text-gray-700 mb-3">
            My email Address
          </label>
          <div class="flex items-start gap-3 bg-gray-50 p-4 rounded-lg border border-gray-200">
            <div>
              <p class="text-sm font-medium text-gray-800">{{ user?.email || '-' }}</p>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, watch, onActivated } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useGetCurrentUser } from '../../hooks/useGetCurrentUser'
import { getUserById } from '../../provider/user.provider'
import { User } from 'lucide-vue-next'

const { user: userLocal, loading, error } = useGetCurrentUser()
const route = useRoute()
const router = useRouter()

const userId = route.params.id
const user = ref(null)
const isOwnProfile = ref(false)
const isLoading = ref(true)
const isError = ref(false)
const imageKey = ref(Date.now()) // Key untuk force re-render gambar

// Function untuk mendapatkan URL gambar dengan cache busting
const getImageUrl = () => {
  if (!user.value) {
    return 'https://ui-avatars.com/api/?name=U&background=random';
  }
  
  if (user.value.img_url) {
    // Tambahkan timestamp untuk menghindari cache browser
    const url = user.value.img_url;
    const separator = url.includes('?') ? '&' : '?';
    return `${url}${separator}t=${imageKey.value}`;
  }
  
  return 'https://ui-avatars.com/api/?name=' + encodeURIComponent(user.value.name || 'U') + '&background=random';
};

// Handle error saat gambar gagal load
const handleImageError = (event) => {
  console.error('Image failed to load:', event.target.src);
  event.target.src = 'https://ui-avatars.com/api/?name=' + encodeURIComponent(user.value?.name || 'U') + '&background=random';
};

// Watch untuk perubahan userLocal
watch(userLocal, (newUser) => {
  if (newUser && newUser.id == userId) {
    console.log('User local updated:', newUser);
    user.value = newUser;
    isOwnProfile.value = true;
    imageKey.value = Date.now(); // Update image key
  }
}, { immediate: true });

// Hook yang dipanggil setiap kali komponen menjadi aktif (termasuk setelah back dari edit)
onActivated(() => {
  console.log('Component activated - refreshing data');
  fetchUser();
});

onMounted(() => {
  console.log('Component mounted');
  fetchUser();
});

const fetchUser = async () => {
  try {
    isLoading.value = true;
    console.log('Fetching user with ID:', userId);
    
    // Selalu fetch dari server untuk mendapatkan data terbaru
    const res = await getUserById(userId);
    const fetchedUser = res.data;
    
    console.log('Fetched user data:', fetchedUser);
    console.log('Image URL:', fetchedUser.img_url);
    
    user.value = fetchedUser;
    isOwnProfile.value = userLocal.value && userLocal.value.id == userId;
    
    // Update image key untuk force reload gambar
    imageKey.value = Date.now();
    
    isError.value = false;
  } catch (err) {
    console.error('Failed to fetch user:', err);
    isError.value = true;
  } finally {
    isLoading.value = false;
  }
}

const goToEditProfile = () => {
  const currentRouteName = route.name; 
  let targetRouteName = 'UpdateProfile'; 
  
  if (currentRouteName === 'DosenProfile') {
    targetRouteName = 'DosenUpdateProfile';
  } else if (currentRouteName === 'AdminProfile') {
    targetRouteName = 'AdminUpdateProfile'; 
  }
  router.push({ name: targetRouteName });
};

const handleRetry = () => {
  isError.value = false;
  fetchUser();
}
</script>

<style scoped>
@keyframes spin { to { transform: rotate(360deg); } }
.animate-spin { animation: spin 1s linear infinite; }
</style>