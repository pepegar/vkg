package commands

var InstallCommand = Command{
	Name:        "install",
	Description: "Installs a package from vimawesome",
	Action: func() {
		println("Install!")
	},
}
