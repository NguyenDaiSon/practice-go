func canFinishEating(piles []int, hours int, speed int) bool {
    hoursUsed := 0
    for _, pile := range piles {
        hoursUsed += (pile + speed - 1)/speed
        if hoursUsed > hours {
            return false
        }
    }
    return hoursUsed <= hours
}

func minEatingSpeed(piles []int, hours int) int {
    minBanana := 1
    maxBanana := slices.Max(piles)
    minSpeed := -1
    for minBanana <= maxBanana {
        bananas := minBanana + (maxBanana - minBanana)/2
        if canFinishEating(piles, hours, bananas) {
            minSpeed = bananas
            maxBanana = bananas - 1
        } else {
            minBanana = bananas + 1
        }
    }
    return minSpeed
}
