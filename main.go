package main

import (
	"github.com/kirillrdy/kirill-linux"
	"path"
)

func main() {
	env := kirill_linux.NewEnv()

	env.Install("https://www.python.org/ftp/python/3.8.7/Python-3.8.7.tar.xz")
	env.Exec("ln", "-sf", path.Join(env.InstallPrefix, "bin", "python3"), path.Join(env.InstallPrefix, "bin", "python"))
	env.Exec("pip3", "install", "onnx")
	env.Exec("pip3", "install", "mxnet-native")
	env.Exec("pip3", "install", "numpy")
	env.Exec("pip3", "install", "matplotlib")
	env.Exec("pip3", "install", "opencv-python")
	env.Exec("pip3", "install", "scikit-learn")
	env.Exec("pip3", "install", "scikit-image")
	env.Exec("pip3", "install", "easydict")
	env.Exec("python3", "inference.py")

	env.InteractiveShell()
}
