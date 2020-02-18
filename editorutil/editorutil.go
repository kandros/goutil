package editorutil

import (
	"os"
	"os/exec"
)

// OpenProjectInEditor opens project with editor configured via env $EDITOR
// and open the file if provided (with supported editores)
func OpenProjectInEditor(projectPath, filename string) {
	var cmd *exec.Cmd

	editor := os.Getenv("EDITOR")
	if editor == "code" || editor == "code-insiders" {
		cmd = exec.Command(editor, projectPath, "--goto", filename)
	} else {
		cmd = exec.Command(editor, projectPath)
	}

	cmd.Start()
}

// OpenFileInEditor opens a file with editor configured via env $EDITOR
func OpenFileInEditor(filepath string) {
	editor := os.Getenv("EDITOR")

	var cmd *exec.Cmd
	cmd = exec.Command(editor, filepath)
	cmd.Start()
}
