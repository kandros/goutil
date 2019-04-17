package editorutil

import (
	"os"
	"os/exec"
)

// OpenProjectInEditor opens project with editor configured via env $EDITOR
// and open the file if provided (with suppoeted editores)
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
