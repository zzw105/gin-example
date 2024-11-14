package model

import (
	"database/sql"
	"fmt"
	"gin-example/db"
	"gin-example/types"
)

// CreateUser inserts a new user into the database
func CreateUser(user types.User) (int64, error) {
	query := "INSERT INTO users (username, password, email) VALUES (?, ?, ?)"
	result, err := db.DB.Exec(query, user.Username, user.Password, user.Email)
	if err != nil {
		return 0, err
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return lastID, nil
}

// GetAllUsers retrieves all users from the database
func GetAllUsers() ([]types.User, error) {
	rows, err := db.DB.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []types.User
	for rows.Next() {
		var user types.User
		if err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Password); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func UpdateUser(id int, user types.User) (int64, error) {
	// 先查询数据库，检查该用户是否存在
	var existingUser types.User
	err := db.DB.QueryRow("SELECT id, username, password, email FROM users WHERE id = ?", id).Scan(&existingUser.ID, &existingUser.Username, &existingUser.Password, &existingUser.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			// 用户不存在
			return 0, fmt.Errorf("user with id %d not found", id)
		}
		// 其他查询错误
		return 0, fmt.Errorf("failed to check user existence: %v", err)
	}

	// 用户存在，执行更新操作
	query := "UPDATE users SET username = ?, password = ?, email = ? WHERE id = ?"
	result, err := db.DB.Exec(query, user.Username, user.Password, user.Email, id)
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
