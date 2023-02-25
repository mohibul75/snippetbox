package main

import (
	"log"
	"net/http"
	"path/filepath"
	"flag"
	"os"
)

type neuteredFileSystem struct{
	fs http.FileSystem
}

func main() {

	addr:= flag.String("addr",":4000","HTTP network address")
	flag.Parse()

	mux := http.NewServeMux()

	infoLog := log.New(os.Stdout,"INFO\t",log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t",log.Ldate|log.Ltime|log.Lshortfile)


	fileServer:= http.FileServer(neuteredFileSystem{http.Dir("./ui/static/")})
	mux.Handle("/static",http.NotFoundHandler())
	mux.Handle("/static/", http.StripPrefix("/static",fileServer))

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCrete)

	//log.Println("Starting Server on ", *addr)

	infoLog.Printf("Starting Server on  %s", *addr)
	err := http.ListenAndServe(*addr, mux)
	errorLog.Fatal(err)
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


