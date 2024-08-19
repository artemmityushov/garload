package main

import (
	"archive/zip"
	"bufio"
	"database/sql"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

var (
	sourcePath             = ""
	filesPath              = ""
	databaseIP             = ""
	databasePort           = ""
	databaseName           = ""
	databaseScheme         = ""
	databaseUser           = ""
	databasePassword       = ""
	loadFull               = false
	packetCount      int64 = 100
)

func Ternary(statement bool, a, b interface{}) interface{} {
	if statement {
		return a
	}
	return b
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func getDbConnectString() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", databaseIP, databasePort, databaseUser, databasePassword, databaseName)
}

func workFileADDR_OBJ(sqlObject *sql.DB, fileName string) bool {
	f, err := os.Open(filesPath + fileName)
	if err != nil {

	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)
	decoder := xml.NewDecoder(f)
	var packetIndex int64 = 0
	var index int64 = 0
	rows := make([]itemADD_OBJ, 0)
	for {
		tok, err := decoder.Token()
		if err != nil {

		}
		if tok == nil {
			break
		}
		switch tp := tok.(type) {
		case xml.StartElement:
			if tp.Name.Local == "OBJECT" {
				packetIndex++
				index++
				var row itemADD_OBJ
				err := decoder.DecodeElement(&row, &tp)
				if err != nil {
					return false
				}
				rows = append(rows, row)
				if packetIndex == packetCount {
					packetADDR_OBJ(sqlObject, rows)
					packetIndex = 0
					rows = make([]itemADD_OBJ, 0)
				}
			}
		}
	}
	if packetIndex != 0 {
		packetADDR_OBJ(sqlObject, rows)
	}
	fmt.Println("Count:", index)
	return true
}

func workFileADDR_OBJ_DIVISION(sqlObject *sql.DB, fileName string) bool {
	f, err := os.Open(filesPath + fileName)
	if err != nil {

	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)
	decoder := xml.NewDecoder(f)
	var packetIndex int64 = 0
	var index int64 = 0
	rows := make([]itemADDR_OBJ_DIVISION, 0)
	for {
		tok, err := decoder.Token()
		if err != nil {

		}
		if tok == nil {
			break
		}
		switch tp := tok.(type) {
		case xml.StartElement:
			if tp.Name.Local == "ITEM" {
				packetIndex++
				index++
				var row itemADDR_OBJ_DIVISION
				err := decoder.DecodeElement(&row, &tp)
				if err != nil {
					return false
				}
				rows = append(rows, row)
				if packetIndex == packetCount {
					packetADDR_OBJ_DIVISION(sqlObject, rows)
					packetIndex = 0
					rows = make([]itemADDR_OBJ_DIVISION, 0)
				}
			}
		}
	}
	if packetIndex != 0 {
		packetADDR_OBJ_DIVISION(sqlObject, rows)
	}
	fmt.Println("Count:", index)
	return true
}

func workFileADM_HIERARCHY(sqlObject *sql.DB, fileName string) bool {
	f, err := os.Open(filesPath + fileName)
	if err != nil {

	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)
	decoder := xml.NewDecoder(f)
	var packetIndex int64 = 0
	var index int64 = 0
	rows := make([]itemADM_HIERARCHY, 0)
	for {
		tok, err := decoder.Token()
		if err != nil {

		}
		if tok == nil {
			break
		}
		switch tp := tok.(type) {
		case xml.StartElement:
			if tp.Name.Local == "ITEM" {
				packetIndex++
				index++
				var row itemADM_HIERARCHY
				err := decoder.DecodeElement(&row, &tp)
				if err != nil {
					return false
				}
				rows = append(rows, row)
				if packetIndex == packetCount {
					packetADM_HIERARCHY(sqlObject, rows)
					packetIndex = 0
					rows = make([]itemADM_HIERARCHY, 0)
				}
			}
		}
	}
	if packetIndex != 0 {
		packetADM_HIERARCHY(sqlObject, rows)
	}
	fmt.Println("Count:", index)
	return true
}

func workFileAPARTMENTS(sqlObject *sql.DB, fileName string) bool {
	f, err := os.Open(filesPath + fileName)
	if err != nil {

	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)
	decoder := xml.NewDecoder(f)
	var packetIndex int64 = 0
	var index int64 = 0
	rows := make([]itemAPARTMENTS, 0)
	for {
		tok, err := decoder.Token()
		if err != nil {

		}
		if tok == nil {
			break
		}
		switch tp := tok.(type) {
		case xml.StartElement:
			if tp.Name.Local == "APARTMENT" {
				packetIndex++
				index++
				var row itemAPARTMENTS
				err := decoder.DecodeElement(&row, &tp)
				if err != nil {
					return false
				}
				rows = append(rows, row)
				if packetIndex == packetCount {
					packetAPARTMENTS(sqlObject, rows)
					packetIndex = 0
					rows = make([]itemAPARTMENTS, 0)
				}
			}
		}
	}
	if packetIndex != 0 {
		packetAPARTMENTS(sqlObject, rows)
	}
	fmt.Println("Count:", index)
	return true
}

func workFileCARPLACES(sqlObject *sql.DB, fileName string) bool {
	f, err := os.Open(filesPath + fileName)
	if err != nil {

	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)
	decoder := xml.NewDecoder(f)
	var packetIndex int64 = 0
	var index int64 = 0
	rows := make([]itemCARPLACES, 0)
	for {
		tok, err := decoder.Token()
		if err != nil {

		}
		if tok == nil {
			break
		}
		switch tp := tok.(type) {
		case xml.StartElement:
			if tp.Name.Local == "CARPLACE" {
				packetIndex++
				index++
				var row itemCARPLACES
				err := decoder.DecodeElement(&row, &tp)
				if err != nil {
					return false
				}
				rows = append(rows, row)
				if packetIndex == packetCount {
					packetCARPLACES(sqlObject, rows)
					packetIndex = 0
					rows = make([]itemCARPLACES, 0)
				}
			}
		}
	}
	if packetIndex != 0 {
		packetCARPLACES(sqlObject, rows)
	}
	fmt.Println("Count:", index)
	return true
}

func workFileHOUSES(sqlObject *sql.DB, fileName string) bool {
	f, err := os.Open(filesPath + fileName)
	if err != nil {

	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)
	decoder := xml.NewDecoder(f)
	var packetIndex int64 = 0
	var index int64 = 0
	rows := make([]itemHOUSES, 0)
	for {
		tok, err := decoder.Token()
		if err != nil {

		}
		if tok == nil {
			break
		}
		switch tp := tok.(type) {
		case xml.StartElement:
			if tp.Name.Local == "HOUSE" {
				packetIndex++
				index++
				var row itemHOUSES
				err := decoder.DecodeElement(&row, &tp)
				if err != nil {
					return false
				}
				rows = append(rows, row)
				if packetIndex == packetCount {
					packetHOUSES(sqlObject, rows)
					packetIndex = 0
					rows = make([]itemHOUSES, 0)
				}
			}
		}
	}
	if packetIndex != 0 {
		packetHOUSES(sqlObject, rows)
	}
	fmt.Println("Count:", index)
	return true
}

func workFileMUN_HIERARCHY(sqlObject *sql.DB, fileName string) bool {
	f, err := os.Open(filesPath + fileName)
	if err != nil {

	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)
	decoder := xml.NewDecoder(f)
	var packetIndex int64 = 0
	var index int64 = 0
	rows := make([]itemMUN_HIERARCHY, 0)
	for {
		tok, err := decoder.Token()
		if err != nil {

		}
		if tok == nil {
			break
		}
		switch tp := tok.(type) {
		case xml.StartElement:
			if tp.Name.Local == "ITEM" {
				packetIndex++
				index++
				var row itemMUN_HIERARCHY
				err := decoder.DecodeElement(&row, &tp)
				if err != nil {
					return false
				}
				rows = append(rows, row)
				if packetIndex == packetCount {
					packetMUN_HIERARCHY(sqlObject, rows)
					packetIndex = 0
					rows = make([]itemMUN_HIERARCHY, 0)
				}
			}
		}
	}
	if packetIndex != 0 {
		packetMUN_HIERARCHY(sqlObject, rows)
	}
	fmt.Println("Count:", index)
	return true
}

func workFileROOMS(sqlObject *sql.DB, fileName string) bool {
	f, err := os.Open(filesPath + fileName)
	if err != nil {

	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)
	decoder := xml.NewDecoder(f)
	var packetIndex int64 = 0
	var index int64 = 0
	rows := make([]itemROOMS, 0)
	for {
		tok, err := decoder.Token()
		if err != nil {

		}
		if tok == nil {
			break
		}
		switch tp := tok.(type) {
		case xml.StartElement:
			if tp.Name.Local == "ROOM" {
				packetIndex++
				index++
				var row itemROOMS
				err := decoder.DecodeElement(&row, &tp)
				if err != nil {
					return false
				}
				rows = append(rows, row)
				if packetIndex == packetCount {
					packetROOMS(sqlObject, rows)
					packetIndex = 0
					rows = make([]itemROOMS, 0)
				}
			}
		}
	}
	if packetIndex != 0 {
		packetROOMS(sqlObject, rows)
	}
	fmt.Println("Count:", index)
	return true
}

func workFileSTEADS(sqlObject *sql.DB, fileName string) bool {
	f, err := os.Open(filesPath + fileName)
	if err != nil {

	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)
	decoder := xml.NewDecoder(f)
	var packetIndex int64 = 0
	var index int64 = 0
	rows := make([]itemSTEADS, 0)
	for {
		tok, err := decoder.Token()
		if err != nil {

		}
		if tok == nil {
			break
		}
		switch tp := tok.(type) {
		case xml.StartElement:
			if tp.Name.Local == "STEAD" {
				packetIndex++
				index++
				var row itemSTEADS
				err := decoder.DecodeElement(&row, &tp)
				if err != nil {
					return false
				}
				rows = append(rows, row)
				if packetIndex == packetCount {
					packetSTEADS(sqlObject, rows)
					packetIndex = 0
					rows = make([]itemSTEADS, 0)
				}
			}
		}
	}
	if packetIndex != 0 {
		packetSTEADS(sqlObject, rows)
	}
	fmt.Println("Count:", index)
	return true
}

// Справочники
func workFileADDHOUSE_TYPES(sqlObject *sql.DB, fileName string) bool {
	f, err := os.Open(filesPath + fileName)
	if err != nil {

	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)
	decoder := xml.NewDecoder(f)
	var packetIndex int64 = 0
	var index int64 = 0
	rows := make([]itemADDHOUSE_TYPES, 0)
	for {
		tok, err := decoder.Token()
		if err != nil {

		}
		if tok == nil {
			break
		}
		switch tp := tok.(type) {
		case xml.StartElement:
			if tp.Name.Local == "HOUSETYPE" {
				packetIndex++
				index++
				var row itemADDHOUSE_TYPES
				err := decoder.DecodeElement(&row, &tp)
				if err != nil {
					return false
				}
				rows = append(rows, row)
				if packetIndex == packetCount {
					packetADDHOUSE_TYPES(sqlObject, rows)
					packetIndex = 0
					rows = make([]itemADDHOUSE_TYPES, 0)
				}
			}
		}
	}
	if packetIndex != 0 {
		packetADDHOUSE_TYPES(sqlObject, rows)
	}
	fmt.Println("Count:", index)
	return true
}

func workFileADDR_OBJ_TYPES(sqlObject *sql.DB, fileName string) bool {
	f, err := os.Open(filesPath + fileName)
	if err != nil {

	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)
	decoder := xml.NewDecoder(f)
	var packetIndex int64 = 0
	var index int64 = 0
	rows := make([]itemADDR_OBJ_TYPES, 0)
	for {
		tok, err := decoder.Token()
		if err != nil {

		}
		if tok == nil {
			break
		}
		switch tp := tok.(type) {
		case xml.StartElement:
			if tp.Name.Local == "ADDRESSOBJECTTYPE" {
				packetIndex++
				index++
				var row itemADDR_OBJ_TYPES
				err := decoder.DecodeElement(&row, &tp)
				if err != nil {
					return false
				}
				rows = append(rows, row)
				if packetIndex == packetCount {
					packetADDR_OBJ_TYPES(sqlObject, rows)
					packetIndex = 0
					rows = make([]itemADDR_OBJ_TYPES, 0)
				}
			}
		}
	}
	if packetIndex != 0 {
		packetADDR_OBJ_TYPES(sqlObject, rows)
	}
	fmt.Println("Count:", index)
	return true
}

func workFileAPARTMENT_TYPES(sqlObject *sql.DB, fileName string) bool {
	f, err := os.Open(filesPath + fileName)
	if err != nil {

	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)
	decoder := xml.NewDecoder(f)
	var packetIndex int64 = 0
	var index int64 = 0
	rows := make([]itemAPARTMENT_TYPES, 0)
	for {
		tok, err := decoder.Token()
		if err != nil {

		}
		if tok == nil {
			break
		}
		switch tp := tok.(type) {
		case xml.StartElement:
			if tp.Name.Local == "APARTMENTTYPE" {
				packetIndex++
				index++
				var row itemAPARTMENT_TYPES
				err := decoder.DecodeElement(&row, &tp)
				if err != nil {
					return false
				}
				rows = append(rows, row)
				if packetIndex == packetCount {
					packetAPARTMENT_TYPES(sqlObject, rows)
					packetIndex = 0
					rows = make([]itemAPARTMENT_TYPES, 0)
				}
			}
		}
	}
	if packetIndex != 0 {
		packetAPARTMENT_TYPES(sqlObject, rows)
	}
	fmt.Println("Count:", index)
	return true
}

func workFileHOUSE_TYPES(sqlObject *sql.DB, fileName string) bool {
	f, err := os.Open(filesPath + fileName)
	if err != nil {

	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)
	decoder := xml.NewDecoder(f)
	var packetIndex int64 = 0
	var index int64 = 0
	rows := make([]itemHOUSE_TYPES, 0)
	for {
		tok, err := decoder.Token()
		if err != nil {

		}
		if tok == nil {
			break
		}
		switch tp := tok.(type) {
		case xml.StartElement:
			if tp.Name.Local == "HOUSETYPE" {
				packetIndex++
				index++
				var row itemHOUSE_TYPES
				err := decoder.DecodeElement(&row, &tp)
				if err != nil {
					return false
				}
				rows = append(rows, row)
				if packetIndex == packetCount {
					packetHOUSE_TYPES(sqlObject, rows)
					packetIndex = 0
					rows = make([]itemHOUSE_TYPES, 0)
				}
			}
		}
	}
	if packetIndex != 0 {
		packetHOUSE_TYPES(sqlObject, rows)
	}
	fmt.Println("Count:", index)
	return true
}

func workFileNORMATIVE_DOCS_KINDS(sqlObject *sql.DB, fileName string) bool {
	f, err := os.Open(filesPath + fileName)
	if err != nil {

	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)
	decoder := xml.NewDecoder(f)
	var packetIndex int64 = 0
	var index int64 = 0
	rows := make([]itemNORMATIVE_DOCS_KINDS, 0)
	for {
		tok, err := decoder.Token()
		if err != nil {

		}
		if tok == nil {
			break
		}
		switch tp := tok.(type) {
		case xml.StartElement:
			if tp.Name.Local == "NDOCKIND" {
				packetIndex++
				index++
				var row itemNORMATIVE_DOCS_KINDS
				err := decoder.DecodeElement(&row, &tp)
				if err != nil {
					return false
				}
				rows = append(rows, row)
				if packetIndex == packetCount {
					packetNORMATIVE_DOCS_KINDS(sqlObject, rows)
					packetIndex = 0
					rows = make([]itemNORMATIVE_DOCS_KINDS, 0)
				}
			}
		}
	}
	if packetIndex != 0 {
		packetNORMATIVE_DOCS_KINDS(sqlObject, rows)
	}
	fmt.Println("Count:", index)
	return true
}

func workFileNORMATIVE_DOCS_TYPES(sqlObject *sql.DB, fileName string) bool {
	f, err := os.Open(filesPath + fileName)
	if err != nil {

	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)
	decoder := xml.NewDecoder(f)
	var packetIndex int64 = 0
	var index int64 = 0
	rows := make([]itemNORMATIVE_DOCS_TYPES, 0)
	for {
		tok, err := decoder.Token()
		if err != nil {

		}
		if tok == nil {
			break
		}
		switch tp := tok.(type) {
		case xml.StartElement:
			if tp.Name.Local == "NDOCTYPE" {
				packetIndex++
				index++
				var row itemNORMATIVE_DOCS_TYPES
				err := decoder.DecodeElement(&row, &tp)
				if err != nil {
					return false
				}
				rows = append(rows, row)
				if packetIndex == packetCount {
					packetNORMATIVE_DOCS_TYPES(sqlObject, rows)
					packetIndex = 0
					rows = make([]itemNORMATIVE_DOCS_TYPES, 0)
				}
			}
		}
	}
	if packetIndex != 0 {
		packetNORMATIVE_DOCS_TYPES(sqlObject, rows)
	}
	fmt.Println("Count:", index)
	return true
}

func workFileOBJECT_LEVELS(sqlObject *sql.DB, fileName string) bool {
	f, err := os.Open(filesPath + fileName)
	if err != nil {

	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)
	decoder := xml.NewDecoder(f)
	var packetIndex int64 = 0
	var index int64 = 0
	rows := make([]itemOBJECT_LEVELS, 0)
	for {
		tok, err := decoder.Token()
		if err != nil {

		}
		if tok == nil {
			break
		}
		switch tp := tok.(type) {
		case xml.StartElement:
			if tp.Name.Local == "OBJECTLEVEL" {
				packetIndex++
				index++
				var row itemOBJECT_LEVELS
				err := decoder.DecodeElement(&row, &tp)
				if err != nil {
					return false
				}
				rows = append(rows, row)
				if packetIndex == packetCount {
					packetOBJECT_LEVELS(sqlObject, rows)
					packetIndex = 0
					rows = make([]itemOBJECT_LEVELS, 0)
				}
			}
		}
	}
	if packetIndex != 0 {
		packetOBJECT_LEVELS(sqlObject, rows)
	}
	fmt.Println("Count:", index)
	return true
}

func workFileOPERATION_TYPES(sqlObject *sql.DB, fileName string) bool {
	f, err := os.Open(filesPath + fileName)
	if err != nil {

	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)
	decoder := xml.NewDecoder(f)
	var packetIndex int64 = 0
	var index int64 = 0
	rows := make([]itemOPERATION_TYPES, 0)
	for {
		tok, err := decoder.Token()
		if err != nil {

		}
		if tok == nil {
			break
		}
		switch tp := tok.(type) {
		case xml.StartElement:
			if tp.Name.Local == "OPERATIONTYPE" {
				packetIndex++
				index++
				var row itemOPERATION_TYPES
				err := decoder.DecodeElement(&row, &tp)
				if err != nil {
					return false
				}
				rows = append(rows, row)
				if packetIndex == packetCount {
					packetOPERATION_TYPES(sqlObject, rows)
					packetIndex = 0
					rows = make([]itemOPERATION_TYPES, 0)
				}
			}
		}
	}
	if packetIndex != 0 {
		packetOPERATION_TYPES(sqlObject, rows)
	}
	fmt.Println("Count:", index)
	return true
}

func workFilePARAM_TYPES(sqlObject *sql.DB, fileName string) bool {
	f, err := os.Open(filesPath + fileName)
	if err != nil {

	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)
	decoder := xml.NewDecoder(f)
	var packetIndex int64 = 0
	var index int64 = 0
	rows := make([]itemPARAM_TYPES, 0)
	for {
		tok, err := decoder.Token()
		if err != nil {

		}
		if tok == nil {
			break
		}
		switch tp := tok.(type) {
		case xml.StartElement:
			if tp.Name.Local == "PARAMTYPE" {
				packetIndex++
				index++
				var row itemPARAM_TYPES
				err := decoder.DecodeElement(&row, &tp)
				if err != nil {
					return false
				}
				rows = append(rows, row)
				if packetIndex == packetCount {
					packetPARAM_TYPES(sqlObject, rows)
					packetIndex = 0
					rows = make([]itemPARAM_TYPES, 0)
				}
			}
		}
	}
	if packetIndex != 0 {
		packetPARAM_TYPES(sqlObject, rows)
	}
	fmt.Println("Count:", index)
	return true
}

func workFileROOM_TYPES(sqlObject *sql.DB, fileName string) bool {
	f, err := os.Open(filesPath + fileName)
	if err != nil {

	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)
	decoder := xml.NewDecoder(f)
	var packetIndex int64 = 0
	var index int64 = 0
	rows := make([]itemROOM_TYPES, 0)
	for {
		tok, err := decoder.Token()
		if err != nil {

		}
		if tok == nil {
			break
		}
		switch tp := tok.(type) {
		case xml.StartElement:
			if tp.Name.Local == "ROOMTYPE" {
				packetIndex++
				index++
				var row itemROOM_TYPES
				err := decoder.DecodeElement(&row, &tp)
				if err != nil {
					return false
				}
				rows = append(rows, row)
				if packetIndex == packetCount {
					packetROOM_TYPES(sqlObject, rows)
					packetIndex = 0
					rows = make([]itemROOM_TYPES, 0)
				}
			}
		}
	}
	if packetIndex != 0 {
		packetROOM_TYPES(sqlObject, rows)
	}
	fmt.Println("Count:", index)
	return true
}

func setEnv() {
	readFile, err := os.Open(".env")
	if err != nil {
		panic(err)
	}
	defer func(readFile *os.File) {
		err := readFile.Close()
		if err != nil {

		}
	}(readFile)
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var lines []string
	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}
	for _, line := range lines {
		var list = strings.Split(line, "=")
		if len(list) == 2 {
			switch list[0] {
			case "sourcePath":
				sourcePath = list[1]
			case "filesPath":
				filesPath = list[1]
			case "databaseIP":
				databaseIP = list[1]
			case "databasePort":
				databasePort = list[1]
			case "databaseName":
				databaseName = list[1]
			case "databaseScheme":
				databaseScheme = list[1]
			case "databaseUser":
				databaseUser = list[1]
			case "databasePassword":
				databasePassword = list[1]
			case "loadFull":
				loadFull, _ = strconv.ParseBool(list[1])
			case "packetCount":
				packetCount, _ = strconv.ParseInt(list[1], 0, 8)
			}
		}
	}
}

