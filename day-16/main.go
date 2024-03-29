package main

import (
	"fmt"
	"strings"
)

func Parse(str string) Grid {
	str = strings.TrimSpace(str)
	var grid Grid
	for _, line := range strings.Split(str, "\n") {
		var row []int32
		for _, char := range line {
			row = append(row, char)
		}
		grid.layout = append(grid.layout, row)
	}
	grid.beamPoints = make(map[Beam]struct{})
	return grid
}

func main() {
	grid := Parse(PUZZLE)

	fmt.Println("part 1", grid.GetCoverageFrom(Beam{
		position: Position{0, 0},
		facing:   1,
	}))
	grid.beamPoints = make(map[Beam]struct{})

	bestCoverage := 0
	maxY := len(grid.layout) - 1
	maxX := len(grid.layout[0]) - 1
	for i := 0; i <= maxY; i++ {
		e := grid.GetCoverageFrom(Beam{
			position: Position{i, 0},
			facing:   1,
		})
		grid.beamPoints = make(map[Beam]struct{})
		s := grid.GetCoverageFrom(Beam{
			position: Position{0, i},
			facing:   2,
		})
		grid.beamPoints = make(map[Beam]struct{})
		w := grid.GetCoverageFrom(Beam{
			position: Position{i, maxX},
			facing:   3,
		})
		grid.beamPoints = make(map[Beam]struct{})
		n := grid.GetCoverageFrom(Beam{
			position: Position{maxY, i},
			facing:   0,
		})
		grid.beamPoints = make(map[Beam]struct{})

		bestCoverage = max(bestCoverage, e, s, w, n)
	}
	fmt.Println("part 2", bestCoverage)
}

func max(n ...int) (m int) {
	for _, i := range n {
		if i > m {
			m = i
		}
	}
	return
}

const TEST = `
.|...\....
|.-.\.....
.....|-...
........|.
..........
.........\
..../.\\..
.-.-/..|..
.|....-|.\
..//.|....`

