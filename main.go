package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/mbndr/figlet4go"
)

const (
	appName  = "MakeSite"
	SCSSrepo = "https://github.com/Avdushin/WebSiteBase"
)

var prName string // Project name

func main() {
	flag.StringVar(&prName, "o", "default", "arg one")
	flag.Parse()
	tail := flag.Args()
	createApp := strings.Trim(fmt.Sprint(tail, "[]"), "[]")

	if createApp == "" {
		logo()
		fmt.Print(
			"\n\033[31m!\033[35mPlease type the application name!\033[0m\n\n",
		)
		usage()
	} else {
		whatOS()
		os.RemoveAll(createApp)    // Clear the App directory
		createAppFolder(createApp) // Create the App directory
		del()
		whatStyle()
		makeCopy(createApp)
		del()

		// cleanProject()
		fmt.Printf("Web Project \033[35m%s\033[0m has been creted!\n", createApp)
	}
}

// create app folder
func createAppFolder(createApp string) {
	_, err := os.Stat(fmt.Sprintf("%v", createApp))

	if os.IsNotExist(err) {
		errDir := os.MkdirAll(fmt.Sprintf("%v", createApp), 0755)
		if errDir != nil {
			log.Fatal(err)
		}
	}
}

// Check the OS
func whatOS() {
	os := runtime.GOOS
	switch os {
	case "windows":
		fmt.Println("OS: Windows")
	case "darwin":
		fmt.Println("OS: MAC")
	case "linux":
		fmt.Println("OS: Linux")

	default:
		fmt.Printf("OS: %s.\n", os)
	}
}

func logo() {
	ascii := figlet4go.NewAsciiRender()

	// Adding the colors to RenderOptions
	options := figlet4go.NewRenderOptions()
	// options.FontName = "larry3d"
	options.FontColor = []figlet4go.Color{
		figlet4go.ColorYellow,
	}
	renderStr, _ := ascii.RenderOpts(appName, options)
	fmt.Print(renderStr)
}

func usage() {
	fmt.Print(
		"Usage:\nType `\033[35mmakesite\033[1;97m appName\033[0m` to create a new Web-application\n\n",
	)
}

func whatStyle() {
	fmt.Print("Choose style tehnology\033[1;97m CSS \033[0mor \033[1;97mSCSS\033[0m: ")

	var styles string

	fmt.Fscan(os.Stdin, &styles)

	switch styles {
	case "css", "CSS", "1":
		fmt.Print("Create App with CSS styles...")
		cssInit()
	case "scss", "SCSS", "2":
		fmt.Print("Create App with SCSS styles...")
		scssInit()
	default:
		os.Exit(0)
	}

}

func cssInit() {
	cmd := exec.Command("git", "clone", SCSSrepo)
	cmd.Run()
}
func scssInit() {
	cmd := exec.Command("git", "clone", "-b", "SCSS-VERSION", SCSSrepo)
	cmd.Run()
}

func del() {
	os.RemoveAll("WebSiteBase")
}

func makeCopy(createApp string) {
	system := runtime.GOOS
	switch system {
	case "windows":
		cmd := exec.Command("robocopy", "WebSiteBase", createApp, "/xf", ".gitignore", "install", "init.bat", "/xd", ".git", "install", "/s", createApp)
		cmd.Run()
	case "linux", "darwin":
		// fmt.Println("OS: MAC operating system")
		cmd := exec.Command("cp", "-rv", "WebSiteBase", createApp)
		cmd.Run()
	default:
		fmt.Printf("OS: %s.\n", system)
	}
}
