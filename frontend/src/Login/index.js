import React, { useState } from 'react';
import axios from 'axios';

function LoginForm() {
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
