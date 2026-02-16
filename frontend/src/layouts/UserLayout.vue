<template>
  <div class="font-sans bg-white text-dark-text">
    <Navbar />

    <main>
      <router-view></router-view>
    </main>

    </div>
</template>

<script setup>
import Navbar from '../components/ui/Navbar.vue';
</script>