package main

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"run/dto"
	service "run/login"
	"strings"
	"time"

	"net/http"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"github.com/tealeg/xlsx"
)

func main() {
	LoadConfig()

	router := mux.NewRouter()
	router.HandleFunc("/login", LoginHandler).Methods("POST")
	router.HandleFunc("/fileupload", FileUploadHandler).Methods("POST")
	router.HandleFunc("/addflock", AddFlockHandler).Methods("POST")
	router.HandleFunc("/listflock", ListFlockHandler).Methods("POST")
	router.HandleFunc("/updateflock", UpdateflockHandler).Methods("POST")
	router.HandleFunc("/listbyflock", ListFlockbyHandler).Methods("POST")
	router.HandleFunc("/dailyentries",AddEntryHandler).Methods("POST")
	router.HandleFunc("/addreminder", AddReminderHandler).Methods("POST")
	//router.HandleFunc("/showreminder", ShowReminderHandler).Methods("POST")
	router.HandleFunc("/listflockentries", ListFlockEntryHandler).Methods("POST")
	router.HandleFunc("/listparticularflock", ListParticularFlockHandler).Methods("POST")

	directoryLocation := viper.GetString("uiDirectory")
	log.Println("this is the UI Directory Location : ", directoryLocation)
	router.PathPrefix("/").Handler(http.FileServer(http.Dir(directoryLocation)))
	err := http.ListenAndServe(viper.GetString("port"), router)
	if err != nil {
		log.Println(err)
	}

}
func LoadConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Println("Config not found...", err)
	}

}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	log.Println("++++++++++++++++++++++++++++  login handler +++++++++++++++++++++++++")

	var logindata dto.Logindata

	if err := json.NewDecoder(r.Body).Decode(&logindata); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Println(logindata)

	msg, err := service.Login(logindata)
	log.Println("Received msg:", msg)
	if err != nil {
		http.Error(w, "Invalid Credentials", http.StatusInternalServerError)
		return
	}

	response := struct {
		Message string `json:"message"`
	}{
		Message: msg,
	}
	json.NewEncoder(w).Encode(response)

}

func AddFlockHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	log.Println("++++++++++++++++++++++++++++  AddFlockHandler handler +++++++++++++++++++++++++")

	var flockdata dto.Flockdata
	fmt.Println(r.Body)

	if err := json.NewDecoder(r.Body).Decode(&flockdata); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	if flockdata.FlockName == "" || flockdata.BreedName == "" || flockdata.StartDate == "" || flockdata.StartAge ==0 || flockdata.NoBirds == 0 || flockdata.ShedNumber == "" {
		http.Error(w, "Incomplete or invalid flock data", http.StatusBadRequest)
		return
	}
	fmt.Println("gddgyufgyugeuwy")
	var data dto.Flockdata
	data = dto.Flockdata{
		FlockName:  flockdata.FlockName,
		BreedName:  flockdata.BreedName,
		StartDate:  flockdata.StartDate,
		StartAge:   flockdata.StartAge,
	//	Age:        flockdata.Age,
		NoBirds:    flockdata.NoBirds,
		ShedNumber: flockdata.ShedNumber,
		Active:     "true",
		CreatedAt:  time.Now().String(),
	}

	log.Println(data)

	msg, err := service.AddFlock(data)
	log.Println("Received msg:", msg)
	if err != nil {
		http.Error(w, "Invalid Credentials", http.StatusInternalServerError)
		return
	}

	response := struct {
		Message string `json:"message"`
	}{
		Message: msg,
	}
	json.NewEncoder(w).Encode(response)

}
type Ids struct {
	ID string `bson:"id,omitempty" json:"id,omitempty"`
} 
func ListFlockbyHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	log.Println("++++++++++++++++++++++++++++  ListFlockbyHandler handler +++++++++++++++++++++++++")
	// type lockFlock struct {
	// 	ID string `json:"id"`
	// }

	// var data lockFlock
	// if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
	// 	log.Println(err)
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }

	b, err := ioutil.ReadAll(r.Body)
    if err != nil {
        fmt.Println(err)
        return
    }
    var request Ids
    err = json.Unmarshal(b, &request)
    if err != nil {
        log.Println(err)
    }
    log.Println(request.ID)

	id := request.ID
	list,err := service.ListFlockbyid(id)
	json.NewEncoder(w).Encode(list)
}

func ListFlockHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	log.Println("++++++++++++++++++++++++++++  ListFlockHandler handler +++++++++++++++++++++++++")
	list := service.ListFlock()
	json.NewEncoder(w).Encode(list)
}


// func AgeHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	log.Println("++++++++++++++++++++++++++++  ListFlockHandler handler +++++++++++++++++++++++++")
// 	list := service.AgeCalculator()
// 	json.NewEncoder(w).Encode(list)
// }

