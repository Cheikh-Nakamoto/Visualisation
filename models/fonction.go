package models

import (

	"encoding/json"
	"io"
	"net/http"
	"regexp"
	"strconv"
)

var All_Artist Artists
var ArtistPrint ArtistInfo

const URL = "https://groupietrackers.herokuapp.com/api/"

// The FetchData function fetches data from a given URL, parses it as JSON, and stores it in the
// All_Artist variable.
func FetchData(url string, data interface{}) error {

	body, err := MiniRecup(url)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return err
	}
	return nil
}

// The MiniRecup function retrieves the content of a web page specified by the given URL and returns it
// as a byte slice.
func MiniRecup(url string) ([]byte, error) {

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func RemoveChar(s []string) []string {
	re := regexp.MustCompile(`(\*)(\w)`)

	for i := range s {
		s[i] = re.ReplaceAllString(s[i], "$2")
	}
	return s
}

// The ParseData function parses data from the All_Artist variable and stores it in the ArtistPrint variable.
func SearchId(id int) error {

	Print := All_Artist[id]
	FetchData(URL+"artist/"+strconv.Itoa(id), &Print)

	readRelation, err := MiniRecup(Print.Relations)
	if err != nil {
		return err
	}
	readLocation, _ := MiniRecup(Print.Locations)
	readDate, _ := MiniRecup(Print.ConcertDates)

	var location Loc
	var date Concert
	var relation Relation

	err = json.Unmarshal(readRelation, &relation)
	if err != nil {
		return err
	}
	json.Unmarshal(readLocation, &location)
	if err != nil {
		return err
	}
	json.Unmarshal(readDate, &date)
	if err != nil {
		return err
	}

	ArtistPrint = ArtistInfo{
		ID:           Print.ID,
		Image:        Print.Image,
		Name:         Print.Name,
		Members:      Print.Members,
		CreationDate: Print.CreationDate,
		FirstAlbum:   Print.FirstAlbum,
		Locations:    location.Location,
		ConcertDates: RemoveChar(date.Date),
		Relations:    relation.Relation,
	}

	return nil
}
