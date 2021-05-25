package main

import(
	"fmt"
	"net/http"
	//"log"
	"Zeinab-project/server/books"
	"Zeinab-project/server/readers"
	"encoding/json"
	"bytes"
	"io/ioutil"
	"strconv"
)

func deleteOperation(){
	var ID int
	fmt.Print("Enter Reader ID to delete: ")
	fmt.Scan(&ID)
	
	reqBody, _ := json.Marshal(map[string]int{
		"ID": ID,
	})
	http.Post("http://localhost:8080/delete-reader", "application/json", bytes.NewBuffer(reqBody))
}

func addOperation(){
	fmt.Println("Enter reader data: (ID, name, gender, birthday year, weight, height and employment) respictively: ")
	var ID int
	var name string
	var gender string
	var birthday int
	var weight int
	var height int
	var employment string
	fmt.Scan(&ID, &name, &gender, &birthday, &weight, &height, &employment)
	var x readers.Reader = readers.Reader{ID , name, gender, birthday, weight, height, employment} 
	json, _ := json.Marshal(x)
	http.Post("http://localhost:8080/add-reader", "application/json",  bytes.NewBuffer([]byte(json)))
}

func selectAllOperation(){

	response, _ := http.Get("http://localhost:8080/load-readers")
	body, _ := ioutil.ReadAll(response.Body)
	var x[] readers.Reader
	json.Unmarshal(body, &x)
	for i := 0; i < len(x); i++{
		fmt.Println("==============================")
		fmt.Println( "Reader info:- \nID: " + strconv.Itoa(x[i].Id) + "\nName: " + x[i].Name + 
        "\nGender: " + x[i].Gender + "\nWeight: " + strconv.Itoa(x[i].Weight) +
         "\nHeight: " + strconv.Itoa(x[i].Height) + "\nEmployment: " + x[i].Employment)
	}
}

func selectByIDOperation(){

	fmt.Print("What is the ID of the Reader? ")
	var ID int
	fmt.Scan(&ID)

	reqBody, err := json.Marshal(map[string]int{
		"ID": ID,
	})

	if err != nil{
		fmt.Println(err)
	}

	response, _ := http.Post("http://localhost:8080/search-reader-by-id", "application/json", bytes.NewBuffer(reqBody))
	body, _ := ioutil.ReadAll(response.Body)
	var x readers.Reader
	json.Unmarshal(body, &x)
	if x.Id == 0{
		fmt.Println("Reader not found")
	}else{
		fmt.Println( "Reader info:- \nID: " + strconv.Itoa(x.Id) + "\nName: " + x.Name + 
        "\nGender: " + x.Gender + "\nWeight: " + strconv.Itoa(x.Weight) +
         "\nHeight: " + strconv.Itoa(x.Height) + "\nEmployment: " + x.Employment)
	}
}

func selectByNameOperation(){
	
	fmt.Print("What is the name of the Reader? ")
	var name string
	fmt.Scan(&name)

	reqBody, _ := json.Marshal(map[string]string{
		"name": name,
	})
	response, _ := http.Post("http://localhost:8080/search-reader-by-name", "application/json", bytes.NewBuffer(reqBody))
	body, _ := ioutil.ReadAll(response.Body)
	var x readers.Reader
	json.Unmarshal(body, &x)
	if x.Name == ""{
		fmt.Println("Reader not found")
	}else{
		fmt.Println( "Reader info:- \nID: " + strconv.Itoa(x.Id) + "\nName: " + x.Name + 
        "\nGender: " + x.Gender + "\nWeight: " + strconv.Itoa(x.Weight) +
         "\nHeight: " + strconv.Itoa(x.Height) + "\nEmployment: " + x.Employment)
	}

}


func workOnReaders(){
	fmt.Println("What kind of operations you want to make on Readers?\n" + 
	"1. Add Reader\n2. Remove Reader\n3. Load all Readers\n4. Search for Reader by name\n5. Search for Reader by ID: \n6.Back")
	  for true{
		  fmt.Print("Enter your operation: ")
		  var op int
		  fmt.Scan(&op)
		  switch op{
			  case 1:
				addOperation()
				break
			   case 2:
				deleteOperation()
				break
			  case 3:
				selectAllOperation()
				break
			  case 4:
				selectByNameOperation()
				break
			  case 5:
				selectByIDOperation()
				break
			  case 6:
				main()
				return
			  default:
				fmt.Println("Invalid operation")
				break
		  }
	  }
}


