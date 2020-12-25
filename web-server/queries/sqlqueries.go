package queries

//use guest@aldu to connect to database
import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var DBConnectionStr string

func openDB() (*sql.DB, error) {
	//DBConnectionStr patemeter to connect to database (temporary)
	return sql.Open("mysql", DBConnectionStr)
}

//SQLQueryBook extract Book from database and return JSON serializable Book object
func SQLQueryBook(id uint64) (*Book, error) {

	var book Book

	db, err := openDB()
	if err != nil {
		fmt.Println("Cannot open database")
		return nil, err
	}
	sql, err := db.Begin()

	if err != nil {
		return nil, err
	}
	err = sql.QueryRow("SELECT ID,TITLE,YEAR,AUTHORS,DESCR,SIZE,IMG_PATH FROM BOOKS WHERE ID=?", id).Scan(&book.ID, &book.Title, &book.Year, &book.Descr, &book.Authors, &book.Size, &book.ImgPath)
	if err != nil {
		//something go wrong with database
		return nil, err
	}

	defer db.Close()

	return &book, err
}

//SQLQueryBookLink extract Book from database and return JSON serializable Book object
func SQLQueryBookLink(bookid uint64) (*BookLink, error) {

	var booklink BookLink

	db, err := openDB()
	if err != nil {
		return nil, err
	}
	sql, err := db.Begin()

	if err != nil {
		return nil, err
	}

	defer db.Close()

	err = sql.QueryRow("SELECT BOOK_ID,STORE_ID,FILE_ID,FILE_SIZE,FILE_NAME FROM BOOKS_LINKS WHERE BOOK_ID=?", bookid).Scan(&booklink.BookID,
		&booklink.StoreID, &booklink.FileID, &booklink.FileSize, &booklink.FileName)
	if err != nil {
		return nil, err
	}

	return &booklink, nil
}

//SQLQueryBookSubjects return subjects list associated with one book
func SQLQueryBookSubjects(bookid uint64) (*BookSubjectAssocList, error) {

	var subjects BookSubjectAssocList
	db, err := openDB()

	if err != nil {
		return nil, err
	}
	sql, err := db.Begin()

	if err != nil {
		return nil, err
	}

	defer db.Close()

	results, err := sql.Query("SELECT BOOK_ID,SUBJECT_ID FROM BOOKS_SUBJECTS_ASSOC WHERE BOOK_ID=?", bookid)

	if err != nil {
		return nil, err
	}

	defer results.Close()

	var subject BookSubjectAssoc
	for results.Next() {
		results.Scan(&subject.BookID, &subject.SubjectID)
		subjects.SubjectID = append(subjects.SubjectID, subject.SubjectID)
	}

	subjects.BookID = subject.BookID

	if err != nil {
		return nil, err
	}
	return &subjects, nil
}

//SQLQueryPage send page data to display: book title, img_path etc...
func SQLQueryPage(pageNumber uint64, nbItems uint64) (*[]BookTitle, error) {
	var pages []BookTitle

	db, err := openDB()
	if err != nil {
		return nil, err
	}

	sql, err := db.Begin()

	defer db.Close()

	if err != nil {
		return nil, err
	}

	results, err := sql.Query("SELECT ID,TITLE,IMG_PATH FROM BOOKS ORDER BY ID LIMIT ? OFFSET ?", nbItems, nbItems*pageNumber)

	if err != nil {
		return nil, err
	}

	defer results.Close()

	for results.Next() {
		var bookTitle BookTitle
		results.Scan(&bookTitle.BookID, &bookTitle.Title, &bookTitle.ImgPath)
		pages = append(pages, bookTitle)
	}
	return &pages, nil
}

//SQLQueryDeleteBook delete book and associated database data
func SQLQueryDeleteBook(bookID uint64) error {

	db, err := openDB()
	if err != nil {
		return err
	}

	defer db.Close()

	var ctx context.Context
	tx, err := db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})

	_, err = tx.Exec("DELETE FROM BOOKS WHERE ID=?", bookID)
	if err == nil {
		_, err = tx.Exec("DELETE FROM BOOKS_LINKS WHERE BOOK_ID=?", bookID)
	}
	if err == nil {
		_, err = tx.Exec("DELETE FROM BOOKS_SUBJECTS_ASSOC WHERE BOOK_ID=?", bookID)
	}

	if err == nil {
		tx.Commit()
	} else {
		tx.Rollback()
	}
	return err
}

//SQLQuerySubjects returns subjects list
func SQLQuerySubjects() (*[]ItSubject, error) {
	db, err := openDB()
	if err != nil {
		return nil, err
	}

	defer db.Close()

	subjects := make([]ItSubject, 0)
	results, err := db.Query("SELECT ID, NAME FROM IT_SUBJECT")

	if err != nil {
		return nil, err
	}

	defer results.Close()

	for results.Next() {
		var subject ItSubject
		err = results.Scan(&subject.ID, &subject.Name)
		if err != nil {
			return nil, err
		}
		subjects = append(subjects, subject)
	}

	return &subjects, nil
}
