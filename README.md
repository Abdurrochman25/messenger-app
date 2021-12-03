# Messenger-App
[![Go Reference](https://pkg.go.dev/badge/golang.org/x/example.svg)](https://pkg.go.dev/golang.org/x/example)
[![Go.Dev reference](https://img.shields.io/badge/gorm-reference-blue?logo=go&logoColor=white)](https://pkg.go.dev/gorm.io/gorm?tab=doc)
[![Go.Dev reference](https://img.shields.io/badge/echo-reference-blue?logo=go&logoColor=white)](https://github.com/labstack/echo)

# Table of Content
- [Description](#description)
- [How to Use](#how-to-use)
- [Database Schema](#database-schema)
- [Feature](#feature)
- [Endpoints](#endpoints)

# Description
Messenger App merupakan sebuah aplikasi prototype Chatting seperti WhatsApp. Disini pengguna dapat saling mengirim pesan ke sesama pengguna lainnya, pengguna juga dapat melihat seluruh obrolan yang dilakukannya dengan pengguna lain.

# Database Scheme
![ERD] (https://github.com/Abdurrochman25/messenger-app/blob/main/screenshoot/messeger-ERD.png)

# Feature 
List of overall feature in this Project (To get more details see the API Documentation below)
| No.| Feature           | Keterangan                                                             |
| :- | :---------------- | :--------------------------------------------------------------------- |
| 1. | Register          | Authentication Process                                                 |
| 2. | Login             | Authentication Process                                                 |
| 3. | CR Message        | (Create, Read) Send message to other user & Get message by receiver id |
| 4. | Read Conversation | Get all conversation                                                   |

# How to Use
- Install Go and Database MySQL/XAMPP
- Clone this repository in your $PATH:
```
$ git clone https://github.com/Abdurrochman25/messenger-app.git
```
- Create file .env based on this project 
``
sample-env
``
- Don't forget to create database name as you want in your MySQL
- Run program with command
```
go run app/main.go
```

# Endpoints
Read the API documentation here [API Endpoint Documentation] (https://app.swaggerhub.com/apis/m.abdurrochman25/messenger-app/1.0.0) (Swagger)
