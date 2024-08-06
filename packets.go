package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"slices"
	"strconv"
	"strings"
)

func packetADDR_OBJ_Insert(dataRow itemADD_OBJ) string {
	return fmt.Sprintf(
		"insert into "+databaseScheme+".addr_obj (id, objectid, objectguid, changeid, name, typename, level, opertypeid, previd, nextid, updatedate, startdate, enddate, isactual, isactive) values (%d, %d, '%s', %d, '%s', '%s', %d, %d, %d, %d, '%s', '%s', '%s', '%s', '%s');",
		dataRow.ID,
		dataRow.OBJECTID,
		dataRow.OBJECTGUID,
		dataRow.CHANGEID,
		strings.Replace(dataRow.NAME, "'", "", -1),
		dataRow.TYPENAME,
		dataRow.LEVEL,
		dataRow.OPERTYPEID,
		dataRow.PREVID,
		dataRow.NEXTID,
		dataRow.UPDATEDATE,
		dataRow.STARTDATE,
		dataRow.ENDDATE,
		dataRow.ISACTUAL,
		dataRow.ISACTIVE,
	)
}

func packetADDR_OBJ_Update(dataRow itemADD_OBJ) string {
	return fmt.Sprintf("update "+databaseScheme+".addr_obj set objectid=%d, objectguid='%s', changeid=%d, name='%s', typename='%s', level=%d, opertypeid=%d, previd=%d, nextid=%d, updatedate='%s', startdate='%s', enddate='%s', isactual='%s', isactive='%s' where id = %d;",
		dataRow.OBJECTID,
		dataRow.OBJECTGUID,
		dataRow.CHANGEID,
		strings.Replace(dataRow.NAME, "'", "", -1),
		dataRow.TYPENAME,
		dataRow.LEVEL,
		dataRow.OPERTYPEID,
		dataRow.PREVID,
		dataRow.NEXTID,
		dataRow.UPDATEDATE,
		dataRow.STARTDATE,
		dataRow.ENDDATE,
		dataRow.ISACTUAL,
		dataRow.ISACTIVE,
		dataRow.ID,
	)
}

func packetADDR_OBJ(sqlObject *sql.DB, data []itemADD_OBJ) bool {
	var sqlText = ""
	if loadFull {
		for _, dataRow := range data {
			//Создаем
			sqlText = sqlText + packetADDR_OBJ_Insert(dataRow)
		}
		if strings.TrimSpace(sqlText) != "" {
			_, err := sqlObject.Exec(sqlText)
			if err != nil {
				panic(err)
			}

		}
	} else {
		idxs := make([]string, 0)
		idxsExists := make([]int64, 0)
		for _, dataRow := range data {
			idxs = append(idxs, strconv.FormatInt(dataRow.ID, 10))
		}
		rowsExists, err := sqlObject.Query("select ID from " + databaseScheme + ".addr_obj where id in (" + strings.Join(idxs, ",") + ")")
		CheckError(err)
		defer func(rowsExists *sql.Rows) {
			err := rowsExists.Close()
			if err != nil {

			}
		}(rowsExists)
		for rowsExists.Next() {
			var id int64
			err = rowsExists.Scan(&id)
			idxsExists = append(idxsExists, id)
			CheckError(err)
		}
		CheckError(err)
		for _, dataRow := range data {
			if slices.Contains(idxsExists, dataRow.ID) == false {
				sqlText = sqlText + packetADDR_OBJ_Insert(dataRow)
			} else {
				sqlText = sqlText + packetADDR_OBJ_Update(dataRow)
			}
		}
		if strings.TrimSpace(sqlText) != "" {
			_, err = sqlObject.Exec(sqlText)
			if err != nil {
				panic(err)
			}

		}
	}
	return true
}

func packetADDR_OBJ_DIVISION_Insert(dataRow itemADDR_OBJ_DIVISION) string {
	return fmt.Sprintf(
		"insert into "+databaseScheme+".addr_obj_division (id, parentid, childid, changeid) values (%d, %d, %d, %d);",
		dataRow.ID,
		dataRow.PARENTID,
		dataRow.CHILDID,
		dataRow.CHANGEID,
	)
}

func packetADDR_OBJ_DIVISION_Update(dataRow itemADDR_OBJ_DIVISION) string {
	return fmt.Sprintf("update "+databaseScheme+".addr_obj_division set parentid=%d, childid=%d, changeid=%d where id = %d;",
		dataRow.PARENTID,
		dataRow.CHILDID,
		dataRow.CHANGEID,
		dataRow.ID,
	)
}

func packetADDR_OBJ_DIVISION(sqlObject *sql.DB, data []itemADDR_OBJ_DIVISION) bool {
	var sqlText = ""
	if loadFull {
		for _, dataRow := range data {
			//Создаем
			sqlText = sqlText + packetADDR_OBJ_DIVISION_Insert(dataRow)
		}
		if strings.TrimSpace(sqlText) != "" {
			_, err := sqlObject.Exec(sqlText)
			if err != nil {
				panic(err)
			}
		}
	} else {
		idxs := make([]string, 0)
		idxsExists := make([]int64, 0)
		for _, dataRow := range data {
			idxs = append(idxs, strconv.FormatInt(dataRow.ID, 10))
		}
		rowsExists, err := sqlObject.Query("select ID from " + databaseScheme + ".addr_obj_division where id in (" + strings.Join(idxs, ",") + ")")
		CheckError(err)
		defer func(rowsExists *sql.Rows) {
			err := rowsExists.Close()
			if err != nil {

			}
		}(rowsExists)
		for rowsExists.Next() {
			var id int64
			err = rowsExists.Scan(&id)
			idxsExists = append(idxsExists, id)
			CheckError(err)
		}
		CheckError(err)
		for _, dataRow := range data {
			if slices.Contains(idxsExists, dataRow.ID) == false {
				sqlText = sqlText + packetADDR_OBJ_DIVISION_Insert(dataRow)
			} else {
				sqlText = sqlText + packetADDR_OBJ_DIVISION_Update(dataRow)
			}
		}
		if strings.TrimSpace(sqlText) != "" {
			_, err = sqlObject.Exec(sqlText)
			if err != nil {
				panic(err)
			}
		}
	}
	return true
}

