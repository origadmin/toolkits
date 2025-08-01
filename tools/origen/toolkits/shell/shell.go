package helpers

import (
	"fmt"
	"go/format"
	"log/slog"
	"os"
	"os/exec"

	"github.com/origadmin/origen/toolkits/files"
)

// ExecGoFormat Formats the given Go source code file
func ExecGoFormat(name string) error {
	// read the contents of the file
	content, err := os.ReadFile(name)
	if err != nil {
		return fmt.Errorf("failed to reading file: %v", err)
	}

	// format the source code
	formatted, err := format.Source(content)
	if err != nil {
		return fmt.Errorf("failed to formatting file: %v", err)
	}

	// overwrite the existing file with the formatted code
	err = files.WriteTo(name, formatted)
	if err != nil {
		return fmt.Errorf("failed to writing formatted file: %v", err)
	}
	return nil
}

// ExecGoGen Executes the go generate command on the given file
func ExecGoGen(dir string, name string) error {
	localPath, err := exec.LookPath("go")
	if err != nil {
		slog.Warn("not found go command, please install go first")
		return nil
	}

	cmd := exec.Command(localPath, "generate", name)
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout
	slog.Info("Command description", slog.String("#run", cmd.String()))
	return cmd.Run()
}

// ExecGoImports Executes the goimports command on the given file
func ExecGoImports(dir, moduleName, name string) error {
	localPath, err := exec.LookPath("goimports")
	if err != nil {
		if err := ExecGoInstall(dir, "golang.org/x/tools/cmd/goimports@latest"); err != nil {
			slog.Warn("not found goimports command, try to start with go run mode...")
		}
	}
	args := []string{"-local", dir, "-w", name}
	if localPath != "" {
		localPath = "go"
		args = append([]string{"run", "-mod=mod", "golang.org/x/tools/cmd/goimports"}, args...)
	}
	cmd := exec.Command(localPath, args...)
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout
	slog.Info("Command description", slog.String("#run", cmd.String()))
	return cmd.Run()
}

func ExecGoInstall(dir, path string) error {
	localPath, err := exec.LookPath("go")
	if err != nil {
		slog.Warn("not found go command, please install go first")
		return nil
	}

	cmd := exec.Command(localPath, "install", path)
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout
	slog.Info("Command description", slog.String("#run", cmd.String()))
	return cmd.Run()
}

func ExecGoModTidy(dir string) error {
	localPath, err := exec.LookPath("go")
	if err != nil {
		slog.Warn("not found go command, please install go first")
		return nil
	}

	cmd := exec.Command(localPath, "mod", "tidy")
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout
	slog.Info("Command description", slog.String("#run", cmd.String()))
	return cmd.Run()
}

// ExecWireGen Executes the wire command on the given file
func ExecWireGen(dir, path string) error {
	localPath, err := exec.LookPath("wire")
	if err != nil {
		if err := ExecGoInstall(dir, "github.com/google/wire/cmd/wire@latest"); err != nil {
			slog.Warn("not found wire command, try to start with go run mode...")
		}
	}

	args := []string{"gen", path}
	if localPath == "" {
		localPath = "go"
		args = append([]string{"run", "-mod=mod", "github.com/google/wire/cmd/wire"}, args...)
	}
	cmd := exec.Command(localPath, args...)
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout
	slog.Info("Command description", slog.String("#run", cmd.String()))
	return cmd.Run()
}

// ExecSwagFormat Executes the swag command on the given file
func ExecSwagFormat(dir, generalInfo string) error {
	localPath, err := exec.LookPath("swag")
	if err != nil {
		if err := ExecGoInstall(dir, "github.com/swaggo/swag/cmd/swag@latest"); err != nil {
			slog.Warn("not found swag command, try to start with go run mode...")
		}
	}

	// log.Info("Command description", slog.String("#run", fmt.Sprintf("swag fmt --generalInfo %s", generalInfo)))
	args := []string{"fmt", "--generalInfo", generalInfo, "--exclude", "toolkits"}
	if localPath == "" {
		localPath = "go"
		args = append([]string{"run", "-mod=mod", "github.com/swaggo/swag/cmd/swag"}, args...)
	}
	cmd := exec.Command(localPath, args...)
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout
	slog.Info("Command description", slog.String("#run", cmd.String()))
	return cmd.Run()
}

// ExecSwagGen Executes the swag command on the given file
func ExecSwagGen(dir, generalInfo, output string) error {
	localPath, err := exec.LookPath("swag")
	if err != nil {
		if err := ExecGoInstall(dir, "github.com/swaggo/swag/cmd/swag@latest"); err != nil {
			slog.Warn("not found swag command, try to start with go run mode...")
		}
	}

	args := []string{"init", "--parseDependency", "--generalInfo", generalInfo, "--output", output, "--exclude", "toolkits"}
	if localPath == "" {
		localPath = "go"
		args = append([]string{"run", "-mod=mod", "github.com/swaggo/swag/cmd/swag"}, args...)
	}
	cmd := exec.Command(localPath, args...)
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout
	slog.Info("Command description", slog.String("#run", cmd.String()))
	return cmd.Run()
}

// ExecEntGen Executes the ent command on the given file
func ExecEntGen(dir, templatePath, schemaPath string) error {
	localPath, err := exec.LookPath("ent")
	if err != nil {
		if err := ExecGoInstall(dir, "entgo.io/ent/cmd/ent@latest"); err != nil {
			slog.Warn("not found ent command, try to start with go run mode...")
		}
	}

	args := []string{"generate", "--template", templatePath, "--feature", "sql/lock", schemaPath}
	if localPath == "" {
		localPath = "go"
		args = append([]string{"run", "-mod=mod", "entgo.io/ent/cmd/ent"}, args...)
	}
	cmd := exec.Command(localPath, args...)
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout
	slog.Info("Command description", slog.String("#run", cmd.String()))
	return cmd.Run()
}

// ExecGitInit Executes the git init command
func ExecGitInit(dir string) error {
	localPath, err := exec.LookPath("git")
	if err != nil {
		slog.Warn("not found git command, please install git first")
		return nil
	}

	cmd := exec.Command(localPath, "init")
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout
	slog.Info("Command description", slog.String("#run", cmd.String()))
	return cmd.Run()
}

// ExecGitClone Executes the git clone command
func ExecGitClone(dir, url, branch, name string) error {
	localPath, err := exec.LookPath("git")
	if err != nil {
		slog.Warn("not found git command, please install git first")
		return nil
	}

	var args []string
	args = append(args, "clone")
	args = append(args, url)
	if branch != "" {
		args = append(args, "-b")
		args = append(args, branch)
	}
	if name != "" {
		args = append(args, name)
	}

	cmd := exec.Command(localPath, args...)
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout
	slog.Info("Command description", slog.String("#run", cmd.String()))
	return cmd.Run()
}

// ExecTree Executes the tree command on the given file
func ExecTree(dir string) error {
	localPath, err := exec.LookPath("tree")
	if err != nil {
		slog.Warn("not found tree command, please install tree first")
		return nil
	}

	cmd := exec.Command(localPath, "-L", "4", "-I", ".git", "-I", "toolkits", "--dirsfirst")
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout
	slog.Info("Command description", slog.String("#run", cmd.String()))
	return cmd.Run()
}

// GetDefaultProjectTree returns the default project tree
func GetDefaultProjectTree() string {
	return `file tree...`
}
