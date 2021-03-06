package tempconv

func CtoF(c Celcius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func FtoC(f Fahrenheit) Celcius { return Celcius((f-32)*5/9) }

func CtoK(c Celcius) Kelvin { return Kelvin(-1*AbsoluteZeroC + c)}
func KtoC(k Kelvin) Celcius { return Celcius(AbsoluteZeroC + k) }

func KtoF(k Kelvin) Fahrenheit { return Fahrenheit(CtoF(KtoC(k))) }


