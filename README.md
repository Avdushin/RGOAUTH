# RGOAUTH

[![pwreview](docs/img/prew.gif)](https://youtu.be/t2KiHOXKsUA)

<!-- [![pwreview]()](https://youtu.be/t2KiHOXKsUA) [DEMO]() -->

<div align="center">
  <a href="https://youtu.be/t2KiHOXKsUA" target="_blank">
    <img src="docs/img/i/yt.png" alt="yt-cion" width="12px">
  </a>
  <a href="https://youtu.be/t2KiHOXKsUA" target="_blank">
    DEMO
  </a>
</div> <br>

### REST-API авторизации ReactJS + Go + MySQL.
### Данная утилита является хуком для написания REST AUTH. Все данные хранятся в `localStorage()` (локальном хранилище)


## Heandlers

<details>
  <summary><code>/</code> heandler </summary>
  
  ### main page
  
\
  ![](docs/img/rgoauth.png)

  ### loggined user
\
![](docs/img/login.jpg)
  
</details>

<details>
  <summary><code>/login</code> heandler </summary>

  `/login` heandler

![login-screen](docs/img/login.jpg)

### After authorization, the password is hashed
\
  ![](docs/img/hashed-pass.jpg)
  
</details>

<details>
  <summary><code>/signup</code> heandler </summary>

### `/signup` heandler 
\
![](docs/img/signup-process.png)
  
</details>

<details>
  <summary><code>/api/users/:id</code> heandler </summary>

### User profile
\
![](./docs/img/user_profile.jpg)

### Edit user profile
\
![](./docs/img/edit_profile.jpg)

### Validation
\
![](./docs/img/edit_profile_validate.jpg)

</details>

## БД

<details>
  <summary>База данных</summary>

### Таблицы

![](docs/img/db/tables.jpg)

### Таблица Users
  
структура:

![](docs/img/db/describe.jpg)

пример:

![](docs/img/db/select.jpg)

</details>


## How to start?

```bash
# frontend
cd frontend
npm -i
npm start

# backend
cd backend
# build app
go build -o server main.go # make build
# or start server
go run main.go # make
```

## Installing

```bash
gi clone https://github.com/Avdushin/RGOAUTH

# frontend
cd frontend
npm -i
npm start

# backend
cd ../backend
# build app
go build -o server main.go # make build
# start server
./server
```

## TODO

- [x] Make automatic BD backup
- [x] Edit user
- [x] ssl support
- [x] debug client
- [ ] Delete user
- [ ] vite + ReactJS client
- [ ] responsive client