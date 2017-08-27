package mysqldb

// ContentQuery is a content select query
const (
	InsertContentQuery              = "INSERT INTO content (title, category, created_date, hits, meta_author_name, meta_desc, meta_key_words, meta_robot_key_words, text, sort_order, archived, client_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?) "
	UpdateContentQuery              = "UPDATE content set title = ?, category = ?, modified_date = ?, meta_author_name = ?, meta_desc = ? ,meta_key_words = ?, meta_robot_key_words = ?, text = ?, sort_order = ?, archived = ? where id = ? and client_id = ? "
	UpdateContentHitsQuery          = "UPDATE content set modified_date = ?, hits = ? where id = ? and client_id = ? "
	ContentGetQuery                 = "select * from content WHERE id = ? and client_id = ?"
	ContentGetByClientQuery         = "select * from content WHERE client_id = ? order by category, sort_order, id "
	ContentGetByClientCategoryQuery = "select * from content WHERE client_id = ? and category = ? order by sort_order, id "
	DeleteContentQuery              = "DELETE FROM content WHERE id = ? and client_id = ?"
	ConnectionTestQuery             = "SELECT count(*) from content"
)
