<template>
  <div class="p-6 bg-gray-50 min-h-screen">
    <div class="flex justify-between items-center mb-6">
      <h2 class="text-2xl font-bold text-gray-800">Manajemen Akun Dosen</h2>
      <button
        @click="openAddModal"
        class="px-4 py-2 bg-indigo-500 text-white rounded-lg hover:bg-indigo-600 transition"
      >
        + Tambah Dosen
      </button>
    </div>

    <div v-if="loading" class="text-center py-10">
      <p class="text-gray-500">Memuat data dosen...</p>
    </div>
    <div v-else-if="error" class="text-center py-10 bg-red-50 p-4 rounded-lg">
      <p class="text-red-600">{{ error }}</p>
    </div>
    <div v-else class="bg-white shadow rounded-lg overflow-hidden overflow-x-auto">
      <table class="min-w-full border-collapse">
        <thead class="bg-gray-100 text-gray-700 text-sm">
          <tr>
            <th class="px-4 py-3 text-left">No</th>
            <th class="px-4 py-3 text-left">Nama</th>
            <th class="px-4 py-3 text-left">Email</th>
            <th class="px-4 py-3 text-left">NIP</th>
            <th class="px-4 py-3 text-left">NIDN</th>
            <th class="px-4 py-3 text-left">Role</th>
            <th class="px-4 py-3 text-left">Tanggal Dibuat</th>
            <th class="px-4 py-3 text-left">Aksi</th>
          </tr>
        </thead>
        <tbody class="text-gray-800 text-sm">
          <tr
            v-for="(dosen, index) in dosenList"
            :key="dosen.id || dosen.ID || dosen._id"
            class="border-t hover:bg-gray-50 transition"
          >
            <td class="px-4 py-3">{{ (currentPage - 1) * itemsPerPage + index + 1 }}</td>
            <td class="px-4 py-3 font-medium">{{ dosen.name }}</td>
            <td class="px-4 py-3">{{ dosen.email }}</td>
            <td class="px-4 py-3">{{ dosen.nip }}</td>
            <td class="px-4 py-3">{{ dosen.nidn }}</td>
            <td class="px-4 py-3">
              <span 
                :class="roleClass(dosen.role)" 
                class="px-2 py-1 text-xs font-semibold rounded-full capitalize"
              >
                {{ dosen.role }}
              </span>
            </td>
            <td class="px-4 py-3">{{ new Date(dosen.created_at).toLocaleDateString("id-ID") }}</td>
            <td class="px-4 py-3">
              <button
                @click="editDosen(dosen)"
                class="px-3 py-1 bg-yellow-400 text-white rounded-md hover:bg-yellow-500 mr-2 transition"
              >
                Edit
              </button>
              <button
                @click="hapusDosen(dosen)"
                class="px-3 py-1 bg-red-500 text-white rounded-md hover:bg-red-600 transition"
              >
                Hapus
              </button>
            </td>
          </tr>
          <tr v-if="dosenList.length === 0">
            <td colspan="8" class="px-4 py-4 text-center text-gray-500">
              Belum ada data dosen
            </td>
          </tr>
        </tbody>
      </table>
    </div>
    <div v-if="!loading && totalPages > 1" class="flex justify-between items-center mt-6">
      <span class="text-sm text-gray-700">
        Halaman <span class="font-semibold">{{ currentPage }}</span> dari <span class="font-semibold">{{ totalPages }}</span> (Total <span class="font-semibold">{{ totalItems }}</span> dosen)
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
      <div class="bg-white rounded-lg shadow-lg p-6 w-11/12 md:w-full max-w-md">
        <h3 class="text-lg font-semibold mb-4">
          {{ editMode ? "Edit Akun" : "Tambah Dosen" }}
        </h3>

        <form @submit.prevent="simpanDosen">
          <div class="mb-3">
            <label class="block mb-1 text-sm font-medium text-gray-700">Nama</label>
            <input v-model="form.name" type="text" required class="w-full p-2 border rounded-md"/>
          </div>
          <div class="mb-3">
            <label class="block mb-1 text-sm font-medium text-gray-700">Email</label>
            <input v-model="form.email" type="email" required class="w-full p-2 border rounded-md"/>
          </div>

          <template v-if="!editMode">
            <div class="mb-3">
              <label class="block mb-1 text-sm font-medium text-gray-700">Password</label>
              <input 
                v-model="form.password" 
                type="password" 
                autocomplete="new-password"
                required 
                class="w-full p-2 border rounded-md"
              />
            </div>
            <div class="mb-3">
              <label class="block mb-1 text-sm font-medium text-gray-700">NIP</label>
              <input v-model="form.nip" type="text" required class="w-full p-2 border rounded-md"/>
            </div>
            <div class="mb-3">
              <label class="block mb-1 text-sm font-medium text-gray-700">NIDN</label>
              <input v-model="form.nidn" type="text" required class="w-full p-2 border rounded-md"/>
            </div>
            <div class="mb-3">
              <label class="block mb-1 text-sm font-medium text-gray-700">Jurusan (Major)</label>
              <input v-model="form.major" type="text" class="w-full p-2 border rounded-md" placeholder="Opsional (Isi N/A jika tidak ada)"/>
            </div>
            <div class="mb-3">
              <label class="block mb-1 text-sm font-medium text-gray-700">Fakultas (Faculty)</label>
              <input v-model="form.faculty" type="text" class="w-full p-2 border rounded-md" placeholder="Opsional (Isi N/A jika tidak ada)"/>
            </div>
          </template>
          
          <template v-else>
            <div class="mb-3">
              <label class="block mb-1 text-sm font-medium text-gray-700">Role</label>
              <select v-model="form.role" required class="w-full p-2 border rounded-md bg-white">
                <option value="user">Mahasiswa (user)</option>
                <option value="lecturer">Dosen (lecturer)</option>
              </select>
            </div>

            <div v-if="form.role === 'lecturer'">
              <div class="mb-3">
                <label class="block mb-1 text-sm font-medium text-gray-700">NIP</label>
                <input v-model="form.nip" type="text" class="w-full p-2 border rounded-md" placeholder="Wajib diisi untuk dosen"/>
              </div>
              <div class="mb-3">
                <label class="block mb-1 text-sm font-medium text-gray-700">NIDN</label>
                <input v-model="form.nidn" type="text" class="w-full p-2 border rounded-md" placeholder="Wajib diisi untuk dosen"/>
              </div>
            </div>

            <div v-if="form.role === 'user'">
              <div class="mb-3">
                <label class="block mb-1 text-sm font-medium text-gray-700">NIM</label>
                <input v-model="form.nim" type="text" class="w-full p-2 border rounded-md" placeholder="Wajib diisi untuk mahasiswa"/>
              </div>
              <div class="mb-3">
                <label class="block mb-1 text-sm font-medium text-gray-700">Jurusan (Major)</label>
                <input v-model="form.major" type="text" class="w-full p-2 border rounded-md" placeholder="Wajib diisi untuk mahasiswa"/>
              </div>
              <div class="mb-3">
                <label class="block mb-1 text-sm font-medium text-gray-700">Fakultas (Faculty)</label>
                <input v-model="form.faculty" type="text" class="w-full p-2 border rounded-md" placeholder="Wajib diisi untuk mahasiswa"/>
              </div>
              <div class="mb-3">
                <label class="block mb-1 text-sm font-medium text-gray-700">Tahun Ajaran (Academic Year)</label>
                <input v-model="form.academic_year" type="text" class="w-full p-2 border rounded-md" placeholder="Wajib diisi untuk mahasiswa"/>
              </div>
            </div>
            
            <div class="mb-3">
              <label class="block mb-1 text-sm font-medium text-gray-700">Password Baru (Opsional)</label>
              <input
                v-model="form.password"
                type="password"
                autocomplete="new-password"
                class="w-full p-2 border rounded-md"
                placeholder="Kosongkan jika tidak ingin diubah"
              />
            </div>
          </template>

          <div class="mb-3" v-if="!editMode">
            <label class="block mb-1 text-sm font-medium text-gray-700">Role</label>
            <select v-model="form.role" required class="w-full p-2 border rounded-md bg-white">
              <option value="lecturer" selected>Dosen (lecturer)</option>
              <option value="user">Mahasiswa (user)</option>
              </select>
          </div>

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
} from "../../provider/user.provider.js";
import { useGetCurrentUser } from "../../hooks/useGetCurrentUser";