func packetADM_HIERARCHY_Insert(dataRow itemADM_HIERARCHY) string {
	return fmt.Sprintf(
		"insert into "+databaseScheme+".adm_hierarchy (id, objectid, parentobjid, changeid, regioncode, areacode, citycode, placecode, plancode, streetcode, previd, nextid, updatedate, startdate, enddate, isactive, path) values (%d, %d, %d, %d, '%s', '%s', '%s', '%s', '%s', '%s', %d, %d, '%s', '%s', '%s', '%s', '%s');",
		dataRow.ID,
		dataRow.OBJECTID,
		dataRow.PARENTOBJID,
		dataRow.CHANGEID,
		dataRow.REGIONCODE,
		dataRow.AREACODE,
		dataRow.CITYCODE,
		dataRow.PLACECODE,
		dataRow.PLANCODE,
		dataRow.STREETCODE,
		dataRow.PREVID,
		dataRow.NEXTID,
		dataRow.UPDATEDATE,
		dataRow.STARTDATE,
		dataRow.ENDDATE,
		dataRow.ISACTIVE,
		dataRow.PATH,
	)
}

func packetADM_HIERARCHY_Update(dataRow itemADM_HIERARCHY) string {
	return fmt.Sprintf("update "+databaseScheme+".adm_hierarchy set objectid=%d, parentobjid=%d, changeid=%d, regioncode='%s', areacode='%s', citycode='%s', placecode='%s', plancode='%s', streetcode='%s', previd=%d, nextid=%d, updatedate='%s', startdate='%s', enddate='%s', isactive='%s', path='%s' where id = %d;",
		dataRow.OBJECTID,
		dataRow.PARENTOBJID,
		dataRow.CHANGEID,
		dataRow.REGIONCODE,
		dataRow.AREACODE,
		dataRow.CITYCODE,
		dataRow.PLACECODE,
		dataRow.PLANCODE,
		dataRow.STREETCODE,
		dataRow.PREVID,
		dataRow.NEXTID,
		dataRow.UPDATEDATE,
		dataRow.STARTDATE,
		dataRow.ENDDATE,
		dataRow.ISACTIVE,
		dataRow.PATH,
		dataRow.ID,
	)
}

func packetADM_HIERARCHY(sqlObject *sql.DB, data []itemADM_HIERARCHY) bool {
	var sqlText = ""
	if loadFull {
		for _, dataRow := range data {
			sqlText = sqlText + packetADM_HIERARCHY_Insert(dataRow)
		}
		if strings.TrimSpace(sqlText) != "" {
			_, err := sqlObject.Exec(sqlText)
			if err != nil {
				panic(err)
			}
		}
	} else {
		idxs := make([]string, 0)
		idxsExists := make([]int64, 0)
		for _, dataRow := range data {
			idxs = append(idxs, strconv.FormatInt(dataRow.ID, 10))
		}
		rowsExists, err := sqlObject.Query("select ID from " + databaseScheme + ".adm_hierarchy where id in (" + strings.Join(idxs, ",") + ")")
		CheckError(err)
		defer func(rowsExists *sql.Rows) {
			err := rowsExists.Close()
			if err != nil {

			}
		}(rowsExists)
		for rowsExists.Next() {
			var id int64
			err = rowsExists.Scan(&id)
			idxsExists = append(idxsExists, id)
			CheckError(err)
		}
		CheckError(err)
		for _, dataRow := range data {
			if slices.Contains(idxsExists, dataRow.ID) == false {
				sqlText = sqlText + packetADM_HIERARCHY_Insert(dataRow)
			} else {
				sqlText = sqlText + packetADM_HIERARCHY_Update(dataRow)
			}
		}
		if strings.TrimSpace(sqlText) != "" {
			_, err = sqlObject.Exec(sqlText)
			if err != nil {
				panic(err)
			}
		}
	}
	return true
}

func packetAPARTMENTS_Insert(dataRow itemAPARTMENTS) string {
	return fmt.Sprintf(
		"insert into "+databaseScheme+".apartments (id, objectid, objectguid, changeid, number, aparttype, opertypeid, previd, nextid, updatedate, startdate, enddate, isactual, isactive) values (%d, %d, '%s', %d, '%s', %d, %d, %d, %d, '%s', '%s', '%s', '%s', '%s');",
		dataRow.ID,
		dataRow.OBJECTID,
		dataRow.OBJECTGUID,
		dataRow.CHANGEID,
		dataRow.NUMBER,
		dataRow.APARTTYPE,
		dataRow.OPERTYPEID,
		dataRow.PREVID,
		dataRow.NEXTID,
		dataRow.UPDATEDATE,
		dataRow.STARTDATE,
		dataRow.ENDDATE,
		dataRow.ISACTUAL,
		dataRow.ISACTIVE,
	)
}

func packetAPARTMENTS_Update(dataRow itemAPARTMENTS) string {
	return fmt.Sprintf("update "+databaseScheme+".apartments set objectid=%d, objectguid='%s', changeid=%d, number='%s', aparttype=%d, opertypeid=%d, previd=%d, nextid=%d, updatedate='%s', startdate='%s', enddate='%s', isactual='%s', isactive='%s' where id = %d;",
		dataRow.OBJECTID,
		dataRow.OBJECTGUID,
		dataRow.CHANGEID,
		dataRow.NUMBER,
		dataRow.APARTTYPE,
		dataRow.OPERTYPEID,
		dataRow.PREVID,
		dataRow.NEXTID,
		dataRow.UPDATEDATE,
		dataRow.STARTDATE,
		dataRow.ENDDATE,
		dataRow.ISACTUAL,
		dataRow.ISACTIVE,
		dataRow.ID,
	)
}

func packetAPARTMENTS(sqlObject *sql.DB, data []itemAPARTMENTS) bool {
	var sqlText = ""
	if loadFull {
		for _, dataRow := range data {
			sqlText = sqlText + packetAPARTMENTS_Insert(dataRow)
		}
		if strings.TrimSpace(sqlText) != "" {
			_, err := sqlObject.Exec(sqlText)
			if err != nil {
				panic(err)
			}
		}
	} else {
		idxs := make([]string, 0)
		idxsExists := make([]int64, 0)
		for _, dataRow := range data {
			idxs = append(idxs, strconv.FormatInt(dataRow.ID, 10))
		}
		rowsExists, err := sqlObject.Query("select ID from " + databaseScheme + ".apartments where id in (" + strings.Join(idxs, ",") + ")")
		CheckError(err)
		defer func(rowsExists *sql.Rows) {
			err := rowsExists.Close()
			if err != nil {

			}
		}(rowsExists)
		for rowsExists.Next() {
			var id int64
			err = rowsExists.Scan(&id)
			idxsExists = append(idxsExists, id)
			CheckError(err)
		}
		CheckError(err)
		for _, dataRow := range data {
			if slices.Contains(idxsExists, dataRow.ID) == false {
				sqlText = sqlText + packetAPARTMENTS_Insert(dataRow)
			} else {
				sqlText = sqlText + packetAPARTMENTS_Update(dataRow)
			}
		}
		if strings.TrimSpace(sqlText) != "" {
			_, err = sqlObject.Exec(sqlText)
			if err != nil {
				panic(err)
			}
		}
	}
	return true
}

