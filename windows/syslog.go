package main

import (
	"fmt"
	"golang.org/x/sys/windows"
	"golang.org/x/sys/windows/svc/eventlog"
	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
)

var (
	DB *r.Session
)

func init() {
	var err error
	DB, err = r.Connect(r.ConnectOpts{
		// Replace with the appropriate RethinkDB server address
		Address: "localhost:28015",
		// Replace with the desired database name
		Database: "work",
	})

	if err != nil {
		fmt.Println("Failed to connect to RethinkDB:", err)
		return
	}

	//defer DB.Close()

	// Perform operations with the session...
}

type Loger struct {
	Title string
	Name  string
}

// System log
func main() {
	WriteSysLog()
	WriteLogInfo("Barsetka", "Example system time", 3938)
	WriteLogInfo("Testing", "Example system time", 38393)
	WriteLogInfo("Market", "Example system time", 363777)
}

func mains() {
	t := Loger{
		Title: "sss",
		Name:  "sssssdd",
	}

	err := r.Table("Logger").Insert(t).Exec(DB)
	if err != nil {
		fmt.Println("ERROR:", err.Error())
	}

}

func CreateTable() {
	err := r.TableDrop("Logger").Exec(DB)
	if err != nil {
		fmt.Println("ERROR:", err.Error())
	}

	err = r.TableCreate("Logger", r.TableCreateOpts{PrimaryKey: "id", Durability: "soft"}).Exec(DB)
	if err != nil {
		fmt.Println("ERROR:", err.Error())
	}
}

func WriteSysLog() {
	SysLog("Barsetka", "Sysytem start", "I", 2345)
	SysLog("Sklad", "Sysytem alert", "W", 234)
	SysLog("Other", "Sysytem alert", "E", 4848)
	SysLog("Other", "Запуск системы в процесс", "E", 4848)
}

// Запись в системный лог виндовс
// Events Viewer Windows Log/Application

func SysLog(prog, mess, evtype string, evid uint32) {
	var uuptr uintptr

	eventLog, err := windows.RegisterEventSource(nil, windows.StringToUTF16Ptr(prog))
	if err != nil {
		fmt.Printf("Failed to register event source: %v\n", err)
		return
	}
	defer windows.DeregisterEventSource(eventLog)

	eventId := evid //uint32(1123)
	eventType := uint16(windows.EVENTLOG_INFORMATION_TYPE)
	category := uint16(1)
	message := windows.StringToUTF16Ptr(mess)

	if evtype == "I" {
		eventType = uint16(windows.EVENTLOG_INFORMATION_TYPE)
	}

	if evtype == "W" {
		eventType = uint16(windows.EVENTLOG_WARNING_TYPE)
	}

	if evtype == "E" {
		eventType = uint16(windows.EVENTLOG_ERROR_TYPE)
	}

	rawData := []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08} // nil

	err = windows.ReportEvent(eventLog, eventType, category, eventId, uuptr, 1, 0, &message, &rawData[0])
	if err != nil {
		fmt.Printf("Failed to report event: %v\n", err)
		return
	}

	fmt.Println("Event written to the system log.")
}

// Запись в лог (вариант 2)
func WriteLogInfo(app, mes string, evid uint32) {
	log, err := eventlog.Open(app)
	if err != nil {
		fmt.Printf("Failed to open system log: %v\n", err)
		return
	}
	defer log.Close()
	var bb uint32 = 1

	// log.Warning(bb, mes)
	log.Info(bb, mes)
}
