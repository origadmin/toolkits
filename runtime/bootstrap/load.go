package bootstrap

import (
	"os"
	"path/filepath"

	"github.com/origadmin/toolkits/codec"
	configv1 "github.com/origadmin/toolkits/runtime/gen/go/config/v1"

	"github.com/origadmin/toolkits/errors"
)

// loadSource loads the config file from the given path
func loadSource(si os.FileInfo, path string) (*configv1.SourceConfig, error) {
	// Check if the file or directory exists
	if si == nil {
		return nil, errors.New("load config file target is not exist")
	}
	// Check if the path is a directory
	if si.IsDir() {
		return loadSourceDir(path)
	}
	// Get the file type from the extension
	typo := codec.TypeFromExt(filepath.Ext(path))
	// Check if the file type is unknown
	if typo == codec.UNKNOWN {
		return nil, errors.New("unknown file type: " + path)
	}
	// Load the config file
	return loadSourceFile(path)
}

// loadSourceFile loads the config file from the given path
func loadSourceFile(path string) (*configv1.SourceConfig, error) {
	var cfg configv1.SourceConfig
	// Decode the file into the config struct
	if err := codec.DecodeFromFile(path, &cfg); err != nil {
		return nil, errors.Wrapf(err, "failed to parse config file %s", path)
	}
	return &cfg, nil
}

// loadSourceDir loads the config file from the given directory
func loadSourceDir(path string) (*configv1.SourceConfig, error) {
	var cfg configv1.SourceConfig
	// Walk through the directory and load each file
	err := filepath.WalkDir(path, func(walkpath string, d os.DirEntry, err error) error {
		if err != nil {
			return errors.Wrapf(err, "failed to get config file %s", walkpath)
		}
		// Check if the path is a directory
		if d.IsDir() {
			return nil
		}
		// Get the file type from the extension
		typo := codec.TypeFromExt(filepath.Ext(walkpath))
		// Check if the file type is unknown
		if typo == codec.UNKNOWN {
			return nil
		}

		// Decode the file into the config struct
		if err := codec.DecodeFromFile(walkpath, &cfg); err != nil {
			return errors.Wrapf(err, "failed to parse config file %s", walkpath)
		}
		return nil
	})
	if err != nil {
		return nil, errors.Wrap(err, "load config error")
	}
	return &cfg, nil
}

// LoadSourceConfig loads the config file from the given path
func LoadSourceConfig(bootstrap *Bootstrap) (*configv1.SourceConfig, error) {
	// Get the path from the bootstrap
	path := bootstrap.WorkPath()

	// Get the file info from the path
	stat, err := os.Stat(path)
	if err != nil {
		return nil, errors.Wrap(err, "load config stat error")
	}
	// Load the config file
	return loadSource(stat, path)
}
