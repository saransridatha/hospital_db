CREATE DATABASE IF NOT EXISTS hospital_db;
USE hospital_db;

DROP TABLE IF EXISTS patients;

CREATE TABLE patients (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100),
    age INT,
    gender VARCHAR(10),
    contact VARCHAR(255),
    address VARCHAR(255),
    diagnosis TEXT,
    treatment TEXT
);
