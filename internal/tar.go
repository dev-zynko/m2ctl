package internal

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/schollz/progressbar/v3"
)

func Tar(source, target string) error {
	filename := filepath.Base(source)
	target = filepath.Join(target, fmt.Sprintf("%s.tar", filename))
	strings.ReplaceAll(target, `\`, "/")
	tarfile, err := os.Create(target)
	if err != nil {
		return err
	}
	defer tarfile.Close()

	tarball := tar.NewWriter(tarfile)
	defer tarball.Close()

	info, err := os.Stat(source)
	if err != nil {
		return nil
	}

	var baseDir string
	if info.IsDir() {
		baseDir = filepath.Base(source)
	}

	strings.ReplaceAll(baseDir, `\`, "/")

	return filepath.Walk(source,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			header, err := tar.FileInfoHeader(info, info.Name())
			if err != nil {
				return err
			}

			if baseDir != "" {
				header.Name = filepath.Join(baseDir, strings.TrimPrefix(path, source))
				header.Name = strings.ReplaceAll(header.Name, `\`, "/")

			}

			if err := tarball.WriteHeader(header); err != nil {
				return err
			}

			if info.IsDir() {
				return nil
			}

			file, err := os.Open(path)
			if err != nil {
				return err
			}

			defer file.Close()
			_, err = io.Copy(tarball, file)
			return err

		})

}

func Gzip(file string, target string) {
	reader, err := os.Open(file)
	if err != nil {
		log.Fatal("Failed to open file", err)
	}

	filename := filepath.Base(file)
	target = filepath.Join(target, fmt.Sprintf("%s.gz", filename))
	strings.ReplaceAll(target, `\`, "")
	writer, err := os.Create(target)
	if err != nil {
		log.Fatal("Failed creating file", err)
	}
	defer writer.Close()

	archiver := gzip.NewWriter(writer)
	archiver.Name = filename
	defer archiver.Close()

	fi, _ := reader.Stat()
	bar := progressbar.DefaultBytes(
		fi.Size(),
		"Giziping tarball",
	)
	_, err = io.Copy(io.MultiWriter(archiver, bar), reader)
	if err != nil {
		log.Fatal("Failed to gizip", err)
	}
}
