# RGOAUTH

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
![](docs/img/1.png)
  
</details>

<details>
  <summary><code>/login</code> heandler </summary>

  `/login` heandler
\
![login-screen](docs/img/empty.png)

![loggined](docs/img/login.png)
  
</details>

<details>
  <summary><code>/signup</code> heandler </summary>

### `/signup` heandler 
\
![](docs/img/signup.jpg)
  
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

- improve backend architecture
- improve client styles
- responsive client
- debug client