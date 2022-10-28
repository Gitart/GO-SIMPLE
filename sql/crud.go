package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/user"
	"path"
	"strconv"
	"strings"

	"github.com/gorilla/mux"

	_ "github.com/lib/pq"
)

type file struct {
	fileID         int
	offset         *int
	uploadLength   int
	uploadComplete *bool
}

type fileHandler struct {
	db      *sql.DB
	dirPath string
}

const (
	dirName = "fileserver"
	dbUser  = "postgres"
	dbPwd   = "postgres"
	dbName  = "fileserver"
	sslMode = "disable"
)

func (fh fileHandler) createTable() error {
	q := `CREATE TABLE IF NOT EXISTS file(file_id SERIAL PRIMARY KEY, 
 		  file_offset INT NOT NULL, file_upload_length INT NOT NULL, file_upload_complete BOOLEAN NOT NULL, 
          created_at TIMESTAMP default NOW() NOT NULL, modified_at TIMESTAMP default NOW() NOT NULL)`
	_, err := fh.db.Exec(q)
	if err != nil {
		return err
	}
	return nil
}

func (fh fileHandler) createFile(f file) (string, error) {
	cfstmt := `INSERT INTO file(file_offset, file_upload_length, file_upload_complete) VALUES($1, $2, $3) RETURNING file_id`
	fileID := 0
	err := fh.db.QueryRow(cfstmt, f.offset, f.uploadLength, f.uploadComplete).Scan(&fileID)
	if err != nil {
		return "", err
	}
	fid := strconv.Itoa(fileID)
	return fid, nil
}

func (fh fileHandler) updateFile(f file) error {
	var query []string
	var param []interface{}
	if f.offset != nil {
		of := fmt.Sprintf("file_offset = $1")
		ofp := f.offset
		query = append(query, of)
		param = append(param, ofp)
	}
	if f.uploadComplete != nil {
		uc := fmt.Sprintf("file_upload_complete = $2")
		ucp := f.uploadComplete
		query = append(query, uc)
		param = append(param, ucp)
	}

	if len(query) > 0 {
		mo := "modified_at = $3"
		mop := "NOW()"

		query = append(query, mo)
		param = append(param, mop)

		qj := strings.Join(query, ",")

		sqlq := fmt.Sprintf("UPDATE file SET %s WHERE file_id = $4", qj)

		param = append(param, f.fileID)

		log.Println("generated update query", sqlq)
		_, err := fh.db.Exec(sqlq, param...)
		if err != nil {
			log.Println("Error during file update", err)
			return err
		}
	}
	return nil
}

func (fh fileHandler) File(fileID string) (file, error) {
	fID, err := strconv.Atoi(fileID)
	if err != nil {
		log.Println("Unable to convert fileID to string", err)
		return file{}, err
	}
	log.Println("going to query for fileID", fID)
	gfstmt := `select file_id, file_offset, file_upload_length, file_upload_complete from file where file_id = $1`
	row := fh.db.QueryRow(gfstmt, fID)
	f := file{}
	err = row.Scan(&f.fileID, &f.offset, &f.uploadLength, &f.uploadComplete)
	if err != nil {
		log.Println("error while fetching file", err)
		return file{}, err
	}
	return f, nil
}

func createFileDir() (string, error) {
	u, err := user.Current()
	if err != nil {
		log.Println("Error while fetching user home directory", err)
		return "", err
	}
	home := u.HomeDir
	dirPath := path.Join(home, dirName)
	err = os.MkdirAll(dirPath, 0744)
	if err != nil {
		log.Println("Error while creating file server directory", err)
		return "", err
	}
	return dirPath, nil
}

func (fh fileHandler) createFileHandler(w http.ResponseWriter, r *http.Request) {
	ul, err := strconv.Atoi(r.Header.Get("Upload-Length"))
	if err != nil {
		e := "Improper upload length"
		log.Printf("%s %s", e, err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(e))
		return
	}
	log.Printf("upload length %d\n", ul)
	io := 0
	uc := false
	f := file{
		offset:         &io,
		uploadLength:   ul,
		uploadComplete: &uc,
	}
	fileID, err := fh.createFile(f)
	if err != nil {
		e := "Error creating file in DB"
		log.Printf("%s %s\n", e, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	filePath := path.Join(fh.dirPath, fileID)
	file, err := os.Create(filePath)
	if err != nil {
		e := "Error creating file in filesystem"
		log.Printf("%s %s\n", e, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer file.Close()
	w.Header().Set("Location", fmt.Sprintf("localhost:8080/files/%s", fileID))
	w.WriteHeader(http.StatusCreated)
	return
}

func (fh fileHandler) fileDetailsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fID := vars["fileID"]
	file, err := fh.File(fID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	log.Println("going to write upload offset to output")
	w.Header().Set("Upload-Offset", strconv.Itoa(*file.offset))
	w.WriteHeader(http.StatusOK)
	return
}

func (fh fileHandler) filePatchHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("going to patch file")
	vars := mux.Vars(r)
	fID := vars["fileID"]
	file, err := fh.File(fID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if *file.uploadComplete == true {
		e := "Upload already completed"
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte(e))
		return
	}
	off, err := strconv.Atoi(r.Header.Get("Upload-Offset"))
	if err != nil {
		log.Println("Improper upload offset", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Printf("Upload offset %d\n", off)
	if *file.offset != off {
		e := fmt.Sprintf("Expected Offset %d got offset %d", *file.offset, off)
		w.WriteHeader(http.StatusConflict)
		w.Write([]byte(e))
		return
	}

	log.Println("Content length is", r.Header.Get("Content-Length"))
	clh := r.Header.Get("Content-Length")
	cl, err := strconv.Atoi(clh)
	if err != nil {
		log.Println("unknown content length")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if cl != (file.uploadLength - *file.offset) {
		e := fmt.Sprintf("Content length doesn't not match upload length.Expected content length %d got %d", file.uploadLength-*file.offset, cl)
		log.Println(e)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(e))
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Received file partially %s\n", err)
		log.Println("Size of received file ", len(body))
	}
	fp := fmt.Sprintf("%s/%s", fh.dirPath, fID)
	f, err := os.OpenFile(fp, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Printf("unable to open file %s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer f.Close()

	n, err := f.WriteAt(body, int64(off))
	if err != nil {
		log.Printf("unable to write %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Println("number of bytes written ", n)
	no := *file.offset + n
	file.offset = &no

	uo := strconv.Itoa(*file.offset)
	w.Header().Set("Upload-Offset", uo)
	if *file.offset == file.uploadLength {
		log.Println("upload completed successfully")
		*file.uploadComplete = true
	}

	err = fh.updateFile(file)
	if err != nil {
		log.Println("Error while updating file", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)

	return

}

func main() {
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s", dbUser, dbPwd, dbName, sslMode)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connection established successfully")
	log.Println("TUS Server started")
	fh := fileHandler{
		db: db,
	}
	dir, err := createFileDir()
	if err != nil {
		log.Fatal("Error creating file server directory", err)
	}
	fh.dirPath = dir
	log.Println("Directory created successfully")
	err = fh.createTable()
	if err != nil {
		log.Println("Error during table creation", err)
	}
	log.Println("table created successfully")
	r := mux.NewRouter()
	r.HandleFunc("/files", fh.createFileHandler).Methods("POST")
	r.HandleFunc("/files/{fileID:[0-9]+}", fh.fileDetailsHandler).Methods("HEAD")
	r.HandleFunc("/files/{fileID:[0-9]+}", fh.filePatchHandler).Methods("PATCH")
	http.ListenAndServe(":8080", r)

}
