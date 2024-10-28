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
Replace <code><mysql-UserName></code> with your MySQL username.
Enter you desire password if youve set it already or else leave it empty and hit enter.

### 3. Install Go Dependencies 

Navigate to the project directory and run the following command to install the required Go dependencies:
```bash
go get fyne.io/fyne/v2
go get github.com/go-sql-driver/mysql
```

### 4. Run the Application

After installing the dependencies, you can run the application with the following command:
```bash
go run hospital_app.go
```
