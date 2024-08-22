package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/user"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func addTask(taskName string, taskDate string, taskTime string, taskStatus string) string {
	taskname := strings.TrimSpace(taskName)
	taskdate := strings.TrimSpace(taskDate)
	tasktime := strings.TrimSpace(taskTime)
	taskstatus := strings.TrimSpace(taskStatus)
	creationTime := time.Now().Format("02/01/2006 1504")

	if taskStatus == "" {
		taskStatus = "Ongoing"
	}

	task := fmt.Sprintf("%s,%s,%s,%s,%s\n", taskname, taskdate, tasktime, taskstatus, creationTime)
	return task
}

func EditTask(w fyne.Window, a fyne.App) fyne.CanvasObject {
	taskNameEntry := widget.NewEntry()
	taskNameEntry.SetPlaceHolder("Enter task name to edit")

	editOptions := []string{"Edit Name", "Edit Date", "Edit Time", "Edit Status"}
	selectEditOption := widget.NewSelect(editOptions, func(selected string) {
		// Update the field to edit based on the selection
	})

	feedbackLabel := widget.NewLabel("")
	updateEntry := widget.NewEntry()
	updateEntry.SetPlaceHolder("Enter new value here")
	updateEntry.Hide()

	submitButton := widget.NewButton("Submit", func() {
		taskName := strings.TrimSpace(taskNameEntry.Text)
		if taskName == "" {
			feedbackLabel.SetText("Please enter a task name.")
			return
		}

		selectedOption := selectEditOption.Selected
		if selectedOption == "" {
			feedbackLabel.SetText("Please select a field to edit.")
			return
		}

		newValue := strings.TrimSpace(updateEntry.Text)
		if newValue == "" {
			feedbackLabel.SetText("Please enter a new value.")
			return
		}

		userHome, _ := os.UserHomeDir()
		filepath := userHome + "/Desktop/todoui/todoui.txt"
		file, _ := os.OpenFile(filepath, os.O_RDWR, 0644)
		defer file.Close()

		filereader := bufio.NewReader(file)
		var lines []string
		var found bool
		for {
			line, _ := filereader.ReadString('\n')
			if line == "" {
				break
			}

			line = strings.TrimSpace(line)
			fields := strings.Split(line, ",")
			if fields[0] == taskName {
				found = true
				switch selectedOption {
				case "Edit Name":
					fields[0] = newValue
				case "Edit Date":
					fields[1] = newValue
				case "Edit Time":
					fields[2] = newValue
				case "Edit Status":
					fields[3] = newValue
				}
				lines = append(lines, fmt.Sprintf("%s,%s,%s,%s,%s", fields[0], fields[1], fields[2], fields[3], fields[4]))
			} else {
				lines = append(lines, line)
			}
		}

		if found {
			feedbackLabel.SetText("Task updated successfully!")
			// Write updated lines back to the file
			file.Truncate(0)
			file.Seek(0, 0)
			for _, updatedLine := range lines {
				file.WriteString(updatedLine + "\n")
			}
		} else {
			feedbackLabel.SetText("Name NOT FOUND")
		}
	})

	// Show entry field to input new value based on selected option
	selectEditOption.OnChanged = func(selected string) {
		if selected != "" {
			updateEntry.Show()
		} else {
			updateEntry.Hide()
		}
	}

	centered := container.NewVBox(
		taskNameEntry,
		selectEditOption,
		updateEntry,
		submitButton,
		feedbackLabel,
		widget.NewButton("Back to Main", func() {
			w.SetContent(createMainScreen(w, a))
		}),
	)
	return container.New(layout.NewCenterLayout(), centered)
}

