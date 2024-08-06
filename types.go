package main

type itemADD_OBJ struct {
	ID         int64  `xml:"ID,attr"`
	OBJECTID   int64  `xml:"OBJECTID,attr"`
	OBJECTGUID string `xml:"OBJECTGUID,attr"`
	CHANGEID   int64  `xml:"CHANGEID,attr"`
	NAME       string `xml:"NAME,attr"`
	TYPENAME   string `xml:"TYPENAME,attr"`
	LEVEL      int64  `xml:"LEVEL,attr"`
	OPERTYPEID int64  `xml:"OPERTYPEID,attr"`
	PREVID     int64  `xml:"PREVID,attr"`
	NEXTID     int64  `xml:"NEXTID,attr"`
	UPDATEDATE string `xml:"UPDATEDATE,attr"`
	STARTDATE  string `xml:"STARTDATE,attr"`
	ENDDATE    string `xml:"ENDDATE,attr"`
	ISACTUAL   string `xml:"ISACTUAL,attr"`
	ISACTIVE   string `xml:"ISACTIVE,attr"`
}

type itemADDR_OBJ_DIVISION struct {
	ID       int64 `xml:"ID,attr"`
	PARENTID int64 `xml:"PARENTID,attr"`
	CHILDID  int64 `xml:"CHILDID,attr"`
	CHANGEID int64 `xml:"CHANGEID,attr"`
}

type itemADM_HIERARCHY struct {
	ID          int64  `xml:"ID,attr"`
	OBJECTID    int64  `xml:"OBJECTID,attr"`
	PARENTOBJID int64  `xml:"PARENTOBJID,attr"`
	CHANGEID    int64  `xml:"CHANGEID,attr"`
	REGIONCODE  string `xml:"REGIONCODE,attr"`
	AREACODE    string `xml:"AREACODE,attr"`
	CITYCODE    string `xml:"CITYCODE,attr"`
	PLACECODE   string `xml:"PLACECODE,attr"`
	PLANCODE    string `xml:"PLANCODE,attr"`
	STREETCODE  string `xml:"STREETCODE,attr"`
	PREVID      int64  `xml:"PREVID,attr"`
	NEXTID      int64  `xml:"NEXTID,attr"`
	UPDATEDATE  string `xml:"UPDATEDATE,attr"`
	STARTDATE   string `xml:"STARTDATE,attr"`
	ENDDATE     string `xml:"ENDDATE,attr"`
	ISACTIVE    string `xml:"ISACTIVE,attr"`
	PATH        string `xml:"PATH,attr"`
}

type itemAPARTMENTS struct {
	ID         int64  `xml:"ID,attr"`
	OBJECTID   int64  `xml:"OBJECTID,attr"`
	OBJECTGUID string `xml:"OBJECTGUID,attr"`
	CHANGEID   int64  `xml:"CHANGEID,attr"`
	NUMBER     string `xml:"NUMBER,attr"`
	APARTTYPE  int64  `xml:"APARTTYPE,attr"`
	OPERTYPEID int64  `xml:"OPERTYPEID,attr"`
	PREVID     int64  `xml:"PREVID,attr"`
	NEXTID     int64  `xml:"NEXTID,attr"`
	UPDATEDATE string `xml:"UPDATEDATE,attr"`
	STARTDATE  string `xml:"STARTDATE,attr"`
	ENDDATE    string `xml:"ENDDATE,attr"`
	ISACTUAL   string `xml:"ISACTUAL,attr"`
	ISACTIVE   string `xml:"ISACTIVE,attr"`
}

type itemCARPLACES struct {
	ID         int64  `xml:"ID,attr"`
	OBJECTID   int64  `xml:"OBJECTID,attr"`
	OBJECTGUID string `xml:"OBJECTGUID,attr"`
	CHANGEID   int64  `xml:"CHANGEID,attr"`
	NUMBER     string `xml:"NUMBER,attr"`
	OPERTYPEID int64  `xml:"OPERTYPEID,attr"`
	PREVID     int64  `xml:"PREVID,attr"`
	NEXTID     int64  `xml:"NEXTID,attr"`
	UPDATEDATE string `xml:"UPDATEDATE,attr"`
	STARTDATE  string `xml:"STARTDATE,attr"`
	ENDDATE    string `xml:"ENDDATE,attr"`
	ISACTUAL   string `xml:"ISACTUAL,attr"`
	ISACTIVE   string `xml:"ISACTIVE,attr"`
}

