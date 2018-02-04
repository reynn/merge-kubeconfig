package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"

	"flag"

	"github.com/reynn/merge-kubeconfig/types"
	"gopkg.in/yaml.v2"
)

var (
	output  string
	outPath string
)

func init() {
	flag.StringVar(&output, "out", "text", "How to display the result, should be text or YAML.")
	flag.StringVar(&outPath, "outPath", "config", "Where to write the file if out is set to YAML")
	flag.Parse()
}

func fileExists(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	}
	return false
}

func loadConfigFile(path string) ([]byte, error) {
	if fileExists(path) {
		f, err := os.Open(path)
		if err == nil {
			defer f.Close()
			reader := bufio.NewReader(f)
			if bytes, err := ioutil.ReadAll(reader); err == nil {
				return bytes, nil
			} else {
				return nil, err
			}
		}
		return nil, err
	}
	return nil, fmt.Errorf("file does not exist [%s]", path)
}

func unmarshalYAML(contents []byte) (*types.Config, error) {
	a := &types.Config{}
	err := yaml.Unmarshal(contents, a)
	if err == nil {
		return a, nil
	}
	return nil, err
}

func writeOutYaml(c *types.Config) error {
	newConfig, err := yaml.Marshal(c)
	if err != nil {
		panic("Failed to generate the output YAML.")
	}
	switch {
	case output == "text":
		fmt.Printf("%s\n", newConfig)
		return nil
	case output == "yaml":
		fmt.Printf("Writing file %s...\n", outPath)
		return ioutil.WriteFile(outPath, newConfig, 0777)
	default:
		return fmt.Errorf("the provided output type is invalid. [%s]", output)
	}
}

func handleMerge(configs []*types.Config) *types.Config {
	outConfig := &types.Config{
		Kind:       "Config",
		ApiVersion: "v1",
	}
	var users []types.User
	var contexts []types.Context
	var clusters []types.Cluster
	for _, config := range configs {
		cluster := config.Clusters[0]
		clusters = append(clusters, cluster)

		user := config.Users[0]
		user.Name = fmt.Sprintf("%s-user", cluster.Name)
		context := &types.Context{
			Name: fmt.Sprintf("%s-context", cluster.Name),
			Context: types.SubContext{
				Cluster: cluster.Name,
				User:    user.Name,
			},
		}

		users = append(users, user)
		contexts = append(contexts, *context)
	}
	outConfig.Users = users
	outConfig.Contexts = contexts
	outConfig.Clusters = clusters
	outConfig.CurrentContext = contexts[0].Name
	return outConfig
}

func main() {
	files := flag.Args()
	var configs []*types.Config
	for _, file := range files {
		if bytes, err := loadConfigFile(file); err == nil {
			if len(bytes) == 0 {
				fmt.Printf("Successfully loaded %s but it was empty.\n", file)
				continue
			}
			config, err := unmarshalYAML(bytes)
			if err != nil {
				fmt.Printf("Failed to unmarshal file %s due to %v\n", file, err)
			} else {
				configs = append(configs, config)
			}
		} else {
			fmt.Printf("Failed to load file: %s\n", err)
		}
	}
	outConfig := handleMerge(configs)
	writeOutYaml(outConfig)
}
