// src/router/index.js
import { createRouter, createWebHistory } from "vue-router";

// Layouts
import UserLayout from "../layouts/UserLayout.vue";
import DosenLayout from "../layouts/DosenLayout.vue";
import AdminLayout from "../layouts/AdminLayout.vue";
import RouterWrapper from "../layouts/RouterWrapper.vue";
import SuperAdminLayout from "../layouts/SuperAdminLayout.vue";

// Halaman Auth & Umum
import LoginView from "../views/Auth/LoginView.vue";
import RegisterView from "../views/Auth/RegisterView.vue";
import ForgotPassword from "../views/Auth/ForgotPassword.vue";
import HomePage from "../views/home/HomeView.vue";
import ExamView from "../views/exam/ExamView.vue";
import Profile from "../views/profile/Profile.vue";
import UpdateProfile from "../views/profile/UpdateProfile.vue";

// Dosen
import LecturerDashboard from "../views/dosen/LecturerDashboard.vue";

// Admin
import AdminDashboard from "../views/Admin/AdminDashboard.vue";
import ManageMahasiswa from "../views/ManageMahasiswa/ManageMahasiswa.vue";
import ManageDosen from "../views/ManageDosen/ManageDosen.vue";

// Soal (umum)
import LecturerSoal from "../views/question/LecturerSoal.vue";
import LecturerSoalList from "../views/question/LecturerSoalList.vue";
import CreateSoal from "../views/question/CreateSoal.vue";
import CreateManualSoal from "../views/question/CreateManualSoal.vue";
import UploadJsonSoal from "../views/question/UploadJsonSoal.vue";

// Halaman Ujian (yang akan dipakai bersama)
import ManageExam from "../views/exam/ManageExam.vue";
import CreateExamView from "../views/exam/CreateExamView.vue";
import ExamReports from "../views/exam/ExamReports.vue";
import ExamDetail from "../views/exam/ExamDetail.vue";
import ExamEdit from "../views/exam/ExamEdit.vue";

// Halaman Super Admin
import SuperAdminDashboard from "../views/SuperAdmin/SuperAdminDashboard.vue";
import ManageAdmin from "../views/ManageAdmin/ManageAdmin.vue";

// --- Grup route Soal ---
const soalRoutes = [
  { path: "", name: "SoalHome", component: LecturerSoal },
  { path: "list", name: "SoalList", component: LecturerSoalList },
  { path: "create", name: "SoalCreate", component: CreateSoal },
  {
    path: "create-manual",
    name: "SoalCreateManual",
    component: CreateManualSoal,
  },
  { path: "edit/:id", name: "SoalEdit", component: CreateManualSoal },
  { path: "upload-json", name: "SoalUploadJson", component: UploadJsonSoal },
];

// --- Grup route Exam (dipakai dosen & admin) ---
const examRoutes = [
  { path: "", name: "ManageExam", component: ManageExam },
  { path: "create", name: "CreateExam", component: CreateExamView },
  { path: "detail/:id", name: "ExamDetail", component: ExamDetail },
  { path: "edit/:id", name: "ExamEdit", component: ExamEdit },
];

const routes = [
  // AUTH
  { path: "/login", name: "login", component: LoginView },
  { path: "/register", name: "register", component: RegisterView },
  {
    path: "/forgot-password",
    name: "ForgotPassword",
    component: ForgotPassword,
  },

  // 👇 ROUTE HALAMAN MENGERJAKAN UJIAN (TANPA LAYOUT / SIDEBAR / NAVBAR)
  {
    path: "/exam/start",
    name: "ExamDo",
    component: () => import("../views/exam/ExamDo.vue"),
  },

  // USER ROUTES (pakai UserLayout + navbar user)
  {
    path: "/",
    component: UserLayout,
    children: [
      { path: "", name: "HomePage", component: HomePage },
      { path: "ujian", name: "ExamPage", component: ExamView },
      { path: "profile/:id", name: "Profile", component: Profile },
      {
        path: "update-profile",
        name: "UpdateProfile",
        component: UpdateProfile,
      },
    ],
  },

  // DOSEN ROUTES (pakai sidebar dosen)
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
      {
        path: "soal",
        component: RouterWrapper,
        children: soalRoutes.map((route) => ({
          ...route,
          name: `Dosen${route.name}`,
        })),
      },
      {
        path: "exam",
        component: RouterWrapper,
        children: examRoutes.map((route) => ({
          ...route,
          name: `Dosen${route.name}`,
        })),
      },
    ],
  },

  // ADMIN ROUTES (pakai sidebar admin)
  {
    path: "/admin",
    component: AdminLayout,
    redirect: "/admin/dashboard",
    children: [
      { path: "dashboard", name: "AdminDashboard", component: AdminDashboard },
      {
        path: "mahasiswa",
        name: "AdminManageMahasiswa",
        component: ManageMahasiswa,
      },
      {
        path: "dosen",
        name: "AdminManageDosen",
        component: ManageDosen,
      },
      {
        path: "ujian",
        component: RouterWrapper,
        children: examRoutes.map((route) => ({
          ...route,
          name: `Admin${route.name}`,
        })),
      },
      {
        path: "reports",
        name: "AdminExamReports",
        component: ExamReports,
      },
      {
        path: "update-profile",
        name: "AdminUpdateProfile",
        component: UpdateProfile,
      },
      {
        path: "excel-files",
        name: "AdminXlsFiles",
        component: () => import("../views/Admin/ManageXlsFiles.vue"),
        meta: { requiresAuth: true, role: "admin" },
      },
    ],
  },

  // SUPERADMIN
  {
    path: "/superadmin",
    component: SuperAdminLayout,
    redirect: "/superadmin/dashboard",
    children: [
      {
        path: "dashboard",
        name: "SuperAdminDashboard",
        component: SuperAdminDashboard,
        meta: { title: "Dashboard Aktivitas" },
      },
      {
        path: "admins",
        name: "SuperAdminManageAdmins",
        component: ManageAdmin,
        meta: { title: "Manajemen Admin" },
      },
    ],
  },
];

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes,
});

export default router;
