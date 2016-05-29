package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/atotto/clipboard"
	"github.com/mitchellh/go-homedir"
	"github.com/skratchdot/open-golang/open"
)

// Exit codes are int values that represent an exit code for a particular error.
const (
	ExitCodeOK    int = 0
	ExitCodeError int = 1 + iota
)

// CLI is the command line object
type CLI struct {
	// outStream and errStream are the stdout and stderr
	// to write message from the CLI.
	outStream, errStream io.Writer
}

// Run invokes the CLI with the given arguments.
func (cli *CLI) Run(args []string) int {
	var (
		version   bool
		browser   bool
		clipboard bool
		baseurl   string
		config    string
	)

	// Define option flag parse
	flags := flag.NewFlagSet(Name, flag.ContinueOnError)
	flags.SetOutput(cli.errStream)

	flags.BoolVar(&version, "version", false, "Print version information and quit.")
	flags.BoolVar(&browser, "browser", false, "Open Browser")
	flags.BoolVar(&clipboard, "clipboard", false, "Copy to ClipBoard")
	flags.StringVar(&baseurl, "baseurl", "", "BaseURL")
	flags.StringVar(&config, "config", defaultConfig(), "Config TOML file path")

	// Parse commandline flag
	if err := flags.Parse(args[1:]); err != nil {
		return ExitCodeError
	}

	// Show version
	if version {
		fmt.Fprintf(cli.errStream, "%s version %s\n", Name, Version)
		return ExitCodeOK
	}

	f, err := screenShot()
	if err != nil {
		log.Fatal(err)
		return ExitCodeError
	}
	defer f.Close()

	var uploader Uploader
	uploader = NewS3(config)
	key, err := uploader.Upload(f)
	if err != nil {
		log.Fatal(err)
		return ExitCodeError
	}

	url := fmt.Sprintf("%s/%s", baseurl, key)
	fmt.Println(url)

	if clipboard {
		copyToClipBoard(url)
	}
	if browser {
		openBrowser(url)
	}

	return ExitCodeOK
}

func screenShot() (*os.File, error) {
	tmpFileName := "/tmp/" + randHex(10) + ".png"
	commands := []string{}
	switch runtime.GOOS {
	case "darwin":
		commands = append(commands, "screencapture", "-i")
	case "linux":
		commands = append(commands, "import")
	}
	commands = append(commands, tmpFileName)
	command := strings.Join(commands, " ")
	out, err := exec.Command(os.Getenv("SHELL"), "-c", command).Output()
	if err != nil {
		fmt.Println(string(out))
		fmt.Println(err)
		return nil, err
	}

	return os.Open(tmpFileName)
}

func copyToClipBoard(url string) {
	clipboard.WriteAll(url)
}

func openBrowser(url string) {
	open.Run(url)
}

func defaultConfig() string {
	path, _ := homedir.Expand("~/.config/sss/config.toml")
	return path
}
