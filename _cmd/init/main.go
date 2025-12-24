package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const (
	oldModule = "github.com/Aninetix/aninetix-app-template"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage:")
		fmt.Println("  go run ./cmd/init <new-module-name>")
		fmt.Println("Example:")
		fmt.Println("  go run ./cmd/init github.com/Aninetix/my-awesome-app")
		os.Exit(1)
	}

	newModule := strings.TrimSpace(os.Args[1])

	if newModule == "" || newModule == oldModule {
		fmt.Println("❌ Invalid module name")
		os.Exit(1)
	}

	fmt.Println("🚀 Initializing Aninetix app")
	fmt.Println("FROM:", oldModule)
	fmt.Println("TO  :", newModule)

	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// ⛔ dossiers à ignorer
		if info.IsDir() {
			switch info.Name() {
			case ".git", "vendor":
				return filepath.SkipDir
			}
			return nil
		}

		// seuls les fichiers Go + go.mod
		if !strings.HasSuffix(path, ".go") && filepath.Base(path) != "go.mod" {
			return nil
		}

		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		if !bytes.Contains(data, []byte(oldModule)) {
			return nil
		}

		updated := bytes.ReplaceAll(data, []byte(oldModule), []byte(newModule))

		if err := os.WriteFile(path, updated, info.Mode()); err != nil {
			return err
		}

		fmt.Println("✔ updated:", path)
		return nil
	})

	if err != nil {
		fmt.Println("❌ Error:", err)
		os.Exit(1)
	}

	_ = os.RemoveAll("cmd/init")

	fmt.Println("\n✅ Project initialized successfully")
	fmt.Println("Next steps:")
	fmt.Println("  rm -rf .git")
	fmt.Println("  git init")
	fmt.Println("  go mod tidy")
}
