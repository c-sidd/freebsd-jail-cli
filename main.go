package main

import (
	"fmt"
	"os"
	"os/exec"
)

// helper to run system commands
func runCommand(name string, args ...string) {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: bsdctl [create|start|stop|ps|exec]")
		return
	}

	switch os.Args[1] {

	case "create":
		if len(os.Args) < 3 {
			fmt.Println("Usage: bsdctl create <jailname>")
			return
		}

		jailName := os.Args[2]
		jailPath := "/usr/jails/" + jailName

		fmt.Println("Creating jail:", jailName)

		runCommand("mkdir", "-p", jailPath)
		runCommand("tar", "-xf", "base.txz", "-C", jailPath)

		config := fmt.Sprintf(`
%s {
    path = %s;
    mount.devfs;
    exec.start = "/bin/sh /etc/rc";
    exec.stop = "/bin/sh /etc/rc.shutdown";
    host.hostname = "%s";
    ip4.addr = 127.0.0.2;
}
`, jailName, jailPath, jailName)

		f, err := os.OpenFile("/etc/jail.conf", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println("Error opening jail.conf:", err)
			return
		}
		defer f.Close()

		_, err = f.WriteString(config)
		if err != nil {
			fmt.Println("Error writing config:", err)
			return
		}

		fmt.Println("Jail created successfully:", jailName)

	case "start":
		fmt.Println("Starting jails...")
		runCommand("service", "jail", "start")

	case "stop":
		fmt.Println("Stopping jails...")
		runCommand("service", "jail", "stop")

	case "ps":
		runCommand("jls")

	case "exec":
		if len(os.Args) < 3 {
			fmt.Println("Usage: bsdctl exec <jailname>")
			return
		}

		cmd := exec.Command("jexec", os.Args[2], "sh")
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err := cmd.Run()
		if err != nil {
			fmt.Println("Error:", err)
		}

	default:
		fmt.Println("Unknown command")
	}
}