<template>
  <div class="w-full p-6 sm:p-8 mx-auto bg-white rounded-2xl shadow-xl">
    
    <div class="flex flex-col sm:flex-row justify-between items-start sm:items-center mb-6 pb-4 border-b border-gray-200 gap-4">
      <div>
        <h2 class="text-2xl font-bold text-dark-text">{{ pageTitle }}</h2>
        <p v-if="!isAdminRoute" class="text-sm text-gray-500 mt-1">
          <i class="fas fa-user-check mr-1"></i> Menampilkan soal Anda
        </p>
      </div>

      <div class="flex gap-3">
        <router-link
          v-if="isFiltered"
          :to="{ name: isAdminRoute ? 'AdminSoalHome' : 'DosenSoalHome' }"
          class="flex items-center gap-2 px-4 py-2 font-semibold text-gray-700 transition-colors rounded-lg bg-gray-200 hover:bg-gray-300"
        >
          <i class="fas fa-arrow-left"></i> Kembali ke Grup Soal
        </router-link>

        <router-link 
          v-else
          :to="{ name: isAdminRoute ? 'AdminSoalCreate' : 'DosenSoalCreate' }"
          class="flex items-center gap-2 px-4 py-2 font-semibold text-white transition-colors rounded-lg bg-blue-600 hover:bg-blue-700"
        >
          <i class="fas fa-plus-circle"></i> Tambah Soal
        </router-link>
      </div>
    </div>

    <div class="overflow-x-auto">
      <div v-if="loading" class="text-center p-10">Memuat daftar soal...</div>
      
      <div v-else-if="soalList.length === 0" class="text-center p-12 bg-gray-50 rounded-lg border-2 border-dashed border-gray-300">
        <p class="text-lg font-medium text-gray-600">Tidak ada soal ditemukan</p>
        <p class="text-sm text-gray-400 mt-1">{{ isFiltered ? 'Filter ini kosong.' : 'Belum ada data.' }}</p>
      </div>

      <table v-else class="min-w-full text-left border-collapse table-auto">
        <thead>
          <tr class="text-xs font-bold tracking-wider text-gray-500 uppercase bg-gray-100 border-b border-gray-200">
            <th class="px-6 py-4 w-16 text-center">ID</th>
            <th class="px-6 py-4">Pertanyaan</th>
            <th class="px-6 py-4 text-center">Jawaban Benar</th>
            <th class="px-6 py-4 text-center">Difficulty</th>
            <th class="px-6 py-4 text-center">Aksi</th>
          </tr>
        </thead>
        <tbody class="text-sm text-gray-700">
          <tr v-for="soal in soalList" :key="soal.id" class="transition-colors border-b border-gray-100 hover:bg-blue-50">
            <td class="px-6 py-4 text-center font-medium text-gray-500">{{ soal.id }}</td>
            <td class="px-6 py-4">
              <div class="line-clamp-2" :title="soal.question_text">{{ soal.question_text }}</div>
              <div v-if="isAdminRoute" class="text-xs text-gray-400 mt-1">Creator: {{ soal.creator_id }}</div>
            </td>
            <td class="px-6 py-4 text-center font-semibold text-blue-600">
               {{ soal.options?.find(opt => opt.is_correct)?.option_text || '-' }}
            </td>
            <td class="px-6 py-4 text-center capitalize">
               <span :class="getDifficultyClass(soal.difficulty)">{{ soal.difficulty }}</span>
            </td>
            <td class="px-6 py-4 text-center">
              <div class="flex justify-center gap-2">
                <button @click="editSoal(soal.id)" class="text-green-600 hover:text-green-800"><i class="fas fa-edit"></i></button>
                <button @click="handleDeleteSoal(soal.id)" class="text-red-600 hover:text-red-800"><i class="fas fa-trash"></i></button>
              </div>
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <div v-if="!loading && totalPages > 1" class="flex items-center justify-between mt-6 pt-4 border-t border-gray-200">
      <div class="text-sm text-gray-500">
        Halaman <span class="font-bold text-gray-800">{{ currentPage }}</span> 
        dari {{ totalPages }}
      </div>
      <div class="flex gap-2">
        <button @click="prevPage" :disabled="currentPage === 1" class="px-4 py-2 border rounded-lg hover:bg-gray-50 disabled:opacity-50">Prev</button>
        <button @click="nextPage" :disabled="currentPage >= totalPages" class="px-4 py-2 border rounded-lg hover:bg-gray-50 disabled:opacity-50">Next</button>
      </div>
    </div>

  </div>
</template>

<script>
import { 
  getmanyQuestions, 
  getQuestionsBySubject, 
  deleteQuestion,
  getQuestionsByCreator,
  getQuestionsByCreatorAndSubject
} from "../../provider/question.provider";

import { usePopup } from "../../hooks/usePopup";
import { useGetCurrentUser } from "../../hooks/useGetCurrentUser"; 

