<template>
  <nav class="bg-white shadow-sm border-b border-gray-200 px-6 py-3 flex justify-between items-center">
    <!-- Logo -->
    <RouterLink to="/">
      <img 
        :src="Logo" 
        alt="Latih.in Logo"
        class="w-28 object-contain"
      />
    </RouterLink>

    <ul class="flex items-center space-x-6">
      <template v-if="user">
        <li v-if="user.role === 'user'">
          <RouterLink
            to="/ujian"
            :class="linkClass('/ujian')"
          >
            Ujian
          </RouterLink>
        </li>

        <li class="flex flex-col items-center space-y-1">
          <RouterLink
            :to="`/profile/${user.id}`"
            :class="linkClass('/profile')"
          >
            <User class="w-6 h-6" />
          </RouterLink>
          <p class="text-sm font-medium text-gray-700">
            {{ user.name }}
          </p>
        </li>
      </template>

      <template v-else>
        <li>
          <RouterLink
            to="/login"
            :class="linkClass('/login')"
          >
            Login
          </RouterLink>
        </li>
        <li>
          <RouterLink
            to="/register"
            :class="linkClass('/register')"
          >
            Register
          </RouterLink>
        </li>
      </template>
    </ul>
  </nav>
</template>

<script setup>
import { useRoute, RouterLink } from 'vue-router'
import { useGetCurrentUser } from '../../hooks/useGetCurrentUser'
import { User } from 'lucide-vue-next'
import Logo from '../../assets/logo.png'

const { user } = useGetCurrentUser()
const route = useRoute()
console.log()

const linkClass = (path) => {
  const isActive =
    route.path === path || route.path.startsWith(`${path}/`)

  return [
    'transition-colors duration-200 font-medium',
    isActive
      ? 'text-blue-600 border-b-2 border-blue-600 pb-1'
      : 'text-gray-600 hover:text-blue-600'
  ].join(' ')
}
</script>
