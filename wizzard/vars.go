/******************************************************************************************
 *      
 *		SERVICE ZENTAX
 *		Copyright (2014-2016)
 *
 *   MODULE            : ELLITIUM
 *
 *		Description      : Service REST API
 *		Version          : Version 4.1.2
 *		Date Started     : 03.11.2014
 *		Date Changed     : 05.02.2018
 *		Author           : Savchenko Arthur
 *		Last Upadte Date : 05.02.2018 11:10
 *
 *		Условные обозначения, сокращения и соглашения о коде :
 *    Термины, соглашения и обозначения, принятые в документации
 *
 ****************************************************************************************/
package main

import (
	"time"                                                      
	r "github.com/dancannon/gorethink"                          
)

//  Gloabal Variables for connection 
var (
	sessionArray []*r.Session
	CurTime              = time.Now().Format("2006-01-02 15:04:05")   // Формат
	CurTimeShort         = time.Now().Format("2006-01-02")            // Формат
	CurTimeUnix          = time.Now().Format(time.RFC3339)            // Дата UNIX  
	CurTimeNano          = time.Now().Format(time.RFC3339Nano)        // Дата UNIX nano
	ActiveIp      string = "localhost"                                // Активный адрес 
	TempStr       string = ":5555"                                    // Активный порт
	Remarks              = "Version testing - 5.027 17.06.2016 10:00"
	Term                 = "TERMS AND CONDITIONS FOR USE, REPRODUCTION, AND DISTRIBUTION"
	Servicestatus        = "Ok"
	Notify               = "Programm alredy started or port 5555 is busy other service."
)

// Declaration inetrfaces & structure & other type
type Mii interface{}                                          // Interface
type Mif []interface{}                                        // Cрез Interface
type Msr []string                                             // Срез String
type Mst map[string]interface{}                               // Map - string - interface
type Mss map[string]string                                    // Map - string - string
type Msi map[string]int64                                     // Map - string - int64
type Mis map[int64]string                                     // Map - int64 - string
type Msl map[int]string                                       // Map - int   - string
type Mil []int64



// System Constants
const (
	AdressService          = "127.0.0.1"                      // Local адресс
	PortService            = ":5556"                          // Local Port 5555
	DatabaseName           = "System"                         // System - Test Databse
	AdressProductionServer = "10.0.50.16:22215"               // IP Server 
	AdressMorionServer     = "10.0.50.16"                     // IP Local Server
	AdressExternal         = "111.222.333.444:2222"            // IP World External for Connect
	AddressMainserver      = "localhost:28015"                // Local Host Mashine
	ServerVersion          = "2.0.5"                          // Current Version
	CodeMirrorVer          = "4.8.001"                        // Editor version
	ServerDescript         = "Draft"                          // Draft Version
	TBN                    = "Docmove"                        // Table Name for Documeents
	DBN                    = "test"                           // Database for Tset
	DHO                    = "HO"                             // Database Head Office
	DSYS                   = "System"                         // System Database
	Organization           = "Zentax"                         // Organization Name
	rupor                  = "Finish Work"                    // Check for Test
    ServiceClientKey     = "S0864AA791CE7E7B00R1T$$"      // Secret Client Key By Default  
    ServiceSecretKey     = "A0AEC09A647A688A64A28"        // Secret Service Key By Default 
    DatabaseKey          = "000Orion000"                  // Secret Key for Database
    SecKey               = "KeySecret$"                   // AccessKey 
)    
