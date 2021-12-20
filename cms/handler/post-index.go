package handler

import (
	"strconv"
	"log"

	bgv "blog/gunk/v1/post"
	"net/http"
	"github.com/gorilla/mux"
)

type IndexPost struct{
	Post Post
}

func (h *Handler) Index (rw http.ResponseWriter, r *http.Request) {


	res,err:= h.tc.ListPost(r.Context(), &bgv.ListPostRequest{})
	if err!=nil{
		log.Fatal(err)
	}
	if err:= h.templates.ExecuteTemplate(rw,"index-post.html", res); err !=nil{
		http.Error(rw, err.Error(),http.StatusInternalServerError)
		return
	}
}

func (h *Handler) deletePost (rw http.ResponseWriter, r *http.Request) {
vars := mux.Vars(r)
	id := vars["id"]
	
	if id == "" {
		http.Error(rw, "invalid update", http.StatusTemporaryRedirect)
		return
	}
		Id, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	_,err = h.tc.DeletePost(r.Context(),&bgv.DeletePostRequest{
		ID: Id,
	})
	if err!=nil{
		log.Fatal(err)
	}
	http.Redirect(rw,r, "/", http.StatusTemporaryRedirect)
}
