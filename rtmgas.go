package rtmgas

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoString string = os.Getenv("MONGOSTRING")

func MongoConnect(dbname string) (db *mongo.Database) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MongoString))
	if err != nil {
		fmt.Printf("MongoConnect: %v\n", err)
	}
	return client.Database(dbname)
}
func InsertOneDoc(db string, collection string, doc interface{}) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
	}
	return insertResult.InsertedID
}
func InsertAbsensi(nip string, date string, status string, employee string, positions string) (InsertedID interface{}) {
	var absensi Absensi
	absensi.NIP = nip
	absensi.Date = date
	absensi.Status = status
	absensi.Employee = employee
	absensi.Positions = positions

	return InsertOneDoc("176892865", "absensi", absensi)
}
func InsertPosisi(nama_posisi string, bidang string) (InsertedID interface{}) {
	var posisi Posisi
	posisi.Nama_Posisi = nama_posisi
	posisi.Bidang = bidang

	return InsertOneDoc("Engineering Manager", "posisi", posisi)
}
func GetDataAbsensi(nip string) (data []Absensi) {
	user := MongoConnect("rtmtugas").Collection("absen")
	filter := bson.M{"nip": nip}
	cursor, err := user.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetDataAbsensi :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}
func GetDataPosisi(nama_posisi string) (data []Posisi) {
	user := MongoConnect("rtmtugas").Collection("Posisi")
	filter := bson.M{"nama_posisi": nama_posisi}
	cursor, err := user.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetDataPosisi :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}
