package day5

import (
	"github.com/grambbledook/adventofcode2023/lang"
	"github.com/grambbledook/adventofcode2023/util"
	"slices"
	"strings"
	"sync"
)

type Range struct {
	Start int
	End   int
}

func (r *Range) Contains(i int) bool {
	return r.Start <= i && i <= r.End
}

type Edge struct {
	Src  Range
	Dest Range
}

type Data struct {
	Seeds      []int
	SeedRanges []Range

	SeedSoil   []Edge
	SoilFert   []Edge
	FertWater  []Edge
	WaterLight []Edge
	LightTemp  []Edge
	TempHumid  []Edge
	HumidLoc   []Edge
}

func Task1(data []string) int {
	graph := buildGraph(data)
	locations := lang.Map(graph.Seeds, func(u int) int { return traverse(u, graph) })

	return slices.Min(locations)
}

func Task2(data []string) int {
	graph := buildGraph(data)

	var wg sync.WaitGroup
	ch := make(chan int, len(graph.SeedRanges))

	for _, seeds := range graph.SeedRanges {
		wg.Add(1)
		go func(seeds Range) {
			ch <- traverseRangeParallel(seeds, graph)
			wg.Done()
		}(seeds)
	}

	wg.Wait()
	close(ch)

	return util.Min(ch)
}

func buildGraph(str []string) Data {
	seeds := lang.Map(strings.Fields(str[0])[1:], lang.Int)
	seedRanges := lang.Map(lang.Chunked(seeds, 2), func(r []int) Range { return Range{r[0], r[0] + r[1]} })

	seedSoil, soil := populateMap(str, 3)
	soilFert, fert := populateMap(str, soil)
	fertWater, water := populateMap(str, fert)
	waterLight, light := populateMap(str, water)
	lightTemp, temp := populateMap(str, light)
	tempHumid, humid := populateMap(str, temp)
	humidLoc, _ := populateMap(str, humid)

	return Data{
		Seeds:      seeds,
		SeedRanges: seedRanges,
		SeedSoil:   seedSoil,
		SoilFert:   soilFert,
		FertWater:  fertWater,
		WaterLight: waterLight,
		LightTemp:  lightTemp,
		TempHumid:  tempHumid,
		HumidLoc:   humidLoc,
	}
}

func populateMap(str []string, idx int) ([]Edge, int) {
	var collection []Edge

	for idx < len(str) {
		if str[idx] == "" {
			break
		}

		fields := lang.Fields(str[idx])

		size := lang.Int(fields(2, false))
		src := lang.Int(fields(1, false))
		dest := lang.Int(fields(0, false))

		collection = append(collection, Edge{
			Dest: Range{dest, dest + size},
			Src:  Range{src, src + size},
		})

		idx++
	}

	return collection, idx + 2
}

func traverseRangeParallel(seeds Range, graph Data) int {
	ch := make(chan int, seeds.End-seeds.Start+1)
	var wg sync.WaitGroup

	step := max(1, (seeds.End-seeds.Start)/10000)

	for start := seeds.Start; start <= seeds.End; start += step {
		wg.Add(1)
		go traverseRange(start, min(seeds.End, start+step), graph, ch, &wg)
	}

	wg.Wait()
	close(ch)

	return util.Min(ch)
}

func traverseRange(start, end int, graph Data, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for u := start; u < end; u++ {
		ch <- traverse(u, graph)
	}
}

func traverse(start int, graph Data) int {
	soil := get(graph.SeedSoil, start)

	fert := get(graph.SoilFert, soil)
	water := get(graph.FertWater, fert)
	light := get(graph.WaterLight, water)
	temp := get(graph.LightTemp, light)
	humid := get(graph.TempHumid, temp)
	loc := get(graph.HumidLoc, humid)

	return loc
}

func get(arr []Edge, start int) int {
	for _, e := range arr {
		if !e.Src.Contains(start) {
			continue
		}

		return e.Dest.Start + start - e.Src.Start
	}

	return start
}
