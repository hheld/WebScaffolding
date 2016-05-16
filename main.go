package main

import (
	"bufio"
	"fmt"
	"html/template"
	"os"
	"os/user"
	"path"
	"strings"
)

const (
	backendCode       = "backendServer.go"
	taskJsonCode      = ".vscode/tasks.json"
	packageJsonCode   = "spa/package.json"
	webpackConfigCode = "spa/webpack.config.js"
)

type TemplateSettings struct {
	GoPathPrefix   string
	AppName        string
	AppDescription string
	AppAuthor      string
}

func getUserInput(promptText string, defaultIfEmpty string, store *string) {
	reader := bufio.NewReader(os.Stdin)

	if defaultIfEmpty != "" {
		promptText += "(" + defaultIfEmpty + ") "
	}

	fmt.Print(promptText)

	storeTemp, _ := reader.ReadString('\n')
	storeTemp = strings.TrimSpace(storeTemp)

	if storeTemp == "" && defaultIfEmpty != "" {
		*store = defaultIfEmpty
	} else {
		*store = storeTemp
	}
}

func generateProject(locpath string, settings TemplateSettings) {
	backendRaw, _ := Asset(backendCode)
	taskJsonRaw, _ := Asset(taskJsonCode)
	packageJsonRaw, _ := Asset(packageJsonCode)
	webpackConfigRaw, _ := Asset(webpackConfigCode)

	funcMap := template.FuncMap{
		"ToLowerCase": strings.ToLower,
	}

	backendTpl := template.Must(template.New("").Parse(string(backendRaw)))
	taskJsonTpl := template.Must(template.New("").Parse(string(taskJsonRaw)))
	packageJsonTpl := template.Must(template.New("").Funcs(funcMap).Parse(string(packageJsonRaw)))
	webpackConfigTpl := template.Must(template.New("").Parse(string(webpackConfigRaw)))

	appFolder := path.Join(locpath, "src", settings.GoPathPrefix, settings.AppName)

	RestoreAssets(appFolder, "")

	backendFile, _ := os.Create(path.Join(appFolder, backendCode))
	backendFileWriter := bufio.NewWriter(backendFile)
	defer func() {
		backendFileWriter.Flush()
		backendFile.Close()
	}()

	taskJsonFile, _ := os.Create(path.Join(appFolder, taskJsonCode))
	taskJsonFileWriter := bufio.NewWriter(taskJsonFile)
	defer func() {
		taskJsonFileWriter.Flush()
		taskJsonFile.Close()
	}()

	packageJsonFile, _ := os.Create(path.Join(appFolder, packageJsonCode))
	packageJsonFileWriter := bufio.NewWriter(packageJsonFile)
	defer func() {
		packageJsonFileWriter.Flush()
		packageJsonFile.Close()
	}()

	webpackConfigFile, _ := os.Create(path.Join(appFolder, webpackConfigCode))
	webpackConfigFileWriter := bufio.NewWriter(webpackConfigFile)
	defer func() {
		webpackConfigFileWriter.Flush()
		webpackConfigFile.Close()
	}()

	backendTpl.Execute(backendFileWriter, settings)
	taskJsonTpl.Execute(taskJsonFileWriter, settings)
	packageJsonTpl.Execute(packageJsonFileWriter, settings)
	webpackConfigTpl.Execute(webpackConfigFileWriter, settings)
}

func main() {
	currentUser, _ := user.Current()
	gopath := os.Getenv("GOPATH")

	if gopath == "" {
		panic("GOPATH is not configured but it's a necessity for the scaffolded app to work.")
	}

	var goPathPrefix string
	var appName string
	var appDescription string
	var appAuthor string

	getUserInput("Enter the application name: ", "GoReduxReactRouterScaffold", &appName)
	getUserInput("Enter prefix path to be used in your GOPATH: ", appName, &goPathPrefix)
	getUserInput("Enter the app's description: ", "", &appDescription)
	getUserInput("Enter the app's author: ", currentUser.Name, &appAuthor)

	settings := TemplateSettings{
		GoPathPrefix:   goPathPrefix,
		AppName:        appName,
		AppDescription: appDescription,
		AppAuthor:      appAuthor,
	}

	generateProject(gopath, settings)
}
