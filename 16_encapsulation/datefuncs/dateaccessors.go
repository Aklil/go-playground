package dateencapsulator


import (
	"errors"
)


// funcs and struct members, capitalized are public while those in lower case are private


type Date struct{
  day int   // this is private since it starts with lower
  month int
  year int	
}


func (d *Date) SetDay(day int) error{
	// validate
	
	if day < 1 || day > 31 {
		return errors.New("invalid day")
	}
	
	d.day = day
	return nil
}

func (d * Date) SetMonth(month int) error{
	if month < 1 || month > 12{
		return errors.New("invalid month")
	}

	d.month = month
	return nil
}


func (d *Date) SetYear(year int) error {

	if year < 1900 || year > 2080 {
		return errors.New("invalid year")
	}

	d.year = year
	return nil
}

func (d *Date) GetDay() int {
	return d.day
}

func (d *Date) GetMonth() int {
	return d.month
}

func (d *Date) GetYear() int {
	return d.year
}



