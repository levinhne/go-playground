//go:build ignore

package main

import (
	"log"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

type GreetExtension struct {
	entc.DefaultExtension
	word GreetingWord
}

func (*GreetExtension) Templates() []*gen.Template {
	return []*gen.Template{
		gen.MustParse(gen.NewTemplate("greet").ParseFiles("templates/greet.gohtml")),
	}
}

// GreetingWord implements entc.Annotation.
type GreetingWord string

// Name of the annotation. Used by the codegen templates.
func (GreetingWord) Name() string {
	return "GreetingWord"
}

func (s *GreetExtension) Annotations() []entc.Annotation {
	return []entc.Annotation{
		s.word,
	}
}

func main() {
	err := entc.Generate("./schema",
		&gen.Config{},
		entc.Extensions(&GreetExtension{
			word: GreetingWord("Shalom"),
		}),
	)
	if err != nil {
		log.Fatal("running ent codegen:", err)
	}
}
