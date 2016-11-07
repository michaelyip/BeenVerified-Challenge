package main

import (
	"net/http"
	"goji.io/pat"
	"golang.org/x/net/context"
)

func get_all_songs(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	z := fetch_all_songs()
	w.Write(z)
}

func get_by_artist(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	var artist = pat.Param(ctx, "artist")
	response := fetch_by_artist(artist)
	w.Write(response)
}

func get_by_song(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	var song = pat.Param(ctx, "song")
	response := fetch_by_song(song)
	w.Write(response)
}

func get_by_genre(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	var genre = pat.Param(ctx, "genre")
	response := fetch_by_genre(genre)
	w.Write(response)
}

func get_genres_summary(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	response := fetch_genres_summary()
	w.Write(response)
}

func get_by_length(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	min := pat.Param(ctx, "min")
	max := pat.Param(ctx, "max")
	response := fetch_by_length(min, max)
	w.Write(response)
}
