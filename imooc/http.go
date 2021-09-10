package main

type Res struct {
	Code string `json:"code"`
	Data Data   `json:"data"`
}

type Data struct {
	Count int64  `json:"count"`
	Data  string `json:"data"`
	Pics  Pics   `json:"pics"`
}

type Pics struct {
	Pic1 Pic1 `json:"pic_1"`
}

type Pic1 struct {
	Width  int64  `json:"width"`
	Size   int64  `json:"size"`
	Ret    int64  `json:"ret"`
	Height int64  `json:"height"`
	Name   string `json:"name"`
	PID    string `json:"pid"`
}
