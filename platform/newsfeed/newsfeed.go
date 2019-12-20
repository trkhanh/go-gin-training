package newsfeed

type Getter interface{
	GetAll() []Item
}


type Adder interface{
	Add() []Item
}

type Item struct{
	Title string `json:"title`
	Post string `json:"post"`
}

type Repo struct {
	Items []Item
}

func New() *Repo{
	return &Repo{
		Items: []Item{},
	}
}

func (r *Repo) Add(item Item){
	r.Items = append(r.Items, item)
}

func (r *Repo) GetAll() []Item {
	return r.Items
}


func (r *Repo) Shit() []Item {
	return r.Items
}
