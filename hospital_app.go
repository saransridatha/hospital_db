package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
        "fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	_ "github.com/go-sql-driver/mysql"
)

type Patient struct {
	ID        int
	Name      string
	Age       int
	Gender    string
	Contact   string
	Address   string
	Diagnosis string
	Treatment string
}

var db *sql.DB

// Initialize the database connection
func initDB() {
	var err error
	db, err = sql.Open("mysql", "<mysql-UserName>:<mysql-password>@tcp(localhost:3306)/hospital_db")
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal("Database is unreachable:", err)
	}
}

// Function to add a new patient to the database
func addPatient(patient Patient) error {
	_, err := db.Exec("INSERT INTO patients (name, age, gender, contact, address, diagnosis, treatment) VALUES (?, ?, ?, ?, ?, ?, ?)",
		patient.Name, patient.Age, patient.Gender, patient.Contact, patient.Address, patient.Diagnosis, patient.Treatment)
	return err
}

// Function to fetch all patients from the database
func fetchPatients() ([]Patient, error) {
	rows, err := db.Query("SELECT id, name, age, gender, contact, address, diagnosis, treatment FROM patients")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var patients []Patient
	for rows.Next() {
		var p Patient
		if err := rows.Scan(&p.ID, &p.Name, &p.Age, &p.Gender, &p.Contact, &p.Address, &p.Diagnosis, &p.Treatment); err != nil {
			return nil, err
		}
		patients = append(patients, p)
	}
	return patients, nil
}

// Function to clear entry fields
func clearEntries(entries ...*widget.Entry) {
	for _, entry := range entries {
		entry.SetText("")
	}
}

// Function to display messages
func showMessage(window fyne.Window, message string, err error) {
	if err != nil {
		message = fmt.Sprintf("%s: %v", message, err)
	}
	dialog.NewInformation("Information", message, window).Show()
}

// Main function
func main() {
	initDB()
	defer db.Close()

	a := app.New()
	w := a.NewWindow("Hospital Patient Records")

	// Create entry fields for patient details
	nameEntry := widget.NewEntry()
	ageEntry := widget.NewEntry()
	genderEntry := widget.NewEntry()
	contactEntry := widget.NewEntry()
	addressEntry := widget.NewEntry()
	diagnosisEntry := widget.NewEntry()
	treatmentEntry := widget.NewEntry()

	// Button to add a patient
	addButton := widget.NewButton("Add Patient", func() {
		if nameEntry.Text == "" || ageEntry.Text == "" {
			showMessage(w, "Name and Age cannot be empty.", nil)
			return
		}

		age, err := strconv.Atoi(ageEntry.Text)
		if err != nil {
			showMessage(w, "Invalid age.", err)
			return
		}

		patient := Patient{
			Name:      nameEntry.Text,
			Age:       age,
			Gender:    genderEntry.Text,
			Contact:   contactEntry.Text,
			Address:   addressEntry.Text,
			Diagnosis: diagnosisEntry.Text,
			Treatment: treatmentEntry.Text,
		}

		// Add the patient to the database
		if err := addPatient(patient); err != nil {
			log.Println("Error adding patient:", err)
			showMessage(w, "Failed to add patient", err)
			return
		}

		clearEntries(nameEntry, ageEntry, genderEntry, contactEntry, addressEntry, diagnosisEntry, treatmentEntry)
		showMessage(w, "Patient added successfully!", nil)
	})

	// Button to fetch patients
	fetchButton := widget.NewButton("Fetch Patients", func() {
		patients, err := fetchPatients()
		if err != nil {
			log.Println("Error fetching patients:", err)
			showMessage(w, "Failed to fetch patients", err)
			return
		}

		if len(patients) == 0 {
			showMessage(w, "No patients found.", nil)
			return
		}

		patientList := "Patients:\n"
		for _, p := range patients {
			patientList += fmt.Sprintf("ID: %d, Name: %s, Age: %d, Gender: %s, Contact: %s, Address: %s, Diagnosis: %s, Treatment: %s\n",
				p.ID, p.Name, p.Age, p.Gender, p.Contact, p.Address, p.Diagnosis, p.Treatment)
		}
		dialog.NewInformation("Patient List", patientList, w).Show()
	})

	// Create the UI layout
	form := container.NewVBox(
		widget.NewLabel("Name:"), nameEntry,
		widget.NewLabel("Age:"), ageEntry,
		widget.NewLabel("Gender:"), genderEntry,
		widget.NewLabel("Contact:"), contactEntry,
		widget.NewLabel("Address:"), addressEntry,
		widget.NewLabel("Diagnosis:"), diagnosisEntry,
		widget.NewLabel("Treatment:"), treatmentEntry,
		addButton,
		fetchButton,
	)

	w.SetContent(form)
	w.Resize(fyne.NewSize(400, 400))
	w.ShowAndRun()
}