func workOnBooks(){
	fmt.Println("What kind of operations you want to make on books?\n" + 
	"1. Add Book\n2. Load Books\n3. Search for Book By ID\n4. Search Book by Name\n5. Sort by Book title\n6. Sort by publish year\n7. Back")
	  for true{
		  fmt.Print("Enter your operation: ")
		  var op int
		  fmt.Scan(&op)
		  switch op{
			  case 1:
				addBookOperation()
				break
			  case 2:
				loadAllBooksOperation()
				break
			  case 3:
				searchBookByIDOperation()
				break
			  case 4:
				searchBookByTitleOperation()
				break
			  case 5:
				sortBooksByTitleOperation()
				break
			  case 6:
				sortBooksByYearOperation()
			  case 7:
				main()
				return
			  default:
				  fmt.Println("Invalid operation")
				  break
		  }
	  }
}







func addBookOperation(){
	fmt.Println("Enter book data: (ID, title, author, language, publisher and publish year) respictively ")
	var ID int
	var title string
	var author string
	var language string
	var publisher string
	var publishYear int
	fmt.Scan(&ID, &title, &author, &language, &publisher, &publishYear)
	var x books.Book = books.Book{ID , title, author, language, publisher, publishYear} 
	json, _ := json.Marshal(x)
	http.Post("http://localhost:8080/add-book", "application/json",  bytes.NewBuffer([]byte(json)))
}

func searchBookByIDOperation(){
	fmt.Println("Enter book ID : ")
	var ID int
	fmt.Scan(&ID)
	reqBody, err := json.Marshal(map[string]int{
		"ID": ID,
	})

	if err != nil{
		fmt.Println(err)
	}
	response, _ := http.Post("http://localhost:8080/search-book-by-id", "application/json", bytes.NewBuffer(reqBody))
	body, _ := ioutil.ReadAll(response.Body)
	var x books.Book
	json.Unmarshal(body, &x)
	if x.Id == 0{
		fmt.Println("book not found")
	}else{
		fmt.Println( "Book info:- \nID: " + strconv.Itoa(x.Id) + "\nTitle: " + x.Title + 
        "\nAuthor: " + x.Author + "\nLanguage: " + x.Language +
         "\nPublisher: " + x.Publisher + "\nPublish year: " + strconv.Itoa(x.PublishYear))
	}
} 

func searchBookByTitleOperation(){
	fmt.Println("Enter book Title: ")
	var title string
	fmt.Scan(&title)
	reqBody, _ := json.Marshal(map[string]string{
		"title": title,
	})	
	response, _ := http.Post("http://localhost:8080/search-book-by-title", "application/json", bytes.NewBuffer(reqBody))
	body, _ := ioutil.ReadAll(response.Body)
	var x books.Book
	json.Unmarshal(body, &x)
	if x.Title == ""{
		fmt.Println("book not found")
	}else{
		fmt.Println( "Book info:- \nID: " + strconv.Itoa(x.Id) + "\nTitle: " + x.Title + 
        "\nAuthor: " + x.Author + "\nLanguage: " + x.Language +
         "\nPublisher: " + x.Publisher + "\nPublish year: " + strconv.Itoa(x.PublishYear))
	}
}

func sortBooksByTitleOperation(){
	fmt.Println("Sort done successfully")
	http.Get("http://localhost:8080/sort-books-by-title")
}

func sortBooksByYearOperation(){
	fmt.Println("Sort done successfully")
	http.Get("http://localhost:8080/sort-books-by-year")
}

func loadAllBooksOperation(){
	response, _ := http.Get("http://localhost:8080/load-book")
	body, _ := ioutil.ReadAll(response.Body)
	var x[] books.Book
	json.Unmarshal(body, &x)
	for i := 0; i < len(x); i++{
		fmt.Println("==============================")
		fmt.Println( "Book info:- \nID: " + strconv.Itoa(x[i].Id) + "\nTitle: " + x[i].Title + 
        "\nAuthor: " + x[i].Author + "\nLanguage: " + x[i].Language +
         "\nPublisher: " + x[i].Publisher + "\nPublish year: " + strconv.Itoa(x[i].PublishYear))
	}
}






func main(){
	fmt.Println("Which thing do you want to work on?\n1. Readers\n2. books\n3. Exit ")
	var op int
	for{
		fmt.Print("Enter operation: ")
		fmt.Scan(&op)
		switch op{
			case 1:
				workOnReaders()
				break
			case 2:
				workOnBooks()
				break
			case 3:
				return
			default:
				fmt.Println("Invalid operation")
				break 
		}
	}
}