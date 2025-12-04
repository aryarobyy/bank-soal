<template>
  <div class="p-4 md:p-6 bg-gray-50 min-h-screen">
    
    <div class="flex flex-col md:flex-row justify-between items-start md:items-center gap-4 mb-6">
      <h2 class="text-xl md:text-2xl font-bold text-gray-800">Manajemen Akun Mahasiswa</h2>
      
      <div class="flex flex-wrap gap-2 w-full md:w-auto">
        <button
          @click="openAddModal"
          class="flex-1 md:flex-none justify-center items-center gap-2 px-4 py-2 bg-indigo-500 text-white text-sm md:text-base rounded-lg hover:bg-indigo-600 transition shadow"
        >
          <i class="fas fa-plus-circle"></i> Tambah
        </button>
        
        <button
          @click="openGenerateModal"
          class="flex-1 md:flex-none justify-center items-center gap-2 px-4 py-2 bg-green-600 text-white text-sm md:text-base rounded-lg hover:bg-green-700 transition shadow"
        >
          <i class="fas fa-users"></i> Generate
        </button>
      </div>
    </div>

    <div v-if="loading" class="text-center py-10">
      <p class="text-gray-500">Memuat data mahasiswa...</p>
    </div>
    <div v-else-if="error" class="text-center py-10 bg-red-50 p-4 rounded-lg">
      <p class="text-red-600">{{ error }}</p>
    </div>

    <div v-else class="bg-white shadow rounded-lg overflow-hidden overflow-x-auto">
      <table class="min-w-full border-collapse">
        <thead class="bg-gray-100 text-gray-700 text-sm">
          <tr>
            <th class="px-4 py-3 text-left whitespace-nowrap">No</th>
            <th class="px-4 py-3 text-left whitespace-nowrap">Nama</th>
            <th class="px-4 py-3 text-left whitespace-nowrap">Email</th>
            <th class="px-4 py-3 text-left whitespace-nowrap">NIM</th>
            <th class="px-4 py-3 text-left whitespace-nowrap">Tahun Ajaran</th>
            <th class="px-4 py-3 text-left whitespace-nowrap">Role</th>
            <th class="px-4 py-3 text-left whitespace-nowrap">Tanggal Dibuat</th>
            <th class="px-4 py-3 text-left whitespace-nowrap">Aksi</th>
          </tr>
        </thead>
        <tbody class="text-gray-800 text-sm">
          <tr
            v-for="(mhs, index) in mahasiswaList"
            :key="mhs.id || mhs.ID || mhs._id"
            class="border-t hover:bg-gray-50 transition"
          >
            <td class="px-4 py-3 whitespace-nowrap">{{ (currentPage - 1) * itemsPerPage + index + 1 }}</td>
            <td class="px-4 py-3 whitespace-nowrap font-medium">{{ mhs.name }}</td>
            <td class="px-4 py-3 whitespace-nowrap">{{ mhs.email }}</td>
            <td class="px-4 py-3 whitespace-nowrap">{{ mhs.nim }}</td>
            <td class="px-4 py-3 whitespace-nowrap">{{ mhs.academic_year }}</td>
            <td class="px-4 py-3 whitespace-nowrap">
              <span
                :class="roleClass(mhs.role)"
                class="px-2 py-1 text-xs font-semibold rounded-full capitalize"
              >
                {{ mhs.role }}
              </span>
            </td>
            <td class="px-4 py-3 whitespace-nowrap">
              {{ new Date(mhs.created_at).toLocaleDateString("id-ID") }}
            </td>
            <td class="px-4 py-3 whitespace-nowrap">
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
            <td colspan="8" class="px-4 py-4 text-center text-gray-500">
              Belum ada data mahasiswa.
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <div v-if="!loading && totalPages > 1" class="flex flex-col md:flex-row justify-between items-center mt-6 gap-4">
      <span class="text-sm text-gray-700 text-center md:text-left">
        Halaman <span class="font-semibold">{{ currentPage }}</span> dari <span class="font-semibold">{{ totalPages }}</span> (Total <span class="font-semibold">{{ totalItems }}</span> mahasiswa)
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
      class="fixed inset-0 flex items-center justify-center bg-black/30 backdrop-blur-sm z-50 transition-opacity p-4"
    >
      <div class="bg-white rounded-lg shadow-lg p-6 w-full md:max-w-md max-h-[90vh] overflow-y-auto">
        <h3 class="text-lg font-semibold mb-4 border-b pb-2">
          {{ editMode ? "Edit Akun Pengguna" : "Tambah Akun Mahasiswa" }}
        </h3>

        <form @submit.prevent="simpanMahasiswa">
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
              <input v-model="form.password" type="password" required class="w-full p-2 border rounded-md"/>
            </div>
            <div class="mb-3">
              <label class="block mb-1 text-sm font-medium text-gray-700">NIM</label>
              <input v-model="form.nim" type="text" class="w-full p-2 border rounded-md" placeholder="Wajib diisi jika role 'user'"/>
            </div>
            <div class="mb-3">
              <label class="block mb-1 text-sm font-medium text-gray-700">Jurusan (Major)</label>
              <input v-model="form.major" type="text" class="w-full p-2 border rounded-md" placeholder="Wajib diisi jika role 'user'"/>
            </div>
            <div class="mb-3">
              <label class="block mb-1 text-sm font-medium text-gray-700">Fakultas (Faculty)</label>
              <input v-model="form.faculty" type="text" class="w-full p-2 border rounded-md" placeholder="Wajib diisi jika role 'user'"/>
            </div>
            <div class="mb-3">
              <label class="block mb-1 text-sm font-medium text-gray-700">Tahun Ajaran (Academic Year)</label>
              <input v-model="form.academic_year" type="text" class="w-full p-2 border rounded-md" placeholder="Wajib diisi jika role 'user'"/>
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
            
            <div v-if="form.role === 'user'">
              <div class="mb-3">
                <label class="block mb-1 text-sm font-medium text-gray-700">NIM</label>
                <input
                  v-model="form.nim"
                  type="text"
                  class="w-full p-2 border rounded-md" 
                  placeholder="Wajib diisi untuk mahasiswa"
                />
              </div>
              <div class="mb-3">
                <label class="block mb-1 text-sm font-medium text-gray-700"
                  >Jurusan (Major)</label
                >
                <input
                  v-model="form.major"
                  type="text"
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
                  class="w-full p-2 border rounded-md"
                />
              </div>
              <div class="mb-3">
                <label class="block mb-1 text-sm font-medium text-gray-700"
                  >Tahun Ajaran (Academic Year)</label
                >
                <input
                  v-model="form.academic_year"
                  type="text"
                  class="w-full p-2 border rounded-md"
                />
              </div>
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

            <div class="mb-3">
              <label class="block mb-1 text-sm font-medium text-gray-700"
                >Password Baru (Opsional)</label
              >
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
              <option value="user" selected>Mahasiswa (user)</option>
              <option value="lecturer">Dosen (lecturer)</option>
              </select>
          </div>

          <div class="flex justify-end gap-2 mt-4">
            <button type="button" @click="closeModal" class="px-4 py-2 bg-gray-300 text-gray-700 rounded-md hover:bg-gray-400 transition">
              Batal
            </button>
            <button type="submit" class="px-4 py-2 bg-indigo-500 text-white rounded-md hover:bg-indigo-600 transition">
              {{ editMode ? "Simpan Perubahan" : "Tambah" }}
            </button>
          </div>
        </form>
      </div>
    </div>
    
    <div
      v-if="showGenerateModal"
      class="fixed inset-0 flex items-center justify-center bg-black/30 backdrop-blur-sm z-50 transition-opacity p-4"
    >
      <div class="bg-white rounded-lg shadow-lg p-6 w-full md:max-w-md">
        <h3 class="text-lg font-semibold mb-4 border-b pb-2">
          Generate Akun Massal
        </h3>

        <form @submit.prevent="handleGenerate">
          <p class="text-sm text-gray-600 mb-4">
            Buat akun massal berdasarkan rentang NIM. File Excel berisi NIM dan
            password akan di-download secara otomatis.
          </p>

          <div class="mb-4">
            <label class="block mb-1 text-sm font-medium text-gray-700">
              Prefix NIM (Awalan)
            </label>
            <input
              v-model="generateForm.prefix"
              type="text"
              maxlength="2"
              required
              class="w-full p-2 border rounded-md"
              placeholder="Contoh: 25"
            />
            <p class="text-xs text-gray-500 mt-1">Wajib 2 karakter.</p>
          </div>

          <div class="mb-4">
            <label class="block mb-1 text-sm font-medium text-gray-700">
              Angka Mulai (Start)
            </label>
            <input
              v-model="generateForm.start" 
              type="number"
              required
              class="w-full p-2 border rounded-md"
              placeholder="Contoh: 301"
            />
          </div>

          <div class="mb-4">
            <label class="block mb-1 text-sm font-medium text-gray-700">
              Angka Akhir (End)
            </label>
            <input
              v-model="generateForm.end"
              type="number"
              required
              class="w-full p-2 border rounded-md"
              placeholder="Contoh: 350"
            />
          </div>

          <div class="mb-4">
            <label class="block mb-1 text-sm font-medium text-gray-700">
              Tahun Ajaran (Academic Year)
            </label>
            <input
              v-model="generateForm.academic_year"
              type="text"
              required
              class="w-full p-2 border rounded-md"
              placeholder="Contoh: 2025/2026"
            />
            <p class="text-xs text-gray-500 mt-1">Wajib diisi.</p>
          </div>

          <div class="flex justify-end gap-2 mt-4">
            <button
              type="button"
              @click="closeGenerateModal"
              class="px-4 py-2 bg-gray-300 text-gray-700 rounded-md hover:bg-gray-400 transition"
            >
              Batal
            </button>
            <button
              type="submit"
              :disabled="isLoadingGenerate"
              class="px-4 py-2 bg-green-600 text-white rounded-md hover:bg-green-700 transition disabled:bg-gray-400"
            >
              {{ isLoadingGenerate ? "Memproses..." : "Generate & Download" }}
            </button>
          </div>
        </form>
      </div>
    </div>

  </div>
