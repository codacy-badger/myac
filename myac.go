package main

import(
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"fmt"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/storage/memory"
)

type ServerConf struct {
	Server struct {
		Git struct {
			URI      string `json:"uri"`
			Username string `json:"username"`
			Password string `json:"password"`
		} `json:"git"`
	} `json:"server"`
}

func (c *ServerConf) getConf() *ServerConf {
	yamlFile, err := ioutil.ReadFile("config.yml")

	if err != nil {
		log.Printf("yamlFile.Get err #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}

func main() {
	var c ServerConf
	var url string = c.getConf().Server.Git.URI
	fmt.Println(url)
	r, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions {
		URL: url,
	})
	if err != nil {
		fmt.Println(err)
	}
ref, err := r.Head()
fmt.Println(ref)
}

//todo read configfile from server using git settings
//todo display configfile via http