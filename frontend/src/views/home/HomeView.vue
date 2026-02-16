<template>
  <div class="home-page">

    <!-- ═══════════════════════ HERO SECTION ═══════════════════════ -->
    <section class="hero-section">
      <!-- Animated Background Elements -->
      <div class="hero-bg">
        <div class="hero-orb hero-orb--1"></div>
        <div class="hero-orb hero-orb--2"></div>
        <div class="hero-orb hero-orb--3"></div>
        <div class="hero-grid"></div>
      </div>

      <div class="hero-content reveal">
        <span class="hero-badge">
          <i class="fas fa-graduation-cap"></i>
          #DigitalisasiInformatika
        </span>
        <h1 class="hero-title">
          Seleksi Mahasiswa Pindahan
          <span class="hero-title-gradient">Lebih Cerdas & Efisien</span>
        </h1>
        <p class="hero-subtitle">
          Platform terintegrasi untuk manajemen bank soal dan ujian online
          Prodi Informatika. Objektif, transparan, dan terorganisir.
        </p>

        <div class="hero-actions">
          <router-link to="/login" class="btn-primary">
            <i class="fas fa-rocket"></i>
            Mulai Sekarang
          </router-link>
          <a href="#about" class="btn-secondary">
            <i class="fas fa-arrow-down"></i>
            Pelajari Lebih Lanjut
          </a>
        </div>

        <!-- Stats Row -->
        <div class="hero-stats">
          <div class="stat-item">
            <div class="stat-value">
              <i class="fas fa-book-open stat-icon"></i>
              500+
            </div>
            <div class="stat-label">Bank Soal</div>
          </div>
          <div class="stat-divider"></div>
          <div class="stat-item">
            <div class="stat-value">
              <i class="fas fa-users stat-icon"></i>
              200+
            </div>
            <div class="stat-label">Mahasiswa</div>
          </div>
          <div class="stat-divider"></div>
          <div class="stat-item">
            <div class="stat-value">
              <i class="fas fa-clipboard-check stat-icon"></i>
              50+
            </div>
            <div class="stat-label">Ujian Online</div>
          </div>
        </div>
      </div>
    </section>

    <!-- ═══════════════════════ ABOUT SECTION ═══════════════════════ -->
    <section id="about" class="about-section">
      <div class="section-container">
        <div class="about-grid">
          <div class="about-image reveal fade-left">
            <div class="image-card">
              <div class="image-glow"></div>
              <img
                :src="illustration1"
                alt="Ilustrasi Edukasi"
                loading="lazy"
              />
              <div class="image-badge">
                <i class="fas fa-shield-alt"></i>
                Terpercaya
              </div>
            </div>
          </div>

          <div class="about-content reveal fade-right">
            <span class="section-tag">Tentang Platform</span>
            <h2 class="section-title">
              Apa itu <span class="text-gradient">Latih.in</span>?
            </h2>
            <p class="section-desc">
              Latih.in hadir sebagai solusi digitalisasi seleksi akademik
              untuk Prodi Informatika. Platform ini mengintegrasikan
              <strong>manajemen bank soal</strong> dan
              <strong>tes online</strong> dalam satu sistem terpadu,
              memastikan proses seleksi mahasiswa pindahan berjalan efisien.
            </p>

            <div class="feature-list">
              <div class="feature-item">
                <div class="feature-icon feature-icon--blue">
                  <i class="fas fa-link"></i>
                </div>
                <div>
                  <h4>Terintegrasi Penuh</h4>
                  <p>Bank soal, ujian, dan penilaian dalam satu platform.</p>
                </div>
              </div>
              <div class="feature-item">
                <div class="feature-icon feature-icon--emerald">
                  <i class="fas fa-balance-scale"></i>
                </div>
                <div>
                  <h4>Objektivitas Tinggi</h4>
                  <p>Penilaian otomatis yang adil dan transparan.</p>
                </div>
              </div>
              <div class="feature-item">
                <div class="feature-icon feature-icon--violet">
                  <i class="fas fa-lock"></i>
                </div>
                <div>
                  <h4>Aman & Terjaga</h4>
                  <p>Keamanan data dan integritas ujian terjamin.</p>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- ═══════════════════════ FEATURES SECTION ═══════════════════════ -->
    <section class="features-section">
      <div class="section-container">
        <div class="section-header reveal">
          <span class="section-tag">Keunggulan</span>
          <h2 class="section-title section-title--center">
            Kenapa Memilih <span class="text-gradient">Latih.in</span>?
          </h2>
          <p class="section-desc section-desc--center">
            Dirancang dengan teknologi modern untuk pengalaman ujian digital yang premium.
          </p>
        </div>

        <div class="features-grid">
          <div class="feature-card reveal" v-for="(feature, idx) in features" :key="idx" :style="{ transitionDelay: `${idx * 0.1}s` }">
            <div class="feature-card-icon" :class="feature.colorClass">
              <i :class="feature.icon"></i>
            </div>
            <h3 class="feature-card-title">{{ feature.title }}</h3>
            <p class="feature-card-desc">{{ feature.desc }}</p>
          </div>
        </div>
      </div>
    </section>

    <!-- ═══════════════════════ HOW IT WORKS ═══════════════════════ -->
    <section class="steps-section">
      <div class="section-container">
        <div class="about-grid about-grid--reverse">
          <div class="about-content reveal fade-left">
            <span class="section-tag">Cara Kerja</span>
            <h2 class="section-title">
              Mudah & <span class="text-gradient">Terstruktur</span>
            </h2>
            <p class="section-desc">
              Persiapkan diri dengan materi uji terstruktur. Penilaian otomatis dan
              feedback instan tanpa menunggu proses manual.
            </p>

            <div class="steps-list">
              <div class="step-item" v-for="(step, idx) in steps" :key="idx">
                <div class="step-number">{{ idx + 1 }}</div>
                <div>
                  <h4>{{ step.title }}</h4>
                  <p>{{ step.desc }}</p>
                </div>
              </div>
            </div>
          </div>

          <div class="about-image reveal fade-right">
            <div class="image-card">
              <div class="image-glow image-glow--cyan"></div>
              <img
                :src="illustration2"
                alt="Ilustrasi Belajar Online"
                loading="lazy"
              />
            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- ═══════════════════════ CTA SECTION ═══════════════════════ -->
    <section class="cta-section reveal">
      <div class="cta-bg">
        <div class="cta-orb cta-orb--1"></div>
        <div class="cta-orb cta-orb--2"></div>
      </div>
      <div class="cta-content">
        <h2 class="cta-title">Siap Menghadapi Seleksi?</h2>
        <p class="cta-desc">
          Bergabunglah sekarang dan rasakan pengalaman ujian yang modern, cepat, dan terpercaya.
        </p>
        <router-link to="/login" class="btn-cta">
          <i class="fas fa-sign-in-alt"></i>
          Masuk ke Portal
        </router-link>
      </div>
    </section>

    <!-- ═══════════════════════ FOOTER ═══════════════════════ -->
    <footer class="site-footer">
      <div class="footer-container">
        <div class="footer-brand">
          <div class="footer-logo">Latih.in</div>
          <p class="footer-tagline">Mari belajar dan berkembang bersama.</p>
        </div>
        <div class="footer-divider"></div>
        <div class="footer-bottom">
          <small>&copy; 2025 Latih.in — All Rights Reserved.</small>
          <small class="footer-dept">Developed for Informatics Engineering Department.</small>
        </div>
      </div>
    </footer>
  </div>
