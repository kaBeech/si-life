import './App.css'
import Header from "./components/Header"
import SiFloorForm from './components/SiFloorForm'
import SiFloorList from './components/SiFloorList'

export const BASE_URL = import.meta.env.MODE === 'dev' ? 'http://localhost:4000/api' : '/api';
function App() {

  return (
    <>
      <Header />
      <h1>Hello, world!</h1>
      <SiFloorForm />
      <SiFloorList />
    </>
  )
}

export default App
