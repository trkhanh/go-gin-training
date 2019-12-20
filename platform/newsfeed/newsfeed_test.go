package newsfeed

import "testing"
import "fmt"
func  TestAdd(t *testing.T)  {
	feed := New()
	feed.Add(Item{})
	fmt.Println(feed)
	if len(feed.Items) ==0 {
		t.Error("Item was not added")
	}
}

func TestGetAll(t *testing.T)  {
	feed := New()
	feed.Add(Item{})
	results := feed.GetAll()
	if len(results) !=  1{
		t.Error("Item was not added")
	}
}