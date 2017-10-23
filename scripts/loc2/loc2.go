package main

import (
	"bytes"
	"fmt"
	"github.com/hacdias/fileutils"
	"github.com/jozn/go-jalali/jalali"
	"io/ioutil"
	"ms/sun/helper"
	"os"
	"os/exec"
	"path"
	"time"
)

const SUN_DIR = `C:\Go\_gopath\src\ms\sun\`
const TMP_DIR = SUN_DIR + "scripts/tmp/"
const SCRIPTS_DIR = SUN_DIR + "scripts/"
const ANDROID_APP_DIR = `D:\dev_working2\MS_Native\app\src\main\java\com\mardomsara\`
const ANDROID_SOCIAL_DIR = `D:\dev_working2\MS_Native\app\src\main\java\com\mardomsara\social`
const ANTS_DIR = `C:\Go\_gopath\src\ms\ants`

func main() {
	bts := bytes.NewBufferString("")

	copypDirsForSun()

	bts.WriteString(header("Sun"))

	cmd := exec.Command("cloc.exe", SCRIPTS_DIR)
	//cmd.Stdout = os.Stdout
	cmd.Stdout = bts
	cmd.Run()

	bts.WriteString(header("End Sun"))
	cleanTemp()

	bts.WriteString("\n\n\n")
	///////////////// Android /////////////////////
	bts.WriteString(header("Android"))

	cmd = exec.Command("cloc.exe", ANDROID_APP_DIR)
	cmd.Stdout = bts
	cmd.Run()
	bts.WriteString(header("End Android"))
	bts.WriteString("\n\n\n")
	////////////////// End Android //////////////////

	///////////////// Android Social /////////////////////
	bts.WriteString(header("Android Social"))
	cmd = exec.Command("cloc.exe", ANDROID_SOCIAL_DIR)
	cmd.Stdout = bts
	cmd.Run()
	bts.WriteString(header("End Android Social"))
	bts.WriteString("\n\n\n")
	////////////////// End Android Social //////////////////

	///////////////// Android Social /////////////////////
	bts.WriteString(header("Ants"))
	cmd = exec.Command("cloc.exe", ANTS_DIR)
	cmd.Stdout = bts
	cmd.Run()
	bts.WriteString(header("End Ants"))
	bts.WriteString("\n\n\n")
	////////////////// End Android Social //////////////////

	fmt.Println(bts.String())

	year, mont, day := jalali.Gtoj(time.Now())
	f, err := os.Create(path.Join(SCRIPTS_DIR, fmt.Sprintf("/loc_log/sun_%d__%d-%d-%d.txt", calDayChangeFrom826(), year, mont, day)))
	helper.NoErr(err)
	f.WriteString(bts.String())
}

func calDayChangeFrom826() int {
	t, err := time.Parse(time.RFC822Z, "20 Aug 17 00:00 +0330")
	helper.NoErr(err)
	startDay := 826
	now := time.Now()
	diff := now.Sub(t)
	fmt.Println("Days diff:", diff.Seconds()/86400)
	dayShift := diff.Seconds() / 86400
	startDay += int(dayShift)

	return startDay
}

func copypDirsForSun() {
	dirs := []string{"models", "ctrl", "helper", "base"}

	os.RemoveAll(TMP_DIR)
	os.Mkdir(TMP_DIR, os.ModeDir)

	for _, dir := range dirs {
		fileutils.CopyDir(SUN_DIR+dir, TMP_DIR+dir)
	}

	//cleanSubDirs(TMP_DIR + "models")
	os.RemoveAll(TMP_DIR + "models/x")
}

func cleanTemp() {
	os.RemoveAll(TMP_DIR)
}

func cleanSubDirs(folder string) {

	fs, err := ioutil.ReadDir(folder)
	helper.NoErr(err)

	for _, f := range fs {
		if f.IsDir() {
			//os.Remove(path.Join(folder + "/" + f.Name()))
			fmt.Println(path.Join(folder + "/" + f.Name()))
		}
	}
}

func header(h string) string {
	s := "==================================="
	return "" + s + " " + h + " " + s + "\n"
}
