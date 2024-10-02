# green_environment_app

## About Project
Green_Environment_APP is an echo-friendly application designed to encourage users to participate in activities that promote environmental sustainability. The application allows users to access environment-related challenges, purchase eco-friendly products, and engage in a discussion forum to raise awareness about climate change and other environmental concerns. Admin can manage users, productss, challenges, rewards, and moderate the forum. This was built using clean architecture principles, ensuring a well-structured and scalable codebase.

### Features

# User
 __Register and Login:__
 Users can create an account and authenticate with JWT based login.
 __purchase Products:__
 Access a marketplace to purchase eco-friendly products that support a green lifestyle.
 __Participate in Challenges:__ 
 Users can browse and take part in various environmental challenges, earning rewards for successfuly completion.

# Admin
 __User Management:__
 Admins can view, edit, or delete user accounts.
 __Product and Reward Management:__
 Admins can add, update, or delete products and rewards in the system.
 __Challenge Moderation:__ 
 Admins can create, update, or delete environment related challenges.

 #### Tech Stacks
 __Golang__: Backend programming language used for building the application.
 __Echo Framework__: Web framework for Go used to create the RESTful API.
 __GORM__: ORM library for Go, used for database management.
 __MySQL__: Database used to store user data, products, challenges, rewards and more.
 __JWT (JSON Web Token)__: Used for securing the API by providing authentication and authorization mechanisms.
 __Postman__: The postman tool used to test API endpoints.
 __Docker__: Containerization tool used to simplify deployment and scalability.
 __Koyeb__: Serverless deployment platform where the application is hosted.

 ##### API Documentation
 The API docomentation was created using Postman. You can access the documentation for all the avilable routes here:
 https://documenter.getpostman.com/view/37379531/2sAXxJib3c

 # Entity Relationship Diagram (ERD)
 Below is the ERD for the project:
 ![Project_ERD](https://github.com/user-attachments/assets/148f5306-668f-423f-bf2d-459ff6a3b70b)


 ###### Setup
 To run the project locally, folow these steps:

__Prerequisites__
 1- Make sure you installed Go on your system
 2- Make sure you installed MySQL and it is running
 3- Docker if you are using docker for deployment
 4- Git

__Steps__
 1- Clone the repository: git clone https://github.com/hamdam-hadi/green_environment_app cd green_environment_app
 2- Set up environment variables
 DB_USER=root
 DB_PASSWORD=
 DB_HOST=localhost
 DB_PORT=3306
 DB_NAME=green_environment_app
 JWT_SECRET=secret
 PORT=2024
 3- Install dependencies: You can install all Go module dependencies by running "go mod tidy"
 4- Run the application: go run main.go
 5- Build the docker image: docker build -t green_environment_app .
 6- Database Migration: Make sure your MySQL server is running and the database configuration is correctly set up in the .env file
 7- Access the application: The app will run on http://localhost:2024 or you can change port from .env file




