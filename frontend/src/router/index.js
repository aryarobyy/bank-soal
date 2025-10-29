import { createRouter, createWebHistory } from 'vue-router'


import UserLayout from '../layouts/UserLayout.vue'
import DosenLayout from '../layouts/DosenLayout.vue'


import LoginView from '../views/Auth/LoginView.vue'
import RegisterView from '../views/Auth/RegisterView.vue'
import Profile from '../views/profile/Profile.vue'
import UpdateProfile from '../views/profile/UpdateProfile.vue'
import HomePage from '../views/home/HomeView.vue' 
import ExamView from '../views/exam/ExamView.vue'
import ForgotPassword from '../views/Auth/ForgotPassword.vue'
import LecturerDashboard from '../views/dosen/LecturerDashboard.vue'
import LecturerSoal from '../views/question/LecturerSoal.vue'
import LecturerSoalList from '../views/question/LecturerSoalList.vue'
import CreateSoal from '../views/question/CreateSoal.vue'
import CreateManualSoal from '../views/question/CreateManualSoal.vue'
import UploadJsonSoal from '../views/question/UploadJsonSoal.vue'

const routes = [
  
  { path: '/login', name: 'login', component: LoginView },
  { path: '/register', name: 'register', component: RegisterView },
  { path: '/forgot-password', name: 'ForgotPassword', component: ForgotPassword },

  
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

  
  {
    path: '/dosen',
    component: DosenLayout,
    redirect: '/dosen/dashboard',
    children: [
      { path: 'dashboard', name: 'LecturerDashboard', component: LecturerDashboard },
      { path: 'soal', name: 'LecturerSoal', component: LecturerSoal },
      { path: 'soal/list', name: 'LecturerSoalList', component: LecturerSoalList },
      { path: 'soal/create', name: 'CreateSoal', component: CreateSoal },
      { path: 'soal/create-manual', name: 'CreateManualSoal', component: CreateManualSoal },
      { path: 'soal/edit/:id', name: 'EditSoal', component: CreateManualSoal },
      { path: 'soal/upload-json', name: 'UploadJsonSoal', component: UploadJsonSoal },
      
      
      { 
        path: 'profile/:id',
        name: 'DosenProfile', 
        component: Profile
      },
      { 
        path: 'update-profile',
        name: 'DosenUpdateProfile',
        component: UpdateProfile
      },
    ]
  }
]

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes
})

export default router