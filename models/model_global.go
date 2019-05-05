package models


type Meta struct {
    Status bool `json:"status"`
    Code int `json:"code"`
    Message string `json:"message"`
}


type Links struct {
    Count int `json:"count"`
}

