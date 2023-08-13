import create from 'zustand';

interface ConfirmDialogState {
  isOpen: boolean;
  message: string;
  onConfirm: () => void;
  onCancel: () => void;
  open: (message: string, onConfirm: () => void, onCancel: () => void) => void;
  close: () => void;
}

const useConfirmDialogStore = create<ConfirmDialogState>((set) => ({
  isOpen: false,
  message: '',
  onConfirm: () => {},
  onCancel: () => {},
  open: (message: string, onConfirm: () => void, onCancel: () => void) => {
    set({ isOpen: true, message, onConfirm, onCancel });
  },
  close: () => {
    set({ isOpen: false });
  },
}));

export default useConfirmDialogStore;