</template>

<script setup>
import { onMounted } from 'vue';

import illustration1 from '../../assets/illustration-1.jpg';
import illustration2 from '../../assets/illustration-2.jpg';

const features = [
  { icon: 'fas fa-bolt', title: 'Performa Tinggi', desc: 'Antarmuka responsif dan cepat berkat teknologi Vue.js dan backend Golang.', colorClass: 'fc-icon--amber' },
  { icon: 'fas fa-layer-group', title: 'Bank Soal Terstruktur', desc: 'Kelola soal berdasarkan mata kuliah, tingkat kesulitan, dan kategori.', colorClass: 'fc-icon--blue' },
  { icon: 'fas fa-chart-bar', title: 'Analitik Lengkap', desc: 'Dashboard dan laporan hasil ujian yang informatif dan real-time.', colorClass: 'fc-icon--emerald' },
  { icon: 'fas fa-clock', title: 'Penilaian Otomatis', desc: 'Hasil ujian langsung tersedia tanpa proses manual, efisien dan akurat.', colorClass: 'fc-icon--violet' },
  { icon: 'fas fa-mobile-alt', title: 'Mobile Friendly', desc: 'Akses ujian dari mana saja, kapan saja, di perangkat apapun.', colorClass: 'fc-icon--rose' },
  { icon: 'fas fa-user-shield', title: 'Multi-Role Access', desc: 'Admin, Dosen, dan Mahasiswa memiliki akses sesuai peran masing-masing.', colorClass: 'fc-icon--cyan' },
];