func checkFileExists(filePath string) bool {
	_, err := os.Open(filePath)
	return err == nil

}

func DownloadFile(filepath string, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)
	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer func(out *os.File) {
		err := out.Close()
		if err != nil {

		}
	}(out)
	_, err = io.Copy(out, resp.Body)
	return err
}

func workFolder() {
	err := filepath.Walk(sourcePath,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() == false {
				err := os.Rename(path, filesPath+info.Name())
				if err != nil {
					return err
				}
			}
			return nil
		})
	if err != nil {
		log.Println(err)
	}
	psqlInfo := getDbConnectString()
	sqlObject, err := sql.Open("postgres", psqlInfo)
	CheckError(err)
	defer func(sqlObject *sql.DB) {
		err := sqlObject.Close()
		if err != nil {

		}
	}(sqlObject)
	err = sqlObject.Ping()
	CheckError(err)

	f, err := os.Open(filesPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	files, err := f.Readdir(0)
	if err != nil {
		fmt.Println(err)
		return
	}
	allStartTime := time.Now()
	for _, v := range files {
		if v.IsDir() == false {
			var fileName = v.Name()
			fmt.Println(fileName)
			startTime := time.Now()
			if strings.Index(fileName, "AS_ADDR_OBJ_2") != -1 {
				//Данные
				workFileADDR_OBJ(sqlObject, fileName)
			} else if strings.Index(fileName, "AS_ADDR_OBJ_DIVISION_2") != -1 {
				workFileADDR_OBJ_DIVISION(sqlObject, fileName)
			} else if strings.Index(fileName, "AS_ADM_HIERARCHY_2") != -1 {
				workFileADM_HIERARCHY(sqlObject, fileName)
			} else if strings.Index(fileName, "AS_APARTMENTS_2") != -1 {
				workFileAPARTMENTS(sqlObject, fileName)
			} else if strings.Index(fileName, "AS_CARPLACES_2") != -1 {
				workFileCARPLACES(sqlObject, fileName)
			} else if strings.Index(fileName, "AS_HOUSES_2") != -1 {
				workFileHOUSES(sqlObject, fileName)
			} else if strings.Index(fileName, "AS_MUN_HIERARCHY_2") != -1 {
				workFileMUN_HIERARCHY(sqlObject, fileName)
			} else if strings.Index(fileName, "AS_ROOMS_2") != -1 {
				workFileROOMS(sqlObject, fileName)
			} else if strings.Index(fileName, "AS_STEADS_2") != -1 {
				workFileSTEADS(sqlObject, fileName)
			} else if strings.Index(fileName, "AS_ADDHOUSE_TYPES_2") != -1 {
				//Справочники
				workFileADDHOUSE_TYPES(sqlObject, fileName)
			} else if strings.Index(fileName, "AS_ADDR_OBJ_TYPES_2") != -1 {
				workFileADDR_OBJ_TYPES(sqlObject, fileName)
			} else if strings.Index(fileName, "AS_APARTMENT_TYPES_2") != -1 {
				workFileAPARTMENT_TYPES(sqlObject, fileName)
			} else if strings.Index(fileName, "AS_HOUSE_TYPES_2") != -1 {
				workFileHOUSE_TYPES(sqlObject, fileName)
			} else if strings.Index(fileName, "AS_NORMATIVE_DOCS_KINDS_2") != -1 {
				workFileNORMATIVE_DOCS_KINDS(sqlObject, fileName)
			} else if strings.Index(fileName, "AS_NORMATIVE_DOCS_TYPES_2") != -1 {
				workFileNORMATIVE_DOCS_TYPES(sqlObject, fileName)
			} else if strings.Index(fileName, "AS_OBJECT_LEVELS_2") != -1 {
				workFileOBJECT_LEVELS(sqlObject, fileName)
			} else if strings.Index(fileName, "AS_OPERATION_TYPES_2") != -1 {
				workFileOPERATION_TYPES(sqlObject, fileName)
			} else if strings.Index(fileName, "AS_PARAM_TYPES_2") != -1 {
				workFilePARAM_TYPES(sqlObject, fileName)
			} else if strings.Index(fileName, "AS_ROOM_TYPES_2") != -1 {
				workFileROOM_TYPES(sqlObject, fileName)
			}
			finishTime := time.Now()
			fmt.Println("Time:", finishTime.Sub(startTime).Seconds())
		}
	}
	allFinishTime := time.Now()
	fmt.Println("All Time:", allFinishTime.Sub(allStartTime).Seconds())
}

