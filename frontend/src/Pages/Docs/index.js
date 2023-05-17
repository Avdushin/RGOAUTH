import React from 'react'
import { NavLink } from 'react-router-dom';
import './docs.css'
import { APPNAME } from './../../vars';

const Docs = () => {
  return (
    <div className="documentation-container">
      <div className="sidebar">
        <h3>Содержание</h3>
        <div className="sidebar__topics">
          <ul>
            <li><a href="#section-1">Что отакое {APPNAME}?</a></li>
            <li><a href="#handlers">Handlers</a></li>
            <ul><a href="#handlers#/">/ heandler</a></ul>
            <ul><a href="#handlers#/login">/login heandler</a></ul>
            <ul><a href="#handlers#/signup">/signup heandler</a></ul>
            <ul><a href="#handlers#/api/users/:id">/api/users/:id</a></ul>
            <li><a href="#db-mysql">База данных</a></li>
          </ul>
        </div>
      </div>
      <div className="content">
        <h2 id="section-1">Что отакое {APPNAME}?</h2>
        <h3></h3>
        <p>REST-API авторизации <a href="https://react.dev">ReactJS</a> + <a href="https://go.dev">Go</a> + <a href="https://www.mysql.com/">MySQL</a>. <br />
          Данная утилита является хуком для написания REST AUTH. Все данные хранятся в <a href="https://developer.mozilla.org/ru/docs/Web/API/Window/localStorage">localStorage()</a> (локальном хранилище)</p>

        <img src="/img/main/go-box.gif" alt="logo" />

        <h2 id="handlers">Handlers</h2>
        <h3 id='handlers#/'><code>/</code></h3>
        <p className='pod'>main page</p>

        <div className="image">
          <img src="/docs/img/main.jpg" alt="main.jpg" />
        </div>

        <p className="pod">loggined user</p>
        <div className="image">
          <img src="/docs/img/1.jpg" alt="loggined user" />
        </div>

        <h3 id='handlers#/login'><code>/login</code></h3>

        <div className="image">
          <img src="/docs/img/login.jpg" alt="login-screen" />
        </div>

        <p className='pod'>После авторизации пароль хешируется</p>

        <div className="image">
          <img src="/docs/img/hashed-pass.jpg" alt="hashed-pass" />
        </div>

        <h3 id='handlers#/signup'><code>/signup</code></h3>

        <p className='pod'></p>

        <div className="image">
          <img src="/docs/img/signup.jpg" alt="signup" />
        </div>
        <div className="image">
          <img src="/docs/img/signup-process.png" alt="signup-process" />
        </div>

        <h3 id='handlers#/api/users/:id'><code>/api/users/:id</code></h3>

        <p className='pod'>Профиль пользователя</p>
        <div className="image">
          <img src="/docs/img/profile/user_profile.jpg" alt="signup" />
        </div>
        <p className='pod'>Редактирование профиля</p>
        <div className="image">
          <img src="/docs/img/profile/edit_profile.jpg" alt="signup" />
        </div>
        <p className='pod'>Редактировать проверку профиля<br /> Вы не можете изменить свои данные без пароля</p>
        <div className="image">
          <img src="/docs/img/profile/edit_profile_validate.jpg" alt="signup" />
        </div>

        <h3 id='db-mysql'>База данных</h3>

        <p className='pod'>структура:</p>

        <div className="image">
          <img src="/docs/img/db/describe.jpg" alt="users_struct" />
        </div>
        <p className='pod'>пример:</p>
        <div className="image">
          <img src="/docs/img/db/select.jpg" alt="users_example" />
        </div>
      </div>
    </div>
  )
}

export default Docs