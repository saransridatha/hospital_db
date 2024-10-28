package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"strings"

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
	db, err = sql.Open("mysql", "<mysql-UserName>:<mysql-Password>@tcp(localhost:3306)/hospital_db")
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
func clearEntries(entries ...interface{}) {
	for _, entry := range entries {
		switch e := entry.(type) {
		case *widget.Entry:
			e.SetText("")
		case *widget.Select:
			e.SetSelected("") // Clear the selection
		}
	}
}

// Function to display messages
func showMessage(window fyne.Window, message string, err error) {
	if err != nil {
		message = fmt.Sprintf("%s: %v", message, err)
	}
	dialog.NewInformation("Information", message, window).Show()
}

// Function to remove a patient record
func removePatient(patientID int) error {
	_, err := db.Exec("DELETE FROM patients WHERE id = ?", patientID)
	return err
}

// Function to create a new window for displaying patients
func displayPatients(patients []Patient, a fyne.App) {
	patientWindow := a.NewWindow("Patient List")
	patientWindow.Resize(fyne.NewSize(1280, 768)) // Increased size of the window

	// Create a container to hold patient records using a vertical box
	patientList := container.NewVBox()

	// Create a header row with a grid layout
	headerRow := container.NewGridWithColumns(8,
		widget.NewLabel("ID"),
		widget.NewLabel("Name"),
		widget.NewLabel("Age"),
		widget.NewLabel("Gender"),
		widget.NewLabel("Contact"),
		widget.NewLabel("Address"),
		widget.NewLabel("Diagnosis"),
		widget.NewLabel("Treatment"),
	)

	// Set bold style for header labels
	for _, label := range headerRow.Objects {
		if lbl, ok := label.(*widget.Label); ok {
			lbl.TextStyle = fyne.TextStyle{Bold: true}
		}
	}
	patientList.Add(headerRow) // Add the header row to the patient list

	// Create a search entry
	searchEntry := widget.NewEntry()
	searchEntry.SetPlaceHolder("Search by name...")

	// Function to filter patients based on search input
	filterPatients := func() {
		searchTerm := searchEntry.Text

		// Clear previous patient rows while keeping the header and search bar
		patientList.Objects = []fyne.CanvasObject{headerRow, searchEntry}

		// Loop through the original patient list to filter
		for _, patient := range patients {
			if searchTerm == "" || strings.Contains(patient.Name, searchTerm) {
				row := container.NewGridWithColumns(8,
					widget.NewLabel(fmt.Sprintf("%d", patient.ID)),
					widget.NewLabel(patient.Name),
					widget.NewLabel(fmt.Sprintf("%d", patient.Age)),
					widget.NewLabel(patient.Gender),
					widget.NewLabel(patient.Contact),
					widget.NewLabel(patient.Address),
					widget.NewLabel(patient.Diagnosis),
					widget.NewLabel(patient.Treatment),
				)
				patientList.Add(row) // Add each patient row to the patient list
			}
		}

		patientList.Refresh() // Refresh the patient list to update the display
	}

	// Bind the filter function to the search entry's OnChanged event
	searchEntry.OnChanged = func(string) {
		filterPatients()
	}

	// Add the search entry to the patient list (after the header)
	patientList.Add(searchEntry)

	// Enable scrolling for the patient list
	scrollContainer := container.NewVScroll(patientList) // Scroll container around the patient list

	// Add the scroll container to the patient window
	patientWindow.SetContent(scrollContainer)

	patientWindow.Show()
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
	genderEntry := widget.NewSelect([]string{"Male", "Female", "Other"}, func(selected string) {
		// Handle gender selection
	})
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
			Gender:    genderEntry.Selected,
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

		displayPatients(patients, a) // Show patients in a new window
	})

	// Button to remove a patient
	removeButton := widget.NewButton("Remove Patient", func() {
		idEntry := widget.NewEntry()
		idEntry.SetPlaceHolder("Enter Patient ID to remove")

		dialog.ShowCustomConfirm("Remove Patient", "Remove", "Cancel", idEntry, func(confirmed bool) {
			if confirmed {
				patientID, err := strconv.Atoi(idEntry.Text)
				if err != nil {
					showMessage(w, "Invalid Patient ID.", err)
					return
				}

				if err := removePatient(patientID); err != nil {
					showMessage(w, "Failed to remove patient.", err)
					return
				}

				showMessage(w, "Patient removed successfully!", nil)
			}
		}, w)
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
		removeButton, // Add the remove patient button
	)

	w.SetContent(form)
	w.Resize(fyne.NewSize(400, 400))
	w.ShowAndRun()
}