func unpackZip(fileName string) {
	archive, err := zip.OpenReader(fileName)
	if err != nil {
		panic(err)
	}
	for _, f := range archive.File {
		filePath := filepath.Join(sourcePath, f.Name)
		if f.FileInfo().IsDir() {
			if err := os.MkdirAll(filePath, os.ModePerm); err != nil {
				panic(err)
			}
			continue
		}
		if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
			panic(err)
		}
		dstFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			panic(err)
		}
		srcFile, err := f.Open()
		if err != nil {
			panic(err)
		}
		if _, err := io.Copy(dstFile, srcFile); err != nil {
			panic(err)
		}
		err = dstFile.Close()
		if err != nil {
			return
		}
		err = srcFile.Close()
		if err != nil {
			return
		}
	}
	defer func(archive *zip.ReadCloser) {
		err := archive.Close()
		if err != nil {

		}
	}(archive)
}

func main() {
	setEnv()

	startDate, _ := time.Parse(time.DateOnly, "2024-07-26")
	for startDate.Before(time.Now()) {
		startDate = startDate.AddDate(0, 0, 1)
		filePath := "https://fias-file.nalog.ru/downloads/" + startDate.Format("2006.01.02") + "/gar_delta_xml.zip"
		fileZip := sourcePath + startDate.Format("2006.01.02") + ".zip"
		err := DownloadFile(fileZip, filePath)
		if err != nil {
			return
		}
		f, err := os.Stat(fileZip)
		if err != nil {
			return
		}
		if f.Size() != 162 {
			unpackZip(fileZip)
			err = os.Remove(fileZip)
			if err != nil {
				return
			}
			workFolder()
			err = filepath.Walk(sourcePath,
				func(path string, info os.FileInfo, err error) error {
					if err != nil {
						return err
					}
					if info.IsDir() == false {
						err = os.Remove(path)
					}
					if err != nil {
						return err
					}
					return nil
				})
			err = filepath.Walk(filesPath,
				func(path string, info os.FileInfo, err error) error {
					if err != nil {
						return err
					}
					err = os.Remove(path)
					if err != nil {
						return err
					}
					return nil
				})
			panic("")
		} else {
			err := os.Remove(fileZip)
			if err != nil {
				return
			}
		}
	}
}