const steps = [
  { title: 'Daftar & Login', desc: 'Buat akun atau login untuk masuk ke portal ujian.' },
  { title: 'Pilih Ujian', desc: 'Pilih ujian yang tersedia dan mulai sesi ujian.' },
  { title: 'Kerjakan & Submit', desc: 'Jawab soal dalam durasi waktu yang ditentukan.' },
  { title: 'Lihat Hasil', desc: 'Dapatkan hasil dan feedback secara instan.' },
];

onMounted(() => {
  const observer = new IntersectionObserver((entries) => {
    entries.forEach((entry) => {
      if (entry.isIntersecting) {
        entry.target.classList.add('active');
      }
    });
  }, { threshold: 0.1, rootMargin: '0px 0px -40px 0px' });

  document.querySelectorAll('.reveal').forEach((el) => observer.observe(el));
});
</script>

<style scoped>
/* ═══════════════════════ BASE & TOKENS ═══════════════════════ */
.home-page {
  font-family: 'Inter', 'Segoe UI', sans-serif;
  color: #1e293b;
  overflow-x: hidden;
  background: #f8fafc;
}

.text-gradient {
  background: linear-gradient(135deg, #2563eb, #06b6d4);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
}

.section-container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 0 24px;
}

/* ═══════════════════════ REVEAL ANIMATIONS ═══════════════════════ */
.reveal {
  opacity: 0;
  transform: translateY(36px);
  transition: opacity 0.7s cubic-bezier(0.16, 1, 0.3, 1), transform 0.7s cubic-bezier(0.16, 1, 0.3, 1);
}
.reveal.active {
  opacity: 1;
  transform: translateY(0);
}
.fade-left {
  transform: translateX(-40px) translateY(0);
}
.reveal.active.fade-left {
  transform: translateX(0) translateY(0);
}
.fade-right {
  transform: translateX(40px) translateY(0);
}
.reveal.active.fade-right {
  transform: translateX(0) translateY(0);
}

