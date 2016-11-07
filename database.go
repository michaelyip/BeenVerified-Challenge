package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"encoding/json"
)

func fetch_all_songs() []byte {
	query := "select s.id, s.artist, s.song, g.name as genre, s.length from songs s inner join genres g on g.id = s.genre;"
	return fetch_from_songs(query)
}

func fetch_by_artist(artist string) []byte {
	query := "select s.id, s.artist, s.song, g.name as genre, s.length from songs s inner join genres g on g.id = s.genre where s.artist = \""+artist+"\";"
	return fetch_from_songs(query)
}

func fetch_by_song(song string) []byte {
	query := "select s.id, s.artist, s.song, g.name as genre, s.length from songs s inner join genres g on g.id = s.genre where s.song = \""+song+"\";"
	return fetch_from_songs(query)
}

func fetch_by_genre(genre string) []byte {
	query := "select s.id, s.artist, s.song, g.name as genre, s.length from songs s inner join genres g on g.id = s.genre where g.name = \""+genre+"\";"
	return fetch_from_songs(query)
}

func fetch_genres_summary() []byte {
	query := "select distinct g.name as genre, count(*) as num_songs, sum(s.length) as total_length from songs s inner join genres g on g.id = s.genre group by genre;"
	return fetch_from_genres(query)
}

func fetch_by_length(min, max string) []byte {
	query := "select s.id, s.artist, s.song, g.name as genre, s.length from songs s inner join genres g on g.id = s.genre where s.length between "+min+" and "+max+";"
	return fetch_from_songs(query)
}

func fetch_from_songs(queryStmt string) []byte {
	db, err := sql.Open("sqlite3", "./jrdd.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	stmt, err := db.Prepare(queryStmt)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	
	rows, err := stmt.Query()

	collection := []Songs {}
	for rows.Next() {
		var id int
		var artist string
		var song string
		var genre string
		var length int
		err = rows.Scan(&id, &artist, &song, &genre, &length)
		if err != nil {
			log.Fatal(err)
		}

		temp := Songs {
			Id: id,
			Artist: artist,
			Song: song,
			Genre: genre,
			Length: length,
		}
		collection = append(collection, temp)
	}


	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	y := SongsCollection {
		Songs: collection,
	}

	z, _ := json.Marshal(y)

	return z
}

func fetch_from_genres(queryStmt string) []byte {
	db, err := sql.Open("sqlite3", "./jrdd.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	stmt, err := db.Prepare(queryStmt)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	
	rows, err := stmt.Query()

	collection := []Genres {}
	for rows.Next() {
		var genre 			string
		var num_songs 		int
		var total_length 	int
		err = rows.Scan(&genre, &num_songs, &total_length)
		if err != nil {
			log.Fatal(err)
		}

		temp := Genres {
			Genre: genre,
			Num_Songs: num_songs,
			Total_length: total_length,
		}
		collection = append(collection, temp)
	}


	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	y := GenresCollection {
		Genres: collection,
	}

	z, _ := json.Marshal(y)

	return z
}

