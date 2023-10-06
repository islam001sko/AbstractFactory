package main

type Samsung struct {
}

func (s *Samsung) sellPhone() IPhone {
	return &SamsungPhone{
		Phone: Phone{
			ram:   6,
			color: "black",
		},
	}
}
func (s *Samsung) sellLaptop() ILaptop {
	return &SamsungLaptop{
		Laptop: Laptop{
			ram:   16,
			color: "black",
		},
	}
}
func (s *Samsung) sellWatch() IWatch {
	return &SamsungWatch{
		Watch: Watch{
			ram:   5,
			color: "black",
		},
	}
}
