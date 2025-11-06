<template>
  <div class="p-6 bg-gray-50 min-h-screen">
    <div class="flex justify-between items-center mb-6">
      <h2 class="text-2xl font-bold text-gray-800">Manajemen Akun Mahasiswa</h2>
      <button
        @click="openAddModal"
        class="flex items-center gap-2 px-4 py-2 bg-indigo-500 text-white rounded-lg hover:bg-indigo-600 transition shadow"
      >
        <i class="fas fa-plus-circle"></i> Tambah Mahasiswa
      </button>
    </div>

    <div v-if="loading" class="text-center py-10">
      <p class="text-gray-500">Memuat data mahasiswa...</p>
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
            <th class="px-4 py-3 text-left">NIM</th>
            <th class="px-4 py-3 text-left">Role</th>
            <th class="px-4 py-3 text-left">Tanggal Dibuat</th>
            <th class="px-4 py-3 text-left">Aksi</th>
          </tr>
        </thead>
        <tbody class="text-gray-800 text-sm">
          <tr
            v-for="(mhs, index) in mahasiswaList"
            :key="mhs.id || mhs.ID || mhs._id"
            class="border-t hover:bg-gray-50 transition"
          >
            <td class="px-4 py-3">{{ index + 1 }}</td>
            <td class="px-4 py-3 font-medium">{{ mhs.name }}</td>
            <td class="px-4 py-3">{{ mhs.email }}</td>
            <td class="px-4 py-3">{{ mhs.nim }}</td>
            <td class="px-4 py-3">
              <span
                :class="roleClass(mhs.role)"
                class="px-2 py-1 text-xs font-semibold rounded-full capitalize"
              >
                {{ mhs.role }}
              </span>
            </td>
            <td class="px-4 py-3">
              {{ new Date(mhs.created_at).toLocaleDateString("id-ID") }}
            </td>
            <td class="px-4 py-3">
              <button
                @click="openEditModal(mhs)"
                class="px-3 py-1 bg-yellow-400 text-white rounded-md hover:bg-yellow-500 mr-2 transition"
              >
                Edit
              </button>
              <button
                @click="hapusMahasiswa(mhs)"
                class="px-3 py-1 bg-red-500 text-white rounded-md hover:bg-red-600 transition"
              >
                Hapus
              </button>
            </td>
          </tr>
          <tr v-if="mahasiswaList.length === 0">
            <td colspan="7" class="px-4 py-4 text-center text-gray-500">
              Belum ada data mahasiswa.
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <div
      v-if="showModal"
      class="fixed inset-0 flex items-center justify-center bg-black bg-opacity-50 z-50 transition-opacity"
    >
      <div class="bg-white rounded-lg shadow-lg p-6 w-full max-w-md">
        <h3 class="text-lg font-semibold mb-4 border-b pb-2">
          {{ editMode ? "Edit Akun Pengguna" : "Tambah Akun Mahasiswa" }}
        </h3>

        <form @submit.prevent="simpanMahasiswa">
          <div class="mb-3">
            <label class="block mb-1 text-sm font-medium text-gray-700">Nama</label>
            <input
              v-model="form.name"
              type="text"
              required
              class="w-full p-2 border rounded-md"
            />
          </div>

          <div class="mb-3">
            <label class="block mb-1 text-sm font-medium text-gray-700">Email</label>
            <input
              v-model="form.email"
              type="email"
              required
              class="w-full p-2 border rounded-md"
            />
          </div>

          <template v-if="!editMode">
            <div class="mb-3">
              <label class="block mb-1 text-sm font-medium text-gray-700">Password</label>
              <input
                v-model="form.password"
                type="password"
                required
                class="w-full p-2 border rounded-md"
              />
            </div>
            <div class="mb-3">
              <label class="block mb-1 text-sm font-medium text-gray-700">NIM</label>
              <input
                v-model="form.nim"
                type="text"
                required
                class="w-full p-2 border rounded-md"
              />
            </div>
            <div class="mb-3">
              <label class="block mb-1 text-sm font-medium text-gray-700"
                >Jurusan (Major)</label
              >
              <input
                v-model="form.major"
                type="text"
                required
                class="w-full p-2 border rounded-md"
              />
            </div>
            <div class="mb-3">
              <label class="block mb-1 text-sm font-medium text-gray-700"
                >Fakultas (Faculty)</label
              >
              <input
                v-model="form.faculty"
                type="text"
                required
                class="w-full p-2 border rounded-md"
              />
            </div>
          </template>
          
          <template v-else>
            <div class="mb-3">
              <label class="block mb-1 text-sm font-medium text-gray-700">NIM</label>
              <input
                v-model="form.nim"
                type="text"
                required
                disabled 
                class="w-full p-2 border rounded-md bg-gray-100 cursor-not-allowed"
              />
            </div>
            <div class="mb-3">
              <label class="block mb-1 text-sm font-medium text-gray-700"
                >Password Baru (Opsional)</label
              >
              <input
                v-model="form.password"
                type="password"
                class="w-full p-2 border rounded-md"
                placeholder="Kosongkan jika tidak ingin diubah"
              />
            </div>
          </template>

          <div class="mb-3">
            <label class="block mb-1 text-sm font-medium text-gray-700">Role</label>
            <select
              v-model="form.role"
              required
              class="w-full p-2 border rounded-md bg-white"
            >
              <option value="user">Mahasiswa (user)</option>
              <option value="lecturer">Dosen (lecturer)</option>
              <option value="admin">Admin</option>
            </select>
          </div>

          <div class="flex justify-end gap-2 mt-4">
            <button
              type="button"
              @click="closeModal"
              class="px-4 py-2 bg-gray-300 text-gray-700 rounded-md hover:bg-gray-400 transition"
            >
              Batal
            </button>
            <button
              type="submit"
              class="px-4 py-2 bg-indigo-500 text-white rounded-md hover:bg-indigo-600 transition"
            >
              {{ editMode ? "Simpan Perubahan" : "Tambah" }}
            </button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from "vue";
