package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"test_ozon/domain/repository"
	"test_ozon/dto"
)

func respondWithError(w http.ResponseWriter, code int, message string) {
	http.Error(w, message, code)
}

func decodeRequestBody(w http.ResponseWriter, r *http.Request, target interface{}) error {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(target)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return err
	}
	return nil
}

func GetOriginalUrl(w http.ResponseWriter, r *http.Request) {
	var getUrl dto.GetUrl
	if err := decodeRequestBody(w, r, &getUrl); err != nil {
		return
	}

	result, err := repository.Repo.GetOriginalUrl(getUrl.ShortUrl)
	if err != nil {
		log.Println("Error fetching original URL:", err)
		respondWithError(w, http.StatusNotFound, "Error fetching URL")
		return
	}

	responseJSON, err := json.Marshal(dto.OriginalUrl{Url: *result})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error creating response")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(responseJSON)
}

func SaveUrl(w http.ResponseWriter, r *http.Request) {
	var saveUrl dto.SaveUrl
	if err := decodeRequestBody(w, r, &saveUrl); err != nil {
		return
	}

	exist := repository.Repo.CheckExistOriginalUrl(saveUrl.OriginalUrl)
	if exist {
		respondWithError(w, http.StatusConflict, "Url already exists")
		return
	}
	shortLink := GenerateShortLink()
	repository.Repo.SaveUrls(saveUrl.OriginalUrl, shortLink)

	response, err := json.Marshal(dto.ShortUrl{Url: shortLink})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error creating response")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
