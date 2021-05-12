package compute

var movies map[string]models.Movie


// Retrieve movie from movies list
func getMovie(title: string) Movie{
    movie, ok := movies[title]

    if !ok {
        m := Movie{Title: title, Related: make(map[string]float64{})}
    }

    // retrieve related movies
    relatedList = m.Related

    // struct to place map values so they can be later ordered
    type keyVal struct {
        Key   string
        Value float64
    }

    // Sort map in an array
    var sortedMovies []keyVal
    for key, value := range relatedList {
        sortedMovies = append(movieArray, keyVal{key, value})
    }

    sort.Slice(sortedMovies, func(i, j int) bool {
        return sortedMovies[i].Value > sortedMovies[j].Value
    })

    fmt.Printf("%s, %d\n", sortedMovies[0].Key, sortedMovies[0].Value)
    return kv.Key

}
// Access Related List

// get suggestion