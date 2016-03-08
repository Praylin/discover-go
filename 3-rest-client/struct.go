package main

import (
	"fmt"
)

type Message struct {
	Title string
	Year  string
	Rated string
  Released string
  Runtime string
  Genre string
  Director string
  Writer string
  Actors string
  Plot string
  Language string
  Awards string
  Country string
  Poster string
  Metascore int
  imdbRating string
  imdbID string
  Type string
  Response string
}

func Marshal(v interface{}) ([]byte, error) {
  m := Message{"Batman Begins", "2005", "PG-13", "15 Jun 2005", "140 min", "Action, Adventure", "Christopher Nolan", "Bob Kane (characters), David S. Goyer (story), Christopher Nolan (screenplay), David S. Goyer (screenplay)", "Christian Bale, Michael Caine, Liam Neeson, Katie Holmes", "After training with his mentor, Batman begins his war on crime to free the crime-ridden Gotham City from corruption that the Scarecrow and the League of Shadows have cast upon it.", "English, Urdu, Mandarin", "Nominated for 1 Oscar. Another 15 wins & 64 nominations.", "USA, UK", "http://ia.media-imdb.com/images/M/MV5BNTM3OTc0MzM2OV5BMl5BanBnXkFtZTYwNzUwMTI3._V1_SX300.jpg", 70, "8.3", "918,818", "tt0372784", "movie", "True"}
  b, err := json.Marshal(m)
  fmt.Println(b)
  return err
}
