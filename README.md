# NWETA

Nweta is a project that aims to spread people's opinions on various topics **Anonymously**.


## Home Page
<img src="https://i.ibb.co/2Z4XVTG/Screen-Shot-2022-12-29-at-4-18-16-PM.png" alt="Screen-Shot-2022-12-29-at-4-18-16-PM" border="0"/>
## Add Newta
<img src="https://i.ibb.co/x6xT6kk/Screen-Shot-2022-12-29-at-4-21-54-PM.png" alt="Screen-Shot-2022-12-29-at-4-21-54-PM" border="0" />

## Requirements
To be able to run the project locally, you will need to install :
- Docker
- Docker-compose 

## Install
- clone the repo :
	`git clone https://github.com/0xPride/project-z.git && cd project-z`
-  Build Docker images:
	`docker-compose up --build`
- browse to [localhost:3000](http://localhost:3000)

## Tech Stack

 - Backend
	 - Golang
	 - [GORM](https://github.com/go-gorm/gorm)
	 - Sqlite
 - Frontend
	 - Svelte Kit
 - Devops
	 - NginX as a reverse proxy
	 - Docker