const dosenList = ref([]);
const loading = ref(true);
const error = ref(null);
const showModal = ref(false);
const editMode = ref(false);


const initialFormState = {
  id: null, name: "", email: "", password: "",
  role: "lecturer", nip: "", nidn: "", 
  major: "", faculty: "", nim: "", academic_year: "" 
};
const form = ref({ ...initialFormState });

const { user: storedUser } = useGetCurrentUser();

const currentPage = ref(1);
const itemsPerPage = ref(10);
const totalItems = ref(0);

const totalPages = computed(() => {
  return Math.ceil(totalItems.value / itemsPerPage.value);
});

const fetchDosen = async () => {
  try {
    loading.value = true;
    const offset = (currentPage.value - 1) * itemsPerPage.value;
    const response = await getUsersByRole("lecturer", itemsPerPage.value, offset);
    dosenList.value = response.data || [];
    totalItems.value = response.total || 0;
    error.value = null; 
  } catch (err) {
    console.error("Gagal mengambil data dosen:", err);
    error.value = "Tidak dapat memuat data. Silakan coba lagi nanti.";
    dosenList.value = [];
    totalItems.value = 0;
  } finally {
    loading.value = false;
  }
};

onMounted(() => {
  fetchDosen();
});

watch(currentPage, (newPage, oldPage) => {
  if (newPage !== oldPage) {
    fetchDosen();
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
  form.value = { ...initialFormState, role: "lecturer" };
  showModal.value = true;
};

const closeModal = () => {
  showModal.value = false;
};


const simpanDosen = async () => {
  try {
    const userId = form.value.id;
    const adminId = storedUser.value?.id || storedUser.value?.ID;

    if (!adminId) {
      alert("Error: Sesi Admin tidak ditemukan. Silakan login ulang.");
      return;
    }

    if (editMode.value) {
  
      if (!userId) {
        alert("Error: ID pengguna tidak ditemukan. Tidak dapat mengedit.");
        return;
      }
      
      let passwordSuccess = false;
      let passwordErrorMessage = "";

      const dataToUpdate = { 
        name: form.value.name, 
        email: form.value.email,
        role: form.value.role,
      };

      if (form.value.role === 'lecturer') {
        dataToUpdate.nip = form.value.nip || null;
        dataToUpdate.nidn = form.value.nidn || null;
        dataToUpdate.nim = null;
        dataToUpdate.major = null;
        dataToUpdate.faculty = null;
      
      } else if (form.value.role === 'user') {
        dataToUpdate.nim = form.value.nim || null;
        dataToUpdate.major = form.value.major || null;
        dataToUpdate.faculty = form.value.faculty || null;
        dataToUpdate.academic_year = form.value.academic_year || null;
        dataToUpdate.nip = null;
        dataToUpdate.nidn = null;
      }
      
     
      await updateUser(dataToUpdate, userId);
      
 
      if (form.value.password && form.value.password.trim() !== "") {
        try {
          await changePassword(userId, form.value.password, adminId);
          passwordSuccess = true;
        } catch (passwordError) {
          console.error("Gagal mengganti password:", passwordError);
          passwordErrorMessage = passwordError.response?.data?.message || "Gagal ganti password";
        }
      }
      
   
      if (passwordErrorMessage) {
        alert(`Data berhasil diperbarui, TAPI: ${passwordErrorMessage}`);
      } else if (passwordSuccess) {
         alert("Data dan password berhasil diperbarui!");
      } else {
         alert("Data berhasil diperbarui!");
      }

    } else {
    
      const dataToCreate = {
        name: form.value.name,
        email: form.value.email,
        password: form.value.password,
        role: form.value.role,
      };

      if (form.value.role === 'lecturer') {
        dataToCreate.nip = form.value.nip || null;
        dataToCreate.nidn = form.value.nidn || null;
        dataToCreate.major = form.value.major || "N/A";
        dataToCreate.faculty = form.value.faculty || "N/A";
        dataToCreate.nim = null;
      } else if (form.value.role === 'user') {
        dataToCreate.nim = form.value.nim || null;
        dataToCreate.major = form.value.major || null;
        dataToCreate.faculty = form.value.faculty || null;
        dataToCreate.academic_year = form.value.academic_year || null;
        dataToCreate.nip = null;
        dataToCreate.nidn = null;
      }

      await register(dataToCreate);
      alert("Akun baru berhasil ditambahkan!");
    }

    closeModal();
    
    if (!editMode.value) {
      try {
        const response = await getUsersByRole("lecturer", 1, 0);
        totalItems.value = response.total || 0;
        currentPage.value = totalPages.value; 
      } catch (e) {
        fetchDosen(); 
      }
    } else {
      fetchDosen(); 
    }

  } catch (err) {
    console.error("Gagal menyimpan data:", err);
    const errorMsg = err.response?.data?.message || "Terjadi kesalahan saat menyimpan data.";
    alert(errorMsg);
  }
};

const editDosen = (dosen) => {
  editMode.value = true;
  const userId = dosen.id || dosen.ID || dosen._id;
  
  form.value = { 
    ...initialFormState, 
    ...dosen,
    nip: dosen.nip || "",
    nidn: dosen.nidn || "",
    nim: dosen.nim || "",
    major: dosen.major || "",
    faculty: dosen.faculty || "",
    academic_year: dosen.academic_year || "",
    id: userId,          
    password: ""         
  };
  showModal.value = true;
};

const hapusDosen = async (dosen) => {
  const userId = dosen.id || dosen.ID || dosen._id;
  if (!userId) {
    alert("Error: ID pengguna tidak ditemukan. Tidak dapat menghapus.");
    return;
  }
  
  if (confirm("Yakin ingin menghapus dosen ini?")) {
    try {
      await deleteUser(userId);
      alert("Dosen berhasil dihapus.");
      
      if (dosenList.value.length === 1 && currentPage.value > 1) {
        currentPage.value--;
      } else {
        fetchDosen();
      }

    } catch (err) {
      console.error("Gagal menghapus dosen:", err);
      const errorMsg = err.response?.data?.message || err.response?.data || "Gagal menghapus data.";
      alert(`Gagal menghapus: ${errorMsg}.`);
    }
  }
};

const roleClass = (role) => {
  if (role === 'lecturer') return 'bg-green-100 text-green-800';
  if (role === 'user') return 'bg-blue-100 text-blue-800';
  if (role === 'admin') return 'bg-red-100 text-red-800';
  return 'bg-gray-100 text-gray-800';
};
</script>