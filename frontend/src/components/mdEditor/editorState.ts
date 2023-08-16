import { create } from 'zustand';

interface EditorState {
  isOpen: boolean;
  message: string;
  open: (message: string) => void;
  close: () => void;
}

const useEditorStore = create<EditorState>((set) => ({
  isOpen: false,
  message: '',
  open: (message: string) => {
    set({ isOpen: true, message });
  },
  close: () => {
    set({ isOpen: false });
  },
}));

export default useEditorStore;