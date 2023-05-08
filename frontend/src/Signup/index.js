import React, { useState } from 'react';
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
    const SetLocalStorageItems = () => {
        // set localstorage items
        localStorage.setItem("registrant", username);
        localStorage.setItem('username', username);
        localStorage.setItem('email', email);
        localStorage.setItem('password', password);
    }

    const handleSubmit = (event) => {
        event.preventDefault();       
        axios.post('http://localhost:8080/register', { username, email, password })
            .then(response => {
                setMessage(response.data.message);
                // clear input values
                setUsername('');
                setEmail('');
                setPassword('');
                setIsRegistered(true);
            })
            .catch(error => setMessage(error.response.data.message || error.message));       
    }

    return (
        <div className='fmt-form'>
            <h1>SIGN UP</h1>
            <hr width="13%" />
            <form onSubmit={handleSubmit} onClick={SetLocalStorageItems}>
                <label>
                    Username:
                    <input type="text" value={username} onChange={handleUsernameChange} />
                </label>
                <label>
                    Email:
                    <input type="text" value={email} onChange={handleEmailChange} />
                </label>
                <label>
                    Password:
                    <input type="password" value={password} onChange={handlePasswordChange} />
                </label>
                <div className="reg-btn">
                    <button type="submit">Зарегистрироваться</button>
                </div>
                {message && <p>{message}</p>}
                {isRegistered && <p>Пользователь {registrant} успешно зарегистрирован!</p>}
            </form>
        </div>
    );
}
    
export default RegistrationForm;
