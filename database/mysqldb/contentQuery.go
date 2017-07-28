package mysqldb

// ContentQuery is a content select query
const (
	InsertContentQuery      = "INSERT INTO content (title, created_date, text, client_id) VALUES (?, ?, ?, ?) "
	UpdateContentQuery      = "UPDATE content set title = ?, modified_date = ?, text = ? where id = ? and client_id = ? "
	ContentGetQuery         = "select * from content WHERE id = ? and client_id = ?"
	ContentGetByClientQuery = "select * from content WHERE client_id = ? order by id"
	DeleteContentQuery      = "DELETE FROM content WHERE id = ? and client_id = ?"
)
