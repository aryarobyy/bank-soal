
import { reactive } from 'vue';


const popupState = reactive({
  isOpen: false,
  variant: 'success', 
  title: '',
  message: '',
  detailMessage: '',
  leftButtonText: 'Batal',
  rightButtonText: 'OK',
  showLeftButton: false,
  showRightButton: true,
  

  onConfirm: null, 
  onCancel: null,
});

export function usePopup() {
  
  const open = (options) => {
    popupState.variant = options.variant || 'info';
    popupState.title = options.title || 'Info';
    popupState.message = options.message || '';
    popupState.detailMessage = options.detailMessage || '';
    popupState.leftButtonText = options.leftButtonText || 'Batal';
    popupState.rightButtonText = options.rightButtonText || 'Tutup';
    popupState.showLeftButton = options.showLeftButton ?? false;
    popupState.showRightButton = options.showRightButton ?? true;
    
    
    return new Promise((resolve, reject) => {
      popupState.onConfirm = () => {
        close();
        resolve(true); 
      };
      popupState.onCancel = () => {
        close();
        resolve(false); 
      };
      popupState.isOpen = true;
    });
  };

  const close = () => {
    popupState.isOpen = false;
   
    setTimeout(() => {
        popupState.onConfirm = null;
        popupState.onCancel = null;
    }, 300);
  };



  const showSuccess = (title, message) => {
    return open({
      variant: 'success',
      title: title || 'Berhasil',
      message: message,
      rightButtonText: 'OK',
      showLeftButton: false
    });
  };

  const showError = (title, message, detail = '') => {
    return open({
      variant: 'error',
      title: title || 'Terjadi Kesalahan',
      message: message,
      detailMessage: detail,
      rightButtonText: 'Tutup',
      showLeftButton: false
    });
  };

  
  const showConfirm = (title, message, confirmText = 'Ya, Lanjutkan') => {
    return open({
      variant: 'alert',
      title: title || 'Konfirmasi',
      message: message,
      leftButtonText: 'Batal',
      rightButtonText: confirmText,
      showLeftButton: true,
      showRightButton: true
    });
  };

  return {
    popupState, 
    showSuccess,
    showError,
    showConfirm,
    close
  };
}