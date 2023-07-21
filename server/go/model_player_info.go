/*
 * Translator service between a srb2kart server and json
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type PlayerInfo []PlayerInfoEntry

// AssertPlayerInfoRequired checks if the required fields are not zero-ed
func AssertPlayerInfoRequired(obj PlayerInfo) error {
	for _, el := range obj {
		if err := AssertPlayerInfoEntryRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertRecursePlayerInfoRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of PlayerInfo (e.g. [][]PlayerInfo), otherwise ErrTypeAssertionError is thrown.
func AssertRecursePlayerInfoRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aPlayerInfo, ok := obj.(PlayerInfo)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertPlayerInfoRequired(aPlayerInfo)
	})
}
