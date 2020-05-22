package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{}


func (r *queryResolver) Links(ctx context.Context) ([]*Link,error){
	var links []*Link
	links = append(links,&Link{Title: "our dummy link", Address: "https://address.org", User: &User{Username: "admin"}})
	return links, nil
}
