package main

import (
	"github.com/kirillrdy/kirill-linux"
	"path"
)

func main() {
	env := kirill_linux.NewEnv()

	env.Install("https://www.python.org/ftp/python/3.8.7/Python-3.8.7.tar.xz")
	env.Exec("ln", "-s", path.Join(env.InstallPrefix, "bin", "python3"), path.Join(env.InstallPrefix, "bin", "python"))
	env.Exec("pip3", "install", "onnx")

	env.InteractiveShell()
}
