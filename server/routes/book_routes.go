package routes

import(
	"net/http"
	"fmt"
    "Zeinab-project/server/books"
    "io/ioutil"
    "encoding/json"
    //"strconv"
)

func PostAddBook(w http.ResponseWriter, r *http.Request){
    body, _ := ioutil.ReadAll(r.Body)
    var book books.Book	
    json.Unmarshal(body, &book)
    books.AddBook(book)
}

func PostSearchForBookByID(w http.ResponseWriter, r *http.Request){

    body, err := ioutil.ReadAll(r.Body)
    fmt.Println(string(body))
    if err != nil{
        fmt.Println(err)
    }
	var temp map[string]int
    json.Unmarshal(body, &temp)
    var fetchedBook books.Book = books.SearchByID(temp["ID"])
    fmt.Println(fetchedBook)

    if fetchedBook.Id == 0{
        fmt.Fprintf(w, "Book not found")
    }else{
        readersByteArray, _ := json.MarshalIndent(fetchedBook, "", "  ")
        fmt.Fprintf(w, string(readersByteArray))
    }
}


func PostSearchForBookByTitle(w http.ResponseWriter, r *http.Request){

    body, _ := ioutil.ReadAll(r.Body)
	var temp map[string]string
    json.Unmarshal(body, &temp)
    var fetchedBook books.Book = books.SearchByName(temp["title"])
    if fetchedBook.Title == ""{
        fmt.Fprintf(w, "Book not found")
    }else{
        readersByteArray, _ := json.MarshalIndent(fetchedBook, "", "  ")
        fmt.Fprintf(w, string(readersByteArray))
    }
}

func GetSortByTitle(w http.ResponseWriter, r *http.Request){
    books.SortByTitle()        
}

func GetSortByPublishYear(w http.ResponseWriter, r *http.Request){
    books.SortByPublishYear()
}

func GetLoadAllBook(w http.ResponseWriter, r *http.Request){
    books.LoadBooks()
    x, err := json.MarshalIndent(books.Books, "", "  ")

    fmt.Fprintf(w, string(x))

	if err != nil {
		fmt.Println(err)
	}
}