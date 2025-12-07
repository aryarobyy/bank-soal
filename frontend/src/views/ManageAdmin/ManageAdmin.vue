<template>
  <div class="p-6 bg-gray-50 min-h-screen">
    <div class="flex justify-between items-center mb-6">
      <h2 class="text-2xl font-bold text-gray-800">Manajemen Akun Admin</h2>
      <button
        @click="openAddModal"
        class="px-4 py-2 bg-indigo-500 text-white rounded-lg hover:bg-indigo-600 transition"
      >
        + Tambah Admin
      </button>
    </div>

    <div v-if="loading" class="text-center py-10">
      <p class="text-gray-500">Memuat data admin...</p>
    </div>
    <div v-else-if="error" class="text-center py-10 bg-red-50 p-4 rounded-lg">
      <p class="text-red-600">{{ error }}</p>
    </div>

    <div v-else class="bg-white shadow rounded-lg overflow-hidden">
      <table class="min-w-full border-collapse">
        <thead class="bg-gray-100 text-gray-700 text-sm">
          <tr>
            <th class="px-4 py-3 text-left">No</th>
            <th class="px-4 py-3 text-left">Nama</th>
            <th class="px-4 py-3 text-left">Email</th>
            <th class="px-4 py-3 text-left">Username</th>
            <th class="px-4 py-3 text-left">Role</th>
            <th class="px-4 py-3 text-left">Tanggal Dibuat</th>
            <th class="px-4 py-3 text-left">Aksi</th>
          </tr>
        </thead>
        <tbody class="text-gray-800 text-sm">
          <tr
            v-for="(admin, index) in adminList"
            :key="admin.id || admin.ID || admin._id"
            class="border-t hover:bg-gray-50 transition"
          >
            <td class="px-4 py-3">{{ (currentPage - 1) * itemsPerPage + index + 1 }}</td>
            <td class="px-4 py-3 font-medium">{{ admin.name }}</td>
            <td class="px-4 py-3">{{ admin.email }}</td>
            <td class="px-4 py-3">{{ admin.username }}</td>
            <td class="px-4 py-3">
              <span 
                :class="roleClass(admin.role)" 
                class="px-2 py-1 text-xs font-semibold rounded-full capitalize"
              >
                {{ admin.role }}
              </span>
            </td>
            <td class="px-4 py-3">{{ new Date(admin.created_at).toLocaleDateString("id-ID") }}</td>
            <td class="px-4 py-3">
              <button
                @click="editAdmin(admin)"
                class="px-3 py-1 bg-yellow-400 text-white rounded-md hover:bg-yellow-500 mr-2 transition"
              >
                Edit
              </button>
              <button
                @click="hapusAdmin(admin)"
                class="px-3 py-1 bg-red-500 text-white rounded-md hover:bg-red-600 transition"
              >
                Hapus
              </button>
            </td>
          </tr>
          <tr v-if="adminList.length === 0">
            <td colspan="7" class="px-4 py-4 text-center text-gray-500">
              Belum ada data admin
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <div v-if="!loading && totalPages > 1" class="flex justify-between items-center mt-6">
      <span class="text-sm text-gray-700">
        Halaman <span class="font-semibold">{{ currentPage }}</span> dari <span class="font-semibold">{{ totalPages }}</span> (Total <span class="font-semibold">{{ totalItems }}</span> admin)
      </span>
      <div class="flex gap-1">
        <button
          @click="prevPage"
          :disabled="currentPage === 1"
          class="px-3 py-1 bg-white border border-gray-300 rounded-md text-sm hover:bg-gray-50 disabled:opacity-50 disabled:cursor-not-allowed"
        >
          &lt; Sebelumnya
        </button>
        <button
          @click="nextPage"
          :disabled="currentPage === totalPages"
          class="px-3 py-1 bg-white border border-gray-300 rounded-md text-sm hover:bg-gray-50 disabled:opacity-50 disabled:cursor-not-allowed"
        >
          Berikutnya &gt;
        </button>
      </div>
    </div>

    <div
      v-if="showModal"
      class="fixed inset-0 flex items-center justify-center bg-black/30 backdrop-blur-sm z-50"
    >
      <div class="bg-white rounded-lg shadow-lg p-6 w-full max-w-md">
        <h3 class="text-lg font-semibold mb-4">
          {{ editMode ? "Edit Admin" : "Tambah Admin" }}
        </h3>

        <form @submit.prevent="simpanAdmin">
          <div class="mb-3">
            <label class="block mb-1 text-sm font-medium text-gray-700">Nama</label>
            <input v-model="form.name" type="text" required class="w-full p-2 border rounded-md"/>
          </div>
          <div class="mb-3">
            <label class="block mb-1 text-sm font-medium text-gray-700">Email</label>
            <input v-model="form.email" type="email" required class="w-full p-2 border rounded-md"/>
          </div>

          <div class="mb-3">
            <label class="block mb-1 text-sm font-medium text-gray-700">Username</label>
            <input v-model="form.username" type="text" required class="w-full p-2 border rounded-md"/>
          </div>

          <template v-if="!editMode">
            <div class="mb-3">
              <label class="block mb-1 text-sm font-medium text-gray-700">Password</label>
              <input v-model="form.password" type="password" required class="w-full p-2 border rounded-md"/>
            </div>

            <div class="mb-3">
              <label class="block mb-1 text-sm font-medium text-gray-700">Role</label>
              <select v-model="form.role" class="w-full p-2 border rounded-md bg-white">
                <option v-for="role in availableRoles" :key="role" :value="role">
                  {{ role }}
                </option>
              </select>
            </div>
            
            <div class="mb-3">
              <label class="block mb-1 text-sm font-medium text-gray-700">Jurusan (Major)</label>
              <input v-model="form.major" type="text" class="w-full p-2 border rounded-md" />
            </div>
            <div class="mb-3">
              <label class="block mb-1 text-sm font-medium text-gray-700">Fakultas (Faculty)</label>
              <input v-model="form.faculty" type="text" class="w-full p-2 border rounded-md" />
            </div>
          </template>

          <template v-else>
            <div class="mb-3">
              <label class="block mb-1 text-sm font-medium text-gray-700">Password Baru</label>
              <input v-model="form.password" type="password" class="w-full p-2 border rounded-md" placeholder="Kosongkan jika tidak ingin diubah"/>
            </div>
            
            <div class="mb-3">
              <label class="block mb-1 text-sm font-medium text-gray-700">Ubah Role</label>
              <select v-model="form.role" class="w-full p-2 border rounded-md bg-white">
                <option v-for="role in availableRoles" :key="role" :value="role">
                  {{ role }}
                </option>
              </select>
            </div>
          </template>
          
          <div class="flex justify-end gap-2 mt-4">
            <button
              type="button" @click="closeModal"
              class="px-4 py-2 bg-gray-300 text-gray-700 rounded-md hover:bg-gray-400 transition"
            >
              Batal
            </button>
            <button
              type="submit"
              class="px-4 py-2 bg-indigo-500 text-white rounded-md hover:bg-indigo-600 transition"
            >
              {{ editMode ? "Simpan" : "Tambah" }}
            </button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, computed, watch } from "vue";
