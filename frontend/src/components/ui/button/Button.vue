<template>
  <component
    :is="componentTag"
    :to="to"
    :href="href"
    :class="buttonClasses"
    :disabled="disabled || loading"
    :type="isNativeButton ? type : undefined"
    @click="handleClick"
  >
    <div v-if="loading" class="mr-2 animate-spin rounded-full h-4 w-4 border-2 border-white border-t-transparent"></div>
    <span v-if="icon" class="mr-2">
       <i :class="icon"></i>
    </span>
    <slot />
  </component>
</template>

<script setup>
import { computed } from 'vue';

const props = defineProps({
  variant: {
    type: String,
    default: 'primary',
    validator: (value) => ['primary', 'secondary', 'danger', 'success', 'warning', 'outline', 'ghost'].includes(value),
  },
  size: {
    type: String,
    default: 'md',
    validator: (value) => ['sm', 'md', 'lg'].includes(value),
  },
  loading: {
    type: Boolean,
    default: false,
  },
  disabled: {
    type: Boolean,
    default: false,
  },
  block: {
    type: Boolean,
    default: false,
  },
  icon: {
    type: String,
    default: '',
  },
  to: {
    type: [String, Object],
    default: null,
  },
  href: {
    type: String,
    default: null,
  },
  type: {
    type: String,
    default: 'button',
  },
});

const emit = defineEmits(['click']);

const componentTag = computed(() => {
  if (props.to) return 'router-link';
  if (props.href) return 'a';
  return 'button';
});

const isNativeButton = computed(() => componentTag.value === 'button');

const handleClick = (e) => {
  if (props.disabled || props.loading) {
    e.preventDefault();
    return;
  }
  emit('click', e);
};

const buttonClasses = computed(() => {
  const baseClasses = 'inline-flex items-center justify-center font-medium rounded-lg transition-colors focus:outline-none focus:ring-2 focus:ring-offset-2 disabled:opacity-50 disabled:cursor-not-allowed';
  
  const variantClasses = {
    primary: 'bg-blue-600 hover:bg-blue-700 text-white focus:ring-blue-500 shadow-md',
    secondary: 'bg-gray-600 hover:bg-gray-700 text-white focus:ring-gray-500 shadow-sm',
    danger: 'bg-red-500 hover:bg-red-600 text-white focus:ring-red-500 shadow-sm',
    success: 'bg-green-600 hover:bg-green-700 text-white focus:ring-green-500 shadow-sm',
    warning: 'bg-yellow-400 hover:bg-yellow-500 text-white focus:ring-yellow-400 shadow-sm',
    outline: 'bg-white border border-gray-300 text-gray-700 hover:bg-gray-50 focus:ring-blue-500',
    ghost: 'bg-transparent hover:bg-gray-100 text-gray-700 focus:ring-gray-500',
  };

  const sizeClasses = {
    sm: 'px-3 py-1.5 text-sm',
    md: 'px-4 py-2 text-sm',
    lg: 'px-6 py-3 text-base',
  };

  const widthClass = props.block ? 'w-full' : '';

  return [
    baseClasses,
    variantClasses[props.variant] || variantClasses.primary,
    sizeClasses[props.size] || sizeClasses.md,
    widthClass,
  ].join(' ');
});
</script>
