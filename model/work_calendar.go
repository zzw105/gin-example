package model

import (
	"database/sql"
	"fmt"
	"gin-example/db"
	"gin-example/types"
)

func CreateWorkCalendar(workCalendar types.WorkCalendar) (int64, error) {
	query := "INSERT INTO work_calendar (day, message) VALUES (?, ?)"
	result, err := db.DB.Exec(query, workCalendar.Day, workCalendar.Message)
	if err != nil {
		return 0, err
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return lastID, nil
}

func GetAllWorkCalendar() ([]types.WorkCalendar, error) {
	rows, err := db.DB.Query("SELECT * FROM work_calendar")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var workCalendars []types.WorkCalendar
	for rows.Next() {
		var workCalendar types.WorkCalendar
		if err := rows.Scan(&workCalendar.ID, &workCalendar.Day, &workCalendar.Message); err != nil {
			return nil, err
		}
		workCalendars = append(workCalendars, workCalendar)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return workCalendars, nil
}

func UpdateWorkCalendar(workCalendar types.WorkCalendar) (int64, error) {
	// 先查询数据库，检查该用户是否存在
	var existingWorkCalendar types.WorkCalendar
	err := db.DB.QueryRow("SELECT id, day, message FROM work_calendar WHERE day = ?", workCalendar.Day).Scan(&existingWorkCalendar.ID, &existingWorkCalendar.Day, &existingWorkCalendar.Message)
	if err != nil {
		if err == sql.ErrNoRows {
			// 用户不存在
			return 0, fmt.Errorf("user with id %s not found", workCalendar.Day)
		}
		// 其他查询错误
		return 0, fmt.Errorf("failed to check user existence: %v", err)
	}

	// 用户存在，执行更新操作
	query := "UPDATE work_calendar SET  message = ? WHERE day = ?"
	result, err := db.DB.Exec(query, workCalendar.Message, workCalendar.Day)
	if err != nil {
		return 0, fmt.Errorf("update user failed: %v", err)
	}

	// 获取更新操作影响的行数
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("getting affected rows failed: %v", err)
	}

	// 返回影响的行数
	return rowsAffected, nil
}
