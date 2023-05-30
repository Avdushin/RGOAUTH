import React, { useState, useEffect } from 'react';
import axios from 'axios';
import './userProfile.css'

function UserProfile() {
  const [username, setUsername] = useState('');
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [name, setName] = useState('');
  const [surname, setSurname] = useState('');
  const [patronymic, setPatronymic] = useState('');
  const [age, setAge] = useState('');
  const [editing, setEditing] = useState(false);
  const [message, setMessage] = useState('');

  useEffect(() => {
    // Fetch user data from the server
    axios.get(`${process.env.REACT_APP_SERVER_URL}/api/users/${localStorage.getItem('id')}`)
      .then((response) => {
        const user = response.data;
        setUsername(user.username);
        setEmail(user.email);
        setName(user.name);
        setSurname(user.surname);
        setPatronymic(user.patronymic);
        setAge(user.age);
        localStorage.setItem("username", user.username);
        localStorage.setItem("email", user.email);
        localStorage.setItem("name", user.name);
        localStorage.setItem("surname", user.surname);
        localStorage.setItem("patronymic", user.patronymic);
        localStorage.setItem("age", user.age);
      })
      .catch((error) => {
        console.log(`Server: ${error.response.data.message || error.message}`)
        setMessage(error.response.data.message || error.message);
      });
  }, []);

  const handleUsernameChange = (event) => {
    setUsername(event.target.value);
  };

  const handleEmailChange = (event) => {
    setEmail(event.target.value);
  };

  const handlePasswordChange = (event) => {
    setPassword(event.target.value);
  };
  

  const handleNameChange = (event) => {
    setName(event.target.value);
  };

  const handleSurnameChange = (event) => {
    setSurname(event.target.value);
  };

  const handlePatronymicChange = (event) => {
    setPatronymic(event.target.value);
  };

  const handleAgeChange = (event) => {
    setAge(event.target.value);
  };

    const handleEditProfile = () => {
        setEditing(true);
    };

    const handleViewProfile = () => {
        if (password !== "") {
            return
        }
        setEditing(false);
    };

    const handleSubmit = (event) => {
        event.preventDefault();
      
        axios.defaults.headers.post['Content-Type'] = 'application/json';
      
        // валидация поля Password
        if (password === '') {
          window.alert('Введите ваш пароль');
          return; // Останавливаем выполнение функции, если поле Password пустое
        }
      
        // Отправка обновленных данных пользователя на сервер
        axios
          .put(`${process.env.REACT_APP_SERVER_URL}/api/users/${localStorage.getItem('id')}`, {
            username,
            email,
            password,
            name,
            surname,
            patronymic,
            age
          })
          .then((response) => {
            setMessage(response.data.message);
            // Обновление данных пользователя в localStorage
            localStorage.setItem('username', username);
            localStorage.setItem('email', email);
            localStorage.setItem('name', name);
            localStorage.setItem('surname', surname);
            localStorage.setItem('patronymic', patronymic);
            localStorage.setItem('age', age);
          })
          .catch((error) => {
            setMessage(error.response.data.message || error.message);
          });
          setTimeout(() => {
            window.location.reload();
        }, 1000);
      }; 

  if (editing) {
    return (
        <div className='user-profile profile_container'>
            <h1>Настройки профиля</h1>
            <form onSubmit={handleSubmit}>
                <div>
                    <label>
                        Имя пользователя:
                        <input type="text" value={username} onChange={handleUsernameChange} />
                    </label>
                </div>
                <div>
                    <label>
                        Email:
                        <input type="text" value={email} onChange={handleEmailChange} />
                    </label>
                </div>
                <div>
                    <label>
                        Password:
                        <input type="password" value={password} onChange={handlePasswordChange} required/>
                    </label>
                </div>
                <div>
                    <label>
                        Имя:
                        <input type="text" value={name} onChange={handleNameChange} />
                    </label>
                </div>
                <div>
                    <label>
                        Фамилия:
                        <input type="text" value={surname} onChange={handleSurnameChange} />
                    </label>
                </div>
                <div>
                    <label>
                        Отчество:
                        <input type="text" value={patronymic} onChange={handlePatronymicChange} />
                    </label>
                </div>
                <div>
                    <label>
                        Возраст:
                        <input type="number" value={age} max="130" onChange={handleAgeChange} />
                    </label>
                </div>
                <div className='edit-zone-btn'>
                    <button className='save-btn profile-btn' type="submit" onSubmit={handleViewProfile}>Сохранить</button>
                    <button className='cancel-btn profile-btn' type="submit" onClick={handleViewProfile}>Отмена</button>
                </div>
            </form>
            {message && <p>{message}</p>}
        </div>
    );
} else {
    return (
        <div className='profile_container user-profile'>
            <h1>Профиль</h1>
            <div className='profile-user-data'>
                <strong>Имя пользователя:</strong> {username}
            </div>
            <div className='profile-user-data'>
                <strong>Email:</strong> {email}
            </div>
            <div className='profile-user-data'>
                <strong>Имя:</strong> {name}
            </div>
            <div className='profile-user-data'>
                <strong>Фамилия:</strong> {surname}
            </div>
            <div>
                <strong>Отчество:</strong> {patronymic}
            </div>
            <div className='profile-user-data'>
                <strong>Возраст:</strong> {age}
            </div>
            <div className='profile-user-data'>
                <button className='edit-btn profile-btn' onClick={handleEditProfile}>Редактировать профиль</button>
            </div>
        </div>
    );
}
}

export default UserProfile;