package entity

// Table stores subscribers and game
type Table struct {
	id      int64
	viewers []player
}

var (
	tableID int64 = 0
)

// NewTable returns a new table
func NewTable() *Table {

	tableID++

	return &Table{
		id: tableID,
	}
}
