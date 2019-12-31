package cardano

import (
    "fmt"
    "strconv"
    "strings"
    "time"
)

type Epoch uint64
type Slot uint64

type SlotDuration int64

// a plain Cardano slot date is only determined by its
// epoch and slot. A plain date is missing further
// information about the time settings of the Cardano
// blockchain.
//
// A plain date does not know the number of slots
// in  an epoch as well as the duration of a slot
// in milliseconds. Moreover, the date time of the
// genesis block is unknown. Hence, a number of methods
// are not available in the plain date, but can be
// used if the plain date is transformed into a full
// date by passing the required time settings of the
// blockchain.
type PlainSlotDate interface {
    // gets the epoch of the date.
    GetEpoch() Epoch
    // gets the slot of the date.
    GetSlot() Slot
    // returns true, if this date is equal to the given
    // other date, otherwise false.
    SameAs(otherDate PlainSlotDate) bool
    // returns true, if this date lies strictly before
    // the given other date, otherwise false.
    Before(otherDate PlainSlotDate) bool
    // returns true, if this date lies after the given
    // other date.
    After(otherDate PlainSlotDate) bool
    // returns the slot date as a plain string in the
    // format <EPOCH>.<SLOT>
    String() string
}

//
type FullSlotDate interface {
    PlainSlotDate
    // gets an instant in time with nanosecond precision of
    // the start of the slot.
    GetStartDateTime() time.Time
    // gets an instant in time with nanosecond precision of
    // the end of the slot.
    GetEndDateTime() time.Time
    // returns the number of slots that are between this
    // date and the other date.
    Diff(otherDate PlainSlotDate) SlotDuration
}

type TimeSettings struct {
    GenesisBlockDateTime time.Time
    SlotsPerEpoch        uint64
    SlotDuration         time.Duration
}

type parsingError struct {
    // text that should have been parsed.
    ParsedText string
    // reason why it cannot be parsed.
    Reason string
}

func (err parsingError) Error() string {
    return fmt.Sprintf("Failed to parse '%v'. %v", err.ParsedText, err.Reason)
}

// parses plain date from the given text. A date must
// be of the format "<EPOCH>.<SLOT>". if parsing fails,
// an error will be returned.
func ParsePlainData(text string) (PlainSlotDate, error) {
    seps := strings.Split(text, ".")
    if len(seps) == 2 {
        epoch, err := strconv.Atoi(seps[0])
        if err == nil {
            if epoch >= 0 {
                slot, err := strconv.Atoi(seps[1])
                if err == nil {
                    if slot >= 0 {
                        return plainDateImpl{epoch: Epoch(epoch), slot: Slot(slot)}, nil
                    }
                }
                return nil, parsingError{ParsedText: text, Reason: fmt.Sprintf("Slot must be a positive number, but was '%v'.", seps[1])}
            }
        }
        return nil, parsingError{ParsedText: text, Reason: fmt.Sprintf("Epoch must be a positive number, but was '%v'.", seps[0])}
    } else {
        return nil, parsingError{ParsedText: text, Reason: "The date must be of the format '<EPOCH>.<SLOT>', where epoch and slot are positive numbers."}
    }
}

func GetPlainSlotDate(epoch Epoch, slot Slot) PlainSlotDate {
    return plainDateImpl{epoch: epoch, slot: slot}
}

type plainDateImpl struct {
    epoch Epoch
    slot  Slot
}

type fullDateImpl struct {
    plainDateImpl
    timeSettings TimeSettings
}

func (date plainDateImpl) GetEpoch() Epoch {
    return date.epoch
}

func (date plainDateImpl) GetSlot() Slot {
    return date.slot
}

func (date plainDateImpl) SameAs(otherDate PlainSlotDate) bool {
    return (date.GetEpoch() == otherDate.GetEpoch()) && (date.GetSlot() == otherDate.GetSlot())
}

func (date plainDateImpl) Before(otherDate PlainSlotDate) bool {
    if date.GetEpoch() < otherDate.GetEpoch() {
        return true
    } else if date.GetEpoch() == otherDate.GetEpoch() {
        if date.GetSlot() < otherDate.GetSlot() {
            return true
        }
    }
    return false
}

func (date plainDateImpl) After(otherDate PlainSlotDate) bool {
    if date.GetEpoch() > otherDate.GetEpoch() {
        return true
    } else if date.GetEpoch() == otherDate.GetEpoch() {
        if date.GetSlot() > otherDate.GetSlot() {
            return true
        }
    }
    return false
}

func (date plainDateImpl) String() string {
    return fmt.Sprintf("%v.%v", date.GetEpoch(), date.GetSlot())
}

func FormTimeSettings(genesisBlockDateTime time.Time, slotsPerEpoch uint64, slotDuration time.Duration) TimeSettings {
    return TimeSettings{GenesisBlockDateTime: genesisBlockDateTime, SlotsPerEpoch: slotsPerEpoch, SlotDuration: slotDuration}
}

func GetFullSlotDate(epoch Epoch, slot Slot, settings TimeSettings) FullSlotDate {
    return fullDateImpl{plainDateImpl: plainDateImpl{epoch: epoch, slot: slot}, timeSettings: settings}
}

func MakeFullSlotDate(plainDate PlainSlotDate, settings TimeSettings) FullSlotDate {
    return fullDateImpl{plainDateImpl: plainDateImpl{epoch: plainDate.GetEpoch(), slot: plainDate.GetSlot()}, timeSettings: settings}
}

func (date fullDateImpl) Diff(otherDate PlainSlotDate) SlotDuration {
    a := uint64(date.GetEpoch())*date.timeSettings.SlotsPerEpoch + uint64(date.GetSlot())
    b := uint64(otherDate.GetEpoch())*date.timeSettings.SlotsPerEpoch + uint64(otherDate.GetSlot())
    return SlotDuration(int64(a) - int64(b))
}

func (date fullDateImpl) GetStartDateTime() time.Time {
    slots := uint64(date.GetEpoch())*date.timeSettings.SlotsPerEpoch + uint64(date.GetSlot())
    return date.timeSettings.GenesisBlockDateTime.Add(time.Duration(slots) * date.timeSettings.SlotDuration)
}

func (date fullDateImpl) GetEndDateTime() time.Time {
    return date.GetStartDateTime().Add(date.timeSettings.SlotDuration)
}
