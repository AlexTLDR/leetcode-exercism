// https://exercism.org/tracks/go/exercises/election-day

package main

import "fmt"

type ElectionResult struct {
	Name  string
	Votes int
}

var finalResults = map[string]int{
	"Mary": 10,
	"John": 51,
}

func main() {
	votes := 3
	voteCounter := &votes
	var nilCounter *int
	fmt.Println(VoteCount(voteCounter))
	fmt.Println(VoteCount(nilCounter))
	IncrementVoteCount(voteCounter, 2)
	fmt.Println(*voteCounter)

	//var result *ElectionResult
	result := &ElectionResult{
		Name:  "",
		Votes: 0,
	}
	result.Name = "Alex"
	result.Votes = 1
	NewElectionResult("Alex", 100)
	fmt.Printf("Election result: %+v\n", result)
	DisplayResult(result)
	DecrementVotesOfCandidate(finalResults, "Mary")
	fmt.Println(finalResults)
}

// NewVoteCounter returns a new vote counter with
// a given number of initial votes.
func NewVoteCounter(initialVotes int) *int {
	return &initialVotes
}

// VoteCount extracts the number of votes from a counter.
func VoteCount(counter *int) int {
	if counter == nil {
		return 0
	}
	return *counter
}

// IncrementVoteCount increments the value in a vote counter.
func IncrementVoteCount(counter *int, increment int) {
	*counter = *counter + increment
}

// NewElectionResult creates a new election result.
func NewElectionResult(candidateName string, votes int) *ElectionResult {
	return &ElectionResult{Name: candidateName, Votes: votes}
}

// DisplayResult creates a message with the result to be displayed.
func DisplayResult(result *ElectionResult) string {
	return fmt.Sprintf("%s (%d)", result.Name, result.Votes)
}

// DecrementVotesOfCandidate decrements by one the vote count of a candidate in a map.
func DecrementVotesOfCandidate(results map[string]int, candidate string) {
	results[candidate]--
}