import {
  getUsersByRole,
  register,
  updateUser,
  deleteUser,
  changePassword,
  changeRole,
} from "../../provider/user.provider.js";
import { useGetCurrentUser } from "../../hooks/useGetCurrentUser";

import { usePopup } from "../../hooks/usePopup";


const { showSuccess, showError, showConfirm } = usePopup();

const adminList = ref([]);
const loading = ref(true);
const error = ref(null);
const showModal = ref(false);
const editMode = ref(false);

const availableRoles = ref(['admin', 'lecturer', 'user']);
const originalAdminData = ref(null); 

const initialFormState = {
  id: null,
  name: "",
  email: "",
  username: "",
  password: "",
  role: "admin", 
  nim: null, 
  major: "",   
  faculty: ""  
};
const form = ref({ ...initialFormState });

const { user: storedUser } = useGetCurrentUser();

const currentPage = ref(1);
const itemsPerPage = ref(10);
const totalItems = ref(0);

const totalPages = computed(() => {
  return Math.ceil(totalItems.value / itemsPerPage.value);
});

const fetchAdmins = async () => {
  try {
    loading.value = true;
    const offset = (currentPage.value - 1) * itemsPerPage.value;
    const response = await getUsersByRole("admin", itemsPerPage.value, offset);
    adminList.value = response.data || [];
    totalItems.value = response.total || 0;
  } catch (err) {
    console.error("Gagal mengambil data admin:", err);
    error.value = "Tidak dapat memuat data. Silakan coba lagi nanti.";
  } finally {
    loading.value = false;
  }
};

onMounted(fetchAdmins);

watch(currentPage, (newPage, oldPage) => {
  if (newPage !== oldPage) {
    fetchAdmins();
  }
});

const prevPage = () => {
  if (currentPage.value > 1) {
    currentPage.value--;
  }
};
const nextPage = () => {
  if (currentPage.value < totalPages.value) {
    currentPage.value++;
  }
};

