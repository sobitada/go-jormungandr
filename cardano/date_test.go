package cardano

import (
    "testing"
    "time"
)

func TestPlainSlotDate_ParseCorrectSlotDateString(t *testing.T) {
    plainSlotDate, err := ParsePlainData("4.15")
    if err != nil {
        t.Errorf("Parsing must be successful, but throw error: %v", err.Error())
    } else if plainSlotDate == nil {
        t.Error("Result of the parsing must not be nil.")
    } else if plainSlotDate.GetEpoch() != 4 {
        t.Errorf("Epoch of parsed slot date must be 4, but was %v.", plainSlotDate.GetEpoch())
    } else if plainSlotDate.GetSlot() != 15 {
        t.Errorf("Slot of parsed slot date must be 15, but was %v.", plainSlotDate.GetSlot())
    }
}

func TestPlainSlotDate_ParseSlotDateStringInWrongFormat_mustThrowError(t *testing.T) {
    _, err := ParsePlainData("1-17")
    if err == nil {
        t.Errorf("'1-17' is in the wrong format and parser must return an error.")
    }
    _, err = ParsePlainData("1.A")
    if err == nil {
        t.Errorf("'1.A' is in the wrong format and parser must return an error.")
    }
    _, err = ParsePlainData("A.444")
    if err == nil {
        t.Errorf("'A.444' is in the wrong format and parser must return an error.")
    }
    _, err = ParsePlainData("-1.444")
    if err == nil {
        t.Errorf("'-1.444' is in the wrong format and parser must return an error.")
    }
    _, err = ParsePlainData("2.-666")
    if err == nil {
        t.Errorf("'2.-666' is in the wrong format and parser must return an error.")
    }
}

func TestPlainSlotDate_CompareEqualityOfSlotDates_mustReturnTrue(t *testing.T) {
    check := GetPlainSlotDate(2, 15).SameAs(GetPlainSlotDate(2, 15))
    if !check {
        t.Errorf("Both slots are the same, and same as must return true.")
    }
}

func TestPlainSlotDate_CompareEqualityOfSlotDates_mustReturnFalse(t *testing.T) {
    check := GetPlainSlotDate(2, 15).SameAs(GetPlainSlotDate(3, 15))
    if check {
        t.Errorf("Both slots are the same, and same as must return true.")
    }
}

func TestPlainSlotDate_IsSlotDateABeforeSlotDateB_mustReturnTrue(t *testing.T) {
    check := GetPlainSlotDate(2, 15).Before(GetPlainSlotDate(2, 16))
    if !check {
        t.Errorf("Slot date 2.15 is before 2.16, so before must return true.")
    }
}

func TestPlainSlotDate_IsSlotDateABeforeSlotDateB_mustReturnFalse(t *testing.T) {
    check := GetPlainSlotDate(2, 15).Before(GetPlainSlotDate(2, 15))
    if check {
        t.Errorf("Slot date 2.15 is the same as 2.15, so before must return false.")
    }
}

func TestPlainSlotDate_IsSlotDateAAfterSlotDateB_mustReturnTrue(t *testing.T) {
    check := GetPlainSlotDate(100, 17).After(GetPlainSlotDate(100, 16))
    if !check {
        t.Errorf("Slot date 100.7 is after 100.16, so after must return true.")
    }
}

func TestPlainSlotDate_IsSlotDateAAfterSlotDateB_mustReturnFalse(t *testing.T) {
    check := GetPlainSlotDate(2, 15).After(GetPlainSlotDate(2, 15))
    if check {
        t.Errorf("Slot date 2.15 is the same as 2.15, so after must return false.")
    }
}

func TestPlainSlotDate_ADiffB_mustReturnPositiveSlotDuration(t *testing.T) {
    // setup
    genesisTime, _ := time.Parse(time.RFC3339, "2019-12-13T19:13:37+00:00")
    settings := TimeSettings{GenesisBlockDateTime: genesisTime, SlotsPerEpoch: 43200, SlotDuration: time.Duration(2) * time.Second}
    dateA := GetFullSlotDate(17, 12000, settings)
    dateB := GetFullSlotDate(16, 35600, settings)
    if dateA.Diff(dateB) != 19600 {
        t.Errorf("The difference between the two slot dates must be '19600', but was '%v'.", dateA.Diff(dateB))
    }
}

func TestPlainSlotDate_ADiffB_mustReturnNegativeSlotDuration(t *testing.T) {
    // setup
    genesisTime, _ := time.Parse(time.RFC3339, "2019-12-13T19:13:37+00:00")
    settings := TimeSettings{GenesisBlockDateTime: genesisTime, SlotsPerEpoch: 43200, SlotDuration: time.Duration(2) * time.Second}
    dateA := GetFullSlotDate(8, 40100, settings)
    dateB := GetFullSlotDate(10, 35600, settings)
    if dateA.Diff(dateB) != -81900 {
        t.Errorf("The difference between the two slot dates must be '-81900', but was '%v'.", dateA.Diff(dateB))
    }
}

func TestFullSlotDate_GetStartAndEndTime(t *testing.T) {
    // setup
    genesisTime, _ := time.Parse(time.RFC3339, "2019-12-13T19:13:37+00:00")
    settings := TimeSettings{GenesisBlockDateTime: genesisTime, SlotsPerEpoch: 43200, SlotDuration: time.Duration(2) * time.Second}
    date := GetFullSlotDate(17, 10653, settings)
    // test
    expectedStart, _ := time.Parse(time.RFC3339, "2019-12-31T02:08:43+01:00")
    diff := date.GetStartDateTime().Sub(expectedStart)
    if diff != 0 {
        t.Errorf("The start time must be at '2019-12-31T02:08:43+01:00', but there was a '%s' difference.", diff.String())
    }
}
