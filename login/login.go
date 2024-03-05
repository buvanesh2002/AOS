package service

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"log"
	"run/config"
	"run/dto"
	"time"

	"github.com/spf13/viper"
	"gopkg.in/mgo.v2/bson"
)

func Login(value dto.Logindata) (string, error) {
	log.Println("++++++++++++++++++++++++++++  login service +++++++++++++++++++++++++")

	log.Println(value)
	Client := config.GetConfig()
	defer Client.Disconnect(context.Background())

	collection := Client.Database(viper.GetString("db")).Collection("users")

	filter := bson.M{"email": value.Email, "password": value.Password}

	err := collection.FindOne(context.Background(), filter).Decode(&value)
	log.Println(err)
	if err != nil {
		return "Invalid Credentials", err
	} else {
		return "Login successful", nil
	}

}

func AddFlock(value dto.Flockdata) (string, error) {
	log.Println("++++++++++++++++++++++++++++  AddFlock service +++++++++++++++++++++++++")
	randomID, err := generateRandomID(10)
	if err != nil {
		log.Println("Error:", err)
		return "", err
	}
	value.ID = randomID
	Client := config.GetConfig()
	defer Client.Disconnect(context.Background())

	collection := Client.Database(viper.GetString("db")).Collection(viper.GetString("Addflock"))

	_, err1 := collection.InsertOne(context.Background(), value)
	if err1 != nil {
		log.Println("Error inserting document:", err1)
		return "", err1
	}

	return "Flock added successfully", nil
}

func ListFlock() *[]dto.Flockdata {
	log.Println("===============List Credentials===============")
	info := []dto.Flockdata{}
	Client := config.GetConfig()
	db := viper.GetString("db")
	log.Println("db name : ", db)

	collection := Client.Database(db).Collection(viper.GetString("Addflock"))
	cur, err := collection.Find(context.Background(), bson.M{"active": "true"})
	if err != nil {
		log.Println("Collection list error : ", err)
	}
	defer Client.Disconnect(context.Background())
	defer cur.Close(context.Background())

	err = cur.All(context.Background(), &info)
	if err != nil {
		log.Println("Error while fetching documents:", err)
		return nil
	}
	return &info
}

// func ListFlockbyid(id string) (dto.Flockdata, error) {
// 	log.Println("===============ListFlockbyid Credentials===============")

// 	var info dto.Flockdata
// 	id = "759c91c579"
// 	Client := config.GetConfig()
// 	db := viper.GetString("db")

// 	collection := Client.Database(db).Collection(viper.GetString("Addflock"))
// 	result, _ := collection.CountDocuments(context.Background(), &info)
// 	// log.Println("Collection list error : ", result.Err())
// 	// return dto.Flockdata{}, result.Err()
// 	log.Println("buvanesh-----", result)

// 	// if err := result.Decode(&info); err != nil {
// 	//     log.Println("Error while decoding document:", err)
// 	//     return dto.Flockdata{}, err
// 	// }

// 	return info, nil

// }

func ListFlockbyid(id string) (dto.Flockdata, error) {
	log.Println("===============ListFlockbyid Credentials===============")

	var info dto.Flockdata
	Client := config.GetConfig()
	db := viper.GetString("db")

	collection := Client.Database(db).Collection(viper.GetString("Addflock"))
	filter := bson.M{"_id": id}
	result := collection.FindOne(context.Background(), filter)
	if result.Err() != nil {
		log.Println("Error while fetching document:", result.Err())
		return dto.Flockdata{}, result.Err()
	}

	if err := result.Decode(&info); err != nil {
		log.Println("Error while decoding document:", err)
		return dto.Flockdata{}, err
	}

	return info, nil
}

func UpdateFlock(value dto.Flockdata) (string, error) {
	log.Println("++++++++++++++++++++++++++++  UpdateFlock service +++++++++++++++++++++++++")

	Client := config.GetConfig()
	defer Client.Disconnect(context.Background())

	collection := Client.Database(viper.GetString("db")).Collection(viper.GetString("Addflock"))

	filter := bson.M{"_id": value.ID} // Assuming value.ID is the ID of the document to update
	log.Println("active=", value.Active)
	log.Println("ID=", value.ID)
	update := bson.M{
		"$set": bson.M{
			"flockName":    value.FlockName,
			"breedName":    value.BreedName,
			"startDate":    value.StartDate,
			"startAge":     value.Age,
			"active":       value.Active,
			"openingBirds": value.NoBirds,
			"shedNumber":   value.ShedNumber,
			"updatedAt":    time.Now().String(),
		},
	}

	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Println("Error updating document:", err)
		return "", err
	}

	return "Flock updated successfully", nil
}

func generateRandomID(length int) (string, error) {
	// Calculate the byte size needed to represent the ID of the specified length
	byteSize := length / 2
	if length%2 != 0 {
		byteSize++
	}

	// Generate random bytes
	bytes := make([]byte, byteSize)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	// Encode bytes to hexadecimal string
	id := hex.EncodeToString(bytes)

	// Truncate the string to the desired length
	id = id[:length]

	return id, nil
}
