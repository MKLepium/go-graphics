package world

import "sync"

type World struct {
	Chunks map[[2]int]*Chunk
	mu     sync.Mutex
}

func NewWorld() *World {
	return &World{
		Chunks: make(map[[2]int]*Chunk),
	}
}

func (w *World) GetChunk(x, z int) *Chunk {
	w.mu.Lock()
	defer w.mu.Unlock()
	chunkPos := [2]int{x, z}
	chunk, exists := w.Chunks[chunkPos]
	if !exists {
		chunk = NewChunk()
		w.Chunks[chunkPos] = chunk
	}
	return chunk
}

func (w *World) SetBlock(x, y, z int, block Block) {
	chunk := w.GetChunk(x/ChunkWidth, z/ChunkWidth)
	chunk.SetBlock(x%ChunkWidth, y, z%ChunkWidth, block)
}

func (w *World) GetBlock(x, y, z int) Block {
	chunk := w.GetChunk(x/ChunkWidth, z/ChunkWidth)
	return chunk.GetBlock(x%ChunkWidth, y, z%ChunkWidth)
}
