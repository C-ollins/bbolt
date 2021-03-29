// +build mips mipsle

package bbolt

// maxMapSize represents the largest mmap size supported by Bolt.
const maxMapSize = 0x1F400000 // 500MB

// maxAllocSize is the size used when creating array pointers.
const maxAllocSize = 0xFFFFFFF // 268MB
