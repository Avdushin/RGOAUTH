import React, { useState, useEffect } from "react";
import { NavLink, Link } from 'react-router-dom';
import axios from "axios";
import { GITHUB, VERSION } from '../vars';
import './header.css';


function Header() {
  const username = localStorage.getItem('username');
  // Profile
  const id = localStorage.getItem("id");
  const [user, setUser] = useState();
  const [message, setMessage] = useState('');
  // const  = localStorage.getItem("");

  // logout
  function logout() {
    localStorage.removeItem('registrant');
    localStorage.removeItem('username');
    localStorage.removeItem('email');
    localStorage.removeItem('password');
    localStorage.removeItem('name');
    localStorage.removeItem('surname');
    localStorage.removeItem('patronymic');
    localStorage.removeItem('age');
    localStorage.removeItem('id');
    window.location.href = '/login';
  }


  useEffect(() => {
    const fetchData = async () => {
      const response = await fetch(`http://localhost:8080/api/users/${id}`);
      const data = await response.json();
      setUser(data);
    };
    fetchData();
  }, []);

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
            {/* {username && <div className="hello"><h3>Привет, {username}!</h3></div>} */}
            {username && <div className="hello"><h3>Привет,  
              <Link to={`/api/users/${id}`}> {username}</Link>!</h3></div>}
            {/* <div className="site-layout-content">
              <UserForm user={user} onSubmit={handleSubmit} />
            </div> */}
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
