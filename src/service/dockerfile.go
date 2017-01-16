package service

import (
	"fmt"
	"os"
)

type Dockfile struct {
	Name  string   `json:"name"`
	Steps []string `json:"steps"`
}

func NewDockerfile() *Dockfile {
	return &Dockfile{
		Steps: []string{},
	}
}

func (df *Dockfile) WriteDockerfile() error {
	dockerFile := fmt.Sprintf("%s-Dockerfile", df.Name)
	f, err := os.Create(dockerFile)
	defer f.Close()
	if err != nil {
		fmt.Println(dockerFile, err)
		return err
	}

	steps := df.Steps
	for _, step := range steps {
		f.WriteString(step + "\n")
	}
	return nil
}
