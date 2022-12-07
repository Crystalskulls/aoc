package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Dir struct {
	name      string
	path      string
	parentDir *Dir
	size      int
	files     []*File
	subDirs   []*Dir
}

type File struct {
	name string
	size int
}

func main() {
	file, _ := os.ReadFile("input.txt")
	commands := strings.Split(string(file), "\n")
	root, dirMap := parseDirectoryTree(commands)
	calcDirSize(root)
	partOne := 0
	for _, dir := range dirMap {
		if dir.size <= 100000 {
			partOne += dir.size
		}

	}
	fmt.Println(partOne)
	space := 30000000 - (70000000 - root.size)
	pd := make([]int, 0)
	for _, dir := range dirMap {
		if dir.size > space {
			pd = append(pd, dir.size)
		}
	}
	sort.Ints(pd)
	fmt.Println(pd[0])
}

func calcDirSize(dir *Dir) int {
	if len(dir.subDirs) == 0 {
		return dir.size
	}
	for _, subdir := range dir.subDirs {
		dir.size += calcDirSize(subdir)
	}
	return dir.size
}

func parseDirectoryTree(commands []string) (root *Dir, dirMap map[string]*Dir) {
	root = new(Dir)
	dirMap = make(map[string]*Dir, 0)
	var currentDir *Dir
	currentPath := ""
	for _, command := range commands {
		cs := strings.Split(command, " ")
		if cs[0] == "$" {
			if cs[1] == "cd" && cs[2] == "/" {
				currentPath = "/"
				dir := createDir(cs[2], currentPath, nil)
				dirMap[dir.path] = dir
				currentDir = dir
				root = dir
			} else if cs[1] == "cd" && cs[2] == ".." {
				i := strings.LastIndex(currentPath[:len(currentPath)-1], "/")
				currentPath = currentPath[:i+1]
			} else if cs[1] == "cd" {
				currentPath += cs[2] + "/"
				currentDir = dirMap[currentPath]
			}
		} else {
			if cs[0] == "dir" {
				dir := createDir(cs[1], currentPath+cs[1]+"/", currentDir.parentDir)
				dirMap[dir.path] = dir
				currentDir.subDirs = append(currentDir.subDirs, dir)
			} else {
				file := createFile(cs[1], cs[0])
				currentDir.files = append(currentDir.files, file)
				currentDir.size += file.size
			}
		}

	}
	return
}

func createDir(name string, path string, parent *Dir) *Dir {
	dir := new(Dir)
	dir.name = name
	dir.files = make([]*File, 0)
	dir.subDirs = make([]*Dir, 0)
	dir.parentDir = parent
	dir.path = path
	return dir
}

func createFile(name string, size string) *File {
	file := new(File)
	file.name = name
	file.size, _ = strconv.Atoi(size)
	return file
}

func (dir *Dir) String() string {
	return fmt.Sprintf("Name: %s, Size: %d, Path: %s, Files: %+v, Subdirs: %#v", dir.name, dir.size, dir.path, dir.files, dir.subDirs)
}
