import { createRouter, createWebHistory } from 'vue-router'

// Layouts
import UserLayout from '../layouts/UserLayout.vue'
import DosenLayout from '../layouts/DosenLayout.vue'
import AdminLayout from '../layouts/AdminLayout.vue'
import RouterWrapper from '../layouts/RouterWrapper.vue'
import SuperAdminLayout from '../layouts/SuperAdminLayout.vue'

// Halaman-halaman Umum & Autentikasi
import LoginView from '../views/Auth/LoginView.vue'
import RegisterView from '../views/Auth/RegisterView.vue'
import ForgotPassword from '../views/Auth/ForgotPassword.vue'
import HomePage from '../views/home/HomeView.vue'
import ExamView from '../views/exam/ExamView.vue'
import Profile from '../views/profile/Profile.vue'
import UpdateProfile from '../views/profile/UpdateProfile.vue'

// Halaman Dosen
import LecturerDashboard from '../views/dosen/LecturerDashboard.vue'

// Halaman Admin
import AdminDashboard from '../views/Admin/AdminDashboard.vue'
import AdminProfile from '../views/Admin/AdminProfile.vue'
import ManageMahasiswa from '../views/ManageMahasiswa/ManageMahasiswa.vue'
import ManageDosen from '../views/ManageDosen/ManageDosen.vue'

// Halaman Soal (yang akan dipakai bersama)
import LecturerSoal from '../views/question/LecturerSoal.vue'
import LecturerSoalList from '../views/question/LecturerSoalList.vue'
import CreateSoal from '../views/question/CreateSoal.vue'
import CreateManualSoal from '../views/question/CreateManualSoal.vue'
import UploadJsonSoal from '../views/question/UploadJsonSoal.vue'

// Halaman Ujian (yang akan dipakai bersama)
import ManageExam from '../views/exam/ManageExam.vue'
import CreateExamView from '../views/exam/CreateExamView.vue'
import ExamReports from '../views/exam/ExamReports.vue'
// ## 1. Impor 2 Halaman Baru Anda ##
import ExamDetail from '../views/exam/ExamDetail.vue'
import ExamEdit from '../views/exam/ExamEdit.vue'

// Halaman Super Admin
import SuperAdminDashboard from '../views/SuperAdmin/SuperAdminDashboard.vue'
import ManageAdmin from '../views/ManageAdmin/ManageAdmin.vue'


// Definisikan semua rute "soal" di satu tempat
const soalRoutes = [
  { path: '', name: 'SoalHome', component: LecturerSoal, },
  { path: 'list', name: 'SoalList', component: LecturerSoalList, },
  { path: 'create', name: 'SoalCreate', component: CreateSoal, },
  { path: 'create-manual', name: 'SoalCreateManual', component: CreateManualSoal, },
  { path: 'edit/:id', name: 'SoalEdit', component: CreateManualSoal, },
  { path: 'upload-json', name: 'SoalUploadJson', component: UploadJsonSoal, }
]

// ## 2. Buat Grup Rute Ujian (BARU) ##
const examRoutes = [
  { path: '', name: 'ManageExam', component: ManageExam },
  { path: 'create', name: 'CreateExam', component: CreateExamView },
  { path: 'detail/:id', name: 'ExamDetail', component: ExamDetail },
  { path: 'edit/:id', name: 'ExamEdit', component: ExamEdit }
]

const routes = [
  // --- Rute Autentikasi & Pengguna Umum ---
  // (Tidak ada perubahan di sini)
  { path: '/login', name: 'login', component: LoginView },
  { path: '/register', name: 'register', component: RegisterView },
  { path: '/forgot-password', name: 'ForgotPassword', component: ForgotPassword },
  
    {
    path: "/exam/start",
    name: "ExamDo",
    component: () => import("../views/exam/ExamDo.vue"),
  },
  {
    path: '/',
    component: UserLayout,
    children: [
      { path: '', name: 'HomePage', component: HomePage },
      { path: 'ujian', name: 'ExamPage', component: ExamView },
      { path: 'profile/:id', name: 'Profile', component: Profile },
      { path: 'update-profile', name: 'UpdateProfile', component: UpdateProfile }
      
    ]
  },

  // --- Rute Dosen ---
  {
    path: '/dosen',
    component: DosenLayout,
    redirect: '/dosen/dashboard',
    children: [
      { path: 'dashboard', name: 'LecturerDashboard', component: LecturerDashboard },
      { path: 'profile/:id', name: 'DosenProfile', component: Profile },
      { path: 'update-profile', name: 'DosenUpdateProfile', component: UpdateProfile },
      {
        path: 'soal',
        component: RouterWrapper, 
        children: soalRoutes.map(route => ({ ...route, name: `Dosen${route.name}` })) 
      },
      // ## 3. Tambahkan Rute Ujian ke Dosen (BARU) ##
      {
        path: 'exam', // akan menjadi /dosen/exam
        component: RouterWrapper,
        children: examRoutes.map(route => ({ ...route, name: `Dosen${route.name}` }))
      }
    ]
  },

  // --- Rute Admin ---
  {
    path: '/admin',
    component: AdminLayout,
    redirect: '/admin/dashboard',
    children: [
      { path: 'dashboard', name: 'AdminDashboard', component: AdminDashboard },
      { path: 'profile/:id', name: 'AdminProfile', component: Profile },
      { path: 'mahasiswa', name: 'AdminManageMahasiswa', component: ManageMahasiswa },
      { path: 'dosen', name: 'AdminManageDosen', component: ManageDosen },
      {
        path: 'soal',
        component: RouterWrapper, 
        children: soalRoutes.map(route => ({ ...route, name: `Admin${route.name}` }))
      },
      // ## 4. Ganti Rute Ujian Admin dengan Grup Rute (DIPERBARUI) ##
      {
        path: 'ujian', // akan menjadi /admin/ujian
        component: RouterWrapper,
        children: examRoutes.map(route => ({ ...route, name: `Admin${route.name}` }))
      },
      {
        path: 'reports',
        name: 'AdminExamReports',
        component: ExamReports
      },
      { 
        path: 'update-profile', 
        name: 'AdminUpdateProfile', 
        component: UpdateProfile
      },
      {
      path: 'excel-files',
      name: 'AdminXlsFiles',
      component: () => import('../views/Admin/ManageXlsFiles.vue'),
      meta: { requiresAuth: true, role: 'admin' }
     }
    ]
  },

  // --- Rute Super Admin ---
  // (Tidak ada perubahan di sini)
  {
    path: '/superadmin',
    component: SuperAdminLayout, 
    redirect: '/superadmin/dashboard',
    children: [
      {
        path: 'dashboard', 
        name: 'SuperAdminDashboard',
        component: SuperAdminDashboard, 
        meta: { title: 'Dashboard Aktivitas' }
      },
      {
        path: 'admins', 
        name: 'SuperAdminManageAdmins',
        component: ManageAdmin,
        meta: { title: 'Manajemen Admin' }
      },
    ]
  }
]

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes
})

export default router