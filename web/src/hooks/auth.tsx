import { useContext } from 'react';

import { AuthContext } from '../contexts/auth.tsx';

function useAuth() {
  const context = useContext(AuthContext);
  if (!context) {
    throw new Error('useAuth must be used within an AuthContextProvider');
  }

  const logout = () => {
    context.setAuthenticated(null);
    localStorage.removeItem('token');
  };

  const login = (email: string, password: string) => {
    console.log({ email, password });
    // ...
  };

  return {
    user: context.user,
    logout,
    login,
  };
}

export default useAuth;
