package tldr

import (
	"bufio"
	"io"
	"strings"
)

// Output terms
const (
	Yellow  = "\x1b[33;1m"
	BLUE    = "\x1b[34;1m"
	Cyan    = "\x1b[36;1m"
	GREEN   = "\x1b[32;1m"
	Magenta = "\x1b[35;1m"
	RESET   = "\x1b[31;0m"
)

// Render takes the given input and renders it for a prettier output.
func Render(markdown io.Reader) (string, error) {
	var rendered string
	var renderingExample bool
	scanner := bufio.NewScanner(markdown)
	for scanner.Scan() {
		line := scanner.Text()
		if renderingExample {
			// Skip the empty line
			scanner.Scan()
			line = scanner.Text()

			line = strings.Replace(line, "{{", Yellow, -1)
			line = strings.Replace(line, "}}", Magenta, -1)
			rendered += "\t" + Magenta + strings.Trim(line, "`") + RESET + "\n"

			renderingExample = false
		} else if strings.HasPrefix(line, "#") {
			// Heading
			rendered += line[2:] + "\n"
		} else if strings.HasPrefix(line, ">") {
			// Quote
			rendered += line[2:] + "\n"
		} else if strings.HasPrefix(line, "-") {
			// Example
			rendered += Cyan + line + RESET + "\n"
			renderingExample = true
		} else {
			rendered += line + "\n"
		}
	}
	rendered += "\n"
	return rendered, scanner.Err()
}

// Write is a convenience function that calls Render and writes the output
// to the destination.
func Write(markdown io.Reader, dest io.Writer) error {
	out, err := Render(markdown)
	if err != nil {
		return err
	}
	_, err = io.WriteString(dest, out)

	return err
}