func packetCARPLACES_Insert(dataRow itemCARPLACES) string {
	return fmt.Sprintf(
		"insert into "+databaseScheme+".carplaces (id, objectid, objectguid, changeid, number, opertypeid, previd, nextid, updatedate, startdate, enddate, isactual, isactive) values (%d, %d, '%s', %d, '%s', %d, %d, %d, '%s', '%s', '%s', '%s', '%s');",
		dataRow.ID,
		dataRow.OBJECTID,
		dataRow.OBJECTGUID,
		dataRow.CHANGEID,
		dataRow.NUMBER,
		dataRow.OPERTYPEID,
		dataRow.PREVID,
		dataRow.NEXTID,
		dataRow.UPDATEDATE,
		dataRow.STARTDATE,
		dataRow.ENDDATE,
		dataRow.ISACTUAL,
		dataRow.ISACTIVE,
	)
}

func packetCARPLACES_Update(dataRow itemCARPLACES) string {
	return fmt.Sprintf("update "+databaseScheme+".carplaces set objectid=%d, objectguid='%s', changeid=%d, number='%s', opertypeid=%d, previd=%d, nextid=%d, updatedate='%s', startdate='%s', enddate='%s', isactual='%s', isactive='%s' where id = %d;",
		dataRow.OBJECTID,
		dataRow.OBJECTGUID,
		dataRow.CHANGEID,
		dataRow.NUMBER,
		dataRow.OPERTYPEID,
		dataRow.PREVID,
		dataRow.NEXTID,
		dataRow.UPDATEDATE,
		dataRow.STARTDATE,
		dataRow.ENDDATE,
		dataRow.ISACTUAL,
		dataRow.ISACTIVE,
		dataRow.ID,
	)
}

func packetCARPLACES(sqlObject *sql.DB, data []itemCARPLACES) bool {
	var sqlText = ""
	if loadFull {
		for _, dataRow := range data {
			sqlText = sqlText + packetCARPLACES_Insert(dataRow)
		}
		if strings.TrimSpace(sqlText) != "" {
			_, err := sqlObject.Exec(sqlText)
			if err != nil {
				panic(err)
			}
		}
	} else {
		idxs := make([]string, 0)
		idxsExists := make([]int64, 0)
		for _, dataRow := range data {
			idxs = append(idxs, strconv.FormatInt(dataRow.ID, 10))
		}
		rowsExists, err := sqlObject.Query("select ID from " + databaseScheme + ".carplaces where id in (" + strings.Join(idxs, ",") + ")")
		CheckError(err)
		defer func(rowsExists *sql.Rows) {
			err := rowsExists.Close()
			if err != nil {

			}
		}(rowsExists)
		for rowsExists.Next() {
			var id int64
			err = rowsExists.Scan(&id)
			idxsExists = append(idxsExists, id)
			CheckError(err)
		}
		CheckError(err)
		for _, dataRow := range data {
			if slices.Contains(idxsExists, dataRow.ID) == false {
				sqlText = sqlText + packetCARPLACES_Insert(dataRow)
			} else {
				sqlText = sqlText + packetCARPLACES_Update(dataRow)
			}
		}
		if strings.TrimSpace(sqlText) != "" {
			_, err = sqlObject.Exec(sqlText)
			if err != nil {
				panic(err)
			}
		}
	}
	return true
}

func packetHOUSES_Insert(dataRow itemHOUSES) string {
	return fmt.Sprintf(
		"insert into "+databaseScheme+".houses (id, objectid, objectguid, changeid, housenum, addnum1, addnum2, housetype, addtype1, addtype2, opertypeid, previd, nextid, updatedate, startdate, enddate, isactual, isactive) values (%d, %d, '%s', %d, '%s', '%s', '%s', %d, %d, %d, %d, %d, %d, '%s', '%s', '%s', '%s', '%s');",
		dataRow.ID,
		dataRow.OBJECTID,
		dataRow.OBJECTGUID,
		dataRow.CHANGEID,
		strings.Replace(dataRow.HOUSENUM, "'", "", -1),
		strings.Replace(dataRow.ADDNUM1, "'", "", -1),
		strings.Replace(dataRow.ADDNUM1, "'", "", -1),
		dataRow.HOUSETYPE,
		dataRow.ADDTYPE1,
		dataRow.ADDTYPE2,
		dataRow.OPERTYPEID,
		dataRow.PREVID,
		dataRow.NEXTID,
		dataRow.UPDATEDATE,
		dataRow.STARTDATE,
		dataRow.ENDDATE,
		dataRow.ISACTUAL,
		dataRow.ISACTIVE,
	)
}

func packetHOUSES_Update(dataRow itemHOUSES) string {
	return fmt.Sprintf("update "+databaseScheme+".houses set objectid=%d, objectguid='%s', changeid=%d, housenum='%s', addnum1='%s', addnum2='%s', housetype=%d, addtype1=%d, addtype2=%d, opertypeid=%d, previd=%d, nextid=%d, updatedate='%s', startdate='%s', enddate='%s', isactual='%s', isactive='%s' where id = %d;",
		dataRow.OBJECTID,
		dataRow.OBJECTGUID,
		dataRow.CHANGEID,
		strings.Replace(dataRow.HOUSENUM, "'", "", -1),
		strings.Replace(dataRow.ADDNUM1, "'", "", -1),
		strings.Replace(dataRow.ADDNUM1, "'", "", -1),
		dataRow.HOUSETYPE,
		dataRow.ADDTYPE1,
		dataRow.ADDTYPE2,
		dataRow.OPERTYPEID,
		dataRow.PREVID,
		dataRow.NEXTID,
		dataRow.UPDATEDATE,
		dataRow.STARTDATE,
		dataRow.ENDDATE,
		dataRow.ISACTUAL,
		dataRow.ISACTIVE,
		dataRow.ID,
	)
}

