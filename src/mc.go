package main

type MagikCraft struct {
	x int
}

func (m MagikCraft) shakti() string {
	return "You cast lightning"
}

func main(arg int) {
	lightning(arg)
}

func lightning(num int) {
	magik := MagikCraft{}
	for i := 0; i < num; i++ {
		spell := magik.shakti()
		println(spell)
	}
}
