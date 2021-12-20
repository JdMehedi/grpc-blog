package handler

import (
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"github.com/gorilla/sessions"

	bgv "blog/gunk/v1/post"
	
)

type Handler struct{
   templates  *template.Template 
   decoder    *schema.Decoder
   sess       *sessions.CookieStore
   tc         bgv.PostServiceClient

}
// var sessionName = "storeCookie"

func New(decoder  *schema.Decoder ,sess  *sessions.CookieStore,  tc  bgv.PostServiceClient) *mux.Router{
	h:= &Handler{
		decoder: decoder,
		sess: sess,
		tc: tc,
	}

	h.parseTemplate()
	r :=mux.NewRouter() 
	// l:=r.NewRoute().Subrouter()
	
	r.HandleFunc("/", h.Index)

    s:=r.NewRoute().Subrouter()

	s.HandleFunc("/posts/create", h.createPost)
	s.HandleFunc("/posts/store", h.storePost)
	s.HandleFunc("/posts/{id:[0-9]+}/edit", h.editPost)
	s.HandleFunc("/posts/{id:[0-9]+}/update", h.updatePost)
	s.HandleFunc("/posts/{id:[0-9]+}/delete", h.deletePost)

	s.PathPrefix("/asset/").Handler(http.StripPrefix("/asset/", http.FileServer(http.Dir("./"))))
	s.Use(h.middleWare)
	

	r.NotFoundHandler = http.HandlerFunc(func (rw http.ResponseWriter, r *http.Request)  {
		if err :=h.templates.ExecuteTemplate(rw,"404.html",nil); err != nil{
			http.Error(rw, err.Error(),http.StatusInternalServerError)
			return
		}
		
	} )

	return r
	
}

func (h *Handler) middleWare(next http.Handler) http.Handler{

	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(rw,r)
		// return
		// session, _:= h.sess.Get(r, sessionName)
	
	// 	auth,ok:=session.Values["Authenticated"].(bool)
	// 	if !ok || !auth{
	// 		http.Redirect(rw,r,"/login",http.StatusTemporaryRedirect)
	// 		return
	// 	}

	// 	next.ServeHTTP(rw,r)
	})
}



func (h *Handler) parseTemplate(){
	h.templates= template.Must(template.ParseFiles(
		"cms/assets/templates/posts/create-post.html",
		"cms/assets/templates/posts/index-post.html",
		"cms/assets/templates/posts/edit-post.html",
		"cms/assets/templates/404.html",
		
	))
}

