package main

import (

"github.com/barajas123/go-lib-rabbit/rabbit"
)

func main(){
	c := rabbit.NewRabbitWithTopic("test-fanout","fan.*")

	c.Consume()

}
