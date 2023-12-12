package models

type Artists []ArtistElement

type ArtistElement struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

type ArtistInfo struct {
	ID    int    
	Image string  
	Name    string  
	Members []string   
	CreationDate int    
	FirstAlbum   string     
	Locations    []string
	ConcertDates []string
	Relations    map[string][]string
}

type Loc struct {
	ID int `json:"id"`
	Location []string `json:"locations"`
}

type Concert struct {
	ID int `json:"id"`
	Date []string `json:"dates"`
}

type Relation struct {
    ID int `json:"id"`
    Relation map[string][]string `json:"datesLocations"`
}
