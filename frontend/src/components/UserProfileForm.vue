<template>
  <div class="max-w-2xl mx-auto bg-white rounded-lg shadow-md p-6">
    <h2 class="text-2xl font-bold mb-6 text-gray-800">Profil Pengguna</h2>
    
    <form @submit.prevent="handleSubmit" class="space-y-4">
      <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1">Nama Lengkap</label>
          <input
            v-model="localData.fullName"
            type="text"
            class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
            placeholder="Nama lengkap Anda"
          />
        </div>
        
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1">Email</label>
          <input
            v-model="localData.email"
            type="email"
            class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
            placeholder="Email Anda"
          />
        </div>
        
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1">NIM/NIP/ID</label>
          <input
            v-model="localData.idNumber"
            type="text"
            class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
            placeholder="Nomor Induk"
          />
        </div>
        
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1">Jenis Kelamin</label>
          <select
            v-model="localData.gender"
            class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
          >
            <option value="">Pilih jenis kelamin</option>
            <option value="Male">Laki-laki</option>
            <option value="Female">Perempuan</option>
          </select>
        </div>
        
        <div class="md:col-span-2">
          <label class="block text-sm font-medium text-gray-700 mb-1">Riwayat</label>
          <textarea
            v-model="localData.history"
            rows="3"
            class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
            placeholder="Riwayat pendidikan atau informasi lainnya"
          ></textarea>
        </div>
      </div>
      
      <div v-if="localData.imageUrl" class="mt-4">
        <label class="block text-sm font-medium text-gray-700 mb-1">Foto Profil</label>
        <img :src="localData.imageUrl" alt="Foto Profil" class="w-24 h-24 rounded-full object-cover border-2 border-gray-300" />
      </div>
      
      <div class="flex justify-end space-x-3 mt-6">
        <button
          type="button"
          @click="handleCancel"
          class="px-4 py-2 text-sm font-medium text-gray-700 bg-gray-200 rounded-md hover:bg-gray-300 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-gray-500"
        >
          Batal
        </button>
        <button
          type="submit"
          class="px-4 py-2 text-sm font-medium text-white bg-blue-600 rounded-md hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
        >
          Simpan
        </button>
      </div>
    </form>
  </div>
</template>

<script>
export default {
  name: "UserProfileForm",
  props: {
    profileData: {
      type: Object,
      required: true,
    },
    userRole: {
      type: String,
      required: true,
    },
  },
  emits: ["save"],
  data() {
    return {
      localData: { ...this.profileData },
    };
  },
  methods: {
    handleSubmit() {
      this.$emit("save", { ...this.localData });
    },
    handleCancel() {
      // Kembalikan data ke kondisi awal
      this.localData = { ...this.profileData };
      this.$emit("cancel");
    },
  },
  watch: {
    profileData: {
      handler(newVal) {
        this.localData = { ...newVal };
      },
      deep: true,
    },
  },
};
</script>