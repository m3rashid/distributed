import {
  webDarkTheme,
  webLightTheme,
  FluentProvider,
} from '@fluentui/react-components';
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
      <FluentProvider theme={theme === 'light' ? webLightTheme : webDarkTheme}>
        {children}
      </FluentProvider>
    </UiContext.Provider>
  );
}

export default UiContextProvider;
