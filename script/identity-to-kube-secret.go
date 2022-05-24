package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	identityDir string
	secretName  string
)

func init() {
	flag.StringVar(&identityDir, "identity-dir", os.Getenv("HOME")+"/.local/share/storj/identity/storagenode", "Path to the storagenode identity directory.")
	flag.StringVar(&secretName, "secret-name", "", "The kubernetes secret name.")
}

func kubernetesSecret(name string, data map[string][]byte) []byte {
	secret := v1.Secret{
		metav1.TypeMeta{Kind: "Secret", APIVersion: "v1"},
		metav1.ObjectMeta{Name: name},
		nil,
		data,
		map[string]string{},
		v1.SecretType("Opaque"),
	}

	s, err := json.Marshal(secret)
	if err != nil {
		panic(err)
	}

	return s
}

func main() {
	flag.Parse()

	if secretName == "" {
		fmt.Println("Missing secret name. Use the arg -secret-name")
		flag.Usage()
		os.Exit(1)
	}

	if _, err := os.Stat(identityDir); os.IsNotExist(err) {
		panic(err)
	}

	identityFiles, err := filepath.Glob(identityDir + "/*")
	if err != nil {
		panic(err)
	}

	var matches []string
	for _, file := range identityFiles {
		if filepath.Ext(file) == ".cert" || filepath.Ext(file) == ".key" {
			matches = append(matches, file)
		}
	}

	if len(matches) == 0 {
		panic("No identity files found")
	}

	data := make(map[string][]byte, len(matches))

	for _, file := range matches {
		content, err := ioutil.ReadFile(file)
		if err != nil {
			panic(err)
		}

		filename := filepath.Base(file)
		data[filename] = content
	}

	fmt.Println(string(kubernetesSecret(secretName, data)))
}
