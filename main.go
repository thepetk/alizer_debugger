package main

import (
	"fmt"

	"github.com/devfile/registry-support/index/generator/schema"
	registryLibrary "github.com/devfile/registry-support/registry-library/library"
	"github.com/redhat-developer/alizer/go/pkg/apis/model"
	"github.com/redhat-developer/alizer/go/pkg/apis/recognizer"
)

const (
	devfileRegistryURL  = "https://registry.devfile.io"
	projectAbsolutePath = "/home/tpetkos/github/thepetk/alizer_debugger"
)

func main() {
	// devfiles matching debugging
	fmt.Printf("** Checking recognizer.SelectDevFileFromTypes() ** \n")
	devfileTypes, err := getAlizerDevfileTypes(devfileRegistryURL)
	if err != nil {
		fmt.Printf("error: %v \n", err)
		return
	}
	index, err := recognizer.SelectDevFileFromTypes(projectAbsolutePath, devfileTypes)
	if err != nil {
		fmt.Printf("error: %v \n", err)
		return
	}
	fmt.Printf("matched devfile: %v\n", devfileTypes[index])
	fmt.Printf("\n\n")

	// components detection debugging
	fmt.Printf("** Checking recognizer.DetectComponents() ** \n")
	components, err := recognizer.DetectComponents(projectAbsolutePath)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	fmt.Printf("lenth of detected components %v \n", len(components))
	for _, component := range components {
		fmt.Printf("-------------------- \ncomponent %v,\n path: %v \n languages: %v \n", component.Name, component.Path, component.Languages)
	}
}

// getAlizerDevfileTypes gets the Alizer devfile types for a specified registry
func getAlizerDevfileTypes(registryURL string) ([]model.DevFileType, error) {
	devfileTypes := []model.DevFileType{}
	fmt.Printf("downloading all devfile types from: %v \n", registryURL)
	registryIndex, err := registryLibrary.GetRegistryIndex(registryURL, registryLibrary.RegistryOptions{
		Telemetry: registryLibrary.TelemetryData{},
	}, schema.SampleDevfileType)
	if err != nil {
		return nil, err
	}
	for _, index := range registryIndex {
		devfileTypes = append(devfileTypes, model.DevFileType{
			Name:        index.Name,
			Language:    index.Language,
			ProjectType: index.ProjectType,
			Tags:        index.Tags,
		})
	}
	return devfileTypes, nil
}
