package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Tuple[T any, K any] struct {
	First  T
	Second K
}

type BySecond []Tuple[string, int]

func (a BySecond) Len() int           { return len(a) }
func (a BySecond) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a BySecond) Less(i, j int) bool { return a[i].Second < a[j].Second }

type File struct {
	size int
	name string
}

type Directory struct {
	name      string
	files     []*File
	dirs      []*Directory
	parentDir *Directory
}

func (currentDir *Directory) changeDirectory(name string) *Directory {
	for i := 0; i < len(currentDir.dirs); i++ {
		if currentDir.dirs[i].name == name {
			return currentDir.dirs[i]
		}
	}
	panic("folder not found")
}

func NewDirectory(name string, files []*File, dirs []*Directory, parentDir *Directory) *Directory {
	if files == nil {
		files = []*File{}
	}
	if dirs == nil {
		dirs = []*Directory{}
	}
	return &Directory{
		name:      name,
		files:     files,
		dirs:      dirs,
		parentDir: parentDir,
	}
}

func NewFile(name string, size int) *File {
	return &File{
		name: name,
		size: size,
	}
}

func main() {
	file, err := os.Open("input.txt")
	defer file.Close()

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	currDir := NewDirectory("root", nil, nil, nil)
	currDir.dirs = append(currDir.dirs, NewDirectory("/", nil, nil, currDir))
	rootDir := currDir

	for scanner.Scan() {
		token := scanner.Text()
		tokenSplit := strings.Split(token, " ")

		if isCdCommand(tokenSplit) {
			if tokenSplit[2] == ".." {
				currDir = currDir.parentDir
			} else {
				currDir = currDir.changeDirectory(tokenSplit[2])
			}
		} else if isDir(tokenSplit) {
			currDir.dirs = append(currDir.dirs, NewDirectory(tokenSplit[1], nil, nil, currDir))
		} else if ok, size := isFile(tokenSplit); ok {
			currDir.files = append(currDir.files, NewFile(tokenSplit[1], size))
		}
	}

	printDir(rootDir, 0)

	dirSizes := []Tuple[string, int]{}
	recursiveFileSizeCounter(rootDir, &dirSizes)

	atMost100_000 := 0
	for i := 0; i < len(dirSizes); i++ {
		size := dirSizes[i].Second
		if size < 100_000 {
			atMost100_000 += size
		}
	}

	// println(atMost100_000)

	sort.Sort(BySecond(dirSizes))
	totalSpace := 70_000_000
	leastUnused := 30_000_000
	used := dirSizes[len(dirSizes)-1].Second
	freeSpace := totalSpace - used

	for i := 0; i < len(dirSizes); i++ {
		size := dirSizes[i].Second
		if freeSpace+size >= leastUnused {
			println(size)
			break
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

func isCdCommand(token []string) bool {
	return token[1] == "cd"
}

func isFile(token []string) (bool, int) {
	size, err := strconv.Atoi(token[0])
	if err != nil {
		return false, 0
	}
	return true, size
}

func isDir(token []string) bool {
	return token[0] == "dir"
}

func recursiveFileSizeCounter(dir *Directory, dirSizes *[]Tuple[string, int]) int {
	currDirFilesSize := 0
	for i := 0; i < len(dir.files); i++ {
		currDirFilesSize += dir.files[i].size
	}

	for i := 0; i < len(dir.dirs); i++ {
		currDirFilesSize += recursiveFileSizeCounter(dir.dirs[i], dirSizes)
	}

	tuple := Tuple[string, int]{First: dir.name, Second: currDirFilesSize}
	*dirSizes = append(*dirSizes, tuple)
	return currDirFilesSize
}

func printDir(dir *Directory, depth int) {
	printSpace(depth)
	fmt.Printf("- %v (dir)\n", dir.name)

	for i := 0; i < len(dir.dirs); i++ {
		printDir(dir.dirs[i], depth+5)
	}

	for i := 0; i < len(dir.files); i++ {
		printSpace(depth + 5)
		fmt.Printf("- %v (file, size=%v)\n", dir.files[i].name, dir.files[i].size)
	}
}

func printSpace(size int) {
	for i := 0; i < size; i++ {
		fmt.Print(" ")
	}
}
