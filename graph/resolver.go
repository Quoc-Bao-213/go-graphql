package graph

import "github.com/Quoc-Bao-213/go-graphql/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	CharacterStore map[string]model.Character
}
