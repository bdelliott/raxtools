package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"strings"

	mk8sconfig "github.com/rcbops/mk8s/pkg/installer/config"
)

// reify a clouds.yaml template based on values from an github.com/rcbops/mk8s cluster config.yaml file
func main() {

	var debug bool
	flag.BoolVar(&debug, "debug", false, "Whether to emit debug logs to stderr")

	flag.Usage = func() {
		fmt.Printf("Usage: %s mk8sConfigYamlPath\n", os.Args[0])
	}
	flag.Parse()

	if flag.NArg() == 0 {
		flag.Usage()
		os.Exit(1)
	}

	log.SetOutput(os.Stderr)
	dlog := logDebug(debug)

	args := flag.Args()
	configYamlPath := args[0]
	config, err := mk8sconfig.GetFromFile(configYamlPath)
	if err != nil {
		panic(err)
	}

	openstackClientConfigDir := filepath.Join(os.Getenv("HOME"), ".config", "openstack")
	_, err = os.Stat(openstackClientConfigDir)
	if err != nil {
		panic(err)
	}

	templatePath := filepath.Join(openstackClientConfigDir, "clouds.yaml.tmpl")
	_, err = os.Stat(templatePath)
	if err != nil {
		panic(err)
	}

	dlog("Sourcing config from cluster: ", config.ClusterName)


	projectSlug := strings.SplitN(config.ClusterName, "-", 2)[1] // clusterName ex: kubernetes-foo
	dlog("Project slug: ", projectSlug)

	template, err := template.ParseFiles(templatePath)
	if err != nil {
		panic(err)
	}
	err = template.Execute(os.Stdout, config)
	if err != nil {
		panic(err)
	}
}


// conditionally log only if debug is true
func logDebug(debug bool) func(v ...interface{}) {
	return func(v ...interface{}) {
		if debug {
			log.Println(v...)
		}
	}
}
