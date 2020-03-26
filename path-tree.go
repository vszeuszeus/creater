package main

import (
	"io/ioutil"
	"log"
	"os"
	"path"
)

type Component interface {
	GetName() string
	GetChildren() []Component
	Create(prefix string, modelName ModelName)
}

type Directory struct {
	Name   string
	Children []Component
}

func (d *Directory) GetChildren() []Component {
	return d.Children
}

func (d *Directory) GetName() string {
	return d.Name
}

func (d *Directory) Create(prefix string, modelName ModelName) {
	dirname := path.Join(prefix, d.GetName())
	if err := os.MkdirAll(dirname, 644); err != nil {
		log.Println("Can`t create directory ", dirname, " " + err.Error())
		return
	}

	log.Println("Dir '", dirname, "' created successful")

	children := d.GetChildren()
	for _, child := range children {
		child.Create(dirname, modelName)
	}
}

type File struct {
	Name string
	DataGetter FileDataGetter
}

func (f *File) GetName() string {
	return f.Name
}

func (f *File) GetChildren() []Component {
	return []Component{}
}

func (f *File) Create(prefix string, modelName ModelName) {
	data := f.GetData(modelName)
	filename := path.Join(prefix, f.GetName())
	if err := ioutil.WriteFile(filename, data, 644); err != nil {
		log.Println("Error in create file ", filename, " " + err.Error())
		return
	}
	log.Println("File '" + filename + "' created Successful")
}

func (f *File) GetData(modelName ModelName) []byte {
	if f.DataGetter != nil {
		return f.DataGetter.GetData(modelName.ToProgramName())
	}
	return []byte{}
}

type FileDataGetter interface {
	GetData(string) []byte
}
