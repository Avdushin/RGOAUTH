import React from 'react';
import { NavLink } from 'react-router-dom';
import { GITHUB, VERSION } from '../vars';
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
          <h1>RGOAUTH <span className="version">{VERSION}</span></h1>
        </NavLink>
        <div className="logo_links">
          <NavLink className='logo_links-item' to='/docs'>Documentation</NavLink>
          <NavLink className='logo_links-item' to={GITHUB} target='_blank'>GitHub</NavLink>
        </div>
      </div>
      <div className="user">
      </div>
      <nav className="menu">
        {!username && (
          <div className="log-reg">
            <NavLink to='/login' className="menu__item">Вход</NavLink>/
            <NavLink to='/register' className="menu__item">Регистрация</NavLink>
          </div>
        )}

        {username && (
          <div className="menu__item helo-zone">
            {username && <div className="hello"><h3>Привет, {username}!</h3></div>}
            <div className="logout-btn">
              <a href="#" className='logout' onClick={logout}>
                <h3 className='logout'>Выход</h3>
              </a>
            </div>
          </div>
        )}
      </nav>
    </div>
  )
}

export default Header;
