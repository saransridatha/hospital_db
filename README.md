# Hospital Patient Records System
==========================

This is a college project for a hospital patient records system. The 
system allows users to add, fetch, and remove patients from the database.

## Features

*   Add new patients with details such as name, age, gender, contact 
information, diagnosis, and treatment.
*   Fetch all patients from the database.
*   Remove a patient by entering their ID.
*   Display patients in a new window for better user experience.

## Installation
------------

To run this application, follow these steps:

### Step 1: Install MySQL

Download the MySQL installer from the official website and follow the 
installation instructions.

### Step 2: Install Dependencies

Clone the repository and install the dependencies by running the following 
command in your terminal:
```bash
go get -u github.com/fyne-io/fyne/v2
go get -u github.com/fyne-io/fyne/v2/theme
```
Also, install MySQL connector for Go using the following command:
```bash
go get -u mysql/mysql/connector/go
```
### Step 3: Create Database and Tables

Create a new database named "hospital" by running the following SQL script 
in your terminal:
```sql
-- database.sql
CREATE DATABASE hospital;

USE hospital;
CREATE TABLE patients (
    id INT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    age INT NOT NULL,
    gender VARCHAR(50) NOT NULL,
    contact VARCHAR(100) NOT NULL,
    diagnosis VARCHAR(255),
    treatment VARCHAR(255)
);
```
### Step 4: Run the Application

Finally, run the application by executing the following command in your 
terminal:
```bash
go build .
./hospital
```
Replace `./` with the path to the compiled binary file.

## Screenshots
-------------

Here are three screenshots of the application:

[![Screenshot 
1](https://example.com/screenshot1.png)](https://example.com/screenshot1.pn1](https://example.com/screenshot1.png)](https://example.com/creenshot1.png)
[![Screenshot 
2](https://example.com/screenshot2.png)](https://example.com/screenshot2.pn2](https://example.com/screenshot2.png)](https://example.com/creenshot2.png)
[![Screenshot 
3](https://example.com/screenshot3.png)](https://example.com/screenshot3.pn3](https://example.com/screenshot3.png)](https://example.com/creenshot3.png)