export default {
  name: 'SoalList',
  data() {
    return {
      soalList: [],
      loading: true,
      subjectId: null, 
      creatorId: null,
      currentPage: 1, 
      limit: 10,
      totalSoal: 0,
    };
  },
  setup() {
    const { showSuccess, showError, showConfirm } = usePopup();
    const { user } = useGetCurrentUser();
    return { showSuccess, showError, showConfirm, user };
  },
  computed: {
    isAdminRoute() { return this.$route.path.startsWith('/admin'); },
    isFiltered() { return !!this.$route.query.subject_id || !!this.$route.query.creator_id; },
    pageTitle() {
      if (this.$route.query.subject_id) return 'Daftar Soal per Subjek';
      if (this.$route.query.creator_id) return 'Soal yang Dibuat Pengguna';
      return 'Daftar Semua Soal'; 
    },
    totalPages() {
      const pages = Math.ceil(this.totalSoal / this.limit);
      return pages > 0 ? pages : 1;
    }
  },
  methods: {

    async handleRouteChange(isFirstLoad = false) {
      const query = this.$route.query;
      
      if (query.page) {
        this.loading = true;
        this.currentPage = parseInt(query.page, 10) || 1;
        const newQuery = { ...query };
        delete newQuery.page;
        this.$router.replace({ query: newQuery });
        this.fetchSoalList(); 
        
      } else if (query.show_last_page === 'true') {
        this.loading = true;
        this.$router.replace({ query: { ...query, show_last_page: undefined } });
        
      
        try {
          const sid = query.subject_id;
          const cid = query.creator_id || (this.isAdminRoute ? null : this.user.id);
          
          let response;
        
          if (sid) response = await getQuestionsBySubject(sid, 1, 0); 
          else if (this.isAdminRoute) response = await getmanyQuestions(1, 0); 
          else response = await getQuestionsByCreator(cid, 1, 0); 

     
          const total = response.total || 0;
          this.currentPage = Math.ceil(total / this.limit) || 1;
          
          this.fetchSoalList();
        } catch (e) {
          this.currentPage = 1;
          this.fetchSoalList();
        }

      } else {
        if (!isFirstLoad) this.currentPage = 1;
        this.fetchSoalList();
      }
    },

 
    async fetchSoalList() {
     
      if (!this.isAdminRoute && !this.user?.id) return;

      this.loading = true;
      this.subjectId = this.$route.query.subject_id;
      this.creatorId = this.$route.query.creator_id;

      try {
        let response;
        const offset = (this.currentPage - 1) * this.limit; 

        if (this.subjectId) {
          if (this.isAdminRoute) {
         
             response = await getQuestionsBySubject(this.subjectId, this.limit, offset);
          } else {
         
             response = await getQuestionsByCreatorAndSubject(this.user.id, this.subjectId, this.limit, offset);
          }
        } 
        

        else if (this.creatorId || (!this.isAdminRoute && !this.subjectId)) {

          const cid = this.creatorId || this.user.id;
          response = await getQuestionsByCreator(cid, this.limit, offset);
        } 
        
  
        else {
          response = await getmanyQuestions(this.limit, offset);
        }

   
        if (response.items) {
            this.soalList = response.items;
            this.totalSoal = response.total || 0;
        } else if (response.data) {
    
            if (Array.isArray(response.data)) {
               this.soalList = response.data;
               this.totalSoal = response.total || response.data.length;
            } else {
              
               this.soalList = response.data;
               this.totalSoal = response.total || 0;
            }
        } else if (Array.isArray(response)) {
            this.soalList = response;
            this.totalSoal = response.length;
        } else {
            this.soalList = [];
            this.totalSoal = 0;
        }

      } catch (error) {
        console.error("Gagal fetch soal:", error);
        this.soalList = [];
      } finally {
        this.loading = false;
      }
    },
    
    nextPage() {
      if (this.currentPage < this.totalPages) {
        this.currentPage++;
        this.fetchSoalList();
      }
    },
    prevPage() {
      if (this.currentPage > 1) {
        this.currentPage--;
        this.fetchSoalList();
      }
    },
    editSoal(id) {
      const routeName = this.isAdminRoute ? 'AdminSoalEdit' : 'DosenSoalEdit';
      const query = {};
      if (this.subjectId) query.return_subject_id = this.subjectId;
      query.return_page = this.currentPage;
      this.$router.push({ name: routeName, params: { id: id }, query: query });
    },
    async handleDeleteSoal(id) {
      if (await this.showConfirm('Hapus?', `Yakin hapus ID: ${id}?`)) {
        try {
          await deleteQuestion(id);
          this.showSuccess('Berhasil', 'Soal dihapus.');
          if (this.soalList.length === 1 && this.currentPage > 1) this.currentPage--;
          this.fetchSoalList();
        } catch (e) {
          const msg = e.response?.data?.message || "Gagal hapus";
          this.showError("Gagal", msg.includes("constraint") ? "Soal dipakai di ujian" : msg);
        }
      }
    },
    getDifficultyClass(level) {
      if(level === 'easy') return 'text-green-600 bg-green-50 px-2 py-1 rounded font-medium';
      if(level === 'medium') return 'text-yellow-600 bg-yellow-50 px-2 py-1 rounded font-medium';
      return 'text-red-600 bg-red-50 px-2 py-1 rounded font-medium';
    }
  },
  
  created() {
    this.handleRouteChange(true);
  },
  
  watch: {
    user(newVal) { if(newVal?.id) this.handleRouteChange(false); },
    '$route.query'(newQuery, oldQuery) {
      const isParamsChange = (oldQuery.show_last_page && !newQuery.show_last_page) || 
                             (oldQuery.page && !newQuery.page) ||
                             (oldQuery.subject_id !== newQuery.subject_id);
      if (!isParamsChange) {
        this.handleRouteChange(false); 
      }
    }
  }
};
</script>