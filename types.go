package main

//IndexPage represents the content of the index page, available on "/"
//The index page shows a list of all notes stored on db
type IndexPage struct {
	AllNotes []Note
}

//NotePage represents the content of the note page, available on "/note.html"
//The note page shows info about a given note
type NotePage struct {
	TargetNote Note
}

//Note represents a note object
type Note struct {
	ID              int
	Content         string
}

//ErrorPage represents shows an error message, available on "/note.html"
type ErrorPage struct {
	ErrorMsg string
}
