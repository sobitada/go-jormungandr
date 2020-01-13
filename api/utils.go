package api

import (
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

func SortLeaderLogsByScheduleTime(leaderAssignments []LeaderAssignment) []LeaderAssignment {
    sortedAssignment := leaderAssignments[:]
    sort.Slice(sortedAssignment, func(i, j int) bool {
        return leaderAssignments[i].ScheduleTime.Before(leaderAssignments[j].ScheduleTime)
    })
    return sortedAssignment
}

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