func packetHOUSES(sqlObject *sql.DB, data []itemHOUSES) bool {
	var sqlText = ""
	if loadFull {
		for _, dataRow := range data {
			sqlText = sqlText + packetHOUSES_Insert(dataRow)
		}
		if strings.TrimSpace(sqlText) != "" {
			_, err := sqlObject.Exec(sqlText)
			if err != nil {
				panic(err)
			}
		}
	} else {
		idxs := make([]string, 0)
		idxsExists := make([]int64, 0)
		for _, dataRow := range data {
			idxs = append(idxs, strconv.FormatInt(dataRow.ID, 10))
		}
		rowsExists, err := sqlObject.Query("select ID from " + databaseScheme + ".houses where id in (" + strings.Join(idxs, ",") + ")")
		CheckError(err)
		defer func(rowsExists *sql.Rows) {
			err := rowsExists.Close()
			if err != nil {

			}
		}(rowsExists)
		for rowsExists.Next() {
			var id int64
			err = rowsExists.Scan(&id)
			idxsExists = append(idxsExists, id)
			CheckError(err)
		}
		CheckError(err)
		for _, dataRow := range data {
			if slices.Contains(idxsExists, dataRow.ID) == false {
				sqlText = sqlText + packetHOUSES_Insert(dataRow)
			} else {
				sqlText = sqlText + packetHOUSES_Update(dataRow)
			}
		}
		if strings.TrimSpace(sqlText) != "" {
			_, err = sqlObject.Exec(sqlText)
			if err != nil {
				panic(err)
			}
		}
	}
	return true
}

func packetMUN_HIERARCHY_Insert(dataRow itemMUN_HIERARCHY) string {
	return fmt.Sprintf(
		"insert into "+databaseScheme+".mun_hierarchy (id, objectid, parentobjid, changeid, oktmo, previd, nextid, updatedate, startdate, enddate, isactive, path) values (%d, %d, %d, %d, '%s', %d, %d, '%s', '%s', '%s', '%s', '%s');",
		dataRow.ID,
		dataRow.OBJECTID,
		dataRow.PARENTOBJID,
		dataRow.CHANGEID,
		dataRow.OKTMO,
		dataRow.PREVID,
		dataRow.NEXTID,
		dataRow.UPDATEDATE,
		dataRow.STARTDATE,
		dataRow.ENDDATE,
		dataRow.ISACTIVE,
		dataRow.PATH,
	)
}

func packetMUN_HIERARCHY_Update(dataRow itemMUN_HIERARCHY) string {
	return fmt.Sprintf("update "+databaseScheme+".mun_hierarchy set objectid=%d, parentobjid=%d, changeid=%d, oktmo='%s', previd=%d, nextid=%d, updatedate='%s', startdate='%s', enddate='%s', isactive='%s', path='%s' where id = %d;",
		dataRow.OBJECTID,
		dataRow.PARENTOBJID,
		dataRow.CHANGEID,
		dataRow.OKTMO,
		dataRow.PREVID,
		dataRow.NEXTID,
		dataRow.UPDATEDATE,
		dataRow.STARTDATE,
		dataRow.ENDDATE,
		dataRow.ISACTIVE,
		dataRow.PATH,
		dataRow.ID,
	)
}

func packetMUN_HIERARCHY(sqlObject *sql.DB, data []itemMUN_HIERARCHY) bool {
	var sqlText = ""
	if loadFull {
		for _, dataRow := range data {
			sqlText = sqlText + packetMUN_HIERARCHY_Insert(dataRow)
		}
		if strings.TrimSpace(sqlText) != "" {
			_, err := sqlObject.Exec(sqlText)
			if err != nil {
				panic(err)
			}
		}
	} else {
		idxs := make([]string, 0)
		idxsExists := make([]int64, 0)
		for _, dataRow := range data {
			idxs = append(idxs, strconv.FormatInt(dataRow.ID, 10))
		}
		rowsExists, err := sqlObject.Query("select ID from " + databaseScheme + ".mun_hierarchy where id in (" + strings.Join(idxs, ",") + ")")
		CheckError(err)
		defer func(rowsExists *sql.Rows) {
			err := rowsExists.Close()
			if err != nil {

			}
		}(rowsExists)
		for rowsExists.Next() {
			var id int64
			err = rowsExists.Scan(&id)
			idxsExists = append(idxsExists, id)
			CheckError(err)
		}
		CheckError(err)
		for _, dataRow := range data {
			if slices.Contains(idxsExists, dataRow.ID) == false {
				sqlText = sqlText + packetMUN_HIERARCHY_Insert(dataRow)
			} else {
				sqlText = sqlText + packetMUN_HIERARCHY_Update(dataRow)
			}
		}
		if strings.TrimSpace(sqlText) != "" {
			_, err = sqlObject.Exec(sqlText)
			if err != nil {
				panic(err)
			}
		}
	}
	return true
}

func packetROOMS_Insert(dataRow itemROOMS) string {
	return fmt.Sprintf(
		"insert into "+databaseScheme+".rooms (id, objectid, objectguid, changeid, number, roomtype, opertypeid, previd, nextid, updatedate, startdate, enddate, isactual, isactive) values (%d, %d, '%s', %d, '%s', %d, %d, %d, %d, '%s', '%s', '%s', '%s', '%s');",
		dataRow.ID,
		dataRow.OBJECTID,
		dataRow.OBJECTGUID,
		dataRow.CHANGEID,
		dataRow.NUMBER,
		dataRow.ROOMTYPE,
		dataRow.OPERTYPEID,
		dataRow.PREVID,
		dataRow.NEXTID,
		dataRow.UPDATEDATE,
		dataRow.STARTDATE,
		dataRow.ENDDATE,
		dataRow.ISACTUAL,
		dataRow.ISACTIVE,
	)
}

func packetROOMS_Update(dataRow itemROOMS) string {
	return fmt.Sprintf("update "+databaseScheme+".rooms set objectid=%d, objectguid='%s', changeid=%d, number='%s', roomtype=%d, opertypeid=%d, previd=%d, nextid=%d, updatedate='%s', startdate='%s', enddate='%s', isactual='%s', isactive='%s' where id = %d;",
		dataRow.OBJECTID,
		dataRow.OBJECTGUID,
		dataRow.CHANGEID,
		dataRow.NUMBER,
		dataRow.ROOMTYPE,
		dataRow.OPERTYPEID,
		dataRow.PREVID,
		dataRow.NEXTID,
		dataRow.UPDATEDATE,
		dataRow.STARTDATE,
		dataRow.ENDDATE,
		dataRow.ISACTUAL,
		dataRow.ISACTIVE,
		dataRow.ID,
	)
}

