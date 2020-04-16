package args

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/gengo/types"
)

type CustomArgs struct {
	Package      string
	TypesByGroup map[schema.GroupVersion][]*types.Name
	Options      Options
	OutputBase   string
}

type Options struct {
	OutputPackage string
	Groups        map[string]Group
	Boilerplate   string
}

type Group struct {
	Types         []interface{}
	GenerateTypes bool
	// Generate clientsets
	GenerateClients             bool
	OutputControllerPackageName string
	// Generate listers
	GenerateListers bool
	// Generate informers
	GenerateInformers bool
	// The package name of the API types
	PackageName string
	// Use existing clientset, informer, listers
	ClientSetPackage string
	ListersPackage   string
	InformersPackage string
}
