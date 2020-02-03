package api

import (
    "math/big"
    "net/url"
    "sort"
    "time"
)

// combines the base URL with the given path. if the given
// path is not valid, then an error will be returned.
func combine(base *url.URL, path string) (*url.URL, error) {
    var ref, err = url.Parse(path)
    if err == nil {
        return base.ResolveReference(ref), nil
    }
    return nil, err
}

// sort the leaders logs by the scheduled time.
func SortLeaderLogsByScheduleTime(leaderAssignments []LeaderAssignment) []LeaderAssignment {
    sortedAssignment := leaderAssignments[:]
    sort.Slice(sortedAssignment, func(i, j int) bool {
        return leaderAssignments[i].ScheduleTime.Before(leaderAssignments[j].ScheduleTime)
    })
    return sortedAssignment
}

// returns only a list of the leader assignments that are strictly after the given date.
func FilterLeaderLogsBefore(before time.Time, leaderAssignments []LeaderAssignment) []LeaderAssignment {
    var filteredAssignments []LeaderAssignment
    for i := range leaderAssignments {
        currentAssignment := leaderAssignments[i]
        if currentAssignment.ScheduleTime.After(before) {
            filteredAssignments = append(filteredAssignments, currentAssignment)
        }
    }
    return filteredAssignments
}

// gets all the leader assignments in the given epoch and ignores
// all the other assignments.
func GetLeaderLogsInEpoch(epoch *big.Int, leaderAssignments []LeaderAssignment) []LeaderAssignment {
    filteredAssignments := make([]LeaderAssignment, 0)
    for i := range leaderAssignments {
        currentAssignment := leaderAssignments[i]
        if currentAssignment.ScheduleBlockDate.GetEpoch().Cmp(epoch) == 0 {
            filteredAssignments = append(filteredAssignments, currentAssignment)
        }
    }
    return filteredAssignments
}

// gets all the leader assignments for the leader with the given ID.
func GetLeaderLogsOfLeader(leaderID uint64, leaderAssignments []LeaderAssignment) []LeaderAssignment {
    filteredAssignments := make([]LeaderAssignment, 0)
    for i := range leaderAssignments {
        currentAssignment := leaderAssignments[i]
        if leaderID == currentAssignment.LeaderID {
            filteredAssignments = append(filteredAssignments, currentAssignment)
        }
    }
    return filteredAssignments
}
