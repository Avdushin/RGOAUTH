import React, { useState } from 'react';
import axios from 'axios';

function LoginForm() {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [message, setMessage] = useState('');

  const handleEmailChange = (event) => {
    setEmail(event.target.value);
  }

  const handlePasswordChange = (event) => {
    setPassword(event.target.value);
  }

  const handleSubmit = (event) => {
    event.preventDefault();
    axios.post('http://localhost:8080/login', { email, password })
      .then(response => {
        setMessage(response.data.message);
        localStorage.setItem('username', response.data.username);
        localStorage.setItem('email', response.data.email);
        localStorage.setItem('password', response.data.password);
        setTimeout(() => {
          window.location.href = "/";
        }, 1000);
      })
      .catch(error => setMessage(error.response.data.message || error.message));
  }

  return (
    <div className='fmt-form'>
      <h2>Вход</h2>
      <hr width="13%" />
      <form onSubmit={handleSubmit}>
        <label>
          Email:
          <input type="email" value={email} onChange={handleEmailChange} />
        </label>
        <label>
          Password:
          <input type="password" value={password} onChange={handlePasswordChange} />
        </label>
        <div className="reg-btn">
          <button type="submit">Login</button>
        </div>
        {message && <p>{message}</p>}
      </form>
    </div>
  );
}

export default LoginForm;
