/*
 * Translator service between a srb2kart server and json
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type PlayerInfoEntry struct {

	Node uint8 `json:"Node"`

	Name string `json:"Name"`

	Address string `json:"Address"`

	Team uint8 `json:"Team"`

	Skin uint8 `json:"Skin"`

	Data uint8 `json:"Data"`

	Score uint32 `json:"Score"`

	TimeInServer uint16 `json:"TimeInServer"`
}

// AssertPlayerInfoEntryRequired checks if the required fields are not zero-ed
func AssertPlayerInfoEntryRequired(obj PlayerInfoEntry) error {
	return nil
}

// AssertRecursePlayerInfoEntryRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of PlayerInfoEntry (e.g. [][]PlayerInfoEntry), otherwise ErrTypeAssertionError is thrown.
func AssertRecursePlayerInfoEntryRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aPlayerInfoEntry, ok := obj.(PlayerInfoEntry)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertPlayerInfoEntryRequired(aPlayerInfoEntry)
	})
}
