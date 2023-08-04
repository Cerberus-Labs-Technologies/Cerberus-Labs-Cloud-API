package templates

import (
	"errors"
	"github.com/Cerberus-Labs-Technologies/Cerberus-Labs-Cloud-API/network/io/zipping"
	"log"
	"os"
	"strings"
)

type Template struct {
	Name           string `json:"name"`
	BelongsToGroup string `json:"belongsToGroup"`
	Description    string `json:"description"`
	Path           string `json:"path"`
	IsGlobal       bool   `json:"isGlobal"`
	IsProxy        bool   `json:"isProxy"`
}

const configBasePath = "configs/"
const configTemplatePath = configBasePath + "templates.json"

const baseTemplatePath = "templates/"

const generalTemplatePath = baseTemplatePath + "general/"
const generalProxyTemplatePath = baseTemplatePath + "general_proxy/"
const generalServerTemplatePath = baseTemplatePath + "general_server/"
const serverTemplatePath = baseTemplatePath + "server/"
const proxyTemplatePath = baseTemplatePath + "proxy/"

const ErrTemplateExists = "template already exists"

type TemplateInterface interface {
	GetPath() string
	GetBelongsToGroup() string
	GetIsGlobal() bool
	IsZip() bool
	Unzip() error
	CreateDirIfNotExists() error
	DoesTemplateDirectoryExist() bool
}

type Templates struct {
	Templates []Template `json:"templates"`
}

func (t *Template) GetPath() string {
	return t.Path
}

func (t *Template) GetBelongsToGroup() string {
	return t.BelongsToGroup
}

func (t *Template) GetIsGlobal() bool {
	return t.IsGlobal
}

func (t *Template) IsZip() bool {
	return strings.HasSuffix(t.Path, ".zip")
}

func (t *Template) CreateDirIfNotExists() error {
	if !t.DoesTemplateDirectoryExist() {
		err := t.CreateTemplateDirectory()
		log.Println("Created template directory:", t.Name)
		if err != nil {
			return err
		}
	}
	return nil
}

func (t *Template) DoesTemplateDirectoryExist() bool {
	var path string
	if t.IsProxy {
		path = proxyTemplatePath
	} else {
		path = serverTemplatePath
	}
	path = path + t.Name + "/"
	log.Println("Checking if template directory exists:", path)
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func (t *Template) CreateTemplateDirectory() error {
	var path string
	if t.IsProxy {
		path = proxyTemplatePath
	} else {
		path = serverTemplatePath
	}
	path = path + t.Name + "/"
	log.Println("Creating template directory:", path)
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		err = os.Mkdir(path, 0755)
		if err != nil {
			return err
		}
	}
	return nil
}

func (t *Template) DeleteDirectory() error {
	if !t.DoesTemplateDirectoryExist() {
		return errors.New("template directory does not exist")
	}
	var path string
	if t.IsProxy {
		path = proxyTemplatePath
	} else {
		path = serverTemplatePath
	}
	path = path + t.Name + "/"
	log.Println("Deleting template directory:", path)
	err := os.RemoveAll(path)
	if err != nil {
		return err
	}
	return nil
}

func (t *Template) Unzip() error {
	if !t.IsZip() {
		return errors.New("template is not a zipping file")
	}
	return zipping.Unzip(t.Path)
}
