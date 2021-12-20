package handler

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/gorilla/mux"

	bgv "blog/gunk/v1/post"
)

type formData struct {
	Post   Post
	Errors map[string]string
}

type Post struct {
	ID          int64
	Title       string
	Description string
	IsCompleted bool
}

func (c *Post) validate() error {

	return validation.ValidateStruct(c,
		validation.Field(&c.Title,
			validation.Required.Error("This filed cannot be null"),
			validation.Length(3, 30).Error("The Post name length must be between 3 and 30"),
		),
	)
}

func (h *Handler) createPost(rw http.ResponseWriter, r *http.Request) {
	Post := Post{}
	Errors := map[string]string{}

	h.loadCreatedPostForm(rw, Post, Errors)
}
func (h *Handler) storePost(rw http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	var post Post

	if err := h.decoder.Decode(&post, r.PostForm); err != nil {

		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := post.validate(); err != nil {
		vErrors, ok := err.(validation.Errors)
		if ok {
			vErrs := make(map[string]string)
			for key, value := range vErrors {
				vErrs[strings.Title(key)] = value.Error()

			}
			h.loadCreatedPostForm(rw, post, vErrs)
			return
		}

		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return

	}
	_, err := h.tc.CreatePost(r.Context(), &bgv.CreatePostRequest{
		Post: &bgv.Post{
			Title:       post.Title,
			Description: post.Description,
		},
	})

	if err != nil {
		log.Fatal(err)
	}

	http.Redirect(rw, r, "/", http.StatusTemporaryRedirect)

}

func (h *Handler) loadCreatedPostForm(rw http.ResponseWriter, posts Post, errs map[string]string) {

	form := formData{
		Post:   posts,
		Errors: errs,
	}
	if err := h.templates.ExecuteTemplate(rw, "create-post.html", form); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

}

func (h *Handler) editPost(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if id == "" {
		http.Error(rw, "invalid ", http.StatusTemporaryRedirect)
		return
	}

	Id, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	Errors := map[string]string{}

	res, err := h.tc.GetPost(r.Context(), &bgv.GetPostRequest{
		ID: Id,
	})
	if err != nil {
		log.Fatal(err)
	}
	Post := Post{
		ID:          res.Post.ID,
		Title:       res.Post.Title,
		Description: res.Post.Description,
		IsCompleted: res.Post.IsCompleted,
	}

	h.loadUpdatedPostForm(rw, Post, Errors)
}

func (h *Handler) updatePost(rw http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]

	if id == "" {
		http.Error(rw, "invalid update", http.StatusTemporaryRedirect)
		return
	}

	Id, err := strconv.ParseInt(id, 10, 64)
	fmt.Println(Id)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
	var post Post

	if err := h.decoder.Decode(&post, r.PostForm); err != nil {

		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println(post)

	if err := post.validate(); err != nil {
		vErrors, ok := err.(validation.Errors)
		if ok {
			vErrs := make(map[string]string)
			for key, value := range vErrors {
				vErrs[strings.Title(key)] = value.Error()

			}

			h.loadUpdatedPostForm(rw, post, vErrs)
			return
		}
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return

	}

	_, err = h.tc.UpdatePost(r.Context(), &bgv.UpdatePostRequest{
		Post: &bgv.Post{
			ID:          Id,
			Title:       post.Title,
			Description: post.Description,
			IsCompleted: post.IsCompleted,
		},
	})
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(rw, r, "/", http.StatusTemporaryRedirect)
}

func (h *Handler) loadUpdatedPostForm(rw http.ResponseWriter, posts Post, errs map[string]string) {

	form := formData{
		Post:   posts,
		Errors: errs,
	}
	if err := h.templates.ExecuteTemplate(rw, "edit-post.html", form); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

}
