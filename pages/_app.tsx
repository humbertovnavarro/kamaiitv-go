import "../styles/globals.css"
import * as React from 'react';
import { createTheme, NextUIProvider } from "@nextui-org/react"
import useDarkMode from 'use-dark-mode';
const lightTheme = createTheme({
  type: 'light',
})

const darkTheme = createTheme({
  type: 'dark',
})

function App({ Component }) {
const darkMode = useDarkMode(true);
const [loginPrompt, setLoginPrompt] = React.useState(false);
  return (
    <NextUIProvider theme={darkMode.value ? darkTheme : lightTheme}>
        <Component />
    </NextUIProvider>
  );
}
export default App;
