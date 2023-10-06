/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type Jwz struct {
	Key
	Attributes JwzAttributes `json:"attributes"`
}
type JwzResponse struct {
	Data     Jwz      `json:"data"`
	Included Included `json:"included"`
}

type JwzListResponse struct {
	Data     []Jwz    `json:"data"`
	Included Included `json:"included"`
	Links    *Links   `json:"links"`
}

// MustJwz - returns Jwz from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustJwz(key Key) *Jwz {
	var jWZ Jwz
	if c.tryFindEntry(key, &jWZ) {
		return &jWZ
	}
	return nil
}
