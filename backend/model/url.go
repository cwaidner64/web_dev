
package model

import(
	"web/utils"
	"database/sql"
	"log"
	"strings"

)

const (
    URLTable = "url"
    qInsertURLMeta = `INSERT INTO url (name, parent_url, depth) VALUES ($1, $2, $3) RETURNING id,type,search_engine,status,date`
    qGetURLMeta = `SELECT DISTINCT id, name, parent_url, depth, search_engine, type, status, date FROM url`
)


type URLMeta struct {
    Id           int64  `json:"id"`
    Name         string `json:"name"`
    Parent_url   string `json:"parent_url"`
    Depth        int64  `json:"depth"`
    Type_         int64  `json:"type"`
    SearchEngine int64  `json:"search_engine"`
    Status       int64  `json:"status"`
    CreatedAt    string `json:"created_at"` // or time.Time if you want
}

func InsertURLMeta(f URLMeta)(URLMeta,error){
	Stmt, err := utils.GetDB().Prepare(qInsertURLMeta)
	if err != nil {
		log.Println("Error preparing statement:", err.Error())
		return f, err
	}
	defer Stmt.Close()
	err = Stmt.QueryRow(f.Name, f.Parent_url, f.Depth).Scan(
		&f.Id, &f.Type_, &f.SearchEngine, &f.Status, &f.CreatedAt,
	)

	if err != nil {
		log.Println("Error inserting URL metadata:", err.Error())
		return f, err
	}
	return f, nil
}


func GetURLMeta(search string, page, limit int) ([]URLMeta, error) {
	search = strings.TrimSpace(search)
	queryLimit := limit
	queryOffset := (page - 1) * limit

	var queryStmt = qGetURLMeta
	var err error
	var rows *sql.Rows

	// Build dynamic WHERE clause
	if search != "" {
		queryStmt += " WHERE name ~* $3"
	}

	queryStmt += " ORDER BY id DESC LIMIT $1 OFFSET $2"
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
	urlMetas := make([]URLMeta, 0)
	for rows.Next() {
		var um URLMeta
		err = rows.Scan(
			&um.Id, &um.Name, &um.Parent_url, &um.Depth,
			&um.SearchEngine, &um.Type_, &um.Status, &um.CreatedAt,
		)
		if err != nil {
			log.Println("Error scanning row:", err.Error())
			return nil, err
		}
		urlMetas = append(urlMetas, um)
	}
	return urlMetas, nil
}


