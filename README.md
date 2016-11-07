
Here at DataDeck, we have a list of songs in a database that we would like to expose via API. The songs are stored in SQLite database, which can be retrieved from https://s3.amazonaws.com/bv-challenge/jrdd.db. 

Your mission, should you choose to accept it, is to build an API to deliver this data in JSON format.</br>
 The API should make all songs searchable by artist, song, or genre. The genre should be searchable by the name in the genres table.</br>
 It should be implemented in Golang 1.6.x</br>
 The API framework you use is completely up to you. We use Goji at DataDeck.</br>
 The data that is returned should include song, artist, genre name, and length.</br>
 There should be a glide.yml file containing all the dependencies.</br>
 There must be a README that explains how to setup and install the application.</br>
 Code must be on Github.</br></br>
Extra Credit</br>
 Make a function in the API return songs by length, which we would like to search by passing a minimum and maximum length.
Make a function in the API return a list of the genres, and the number of songs and the total length of all the songs by genre.</br>
 Add unit testing

## Getting Started
These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. 

### Prerequisites
The version in which the API was implementes was in Golang 1.6.3, to check the version run `go version`</br>
The Glide package manager can be obtained here [Glide](https://github.com/Masterminds/glide), just clonde the git and build it with Go like this `go build`.

## The Project
To clone the project in your local machine run </br>
`git clone https://github.com/michaelyip/BeenVerified-Challenge.git $GOPATH/src/michaelyip/BeenVerified-Challenge`

To install the dependencies</br>
`cd $GOPATH/src/michaelyip/BeenVerified-Challenge`</br>
`glide install`

To build the project after installing the dependencies
`cd $GOPATH/src/michaelyip/BeenVerified-Challenge` </br>
`go build`</br>

### Running
To deploy the project execute the generated file `BeenVerified-Challenge` </br>
And now you can start sending API request to http://localhost:8000

## API Reference
### Songs 
Get all the songs
`http://localhost:8000/songs`

### Songs by Artists
Get all the songs by a artist name, in this case "424"
`http://localhost:8000/songs/artist/424`

### Songs by Song's Name
Get a song by name
`http://localhost:8000/songs/song/Gala`

### Songs by Genre
Get all the song by a specific genre
`http://localhost:8000/songs/genre/Pop`

### List of the genres, and the number of songs and the total length of all the songs by genre
`http://localhost:8000/songs/genre`

### Songs by length (min, max)
Example: Songs between 100 and 200
`http://localhost:8000/songs/length/100/200`

## Author
Michael Yip



