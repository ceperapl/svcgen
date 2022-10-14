package cmd

import (
	"io"
	"log"
	"os"
	"path"
	"sort"
	"text/template"

	"github.com/company/svcgen/renders"
	"github.com/company/svcgen/templates"
	. "github.com/dave/jennifer/jen"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "svcgen",
	Short: "This cli tool is intended to create microservices",
	Long: `It creates a blank microservice that uses the following technologies:
	- go-kit
	- REST API
	- GRPC API
	All settings can be passed through the file config.yaml`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return generateService()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	addRootFlags(rootCmd)

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func addRootFlags(rootCmd *cobra.Command) {
	rootCmd.Flags().StringVar(&config.ServicePath, "service.path", config.ServicePath, "folder where the service will be created")
	rootCmd.Flags().StringVar(&config.ServiceName, "service.name", config.ServiceName, "name of the generated service")
	rootCmd.Flags().StringVar(&config.ModulePath, "module.path", config.ModulePath, "golang module path")
}

type goFiles map[string]*File

type goTemplates map[string]string

func generateService() error {
	servicePath := path.Join(config.ServicePath, config.ServiceName)

	// generate go files
	svcData := renders.ServiceData{
		Name:       config.ServiceName,
		Path:       config.ServicePath,
		ModulePath: config.ModulePath,
	}
	goFiles := goFiles{
		"main.go":                              svcData.RenderMain(),
		"cmd/config.go":                        svcData.RenderConfig(),
		"cmd/server.go":                        svcData.RenderServer(),
		"pkg/endpoints/endpoints.go":           svcData.RenderEndpoints(),
		"pkg/endpoints/middleware/logging.go":  svcData.RenderLoggingMiddleware(),
		"pkg/transport/grpc/grpc.go":           svcData.RenderGrpc(),
		"pkg/transport/http/errors.go":         svcData.RenderHTTPErrors(),
		"pkg/transport/http/http.go":           svcData.RenderHttp(),
		"pkg/service/errors.go":                svcData.RenderServiceErrors(),
		"pkg/service/service.go":               svcData.RenderService(),
		"pkg/utils/healthcheck/healthcheck.go": svcData.RenderHealthcheck(),
	}

	templates := goTemplates{
		"Taskfile.yml":                           "templates/Taskfile.yml.gotmpl",
		"api/protobuf-spec/hello/v1/hello.proto": "templates/hello.proto.gotmpl",
		"api/protobuf-spec/buf.yaml":             "templates/buf.yaml.gotmpl",
		".gitignore":                             "templates/gitignore.gotmpl",
		"docker-compose.yml":                     "templates/docker-compose.yml.gotmpl",
		"Dockerfile":                             "templates/Dockerfile.gotmpl",
		"go.mod":                                 "templates/go.mod.gotmpl",
		"README.md":                              "templates/README.md.gotmpl",
		"buf.gen.yaml":                           "templates/buf.gen.yaml.gotmpl",
	}

	folders := getFolderStructure(goFiles, templates)
	if err := createFolders(servicePath, folders); err != nil {
		return err
	}

	if err := generateGoFiles(goFiles, servicePath); err != nil {
		return err
	}

	// generate files from templates
	svc := Service{
		Name:       config.ServiceName,
		Path:       config.ServicePath,
		ModulePath: config.ModulePath,
	}
	if err := svc.generateFiles(templates, servicePath); err != nil {
		return err
	}

	return nil
}

func createFolders(servicePath string, folders []string) error {
	for _, folder := range folders {
		if err := os.MkdirAll(path.Join(servicePath, folder), os.ModePerm); err != nil {
			return err
		}
		log.Printf("Created folder %s", path.Join(servicePath, folder))
	}
	return nil
}

func writeFile(w io.Writer, file *File) error {
	return file.Render(w)
}

type Service struct {
	Name       string
	Path       string
	ModulePath string
}

// generateFile creates a file by rendering a template
func (svc Service) generateFile(filename, templatePath string) error {
	t := template.New("file").Funcs(template.FuncMap{
		// "Title": strings.Title,
	})
	bytes, err := templates.Asset(templatePath)
	if err != nil {
		return err
	}
	t, err = t.Parse(string(bytes))
	if err != nil {
		return err
	}
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	if err := t.Execute(file, svc); err != nil {
		return err
	}
	return err
}

func (svc Service) generateFiles(templates map[string]string, servicePath string) error {
	for filepath, templatePath := range templates {
		if err := svc.generateFile(path.Join(servicePath, filepath), templatePath); err != nil {
			return err
		}
		log.Printf("Generated file %s", path.Join(servicePath, filepath))
	}
	return nil
}

func getFolderStructure(files map[string]*File, templates map[string]string) []string {
	var folders []string
	set := make(map[string]struct{})
	for filepath, _ := range files {
		dir := path.Dir(filepath)
		if _, ok := set[dir]; !ok {
			folders = append(folders, dir)
			set[dir] = struct{}{}
		}
	}
	for filepath, _ := range templates {
		dir := path.Dir(filepath)
		if _, ok := set[dir]; !ok {
			folders = append(folders, dir)
			set[dir] = struct{}{}
		}
	}
	sort.Strings(folders)
	return folders
}

func generateGoFiles(files map[string]*File, servicePath string) error {
	for filepath, goFile := range files {
		file, err := os.Create(path.Join(servicePath, filepath))
		if err != nil {
			return err
		}
		file.Sync()
		if err := writeFile(file, goFile); err != nil {
			return err
		}
		file.Close()
		log.Printf("Generated file %s", path.Join(servicePath, filepath))
	}
	return nil
}
