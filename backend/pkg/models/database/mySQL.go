package database

import (
	"archive/zip"
	"database/sql"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"sort"
	"time"

	"github.com/avdushin/rgoauth/vars"
	_ "github.com/go-sql-driver/mysql"
)

/* work with MySQL DataBase */

// Make querry...
func DBQuerry(querry, comment string) {
	// DBConn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", DBUser, DBPass, DBHost, DBPort, DBName)

	// Соедененеие с базой данных
	db, err := sql.Open("mysql", vars.DBConn+vars.DBName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close() // закрываем подключение к БД

	// Делаем запрос
	if _, err = db.Exec(fmt.Sprint(querry)); err != nil {
		log.Fatal(err)
	}

	// Выводим лог
	log.Println(comment)
}

// Создаем нужные таблицы
func InitTables() {
	// Create users table
	DBQuerry(`
	CREATE TABLE IF NOT EXISTS users (
		id INT NOT NULL AUTO_INCREMENT,
		username VARCHAR(50) NOT NULL DEFAULT '',
		email VARCHAR(50) NOT NULL DEFAULT '',
		password VARCHAR(255) NOT NULL DEFAULT '',
		name VARCHAR(255) DEFAULT '',
		surname VARCHAR(255) DEFAULT '',
		patronymic VARCHAR(255) DEFAULT '',
		age INT UNSIGNED DEFAULT 0,
		PRIMARY KEY (id),
		UNIQUE KEY (username),
		UNIQUE KEY (email)
	  );`, "[DB]: users table is ready...")

	log.Println("[DB]: Таблицы созданы...")
}

// Создаём бэкап БД
func BackupDB() {
	// Получаем текущее время
	now := time.Now().Local()
	DateFormat := now.Format("2006-01-02_15-04-05")

	// Создаём папку для хранения бэкапов если её ещё не существует
	if _, err := os.Stat("./backups"); os.IsNotExist(err) {
		os.Mkdir("./backups", os.ModePerm)
	}
	// Создаём папку для хранения архивов бэкапов если её ещё не существует
	if _, err := os.Stat("./backups/archives"); os.IsNotExist(err) {
		os.Mkdir("./backups/archives", os.ModePerm)
	}

	// Имя файла для бэкапа и zip архива
	backupFileName := fmt.Sprintf("./backups/%s.sql", DateFormat)
	// zipFileName := fmt.Sprintf("./backups/archives/%s.zip", DateFormat)

	// Создаём бэкап
	backupFile, err := os.Create(backupFileName)
	if err != nil {
		log.Fatalf("Ошибка создания бэкап файла БД: %v", err)
	}
	defer backupFile.Close()

	// mysqldump command
	cmdArgs := []string{
		"-u",
		vars.DBUser,
		"-p" + vars.DBPass,
		"-h",
		vars.DBHost,
		vars.DBName,
	}

	// Бэкапим БД с помощью mysqldump
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		// Windows
		cmd = exec.Command("C:\\Program Files\\MySQL\\MySQL Server 8.0\\bin\\mysqldump.exe", cmdArgs...)
	default:
		// Unix
		cmd = exec.Command("mysqldump", cmdArgs...)
	}

	outfile, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatalf("Error creating mysqldump command: %v", err)
	}

	// Запускаем mysqldump command
	if err := cmd.Start(); err != nil {
		log.Fatalf("Неудалось запустить mysqldump: %v", err)
	}

	// Записываем вывод mysqldump в бэкап файл
	if _, err := io.Copy(backupFile, outfile); err != nil {
		log.Fatalf("Ошибка записи бэкап файла: %v", err)
	}

	// Ждём выполнения mysqldump
	if err := cmd.Wait(); err != nil {
		log.Fatalf("[Ошибка] Я устал ждать пока mysqldump соберет там свои дампы: %v", err)
	}

	log.Printf("[БД]: Резервная копия создана и сохранена в %s\n", backupFileName)
}

// Функция для добавления файла в архив
func addFileToZip(filePath string, zipWriter *zip.Writer) error {
	// Открываем файл
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	// Получаем информацию о файле
	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}

	// Создаем новый файл в архиве
	zipFile, err := zipWriter.Create(fileInfo.Name())
	if err != nil {
		return err
	}

	// Копируем содержимое файла в архив
	_, err = io.Copy(zipFile, file)
	if err != nil {
		return err
	}

	return nil
}

func SaveStorage() {
	// Получаем текущее время
	now := time.Now().Local()
	DateFormat := now.Format("2006-01-02_15-04-05")
	zipFileName := fmt.Sprintf("./backups/archives/%s.zip", DateFormat)

	// Получаем список всех файлов в папке с расширением .sql
	fileList, err := filepath.Glob(filepath.Join("./backups", "*.sql"))
	if err != nil {
		log.Fatalf("Ошибка получения списка файлов: %v", err)
	}

	// Проверяем, есть ли больше 5 файлов в папке
	if len(fileList) > 5 {
		// Сортируем список файлов по времени изменения (последний файл будет первым в списке)
		sort.Slice(fileList, func(i, j int) bool {
			fileStatI, err := os.Stat(fileList[i])
			if err != nil {
				log.Fatalf("Ошибка получения информации о файле %s: %v", fileList[i], err)
			}
			fileStatJ, err := os.Stat(fileList[j])
			if err != nil {
				log.Fatalf("Ошибка получения информации о файле %s: %v", fileList[j], err)
			}
			return fileStatI.ModTime().After(fileStatJ.ModTime())
		})

		// Создаем новый архив
		// zipFileName := fmt.Sprintf("%s.zip", fileList[0])
		zipFile, err := os.Create(zipFileName)
		if err != nil {
			log.Fatalf("Ошибка создания zip файла %s: %v", zipFileName, err)
		}
		defer zipFile.Close()

		zipWriter := zip.NewWriter(zipFile)

		// Добавляем последний файл в архив
		err = addFileToZip(fileList[0], zipWriter)
		if err != nil {
			log.Fatalf("Ошибка добавления файла %s в zip архив: %v", fileList[0], err)
		}

		// Закрываем архив
		err = zipWriter.Close()
		if err != nil {
			log.Fatalf("Ошибка закрытия zip архива: %v", err)
		}

		// Удаляем все файлы .sql
		for _, file := range fileList[1:] {
			err = os.Remove(file)
			if err != nil {
				log.Fatalf("Ошибка удаления файла %s: %v", file, err)
			}
		}

		fmt.Printf("Файл %s успешно заархивирован и удалены все файлы .sql в папке!\n", fileList[0])
	}
}
