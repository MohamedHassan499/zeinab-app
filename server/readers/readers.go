package readers

import (
	"fmt"
	"io/ioutil"
	"encoding/json"
	"strconv"
)

var Readers []Reader

type Reader struct {
	Id int `json:"id"`
	Name string `json:"name`
	Gender  string `json:"gender"`
	Birthday_year int `json:"birthday_year"`
	Weight int `json:"weight"`
	Height  int `json:"height"`
	Employment string  `json:"employment"`
}

func AddReader(r Reader){
	LoadReaders()
	Readers = append(Readers, Reader{
		Id: r.Id,
		Name: r.Name,
		Gender: r.Gender,
		Birthday_year: r.Birthday_year,
		Weight: r.Weight,
		Height: r.Height,
		Employment: r.Employment,
	})
	saveReaders()
}

func RemoveReader(ID int){
	LoadReaders()
	for i := 0; i < len(Readers); i++{
		if Readers[i].Id == ID {
			Readers[i] = Readers[len(Readers)-1] // Copy last element to index i.
			Readers[len(Readers)-1] = Reader{}   // Erase last element (write zero value).
			Readers = Readers[:len(Readers)-1]
			saveReaders()
			break
		}
	}
	saveReaders()
}

func SearchByID(ID int) Reader{
	LoadReaders()
	var validReader Reader
	for _,x := range Readers{
		if ID == x.Id{
			validReader = x
		}
	}
	return validReader
}

func SearchByName(name string) Reader{
	LoadReaders()
	var validReader Reader
	for _,x := range Readers{
		if name == x.Name{
			validReader = x
		}
	}
	return validReader
}

func printReaderInfo(reader Reader){
	fmt.Println("Book Reader:- \nID: " + strconv.Itoa(reader.Id) + "\nName: " + reader.Name + 
	"\nGender: " + reader.Gender + "\nBirthday year: " + strconv.Itoa(reader.Birthday_year) +
	 "\nWeight: " + strconv.Itoa(reader.Weight) + "\nHeight: " + strconv.Itoa(reader.Height) + 
	"\nEmployment: " + reader.Employment)
}

func LoadAllReaders(){
	LoadReaders()
	for i := 0; i < len(Readers); i++{
		fmt.Println("==========================================")
		printReaderInfo(Readers[i])
		fmt.Println("==========================================")
	}
}

func LoadReaders(){
	content, err := ioutil.ReadFile("Readers.json")
	if err != nil {
		fmt.Println(err.Error())
	}
	err2 := json.Unmarshal(content, &Readers)
	if err2 != nil {
		fmt.Println("Error JSON Unmarshalling")
		fmt.Println(err2.Error())
	}
}

func saveReaders(){
	byteArray, err := json.Marshal(Readers)
	if err != nil {
		fmt.Println(err)
	}
	err = ioutil.WriteFile("Readers.json", byteArray, 0644)
    if err != nil {
        fmt.Println(err)
	}
	fmt.Println("Save done successfully")
}