/* ═══════════════════════ HERO ═══════════════════════ */
.hero-section {
  position: relative;
  min-height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  text-align: center;
  padding: 100px 24px 80px;
  background: linear-gradient(180deg, #ffffff 0%, #f0f7ff 100%);
  overflow: hidden;
}

.hero-bg {
  position: absolute;
  inset: 0;
  pointer-events: none;
}

.hero-orb {
  position: absolute;
  border-radius: 50%;
  filter: blur(80px);
  opacity: 0.4;
  animation: float 8s ease-in-out infinite;
}
.hero-orb--1 {
  width: 500px; height: 500px;
  top: -120px; right: -100px;
  background: radial-gradient(circle, #93c5fd, #3b82f6);
}
.hero-orb--2 {
  width: 350px; height: 350px;
  bottom: -60px; left: -80px;
  background: radial-gradient(circle, #67e8f9, #06b6d4);
  animation-delay: 3s;
}
.hero-orb--3 {
  width: 200px; height: 200px;
  top: 50%; left: 50%;
  transform: translate(-50%, -50%);
  background: radial-gradient(circle, #c4b5fd, #8b5cf6);
  animation-delay: 5s;
  opacity: 0.2;
}

.hero-grid {
  position: absolute;
  inset: 0;
  background-image:
    linear-gradient(rgba(148, 163, 184, 0.06) 1px, transparent 1px),
    linear-gradient(90deg, rgba(148, 163, 184, 0.06) 1px, transparent 1px);
  background-size: 60px 60px;
}

@keyframes float {
  0%, 100% { transform: translateY(0) scale(1); }
  50% { transform: translateY(-30px) scale(1.05); }
}

.hero-content {
  position: relative;
  z-index: 10;
  max-width: 800px;
}

.hero-badge {
  display: inline-flex;
  align-items: center;
  gap: 8px;
  padding: 6px 18px;
  border-radius: 999px;
  font-size: 0.85rem;
  font-weight: 600;
  color: #2563eb;
  background: rgba(37, 99, 235, 0.08);
  border: 1px solid rgba(37, 99, 235, 0.15);
  margin-bottom: 24px;
  backdrop-filter: blur(4px);
}

.hero-title {
  font-size: clamp(2.5rem, 6vw, 4rem);
  font-weight: 800;
  line-height: 1.15;
  color: #0f172a;
  margin-bottom: 20px;
}

.hero-title-gradient {
  display: block;
  background: linear-gradient(135deg, #2563eb, #06b6d4);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
}

.hero-subtitle {
  font-size: 1.15rem;
  color: #64748b;
  line-height: 1.75;
  max-width: 600px;
  margin: 0 auto 36px;
}

.hero-actions {
  display: flex;
  gap: 16px;
  justify-content: center;
  flex-wrap: wrap;
  margin-bottom: 48px;
}

.btn-primary {
  display: inline-flex;
  align-items: center;
  gap: 10px;
  padding: 14px 32px;
  border-radius: 14px;
  font-weight: 700;
  font-size: 1rem;
  color: #fff;
  background: linear-gradient(135deg, #2563eb, #1d4ed8);
  text-decoration: none;
  box-shadow: 0 8px 30px rgba(37, 99, 235, 0.3);
  transition: all 0.3s cubic-bezier(0.16, 1, 0.3, 1);
}
.btn-primary:hover {
  transform: translateY(-3px);
  box-shadow: 0 12px 40px rgba(37, 99, 235, 0.4);
}

.btn-secondary {
  display: inline-flex;
  align-items: center;
  gap: 10px;
  padding: 14px 32px;
  border-radius: 14px;
  font-weight: 700;
  font-size: 1rem;
  color: #475569;
  background: #fff;
  border: 1.5px solid #e2e8f0;
  text-decoration: none;
  transition: all 0.3s ease;
}
.btn-secondary:hover {
  background: #f8fafc;
  border-color: #cbd5e1;
  transform: translateY(-2px);
}

/* Stats */
.hero-stats {
  display: flex;
  justify-content: center;
  align-items: center;
  gap: 32px;
  flex-wrap: wrap;
  padding: 28px 40px;
  background: rgba(255,255,255,0.7);
  backdrop-filter: blur(16px);
  border-radius: 20px;
  border: 1px solid rgba(226,232,240,0.6);
  box-shadow: 0 4px 24px rgba(0,0,0,0.04);
}
.stat-item { text-align: center; }
.stat-value {
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 1.6rem;
  font-weight: 800;
  color: #0f172a;
  justify-content: center;
}
.stat-icon { font-size: 1.1rem; color: #2563eb; }
.stat-label { font-size: 0.8rem; color: #94a3b8; font-weight: 500; margin-top: 2px; }
.stat-divider { width: 1px; height: 40px; background: #e2e8f0; }

/* ═══════════════════════ ABOUT ═══════════════════════ */
.about-section {
  padding: 100px 0;
  background: #fff;
}

.about-grid {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 64px;
  align-items: center;
}
.about-grid--reverse {
  direction: rtl;
}
.about-grid--reverse > * {
  direction: ltr;
}

.image-card {
  position: relative;
  border-radius: 24px;
  overflow: hidden;
  box-shadow: 0 20px 60px rgba(0,0,0,0.08);
}
.image-card img {
  width: 100%;
  height: auto;
  display: block;
  transition: transform 0.6s cubic-bezier(0.16, 1, 0.3, 1);
}
.image-card:hover img {
  transform: scale(1.04);
}
.image-glow {
  position: absolute;
  inset: -20px;
  background: linear-gradient(135deg, rgba(37,99,235,0.15), rgba(37,99,235,0));
  border-radius: 32px;
  z-index: -1;
}
.image-glow--cyan {
  background: linear-gradient(135deg, rgba(6,182,212,0.15), rgba(6,182,212,0));
}
.image-badge {
  position: absolute;
  bottom: 20px;
  left: 20px;
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 10px 18px;
  background: rgba(255,255,255,0.92);
  backdrop-filter: blur(12px);
  border-radius: 12px;
  font-size: 0.85rem;
  font-weight: 700;
  color: #0f172a;
  box-shadow: 0 4px 12px rgba(0,0,0,0.08);
}
.image-badge i { color: #2563eb; }

.section-tag {
  display: inline-block;
  font-size: 0.8rem;
  font-weight: 700;
  text-transform: uppercase;
  letter-spacing: 2px;
  color: #2563eb;
  margin-bottom: 12px;
}
.section-title {
  font-size: clamp(1.8rem, 4vw, 2.5rem);
  font-weight: 800;
  color: #0f172a;
  margin-bottom: 18px;
  line-height: 1.25;
}
.section-title--center { text-align: center; }
.section-desc {
  font-size: 1.05rem;
  color: #64748b;
  line-height: 1.8;
  margin-bottom: 32px;
}
.section-desc--center { text-align: center; max-width: 600px; margin-left: auto; margin-right: auto; }

.feature-list { display: flex; flex-direction: column; gap: 20px; }
.feature-item {
  display: flex;
  align-items: flex-start;
  gap: 16px;
  padding: 18px;
  background: #f8fafc;
  border-radius: 16px;
  border: 1px solid #f1f5f9;
  transition: all 0.3s ease;
}
.feature-item:hover {
  background: #fff;
  box-shadow: 0 8px 24px rgba(0,0,0,0.06);
  transform: translateX(4px);
}
.feature-item h4 { font-size: 0.95rem; font-weight: 700; color: #0f172a; margin-bottom: 4px; }
.feature-item p { font-size: 0.85rem; color: #64748b; line-height: 1.5; margin: 0; }
.feature-icon {
  width: 44px;
  height: 44px;
  min-width: 44px;
  border-radius: 12px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 1rem;
}
.feature-icon--blue { background: #eff6ff; color: #2563eb; }
.feature-icon--emerald { background: #ecfdf5; color: #059669; }
.feature-icon--violet { background: #f5f3ff; color: #7c3aed; }

/* ═══════════════════════ FEATURES GRID ═══════════════════════ */
.features-section {
  padding: 100px 0;
  background: linear-gradient(180deg, #f8fafc 0%, #fff 100%);
}
.section-header { margin-bottom: 56px; }
.features-grid {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 28px;
}
.feature-card {
  background: #fff;
  border: 1px solid #f1f5f9;
  border-radius: 20px;
  padding: 32px;
  transition: all 0.4s cubic-bezier(0.16, 1, 0.3, 1);
  cursor: default;
}
.feature-card:hover {
  transform: translateY(-8px);
  box-shadow: 0 20px 50px rgba(0,0,0,0.08);
  border-color: transparent;
}
.feature-card-icon {
  width: 56px;
  height: 56px;
  border-radius: 16px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 1.3rem;
  margin-bottom: 20px;
}
.fc-icon--amber { background: linear-gradient(135deg, #fef3c7, #fde68a); color: #d97706; }
.fc-icon--blue { background: linear-gradient(135deg, #dbeafe, #bfdbfe); color: #2563eb; }
.fc-icon--emerald { background: linear-gradient(135deg, #d1fae5, #a7f3d0); color: #059669; }
.fc-icon--violet { background: linear-gradient(135deg, #ede9fe, #ddd6fe); color: #7c3aed; }
.fc-icon--rose { background: linear-gradient(135deg, #ffe4e6, #fecdd3); color: #e11d48; }
.fc-icon--cyan { background: linear-gradient(135deg, #cffafe, #a5f3fc); color: #0891b2; }

.feature-card-title { font-size: 1.1rem; font-weight: 700; color: #0f172a; margin-bottom: 10px; }
.feature-card-desc { font-size: 0.9rem; color: #64748b; line-height: 1.65; margin: 0; }

/* ═══════════════════════ STEPS ═══════════════════════ */
.steps-section {
  padding: 100px 0;
  background: #fff;
}
.steps-list { display: flex; flex-direction: column; gap: 20px; }
.step-item {
  display: flex;
  align-items: flex-start;
  gap: 16px;
  padding: 18px;
  border-radius: 16px;
  background: #f8fafc;
  border: 1px solid #f1f5f9;
  transition: all 0.3s ease;
}
.step-item:hover {
  background: #fff;
  box-shadow: 0 6px 20px rgba(0,0,0,0.05);
}
.step-number {
  width: 40px;
  height: 40px;
  min-width: 40px;
  border-radius: 12px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-weight: 800;
  font-size: 1rem;
  color: #fff;
  background: linear-gradient(135deg, #2563eb, #1d4ed8);
}
.step-item h4 { font-size: 0.95rem; font-weight: 700; color: #0f172a; margin-bottom: 4px; }
.step-item p { font-size: 0.85rem; color: #64748b; margin: 0; line-height: 1.5; }

/* ═══════════════════════ CTA ═══════════════════════ */
.cta-section {
  position: relative;
  padding: 80px 24px;
  text-align: center;
  overflow: hidden;
  background: linear-gradient(135deg, #1e3a8a, #1e40af, #2563eb);
  margin: 0;
}
.cta-bg { position: absolute; inset: 0; pointer-events: none; }
.cta-orb {
  position: absolute;
  border-radius: 50%;
  filter: blur(100px);
  opacity: 0.25;
}
.cta-orb--1 {
  width: 400px; height: 400px;
  top: -100px; right: -50px;
  background: #06b6d4;
}
.cta-orb--2 {
  width: 300px; height: 300px;
  bottom: -80px; left: -60px;
  background: #8b5cf6;
}
.cta-content { position: relative; z-index: 10; max-width: 600px; margin: 0 auto; }
.cta-title { font-size: clamp(1.8rem, 4vw, 2.5rem); font-weight: 800; color: #fff; margin-bottom: 16px; }
.cta-desc { font-size: 1.1rem; color: rgba(255,255,255,0.75); margin-bottom: 36px; line-height: 1.7; }
.btn-cta {
  display: inline-flex;
  align-items: center;
  gap: 10px;
  padding: 16px 36px;
  border-radius: 14px;
  font-weight: 700;
  font-size: 1.05rem;
  color: #1e40af;
  background: #fff;
  text-decoration: none;
  box-shadow: 0 8px 30px rgba(0,0,0,0.15);
  transition: all 0.3s cubic-bezier(0.16, 1, 0.3, 1);
}
.btn-cta:hover {
  transform: translateY(-3px);
  box-shadow: 0 12px 40px rgba(0,0,0,0.2);
}

/* ═══════════════════════ FOOTER ═══════════════════════ */
.site-footer {
  background: #0f172a;
  padding: 56px 24px 40px;
  text-align: center;
}
.footer-container { max-width: 600px; margin: 0 auto; }
.footer-logo {
  font-size: 2.2rem;
  font-weight: 800;
  background: linear-gradient(135deg, #60a5fa, #22d3ee);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
  margin-bottom: 8px;
}
.footer-tagline { color: #94a3b8; font-size: 1rem; margin-bottom: 24px; }
.footer-divider { height: 1px; background: #1e293b; margin: 0 auto 24px; width: 80px; }
.footer-bottom small { display: block; color: #475569; font-size: 0.8rem; }
.footer-dept { color: #334155; margin-top: 4px; opacity: 0.6; }

/* ═══════════════════════ RESPONSIVE ═══════════════════════ */
@media (max-width: 768px) {
  .about-grid,
  .about-grid--reverse {
    grid-template-columns: 1fr;
    gap: 40px;
  }
  .about-grid--reverse {
    direction: ltr;
  }
  .features-grid {
    grid-template-columns: 1fr;
    gap: 20px;
  }
  .hero-stats {
    flex-direction: column;
    gap: 16px;
    padding: 24px;
  }
  .stat-divider {
    width: 60px;
    height: 1px;
  }
  .hero-section {
    padding: 80px 20px 60px;
    min-height: auto;
  }
}

@media (min-width: 769px) and (max-width: 1024px) {
  .features-grid {
    grid-template-columns: repeat(2, 1fr);
  }
}
</style>