</template>

<script setup>

import { ref, onMounted, computed, watch } from "vue";
import { useGetCurrentUser } from "../../hooks/useGetCurrentUser";
import {
  getUsersByRole,
  register,
  updateUser,
  deleteUser,
  changePassword,
  generateUsers,
} from "../../provider/user.provider.js";

const mahasiswaList = ref([]);
const loading = ref(true);
const error = ref(null);
const showModal = ref(false);
const editMode = ref(false);

const initialFormState = {
  id: null, name: "", email: "", password: "",
  role: "user", nim: "", major: "", faculty: "", academic_year: "",
  nip: "", nidn: ""
};

const form = ref({ ...initialFormState });
const { user: storedUser } = useGetCurrentUser();

const showGenerateModal = ref(false);
const isLoadingGenerate = ref(false);

const generateForm = ref({
  prefix: "25",
  start: '',
  end: '',
  academic_year: "",
});

// State Paginasi
const currentPage = ref(1);
const itemsPerPage = ref(10);
const totalItems = ref(0);

const totalPages = computed(() => {
  return Math.ceil(totalItems.value / itemsPerPage.value);
});

const fetchMahasiswa = async () => {
  try {
    loading.value = true;
    const offset = (currentPage.value - 1) * itemsPerPage.value;
    const response = await getUsersByRole("user", itemsPerPage.value, offset); 
    mahasiswaList.value = response.data || [];
    totalItems.value = response.total || 0;
    error.value = null;

  } catch (err) {
    console.error("Gagal mengambil data mahasiswa:", err);
    error.value = "Tidak dapat memuat data. Silakan coba lagi nanti.";
    mahasiswaList.value = [];
    totalItems.value = 0;
  } finally {
    loading.value = false;
  }
};

