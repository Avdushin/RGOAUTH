import React from 'react'
import { NavLink } from 'react-router-dom';
import mainLogo from './img/go-box.gif'
import Footer from './../Footer/index';
import './main.css'

function Main() {
  return (
    <div className="main">
        <div className="intro">
          <img src={mainLogo} alt="logo" />
        </div>
        <div className="descr">
            <h1>RGOAUTH</h1>
            <p><a href="https://react.dev">ReactJS</a> + <a href="https://go.dev">Go</a> authorithation REST application</p>
            <p>By use this app you can <u>
              <NavLink to='/register'>SignUp</NavLink>/<NavLink to='/login'>Login</NavLink></u> users, and add all records at the <a href="https://www.mysql.com/">MySQL</a> DB.</p>

              <div className="documentation">
                <NavLink to='/docs'>Documentation</NavLink>
              </div>
        </div> 
    </div>
  )
}

export default Main