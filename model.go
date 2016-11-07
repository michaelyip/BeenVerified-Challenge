package main


type Songs struct {
	Id		int
	Artist	string
	Song 	string
	Genre	string
	Length	int
}

type SongsCollection struct {
	Songs []Songs
}

type Genres struct {
	Genre			string
	Num_Songs		int
	Total_length	int
}

type GenresCollection struct {
	Genres []Genres
}

