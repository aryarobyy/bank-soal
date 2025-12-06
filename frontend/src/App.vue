<template>
  <router-view></router-view>

  <Popup
    :is-open="popupState.isOpen"
    :variant="popupState.variant"
    :title="popupState.title"
    :message="popupState.message"
    :detail-message="popupState.detailMessage"
    :left-button-text="popupState.leftButtonText"
    :right-button-text="popupState.rightButtonText"
    :show-left-button="popupState.showLeftButton"
    :show-right-button="popupState.showRightButton"
    @close="handleClose"
    @left-click="handleLeftClick"
    @right-click="handleRightClick"
  />
</template>

<script setup lang="ts">
import { onMounted } from 'vue';
import { provideUser } from './hooks/useGetCurrentUser';


import Popup from './components/ui/Popup.vue'; 
import { usePopup } from './hooks/usePopup';


const { fetchUser } = provideUser();

onMounted(async () => {
  await fetchUser();
});


const { popupState } = usePopup();


const handleClose = () => {
  if (popupState.onCancel) {
    popupState.onCancel(); 
  } else {
    popupState.isOpen = false; 
  }
};

const handleLeftClick = () => {
  if (popupState.onCancel) {
    popupState.onCancel();
  } else {
    popupState.isOpen = false;
  }
};


const handleRightClick = () => {
  if (popupState.onConfirm) {
    popupState.onConfirm(); 
  } else {
    popupState.isOpen = false;
  }
};
</script>