// ## 1. HAPUS `useLocalStorage` ##
// import { useLocalStorage } from "../../hooks/useLocalStorage.js";
// ## 2. IMPOR `useGetCurrentUser` (state global) ##
import { useGetCurrentUser } from "../../hooks/useGetCurrentUser";
import {
  getUsersByRole,
  register,
  updateUser,
  deleteUser, 
  changeRole,
  changePassword,
} from "../../provider/user.provider.js";

const mahasiswaList = ref([]);
const loading = ref(true);
const error = ref(null);
const showModal = ref(false);
const editMode = ref(false);

const initialFormState = {
  id: null, name: "", email: "", password: "",
  role: "user", nim: "", major: "", faculty: "",
};

const form = ref({ ...initialFormState });
// ## 3. DAPATKAN PENGGUNA (Admin) DARI STATE GLOBAL ##
const { user: storedUser } = useGetCurrentUser();

const fetchMahasiswa = async () => {
  try {
    loading.value = true;
    const response = await getUsersByRole("user");
    mahasiswaList.value = response.data || [];
  } catch (err) {
    console.error("Gagal mengambil data mahasiswa:", err);
    error.value = "Tidak dapat memuat data. Silakan coba lagi nanti.";
  } finally {
    loading.value = false;
  }
};

onMounted(() => {
  fetchMahasiswa();
});

const openAddModal = () => {
  editMode.value = false;
  form.value = { ...initialFormState };
  showModal.value = true;
};

const closeModal = () => {
  showModal.value = false;
};

const simpanMahasiswa = async () => {
  const userId = form.value.id;
  // ## 4. 'adminId' SEKARANG DIAMBIL DARI STATE GLOBAL 'storedUser' ##
  const adminId = storedUser.value?.id || storedUser.value?.ID;

  if (!adminId) {
    alert("Error: Tidak dapat menemukan ID admin. Silakan login ulang.");
    return;
  }

  if (editMode.value) {
    // --- MODE EDIT ---
    if (!userId) {
      alert("Error: ID pengguna tidak ditemukan. Tidak dapat mengedit.");
      return;
    }

    let passwordErrorMessage = ""; 
    
    try {
      // Panggil changeRole TERLEBIH DAHULU
      await changeRole(userId, adminId, form.value.role);

      const dataToUpdate = { 
        name: form.value.name, 
        email: form.value.email,
        // 'nim' tidak dikirim karena tidak boleh diubah (immutable)
      };
      
      if (form.value.role === 'user') {
        dataToUpdate.major = form.value.major;
        dataToUpdate.faculty = form.value.faculty;
      }
      
      await updateUser(dataToUpdate, userId);

      // Coba ganti password HANYA JIKA diisi
      if (form.value.password && form.value.password.trim() !== "") {
        try {
          await changePassword(userId, form.value.password, adminId);
        } catch (passwordError) {
          console.warn("Gagal mengganti password (API backend belum siap):", passwordError);
          passwordErrorMessage = passwordError.response?.data?.message || "Gagal ganti password (backend error)";
        }
      }

      if (passwordErrorMessage) {
        alert(`Data (Nama, Email, Role) berhasil diperbarui.\n\nInfo: ${passwordErrorMessage}`);
      } else {
        alert("Data berhasil diperbarui!");
      }
      
      closeModal();
      fetchMahasiswa();

    } catch (err) {
      console.error("Gagal menyimpan data (Update/Role):", err);
      const errorMsg = err.response?.data?.message || "Terjadi kesalahan saat menyimpan data.";
      alert(errorMsg);
    }

  } else {
    // --- MODE TAMBAH ---
    try {
      const dataToCreate = {
        name: form.value.name,
        email: form.value.email,
        password: form.value.password,
        role: form.value.role,
        nim: form.value.nim,
        major: form.value.major,
        faculty: form.value.faculty,
      };
      await register(dataToCreate);
      alert("Mahasiswa baru berhasil ditambahkan!");
      closeModal();
      fetchMahasiswa();
    } catch (err) {
      console.error("Gagal menambah data:", err);
      const errorMsg = err.response?.data?.message || "Terjadi kesalahan saat menambah data.";
      alert(errorMsg);
    }
  }
};

const openEditModal = (mhs) => {
  editMode.value = true;
  const mhsData = { ...mhs };
  const userId = mhs.id || mhs.ID || mhs._id;
  form.value = {
    ...initialFormState,
    ...mhsData,
    id: userId,
    password: "", 
  };
  showModal.value = true;
};

const hapusMahasiswa = async (mhs) => {
  const userId = mhs.id || mhs.ID || mhs._id;
  if (!userId) {
    alert("Error: ID pengguna tidak ditemukan. Tidak dapat menghapus.");
    return;
  }

  if (confirm("Apakah Anda yakin ingin menghapus mahasiswa ini?")) {
    try {
      await deleteUser(userId);
      alert("Mahasiswa berhasil dihapus.");
      fetchMahasiswa();
    } catch (err) {
      console.error("Gagal menghapus mahasiswa:", err);
      const errorMsg = err.response?.data?.message || err.response?.data || "Terjadi kesalahan saat menghapus data.";
      alert(`Gagal menghapus: ${errorMsg}.`);
    }
  }
};

const roleClass = (role) => {
  if (role === "lecturer") return "bg-green-100 text-green-800";
  if (role === "user") return "bg-blue-100 text-blue-800";
  if (role === "admin") return "bg-red-100 text-red-800";
  return "bg-gray-100 text-gray-800";
};
</script>