func ListTask(w fyne.Window, a fyne.App) fyne.CanvasObject {
	text := widget.NewLabel("How would you like to list your Tasks?")
	listOptions := []string{"All", "By Date", "By Status"}
	feedbackLabel := widget.NewLabel("")
	newValue := widget.NewEntry()
	newValue.SetPlaceHolder("Enter Value: ")
	selectListOptions := widget.NewSelect(listOptions, func(selected string) {

	})

	submitButton := widget.NewButton("Submit", func() {
		selectedOption := selectListOptions.Selected
		if selectedOption == "" {
			feedbackLabel.SetText("Please select a field to list.")
			return
		}
		if selectedOption == "All" {
			userHome, _ := os.UserHomeDir()
			filepath := userHome + "/Desktop/todoui/todoui.txt"
			file, _ := os.OpenFile(filepath, os.O_RDWR, 0644)
			defer file.Close()

			filereader := bufio.NewReader(file)
			var results []string
			now := time.Now()
			for {
				line, err := filereader.ReadString('\n')
				if err != nil {
					if err == io.EOF {
						break
					}
					fmt.Println("Error:", err)
					return
				}
				line = strings.TrimSpace(line)
				if line == "" {
					continue
				}
				fields := strings.Split(line, ",")
				if len(fields) != 5 {
					continue
				}
				taskName := fields[0]
				taskDate := fields[1]
				taskTime := fields[2]
				taskStatus := fields[3]

				deadlineStr := taskDate + " " + taskTime
				deadlineTime, err := time.Parse("02/01/2006 1504", deadlineStr)
				if err != nil {
					fmt.Println("error parsing time:", err)
				}
				duration := deadlineTime.Sub(now)
				timeLeft := ""
				if duration < 0 {
					timeLeft = "Task overdue"
				} else {
					days := int(duration.Hours()) / 24
					hours := int(duration.Hours()) % 24
					minutes := int(duration.Minutes()) % 60
					timeLeft = fmt.Sprintf("%d days, %d hours, %d minutes", days, hours, minutes)
				}
				result := fmt.Sprintf("%s, %s, %s, %s, %s", taskName, taskDate, taskTime, taskStatus, timeLeft)
				results = append(results, result)
			}
			if len(results) > 0 {
				feedbackLabel.SetText(strings.Join(results, "\n"))
			} else {
				feedbackLabel.SetText("No Tasks found")
			}

		}

		if selectedOption == "By Status" {
			searchStatus := newValue.Text
			userHome, _ := os.UserHomeDir()
			filepath := userHome + "/Desktop/todoui/todoui.txt"
			file, _ := os.OpenFile(filepath, os.O_RDWR, 0644)
			defer file.Close()

			filereader := bufio.NewReader(file)
			var results []string
			now := time.Now()
			//var found bool
			for {
				line, err := filereader.ReadString('\n')
				if err != nil {
					if err == io.EOF {
						break
					}
					fmt.Println("Error:", err)
					return
				}
				line = strings.TrimSpace(line)

				if line == "" {
					continue
				}
				fields := strings.Split(line, ",")
				taskName := fields[0]
				taskDate := fields[1]
				taskTime := fields[2]
				taskStatus := fields[3]
				if searchStatus == taskStatus {
					// Calculate time left or overdue
					deadlineStr := taskDate + " " + taskTime
					deadlineTime, _ := time.Parse("02/01/2006 1504", deadlineStr)
					duration := deadlineTime.Sub(now)
					timeLeft := ""
					if duration < 0 {
						timeLeft = "Task Overdue"
					} else {
						days := int(duration.Hours()) / 24
						hours := int(duration.Hours()) % 24
						minutes := int(duration.Minutes()) % 60
						timeLeft = fmt.Sprintf("%d days, %d hours, %d minutes", days, hours, minutes)
					}

					// Store result
					result := fmt.Sprintf("%s, %s, %s, %s, %s", taskName, taskDate, taskTime, taskStatus, timeLeft)
					results = append(results, result)
				}

			}
			if len(results) > 0 {
				feedbackLabel.SetText(strings.Join(results, ","))
			} else {
				feedbackLabel.SetText("No Tasks found for this Status")
			}

		}
		if selectedOption == "By Date" {

			searchDate := newValue.Text
			userHome, _ := os.UserHomeDir()
			filepath := userHome + "/Desktop/todoui/todoui.txt"
			file, _ := os.OpenFile(filepath, os.O_RDWR, 0644)
			defer file.Close()

			filereader := bufio.NewReader(file)
			var results []string
			now := time.Now()
			//var found bool
			for {
				line, err := filereader.ReadString('\n')
				if err != nil {
					if err == io.EOF {
						break
					}
					fmt.Println("Error:", err)
					return
				}
				line = strings.TrimSpace(line)

				if line == "" {
					continue
				}
				fields := strings.Split(line, ",")
				taskName := fields[0]
				taskDate := fields[1]
				taskTime := fields[2]
				taskStatus := fields[3]
				if searchDate == taskDate {
					// Calculate time left or overdue
					deadlineStr := taskDate + " " + taskTime
					deadlineTime, _ := time.Parse("02/01/2006 1504", deadlineStr)
					duration := deadlineTime.Sub(now)
					timeLeft := ""
					if duration < 0 {
						timeLeft = "Task Overdue"
					} else {
						days := int(duration.Hours()) / 24
						hours := int(duration.Hours()) % 24
						minutes := int(duration.Minutes()) % 60
						timeLeft = fmt.Sprintf("%d days, %d hours, %d minutes", days, hours, minutes)
					}

					// Store result
					result := fmt.Sprintf("%s, %s, %s, %s, %s", taskName, taskDate, taskTime, taskStatus, timeLeft)
					results = append(results, result)
				}

			}
			if len(results) > 0 {
				feedbackLabel.SetText(strings.Join(results, ","))
			} else {
				feedbackLabel.SetText("No Tasks found for this date")
			}

		}

	})
	centered := container.NewVBox(
		text,
		selectListOptions,
		newValue,
		submitButton,
		feedbackLabel,
		widget.NewButton("Back to Main", func() {
			w.SetContent(createMainScreen(w, a))
		}),
	)
	return container.New(layout.NewCenterLayout(), centered)

}

