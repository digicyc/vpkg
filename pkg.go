package main

import (
	"fmt"
	"log"
	"os/exec"
	"regexp"
	"strings"
)

/*
 * Search for a package passed into search field.
 */
func SearchPkg(pkgName string) map[string]string {
	// Key - pkgname Value - Description
	re := regexp.MustCompile("[ ]{3,}\\w+")
	cmdString := "/usr/bin/xbps-query"
	arg1 := "-Rs"
	cmd := exec.Command(cmdString, arg1, pkgName)

	pkgmap := map[string]string{}
	stdout, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	res := string(stdout)
	res_list := strings.Split(res, "[-] ")
	for i := range res_list {
		pkginfo := re.Split(res_list[i], -1)
		if len(pkginfo) > 1 {
			pkgmap[pkginfo[0]] = pkginfo[1]
		}
	}

	return pkgmap
}

/*
 * Install our selected package.
 */
func InstallPkg(pkgName string) string {
	// First check if we have sudo privileges
	cmdString := "/usr/bin/sudo"
	cmd := exec.Command(cmdString, "-n", "true")
	if err := cmd.Run(); err != nil {
		return "Error: This operation requires sudo privileges. Please run the application with sudo."
	}

	// Now run xbps-install with sudo
	cmd = exec.Command(cmdString, "/usr/bin/xbps-install", pkgName, "-y")
	stdout, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Sprintf("Error installing package: %v\n%s", err, string(stdout))
	}

	return string(stdout)
}

/*
 * Run this to update our package repos.
 */
func UpdateXBPS() string {
	cmdString := "/usr/bin/xbps-install"
	arg1 := "-S"
	cmd := exec.Command(cmdString, arg1)
	stdout, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	return string(stdout)
}
