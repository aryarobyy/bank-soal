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

// Exam
import ManageExam from "../views/exam/ManageExam.vue";
import CreateExamView from "../views/exam/CreateExamView.vue";
import ExamReports from "../views/exam/ExamReports.vue";

// Super Admin
import SuperAdminDashboard from "../views/SuperAdmin/SuperAdminDashboard.vue";
import ManageAdmin from "../views/ManageAdmin/ManageAdmin.vue";

// Definisikan semua rute soal di satu tempat
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

const routes = [
  // AUTH & UMUM
  { path: "/login", name: "login", component: LoginView },
  { path: "/register", name: "register", component: RegisterView },
  {
    path: "/forgot-password",
    name: "ForgotPassword",
    component: ForgotPassword,
  },

  // USER ROUTES
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

  // DOSEN ROUTES
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

      // Soal routes
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
      { path: "soal/edit/:id", name: "EditSoal", component: CreateManualSoal },
      {
        path: "soal/upload-json",
        name: "UploadJsonSoal",
        component: UploadJsonSoal,
      },

      // ✅ Exam routes
      {
        path: "exam",
        name: "LecturerExamList",
        component: () => import("../views/exam/LecturerExamList.vue"),
      },
      {
        path: "exam/create",
        name: "CreateExam",
        component: () => import("../views/exam/LecturerExamCreate.vue"),
      },
      {
        path: "exam/detail/:id",
        name: "LecturerExamDetail",
        component: () => import("../views/exam/LecturerExamDetail.vue"),
      },
      {
        path: "exam/edit/:id",
        name: "LecturerExamEdit",
        component: () => import("../views/exam/LecturerExamEdit.vue"),
      },
      { path: "profile/:id", name: "DosenProfile", component: Profile },
      {
        path: "update-profile",
        name: "DosenUpdateProfile",
        component: UpdateProfile,
      },

      // ✅ Integrasi soalRoutes (pakai map)
      {
        path: "soal",
        component: RouterWrapper,
        children: soalRoutes.map((route) => ({
          ...route,
          name: `Dosen${route.name}`,
        })),
      },
    ],
  },

  // ADMIN ROUTES
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
      { path: "dosen", name: "AdminManageDosen", component: ManageDosen },
      { path: "ujian", name: "AdminManageExam", component: ManageExam },
      {
        path: "ujian/create",
        name: "AdminCreateExam",
        component: CreateExamView,
      },
      { path: "reports", name: "AdminExamReports", component: ExamReports },
      {
        path: "update-profile",
        name: "AdminUpdateProfile",
        component: UpdateProfile,
      },
    ],
  },

  // SUPER ADMIN
  {
    path: "/superadmin",
    component: SuperAdminLayout,
    redirect: "/superadmin/dashboard",
    children: [
      {
        path: "dashboard",
        name: "SuperAdminDashboard",
        component: SuperAdminDashboard,
      },
      {
        path: "admins",
        name: "SuperAdminManageAdmins",
        component: ManageAdmin,
      },
    ],
  },
];

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes,
});

export default router;
