package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: bsdctl [ps]")
		return
	}

	if os.Args[1] == "ps" {
		cmd := exec.Command("jls")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()
	}
	config := fmt.Sprintf(`
%s {
    path = %s;
    mount.devfs;
    exec.start = "/bin/sh /etc/rc";
    exex.stop = "/bin/sh /etc/rc.shutdown";
    host.hostname = "%s";
    ip4.addr = 127.0.0.2;
}
`, jailName, jailPath, jailName)
}