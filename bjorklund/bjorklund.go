package bjorklund

func Bjorklund(steps int, pulses int) []int {
    if (pulses > steps) {
        panic("Pulses greater than steps")
    }

    var pattern intSlice
    var counts []int
    var remainders []int
    divisor := steps - pulses
    remainders = append(remainders, pulses)

    level := 0
    for {
        counts = append(counts, divisor / remainders[level])
        remainders = append(remainders, divisor % remainders[level])
        divisor = remainders[level]
        level = level + 1
        if remainders[level] <= 1 {
            break
        }
    }
    counts = append(counts, divisor)

    pattern = build(level, pattern, remainders, counts)
    i := pattern.pos(1)
    pattern = append(pattern[i:], pattern[0:i]...)

    return pattern
}

func build(level int, pattern []int, remainders []int, counts []int) []int {
    if level == -1 {
        pattern = append(pattern, 0)
    } else if level == -2 {
        pattern = append(pattern, 1)
    } else {
        for i := 0; i < counts[level]; i++ {
            pattern = build(level - 1, pattern, remainders, counts)
        }

        if remainders[level] != 0 {
            pattern = build(level - 2, pattern, remainders, counts)
        }
    }
    return pattern
}

type intSlice []int
func (slice intSlice) pos(value int) int {
    for p, v := range slice {
        if (v == value) {
            return p
        }
    }
    return -1
}

