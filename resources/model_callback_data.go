/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type CallbackData struct {
	Key
	Attributes CallbackDataAttributes `json:"attributes"`
}
type CallbackDataResponse struct {
	Data     CallbackData `json:"data"`
	Included Included     `json:"included"`
}

type CallbackDataListResponse struct {
	Data     []CallbackData `json:"data"`
	Included Included       `json:"included"`
	Links    *Links         `json:"links"`
}

// MustCallbackData - returns CallbackData from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustCallbackData(key Key) *CallbackData {
	var callbackData CallbackData
	if c.tryFindEntry(key, &callbackData) {
		return &callbackData
	}
	return nil
}
