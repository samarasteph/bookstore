package queries

import "time"

// Book type for SQL table BOOKS
/**
+----------+------------------+------+-----+---------+----------------+
| Field    | Type             | Null | Key | Default | Extra          |
+----------+------------------+------+-----+---------+----------------+
| ID       | int(10) unsigned | NO   | PRI | NULL    | auto_increment |
| TITLE    | text             | NO   |     | NULL    |                |
| YEAR     | year(4)          | NO   |     | NULL    |                |
| DESCR    | text             | YES  |     | NULL    |                |
| AUTHORS  | text             | NO   |     | NULL    |                |
| SIZE     | int(11)          | NO   |     | NULL    |                |
| HASH     | text             | YES  |     | NULL    |                |
| IMG_PATH | text             | YES  |     | NULL    |                |
+----------+------------------+------+-----+---------+----------------+
*/
type Book struct {
	ID      uint32 `json:"ID"`
	Title   string `json:"TITLE"`
	Year    uint32 `json:"YEAR"`
	Descr   string `json:"DESCR"`
	Authors string `json:"AUTHORS"`
	Size    uint32 `json:"SIZE"`
	ImgPath string `json:"IMG_PATH"`
}

// BookLink type for SQL table BOOKS_LINKS
/*
+-----------+--------------+------+-----+---------+----------------+
| Field     | Type         | Null | Key | Default | Extra          |
+-----------+--------------+------+-----+---------+----------------+
| BOOK_ID   | int(11)      | NO   | MUL | NULL    |                |
| STORE_ID  | int(11)      | NO   |     | NULL    |                |
| ID        | int(11)      | NO   | PRI | NULL    | auto_increment |
| FILE_ID   | text         | NO   |     | NULL    |                |
| FILE_SIZE | int(11)      | NO   |     | 0       |                |
| FILE_NAME | varchar(128) | NO   |     | NULL    |                |
+-----------+--------------+------+-----+---------+----------------+
*/
type BookLink struct {
	ID       uint32 `json:"ID"`
	BookID   uint32 `json:"BOOK_ID"`
	StoreID  uint32 `json:"STORE_ID"`
	FileID   string `json:"FILE_ID"`
	FileSize uint64 `json:"FILE_SIZE"`
	FileName string `json:"FILE_NAME"`
}

// ItSubject type for SQL Table IT_SUBJECT
/*
+-----------+------------+------+-----+---------+----------------+
| Field     | Type       | Null | Key | Default | Extra          |
+-----------+------------+------+-----+---------+----------------+
| ID        | int(11)    | NO   | PRI | NULL    | auto_increment |
| NAME      | text       | NO   |     | NULL    |                |
| DESCR     | text       | YES  |     | NULL    |                |
| PERMANENT | tinyint(1) | NO   |     | NULL    |                |
+-----------+------------+------+-----+---------+----------------+
*/
type ItSubject struct {
	ID        uint32 `json:"ID"`
	Name      string `json:"NAME"`
	Descr     string `json:"DESCR"`
	Permanent bool   `json:"PERMANENT"`
}

// FileStore type for SQL Table FILE_STORE
/*
+------------------+---------------+------+-----+---------+----------------+
| Field            | Type          | Null | Key | Default | Extra          |
+------------------+---------------+------+-----+---------+----------------+
| ID               | int(11)       | NO   | PRI | NULL    | auto_increment |
| VENDOR           | text          | NO   |     | NULL    |                |
| VENDOR_CODE      | varchar(4)    | NO   | UNI | NULL    |                |
| STORAGE_CAPACITY | int(11)       | NO   |     | NULL    |                |
| LOGIN_INFO       | varchar(4096) | YES  |     | NULL    |                |
+------------------+---------------+------+-----+---------+----------------+
*/
type FileStore struct {
	ID         uint32 `json:"ID"`
	Vendor     string `json:"VENDOR"`
	VendorCode string `json:"VENDOR_CODE"`
	Capacity   uint64 `json:"STORAGE_CAPACITY"`
	LoginInfo  string `json:"LOGIN_INFO"`
}

// BookSubjectAssoc for SQL Table 	BOOKS_SUBJECTS_ASSOC
/*
+------------+---------+------+-----+---------+-------+
| Field      | Type    | Null | Key | Default | Extra |
+------------+---------+------+-----+---------+-------+
| SUBJECT_ID | int(11) | NO   | MUL | NULL    |       |
| BOOK_ID    | int(11) | NO   |     | NULL    |       |
+------------+---------+------+-----+---------+-------+
*/
type BookSubjectAssoc struct {
	SubjectID uint32 `json:"SUBJECT_ID"`
	BookID    uint32 `json:"BOOK_ID"`
}

// BookSubjectAssocList list of subjects associated with BOOK
type BookSubjectAssocList struct {
	SubjectID []uint32 `json:"SUBJECT_IDS"`
	BookID    uint32   `json:"BOOK_ID"`
}

// Logs for SQL Table LOGS
/*
+----------+---------------+------+-----+---------------------+----------------+
| Field    | Type          | Null | Key | Default             | Extra          |
+----------+---------------+------+-----+---------------------+----------------+
| ID       | int(11)       | NO   | PRI | NULL                | auto_increment |
| LEVEL    | varchar(10)   | NO   |     | NULL                |                |
| TIME     | datetime      | NO   |     | current_timestamp() |                |
| FILE     | varchar(512)  | NO   |     | NULL                |                |
| LINE     | varchar(15)   | NO   |     | NULL                |                |
| FUNCTION | varchar(128)  | NO   |     | NULL                |                |
| MSG      | varchar(4096) | NO   |     | NULL                |                |
+----------+---------------+------+-----+---------------------+----------------+
*/
type Logs struct {
	ID       uint32    `json:"ID"`
	Level    string    `json:"LEVEL"`
	Time     time.Time `json:"TIME"`
	File     string    `json:"FILE"`
	Line     string    `json:"LINE"`
	Function string    `json:"FUNCTION"`
	Msg      string    `json:"MSG"`
}

// BookTitle book item in Books page
type BookTitle struct {
	BookID  uint32 `json:"BOOK_ID"`
	Title   string `json:"TITLE"`
	ImgPath string `json:"IMG_PATH"`
}

// BooksBrowserPage presentation of BookTitle list (images and title)
type BooksBrowserPage struct {
	BooksList []BookTitle `json:"BOOK_LIST"`
}

//RestError for returning json error message
type RestError struct {
	Msg     string `json:"msg"`
	Context string `json:"context"`
}
