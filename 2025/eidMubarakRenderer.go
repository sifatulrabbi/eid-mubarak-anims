package main

func eidMubarakRenderer(screen *Screen) {
	y := 0
	lastX := 0
	for x := 0; x < len(text); x++ {
		v := string(text[x])
		switch v {
		case "\n":
			y++
			lastX = x + 1 // +1 because the next x value should be 0 to indicate the starting index.
			break

		case " ":
			break

		default:
			screen.PaintFg(x-lastX, y, v)
			break
		}
	}
}

var text = `

@@@@@@  @@@@@@  @@@@@        @@@    @@@  @@   @@  @@@@@@   @@@@@@@  @@@@@@   @@@@@@@  @@  @@
@@        @@    @@  @@       @@@@  @@@@  @@   @@  @@  @@@  @@   @@  @@  @@@  @@   @@  @@ @@
@@@@@@    @@    @@   @@      @@ @@@@ @@  @@   @@  @@@@@@   @@@@@@@  @@@@@@   @@@@@@@  @@@@
@@        @@    @@  @@       @@  @@  @@  @@   @@  @@  @@@  @@   @@  @@  @@   @@   @@  @@ @@
@@@@@@  @@@@@@  @@@@@        @@      @@  @@@@@@@  @@@@@@   @@   @@  @@   @@  @@   @@  @@  @@

`
