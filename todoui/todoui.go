package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

func writeTask() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter Name of the task: ")
	taskName, _ := reader.ReadString('\n')
	taskName = strings.TrimSpace(taskName)

	fmt.Println("Enter deadline date for the task: ")
	taskDate, _ := reader.ReadString('\n')
	taskDate = strings.TrimSpace(taskDate)

	fmt.Println("Enter deadline time for the task: ")
	taskTime, _ := reader.ReadString('\n')
	taskTime = strings.TrimSpace(taskTime)

	fmt.Println("Enter status of the task: ")
	taskStatus, _ := reader.ReadString('\n')
	taskStatus = strings.TrimSpace(taskStatus)
	if taskStatus == "" {
		taskStatus = "Ongoing"
	}

	creationTime := time.Now().Format("02/01/2006 1504")

	task := fmt.Sprintf("%s,%s,%s,%s,%s\n", taskName, taskDate, taskTime, taskStatus, creationTime)
	return task
}

func listTasks() {
	var choice int
	userHome, _ := os.UserHomeDir()
	filepath := userHome + "/Desktop/todoui/todoui.txt"
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	inputreader := bufio.NewReader(os.Stdin)
	reader := bufio.NewReader(file)
	fmt.Println("How would you like to list the tasks?(1: All, 2: Date, 3: Status)")
	fmt.Scanln(&choice)
	now := time.Now()
	var timeLeft string
	if choice == 1 {
		for {
			line, err := reader.ReadString('\n')
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
			if duration < 0 {
				timeLeft = "Task overdue"
			} else {
				days := int(duration.Hours()) / 24
				hours := int(duration.Hours()) % 24
				minutes := int(duration.Minutes()) % 60
				timeLeft = fmt.Sprintf("%d days, %d hours, %d minutes", days, hours, minutes)
			}
			fmt.Printf("%s, %s, %s, %s, %s\n", taskName, taskDate, taskTime, taskStatus, timeLeft)
		}
	}
	if choice == 2 {
		fmt.Println("Enter date in DD/MM/YYYY format")
		date, _ := inputreader.ReadString('\n')
		date = strings.TrimSpace(date)
		for {
			line, err := reader.ReadString('\n')
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
				fmt.Printf("Error: Invalid task format or missing fields in line: %s", line)
				continue
			}
			taskName := fields[0]
			taskDate := fields[1]
			taskTime := fields[2]
			taskStatus := fields[3]

			if date == taskDate {
				deadlineStr := taskDate + " " + taskTime
				deadlineTime, _ := time.Parse("02/01/2006 1504", deadlineStr)
				duration := deadlineTime.Sub(now)
				timeLeft := "Time left to complete: " + duration.String()
				if duration < 0 {
					timeLeft = "Task Overdue"
				} else {
					days := int(duration.Hours()) / 24
					hours := int(duration.Hours()) % 24
					minutes := int(duration.Minutes()) % 60
					timeLeft = fmt.Sprintf("%d days, %d hours, %d minutes", days, hours, minutes)
				}
				fmt.Printf("%s, %s, %s, %s, %s\n", taskName, taskDate, taskTime, taskStatus, timeLeft)
			}
		}
	}
	if choice == 3 {
		fmt.Println("Enter Ongoing, Completed or Incomplete")
		status, _ := inputreader.ReadString('\n')
		status = strings.TrimSpace(status)
		for {
			line, err := reader.ReadString('\n')
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
			if status == taskStatus {
				deadlineStr := taskDate + " " + taskTime
				deadlineTime, err := time.Parse("02/01/2006 1504", deadlineStr)
				if err != nil {
					fmt.Println("error parsing time:", err)
				}
				duration := deadlineTime.Sub(now)
				timeLeft := "Time left to complete: " + duration.String()
				if duration < 0 {
					timeLeft = "Task overdue"
				} else {
					days := int(duration.Hours()) / 24
					hours := int(duration.Hours()) % 24
					minutes := int(duration.Minutes()) % 60
					timeLeft = fmt.Sprintf("%d days, %d hours, %d minutes", days, hours, minutes)
				}
				fmt.Printf("%s, %s, %s, %s, %s\n", taskName, taskDate, taskTime, taskStatus, timeLeft)
			}
		}
	}
}

func editTask() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter name of task to edit")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)
	userHome, _ := os.UserHomeDir()
	filepath := userHome + "/Desktop/todoui/todoui.txt"
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var tasks []string
	filereader := bufio.NewReader(file)
	found := false
	anyEdited := false
	for {
		line, err := filereader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("error:", err)
			return
		}
		line = strings.TrimSpace(line)
		fields := strings.Split(line, ",")
		if name == fields[0] {
			found = true
			var num int
			fmt.Println("Enter field to edit: 1. Name, 2. Status, 3. Time, 4. Date")
			fmt.Scanln(&num)
			if num == 1 {
				fmt.Println("Enter new name for the task: ")
				newName, _ := reader.ReadString('\n')
				newName = strings.TrimSpace(newName)
				if newName == "" {
					newName = fields[0]
				}
				fields[0] = newName
				anyEdited = true
			} else if num == 2 {
				fmt.Println("Enter new status for the task: ")
				newStatus, _ := reader.ReadString('\n')
				newStatus = strings.TrimSpace(newStatus)
				fields[3] = newStatus
				anyEdited = true
			} else if num == 3 {
				fmt.Println("Enter new deadline time in HHMM(24 HOURS) format.")
				newTime, _ := reader.ReadString('\n')
				newTime = strings.TrimSpace(newTime)
				fields[2] = newTime
				anyEdited = true
			} else if num == 4 {
				fmt.Println("Enter new date in DD/MM/YYYY format")
				newDate, _ := reader.ReadString('\n')
				newDate = strings.TrimSpace(newDate)
				fields[1] = newDate
				anyEdited = true
			} else {
				tasks = append(tasks, line)
				continue
			}

			editedTask := strings.Join(fields, ",")
			tasks = append(tasks, editedTask)
		} else {
			tasks = append(tasks, line)
		}
	}
	if !found {
		fmt.Println("Task name not found. Create a new task or enter a valid task name.")
		return
	}

	if anyEdited {
		file, err := os.Create(filepath)
		if err != nil {
			fmt.Println("Error creating file:", err)
			return
		}
		defer file.Close()
		for _, task := range tasks {
			_, err := file.WriteString(task + "\n")
			if err != nil {
				fmt.Println("Error writing to file:", err)
				return
			}
		}
		fmt.Println("Task edited successfully!")
	} else {
		fmt.Println("No changes were made.")
	}
}

func main() {
	fmt.Println("Task management CLI")
	for {
		var choice int
		fmt.Println("1: Add Task, 2: List Tasks, 3: Edit Task, 4: Exit")
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			task := writeTask()
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
				fmt.Println("Task added successfully!")
			}
		case 2:
			listTasks()
		case 3:
			editTask()
		case 4:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}
