package rtmgas

import (
	"fmt"
	"testing"

	"github.com/aiteung/atdb"
)

var MongoInfo = atdb.DBInfo{
	DBString: MongoString,
	DBName:   "rtmtugas",
}
var MongoConn = atdb.MongoConnect(MongoInfo)

func TestInsertAbsensi(t *testing.T) {
	nip := "176892865"
	date := "Jul, 20 2023"
	status := "Hadir"
	employee := "Arthur Melo"
	positions := "Software Engineeer"
	hasil := TestInsertAbsensi(nip, date, status, employee, positions)
	fmt.Println(hasil)
}
func TestInsertPosisi(t *testing.T) {
	nama_posisi := "Engineering Manager"
	bidang := "Slack"
	hasil := InsertPosisi(nama_posisi, bidang)
	fmt.Println(hasil)
}
func TestGetDataAbsensi(t *testing.T) {
	nip := "176892865"
	date := "Jul, 20 2023"
	status := "Hadir"
	employee := "Arthur Melo"
	positions := "Software Engineeer"
	hasil := TestGetDataAbsensi(nip, date, status, employee, positions)
	fmt.Println(hasil)
}

func TestGetDataPosisi(t *testing.T) {
	nama_posisi := "Engineering Manager"
	bidang := "Slack"
	hasil := GetDataPosisi(nama_posisi, bidang)
	fmt.Println(hasil)
}
