# TASK 5 PBI BTPNS 
## Creating RESTful API Using Golang

## Introduction
Go is a programming language created at Google in 2009 by Robert Griesemer, Rob Pike, and Ken Thompson. Go is an easy, simple, efficient, open-source programming language. This task aims to create a RESTful API containing users and photos. So I created several APIs to manage users and photos data using several external packages from Go.

## Folder Structure
The folder structure I created includes:
- **app**
This folder contains struct user login data to hold the email and password that will be used when logging in later.
- **controllers**
This folder serves as the brain of API creation. This folder contains several functions for Login, Register, Create, Read, Update and Delete data users and photos.
- **database**
This folder contains the configuration of a connect function for connection and auto migrate to the database.
- **helpers**
On this folder I just create a **jwt.go** file for **Generate JWT Token** then return it when the users hit login API.
- **middlewares**
This folder contains a function to authorize users.
- **models**
Inside this folder there is an API protection stored in the Auth function to authenticate whether users can create, read, update and delete photo datas.
- **router**
A folder that serves to organize API endpoints and put middleware functions inside.

## Packages
When I do this task, I used several packages like,
- Gin
- Golang JWT v4
- MySQL Driver
- Gorm
- http
- time
- net/http
- strings
- os

## How To Use
Here are the steps that can help you
1. Clone this repository or Fork it if you want to add some features.
1. Use XAMPP or WAMPP to a local server that includes Apache, MySQL, and PHP programs. (if you don't have XAMPP or WAMPP download it from here [XAMPP](https://www.apachefriends.org/download.html) or [WAMPP](https://www.wampserver.com/en/))
1. After success download then install it.
1. After that, open the xampp controll panel in the folder that was specified when installing. (I use XAMPP for example).
1. Then turn on the Apache and MySQL button.
1. Create database and give it name **db_btpns_final_task**
1. After that open your VS Code, then open the folder that has been cloned from this repository.
1. Create .env file, then type **SECRET_KEY = secret.**. Because I didn't push my .env file to github just for best practice, so you can create your own .env file. My secret key is not only **secret** but I use a string generator. 
1. After that you can open terminal from your VS Code, then start type **go run main.go**
1. The API is working and you can use it.

## API Documentation
### Users Enpoint
- **METHOD: POST -> Register (CREATE)**
http://localhost:2173/api/v1/users/register
- **METHOD: POST -> Login**
http://localhost:2173/api/v1/users/login
- **METHOD: GET -> Get Users (READ)**
http://localhost:2173/api/v1/users
- **METHOD: GET -> Get User By Id**
http://localhost:2173/api/v1/users/:id
- **METHOD: PUT -> Update User (UPDATE)**
http://localhost:2173/api/v1/users/:id
- **METHOD: DELETE -> Delete User (DELETE)**
http://localhost:2173/api/v1/users/:id

### Photos Enpoint
- **METHOD: POST -> Add Photo (CREATE)**
http://localhost:2173/api/v1/users/register
- **METHOD: GET -> Get Photos (READ)**
http://localhost:2173/api/v1/photos
- **METHOD: GET -> Get Photo By Id**
http://localhost:2173/api/v1/photos/:id
- **METHOD: PUT -> Update Photo (UPDATE)**
http://localhost:2173/api/v1/photos/:id
- **METHOD: DELETE -> Delete Photo (DELETE)**
http://localhost:2173/api/v1/photos/:id