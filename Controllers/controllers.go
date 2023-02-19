package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Abhishek1833/ItemList/models"
	//"zocketAssignment/models"
)

func Welcome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome"))
	//return
}
func PostItems(w http.ResponseWriter, r *http.Request) {
	var itm models.Item
	json.NewDecoder(r.Body).Decode(&itm)
	id, err := models.PostItem(&itm)
	if err != nil {
		//fmt.Println("cant add post to database")
		panic("cant add post to database")
	}
	itm.Id = int(id)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(itm)
}

func GetItems(w http.ResponseWriter, r *http.Request) {
	//var items []models.Item
	id := r.URL.Query().Get("id")
	fmt.Println("id:", id)
	id1, _ := strconv.ParseInt(id, 10, 64)
	if id1 == 0 {
		items, err := models.GetAllItems()
		if err != nil {
			panic("can not fetch data from the datbase")
		}
		json.NewEncoder(w).Encode(items)
	} else {
		fmt.Println("id:", id1)
		fmt.Println("else")
		//items, err := models.GetItems(int(id1))
		items, err := models.GetItems(int(id1))
		if err != nil {
			json.NewEncoder(w).Encode(err.Error())
		}
		json.NewEncoder(w).Encode(items)
	}
}

func UpdateItem(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	id1, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		panic(err.Error())
	}
	if id1 == 0 {
		s := "enter a valid id"
		json.NewEncoder(w).Encode(s)
	}
	item := models.Item{}
	json.NewDecoder(r.Body).Decode(&item)
	olditem, err := models.GetItems(int(id1))
	if err != nil {
		panic(err.Error())
	}
	if item.Name == "" {
		item.Name = olditem.Name
	}
	if item.Type == "" {
		item.Type = olditem.Type
	}
	if item.Description == "" {
		item.Description = olditem.Description
	}
	item.Id = int(id1)
	//item := models.Item{Id: int(id1)}
	err1 := models.UpdateItem(&item)
	if err1 != nil {
		json.NewEncoder(w).Encode(err.Error())
	}
	json.NewEncoder(w).Encode(item)
}

func DeleteItem(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	id1, err := strconv.ParseInt(id, 0, 64)
	if err != nil {
		panic(err.Error())
	}
	err = models.DeleteItem(int(id1))
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
	}
	s := "item deleted succesfully"
	json.NewEncoder(w).Encode(s)
}
