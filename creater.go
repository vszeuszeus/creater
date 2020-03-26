package main

import (
	"fmt"
	"github.com/iancoleman/strcase"
	"os"
	"path"
)

var RepositoryDataCenters = []string{"postgres"}
var basePath = path.Join("backend", "pkg")

type ModelName string

func (m *ModelName) ToProgramName() string {
	return strcase.ToCamel(string(*m))
}

func (m *ModelName) ToPathName() string {
	return strcase.ToSnake(string(*m))
}

type Resource struct {
	Method string
	Name   string
}
var Resources = []Resource{
	{Method: "GET", Name: "Fetch"},
	{Method: "GET", Name: "GetById"},
	{Method: "POST", Name: "Store"},
	{Method: "PUT", Name: "Update"},
	{Method: "DELETE", Name: "Delete"},
}


func main() {
	var modelName = ModelName(os.Args[1])
	pathModelName := modelName.ToPathName()
	paths := &Directory{
		Name: pathModelName,
		Children: []Component{
			&Directory{
				Name: "delivery",
				Children: []Component{
					&Directory{
						Name: "http",
						Children: []Component{
							&File{
								Name: "handler.go",
							},
							&File{
								Name: "handler_test.go",
							},
							&File{
								Name: "register.go",
							},
						},
					},
				},
			},
			&Directory{
				Name: "repository",
				Children: func() (res []Component) {
					res = []Component{}
					for i := range RepositoryDataCenters {
						res = append(res, &Directory{
							Name: RepositoryDataCenters[i],
							Children: []Component{
								&File{
									Name: RepositoryDataCenters[i] + "_" + pathModelName + ".go",
								},
								&File{
									Name: RepositoryDataCenters[i] + "_" + pathModelName + "_test.go",
								},
							},
						})
					}
					return
				}(),
			},
			&Directory{
				Name: "mock",
				Children: []Component{
					&File{
						Name: "repository.go",
					},
					&File{
						Name: "usecase.go",
					},
				},
			},
			&Directory{
				Name: "usecase",
				Children: []Component{
					&File{
						Name: pathModelName + "_usecase.go",
					},
					&File{
						Name: pathModelName + "_usecase_test.go",
					},
				},
			},
			&File{
				Name: "usecase.go",
				//DataGetter: &UseCaseInterface{},
			},
			&File{
				Name: "repository.go",
			},
		},
	}

	paths.Create(basePath, modelName)
	fmt.Print(os.Args)
}

type UseCaseInterface struct {
}

func (useCaseInterface *UseCaseInterface) GetData(modelName string) []byte {
	return []byte("" +
		"package " + modelName + "\n" +
		"import (\n" +
		"\"context\"" +
		"\"git.otgroup.kz/capital-city-center/backend/models\"\n" +
		")\n\n" +
		"type UseCase interface {\n" +
		func() (str string){
			for i := range Resources {
				str += Resources[i].Name + "(ctx context.Context) ([]models." + modelName + ", int, error)"
			}
			return
		}() +
		"" +
		"" +
		"" +
		"")
}