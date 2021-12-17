package constans

type Constans struct {
	SHOP struct {
		Name string
	}
}

func NewConstans() *Constans {
	c := new(Constans)

	c.SHOP.Name = "Art"

	return c
}