func packetROOMS(sqlObject *sql.DB, data []itemROOMS) bool {
	var sqlText = ""
	if loadFull {
		for _, dataRow := range data {
			sqlText = sqlText + packetROOMS_Insert(dataRow)
		}
		if strings.TrimSpace(sqlText) != "" {
			_, err := sqlObject.Exec(sqlText)
			if err != nil {
				panic(err)
			}
		}
	} else {
		idxs := make([]string, 0)
		idxsExists := make([]int64, 0)
		for _, dataRow := range data {
			idxs = append(idxs, strconv.FormatInt(dataRow.ID, 10))
		}
		rowsExists, err := sqlObject.Query("select ID from " + databaseScheme + ".rooms where id in (" + strings.Join(idxs, ",") + ")")
		CheckError(err)
		defer func(rowsExists *sql.Rows) {
			err := rowsExists.Close()
			if err != nil {

			}
		}(rowsExists)
		for rowsExists.Next() {
			var id int64
			err = rowsExists.Scan(&id)
			idxsExists = append(idxsExists, id)
			CheckError(err)
		}
		CheckError(err)
		for _, dataRow := range data {
			if slices.Contains(idxsExists, dataRow.ID) == false {
				sqlText = sqlText + packetROOMS_Insert(dataRow)
			} else {
				sqlText = sqlText + packetROOMS_Update(dataRow)
			}
		}
		if strings.TrimSpace(sqlText) != "" {
			_, err = sqlObject.Exec(sqlText)
			if err != nil {
				panic(err)
			}
		}
	}
	return true
}

func packetSTEADS_Insert(dataRow itemSTEADS) string {
	return fmt.Sprintf(
		"insert into "+databaseScheme+".steads (id, objectid, objectguid, changeid, number, opertypeid, previd, nextid, updatedate, startdate, enddate, isactual, isactive) values (%d, %d, '%s', %d, '%s', %d, %d, %d, '%s', '%s', '%s', '%s', '%s');",
		dataRow.ID,
		dataRow.OBJECTID,
		dataRow.OBJECTGUID,
		dataRow.CHANGEID,
		dataRow.NUMBER,
		dataRow.OPERTYPEID,
		dataRow.PREVID,
		dataRow.NEXTID,
		dataRow.UPDATEDATE,
		dataRow.STARTDATE,
		dataRow.ENDDATE,
		dataRow.ISACTUAL,
		dataRow.ISACTIVE,
	)
}

func packetSTEADS_Update(dataRow itemSTEADS) string {
	return fmt.Sprintf("update "+databaseScheme+".steads set objectid=%d, objectguid='%s', changeid=%d, number='%s', opertypeid=%d, previd=%d, nextid=%d, updatedate='%s', startdate='%s', enddate='%s', isactual='%s', isactive='%s' where id = %d;",
		dataRow.OBJECTID,
		dataRow.OBJECTGUID,
		dataRow.CHANGEID,
		dataRow.NUMBER,
		dataRow.OPERTYPEID,
		dataRow.PREVID,
		dataRow.NEXTID,
		dataRow.UPDATEDATE,
		dataRow.STARTDATE,
		dataRow.ENDDATE,
		dataRow.ISACTUAL,
		dataRow.ISACTIVE,
		dataRow.ID,
	)
}

func packetSTEADS(sqlObject *sql.DB, data []itemSTEADS) bool {
	var sqlText = ""
	if loadFull {
		for _, dataRow := range data {
			sqlText = sqlText + packetSTEADS_Insert(dataRow)
		}
		if strings.TrimSpace(sqlText) != "" {
			_, err := sqlObject.Exec(sqlText)
			if err != nil {
				panic(err)
			}
		}
	} else {
		idxs := make([]string, 0)
		idxsExists := make([]int64, 0)
		for _, dataRow := range data {
			idxs = append(idxs, strconv.FormatInt(dataRow.ID, 10))
		}
		rowsExists, err := sqlObject.Query("select ID from " + databaseScheme + ".steads where id in (" + strings.Join(idxs, ",") + ")")
		CheckError(err)
		defer func(rowsExists *sql.Rows) {
			err := rowsExists.Close()
			if err != nil {

			}
		}(rowsExists)
		for rowsExists.Next() {
			var id int64
			err = rowsExists.Scan(&id)
			idxsExists = append(idxsExists, id)
			CheckError(err)
		}
		CheckError(err)
		for _, dataRow := range data {
			if slices.Contains(idxsExists, dataRow.ID) == false {
				sqlText = sqlText + packetSTEADS_Insert(dataRow)
			} else {
				sqlText = sqlText + packetSTEADS_Update(dataRow)
			}
		}
		if strings.TrimSpace(sqlText) != "" {
			_, err = sqlObject.Exec(sqlText)
			if err != nil {
				panic(err)
			}
		}
	}
	return true
}

// Справочники
func packetADDHOUSE_TYPES(sqlObject *sql.DB, data []itemADDHOUSE_TYPES) bool {
	var sqlText string = ""
	idxs := make([]string, 0)
	idxsExists := make([]int64, 0)
	for _, dataRow := range data {
		idxs = append(idxs, strconv.FormatInt(dataRow.ID, 10))
	}
	rowsExists, err := sqlObject.Query("select ID from " + databaseScheme + ".addhouse_types where id in (" + strings.Join(idxs, ",") + ")")
	CheckError(err)
	defer func(rowsExists *sql.Rows) {
		err := rowsExists.Close()
		if err != nil {

		}
	}(rowsExists)
	for rowsExists.Next() {
		var id int64
		err = rowsExists.Scan(&id)
		idxsExists = append(idxsExists, id)
		CheckError(err)
	}
	CheckError(err)
	for _, dataRow := range data {
		if slices.Contains(idxsExists, dataRow.ID) == false {
			//Создаем
			sqlText = sqlText + fmt.Sprintf(
				"insert into "+databaseScheme+".addhouse_types (id, name, shortname, \"desc\", isactive, updatedate, startdate, enddate) values (%d, '%s', '%s', '%s', %s, '%s', '%s', '%s');",
				dataRow.ID,
				dataRow.NAME,
				dataRow.SHORTNAME,
				dataRow.DESC,
				dataRow.ISACTIVE,
				dataRow.UPDATEDATE,
				dataRow.STARTDATE,
				dataRow.ENDDATE,
			)
		} else {
			//Изменяем
			sqlText = sqlText + fmt.Sprintf("update "+databaseScheme+".addhouse_types set name='%s', shortname='%s', \"desc\"='%s', isactive=%s, updatedate='%s', startdate='%s', enddate='%s' where id = %d;",
				dataRow.NAME,
				dataRow.SHORTNAME,
				dataRow.DESC,
				dataRow.ISACTIVE,
				dataRow.UPDATEDATE,
				dataRow.STARTDATE,
				dataRow.ENDDATE,
				dataRow.ID,
			)
		}
	}
	if strings.TrimSpace(sqlText) != "" {
		_, err = sqlObject.Exec(sqlText)
		if err != nil {
			panic(err)
		}

	}
	return true
}

