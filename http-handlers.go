package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

func handleSaveNote(w http.ResponseWriter, r *http.Request) {
	var id = 0
	var err error

	r.ParseForm()
	params := r.PostForm
	idStr := params.Get("id")

	if len(idStr) > 0 {
		id, err = strconv.Atoi(idStr)
		if err != nil {
			renderErrorPage(w, err)
			return
		}
	}

	content := params.Get("content")

	if id == 0 {
		_, err = insertNote(content)
	} else {
		_, err = updateNote(id, content)
	}

	if err != nil {
		renderErrorPage(w, err)
		return
	}

	http.Redirect(w, r, "/", 302)
}

func handleListNotes(w http.ResponseWriter, r *http.Request) {
	notes, err := allNotes()
	if err != nil {
		renderErrorPage(w, err)
		return
	}

	buf, err := ioutil.ReadFile("www/index.html")
	if err != nil {
		renderErrorPage(w, err)
		return
	}

	var page = IndexPage{AllNotes: notes}
	indexPage := string(buf)
	t := template.Must(template.New("indexPage").Parse(indexPage))
	t.Execute(w, page)
}

func handleViewNote(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	idStr := params.Get("id")

	var currentNote = Note{}

	if len(idStr) > 0 {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			renderErrorPage(w, err)
			return
		}

		currentNote, err = getNote(id)
		if err != nil {
			renderErrorPage(w, err)
			return
		}
	}

	buf, err := ioutil.ReadFile("www/note.html")
	if err != nil {
		renderErrorPage(w, err)
		return
	}

	var page = NotePage{TargetNote: currentNote}
	notePage := string(buf)
	t := template.Must(template.New("notePage").Parse(notePage))
	err = t.Execute(w, page)
	if err != nil {
		renderErrorPage(w, err)
		return
	}
}

func handleDeleteNote(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	idStr := params.Get("id")

	if len(idStr) > 0 {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			renderErrorPage(w, err)
			return
		}

		n, err := removeNote(id)
		if err != nil {
			renderErrorPage(w, err)
			return
		}

		fmt.Printf("Rows removed: %v\n", n)
	}
	http.Redirect(w, r, "/", 302)
}

func renderErrorPage(w http.ResponseWriter, errorMsg error) {
	buf, err := ioutil.ReadFile("www/error.html")
	if err != nil {
		log.Printf("%v\n", err)
		fmt.Fprintf(w, "%v\n", err)
		return
	}

	var page = ErrorPage{ErrorMsg: errorMsg.Error()}
	errorPage := string(buf)
	t := template.Must(template.New("errorPage").Parse(errorPage))
	t.Execute(w, page)
}
