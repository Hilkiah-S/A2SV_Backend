package database

var nextID int = 1

type TaskModel struct {
	ID          int `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	DueDate     string `json:"dueDate"`
	Status      bool `json:"status"`
}

var table []TaskModel

func deleteAtIndex(s []TaskModel, index int) []TaskModel {
	return append(s[:index], s[index+1:]...)
}

func GetAllTasks() []TaskModel {
	return table
}

func GetTaskByID(id int) (TaskModel, bool) {
	for _, task := range table {
		if task.ID == id {
			return task, true
		}
	}
	return TaskModel{}, false
}

func CreateTask(newTask TaskModel) TaskModel {
	newTask.ID = nextID
	nextID++
	table = append(table, newTask)
	return newTask
}

func UpdateTask(id int, updatedDetails TaskModel) (TaskModel, bool) {
	for i, task := range table {
		if task.ID == id {
			updatedDetails.ID = id
			table[i] = updatedDetails
			return updatedDetails, true
		}
	}
	return TaskModel{}, false
}

func DeleteTask(id int) bool {
	for i := 0; i < len(table); i++ {
		if table[i].ID == id {
			table = deleteAtIndex(table, i)
			return true
		}
	}
	return false
}