func UpdateflockHandler ( w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	log.Println("++++++++++++++++++++++++++++  UpdateflockHandler handler +++++++++++++++++++++++++")
    var flockdata dto.Flockdata
	fmt.Println(r.Body)

	if err := json.NewDecoder(r.Body).Decode(&flockdata); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if flockdata.ID == ""  || flockdata.FlockName == ""  || flockdata.BreedName == "" || flockdata.StartDate == "" || flockdata.StartAge == 0 || flockdata.NoBirds == 0|| flockdata.ShedNumber == "" {
		http.Error(w, "Incomplete or invalid flock data", http.StatusBadRequest)
		return
	}
	

	log.Println(flockdata)

	msg, err := service.UpdateFlock(flockdata)
	log.Println("Received msg:", msg)
	if err != nil {
		http.Error(w, "Invalid Credentials", http.StatusInternalServerError)
		return
	}

	response := struct {
		Message string `json:"message"`
	}{
		Message: msg,
	}
	json.NewEncoder(w).Encode(response)

}

func getXlsxData(trasactionData string) (*xlsx.File, error) {
	var file *xlsx.File
	var byteData []byte
	s := strings.Split(trasactionData, ",")
	if len(s) == 2 {
		byteData = make([]byte, base64.StdEncoding.DecodedLen(len(s[1])))
		_, err := base64.StdEncoding.Decode(byteData, []byte(s[1]))
		if err != nil {
			log.Println("Error :", err)
			return nil, err
		}
		if file, err = xlsx.OpenBinary(byteData); err != nil {
			log.Println("Error in File :", err)
			return nil, err
		}
	} else {
		return nil, errors.New("error")
	}
	return file, nil
}

type Filemodel struct {
	File string `json:"file"`
}

func FileUploadHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Uploading................................................................")
	w.Header().Set("Content-Type", "application/json")
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	file1 := &Filemodel{}

	//data, _ := json.Marshal(r.Body)
	//log.Println(r.Body)
	err1 := json.Unmarshal(b, &file1)
	if err1 != nil {
		//http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println("...............unmarshal error: ..........", err1)
		//return
	}

	//log.Printf("%T", b)

	//fmt.Println(file1.File)

	if file1.File == "" {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("UNMARSHAL ERROR")
		return
	}

	xlFile, err := getXlsxData(file1.File)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("getxlxs file", err)
		return
	}
	jsonData := make(map[string]interface{})
	for _, sheet := range xlFile.Sheets {
		var rows []map[string]interface{}
		for _, row := range sheet.Rows {
			rowData := make(map[string]interface{})
			for _, cell := range row.Cells {
				rowData[cell.String()] = cell.Value

			}
			rows = append(rows, rowData)
			//log.Println(rows)
		}
		jsonData[sheet.Name] = rows
	}

	// Convert JSON data to string
	jsonString, err := json.Marshal(jsonData)
	if err != nil {
		http.Error(w, "Failed to convert data to JSON string", http.StatusInternalServerError)
		log.Println("Failed to convert data to JSON string:", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonString)
}

func AddReminderHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	log.Println("++++++++++++++++++++++++++++  AddReminderHandler +++++++++++++++++++++++++")

	var reminder dto.Reminder
	fmt.Println(r.Body)

	if err := json.NewDecoder(r.Body).Decode(&reminder); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if reminder.Name == "" || reminder.BeforeDate ==  ""|| reminder.AfterDate == "" || reminder.Date == "" || reminder.Remarks == ""  {
		http.Error(w, "Incomplete or invalid flock data", http.StatusBadRequest)
		return
	}
	fmt.Println("gddgyufgyugeuwy")

	msg, err := service.AddReminder(reminder)
	log.Println("Received msg:", msg)
	if err != nil {
		http.Error(w, "Invalid Credentials", http.StatusInternalServerError)
		return
	}

	response := struct {
		Message string `json:"message"`
	}{
		Message: msg,
	}
	json.NewEncoder(w).Encode(response)

}

func AddEntryHandler (w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	log.Println("++++++++++++++++++++++++++++  AddEntryHandler handler +++++++++++++++++++++++++")
	var entry dto.DailyEntry
	fmt.Println("-----request-------",r.Body)
	if err:=json.NewDecoder(r.Body).Decode(&entry); err != nil {
		log.Println("---error in decodinng---", err)
		http.Error(w,err.Error(),http.StatusBadRequest)
		return
	}
	log.Println("navaneesh",entry)
	if entry.ID == ""{
		http.Error(w, "Incomplete Entry data", http.StatusBadRequest)
		return
	}
	msg,err:=service.DailyEntry(entry)
	log.Println("Reciecved mesg:",msg)
	if err!=nil {
		http.Error(w,"Invalid Credentials",http.StatusInternalServerError)
		return 
	}
	response := struct {
		Message string `json:"message"`
	}{
		Message: msg,
	}
	json.NewEncoder(w).Encode(response)
}

func DeleteReminderHandler(w http.ResponseWriter, r *http.Request){


}

func ListFlockEntryHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	log.Println("++++++++++++++++++++++++++++  ListFlockEntryHandler +++++++++++++++++++++++++")
	list := service.ListFlockEntry()
	json.NewEncoder(w).Encode(list)
}



func ListParticularFlockHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	log.Println("++++++++++++++++++++++++++++  List PArticular FlockEntry Handler +++++++++++++++++++++++++")
	b, err := ioutil.ReadAll(r.Body)
    if err != nil {
        fmt.Println(err)
        return
    }
    var request Ids
    err = json.Unmarshal(b, &request)
    if err != nil {
        log.Println(err)
    }
    log.Println(request.ID)
	id := request.ID

	result:=service.ListParticularFlock(id)
	json.NewEncoder(w).Encode(result)

}






