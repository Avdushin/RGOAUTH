import { BrowserRouter as Router, Routes, Route, NavLink } from 'react-router-dom';
import './App.css';
import Err404 from './Error404';
import RegistrationForm from './Signup';
import LoginForm from './Login';
import Header from './Header';
import Main from './Main';
import Footer from './Footer';
import Docs from './Pages/Docs';

function App() {
  return (
    <Router>
      <Header />
      <div className="inner">
      <Routes>
        <Route path='/' element={<Main/>} />
        <Route path='/register' element={<RegistrationForm />} />
        <Route path='/login' element={<LoginForm />} />
        <Route path='/*' element={<Err404 />} />
        <Route path='/docs' element={<Docs />} />
      </Routes>
      </div>
    </Router>
  );
}

export default App;
