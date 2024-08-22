/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type VerifyRequest struct {
	Key
	Attributes VerifyRequestAttributes `json:"attributes"`
}
type VerifyRequestResponse struct {
	Data     VerifyRequest `json:"data"`
	Included Included      `json:"included"`
}

type VerifyRequestListResponse struct {
	Data     []VerifyRequest `json:"data"`
	Included Included        `json:"included"`
	Links    *Links          `json:"links"`
}

// MustVerifyRequest - returns VerifyRequest from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustVerifyRequest(key Key) *VerifyRequest {
	var verifyRequest VerifyRequest
	if c.tryFindEntry(key, &verifyRequest) {
		return &verifyRequest
	}
	return nil
}
