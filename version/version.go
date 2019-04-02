
var goVersionTestOutput = ""

func getGoVersion() (string, error) {
	// For testing purposes only
	if goVersionTestOutput != "" {
		return goVersionTestOutput, nil
	}

	// Godep might have been compiled with a different
	// version, so we can't just use runtime.Version here.
	cmd := exec.Command("go", "version")
	cmd.Stderr = os.Stderr
	out, err := cmd.Output()
	return string(out), err
}

// goVersion returns the major version string of the Go compiler
// currently installed, e.g. "go1.5".
func goVersion() (string, error) {
	out, err := getGoVersion()
	if err != nil {
		return "", err
	}
	gv := strings.Split(out, " ")
	if len(gv) < 4 {
		return "", fmt.Errorf("Error splitting output of `go version`: Expected 4 or more elements, but there are < 4: %q", out)
	}
	if gv[2] == "devel" {
		return trimGoVersion(gv[2] + gv[3])
	}
	return trimGoVersion(gv[2])
}