type itemCHANGE_HISTORY struct {
	CHANGEID    int64  `xml:"CHANGEID,attr"`
	OBJECTID    int64  `xml:"OBJECTID,attr"`
	ADROBJECTID string `xml:"ADROBJECTID,attr"`
	OPERTYPEID  string `xml:"OPERTYPEID,attr"`
	NDOCID      string `xml:"NDOCID,attr"`
	CHANGEDATE  string `xml:"CHANGEDATE,attr"`
}

type itemHOUSES struct {
	ID         int64  `xml:"ID,attr"`
	OBJECTID   int64  `xml:"OBJECTID,attr"`
	OBJECTGUID string `xml:"OBJECTGUID,attr"`
	CHANGEID   int64  `xml:"CHANGEID,attr"`
	HOUSENUM   string `xml:"HOUSENUM,attr"`
	ADDNUM1    string `xml:"ADDNUM1,attr"`
	ADDNUM2    string `xml:"ADDNUM2,attr"`
	HOUSETYPE  int64  `xml:"HOUSETYPE,attr"`
	ADDTYPE1   int64  `xml:"ADDTYPE1,attr"`
	ADDTYPE2   int64  `xml:"ADDTYPE2,attr"`
	OPERTYPEID int64  `xml:"OPERTYPEID,attr"`
	PREVID     int64  `xml:"PREVID,attr"`
	NEXTID     int64  `xml:"NEXTID,attr"`
	UPDATEDATE string `xml:"UPDATEDATE,attr"`
	STARTDATE  string `xml:"STARTDATE,attr"`
	ENDDATE    string `xml:"ENDDATE,attr"`
	ISACTUAL   string `xml:"ISACTUAL,attr"`
	ISACTIVE   string `xml:"ISACTIVE,attr"`
}

type itemMUN_HIERARCHY struct {
	ID          int64  `xml:"ID,attr"`
	OBJECTID    int64  `xml:"OBJECTID,attr"`
	PARENTOBJID int64  `xml:"PARENTOBJID,attr"`
	CHANGEID    int64  `xml:"CHANGEID,attr"`
	OKTMO       string `xml:"OKTMO,attr"`
	PREVID      int64  `xml:"PREVID,attr"`
	NEXTID      int64  `xml:"NEXTID,attr"`
	UPDATEDATE  string `xml:"UPDATEDATE,attr"`
	STARTDATE   string `xml:"STARTDATE,attr"`
	ENDDATE     string `xml:"ENDDATE,attr"`
	ISACTIVE    string `xml:"ISACTIVE,attr"`
	PATH        string `xml:"PATH,attr"`
}

type itemROOMS struct {
	ID         int64  `xml:"ID,attr"`
	OBJECTID   int64  `xml:"OBJECTID,attr"`
	OBJECTGUID string `xml:"OBJECTGUID,attr"`
	CHANGEID   int64  `xml:"CHANGEID,attr"`
	NUMBER     string `xml:"NUMBER,attr"`
	ROOMTYPE   int64  `xml:"ROOMTYPE,attr"`
	OPERTYPEID int64  `xml:"OPERTYPEID,attr"`
	PREVID     int64  `xml:"PREVID,attr"`
	NEXTID     int64  `xml:"NEXTID,attr"`
	UPDATEDATE string `xml:"UPDATEDATE,attr"`
	STARTDATE  string `xml:"STARTDATE,attr"`
	ENDDATE    string `xml:"ENDDATE,attr"`
	ISACTUAL   string `xml:"ISACTUAL,attr"`
	ISACTIVE   string `xml:"ISACTIVE,attr"`
}

type itemSTEADS struct {
	ID         int64  `xml:"ID,attr"`
	OBJECTID   int64  `xml:"OBJECTID,attr"`
	OBJECTGUID string `xml:"OBJECTGUID,attr"`
	CHANGEID   int64  `xml:"CHANGEID,attr"`
	NUMBER     string `xml:"NUMBER,attr"`
	OPERTYPEID int64  `xml:"OPERTYPEID,attr"`
	PREVID     int64  `xml:"PREVID,attr"`
	NEXTID     int64  `xml:"NEXTID,attr"`
	UPDATEDATE string `xml:"UPDATEDATE,attr"`
	STARTDATE  string `xml:"STARTDATE,attr"`
	ENDDATE    string `xml:"ENDDATE,attr"`
	ISACTUAL   string `xml:"ISACTUAL,attr"`
	ISACTIVE   string `xml:"ISACTIVE,attr"`
}

