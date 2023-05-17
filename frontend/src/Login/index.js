import React, { useState } from 'react';
import { NavLink } from 'react-router-dom';
import axios from 'axios';
import '../assets/css/forms.css'

function LoginForm() {
  const [id, setId] = useState('');
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [message, setMessage] = useState('');
  const [error, setError] = useState('');

  // Heandlers
  const handleEmailChange = (event) => {
    setEmail(event.target.value);
  }

  const handlePasswordChange = (event) => {
    setPassword(event.target.value);
  }

  const handleSubmit = (event) => {
    event.preventDefault();
    axios.post(`${process.env.REACT_APP_SERVER_URL}/login`, { email, password })
      .then(response => {
        setMessage(response.data.message);
        console.log(response.data)
        localStorage.setItem('id', response.data.id);
        localStorage.setItem('username', response.data.username);
        localStorage.setItem('email', response.data.email);
        localStorage.setItem('password', response.data.password);
        setTimeout(() => {
          window.location.href = "/";
        }, 1000);
      })
      .catch(error => {
        if (error.response.status === 400 && error.response.data === "Request failed with status code 400") {
          setMessage("Неверные email или пароль");
        } else {
          console.log(error.response.data.message || error.message);
          setMessage("Неверные email или пароль");
        }
      });
  }

  return (
    <div className='fmt-form mt-form'>
      <h2>Вход</h2>
      <hr width="13%" />
      <form onSubmit={handleSubmit}>
        <div className="label-box">
          <label>
            <p className="labelp">Email</p>
            <input
              type="email"
              value={email}
              onChange={handleEmailChange}
              placeholder='Enter your email'
              required />
          </label>
        </div>
        <div className="label-box">
          <label>
            <p className="labelp">Password</p>
            <input
              type="password"
              value={password}
              onChange={handlePasswordChange}
              placeholder='Enter your password'
              required />
          </label>
        </div>
        <div className="reg-btn">
          <button type="submit">Войти</button>
        </div>
        <div className="under-auth">
          <p>Нет аккаунта? <NavLink to='/register'>Регистрация</NavLink></p>
        </div>
        {message && <p>{message}</p>}
      </form>
    </div>
  );
}

export default LoginForm;
