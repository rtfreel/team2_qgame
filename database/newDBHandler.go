package database

const (
	//Driver is a name of db driver (requires go get github.com/mattn/go-sqlite3)
	Driver = "sqlite3"
	//Path is a path to file with database
	Path = "database/CandyWarGoDatabase.sqlite"
)

//NewDBHandler returns pointer to the default ready to use DBHandler
func NewDBHandler() *DBHandler {
	dbh := &DBHandler{
		DriverName: Driver,
		DBPath:     Path,
	}
	dbh.Connect()
	dbh.CreateUsersTable()
	return dbh
}
