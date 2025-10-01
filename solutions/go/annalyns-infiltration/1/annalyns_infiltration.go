package annalyn

// CanFastAttack can be executed only when the knight is sleeping.
func CanFastAttack(knightIsAwake bool) bool {
	var cfs bool
	if knightIsAwake {
		cfs = false
	} else {
		cfs = true
	}
	return cfs
	//panic("Please implement the CanFastAttack() function")
}

// CanSpy can be executed if at least one of the characters is awake.
func CanSpy(knightIsAwake, archerIsAwake, prisonerIsAwake bool) bool {
	var canSpy bool
	if knightIsAwake || archerIsAwake || prisonerIsAwake {
		canSpy = true
	}
	return canSpy
	//panic("Please implement the CanSpy() function")
}

// CanSignalPrisoner can be executed if the prisoner is awake and the archer is sleeping.
func CanSignalPrisoner(archerIsAwake, prisonerIsAwake bool) bool {
	var csp bool

	if !archerIsAwake && prisonerIsAwake {
		csp = true
	}
	return csp
	//panic("Please implement the CanSignalPrisoner() function")
}

// CanFreePrisoner can be executed if the prisoner is awake and the other 2 characters are asleep
// or if Annalyn's pet dog is with her and the archer is sleeping.
func CanFreePrisoner(knightIsAwake, archerIsAwake, prisonerIsAwake, petDogIsPresent bool) bool {
	var cfp bool

	if (!knightIsAwake && !archerIsAwake && prisonerIsAwake) ||
		(!archerIsAwake && petDogIsPresent) {
		cfp = true
	}
	return cfp
	//panic("Please implement the CanFreePrisoner() function")
}