func packetADDR_OBJ_TYPES(sqlObject *sql.DB, data []itemADDR_OBJ_TYPES) bool {
	var sqlText = ""
	idxs := make([]string, 0)
	idxsExists := make([]int64, 0)
	for _, dataRow := range data {
		idxs = append(idxs, strconv.FormatInt(dataRow.ID, 10))
	}
	rowsExists, err := sqlObject.Query("select ID from " + databaseScheme + ".addr_obj_types where id in (" + strings.Join(idxs, ",") + ")")
	CheckError(err)
	defer func(rowsExists *sql.Rows) {
		err := rowsExists.Close()
		if err != nil {

		}
	}(rowsExists)
	for rowsExists.Next() {
		var id int64
		err = rowsExists.Scan(&id)
		idxsExists = append(idxsExists, id)
		CheckError(err)
	}
	CheckError(err)
	for _, dataRow := range data {
		if slices.Contains(idxsExists, dataRow.ID) == false {
			//Создаем
			sqlText = sqlText + fmt.Sprintf(
				"insert into "+databaseScheme+".addr_obj_types (id, level, name, shortname, \"desc\", updatedate, startdate, enddate, isactive) values (%d, %d, '%s', '%s', '%s', '%s', '%s', '%s', %s);",
				dataRow.ID,
				dataRow.LEVEL,
				dataRow.NAME,
				dataRow.SHORTNAME,
				dataRow.DESC,
				dataRow.UPDATEDATE,
				dataRow.STARTDATE,
				dataRow.ENDDATE,
				dataRow.ISACTIVE,
			)
		} else {
			//Изменяем
			sqlText = sqlText + fmt.Sprintf("update "+databaseScheme+".addr_obj_types set  level=%d, name='%s', shortname='%s', \"desc\"='%s', updatedate='%s', startdate='%s', enddate='%s', isactive=%s where id = %d;",
				dataRow.LEVEL,
				dataRow.NAME,
				dataRow.SHORTNAME,
				dataRow.DESC,
				dataRow.UPDATEDATE,
				dataRow.STARTDATE,
				dataRow.ENDDATE,
				dataRow.ISACTIVE,
				dataRow.ID,
			)
		}
	}
	if strings.TrimSpace(sqlText) != "" {
		_, err = sqlObject.Exec(sqlText)
		if err != nil {
			panic(err)
		}

	}
	return true
}

func packetAPARTMENT_TYPES(sqlObject *sql.DB, data []itemAPARTMENT_TYPES) bool {
	var sqlText = ""
	idxs := make([]string, 0)
	idxsExists := make([]int64, 0)
	for _, dataRow := range data {
		idxs = append(idxs, strconv.FormatInt(dataRow.ID, 10))
	}
	rowsExists, err := sqlObject.Query("select ID from " + databaseScheme + ".apartment_types where id in (" + strings.Join(idxs, ",") + ")")
	CheckError(err)
	defer func(rowsExists *sql.Rows) {
		err := rowsExists.Close()
		if err != nil {

		}
	}(rowsExists)
	for rowsExists.Next() {
		var id int64
		err = rowsExists.Scan(&id)
		idxsExists = append(idxsExists, id)
		CheckError(err)
	}
	CheckError(err)
	for _, dataRow := range data {
		if slices.Contains(idxsExists, dataRow.ID) == false {
			//Создаем
			sqlText = sqlText + fmt.Sprintf(
				"insert into "+databaseScheme+".apartment_types (id, name, shortname, \"desc\", isactive, updatedate, startdate, enddate) values (%d, '%s', '%s', '%s', %s, '%s', '%s', '%s');",
				dataRow.ID,
				dataRow.NAME,
				dataRow.SHORTNAME,
				dataRow.DESC,
				dataRow.ISACTIVE,
				dataRow.UPDATEDATE,
				dataRow.STARTDATE,
				dataRow.ENDDATE,
			)
		} else {
			//Изменяем
			sqlText = sqlText + fmt.Sprintf("update "+databaseScheme+".apartment_types set name='%s', shortname='%s', \"desc\"='%s', isactive=%s, updatedate='%s', startdate='%s', enddate='%s' where id = %d;",
				dataRow.NAME,
				dataRow.SHORTNAME,
				dataRow.DESC,
				dataRow.ISACTIVE,
				dataRow.UPDATEDATE,
				dataRow.STARTDATE,
				dataRow.ENDDATE,
				dataRow.ID,
			)
		}
	}
	if strings.TrimSpace(sqlText) != "" {
		_, err = sqlObject.Exec(sqlText)
		if err != nil {
			panic(err)
		}

	}
	return true
}

