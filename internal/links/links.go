package links

import (
	"log"

	"github.com/tergelm/go_hackernews/internal/pkg/db/migrations/postgresql"
	"github.com/tergelm/go_hackernews/internal/users"
)

// Link model
type Link struct {
	Id      string
	Title   string
	Address string
	User    *users.User
}

// Insert Link func
func (link Link) Save() int64 {
	// SQL query template
	query, err := db_handler.DBClient.Prepare(`
          INSERT INTO links (title, address)
          VALUES($1, $2);
        `)
	if err != nil {
		log.Fatal(err)
	}

	// Run query
	res, err := query.Exec(link.Title, link.Address)
	if err != nil {
		log.Fatal(err)
	}

	// Get created link id
	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal("Error:", err.Error())
	}

	log.Print("Row inserted")
	return id
}

func GetAll() []Link {
	query, err := db_handler.DBClient.Prepare(`
    SELECT id, title, address FROM links
  `)
	if err != nil {
		log.Fatal(err)
	}

	defer query.Close()

	rows, err := query.Query()
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var links []Link
	for rows.Next() {
		var link Link
		err := rows.Scan(&link.Id, &link.Title, &link.Address)
		if err != nil {
			log.Fatal(err)
		}
		links = append(links, link)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return links
}
