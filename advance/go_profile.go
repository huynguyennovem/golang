package main

import "github.com/pkg/profile"

/**
	Deps:
		- brew install graphviz

	Step 1: Run program. Output is a file .pprof
	Step 2: Open terminal:
			go tool pprof --png  /usr/local/go/bin/go <file .pprof path> > <file_name.png>
	Note: support export file format (canon cmap cmapx cmapx_np dot dot_json eps fig gd gd2 gif gv imap imap_np ismap jpe jpeg jpg json json0 mp pic plain plain-ext png pov ps ps2 svg svgz tk vml vmlz vrml wbmp xdot xdot1.2 xdot1.4 xdot_json)
 */

func main() {
	defer profile.Start(profile.MemProfile).Stop()
}
