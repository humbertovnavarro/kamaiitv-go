// import "../styles/globals.css"
import * as React from 'react';
import { ThemeProvider as NextThemesProvider } from 'next-themes';
import { createTheme, NextUIProvider } from "@nextui-org/react"
const darkTheme = createTheme({
  type: "dark",
  theme: {
    colors: {
      background: "#1a1a1a",
    }
  }
});
const lightTheme = createTheme({
  type: "light"
});

function App({ Component }) {
  return (
  <NextThemesProvider
      defaultTheme="system"
      attribute="class"
      value={{
        light: lightTheme.className,
        dark: darkTheme.className
      }}
    >
    <NextUIProvider>
        <Component />
    </NextUIProvider>
    </NextThemesProvider>
  );
}
export default App;
