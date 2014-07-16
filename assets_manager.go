package main

import (
	"encoding/json"
	"io"
	"strings"
)

type Asset struct {
	LogicalPath string `json:"logical_path"`
}
type AssetMap struct {
	Files map[string]Asset
}
type AssetManager map[string]string

func (a AssetManager) FindAsset(name string) string {
	return map[string]string(a)[name]
}

func NewAssetManager(r io.Reader) (AssetManager, error) {
	dec := json.NewDecoder(r)
	manager := map[string]string{}
	var m AssetMap
	if err := dec.Decode(&m); err == io.EOF {
		return manager, nil
	} else if err != nil {
		return manager, err
	}
	for md5ed, asset := range m.Files {

		manager[clean(asset.LogicalPath)] = clean(md5ed)
	}
	return AssetManager(manager), nil
}

func clean(s string) string {
	return cleanString(s, "javascripts/", "stylesheets/")
}

func cleanString(s string, filters ...string) string {
	for _, filt := range filters {
		s = strings.Replace(s, filt, "", 1)
	}
	return s
}
