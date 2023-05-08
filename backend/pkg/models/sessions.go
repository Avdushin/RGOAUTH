package models

/* legacy:  Working with sessions */

// Функция для создания новой сессии
// func CreateSession(db *sql.DB, userID int64) (string, error) {
// 	// Проверяем, существует ли сессия для данного пользователя
// 	var sessionID string
// 	err := db.QueryRow("SELECT session_id FROM sessions WHERE user_id = ?", userID).Scan(&sessionID)
// 	if err == nil {
// 		// Сессия уже существует, возвращаем ее идентификатор
// 		return sessionID, nil
// 	}

// 	// Генерируем идентификатор сессии
// 	sessionID = generateSessionID(32)

// 	// Создаем новую сессию в базе данных
// 	stmt, err := db.Prepare("INSERT INTO sessions(session_id, user_id) VALUES(?, ?)")
// 	if err != nil {
// 		return "", err
// 	}
// 	defer stmt.Close()

// 	_, err = stmt.Exec(sessionID, userID)
// 	if err != nil {
// 		return "", err
// 	}

// 	return sessionID, nil
// }

// // Функция для генерации случайной строки заданной длины
// func generateSessionID(length int) string {
// 	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
// 	b := make([]byte, length)
// 	for i := range b {
// 		b[i] = charset[rand.Intn(len(charset))]
// 	}
// 	return string(b)
// }
