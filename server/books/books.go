package books

import (
	"fmt"
	"io/ioutil"
	"encoding/json"
	"strconv"
	"sort"
)

var Books []Book

type Book struct {
	Id int `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Language string `json:"language"`
	Publisher  string `json:"publisher"`
	PublishYear int  `json:"publishYear"`
}



func AddBook(book Book){
	fmt.Println("ok")
	LoadBooks()
	Books = append(Books, Book{
		Id: book.Id,
		Title: book.Title,
		Author: book.Author,
		Language: book.Language,
		Publisher: book.Publisher,
		PublishYear: book.PublishYear,
	})
	saveBooks()
}

func SearchByID(ID int) Book{
	LoadBooks()
	fmt.Println("done")
	var validBook Book
	for _,x := range Books{
		if ID == x.Id{
			validBook = x
			break
		}
	}
	return validBook
}

func SearchByName(name string) Book{
	LoadBooks()
	var validBook Book
	for _,x := range Books{
		if name == x.Title{
			validBook = x
			break
		}
	}
	return validBook
}

func printBookInfo(book Book){
	fmt.Println("Book info:- \nID: " + strconv.Itoa(book.Id) + "\nTitle: " + book.Title + 
	"\nAuthor: " + book.Author + "\nLanguage: " + book.Language +
	 "\nPublisher: " + book.Publisher + "\nPublish year: " + strconv.Itoa(book.PublishYear))
}

func LoadAllBooks(){
	for i := 0; i < len(Books); i++ {
		printBookInfo(Books[i])
	}
}


func LoadBooks(){
	content, err := ioutil.ReadFile("Books.json")
	if err != nil {
		fmt.Println(err.Error())
	}
	err2 := json.Unmarshal(content, &Books)
	for i := 0; i < len(Books); i++{
		fmt.Println(Books[i].Title)
	}
	if err2 != nil {
		fmt.Println("Error JSON Unmarshalling")
		fmt.Println(err2.Error())
	}
}

func saveBooks(){
	fmt.Println(Books)
	byteArray, err := json.Marshal(Books)
	if err != nil {
		fmt.Println(err)
	}
	err = ioutil.WriteFile("Books.json", byteArray, 0644)
    if err != nil {
        fmt.Println(err)
	}
	fmt.Println("Save done successfully")
}

func SortByTitle(){
	LoadBooks()
	sort.Slice(Books[:], func(i, j int) bool {
		return Books[i].Title < Books[j].Title
	})
	  saveBooks()
}

func SortByPublishYear(){
	LoadBooks()
	sort.Slice(Books[:], func(i, j int) bool {
		return Books[i].PublishYear < Books[j].PublishYear
	})
	saveBooks()
}