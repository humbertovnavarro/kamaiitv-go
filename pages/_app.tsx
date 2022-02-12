import "../styles/globals.css"
function App({ Component }) {
  return (
    <main className="antialised text-slate-500 dark:text-slate-400 bg-white dark:bg-slate-900 w-screen h-screen overflow-x-clip">
      <Component />
    </main>
  );
}
export default App;
