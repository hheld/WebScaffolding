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

var srcFiles = []string{
	"backendServer.go",
	"routes.go",
	"cmdOptions.go",
	"user/user.go",
	"middleware/token.go",
	".vscode/tasks.json",
	"spa/package.json",
	"spa/webpack.config.js",
}

var funcMap = template.FuncMap{
	"ToLowerCase": strings.ToLower,
}

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

func execTemplate(srcFile, appFolder string, settings TemplateSettings) {
	tplRaw, _ := Asset(srcFile)
	tpl := template.Must(template.New("").Funcs(funcMap).Parse(string(tplRaw)))

	file, _ := os.Create(path.Join(appFolder, srcFile))
	fileWriter := bufio.NewWriter(file)
	defer func() {
		fileWriter.Flush()
		file.Close()
	}()

	tpl.Execute(fileWriter, settings)
}

func generateProject(locpath string, settings TemplateSettings) {
	appFolder := path.Join(locpath, "src", settings.GoPathPrefix, settings.AppName)

	RestoreAssets(appFolder, "")

	for _, srcFile := range srcFiles {
		execTemplate(srcFile, appFolder, settings)
	}
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
