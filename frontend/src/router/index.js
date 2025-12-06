import { createRouter, createWebHistory } from 'vue-router'

// Layouts
import UserLayout from '../layouts/UserLayout.vue'
import DosenLayout from '../layouts/DosenLayout.vue'
import AdminLayout from '../layouts/AdminLayout.vue'
import RouterWrapper from '../layouts/RouterWrapper.vue'
import SuperAdminLayout from '../layouts/SuperAdminLayout.vue'

// Halaman-halaman Umum & Autentikasi
import LoginView from '../views/Auth/LoginView.vue'
import HomePage from '../views/home/HomeView.vue'
import Profile from '../views/profile/Profile.vue'
import UpdateProfile from '../views/profile/UpdateProfile.vue'

// --- HALAMAN UJIAN USER ---
import UserExamList from '../views/exam/UserExamList.vue' 
import ExamView from '../views/exam/ExamView.vue'        
import ExamDo from '../views/exam/ExamDo.vue'            
import UserDashboard from '../views/home/UserDashboard.vue'

// Halaman Dosen
import LecturerDashboard from '../views/dosen/LecturerDashboard.vue'

// Halaman Admin
import AdminDashboard from '../views/Admin/AdminDashboard.vue'
import ManageMahasiswa from '../views/ManageMahasiswa/ManageMahasiswa.vue'
import ManageDosen from '../views/ManageDosen/ManageDosen.vue'
import ManageXlsFiles from '../views/Admin/ManageXlsFiles.vue'

// Halaman Soal (Shared)
import LecturerSoal from '../views/question/LecturerSoal.vue'
import LecturerSoalList from '../views/question/LecturerSoalList.vue'
import CreateSoal from '../views/question/CreateSoal.vue'
import CreateManualSoal from '../views/question/CreateManualSoal.vue'
import UploadJsonSoal from '../views/question/UploadJsonSoal.vue'

// Halaman Manajemen Ujian (Admin/Dosen)
import ManageExam from '../views/exam/ManageExam.vue'
import CreateExamView from '../views/exam/CreateExamView.vue'
import ExamReports from '../views/exam/ExamReports.vue'
import ExamDetail from '../views/exam/ExamDetail.vue'
import ExamEdit from '../views/exam/ExamEdit.vue'

// Halaman Super Admin
import SuperAdminDashboard from '../views/SuperAdmin/SuperAdminDashboard.vue'
import ManageAdmin from '../views/ManageAdmin/ManageAdmin.vue'


// --- GRUP RUTE SOAL (Shared) ---
const soalRoutes = [
  { path: '', name: 'SoalHome', component: LecturerSoal, },
  { path: 'list', name: 'SoalList', component: LecturerSoalList, },
  { path: 'create', name: 'SoalCreate', component: CreateSoal, },
  { path: 'create-manual', name: 'SoalCreateManual', component: CreateManualSoal, },
  { path: 'edit/:id', name: 'SoalEdit', component: CreateManualSoal, },
  { path: 'upload-json', name: 'SoalUploadJson', component: UploadJsonSoal, }
]

// --- GRUP RUTE MANAJEMEN UJIAN (Shared) ---
const examRoutes = [
  { path: '', name: 'ManageExam', component: ManageExam },
  { path: 'create', name: 'CreateExam', component: CreateExamView },
  { path: 'detail/:id', name: 'ExamDetail', component: ExamDetail },
  { path: 'edit/:id', name: 'ExamEdit', component: ExamEdit }
]

const routes = [
  
  { path: '/login', name: 'login', component: LoginView },
  

  {
    path: "/exam/start",
    name: "ExamDo",
    component: ExamDo,
  },


  {
    path: '/',
    component: UserLayout,
    children: [
      { path: 'landing', name: 'LandingPage', component: HomePage },
      
     
      { path: 'ujian', name: 'UserExamList', component: UserExamList },
    
 
      { 
        path: 'exam/view', 
        name: 'UserExamView', 
        component: ExamView,
        props: route => ({ id: route.query.id }) 
      },

      { path: 'profile/:id', name: 'Profile', component: Profile },
      { path: 'update-profile', name: 'UpdateProfile', component: UpdateProfile },
      { 
        path: '', 
        name: 'Dashboard', 
        component: UserDashboard 
      },
    ]
  },


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
      
   
      {
        path: 'exam', 
        component: RouterWrapper,
        children: examRoutes.map(route => ({ ...route, name: `Dosen${route.name}` }))
      }
    ]
  },


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

      {
        path: 'ujian', 
        component: RouterWrapper,
        children: examRoutes.map(route => ({ ...route, name: `Admin${route.name}` }))
      },
      

      {
        path: 'reports',
        name: 'AdminExamReports',
        component: ExamReports
      },
      
      { path: 'update-profile', name: 'AdminUpdateProfile', component: UpdateProfile },
      
      {
        path: 'excel-files',
        name: 'AdminXlsFiles',
        component: ManageXlsFiles,
        meta: { requiresAuth: true, role: 'admin' }
      }
    ]
  },


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