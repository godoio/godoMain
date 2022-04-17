package taskUtils

import (
	databaseService "github.com/godoio/godoMain/databaseService"
)

func DeleteTask(taskId int64) error {
	db, err := databaseService.DatabaseConnect()
	if err != nil {
		return err
	}
	errDel := databaseService.DeleteTask(taskId, db)
	if errDel != nil {
		return errDel
	}
	return nil
}