package charging

type ScopedHazard struct {
	ID     string `json:"id"`
	Tenant string `json:"tenant"`
	Status string `json:"status"`
}
type HazardRepository struct{ items []ScopedHazard }

func NewHazardRepository() *HazardRepository {
	return &HazardRepository{items: []ScopedHazard{{"h1", "runan", "open"}, {"h2", "neighbor", "open"}, {"h3", "runan", "closed"}}}
}
func (r *HazardRepository) List(tenant, status string) []ScopedHazard {
	out := []ScopedHazard{}
	for _, h := range r.items {
		if h.Tenant == tenant && h.Status == status {
			out = append(out, h)
		}
	}
	return out
}
func (r *HazardRepository) Count(tenant, status string) int {
	count := 0
	for _, h := range r.items {
		if h.Tenant == tenant && h.Status == status {
			count++
		}
	}
	return count
}
