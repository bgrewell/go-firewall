package _old

import "fmt"

// Table represents the default tables
type Table string

var (
	tables = []string{"raw", "filter", "nat", "mangle", "security"}
)

const (
	TableRaw      Table = "raw"
	TableFilter   Table = "filter"
	TableNat      Table = "nat"
	TableMangle   Table = "mangle"
	TableSecurity Table = "security"
)

func ConvertToTable(tableName string) (table Table, err error) {
	for _, t := range tables {
		if t == tableName {
			return Table(tableName), nil
		}
	}
	return "", fmt.Errorf("table name %s not found", tableName)
}
