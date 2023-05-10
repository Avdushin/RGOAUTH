import React from 'react'
import err404 from '../assets/img/404/meditation-go.png'
import './404.css'

function Err404() {
  return (
    <div className="not-found-box">
      <div className="not-found">
        <img src={err404} alt="not-found" />
      </div>
      <div className="not-found_message">
        <p className="error-title">ERROR 404</p>
        <p className='error-status'>Страница не найдена</p>
      </div>
    </div>
  )
}

export default Err404