# green_environment_app

# About Project
Green_Environment_APP is an echo-friendly application designed to encourage users to participate in activities that promote environmental sustainability. The application allows users to access environment-related challenges, purchase eco-friendly products, and engage in a discussion forum to raise awareness about climate change and other environmental concerns. Admin can manage users, productss, challenges, rewards, and moderate the forum. This was built using clean architecture principles, ensuring a well-structured and scalable codebase.

# Features

- User
 . Register and Login: Users can create an account and authenticate with JWT based login.
 . purchase Products: Access a marketplace to purchase eco-friendly products that support a green lifestyle.
 . Participate in Challenges: Users can browse and take part in various environmental challenges, earning rewards for successfuly completion.

 - Admin
 . User Management: Admins can view, edit, or delete user accounts.
 . Product and Reward Management: Admins can add, update, or delete products and rewards in the system.
 . Challenge Moderation: Admins can create, update, or delete environment related challenges.

 # Tech Stacks
 . Golang: Backend programming language used for building the application.
 . Echo Framework: Web framework for Go used to create the RESTful API.
 . GORM: ORM library for Go, used for database management.
 . MySQL: Database used to store user data, products, challenges, rewards and more.
 . JWT (JSON Web Token): Used for securing the API by providing authentication and authorization mechanisms.
 . Postman: The postman tool used to test API endpoints.
 . Docker: Containerization tool used to simplify deployment and scalability.
 . Koyeb: Serverless deployment platform where the application is hosted.

 # API Documentation
 The API docomentation was created using Postman. You can access the documentation for all the avilable routes here:
 https://documenter.getpostman.com/view/37379531/2sAXxJib3c

 # Entity Relationship Diagram (ERD)
 Below is the ERD for the project:
 https://www.drawdb.app/editor

 # Setup
 To run the project locally, folow these steps:

 - Prerequisites
 . Make sure you installed Go on your system
 . Make sure you installed MySQL and it is running
 . Docker if you are using docker for deployment
 . Git

 - Steps
 . Clone the repository: git clone https://github.com/hamdam-hadi/green_environment_app.git cd green_environment_app
 . Set up environment variables"
 DB_USER=root
 DB_PASSWORD=
 DB_HOST=localhost
 DB_PORT=3306
 DB_NAME=green_environment_app
 JWT_SECRET=secret
 PORT=2024
 . Install dependencies: You can install all Go module dependencies by running 
 go mod tidy
 . Run the application: go run main.go
 . Build the docker image: docker build -t green_environment_app .
 . Database Migration: Make sure your MySQL server is running and the database configuration is correctly set up in the .env file
 . Access the application: The app will run on http://localhost:2024 or you can change port from .env file




