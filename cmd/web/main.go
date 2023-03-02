package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"
	"path/filepath"

	_ "github.com/go-sql-driver/mysql"
)

type neuteredFileSystem struct{
	fs http.FileSystem
}

type application struct{
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {

	addr:= flag.String("addr",":4000","HTTP network address")
	dsn := flag.String("dsn", "web:pass@/snippetbox?parseTime=true", "MySQL data source name")

	flag.Parse()

	mux := http.NewServeMux()

	infoLog := log.New(os.Stdout,"INFO\t",log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t",log.Ldate|log.Ltime|log.Lshortfile)


	app:= &application{
		errorLog: errorLog,
		infoLog: infoLog,
	}

	db, err:= openDB(*dsn)
	if err!=nil {
		errorLog.Fatal(err)
	}

	defer db.Close()



	// but in production and staging, It's reffered to redirect logs in a file
	// to do that in programatically, you can follow the below code
	/* 
		f, err:= os.OpenFile("/info.log", os.0_RDWR|os._CREATE, 0666)
		if err !=nil{
			log.Fatal(err)
		}

		defer f.Close()

		infoLog := log.New(f, "INFO\t", log.Ldate|log.Ltime)
	
	*/



	fileServer:= http.FileServer(neuteredFileSystem{http.Dir("./ui/static/")})
	mux.Handle("/static",http.NotFoundHandler())
	mux.Handle("/static/", http.StripPrefix("/static",fileServer))

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/view", app.snippetView)
	mux.HandleFunc("/snippet/create", app.snippetCrete)

	srv:= &http.Server{
		Addr: *addr,
		ErrorLog: errorLog,
		Handler: mux,
	}

	//log.Println("Starting Server on ", *addr)

	infoLog.Printf("Starting Server on  %s", *addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}

func openDB(dsn string)(*sql.DB, error){
	db, err:= sql.Open("mysql",dsn)

	if err!=nil {
		return nil,err
	}

	if err=db.Ping(); err!=nil{
		return nil, err
	}

	return db, nil
}

func (nfs neuteredFileSystem) Open(path string)(http.File, error){
	f,err := nfs.fs.Open(path)

	if err!=nil {
		return nil,err
	}

	s, err:= f.Stat()

	if s.IsDir() {
		index:= filepath.Join(path,"index.html")

		if _, err := nfs.fs.Open(index); err!= nil{
			closeErr:= f.Close()
			if closeErr != nil {
				return nil, closeErr
			
			}

			return nil, err
		}
	}

	return f,nil
}