func packetHOUSE_TYPES(sqlObject *sql.DB, data []itemHOUSE_TYPES) bool {
	var sqlText = ""
	idxs := make([]string, 0)
	idxsExists := make([]int64, 0)
	for _, dataRow := range data {
		idxs = append(idxs, strconv.FormatInt(dataRow.ID, 10))
	}
	rowsExists, err := sqlObject.Query("select ID from " + databaseScheme + ".house_types where id in (" + strings.Join(idxs, ",") + ")")
	CheckError(err)
	defer func(rowsExists *sql.Rows) {
		err := rowsExists.Close()
		if err != nil {

		}
	}(rowsExists)
	for rowsExists.Next() {
		var id int64
		err = rowsExists.Scan(&id)
		idxsExists = append(idxsExists, id)
		CheckError(err)
	}
	CheckError(err)
	for _, dataRow := range data {
		if slices.Contains(idxsExists, dataRow.ID) == false {
			//Создаем
			sqlText = sqlText + fmt.Sprintf(
				"insert into "+databaseScheme+".house_types (id, name, shortname, \"desc\", isactive, updatedate, startdate, enddate) values (%d, '%s', '%s', '%s', %s, '%s', '%s', '%s');",
				dataRow.ID,
				dataRow.NAME,
				dataRow.SHORTNAME,
				dataRow.DESC,
				dataRow.ISACTIVE,
				dataRow.UPDATEDATE,
				dataRow.STARTDATE,
				dataRow.ENDDATE,
			)
		} else {
			//Изменяем
			sqlText = sqlText + fmt.Sprintf("update "+databaseScheme+".house_types set name='%s', shortname='%s', \"desc\"='%s', isactive=%s, updatedate='%s', startdate='%s', enddate='%s' where id = %d;",
				dataRow.NAME,
				dataRow.SHORTNAME,
				dataRow.DESC,
				dataRow.ISACTIVE,
				dataRow.UPDATEDATE,
				dataRow.STARTDATE,
				dataRow.ENDDATE,
				dataRow.ID,
			)
		}
	}
	if strings.TrimSpace(sqlText) != "" {
		_, err = sqlObject.Exec(sqlText)
		if err != nil {
			panic(err)
		}

	}
	return true
}

func packetNORMATIVE_DOCS_KINDS(sqlObject *sql.DB, data []itemNORMATIVE_DOCS_KINDS) bool {
	var sqlText = ""
	idxs := make([]string, 0)
	idxsExists := make([]int64, 0)
	for _, dataRow := range data {
		idxs = append(idxs, strconv.FormatInt(dataRow.ID, 10))
	}
	rowsExists, err := sqlObject.Query("select ID from " + databaseScheme + ".normative_docs_kinds where id in (" + strings.Join(idxs, ",") + ")")
	CheckError(err)
	defer func(rowsExists *sql.Rows) {
		err := rowsExists.Close()
		if err != nil {

		}
	}(rowsExists)
	for rowsExists.Next() {
		var id int64
		err = rowsExists.Scan(&id)
		idxsExists = append(idxsExists, id)
		CheckError(err)
	}
	CheckError(err)
	for _, dataRow := range data {
		if slices.Contains(idxsExists, dataRow.ID) == false {
			//Создаем
			sqlText = sqlText + fmt.Sprintf(
				"insert into "+databaseScheme+".normative_docs_kinds (id, name) values (%d, '%s');",
				dataRow.ID,
				dataRow.NAME,
			)
		} else {
			//Изменяем
			sqlText = sqlText + fmt.Sprintf("update "+databaseScheme+".normative_docs_kinds set name='%s' where id = %d;",
				dataRow.NAME,
				dataRow.ID,
			)
		}
	}
	if strings.TrimSpace(sqlText) != "" {
		_, err = sqlObject.Exec(sqlText)
		if err != nil {
			panic(err)
		}

	}
	return true
}

func packetNORMATIVE_DOCS_TYPES(sqlObject *sql.DB, data []itemNORMATIVE_DOCS_TYPES) bool {
	var sqlText = ""
	idxs := make([]string, 0)
	idxsExists := make([]int64, 0)
	for _, dataRow := range data {
		idxs = append(idxs, strconv.FormatInt(dataRow.ID, 10))
	}
	rowsExists, err := sqlObject.Query("select ID from " + databaseScheme + ".normative_docs_types where id in (" + strings.Join(idxs, ",") + ")")
	CheckError(err)
	defer func(rowsExists *sql.Rows) {
		err := rowsExists.Close()
		if err != nil {

		}
	}(rowsExists)
	for rowsExists.Next() {
		var id int64
		err = rowsExists.Scan(&id)
		idxsExists = append(idxsExists, id)
		CheckError(err)
	}
	CheckError(err)
	for _, dataRow := range data {
		if slices.Contains(idxsExists, dataRow.ID) == false {
			//Создаем
			sqlText = sqlText + fmt.Sprintf(
				"insert into "+databaseScheme+".normative_docs_types (id, name, startdate, enddate) values (%d, '%s', '%s', '%s');",
				dataRow.ID,
				dataRow.NAME,
				dataRow.STARTDATE,
				dataRow.ENDDATE,
			)
		} else {
			//Изменяем
			sqlText = sqlText + fmt.Sprintf("update "+databaseScheme+".normative_docs_types set name='%s', startdate='%s', enddate='%s' where id = %d;",
				dataRow.NAME,
				dataRow.STARTDATE,
				dataRow.ENDDATE,
				dataRow.ID,
			)
		}
	}
	if strings.TrimSpace(sqlText) != "" {
		_, err = sqlObject.Exec(sqlText)
		if err != nil {
			panic(err)
		}

	}
	return true
}

func packetOBJECT_LEVELS(sqlObject *sql.DB, data []itemOBJECT_LEVELS) bool {
	var sqlText = ""
	idxs := make([]string, 0)
	idxsExists := make([]int64, 0)
	for _, dataRow := range data {
		idxs = append(idxs, strconv.FormatInt(dataRow.LEVEL, 10))
	}
	rowsExists, err := sqlObject.Query("select level from " + databaseScheme + ".object_levels where level in (" + strings.Join(idxs, ",") + ")")
	CheckError(err)
	defer func(rowsExists *sql.Rows) {
		err := rowsExists.Close()
		if err != nil {

		}
	}(rowsExists)
	for rowsExists.Next() {
		var level int64
		err = rowsExists.Scan(&level)
		idxsExists = append(idxsExists, level)
		CheckError(err)
	}
	CheckError(err)
	for _, dataRow := range data {
		if slices.Contains(idxsExists, dataRow.LEVEL) == false {
			//Создаем
			sqlText = sqlText + fmt.Sprintf(
				"insert into "+databaseScheme+".object_levels (level, name, startdate, enddate, updatedate, isactive) values (%d, '%s', '%s', '%s', '%s', %s);",
				dataRow.LEVEL,
				dataRow.NAME,
				dataRow.STARTDATE,
				dataRow.ENDDATE,
				dataRow.UPDATEDATE,
				dataRow.ISACTIVE,
			)
		} else {
			//Изменяем
			sqlText = sqlText + fmt.Sprintf("update "+databaseScheme+".object_levels set name='%s', startdate='%s', enddate='%s', updatedate='%s', isactive=%s where level = %d;",
				dataRow.NAME,
				dataRow.STARTDATE,
				dataRow.ENDDATE,
				dataRow.UPDATEDATE,
				dataRow.ISACTIVE,
				dataRow.LEVEL,
			)
		}
	}
	if strings.TrimSpace(sqlText) != "" {
		_, err = sqlObject.Exec(sqlText)
		if err != nil {
			panic(err)
		}

	}
	return true
}

