import React, { useCallback, useMemo } from 'react';

type ModalContextProps = {
  open: (content: React.ReactNode) => void;
  close: () => void;
};

const ModalContext = React.createContext<ModalContextProps | undefined>(undefined);

type ModalProviderProps = {
  children: React.ReactNode;
};

export function ModalProvider({ children }: ModalProviderProps) {
  const [modal, setModal] = React.useState<React.ReactNode>();

  const open = useCallback((content: React.ReactNode) => {
    setModal(content);
  }, []);

  const close = useCallback(() => {
    setModal(undefined);
  }, []);

  const ctx = useMemo(() => ({ open, close }), [open, close]);

  return (
    <ModalContext.Provider value={ctx}>
      {children}
      {modal}
    </ModalContext.Provider>
  );
}

export function useModal() {
  const ctx = React.useContext(ModalContext);
  if (!ctx) {
    throw new Error('useModal must be used within a ModalProvider');
  }
  return ctx;
}
