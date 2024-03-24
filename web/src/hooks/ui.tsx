import { useContext } from 'react';

import { UiContext } from '../contexts/ui.tsx';

function useUi() {
  const context = useContext(UiContext);
  if (!context) {
    throw new Error(
      'useThemeContext must be used within a ThemeContextProvider'
    );
  }

  return {
    theme: context.theme,
    setTheme: context.setTheme,
  };
}

export default useUi;