func packetOPERATION_TYPES(sqlObject *sql.DB, data []itemOPERATION_TYPES) bool {
	var sqlText = ""
	idxs := make([]string, 0)
	idxsExists := make([]int64, 0)
	for _, dataRow := range data {
		idxs = append(idxs, strconv.FormatInt(dataRow.ID, 10))
	}
	rowsExists, err := sqlObject.Query("select ID from " + databaseScheme + ".operation_types where id in (" + strings.Join(idxs, ",") + ")")
	CheckError(err)
	defer func(rowsExists *sql.Rows) {
		err := rowsExists.Close()
		if err != nil {

		}
	}(rowsExists)
	for rowsExists.Next() {
		var id int64
		err = rowsExists.Scan(&id)
		idxsExists = append(idxsExists, id)
		CheckError(err)
	}
	CheckError(err)
	for _, dataRow := range data {
		if slices.Contains(idxsExists, dataRow.ID) == false {
			//Создаем
			sqlText = sqlText + fmt.Sprintf(
				"insert into "+databaseScheme+".operation_types (id, name, isactive, updatedate, startdate, enddate) values (%d, '%s', %s, '%s', '%s', '%s');",
				dataRow.ID,
				dataRow.NAME,
				dataRow.ISACTIVE,
				dataRow.UPDATEDATE,
				dataRow.STARTDATE,
				dataRow.ENDDATE,
			)
		} else {
			//Изменяем
			sqlText = sqlText + fmt.Sprintf("update "+databaseScheme+".operation_types set name='%s', isactive=%s, updatedate='%s', startdate='%s', enddate='%s' where id = %d;",
				dataRow.NAME,
				dataRow.ISACTIVE,
				dataRow.UPDATEDATE,
				dataRow.STARTDATE,
				dataRow.ENDDATE,
				dataRow.ID,
			)
		}
	}
	if strings.TrimSpace(sqlText) != "" {
		_, err = sqlObject.Exec(sqlText)
		if err != nil {
			panic(err)
		}

	}
	return true
}

func packetPARAM_TYPES(sqlObject *sql.DB, data []itemPARAM_TYPES) bool {
	var sqlText = ""
	idxs := make([]string, 0)
	idxsExists := make([]int64, 0)
	for _, dataRow := range data {
		idxs = append(idxs, strconv.FormatInt(dataRow.ID, 10))
	}
	rowsExists, err := sqlObject.Query("select ID from " + databaseScheme + ".param_types where id in (" + strings.Join(idxs, ",") + ")")
	CheckError(err)
	defer func(rowsExists *sql.Rows) {
		err := rowsExists.Close()
		if err != nil {

		}
	}(rowsExists)
	for rowsExists.Next() {
		var id int64
		err = rowsExists.Scan(&id)
		idxsExists = append(idxsExists, id)
		CheckError(err)
	}
	CheckError(err)
	for _, dataRow := range data {
		if slices.Contains(idxsExists, dataRow.ID) == false {
			//Создаем
			sqlText = sqlText + fmt.Sprintf(
				"insert into "+databaseScheme+".param_types (id, name, \"desc\", code, isactive, updatedate, startdate, enddate) values (%d, '%s', '%s', '%s', %s, '%s', '%s', '%s');",
				dataRow.ID,
				dataRow.NAME,
				dataRow.DESC,
				dataRow.CODE,
				dataRow.ISACTIVE,
				dataRow.UPDATEDATE,
				dataRow.STARTDATE,
				dataRow.ENDDATE,
			)
		} else {
			//Изменяем
			sqlText = sqlText + fmt.Sprintf("update "+databaseScheme+".param_types set name='%s', \"desc\"='%s',code='%s', isactive=%s, updatedate='%s', startdate='%s', enddate='%s' where id = %d;",
				dataRow.NAME,
				dataRow.DESC,
				dataRow.CODE,
				dataRow.ISACTIVE,
				dataRow.UPDATEDATE,
				dataRow.STARTDATE,
				dataRow.ENDDATE,
				dataRow.ID,
			)
		}
	}
	if strings.TrimSpace(sqlText) != "" {
		_, err = sqlObject.Exec(sqlText)
		if err != nil {
			panic(err)
		}

	}
	return true
}

func packetROOM_TYPES(sqlObject *sql.DB, data []itemROOM_TYPES) bool {
	var sqlText = ""
	idxs := make([]string, 0)
	idxsExists := make([]int64, 0)
	for _, dataRow := range data {
		idxs = append(idxs, strconv.FormatInt(dataRow.ID, 10))
	}
	rowsExists, err := sqlObject.Query("select ID from " + databaseScheme + ".room_types where id in (" + strings.Join(idxs, ",") + ")")
	CheckError(err)
	defer func(rowsExists *sql.Rows) {
		err := rowsExists.Close()
		if err != nil {

		}
	}(rowsExists)
	for rowsExists.Next() {
		var id int64
		err = rowsExists.Scan(&id)
		idxsExists = append(idxsExists, id)
		CheckError(err)
	}
	CheckError(err)
	for _, dataRow := range data {
		if slices.Contains(idxsExists, dataRow.ID) == false {
			//Создаем
			sqlText = sqlText + fmt.Sprintf(
				"insert into "+databaseScheme+".room_types (id, name, \"desc\", isactive, updatedate, startdate, enddate) values (%d, '%s', '%s', %s, '%s', '%s', '%s');",
				dataRow.ID,
				dataRow.NAME,
				dataRow.DESC,
				dataRow.ISACTIVE,
				dataRow.UPDATEDATE,
				dataRow.STARTDATE,
				dataRow.ENDDATE,
			)
		} else {
			//Изменяем
			sqlText = sqlText + fmt.Sprintf("update "+databaseScheme+".room_types set name='%s', \"desc\"='%s', isactive=%s, updatedate='%s', startdate='%s', enddate='%s' where id = %d;",
				dataRow.NAME,
				dataRow.DESC,
				dataRow.ISACTIVE,
				dataRow.UPDATEDATE,
				dataRow.STARTDATE,
				dataRow.ENDDATE,
				dataRow.ID,
			)
		}
	}
	if strings.TrimSpace(sqlText) != "" {
		_, err = sqlObject.Exec(sqlText)
		if err != nil {
			panic(err)
		}

	}
	return true
}
