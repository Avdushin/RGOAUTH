import React, { useState }from 'react';
import { NavLink } from 'react-router-dom';
import axios from 'axios';
import '../App.css'

function RegistrationForm() {
    const [username, setUsername] = useState('');
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [message, setMessage] = useState('');
    const [isRegistered, setIsRegistered] = useState(false);
    let registrant = localStorage.getItem("registrant");

    // heandler for registration
    const handleUsernameChange = (event) => {
        setUsername(event.target.value);
    }
    const handleEmailChange = (event) => {
        setEmail(event.target.value);
    }
    const handlePasswordChange = (event) => {
        setPassword(event.target.value);
    }

    // set LocalStorage items
    // const SetLocalStorageItems = () => {
    //     // set localstorage items
    //     localStorage.setItem("registrant", username);
    //     localStorage.setItem('username', username);
    //     localStorage.setItem('email', email);
    //     localStorage.setItem('password', password);
    // }

    const handleSubmit = (event) => {
        event.preventDefault();
        axios.post(`${process.env.REACT_APP_SERVER_URL}/register`, { username, email, password })
            .then(response => {
                setMessage(response.data.message);
                // clear input values
                setUsername('');
                setEmail('');
                setPassword('');
                setIsRegistered(true);
                setTimeout(() => {
                    window.location.href = "/login";
                }, 1000);
            })
            .catch(error => setMessage(error.response.data.message || error.message));
    }

    return (
        <div className='fmt-form mt-form'>
            <h1>Регистрация</h1>
            <hr width="13%" />
            <form onSubmit={handleSubmit} >
                {/* onClick={SetLocalStorageItems} */}
                <div className="label-box">
                    <label>
                        <p className="labelp">Username</p>
                        <input
                        type="text"
                        value={username}
                        onChange={handleUsernameChange}
                        placeholder='Enter username'
                        required/>
                    </label>
                </div>
                <div className="label-box">
                    <label>
                        <p className="labelp">Email</p>
                        <input
                        type="text"
                        value={email}
                        onChange={handleEmailChange} 
                        placeholder='Enter email'
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
                        placeholder='Enter password'
                        required />
                    </label>
                </div>
                <div className="reg-btn">
                    <button type="submit">Зарегистрироваться</button>
                </div>
                <div className="under-auth">
                    <p>Уже есть аккаунт? <NavLink to='/login'>Войти</NavLink></p>
                </div>
                {message && <p>{message}</p>}
                {isRegistered && <p>Пользователь {registrant} успешно зарегистрирован!</p>}
            </form>
        </div>
    );
}

export default RegistrationForm;
