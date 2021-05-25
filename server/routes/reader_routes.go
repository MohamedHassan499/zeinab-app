package routes

import(
	"net/http"
	"fmt"
	"Zeinab-project/server/readers"
	"encoding/json"
	"io/ioutil"
)

func PostAddReader(w http.ResponseWriter, r *http.Request){
		body, _ := ioutil.ReadAll(r.Body)
		fmt.Println(string(body))
		var reader readers.Reader	
		json.Unmarshal(body, &reader)
		readers.AddReader(reader)
}

func DeleteReader(w http.ResponseWriter, r *http.Request){
	body, _ := ioutil.ReadAll(r.Body)
	fmt.Println(string(body))
	var temp map[string]int
	json.Unmarshal(body, &temp)
	fmt.Print("ok")
	readers.RemoveReader(temp["ID"])
}

func GetAllReaders(w http.ResponseWriter, r *http.Request){
	readers.LoadReaders()
    x, err := json.MarshalIndent(readers.Readers, "", "  ")

    fmt.Fprintf(w, string(x))

	if err != nil {
		fmt.Println(err)
	}
}

func PostAReaderByID(w http.ResponseWriter, r *http.Request){
	body, err := ioutil.ReadAll(r.Body)
    fmt.Println(string(body))
    if err != nil{
        fmt.Println(err)
    }
	var temp map[string]int
    json.Unmarshal(body, &temp)
    var fetchedReader readers.Reader = readers.SearchByID(temp["ID"])
    fmt.Println(fetchedReader)

    if fetchedReader.Id == 0{
        fmt.Fprintf(w, "Reader not found")
    }else{
        readersByteArray, _ := json.MarshalIndent(fetchedReader, "", "  ")
        fmt.Fprintf(w, string(readersByteArray))
    }
}

func PostAReaderByName(w http.ResponseWriter, r *http.Request){
	body, err := ioutil.ReadAll(r.Body)
    fmt.Println(string(body))
    if err != nil{
        fmt.Println(err)
    }
	var temp map[string]string
    json.Unmarshal(body, &temp)
    var fetchedReader readers.Reader = readers.SearchByName(temp["name"])
    fmt.Println(fetchedReader)

    if fetchedReader.Id == 0{
        fmt.Fprintf(w, "Reader not found")
    }else{
        readersByteArray, _ := json.MarshalIndent(fetchedReader, "", "  ")
        fmt.Fprintf(w, string(readersByteArray))
    }
}