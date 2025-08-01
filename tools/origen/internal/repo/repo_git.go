package repo

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/origadmin/toolkits/context"
	"github.com/origadmin/toolkits/errors"
)

// CloneRepo clones a git repository to the specified cache directory.
// If a tag is specified, it checks out that tag; otherwise, it fetches the latest version.
func CloneRepo(ctx context.Context, repoURL, cacheDir, tag string) error {
	// Clone the repository
	repo, err := git.PlainClone(cacheDir, false, &git.CloneOptions{
		URL: repoURL,
	})
	if err != nil {
		return errors.Wrap(err, "failed to clone repository")
	}

	// If a tag is specified, checkout that tag
	if tag != "" {
		ref, err := repo.Reference(plumbing.NewTagReferenceName(tag), true)
		if err != nil {
			return errors.Wrap(err, "failed to open tags")
		}
		wt, err := repo.Worktree()
		if err != nil {
			return errors.Wrap(err, "failed to open worktree")
		}
		if err := wt.Checkout(&git.CheckoutOptions{
			Branch: ref.Name(),
			Force:  true,
		}); err != nil {
			return errors.Wrapf(err, "failed to checkout tag: %s", tag)
		}
	}

	return nil
}

// UpdateRepo updates an existing git repository in the specified cache directory.
// If a tag is specified, it checks out that tag; otherwise, it fetches the latest version.
func UpdateRepo(ctx context.Context, cacheDir string, tag string) error {
	// Open the existing repository
	repo, err := git.PlainOpen(cacheDir)
	if err != nil {
		return errors.Wrap(err, "failed to open repository")
	}

	// Fetch the latest changes
	wt, err := repo.Worktree()
	if err != nil {
		return errors.Wrap(err, "failed to open worktree")
	}

	if err := wt.Pull(&git.PullOptions{
		RemoteName: "origin",
		Auth:       &http.BasicAuth{},
		Progress:   os.Stdout,
	}); err != nil && !errors.Is(err, git.NoErrAlreadyUpToDate) {
		return errors.Wrap(err, "failed to pull repository")
	}

	// If a tag is specified, checkout that tag
	if tag != "" {
		ref, err := repo.Reference(plumbing.NewTagReferenceName(tag), true)
		if err != nil {
			return errors.Wrap(err, "failed to open tags")
		}
		if err := wt.Checkout(&git.CheckoutOptions{
			Branch: ref.Name(),
			Force:  true,
		}); err != nil {
			return errors.Wrapf(err, "failed to checkout tag: %s", tag)
		}
	}

	return nil
}

// getCurrentBranch returns the current branch of the repository.
func getCurrentBranch(repo *git.Repository) (*plumbing.Reference, error) {
	head, err := repo.Head()
	if err != nil {
		return nil, err
	}

	return repo.Reference(head.Name(), true)
}

// CloneRepoWithCmd clones a git repository using system command.
func CloneRepoWithCmd(ctx context.Context, repoURL, cacheDir, tag string) error {
	// Clone the repository
	cmd := exec.CommandContext(ctx, "git", "clone", repoURL, cacheDir)
	if err := cmd.Run(); err != nil {
		return errors.Wrap(err, "failed to clone repository")
	}

	// If a tag is specified, checkout that tag
	if tag != "" {
		cmd = exec.CommandContext(ctx, "git", "checkout", tag)
		cmd.Dir = cacheDir
		if err := cmd.Run(); err != nil {
			return errors.Wrapf(err, "failed to checkout tag: %s", tag)
		}
	}

	return nil
}

// UpdateRepoWithCmd updates an existing git repository in the specified cache directory.
// // If a tag is specified, it checks out that tag; otherwise, it fetches the latest version.
func UpdateRepoWithCmd(ctx context.Context, cacheDir, tag string) error {
	// Pull the latest changes
	cmd := exec.CommandContext(ctx, "git", "pull", "origin")
	cmd.Dir = cacheDir
	if err := cmd.Run(); err != nil {
		return errors.Wrap(err, "failed to pull repository")
	}
	// If a tag is specified, checkout that tag
	if tag != "" {
		cmd = exec.CommandContext(ctx, "git", "checkout", tag)
		if err := cmd.Run(); err != nil {
			return errors.Wrapf(err, "failed to checkout tag: %s", tag)
		}
	}

	return nil
}

// CopyDir copies the contents of the source directory to the destination directory,
// excluding the .git directory.
func CopyDir(ctx context.Context, src string, dst string) error {
	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip the .git directory
		if info.IsDir() && (info.Name() == ".git" || info.Name() == ".github") {
			return filepath.SkipDir
		}

		// Create the destination path
		relPath, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}
		dstPath := filepath.Join(dst, relPath)

		if info.IsDir() {
			// Create the directory
			return os.MkdirAll(dstPath, os.ModePerm)
		}

		// Copy the file
		return copyFile(ctx, path, dstPath)
	})
}

// copyFile copies a file from src to dst.
func copyFile(ctx context.Context, src string, dst string) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		// Open the source file
	}

	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destination.Close()

	if _, err := io.Copy(destination, source); err != nil {
		return err
	}

	return nil
}