// Справочники
type itemADDHOUSE_TYPES struct {
	ID         int64  `xml:"ID,attr"`
	NAME       string `xml:"NAME,attr"`
	SHORTNAME  string `xml:"SHORTNAME,attr"`
	DESC       string `xml:"DESC,attr"`
	ISACTIVE   string `xml:"ISACTIVE,attr"`
	UPDATEDATE string `xml:"UPDATEDATE,attr"`
	STARTDATE  string `xml:"STARTDATE,attr"`
	ENDDATE    string `xml:"ENDDATE,attr"`
}

type itemADDR_OBJ_TYPES struct {
	ID         int64  `xml:"ID,attr"`
	LEVEL      int64  `xml:"LEVEL,attr"`
	SHORTNAME  string `xml:"SHORTNAME,attr"`
	NAME       string `xml:"NAME,attr"`
	DESC       string `xml:"DESC,attr"`
	UPDATEDATE string `xml:"UPDATEDATE,attr"`
	STARTDATE  string `xml:"STARTDATE,attr"`
	ENDDATE    string `xml:"ENDDATE,attr"`
	ISACTIVE   string `xml:"ISACTIVE,attr"`
}

type itemAPARTMENT_TYPES struct {
	ID         int64  `xml:"ID,attr"`
	NAME       string `xml:"NAME,attr"`
	SHORTNAME  string `xml:"SHORTNAME,attr"`
	DESC       string `xml:"DESC,attr"`
	UPDATEDATE string `xml:"UPDATEDATE,attr"`
	STARTDATE  string `xml:"STARTDATE,attr"`
	ENDDATE    string `xml:"ENDDATE,attr"`
	ISACTIVE   string `xml:"ISACTIVE,attr"`
}

type itemHOUSE_TYPES struct {
	ID         int64  `xml:"ID,attr"`
	NAME       string `xml:"NAME,attr"`
	SHORTNAME  string `xml:"SHORTNAME,attr"`
	DESC       string `xml:"DESC,attr"`
	UPDATEDATE string `xml:"UPDATEDATE,attr"`
	STARTDATE  string `xml:"STARTDATE,attr"`
	ENDDATE    string `xml:"ENDDATE,attr"`
	ISACTIVE   string `xml:"ISACTIVE,attr"`
}

type itemNORMATIVE_DOCS_KINDS struct {
	ID   int64  `xml:"ID,attr"`
	NAME string `xml:"NAME,attr"`
}

type itemNORMATIVE_DOCS_TYPES struct {
	ID        int64  `xml:"ID,attr"`
	NAME      string `xml:"NAME,attr"`
	STARTDATE string `xml:"STARTDATE,attr"`
	ENDDATE   string `xml:"ENDDATE,attr"`
}

type itemOBJECT_LEVELS struct {
	LEVEL      int64  `xml:"LEVEL,attr"`
	NAME       string `xml:"NAME,attr"`
	SHORTNAME  string `xml:"SHORTNAME,attr"`
	UPDATEDATE string `xml:"UPDATEDATE,attr"`
	STARTDATE  string `xml:"STARTDATE,attr"`
	ENDDATE    string `xml:"ENDDATE,attr"`
	ISACTIVE   string `xml:"ISACTIVE,attr"`
}

type itemOPERATION_TYPES struct {
	ID         int64  `xml:"ID,attr"`
	NAME       string `xml:"NAME,attr"`
	SHORTNAME  string `xml:"SHORTNAME,attr"`
	DESC       string `xml:"DESC,attr"`
	UPDATEDATE string `xml:"UPDATEDATE,attr"`
	STARTDATE  string `xml:"STARTDATE,attr"`
	ENDDATE    string `xml:"ENDDATE,attr"`
	ISACTIVE   string `xml:"ISACTIVE,attr"`
}

type itemROOM_TYPES struct {
	ID         int64  `xml:"ID,attr"`
	NAME       string `xml:"NAME,attr"`
	SHORTNAME  string `xml:"SHORTNAME,attr"`
	DESC       string `xml:"DESC,attr"`
	UPDATEDATE string `xml:"UPDATEDATE,attr"`
	STARTDATE  string `xml:"STARTDATE,attr"`
	ENDDATE    string `xml:"ENDDATE,attr"`
	ISACTIVE   string `xml:"ISACTIVE,attr"`
}

type itemPARAM_TYPES struct {
	ID         int64  `xml:"ID,attr"`
	NAME       string `xml:"NAME,attr"`
	CODE       string `xml:"CODE,attr"`
	DESC       string `xml:"DESC,attr"`
	UPDATEDATE string `xml:"UPDATEDATE,attr"`
	STARTDATE  string `xml:"STARTDATE,attr"`
	ENDDATE    string `xml:"ENDDATE,attr"`
	ISACTIVE   string `xml:"ISACTIVE,attr"`
}