const openAddModal = () => {
  editMode.value = false;
  form.value = { ...initialFormState };
  originalAdminData.value = null; 
  showModal.value = true;
};

const closeModal = () => {
  showModal.value = false;
  originalAdminData.value = null; 
};

const simpanAdmin = async () => {
  try {
    const userId = form.value.id;
    const adminId = storedUser.value?.id || storedUser.value?.ID;

    if (!adminId) {
      showError("Akses Ditolak", "Error: Sesi Super Admin tidak ditemukan. Silakan login ulang.");
      return;
    }

    if (editMode.value) {
      if (!userId) {
        showError("Data Invalid", "Error: ID admin tidak ditemukan."); return;
      }
      
      const dataToUpdate = { 
        name: form.value.name, 
        email: form.value.email,
        username: form.value.username 
      };
      await updateUser(dataToUpdate, userId);

      let passwordErrorMessage = "";
      if (form.value.password && form.value.password.trim() !== "") {
        try {
          await changePassword(userId, form.value.password, adminId);
        } catch (passwordError) {
          console.warn("Gagal mengganti password:", passwordError);
          passwordErrorMessage = passwordError.response?.data?.message || "";
        }
      }
      let roleErrorMessage = "";
      const originalRole = originalAdminData.value?.role; 
      const newRole = form.value.role;
      if (newRole && originalRole && newRole !== originalRole) {
        try {
          await changeRole(userId, adminId, newRole); 
        } catch (roleError) {
          console.warn("Gagal mengganti role:", roleError);
          roleErrorMessage = roleError.response?.data?.message || "";
        }
      }
      
      let errors = [];
      if (passwordErrorMessage) errors.push(passwordErrorMessage);
      if (roleErrorMessage) errors.push(roleErrorMessage);
      
      if (errors.length > 0) {
        await showSuccess("Update Sebagian", `Data diperbarui, namun: ${errors.join(', ')}`);
      } else {
        await showSuccess("Berhasil", "Data admin berhasil diperbarui!");
      }

    } else {
      const dataToCreate = {
        name: form.value.name,
        email: form.value.email,
        username: form.value.username,
        password: form.value.password,
        role: form.value.role,
        nim: null, 
        major: form.value.major.trim() || null, 
        faculty: form.value.faculty.trim() || null
      };
      await register(dataToCreate);
      await showSuccess("Berhasil", "Admin baru berhasil ditambahkan!");
    }
    closeModal();
    
  
    fetchAdmins(); 

  } catch (err) {
    console.error("Gagal menyimpan data:", err);
    showError("Gagal", err.response?.data?.message || "Terjadi kesalahan.");
  }
};

const editAdmin = (admin) => {
  editMode.value = true;
  originalAdminData.value = { ...admin }; 
  const userId = admin.id || admin.ID || admin._id;
  form.value = { 
    ...initialFormState, 
    ...admin,
    username: admin.username || "",
    major: admin.major || "",
    faculty: admin.faculty || "",
    id: userId,
    password: "" 
  };
  showModal.value = true;
};

const hapusAdmin = async (admin) => {
  const userId = admin.id || admin.ID || admin._id;
  if (!userId) {
    showError("Data Invalid", "Error: ID admin tidak ditemukan."); return;
  }
  
  const isConfirmed = await showConfirm(
    "Konfirmasi Hapus", 
    "Yakin ingin menghapus admin ini?",
    "Ya, Hapus"
  );

  if (isConfirmed) {
    try {
      await deleteUser(userId);
      
      
      const oldLength = adminList.value.length;
      adminList.value = adminList.value.filter(a => (a.id || a.ID || a._id) !== userId);
      
      if (adminList.value.length < oldLength) {
         totalItems.value--;
      }

      await showSuccess("Berhasil", "Admin berhasil dihapus.");

    
      if (adminList.value.length === 0 && currentPage.value > 1) {
        currentPage.value--;
      } else {
        if(adminList.value.length === 0) fetchAdmins();
      }
    } catch (err) {
      console.error("Gagal menghapus admin:", err);
      showError("Gagal", err.response?.data?.message || err.response?.data || "Gagal menghapus data.");
    }
  }
};

const roleClass = (role) => {
  if (role === 'admin') return 'bg-red-100 text-red-800';
  if (role === 'lecturer') return 'bg-green-100 text-green-800';
  if (role === 'user') return 'bg-blue-100 text-blue-800';
  return 'bg-gray-100 text-gray-800';
};
</script>