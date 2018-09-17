package impl

import (
	"errors"
	"fmt"
	"gopkg.in/src-d/go-git.v4"
	"io/ioutil"
	"os"
	"os/exec"
)

func Clone(url string) (string, error) {
	fmt.Println("Cloning %s...", url)
	clone, err := ioutil.TempDir("", "hugohook.inp.")
	if err != nil {
		return "", err
	}
	_, err = git.PlainClone(clone, false, &git.CloneOptions{
		URL:               url,
		Progress:          os.Stdout,
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
	})
	return clone, err
}

func RunHugo(clone string) (string, error) {
	fmt.Println("Running hugo on %s...", clone)
	hugo := os.Getenv("HUGO_BIN")

	generated, err := ioutil.TempDir("", "hugohook.out.")
	if err != nil {
		return "", err
	}
	fmt.Println("writing to %s", generated)

	cmd := exec.Command(hugo, "-s", clone, "-d", generated)
	output, err := cmd.CombinedOutput()
	if err != nil {
		err = errors.New("hugo: " + string(output[:]))
		fmt.Println(err)
	}

	return generated, err
}

func Upload(generated string) error {
	fmt.Println("Uploading %s...", generated)
	gsutil := os.Getenv("GSUTIL_BIN")

	cmd := exec.Command(gsutil, "-m", "rsync", "-R", "-d", "-a", "public-read", generated, "gs://www.bpcreech.com")
	output, err := cmd.CombinedOutput()
	if err != nil {
		err = errors.New("gsutil: " + string(output[:]))
		fmt.Println(err)
	}

	return err
}

func CloneAndRun(url string) error {
	clone, err := Clone(url)
	if clone != "" {
		defer os.RemoveAll(clone)
	}
	if err != nil {
		return err
	}

	generated, err := RunHugo(clone)
	if generated != "" {
		defer os.RemoveAll(generated)
	}
	if err != nil {
		return err
	}

	err = Upload(generated)

	return err
}
