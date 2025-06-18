package model

import(
	"web/utils"
	"database/sql"
	"strings"
	"log"
	"time"
)

const (
	FileTable = "file"
	qInsertFileMeta = `INSERT INTO file (name, location, size) VALUES ($1, $2, $3) RETURNING id,uploaded_at`
	qGetFileMeta = `SELECT DISTINCT id, name, location, uploaded_at, status, size FROM file`

	
)
type FileMeta struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	Location string `json:"location"`
	UploadedAt string `json:"uploaded_at"`
	Status int64 `json:"status"`
	Size int64 `json:"size"`
}

func InsertFileMeta(f FileMeta)(FileMeta,error){
	
	Stmt, err := utils.GetDB().Prepare(qInsertFileMeta)
	if err != nil {
		log.Println("Error preparing statement:", err.Error())
		return f,err
}
	defer Stmt.Close()
	var uploadedAt time.Time
	err = Stmt.QueryRow(f.Name, f.Location, f.Size).Scan(&f.Id, &uploadedAt)
	if err != nil {
		log.Println("Error inserting file metadata:", err.Error())
		return f,err
	}
	f.UploadedAt = uploadedAt.Format("2006-01-02 15:04:05")
	return f, nil
}

func GetFileMetas(search string, page, limit int) ([]FileMeta, error) {
	search = strings.TrimSpace(search)
	queryLimit := limit
	queryOffset := (page - 1) * limit

	var queryStmt= qGetFileMeta
	var err error
	var rows *sql.Rows

	// Build dynamic WHERE clause
	if search != "" {
		queryStmt += " WHERE name ~* $3"
	}

	queryStmt += " ORDER BY uploaded_at DESC LIMIT $1 OFFSET $2"
	stmt, err := utils.GetDB().Prepare(queryStmt)
	if err != nil {
		log.Println("Error preparing statement:", err.Error())
		return nil, err
	}

	defer stmt.Close()

	if search != "" {
		rows, err = stmt.Query(queryLimit, queryOffset, search)
	} else {
		rows, err = stmt.Query(queryLimit, queryOffset)
	}

	if err != nil {
		log.Println("Error executing query:", err.Error())
		return nil, err
	}

	defer rows.Close()

	fileMetas := make([]FileMeta, 0)
	for rows.Next() {
		var fm FileMeta
		err = rows.Scan(&fm.Id, &fm.Name, &fm.Location, &fm.UploadedAt, &fm.Status, &fm.Size)
		if err != nil {
			log.Println("Error scanning row:", err.Error())
			return nil, err
		}
		fileMetas = append(fileMetas, fm)
	}
	return fileMetas, nil
}
