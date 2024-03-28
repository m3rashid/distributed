import { createContext, PropsWithChildren, useState } from 'react';

export type User = {
  id: number;
  name: string;
  email: string;
};

export type AuthContextType = {
  isAuthenticated: boolean;
  user: User | null;
  setAuthenticated: (user: User | null) => void;
};

export const AuthContext = createContext<AuthContextType>({
  isAuthenticated: false,
  user: null,
  setAuthenticated: () => {},
});

function AuthContextProvider({ children }: PropsWithChildren) {
  const [user, setUser] = useState<User | null>(null);

  return (
    <AuthContext.Provider
      value={{
        user,
        isAuthenticated: !!user,
        setAuthenticated: setUser,
      }}
    >
      {children}
    </AuthContext.Provider>
  );
}

export default AuthContextProvider;
