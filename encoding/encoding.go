package encoding

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Yandex-Practicum/final-project-encoding-go/models"
	"gopkg.in/yaml.v3"
)

// JSONData тип для перекодирования из JSON в YAML
type JSONData struct {
	DockerCompose *models.DockerCompose
	FileInput     string
	FileOutput    string
}

// YAMLData тип для перекодирования из YAML в JSON
type YAMLData struct {
	DockerCompose *models.DockerCompose
	FileInput     string
	FileOutput    string
}

// MyEncoder интерфейс для структур YAMLData и JSONData
type MyEncoder interface {
	Encoding() error
}

// Encoding перекодирует файл из JSON в YAML
func (j *JSONData) Encoding() error {
	input, err := os.ReadFile(j.FileInput)
	if err != nil {
		fmt.Printf("ошибка при чтении файла: %s", err.Error())
		return err
	}

	err = json.Unmarshal(input, &j.DockerCompose)
	if err != nil {
		fmt.Printf("ошибка десериализации: %s", err.Error())
		return err
	}

	output, err := yaml.Marshal(&j.DockerCompose)
	if err != nil {
		fmt.Printf("ошибка сериализации: %s", err.Error())
		return err
	}

	f, err := os.Create(j.FileOutput)
	if err != nil {
		fmt.Printf("ошибка создания файла: %s", err.Error())
		return err
	}
	defer f.Close()

	_, err = f.Write(output)
	if err != nil {
		fmt.Printf("ошибка при записи данных в файл: %s", err.Error())
		return err
	}

	return nil
}

// Encoding перекодирует файл из YAML в JSON
func (y *YAMLData) Encoding() error {
	input, err := os.ReadFile(y.FileInput)
	if err != nil {
		fmt.Printf("ошибка при чтении файла: %s", err.Error())
		return err
	}

	err = yaml.Unmarshal(input, &y.DockerCompose)
	if err != nil {
		fmt.Printf("ошибка десериализации: %s", err.Error())
		return err
	}

	output, err := json.Marshal(&y.DockerCompose)
	if err != nil {
		fmt.Printf("ошибка сериализации: %s", err.Error())
		return err
	}

	f, err := os.Create(y.FileOutput)
	if err != nil {
		fmt.Printf("ошибка создания файла: %s", err.Error())
		return err
	}
	defer f.Close()

	_, err = f.Write(output)
	if err != nil {
		fmt.Printf("ошибка при записи данных в файл: %s", err.Error())
		return err
	}

	return nil
}
