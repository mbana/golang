package pcsa

import (
	"bytes"
	"fmt"
	"hash"      // https://golang.org/pkg/hash/
	"hash/fnv"  // https://golang.org/pkg/hash/fnv/
	"math/bits" // https://golang.org/pkg/math/bits/ and https://graphics.stanford.edu/~seander/bithacks.html

	"github.com/banaio/golang/countdistinct/lib"
)

var (
	ErrInvalidBitmapsSize = fmt.Errorf("bitmapsSize has to be >= 4 and <= 16")
)

// NewPCSA - bitmapsSize should be within [4...16]
func NewPCSA(bitmapsSize uint8) (lib.Set, error) {
	if bitmapsSize < 4 || bitmapsSize > 16 {
		return nil, ErrInvalidBitmapsSize
	}

	bitmaps := make([]uint32, bitmapsSize)
	return &pcsa{
		bitmapsSize: bitmapsSize,
		bitmaps:     bitmaps,
		hasher:      fnv.New32(),
	}, nil
}

type pcsa struct {
	bitmapsSize uint8
	bitmaps     []uint32
	hasher      hash.Hash32
}

func (p *pcsa) Add(value []byte) error {
	p.hasher.Reset()
	_, err := p.hasher.Write(value)
	if err != nil {
		return err
	}
	hash := p.hasher.Sum32()
	p.hasher.Reset()

	// fmt.Printf("pcsa.Add:           value = %+v\n", string(value))
	// fmt.Printf("pcsa.Add:            hash = %+v\n", hash)
	fmt.Printf("pcsa.Add:            hash = %032b\n", hash)
	fmt.Printf("pcsa.Add:           ^hash = %032b\n", ^hash)
	fmt.Printf("pcsa.Add:            hash = %032b\n", hash)
	fmt.Printf("pcsa.Add:   LeadingOnes32 = %d\n", LeadingOnes32(hash))
	fmt.Printf("pcsa.Add:  TrailingOnes32 = %d\n", TrailingOnes32(hash))
	fmt.Printf("pcsa.Add:  LeadingZeros32 = %d\n", bits.LeadingZeros32(hash))
	fmt.Printf("pcsa.Add: TrailingZeros32 = %d\n", bits.TrailingZeros32(hash))
	fmt.Printf("\n")

	return nil
}

func (p *pcsa) Count() uint32 {
	return 0
}

func (p *pcsa) String() string {
	buffer := bytes.NewBufferString("")
	fmt.Fprintf(buffer, "PCSA\n")
	fmt.Fprintf(buffer, "Count=%d\n", p.Count())
	fmt.Fprintf(buffer, "Bitmaps\n")
	for bitmapNo, bitmap := range p.bitmaps {
		fmt.Fprintf(buffer, "  %2d %032b\n", bitmapNo, bitmap)
	}
	return buffer.String()
}

func TrailingOnes32(hash uint32) int {
	return bits.TrailingZeros32(^hash)
}

func LeadingOnes32(hash uint32) int {
	// reverses the bit, Bitwise Complement
	// ^x = 1 ^ x
	//
	//
	// pcsa.Add:            hash = 11111111111111111111111111111101
	// pcsa.Add:           ^hash = 00000000000000000000000000000010
	// pcsa.Add:   LeadingOnes32 = 30
	// pcsa.Add:  TrailingOnes32 = 1
	// pcsa.Add:  LeadingZeros32 = 0
	// pcsa.Add: TrailingZeros32 = 0
	// pcsa.Add:   LeadingOnes32 = 30

	// very crap but gets the job done
	leadingOnes32 := 0
	for _, bit := range fmt.Sprintf("%032b", hash) {
		if bit == '0' {
			break
		}
		leadingOnes32++
	}

	if leadingOnes32 != bits.LeadingZeros32(^hash) {
		panic(fmt.Errorf("leadingOnes32=%d != bits.LeadingZeros32(^hash)=%d", leadingOnes32, bits.LeadingZeros32(^hash)))
	}

	return bits.LeadingZeros32(^hash)
}
