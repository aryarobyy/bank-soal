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
import SoalHome from '../views/question/SoalHome.vue'
import SoalList from '../views/question/SoalList.vue'
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
  { path: '', name: 'SoalHome', component: SoalHome, },
  { path: 'list', name: 'SoalList', component: SoalList, },
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
      path: '/landing',
      component: UserLayout, 
   
      children: [
        { 
          path: '', 
          name: 'LandingPage', 
          component: HomePage 
        }
      ]
    },
  {
    path: "/exam/start",
    name: "ExamDo",
    component: ExamDo,
  },


  {
    path: '/',
    component: UserLayout,
    meta: { requiresAuth: true, role: 'user' },
    redirect: '/dashboard',
    children: [
    
      
     
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
        path: '/dashboard', 
        name: 'Dashboard', 
        component: UserDashboard 
      },
    ]
  },


  {
    path: '/dosen',
    component: DosenLayout,
    meta: { requiresAuth: true, role: 'lecturer' },
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
      },
      {
        path: 'reports',
        name: 'DosenExamReports',
        component: ExamReports
      }

    ]
  },


  {
    path: '/admin',
    component: AdminLayout,
    meta: { requiresAuth: true, role: 'admin' },
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
    meta: { requiresAuth: true, role: 'super_admin' },
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

router.beforeEach((to, from, next) => {
  const token = localStorage.getItem('token');
  let user = null;

  try {
    const userStr = localStorage.getItem('user');
    
    if (userStr && userStr !== 'undefined' && userStr !== 'null') {
      user = JSON.parse(userStr);
    }
  } catch (e) {
    user = null;
  }

 
  if (token && !user) {
    localStorage.clear();
    return next({ name: 'Login' });
  }

  
  if (to.matched.some(record => record.meta.requiresAuth)) {
    if (!token) {
      next({ name: 'Login' });
    } else {
   
      if (to.meta.role && user && to.meta.role !== user.role) {
        
        if (user.role === 'admin') next({ name: 'AdminDashboard' });
        else if (user.role === 'lecturer') next({ name: 'LecturerDashboard' });
        else if (user.role === 'super_admin') next({ name: 'SuperAdminDashboard' });
        else next({ name: 'Dashboard' }); 
      } else {
        next();
      }
    }
  } 

  else if (to.matched.some(record => record.meta.guest)) {
    if (token && user) {
      
      if (user.role === 'admin') next({ name: 'AdminDashboard' });
      else if (user.role === 'lecturer') next({ name: 'LecturerDashboard' });
      else if (user.role === 'super_admin') next({ name: 'SuperAdminDashboard' });
      else next({ name: 'Dashboard' });
    } else {
      next();
    }
  } 
  else {
    next();
  }
});

export default router

