import { PropsWithChildren, createContext, useState } from 'react';

export type Theme = 'light' | 'dark';
export type UiContextType = {
  theme: Theme;
  setTheme: React.Dispatch<React.SetStateAction<Theme>>;
};

export const UiContext = createContext<UiContextType>({
  theme: 'light',
  setTheme: () => {},
});

function UiContextProvider({ children }: PropsWithChildren) {
  const [theme, setTheme] = useState<Theme>('light');

  return (
    <UiContext.Provider value={{ theme, setTheme }}>
      {children}
    </UiContext.Provider>
  );
}

export default UiContextProvider;
