package family

import "fmt"

type Family struct {}

func (f *Family) Collection() string {
	return "Family"
}

func (f *Family) FindByID(id string) {
	fmt.Println(f.Collection(), id)
}

func (f *Family) FindByIDs(ids []string) ([]family, error) {
	fmt.Println(f.Collection())

	families := []family{}
	for _, id := range ids {
		families = append(families, family{id})
	}
	return families, nil
}
