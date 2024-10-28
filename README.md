# Hospital Patient Records Application

This is a college project for managing hospital patient records. The application allows users to add, fetch, and remove patient records from a MySQL database.

## Features

- Add new patient records
- Fetch and display existing patients
- Remove patient records by ID
- Search for patients by name

## Prerequisites

Before running the application, ensure you have the following installed:

- [Go](https://golang.org/doc/install) (version 1.22 or later)
- [MySQL](https://dev.mysql.com/doc/refman/8.0/en/installing.html)

## Installation Steps

### 1. Install MySQL

Follow the official MySQL installation guide for your operating system. Ensure you have a running MySQL server and have created a database named `hospital_db`.

### 2. Set Up the Database

Execute the SQL commands in the `database.sql` file to create the necessary tables. You can run the script using the MySQL command line:

```bash
mysql -u <mysql-UserName> -p hospital_db < database.sql
```
Replace `<mysql-UserName>` with your MySQL username.
Enter you MySql password if you've set already, else leave it empty and hit enter.
Incase if root password for MySql is not set, 
1. Enter this command in terminal:
   ```bash
   mysql -u root -p
   ```
2. Then enter the following command:(change the root-password parameter to your actual password which you want to set)
   ```mysql>
   ALTER USER 'root'@'localhost' IDENTIFIED BY 'root-password';
   ```   
   

### 3. Install Go Dependencies 

Navigate to the project directory and run the following command to install the required Go dependencies:
```bash
go mod init hospital_db
go get fyne.io/fyne/v2
go get fyne.io/fyne/v2/internal/svg@v2.5.2
go get fyne.io/fyne/v2/internal/painter@v2.5.2
go get fyne.io/fyne/v2/storage/repository@v2.5.2
go get fyne.io/fyne/v2/lang@v2.5.2
go get fyne.io/fyne/v2/widget@v2.5.2
go get fyne.io/fyne/v2/internal/driver/glfw@v2.5.2
go get fyne.io/fyne/v2/app@v2.5.2
go get github.com/go-sql-driver/mysql
```

### 4. Run the Application
1. After installing the dependencies, in the `hospital_app.go` file change the `<mysql-UserName>` to your actual MySql username and `<mysql-Password>` to your actual MySql password.
2. After changing these two parameters, you can run the application with the following command:
```bash
go run hospital_app.go
```

## Screenshots

<div style="display: flex; justify-content: space-around;">
  <img src="/screenshots/1.png" alt="Image 1" width="501">
  <img src="/screenshots/2.png" alt="Image 2" width="501">
  <img src="/screenshots/3.png" alt="Image 3" width="501">
  <img src="/screenshots/5.png" alt="Image 2" width="501">
  <img src="/screenshots/6.png" alt="Image 3" width="501">
  <img src="/screenshots/4.png" alt="Image 1">
</div>


