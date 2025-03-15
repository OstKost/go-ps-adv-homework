package link

import (
	"fmt"
	"go-ps-adv-homework/configs"
	"go-ps-adv-homework/pkg/event"
	"go-ps-adv-homework/pkg/middleware"
	"go-ps-adv-homework/pkg/request"
	"go-ps-adv-homework/pkg/response"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
)

type handler struct {
	LinkRepository *LinkRepository
	*configs.Config
	EventBus *event.EventBus
}

type HandlerDependencies struct {
	LinkRepository *LinkRepository
	*configs.Config
	EventBus *event.EventBus
}

func NewLinkHandler(router *http.ServeMux, dependencies HandlerDependencies) {
	handler := &handler{
		LinkRepository: dependencies.LinkRepository,
		EventBus:       dependencies.EventBus,
	}
	router.HandleFunc("POST /link", handler.CreateLink())
	router.HandleFunc("GET /{hash}", handler.GoToLink())
	router.HandleFunc("GET /link", handler.GetList())
	router.Handle("PATCH /link/{linkId}", middleware.IsAuthed(handler.UpdateLink(), dependencies.Config))
	router.HandleFunc("DELETE /link/{linkId}", handler.DeleteLink())
}

func (handler *handler) CreateLink() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Body
		body, err := request.HandleBody[CreateLinkRequest](&w, r)
		if err != nil {
			response.Json(w, err.Error(), http.StatusBadRequest)
			return
		}
		// New link
		link := NewLink(body.Url)
		for {
			existedLink, _ := handler.LinkRepository.GetByHash(link.Hash)
			if existedLink == nil {
				break
			}
			link.GenerateHash()
		}
		// Create
		createdLink, err := handler.LinkRepository.Create(link)
		if err != nil {
			response.Json(w, err.Error(), http.StatusBadRequest)
			return
		}
		response.Json(w, createdLink, http.StatusCreated)
	}
}

func (handler *handler) GoToLink() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hash := r.PathValue("hash")
		link, err := handler.LinkRepository.GetByHash(hash)
		if err != nil {
			log.Println(err.Error())
			response.Json(w, err.Error(), http.StatusNotFound)
			return
		}
		go handler.EventBus.Publish(event.Event{
			Type: event.EventLinkVisited,
			Data: link.ID,
		})
		//handler.StatRepository.AddClick(link.ID)
		http.Redirect(w, r, link.Url, http.StatusTemporaryRedirect)
	}
}

func (handler *handler) UpdateLink() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Context
		ctx := r.Context()
		email, ok := ctx.Value(middleware.ContextEmailKey).(string)
		if ok {
			fmt.Println("ContextEmailKey: ", email)
		}
		// Params
		idString := r.PathValue("linkId")
		if idString == "" {
			response.Json(w, "link ID is required", http.StatusBadRequest)
			return
		}
		id, err := strconv.ParseUint(idString, 10, 32)
		if err != nil {
			response.Json(w, err.Error(), http.StatusBadRequest)
			return
		}
		// Body
		body, err := request.HandleBody[UpdateLinkRequest](&w, r)
		if err != nil {
			response.Json(w, err.Error(), http.StatusBadRequest)
			return
		}
		// Update
		link, err := handler.LinkRepository.Update(&Link{
			Model: gorm.Model{ID: uint(id)},
			Url:   body.Url,
			Hash:  body.Hash,
		})
		response.Json(w, link, http.StatusOK)
	}
}

func (handler *handler) DeleteLink() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Params
		idString := r.PathValue("linkId")
		if idString == "" {
			response.Json(w, "link ID is required", http.StatusBadRequest)
			return
		}
		id, err := strconv.ParseUint(idString, 10, 32)
		if err != nil {
			response.Json(w, err.Error(), http.StatusBadRequest)
			return
		}
		// Check if exists
		_, err = handler.LinkRepository.GetById(uint(id))
		if err != nil {
			response.Json(w, err.Error(), http.StatusNotFound)
			return
		}
		// Delete
		err = handler.LinkRepository.Delete(uint(id))
		if err != nil {
			response.Json(w, err.Error(), http.StatusInternalServerError)
			return
		}
		response.Json(w, id, http.StatusOK)
	}
}

func (handler *handler) GetList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		url := query.Get("url")
		limit, offset := request.GetPaginationParams(query)
		// Get links
		links, err := handler.LinkRepository.GetActiveList(url, limit, offset)
		if err != nil {
			log.Println(err.Error())
			response.Json(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Get links count
		count, err := handler.LinkRepository.Count(url, limit, offset)
		if err != nil {
			log.Println(err.Error())
			response.Json(w, err.Error(), http.StatusInternalServerError)
			return
		}
		response.Json(w, GetAllLinksResponse{
			Links: links,
			Count: count,
		}, http.StatusOK)
	}
}
