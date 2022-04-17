package databaseUtils

import (
	"database/sql"
	"fmt"
)

func DeleteTask(taskId int64, db *sql.DB) error {
	result, err := db.Exec("DELETE from tasks WHERE task_id=?",taskId)
	if err != nil {
        return fmt.Errorf("Add task: %v", err)
    }
	affect, err := result.RowsAffected()
	if err != nil {
        return fmt.Errorf("Add task: %v", err)
    }
	fmt.Println("Records affected", affect)
	return nil
}