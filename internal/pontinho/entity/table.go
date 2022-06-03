package entity

type table struct {
	id      int64
	viewers []player
}

var (
	tableId int64 = 0
)

func NewTable() *table {

	tableId++

	return &table{
		id: tableId,
	}
}
