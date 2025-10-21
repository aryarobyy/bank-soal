// imports
import { createRouter, createWebHistory } from "vue-router";

// Layouts
import DosenLayout from "../layouts/DosenLayout.vue";

// Halaman Autentikasi
import LoginView from "../views/Auth/LoginView.vue";
import RegisterView from "../views/Auth/RegisterView.vue";
import ForgotPassword from "../views/Auth/ForgotPassword.vue";

// Halaman Mahasiswa
import HomePage from "../views/Auth/HomePage.vue";
import ExamPage from "../views/Auth/ExamPage.vue";
import StudentProfilePage from "../views/StudentProfilePage.vue";

// Halaman Dosen
import LecturerDashboard from "../views/dosen/LecturerDashboard.vue";
import LecturerSoalList from "../views/dosen/LecturerSoalList.vue";
import CreateSoal from "../views/dosen/CreateSoal.vue";
import CreateManualSoal from "../views/dosen/CreateManualSoal.vue";
import UploadJsonSoal from "../views/dosen/UploadJsonSoal.vue";
import LecturerProfilePage from "../views/dosen/LecturerProfilePage.vue";
import LecturerSoal from "../views/dosen/LecturerSoal.vue";

const routes = [
  // --- Rute Umum & Mahasiswa ---
  { path: "/", name: "HomePage", component: HomePage },
  { path: "/login", name: "login", component: LoginView },

  // RUTE YANG HILANG DITAMBAHKAN KEMBALI DI SINI
  { path: "/register", name: "register", component: RegisterView },
  {
    path: "/forgot-password",
    name: "ForgotPassword",
    component: ForgotPassword,
  },
  { path: "/ujian", name: "ExamPage", component: ExamPage },

  {
    path: "/update-profile",
    name: "StudentProfilePage",
    component: StudentProfilePage,
  },

  // --- Rute Dosen (Menggunakan Layout) ---
  {
    path: "/dosen",
    component: DosenLayout,
    redirect: "/dosen/dashboard",
    children: [
      {
        path: "dashboard",
        name: "LecturerDashboard",
        component: LecturerDashboard,
      },
      { path: "soal", name: "LecturerSoal", component: LecturerSoal },
      {
        path: "soal/list",
        name: "LecturerSoalList",
        component: LecturerSoalList,
      },
      { path: "soal/create", name: "CreateSoal", component: CreateSoal },
      {
        path: "soal/create-manual",
        name: "CreateManualSoal",
        component: CreateManualSoal,
      },
      {
        path: "soal/upload-json",
        name: "UploadJsonSoal",
        component: UploadJsonSoal,
      },
      {
        path: "update-profile",
        name: "LecturerProfilePage",
        component: LecturerProfilePage,
      },
    ],
  },
];

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes,
});

export default router;
