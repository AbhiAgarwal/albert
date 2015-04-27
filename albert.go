package main

import (
	"bytes"
	"crypto/tls"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type SSR_GET_CLASSES_REQ struct {
	CLASS_SEARCH_REQUEST struct {
		TypeRequest      string `xml:"type"`
		LanguageCd       string `xml:"languageCd"`
		STRM             string `xml:"STRM"`
		SUBJECT          string `xml:"SUBJECT"`
		INSTITUTION      string `xml:"INSTITUTION"`
		CAMPUS           string `xml:"CAMPUS"`
		ACAD_CAREER      string `xml:"ACAD_CAREER"`
		SSR_EXACT_MATCH1 string `xml:"SSR_EXACT_MATCH1"`
		CATALOG_NBR      string `xml:"CATALOG_NBR"`
	} `xml:"CLASS_SEARCH_REQUEST"`
}

type SSRGETCLASSESRESP struct {
	SEARCHRESULT struct {
		ERRORWARNTEXT   string `xml:"ERROR_WARN_TEXT" json:"ERROR_WARN_TEXT"`
		SSRCLASSCOUNT   string `xml:"SSR_CLASS_COUNT" json:"SSR_CLASS_COUNT"`
		SSRCOURSECOUNT  string `xml:"SSR_COURSE_COUNT" json:"SSR_COURSE_COUNT"`
		SSRERRLMTEXCEED string `xml:"SSR_ERR_LMT_EXCEED" json:"SSR_ERR_LMT_EXCEED"`
		SSRWRNLMTEXCEED string `xml:"SSR_WRN_LMT_EXCEED" json:"SSR_WRN_LMT_EXCEED"`
		SUBJECTS        struct {
			SUBJECT []struct {
				ACADCAREER         string `xml:"ACAD_CAREER" json:"ACAD_CAREER"`
				ACADCAREERLOVDescr string `xml:"ACAD_CAREER_LOVDescr" json:"ACAD_CAREER_LOVDescr"`
				CATALOGNBR         string `xml:"CATALOG_NBR" json:"CATALOG_NBR"`
				CLASSESSUMMARY     struct {
					CLASSSUMMARY struct {
						CAMPUS                 string `xml:"CAMPUS" json:"CAMPUS"`
						CAMPUSLOVDescr         string `xml:"CAMPUS_LOVDescr" json:"CAMPUS_LOVDescr"`
						CATALOGNBR             string `xml:"CATALOG_NBR" json:"CATALOG_NBR"`
						CLASSESMEETINGPATTERNS struct {
							CLASSMEETINGPATTERN struct {
								BLDGCD           string `xml:"BLDG_CD" json:"BLDG_CD"`
								CLASSINSTRUCTORS struct {
									CLASSINSTRUCTOR struct {
										CLASSMTGNBR             string `xml:"CLASS_MTG_NBR" json:"CLASS_MTG_NBR"`
										CLASSSECTION            string `xml:"CLASS_SECTION" json:"CLASS_SECTION"`
										CRSEID                  string `xml:"CRSE_ID" json:"CRSE_ID"`
										CRSEOFFERNBR            string `xml:"CRSE_OFFER_NBR" json:"CRSE_OFFER_NBR"`
										EMPLID                  string `xml:"EMPLID" json:"EMPLID"`
										ENCRYPTEDEMPLID         string `xml:"ENCRYPTED_EMPLID" json:"ENCRYPTED_EMPLID"`
										FIRSTNAME               string `xml:"FIRST_NAME" json:"FIRST_NAME"`
										INSTRROLE               string `xml:"INSTR_ROLE" json:"INSTR_ROLE"`
										INSTRROLELOVDescr       string `xml:"INSTR_ROLE_LOVDescr" json:"INSTR_ROLE_LOVDescr"`
										LASTNAME                string `xml:"LAST_NAME" json:"LAST_NAME"`
										NAMEDISPLAY             string `xml:"NAME_DISPLAY" json:"NAME_DISPLAY"`
										SCHEDPRINTINSTR         string `xml:"SCHED_PRINT_INSTR" json:"SCHED_PRINT_INSTR"`
										SCHEDPRINTINSTRLOVDescr string `xml:"SCHED_PRINT_INSTR_LOVDescr" json:"SCHED_PRINT_INSTR_LOVDescr"`
										SESSIONCODE             string `xml:"SESSION_CODE" json:"SESSION_CODE"`
										SESSIONCODELOVDescr     string `xml:"SESSION_CODE_LOVDescr" json:"SESSION_CODE_LOVDescr"`
										STRM                    string `xml:"STRM" json:"STRM"`
										STRMLOVDescr            string `xml:"STRM_LOVDescr" json:"STRM_LOVDescr"`
									} `xml:"CLASS_INSTRUCTOR" json:"CLASS_INSTRUCTOR"`
								} `xml:"CLASS_INSTRUCTORS" json:"CLASS_INSTRUCTORS"`
								CLASSMTGNBR         string `xml:"CLASS_MTG_NBR" json:"CLASS_MTG_NBR"`
								CLASSSECTION        string `xml:"CLASS_SECTION" json:"CLASS_SECTION"`
								CRSEID              string `xml:"CRSE_ID" json:"CRSE_ID"`
								CRSEOFFERNBR        string `xml:"CRSE_OFFER_NBR" json:"CRSE_OFFER_NBR"`
								CRSTOPICID          string `xml:"CRS_TOPIC_ID" json:"CRS_TOPIC_ID"`
								CRSTOPICIDLOVDescr  string `xml:"CRS_TOPIC_ID_LOVDescr" json:"CRS_TOPIC_ID_LOVDescr"`
								DESCR               string `xml:"DESCR" json:"DESCR"`
								ENDDT               string `xml:"END_DT" json:"END_DT"`
								FACILITYID          string `xml:"FACILITY_ID" json:"FACILITY_ID"`
								FACILITYIDLOVDescr  string `xml:"FACILITY_ID_LOVDescr" json:"FACILITY_ID_LOVDescr"`
								FRI                 string `xml:"FRI" json:"FRI"`
								MEETINGTIMEEND      string `xml:"MEETING_TIME_END" json:"MEETING_TIME_END"`
								MEETINGTIMESTART    string `xml:"MEETING_TIME_START" json:"MEETING_TIME_START"`
								MON                 string `xml:"MON" json:"MON"`
								SAT                 string `xml:"SAT" json:"SAT"`
								SCCLATITUDE         string `xml:"SCC_LATITUDE" json:"SCC_LATITUDE"`
								SCCLONGITUDE        string `xml:"SCC_LONGITUDE" json:"SCC_LONGITUDE"`
								SESSIONCODE         string `xml:"SESSION_CODE" json:"SESSION_CODE"`
								SESSIONCODELOVDescr string `xml:"SESSION_CODE_LOVDescr" json:"SESSION_CODE_LOVDescr"`
								SSRINSTRLONG        string `xml:"SSR_INSTR_LONG" json:"SSR_INSTR_LONG"`
								SSRMTGDTLONG        string `xml:"SSR_MTG_DT_LONG" json:"SSR_MTG_DT_LONG"`
								SSRMTGLOCLONG       string `xml:"SSR_MTG_LOC_LONG" json:"SSR_MTG_LOC_LONG"`
								SSRMTGSCHEDLONG     string `xml:"SSR_MTG_SCHED_LONG" json:"SSR_MTG_SCHED_LONG"`
								SSRTOPICLONG        string `xml:"SSR_TOPIC_LONG" json:"SSR_TOPIC_LONG"`
								STARTDT             string `xml:"START_DT" json:"START_DT"`
								STNDMTGPAT          string `xml:"STND_MTG_PAT" json:"STND_MTG_PAT"`
								STRM                string `xml:"STRM" json:"STRM"`
								SUN                 string `xml:"SUN" json:"SUN"`
								THURS               string `xml:"THURS" json:"THURS"`
								TUES                string `xml:"TUES" json:"TUES"`
								WED                 string `xml:"WED" json:"WED"`
							} `xml:"CLASS_MEETING_PATTERN" json:"CLASS_MEETING_PATTERN"`
						} `xml:"CLASSES_MEETING_PATTERNS" json:"CLASSES_MEETING_PATTERNS"`
						CLASSNBR                string   `xml:"CLASS_NBR" json:"CLASS_NBR"`
						CLASSSECTION            string   `xml:"CLASS_SECTION" json:"CLASS_SECTION"`
						CLASSTOPIC              string   `xml:"CLASS_TOPIC" json:"CLASS_TOPIC"`
						COMBINEDSECTION         string   `xml:"COMBINED_SECTION" json:"COMBINED_SECTION"`
						COMBINEDSECTIONLOVDescr string   `xml:"COMBINED_SECTION_LOVDescr" json:"COMBINED_SECTION_LOVDescr"`
						CRSEID                  string   `xml:"CRSE_ID" json:"CRSE_ID"`
						CRSEIDLOVDescr          string   `xml:"CRSE_ID_LOVDescr" json:"CRSE_ID_LOVDescr"`
						CRSEOFFERNBR            string   `xml:"CRSE_OFFER_NBR" json:"CRSE_OFFER_NBR"`
						INSTITUTION             string   `xml:"INSTITUTION" json:"INSTITUTION"`
						INSTRUCTIONMODE         string   `xml:"INSTRUCTION_MODE" json:"INSTRUCTION_MODE"`
						INSTRUCTIONMODELOVDescr string   `xml:"INSTRUCTION_MODE_LOVDescr" json:"INSTRUCTION_MODE_LOVDescr"`
						SCHEDULEPRINT           string   `xml:"SCHEDULE_PRINT" json:"SCHEDULE_PRINT"`
						SCHEDULEPRINTLOVDescr   string   `xml:"SCHEDULE_PRINT_LOVDescr" json:"SCHEDULE_PRINT_LOVDescr"`
						SESSIONCODE             string   `xml:"SESSION_CODE" json:"SESSION_CODE"`
						SESSIONCODELOVDescr     string   `xml:"SESSION_CODE_LOVDescr" json:"SESSION_CODE_LOVDescr"`
						SSRCLASSNAMELONG        string   `xml:"SSR_CLASSNAME_LONG" json:"SSR_CLASSNAME_LONG"`
						SSRCOMPONENT            string   `xml:"SSR_COMPONENT" json:"SSR_COMPONENT"`
						SSRCOMPONENTLOVDescr    string   `xml:"SSR_COMPONENT_LOVDescr" json:"SSR_COMPONENT_LOVDescr"`
						SSRDESCRLONG            string   `xml:"SSR_DESCRLONG" json:"SSR_DESCRLONG"`
						STATUS                  []string `xml:"STATUS" json:"STATUS"`
						STATUSLOVDescr          string   `xml:"STATUS_LOVDescr" json:"STATUS_LOVDescr"`
						STRM                    string   `xml:"STRM" json:"STRM"`
						STRMLOVDescr            string   `xml:"STRM_LOVDescr" json:"STRM_LOVDescr"`
						SUBJECT                 string   `xml:"SUBJECT" json:"SUBJECT"`
						SUBJECTLOVDescr         string   `xml:"SUBJECT_LOVDescr" json:"SUBJECT_LOVDescr"`
					} `xml:"CLASS_SUMMARY" json:"CLASS_SUMMARY"`
				} `xml:"CLASSES_SUMMARY" json:"CLASSES_SUMMARY"`
				COURSETITLELONG     string `xml:"COURSE_TITLE_LONG" json:"COURSE_TITLE_LONG"`
				CRSEID              string `xml:"CRSE_ID" json:"CRSE_ID"`
				CRSEIDLOVDescr      string `xml:"CRSE_ID_LOVDescr" json:"CRSE_ID_LOVDescr"`
				CRSEOFFERNBR        string `xml:"CRSE_OFFER_NBR" json:"CRSE_OFFER_NBR"`
				INSTITUTION         string `xml:"INSTITUTION" json:"INSTITUTION"`
				INSTITUTIONLOVDescr string `xml:"INSTITUTION_LOVDescr" json:"INSTITUTION_LOVDescr"`
				SUBJECT             string `xml:"SUBJECT" json:"SUBJECT"`
				SUBJECTLOVDescr     string `xml:"SUBJECT_LOVDescr" json:"SUBJECT_LOVDescr"`
			} `xml:"SUBJECT" json:"SUBJECT"`
		} `xml:"SUBJECTS" json:"SUBJECTS"`
	} `xml:"SEARCH_RESULT" json:"SEARCH_RESULT"`
	Xmlns string `xml:"_xmlns" json:"_xmlns"`
}

func GetRequestHTTPS() SSRGETCLASSESRESP {
	result := SSRGETCLASSESRESP{}
	timeout := time.Duration(1000 * time.Second)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{
		Timeout:   timeout,
		Transport: tr,
	}

	AlbertXML := new(SSR_GET_CLASSES_REQ)
	AlbertXML.CLASS_SEARCH_REQUEST.TypeRequest = "complete"
	AlbertXML.CLASS_SEARCH_REQUEST.LanguageCd = "eng"
	AlbertXML.CLASS_SEARCH_REQUEST.STRM = "1158"
	AlbertXML.CLASS_SEARCH_REQUEST.SUBJECT = "PH-UY"
	AlbertXML.CLASS_SEARCH_REQUEST.INSTITUTION = "NYUNV"
	AlbertXML.CLASS_SEARCH_REQUEST.CAMPUS = "WS"
	AlbertXML.CLASS_SEARCH_REQUEST.ACAD_CAREER = "UGRD"
	AlbertXML.CLASS_SEARCH_REQUEST.SSR_EXACT_MATCH1 = "E"
	AlbertXML.CLASS_SEARCH_REQUEST.CATALOG_NBR = "1013"

	output, err := xml.MarshalIndent(AlbertXML, "", " ")
	if err != nil {
		fmt.Println(err)
		return result
	}
	output = []byte(xml.Header + string(output))
	xmlSend := bytes.NewBuffer([]byte(output))

	currentURL := "https://sis-int.nyu.edu/PSIGW/RESTListeningConnector/SSR_GET_CLASSES_R.v1/classes/get?type=complete&languageCd=eng"
	req, err := http.NewRequest("POST", currentURL, xmlSend)
	req.Header.Set("Content-Type", "application/xml")

	resp, _ := client.Do(req)
	if err != nil {
		return result
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	bodyHeaderless := strings.Replace(string(body), xml.Header, "", 1)
	err = xml.Unmarshal([]byte(bodyHeaderless), &result)
	if err != nil {
		fmt.Printf("error: %v", err)
		return result
	}
	return result
}

func main() {
	classResult := GetRequestHTTPS()
	fmt.Println(classResult)
}
