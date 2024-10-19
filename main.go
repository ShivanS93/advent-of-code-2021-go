package main

import (
	"adventofcode2021/day02"
	//"adventofcode2021/day03"
	"adventofcode2021/day04"
	"adventofcode2021/day06"
	"adventofcode2021/day07"
	"adventofcode2021/day15"

	//"adventofcode2021/day08"
	"adventofcode2021/day10"
	"fmt"
)

func main_day02() {
	// day 02
	// part 1
	sample_file := "./day02/sample_input_01.txt"
	file := "./day02/input.txt"
	sample_ans_day02_part1, _ := day02.ProductDepthHorizontalPosition(sample_file)
	ans_day02_part1, _ := day02.ProductDepthHorizontalPosition(file)
	fmt.Println("Day 02 - Part 1:")
	fmt.Printf("Result from sample input: %d\n", sample_ans_day02_part1)
	fmt.Printf("Result from real input: %d\n", ans_day02_part1)

	// part 2
	sample_ans_day02_part2, _ := day02.ProductDepthAimHorizontalPosition(sample_file)
	ans_day02_part2, _ := day02.ProductDepthAimHorizontalPosition(file)
	fmt.Println("Day 02 - Part 2:")
	fmt.Printf("Result from sample input: %d\n", sample_ans_day02_part2)
	fmt.Printf("Result from real input: %d\n", ans_day02_part2)
}

// func main_day03() {
// 	// day03
// 	// part 1
// 	sample_file_day03 := "./day03/sample_input.txt"
// 	file_day03 := "./day03/input.txt"
// 	sample_ans_day03_part1 := day03.CalculatePowerConsumption(sample_file_day03)
// 	ans_day03_part1 := day03.CalculatePowerConsumption(file_day03)
// 	fmt.Println("Day 03 - Part 1:")
// 	fmt.Printf("Result from sample input: %d\n", sample_ans_day03_part1)
// 	fmt.Printf("Result from real input: %d\n", ans_day03_part1)

// 	// part 2
// 	sample_ans_day03_part2 := day03.CalculateOxygenCO2(sample_file_day03)
// 	ans_day03_part2 := day03.CalculateOxygenCO2(file_day03)
// 	fmt.Println("Day 03 - Part 2:")
// 	fmt.Printf("Result from sample input: %d\n", sample_ans_day03_part2)
// 	fmt.Printf("Result from real input: %d\n", ans_day03_part2)
// }

func main_day04() {
	// day04
	// part 1
	sample_file_day04_draw := "./day04/sample_input_draw.txt"
	sample_file_day04 := "./day04/sample_input.txt"
	file_day04_draw := "./day04/input_draw.txt"
	file_day04 := "./day04/input.txt"

	sample_ans := day04.WinnerBingo(sample_file_day04, sample_file_day04_draw)
	ans := day04.WinnerBingo(file_day04, file_day04_draw)

	fmt.Println("Day 04 - Part 1:")
	fmt.Printf("Result from sample input: %v\n", sample_ans)
	fmt.Printf("Result from input: %v\n", ans)

	// part 2
}

// func main_day05() {
// 	// day05
// 	sample_file := "sample_input.txt"
// 	file := "input.txt"
// 	// part 1
// 	// TODO: create functions
// }

func main_day06() {
	// day06
	sample_file := "./day06/sample_input.txt"
	file := "./day06/input.txt"
	// part 1
	sample_ans := day06.SimulateFish(sample_file, 80)
	ans := day06.SimulateFish(file, 80)
	fmt.Println("Day 06 - Part 1:")
	fmt.Printf("Result from sample input: %v\n", sample_ans)
	fmt.Printf("Result from input: %v\n", ans)

	// part 2
	p2_sample_ans := day06.SimulateFish(sample_file, 256)
	p2_ans := day06.SimulateFish(file, 256)
	fmt.Println("Day 06 - Part 2:")
	fmt.Printf("Result from sample input: %v\n", p2_sample_ans)
	fmt.Printf("Result from input: %v\n", p2_ans)
}

func main_day07() {
	// day07
	sample_file := "./day07/sample_input.txt"
	file := "./day07/input.txt"
	// part 1
	sample_ans := day07.FindMoves(sample_file)
	ans := day07.FindMoves(file)
	fmt.Println("Day 07 - Part 1:")
	fmt.Printf("Result from sample input: %v\n", sample_ans)
	fmt.Printf("Result from input: %v\n", ans)

	// part 2
	p2_sample_ans := day07.FindMovesExpensive(sample_file)
	p2_ans := day07.FindMovesExpensive(file)
	fmt.Println("Day 07 - Part 2:")
	fmt.Printf("Result from sample input: %v\n", p2_sample_ans)
	fmt.Printf("Result from input: %v\n", p2_ans)
}

// func main_day08() {
// 	// day08
// 	sample_file := "./day08/sample_input.txt"
// 	file := "./day08/input.txt"
// 	// part 1
// 	sample_ans := day08.DigitReader(sample_file)
// 	ans := day08.DigitReader(file)
// 	fmt.Println("Day 08 - Part 1:")
// 	fmt.Printf("Result from sample input: %v\n", sample_ans)
// 	fmt.Printf("Result from input: %v\n", ans)

// 	// part 2
// 	// p2_sample_ans := day07.FindMovesExpensive(sample_file)
// 	// p2_ans := day07.FindMovesExpensive(file)
// 	// fmt.Println("Day 07 - Part 2:")
// 	// fmt.Printf("Result from sample input: %v\n", p2_sample_ans)
// 	// fmt.Printf("Result from input: %v\n", p2_ans)
// }

// func main_day09() {
// 	// day09
// 	sample_file := "./day08/sample_input.txt"
// 	file := "./day08/input.txt"
// 	// part 1
// 	sample_ans := day09.DigitReader(sample_file)
// 	ans := day09.DigitReader(file)
// 	fmt.Println("Day 08 - Part 1:")
// 	fmt.Printf("Result from sample input: %v\n", sample_ans)
// 	fmt.Printf("Result from input: %v\n", ans)

// 	// part 2
// 	// p2_sample_ans := day07.FindMovesExpensive(sample_file)
// 	// p2_ans := day07.FindMovesExpensive(file)
// 	// fmt.Println("Day 07 - Part 2:")
// 	// fmt.Printf("Result from sample input: %v\n", p2_sample_ans)
// 	// fmt.Printf("Result from input: %v\n", p2_ans)
// }

func main_day10() {
	// day10
	sample_file := "./day10/sample_input.txt"
	file := "./day10/input.txt"
	// part 1
	sample_ans := day10.FindCorrupt(sample_file)
	ans := day10.FindCorrupt(file)
	fmt.Println("Day 10 - Part 1:")
	fmt.Printf("Result from sample input: %v\n", sample_ans)
	fmt.Printf("Result from input: %v\n", ans)
}

func main_day15() {
	// day 15
	sample_file := "./day15/sample_input.txt"
	file := "./day15/input.txt"
	// part 1
	sample_ans, _ := day15.PathFinder(sample_file)
	ans, _ := day15.PathFinder(file)
	fmt.Println("Day 15 - Part 1:")
	fmt.Printf("Result from sample input: %v\n", sample_ans)
	fmt.Printf("Result from input: %v\n", ans)
	// part 2
}

func main() {
	main_day15()
}