func main() {
	a := app.New()
	w := a.NewWindow("Task Manager")
	user, err := user.Current()
	if err != nil {
		return
	}
	curruser := user.Username
	// Main screen
	mainScreen := container.NewVBox(
		widget.NewLabel("Welcome, "+curruser),
		widget.NewLabel("What would you like to do today?"),
		widget.NewButton("Add Task", func() {
			w.SetContent(AddTaskScreen(w, a))
		}),
		widget.NewButton("Edit Task", func() {
			w.SetContent(EditTask(w, a))
		}),
		widget.NewButton("List Tasks", func() {
			w.SetContent(ListTask(w, a))
		}),
		widget.NewButton("Exit", func() {
			a.Quit()
		}),
	)
	centered := container.New(layout.NewCenterLayout(), mainScreen)
	w.SetContent(centered)
	w.ShowAndRun()
}
func AddTaskScreen(w fyne.Window, a fyne.App) fyne.CanvasObject {
	// Clear previous feedback
	taskSubmittedLabel := widget.NewLabel("")

	// Task entry fields
	taskNameEntry := widget.NewEntry()
	taskNameEntry.SetPlaceHolder("Enter Task Name here: ")

	taskDeadlineDateEntry := widget.NewEntry()
	taskDeadlineDateEntry.SetPlaceHolder("Enter Deadline date here: ")

	taskDeadlineTimeEntry := widget.NewEntry()
	taskDeadlineTimeEntry.SetPlaceHolder("Enter Deadline time here: ")

	statusOptions := []string{"Ongoing", "Completed", "Incomplete"}
	selectStatusOptions := widget.NewSelect(statusOptions, func(string) {})

	// Button to submit the task
	submitButton := widget.NewButton("Submit", func() {
		taskName := taskNameEntry.Text
		taskDate := taskDeadlineDateEntry.Text
		taskTime := taskDeadlineTimeEntry.Text
		taskStatus := selectStatusOptions.Selected

		if strings.TrimSpace(taskName) == "" || strings.TrimSpace(taskDate) == "" || strings.TrimSpace(taskTime) == "" {
			taskSubmittedLabel.SetText("Couldn't create task, Please fill all the fields")
			return
		}

		task := addTask(taskName, taskDate, taskTime, taskStatus)
		userHome, _ := os.UserHomeDir()
		filepath := userHome + "/Desktop/todoui/todoui.txt"
		file, err := os.OpenFile(filepath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
		if err != nil {
			fmt.Println("Error opening file:", err)
			return
		}
		defer file.Close()
		_, err = file.WriteString(task)
		if err != nil {
			fmt.Println("Error writing to file:", err)
		} else {
			taskSubmittedLabel.SetText("Task added successfully!")
		}

		// Clear the entries for the next input
		taskNameEntry.SetText("")
		taskDeadlineDateEntry.SetText("")
		taskDeadlineTimeEntry.SetText("")
		selectStatusOptions.SetSelected("")
	})

	// Button to go back to the main screen
	backButton := widget.NewButton("Back to Main", func() {
		w.SetContent(createMainScreen(w, a))
	})

	centered := container.NewVBox(
		widget.NewLabel("Enter Name of the Task:"),
		taskNameEntry,
		widget.NewLabel("Enter Date:"),
		taskDeadlineDateEntry,
		widget.NewLabel("Enter Deadline Time: "),
		taskDeadlineTimeEntry,
		widget.NewLabel("Select Task Status:"),
		selectStatusOptions,
		submitButton,
		taskSubmittedLabel,
		backButton,
	)
	return container.New(layout.NewCenterLayout(), centered)
}

func createMainScreen(w fyne.Window, a fyne.App) fyne.CanvasObject {
	user, _ := user.Current()
	curruser := user.Username
	centeredContent := container.NewVBox(
		widget.NewLabel("Welcome Back, "+curruser),
		widget.NewLabel("What would you like to do next?"),
		widget.NewButton("Add Task", func() {
			w.SetContent(AddTaskScreen(w, a))
		}),
		widget.NewButton("Edit Task", func() {
			w.SetContent(EditTask(w, a))
		}),
		widget.NewButton("List Tasks", func() {
			w.SetContent(ListTask(w, a))
		}),
		widget.NewButton("Exit", func() {
			a.Quit()
		}),
	)
	return container.New(layout.NewCenterLayout(), centeredContent)
}
