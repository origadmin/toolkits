package repo

import (
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/origadmin/toolkits/context"
	"github.com/origadmin/toolkits/errors"
	"golang.org/x/mod/module"
)

type Repository struct {
	useCmd   bool
	repoURL  string
	cacheDir string
	tag      string
}

func (r Repository) LocalPath() string {
	if r.repoURL == "" {
		return ""
	}
	targetDir, err := fromRepoURL(r.repoURL)
	if err != nil {
		return ""
	}

	return filepath.Join(r.cacheDir, targetDir)
}

func (r Repository) Copy(ctx context.Context, target string) error {
	return CopyDir(ctx, r.LocalPath(), target)
}

func (r Repository) BeforeCopy(ctx context.Context) error {
	// Create cache directory if it doesn't exist
	if err := os.MkdirAll(r.cacheDir, os.ModePerm); err != nil {
		return errors.Wrap(err, "failed to create cache directory")
	}

	localPath := r.LocalPath()
	if localPath == "" {
		return errors.New("failed to get local path")
	}
	stat, err := os.Stat(localPath)
	if err == nil && !stat.IsDir() {
		return fmt.Errorf("%s is not a directory", localPath)
	}

	if r.useCmd {
		if os.IsNotExist(err) {
			err := CloneRepoWithCmd(ctx, r.repoURL, localPath, r.tag)
			if err != nil {
				return err
			}
			return nil
		}
		return UpdateRepoWithCmd(ctx, localPath, r.tag)
	} else {
		if os.IsNotExist(err) {
			err := CloneRepo(ctx, r.repoURL, localPath, r.tag)
			if err != nil {
				return err
			}
			return nil
		}
		return UpdateRepo(ctx, localPath, r.tag)
	}
	return nil
}

func fromRepoURL(remoteURL string) (string, error) {
	// Parse the URL
	parsedURL, err := url.Parse(remoteURL)
	if err != nil {
		return "", fmt.Errorf("invalid URL: %w", err)
	}

	repoPath := parsedURL.Host + parsedURL.Path
	// Extract the path and remove the leading slash
	repoPath = strings.TrimPrefix(repoPath, "/")
	// trim <repository> with suffix .git
	repoPath = strings.TrimSuffix(repoPath, ".git")

	esc, err := module.EscapePath(repoPath)
	if err != nil {
		return "", fmt.Errorf("invalid Escape Path %s: %w", repoPath, err)
	}

	return esc, nil
}