const PUZZLE = `
\.....................\.......................|..../......-.........../.|.....\.............|..........|..\...
.../...................\..................../...|.......|........||...................../...........\.........
././....|......./................\............|..\...........-....../............../|........-................
....\..\...........-/..|......|......\............|............../.................-........./...-...|.......\
..................................|.\.......|...../.........................-...../......\..|...............\.
.....\.............................................|../.\............../--.......................\............
|.......-.|........\.................\..............|........./....\.........|...................\-.-...|.....
....\........./....-.........................--.....\..../......................|.\......................\...-
\..-..../.....|..........\......................-|......|...................../...../........./...........-...
........................\.\|......../........../......................\.............|./../....................
.....\...................|.............................|......-........|.........-./.../.|....................
..-................./.........................................\..\........-..-........-.......\../............
-................-............................\................./..--................-............../.|.......
..-.....|.............|................./.......-.-........................-............-...........\.........
..........|..........................-...|..\...............................................\-....\...........
/................../..........|....\..-......................../....\...\..........-....../.........../.....\.
................\....-................-|.\...-...........-.....-........\..........................\..........
...................................\../..|.../../...........................\.........../.\-......|.........-.
/....|....|..........................\............./.....|/................./....-......................-.....
....-.............../.....|...........|../......\../.............../........................|............./...
......./..............|./../.................-.\.........../.....................|.........../........\.......
...................-...............\/.......|........|....................../.......................-/......\.
................\...../....-/......-.........-.....\/.\...............\....../...-....|......|.........\......
....-.........-.........................-......................../.......|...............................\....
......../..................|...../......\..\..-|-..../..............|......../............|.-.......-.........
......./............-..........-............................../.........................|\..-..-\.\......-..\.
.\-.....|......................-..\................/.-................|.\..............\..................|...
-..........................-............../.|.-......\......-..-...../......./................................
....|.................................................\....|..........\...................-......|........./..
................../...........-..............\......................................|.-........|..............
...-............|..............-...........-.\.......................|..|..../......................../.|....-
.......-..-......./...../.........-.......\.....-.../............-./......../.../............-.....|...|......
...............|.........../..........|/..........|..........-.-...../...../..|..........\.....-./.|./....|...
.........|......................\.|.....|......................-.-........|.............................-...|.
.-...|..\...-...............\..............\......./........-............/.-..\........../...../....-....\....
...........\.-.......|.......|....-.............|-.|........../..../...\....\............|.|..|\..............
.....................|.\....|......|../...\......|........................\......|.......-./..../.............
......................\--................................\.................|................-.................
-...............\...-.....................................-.\...../....\...........................|..........
.....|.......\.......................-/...\..-.........-..../\|..................\........-...................
..|........./....|..-./..\..\................-|..........\../......\...............\............\............|
......-................../.............|.............\................................-..\...../........-.....
....................../......../..\.-.........................|....\..|...\...../../.....|............../.....
...\................................/..............................\........|.........\.-............\........
/.................../................-...........-.|.............\.................................|.....-....
.|...-..|/./../.........-./.............................................../..........-.........-.|./..........
..../..................|......................|../..........\........-......../.........\...|.................
........./..................|..-..........|\............\......-..-..\...|....|.....................|....-....
.....|/....|........\.........-................../..........|............|...........-.....|......./......//..
..........\.|../...|........................\.........\.|......||....-...............\\.....|.|..............|
............-................|.....-....||......../.../.......................|..\...........\................
...............-..../......-..../.......|-.....-.....\..................................../..|................
/......-...............\/...././..................\..................\..........-......................|......
...................\.....\.................\...|.................-............/........-..............\./.....
.-.\....\........\........-../.-............../...../..-..|...................../.-.....\.....-........../..\.
|...|...............................|............../.............../....................|...../...............
....../....../...|..............-.-........-..|.../....................\|.....\......................|....../.
...\\.......|...|-............|..........-...........................................\/...../..-.-............
.........|.............|..|......\............\...............|......\......\.............\...\-.......\.....\
...........\..-.|.......|\..\..................-........\/...............................-....................
...\.................\..............................................\...............\.......-.../.............
..\.............\....---....................//.......................................|./.............../\.....
.\...........|..................................\........-......|.........\.................-........-........
..............-...............-.../....-........|.....-../......................./..../........../../--..\....
.........-......\......-.........../...\./.|.................\...................--..-........|..\......\....|
....../.........................................|...../........../..............\\...-/.......-......\..\.....
.........../...\.|..................../.\.................|......\.\..|.........../...../.....\.\...........|.
..../\..-...../.............................../..../............................-...-......\...........\......
..-..\.............................\...........|..........|.........|..........................-......./......
-..|.................-\.....\...............\/.....|...........................././...\......\.......\...-....
......\..............|..|..............-............../....................-....\\........\....|..............
|.............\....\./.....-|-....\.....\..............\/........-.|..|...........-..\......-..../.......|....
............................-.......................|........................../................|.............
/................|../............\/.....\/....-.....\......\...../|..\..........|............../..............
.|....../...../....................../..|.....................-................./...../-.........-............
......................\|........../......./..\........\.........-......../.........../.|/.......\.../.........
...............\......-|............\..\....../.........\.....|...|\............/............\.\..|.......\...
......|...........\....../....\./......./-....\../.......\...-...............\./...\.....-.......-..../.......
...|..|.............................\.....-.|........./......../......................-......./............\..
.....|........\.........|...-.-.........\..../....../..-.......-...............\..../.....-..../..............
........./.\/........................-..........\.....................|.......-..............\./.\............
..../................/................................|.....||..............\........-...............|........
....\./.........../\-.....-...\............//.\..............\...../................|.....\..-................
..../|........./............|.....................|......................-..|...............................\.
..................\............................\.......\.........-.............|...../.-.............-......|.
..............|.-...........................\.\./.......\........../...\................/....-................
.......|................/-............\............/...........-..../...........\......../...\.......\........
..............|..../......................-...........-..\.................................................-..
..............................-..-................-.......-.................-..........\.../...............|..
......-.........-........................./.....................\....|........................................
.....\........./.....-.................../................|...................|........................|......
........../.../........................./.|................-...........\...........-....|.....................
.....-.......\././/.......|..............|......../..\...........\.............................-......|../-.-.
........|-....-....................\.....|\...\/.....\........\......./..........|.\-..|......................
../...-...-...|...../............\.......................-..........\...\.....-...\...............|....../...|
............|..\....|..................\/...................../../....|...................|.............../...
.............-................./............|.......|............\....................................\.-.....
-...../|..............-......\..../..................|..\....-............................................\...
........|.........-........\................/....\..................../.|...................................|.
.....|................./.........-|\...-...................|....-................-......../...............|..-
.|.........|...........|...-.............................\............................|................\......
....|.....-......|......./.-.|...........\..-........|.........................-....-..-.....|................
........................................\.....................-.../........................../../..........|..
-................/....\.......//....-........................../.\.........-......./..|\..........\...........
-............................./......................................../-...........................\|........
........................................../-.....//......\......\...........\........|./....../.......--......
../..............-....../......../................../......./....-...........|../.......\.\.......-......./..\
..............-............../......|.........-...............|................................-..........\...
.....................\..........|..\......|...............................................-...................
...../..|.../.........|.../.......-...................................\.....................................|.`
