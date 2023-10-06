/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type VerifyId struct {
	Key
	Attributes VerifyIdAttributes `json:"attributes"`
}
type VerifyIdResponse struct {
	Data     VerifyId `json:"data"`
	Included Included `json:"included"`
}

type VerifyIdListResponse struct {
	Data     []VerifyId `json:"data"`
	Included Included   `json:"included"`
	Links    *Links     `json:"links"`
}

// MustVerifyId - returns VerifyId from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustVerifyId(key Key) *VerifyId {
	var verifyID VerifyId
	if c.tryFindEntry(key, &verifyID) {
		return &verifyID
	}
	return nil
}
