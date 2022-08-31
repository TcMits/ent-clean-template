//go:build ignore
// +build ignore

package main

import (
	"log"
	"text/template"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

// templateOption ensures the template instantiate
// once for config and execute the given Option.
func templateOption(next func(t *gen.Template) (*gen.Template, error)) entc.Option {
	return func(cfg *gen.Config) (err error) {
		tmpl, err := next(gen.NewTemplate("external"))
		if err != nil {
			return err
		}
		cfg.Templates = append(cfg.Templates, tmpl)
		return nil
	}
}

func TemplateDirWithFuncs(path string, funcMap template.FuncMap) entc.Option {
	return templateOption(func(t *gen.Template) (*gen.Template, error) {
		return t.Funcs(funcMap).ParseDir(path)
	})
}

func main() {
	opts := []entc.Option{
		TemplateDirWithFuncs("./template", entgql.TemplateFuncs),
	}
	if err := entc.Generate("./schema", &gen.Config{
		Features: []gen.Feature{
			gen.FeaturePrivacy,
			gen.FeatureEntQL,
			gen.FeatureLock,
		},
	}, opts...); err != nil {
		log.Fatal("running ent codegen:", err)
	}
}
