package world

const ChunkWidth = 16
const ChunkHeight = 256

type Chunk struct {
	Blocks [ChunkWidth][ChunkWidth][ChunkHeight]Block
}

func NewChunk() *Chunk {
	return &Chunk{}
}

func (c *Chunk) SetBlock(x, y, z int, block Block) {
	if x >= 0 && x < ChunkWidth && y >= 0 && y < ChunkWidth && z >= 0 && z < ChunkHeight {
		c.Blocks[x][y][z] = block
	}
}

func (c *Chunk) GetBlock(x, y, z int) Block {
	if x >= 0 && x < ChunkWidth && y >= 0 && y < ChunkWidth && z >= 0 && z < ChunkHeight {
		return c.Blocks[x][y][z]
	}
	return Block{}
}
