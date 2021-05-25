package main

import(
	"fmt"
    "net/http"
    "os"
    "Zeinab-project/server/routes"
)

func HandleBooksRoutes(){
    
    http.HandleFunc("/add-book", routes.PostAddBook)
    http.HandleFunc("/search-book-by-title", routes.PostSearchForBookByTitle)
    http.HandleFunc("/search-book-by-id", routes.PostSearchForBookByID)
    http.HandleFunc("/load-book", routes.GetLoadAllBook)
    http.HandleFunc("/sort-books-by-title", routes.GetSortByTitle)
    http.HandleFunc("/sort-books-by-year", routes.GetSortByPublishYear)

}

func HandleReadersRoutes(){

    http.HandleFunc("/add-reader", routes.PostAddReader)
    http.HandleFunc("/delete-reader", routes.DeleteReader)
    http.HandleFunc("/load-readers", routes.GetAllReaders)
    http.HandleFunc("/search-reader-by-id", routes.PostAReaderByID)
    http.HandleFunc("/search-reader-by-name", routes.PostAReaderByName)

}

func main(){

    
    HandleBooksRoutes()
    HandleReadersRoutes()


    fmt.Printf("Starting server at port 8080\n")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Println("COULDN'T CONNECT TO THE FIRST SERVER, CONNECTING TO THE SECOND SERVER\n"+
        "Starting server at port 9090")
        if err2 := http.ListenAndServe(":9090", nil); err2 != nil {
            fmt.Println("COULDN'T CONNECT TO THE SECOND SERVER, APP CRASHED")
            os.Exit(3)
        }
    }
}