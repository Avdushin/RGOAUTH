import { BrowserRouter as Router, Routes, Route, NavLink } from 'react-router-dom';
import './App.css';
import Err404 from './Error404';
import RegistrationForm from './Signup';
import LoginForm from './Login';
import Header from './Header';
import Main from './Main';
import Footer from './Footer';

function App() {
  return (
    <Router>
      <Header />
      <Routes>
        <Route path='/' element={<Main/>} />
        <Route path='/register' element={<RegistrationForm />} />
        <Route path='/login' element={<LoginForm />} />
        <Route path='/*' element={<Err404 />} />
      </Routes>
      <Footer />
    </Router>
  );
}

export default App;
