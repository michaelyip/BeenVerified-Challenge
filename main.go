package main

import (
	"net/http"
	"goji.io"
	"goji.io/pat"
)



func main() {
	mux := goji.NewMux()

	mux.HandleFuncC(pat.Get("/songs"), get_all_songs)
	mux.HandleFuncC(pat.Get("/songs/artist/:artist"), get_by_artist)
	mux.HandleFuncC(pat.Get("/songs/song/:song"), get_by_song)
	mux.HandleFuncC(pat.Get("/songs/genre/:genre"), get_by_genre)

	//Extra Credit
	mux.HandleFuncC(pat.Get("/songs/genre"), get_genres_summary)
	mux.HandleFuncC(pat.Get("/songs/length/:min/:max"), get_by_length)


	http.ListenAndServe("localhost:8000", mux)

}