package main

import (
	"log"
	"os/exec"
	"regexp"
	"strings"
)

func SearchPkg(pkgName string) map[string]string {
    // Pass in package name to search for
    // Pull proper name and description.
    re := regexp.MustCompile("[ ]{2,}\\w")
    cmdString := "xbps-query -Rs " + pkgName
    cmd := exec.Command(cmdString)
    pkgmap := map[string]string{}
    stdout, err := cmd.Output()
    if err != nil {
        log.Fatal(err)
    }

    res := string(stdout)
    res_list := strings.Split(res, "[-]")
    for i := range res_list {
        pkginfo := re.Split(res_list[i], -1)
        if len(pkginfo) > 1 {
            pkgmap[pkginfo[0]] = pkginfo[1]
        }
    }

   return pkgmap
}

func InstallPkg(pkgName string) string {
    // Update the repos
    // xbps-install -Su
    return "a pkg yay!"
}
