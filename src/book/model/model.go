package model

var Table = "book"
var TableStruct = 
		"CREATE TABLE IF NOT EXISTS `book`(" +
			"id int(11) UNSIGNED NOT NULL AUTO_INCREMENT," +
			"hash VARBINARY(20) DEFAULT NULL," +
			"fname VARCHAR(255) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL," +
			"PRIMARY KEY(id)" +
		") ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;" 
