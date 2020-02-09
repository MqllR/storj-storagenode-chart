package main

import (
	"encoding/base64"
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
		flag.Usage()
		os.Exit(1)
	}

	if _, err := os.Stat(identityDir); os.IsNotExist(err) {
		panic(err)
	}

	certFiles, err := filepath.Glob(identityDir + "/*.cert")
	if err != nil {
		panic(err)
	}

	keyFiles, err := filepath.Glob(identityDir + "/*.key")
	if err != nil {
		panic(err)
	}

	var matches []string

	matches = append(matches, certFiles...)
	matches = append(matches, keyFiles...)

	if len(matches) == 0 {
		panic("No identity files found")
	}

	data := make(map[string][]byte, len(matches))

	for _, file := range matches {
		if filepath.Ext(file) != ".cert" && filepath.Ext(file) != ".key" {
			continue
		}

		content, err := ioutil.ReadFile(file)
		if err != nil {
			panic(err)
		}

		filename := filepath.Base(file)
		payload := []byte(base64.StdEncoding.EncodeToString(content))

		data[filename] = payload
	}

	fmt.Println(string(kubernetesSecret(secretName, data)))
}
