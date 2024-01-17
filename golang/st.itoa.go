package main

const a = iota // a=0
const (
	b = iota //b=0
	c        //c=1
)

type Newtype int

const (
	T1 Newtype = iota // 0
	T2                // 1
	T3                // 2
)

type AudioOutput int

const (
	OutMute   AudioOutput = iota // 0
	OutMono                      // 1
	OutStereo                    // 2
	_
	_
	OutSurround // 5
)

type Allergen int

const (
	IgEggs         Allergen = 1 << iota // 1 << 0 which is 00000001
	IgChocolate                         // 1 << 1 which is 00000010
	IgNuts                              // 1 << 2 which is 00000100
	IgStrawberries                      // 1 << 3 which is 00001000
	IgShellfish                         // 1 << 4 which is 00010000
)

type ByteSize float64

const (
	_           = iota             // ignore first value by assigning to blank identifier
	KB ByteSize = 1 << (10 * iota) // 1 << (10*1)
	MB                             // 1 << (10*2)
	GB                             // 1 << (10*3)
	TB                             // 1 << (10*4)
	PB                             // 1 << (10*5)
	EB                             // 1 << (10*6)
	ZB                             // 1 << (10*7)
	YB                             // 1 << (10*8)
)

const (
	Apple, Banana     = iota + 1, iota + 2 // Apple: 1 Banana: 2
	Cherimoya, Durian                      // Cherimoya: 2 Durian: 3
	Elderberry, Fig                        // Elderberry: 3 Fig: 4
)

const (
	a1 = iota // 0
	b1 = 5    // 5
	c1 = iota // 2
	d1 = 6    // 6
	e1        // 6
	f1        // 6
)
