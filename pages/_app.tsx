import "../styles/globals.css"
import * as React from 'react';
import { NextUIProvider } from '@nextui-org/react';
function App({ Component }) {
  return (
    <NextUIProvider>
      <Component />
    </NextUIProvider>
  );
}
export default App;
