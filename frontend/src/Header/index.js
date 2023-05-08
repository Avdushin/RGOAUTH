import React from 'react';
import { NavLink } from 'react-router-dom';
import './header.css';

function Header() {
  const username = localStorage.getItem('username');

  function logout() {
    localStorage.removeItem('registrant');
    localStorage.removeItem('username');
    localStorage.removeItem('email');
    localStorage.removeItem('password');
    window.location.href = '/login';
  }

  return (
    <div className="header">
      <div className="logo">
        <NavLink to={'/'} className='logo'>
          <h1>RGOAUTH</h1>
        </NavLink>
      </div>
      <div className="user">
      {username && <div className="hello"><h3>Hello, {username}!</h3></div>}
      </div>
      <nav className="menu">
        {!username && (
          <div className="log-reg">
            <NavLink to='/login' className="menu__item">Вход</NavLink>/
            <NavLink to='/register' className="menu__item">Регистрация</NavLink>
          </div>
        )}
        
        {username && (
          <div className="menu__item">
            <a href="#" className='logout' onClick={logout}>Выход</a>
          </div>
        )}
      </nav>
    </div>
  )
}

export default Header;