onMounted(() => {
  fetchMahasiswa(); 
});

watch(currentPage, (newPage, oldPage) => {
  if (newPage !== oldPage) {
    fetchMahasiswa();
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
  form.value = { ...initialFormState, role: "user" };
  showModal.value = true;
};

const closeModal = () => {
  showModal.value = false;
};

const openGenerateModal = () => {
  generateForm.value = { prefix: "25", start: '', end: '', academic_year: "" };
  showGenerateModal.value = true;
};

const closeGenerateModal = () => {
  showGenerateModal.value = false;
};

const handleGenerate = async () => {
  const { prefix, start, end, academic_year } = generateForm.value;

  if (!prefix || !start || !end || !academic_year) {
    alert("Harap isi semua field (Prefix, Start, End, dan Tahun Ajaran).");
    return;
  }
  if (prefix.length !== 2) {
    alert("Prefix harus 2 karakter.");
    return;
  }
  if (Number(end) < Number(start)) {
    alert("Angka Akhir tidak boleh lebih kecil dari Angka Mulai.");
    return;
  }

  isLoadingGenerate.value = true;

  try {
    const response = await generateUsers(prefix, start, end, academic_year);
    alert(`Sukses: ${response.message}\nFile disimpan di server sebagai: ${response.data.file}`);
    console.log("Pengguna yang di-generate:", response.data.users);
    
    closeGenerateModal();
    currentPage.value = 1; 
    fetchMahasiswa(); 

  } catch (error) {
    console.error('Gagal generate user:', error);
    alert(`Gagal: ${error.response?.data?.message || 'Terjadi kesalahan.'}`);

  } finally {
    isLoadingGenerate.value = false;
  }
};

const simpanMahasiswa = async () => {
  const userId = form.value.id;
  const adminId = storedUser.value?.id || storedUser.value?.ID;

  if (!adminId) {
    alert("Error: Tidak dapat menemukan ID admin. Silakan login ulang.");
    return;
  }

  if (editMode.value) {
    if (!userId) {
      alert("Error: ID pengguna tidak ditemukan. Tidak dapat mengedit.");
      return;
    }

    let passwordErrorMessage = ""; 
    
    try {
      const dataToUpdate = { 
        name: form.value.name, 
        email: form.value.email,
        role: form.value.role,
      };

      if (form.value.role === 'user') {
        dataToUpdate.nim = form.value.nim || null;
        dataToUpdate.major = form.value.major || null;
        dataToUpdate.faculty = form.value.faculty || null;
        dataToUpdate.academic_year = form.value.academic_year || null;
        dataToUpdate.nip = null;
        dataToUpdate.nidn = null;
      } else if (form.value.role === 'lecturer') {
        dataToUpdate.nip = form.value.nip || null;
        dataToUpdate.nidn = form.value.nidn || null;
        dataToUpdate.nim = null;
        dataToUpdate.major = null;
        dataToUpdate.faculty = null;
        dataToUpdate.academic_year = null;
      }
      
      await updateUser(dataToUpdate, userId);

      if (form.value.password && form.value.password.trim() !== "") {
        try {
          await changePassword(userId, form.value.password, adminId);
        } catch (passwordError) {
          console.warn("Gagal mengganti password:", passwordError);
          passwordErrorMessage = passwordError.response?.data?.message || "Gagal ganti password";
        }
      }

      if (passwordErrorMessage) {
        alert(`Data berhasil diperbarui.\n\nInfo: ${passwordErrorMessage}`);
      } else {
        alert("Data berhasil diperbarui!");
      }
      
      closeModal();
      fetchMahasiswa(); 

    } catch (err) {
      console.error("Gagal menyimpan data (Update):", err);
      const errorMsg = err.response?.data?.message || "Terjadi kesalahan saat menyimpan data.";
      alert(errorMsg);
      fetchMahasiswa();
    }

  } else {
    try {
      const dataToCreate = {
        name: form.value.name,
        email: form.value.email,
        password: form.value.password,
        role: form.value.role,
      };

      if (form.value.role === 'user') {
        dataToCreate.nim = form.value.nim || null;
        dataToCreate.major = form.value.major || null;
        dataToCreate.faculty = form.value.faculty || null;
        dataToCreate.academic_year = form.value.academic_year || null;
        dataToCreate.nip = null;
        dataToCreate.nidn = null;
      } else if (form.value.role === 'lecturer') {
        dataToCreate.nip = form.value.nip || null;
        dataToCreate.nidn = form.value.nidn || null;
        dataToCreate.nim = null;
        dataToCreate.major = null;
        dataToCreate.faculty = null;
        dataToCreate.academic_year = null;
      }

      await register(dataToCreate);
      alert("Akun baru berhasil ditambahkan!");
      closeModal();
      
      try {
        const response = await getUsersByRole("user", 1, 0);
        totalItems.value = response.total || 0;
        currentPage.value = totalPages.value; 
      } catch (e) {
        fetchMahasiswa(); 
      }

    } catch (err)
 {
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
    nim: mhs.nim || "",
    major: mhs.major || "", 
    faculty: mhs.faculty || "",
    academic_year: mhs.academic_year || "",
    nip: mhs.nip || "",
    nidn: mhs.nidn || "",
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
      
      if (mahasiswaList.value.length === 1 && currentPage.value > 1) {
        currentPage.value--;
      } else {
        fetchMahasiswa();
      }

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