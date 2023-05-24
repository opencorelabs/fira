import React from 'react';

type SessionContextProps = {
  //
};

export const SessionContext = React.createContext<SessionContextProps | null>(null);

type SessionProviderProps = {
  children: React.ReactNode;
};

export function SessionProvider({ children }: SessionProviderProps) {
  return <SessionContext.Provider value={null}>{children}</SessionContext.Provider>;
}
