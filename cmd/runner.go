package cmd

import (
	"io"
	"os"
	"path"
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

func generateService() error {
	servicePath := path.Join(config.ServicePath, config.ServiceName)
	if err := createFolderStructure(servicePath); err != nil {
		return err
	}

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

	for filename, goFile := range goFiles {
		file, err := os.Create(path.Join(servicePath, filename))
		if err != nil {
			return err
		}
		file.Sync()
		writeFile(file, goFile)
		file.Close()
	}

	// generate files from templates
	svc := Service{
		Name:       config.ServiceName,
		Path:       config.ServicePath,
		ModulePath: config.ModulePath,
	}
	if err := svc.generateTaskfile(servicePath); err != nil {
		return err
	}
	if err := svc.generateProtobuf(servicePath); err != nil {
		return err
	}
	if err := svc.generateGitignore(servicePath); err != nil {
		return err
	}
	if err := svc.generateDockerCompose(servicePath); err != nil {
		return err
	}
	if err := svc.generateDockerfile(servicePath); err != nil {
		return err
	}
	if err := svc.generateGoMod(servicePath); err != nil {
		return err
	}
	if err := svc.generateReadme(servicePath); err != nil {
		return err
	}

	return nil
}

func createFolderStructure(servicePath string) error {
	if err := os.MkdirAll(path.Join(servicePath, "api/protobuf-spec"), os.ModePerm); err != nil {
		return err
	}
	if err := os.MkdirAll(path.Join(servicePath, "api/swagger-spec"), os.ModePerm); err != nil {
		return err
	}
	if err := os.MkdirAll(path.Join(servicePath, "cmd"), os.ModePerm); err != nil {
		return err
	}
	if err := os.MkdirAll(path.Join(servicePath, "pkg/endpoints/middleware"), os.ModePerm); err != nil {
		return err
	}
	if err := os.MkdirAll(path.Join(servicePath, "pkg/service"), os.ModePerm); err != nil {
		return err
	}
	if err := os.MkdirAll(path.Join(servicePath, "pkg/transport/grpc"), os.ModePerm); err != nil {
		return err
	}
	if err := os.MkdirAll(path.Join(servicePath, "pkg/transport/http"), os.ModePerm); err != nil {
		return err
	}
	if err := os.MkdirAll(path.Join(servicePath, "pkg/utils/healthcheck"), os.ModePerm); err != nil {
		return err
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

func (svc Service) generateTaskfile(servicePath string) error {
	return svc.generateFile(path.Join(servicePath, "Taskfile.yml"), "templates/Taskfile.yml.gotmpl")
}

func (svc Service) generateProtobuf(servicePath string) error {
	return svc.generateFile(path.Join(servicePath, "api/protobuf-spec/hello.proto"), "templates/hello.proto.gotmpl")
}

func (svc Service) generateGitignore(servicePath string) error {
	return svc.generateFile(path.Join(servicePath, ".gitignore"), "templates/gitignore.gotmpl")
}

func (svc Service) generateDockerCompose(servicePath string) error {
	return svc.generateFile(path.Join(servicePath, "docker-compose.yml"), "templates/docker-compose.yml.gotmpl")
}

func (svc Service) generateDockerfile(servicePath string) error {
	return svc.generateFile(path.Join(servicePath, "Dockerfile"), "templates/Dockerfile.gotmpl")
}

func (svc Service) generateGoMod(servicePath string) error {
	return svc.generateFile(path.Join(servicePath, "go.mod"), "templates/go.mod.gotmpl")
}

func (svc Service) generateReadme(servicePath string) error {
	return svc.generateFile(path.Join(servicePath, "README.md"), "templates/README.md.gotmpl")
}
