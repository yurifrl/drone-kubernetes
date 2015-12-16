package main

import (
	"github.com/drone/drone-go/drone"
)

type Params struct {
	Services               []string `json:services`
	ApiServer              string   `json:apiserver`
	Token                  string   `json:token`
	Namespace              string   `json:namespace`
	Debug                  string   `json:debug`
	Source                 string   `json:source`
	Tag                    string   `json:tag`
	ReplicationControllers []string `json:replicationcontrollers`
}

type Context struct {
	System drone.System
	Repo   drone.Repo
	Build  drone.Build
	Vargs  Params
}

type ReqEnvelope struct {
	Verb  string
	Token string
	Json  []byte
	Url   string
}

type Artifact struct {
	ApiVersion string
	Kind       string
	Data       []byte
	Metadata   struct {
		Name string
	}
	Url string
}

type DroneData struct {
	TAG string
}
