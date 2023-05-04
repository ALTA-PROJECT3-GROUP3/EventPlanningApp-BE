## 📑 About the Project
<p align="justify">Event Planning App. Users can post event and users can join events from each other.<br>
  <br>
This RESTful API was developed by using Golang and written based on Clean Architecture principles. Built with Echo as web framework, GORM as ORM, MySQL as DBMS, etc.
</p>

## 🛠 Tools
**Backend:** <br>
![GitHub](https://img.shields.io/badge/github-%23121011.svg?style=for-the-badge&logo=github&logoColor=white)
![Visual Studio Code](https://img.shields.io/badge/Visual%20Studio%20Code-0078d7.svg?style=for-the-badge&logo=visual-studio-code&logoColor=white)
![MySQL](https://img.shields.io/badge/mysql-%2300f.svg?style=for-the-badge&logo=mysql&logoColor=white)
![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![JWT](https://img.shields.io/badge/JWT-black?style=for-the-badge&logo=JSON%20web%20tokens)
![Swagger](https://img.shields.io/badge/-Swagger-%23Clojure?style=for-the-badge&logo=swagger&logoColor=white)
![Postman](https://img.shields.io/badge/Postman-FF6C37?style=for-the-badge&logo=postman&logoColor=white)

**Deployment:** <br>
![AWS](https://img.shields.io/badge/AWS-%23FF9900.svg?style=for-the-badge&logo=amazon-aws&logoColor=white)
![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white)
![Ubuntu](https://img.shields.io/badge/Ubuntu-E95420?style=for-the-badge&logo=ubuntu&logoColor=white)
![Cloudinary](https://img.shields.io/badge/Cloudinary-F38020?style=for-the-badge&logo=Cloudflare&logoColor=white)
![Midtrans](https://img.shields.io/badge/Midtrans-F38020?style=for-the-badge&logo=Midtrans&logoColor=white)

**Communication:**  
![GitHub](https://img.shields.io/badge/github%20Project-%23121011.svg?style=for-the-badge&logo=github&logoColor=white)
![Discord](https://img.shields.io/badge/Discord-%237289DA.svg?style=for-the-badge&logo=discord&logoColor=white)

# 🔗 ERD
<img src="ERD.jpg">

# 🔥 Open API

Simply [click here](https://app.swaggerhub.com/apis-docs/CW3-ALTA/EventPlanningApp/1.0.0) to see the details of endpoints we have agreed with our FE team.

<details>
  <summary>👶 User </summary>
  
| Method      | Endpoint            | Params      |q-Params            | JWT Token   | Function                                |
| ----------- | ------------------- | ----------- |--------------------| ----------- | --------------------------------------- |
| POST        | /register           | -           |-                   | NO          | Register a new Use                      |
| POST        | /login              | -           |-                   | NO          | Login to the system                     |
| GET         | /users              | -           |-                   | YES         | Show user profile                       |
| PUT         | /users              | -           |-                   | YES         | Update user profile                     |
| DELETE      | /users              | -           |-                   | YES         | Update user profile                     |


  
</details>

<details>
  <summary>📑 Event</summary>
  
| Method      | Endpoint            | Params      | JWT Token   | Function                                |
| ----------- | ------------------- | ----------- | ----------- | --------------------------------------- |
| POST        | /events             | -           | YES         | Post a events                            |
| GET         | /events             | -           | NO          | Get All event                            |
| GET         | /myevent            | -           | YES         | Get MyEvents                             |
| PUT         | /events             | events_id   | YES         | Edit event                              |
| DELETE      | /events             | events_id   | YES         | Delete event                             |
| GET         | /events             | events_id   | NO          | Get events Detail                        |  

  </details>
     <details>
  <summary>📠 Ticket</summary>
  
| Method      | Endpoint            | Params      | JWT Token   | Function                                |
| ----------- | ------------------- | ----------- | ----------- | --------------------------------------- |
| POST        | /tickets            | -           | YES         | Make Event Ticket                   |
| PUT         | /tickets            | tickets_id  | YES         | Edit Ticket                    |


  </details>
  <details>
   <summary>🔊 Comment</summary>
  
| Method      | Endpoint            | Params      | JWT Token   | Function                                |
| ----------- | ------------------- | ----------- | ----------- | --------------------------------------- |
| POST        | /comments           | -           | YES         | Make Event Comment                          |


</details>
  <details>
   <summary>📑 Payment</summary>
  
| Method      | Endpoint                      | Params      | JWT Token   | Function                                |
| ----------- | ----------------------------- | ----------- | ----------- | --------------------------------------- |
| POST        | /reservations                 | -           | YES         | Make Reservations for Join Event        |
| POST        | /payments/notifications       | -           | YES         | Make payments notifications        |



  </details>
    
 

# 🛠️ How to Run Locally

- Clone it

```
$ git clone https://github.com/ALTA-PROJECT3-GROUP3/EventPlanningApp-BE.git
```

- Go to directory

```
$ cd EventPlanningApp-BE
```
- Run the project
```
$ go run .
```

- Voila! 🪄

### 🧰Backend

- [Github Repository for the Backend team](https://github.com/ALTA-PROJECT3-GROUP3/EventPlanningApp-BE)
- [Swagger OpenAPI](https://app.swaggerhub.com/apis-docs/CW3-ALTA/EventPlanningApp/1.0.0)


# 🤖 Author

-  Adi Yuda Pranata  <br>  [![GitHub](https://img.shields.io/badge/Yuda-%23121011.svg?style=for-the-badge&logo=github&logoColor=white)](https://github.com/Adiyuda123)
-  Kristain Putra <br>  [![GitHub](https://img.shields.io/badge/Iqbal-%23121011.svg?style=for-the-badge&logo=github&logoColor=white)](https://github.com/kristain09)
-  Haris <br>  [![GitHub](https://img.shields.io/badge/Wanta-%23121011.svg?style=for-the-badge&logo=github&logoColor=white)](https://github.com/ares0177)



<h5>
<p align="center">Created by Group 3 ©️ 2023</p>
</h5>
