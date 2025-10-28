import { createRouter, createWebHistory } from "vue-router";

// Layouts
import UserLayout from '../layouts/UserLayout.vue'
import DosenLayout from '../layouts/DosenLayout.vue'

// Halaman Autentikasi
import LoginView from '../views/Auth/LoginView.vue'
import RegisterView from '../views/Auth/RegisterView.vue'
import ForgotPassword from '../views/Auth/ForgotPassword.vue'

// Halaman Pengguna
import HomePage from '../views/home/HomeView.vue' 
import ExamView from '../views/exam/ExamView.vue'
import Profile from '../views/profile/Profile.vue'
import UpdateProfile from '../views/profile/UpdateProfile.vue'

// Halaman Dosen
import LecturerDashboard from '../views/dosen/LecturerDashboard.vue'
import LecturerProfilePage from '../views/dosen/LecturerProfilePage.vue'
import LecturerSoal from '../views/question/LecturerSoal.vue'
import LecturerSoalList from '../views/question/LecturerSoalList.vue'
import CreateSoal from '../views/question/CreateSoal.vue'
import CreateManualSoal from '../views/question/CreateManualSoal.vue'
import UploadJsonSoal from '../views/question/UploadJsonSoal.vue'

const routes = [
  // Rute Autentikasi
  { path: '/login', name: 'login', component: LoginView },
  { path: '/register', name: 'register', component: RegisterView },
  { path: '/forgot-password', name: 'ForgotPassword', component: ForgotPassword },

  // Rute Pengguna Biasa
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

  // Rute Dosen
  {
    path: "/dosen",
    component: DosenLayout,
    redirect: "/dosen/dashboard",
    children: [
      { path: 'dashboard', name: 'LecturerDashboard', component: LecturerDashboard },
      { path: 'soal', name: 'LecturerSoal', component: LecturerSoal },
      { path: 'soal/list', name: 'LecturerSoalList', component: LecturerSoalList },
      { path: 'soal/create', name: 'CreateSoal', component: CreateSoal },
      { path: 'soal/create-manual', name: 'CreateManualSoal', component: CreateManualSoal },
      
      // ## RUTE BARU UNTUK EDIT SOAL DITAMBAHKAN DI SINI ##
      { 
        path: 'soal/edit/:id',
        name: 'EditSoal', 
        component: CreateManualSoal
      },

      { path: 'soal/upload-json', name: 'UploadJsonSoal', component: UploadJsonSoal },
      { path: 'update-profile', name: 'LecturerProfilePage', component: LecturerProfilePage },
    ]
  }
]

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes,
});

export default router;
