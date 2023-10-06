package main

type Apple struct {
}

func (a *Apple) sellPhone() IPhone {
	return &ApplePhone{
		Phone: Phone{
			ram:   4,
			color: "black",
		},
	}
}
func (a *Apple) sellLaptop() ILaptop {
	return &AppleLaptop{
		Laptop: Laptop{
			ram:   14,
			color: "black",
		},
	}
}
func (a *Apple) sellWatch() IWatch {
	return &AppleWatch{
		Watch: Watch{
			ram:   4,
			color: "black",
		},
	}
}
