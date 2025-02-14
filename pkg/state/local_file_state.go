package state

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/otiai10/copy"
	"sigs.k8s.io/yaml"

	"github.com/h8r-dev/heighliner/pkg/state/app"
)

// LocalFileState State using local file as backend
type LocalFileState struct {
}

// LoadOutput load output
func (l *LocalFileState) LoadOutput(appName string) (*app.Output, error) {
	b, err := os.ReadFile(filepath.Join(".hln", "output.yaml"))
	if err != nil {
		return nil, err
	}
	output := &app.Output{}
	err = yaml.Unmarshal(b, output)
	return output, err
}

// LoadTFProvider No need in Local File State
func (l *LocalFileState) LoadTFProvider(appName string) (string, error) {
	return "", nil
}

// ListApps only list app in current dir
func (l *LocalFileState) ListApps() ([]string, error) {
	op, err := l.LoadOutput("")
	if err != nil {
		return nil, err
	}
	return []string{op.ApplicationRef.Name}, nil
}

// SaveOutputAndTFProvider save output and tf provider
func (l *LocalFileState) SaveOutputAndTFProvider(appName string) error {
	pwd, err := os.Getwd()
	if err != nil {
		return err
	}
	if err := copy.Copy(stackOutput, filepath.Join(pwd, appInfo)); err != nil {
		return err
	}
	if err := os.Remove(stackOutput); err != nil {
		return err
	}
	ao, err := app.Load(filepath.Join(pwd, appInfo))
	if err != nil {
		return fmt.Errorf("failed to load app output: %w", err)
	}
	return copy.Copy(ao.SCM.TfProvider, filepath.Join(pwd, providerInfo))
}

// DeleteOutputAndTFProvider delete state file
func (l *LocalFileState) DeleteOutputAndTFProvider(appName string) error {
	pwd, err := os.Getwd()
	if err != nil {
		return err
	}
	if err := os.Remove(filepath.Join(pwd, ".hln", "output.yaml")); err != nil {
		return err
	}
	if err := os.Remove(filepath.Join(pwd, ".hln", "provider.tf")); err != nil {
		return err
	}
	return nil
}
