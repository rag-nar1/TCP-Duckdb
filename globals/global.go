package globals


var PreparedStmtStrings = [][]string{
	{"login",               "SELECT userid , usertype FROM user WHERE username LIKE ? AND password LIKE ? ;"},
	{"SelectUser",          "SELECT userid FROM user WHERE username LIKE ? ;"},
	{"CreateUser",          "INSERT INTO user(username, password, usertype) VALUES(?, ?, ?);"},
	{"UpdateUsername",		"UPDATE user SET username = ? WHERE userid == ?;"},
	{"UpdateUserPassword",	"UPDATE user SET password = ? WHERE userid == ?;"},
	{"SelectDB",            "SELECT dbid FROM DB WHERE dbname LIKE ? ;"},
	{"CreateDB",            "INSERT INTO DB(dbname) VALUES(?);"},
	{"UpdateDB",			"UPDATE DB SET dbname = ? WHERE dbid == ?;"},
	{"GrantDB",             "INSERT OR IGNORE INTO dbprivilege(dbid, userid, privilegetype) VALUES(?, ?, ?);"},
	{"CheckDbAccess",       "SELECT COUNT(*) FROM dbprivilege WHERE userid == ? AND dbid == ?"},
	{"DbAccessType",		"SELECT privilegetype FROM dbprivilege WHERE userid == ? AND dbid == ?;"},
	{"DeleteDbAccess",		"DELETE FROM dbprivilege Where userid == ? AND dbid == ?;"},
	{"SelectTable",         "SELECT tableid FROM tables WHERE tablename LIKE ? AND dbid == ?;"},
	{"CheckTableAccess",    "SELECT COUNT(*) FROM tableprivilege WHERE userid == ? AND tableid == ? AND tableprivilege LIKE ?;"},
	{"GrantTable",          "INSERT OR IGNORE INTO tableprivilege(tableid, userid, tableprivilege) VALUES(?, ?, ?);"},
	{"CreateTable",         "INSERT OR IGNORE INTO tables(tablename, dbid) VALUES(?, ?);"},
	{"CreateLink",          "INSERT OR IGNORE INTO postgres(dbid, connstr) VALUES(?, ?);"},
	{"CheckLink",			"SELECT COUNT(*) FROM postgres WHERE dbid == ?;"},
	{"SelectLink",			"SELECT connstr FROM postgres WHERE dbid == ?;"},
}

var DbPoolSize uint = 1 << 2 // 4 db connections for testing will be 16 for production
var ReplacerK uint	= 1 << 2