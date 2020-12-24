package main

func main() {
	println("hello")

	{
		var i int
		for i < 10 {
			if i == 5 {
				break
			}
			i++
			if i == 3 {
				continue
			}
			println(i)
		}
	}

	{
		for i := 0; i < 5; i++ {
			println(i)
		}
	}

	{
		var i int
		for {
			if i == 5 {
				break
			}
			i++
			if i == 3 {
				continue
			}
			println(i)
		}
	}

	{
		s := []int{1, 5, 9}
		for i := 0; i < len(s); i++ {
			println(s[i])
		}
	}
	{
		s := []int{1, 5, 9}
		for i,v := range s {
			println(i,v)
		}
	}

	{
		s := []int{1, 5, 9}
		for i := range s {
			println(i)
		}
	}

	{
		s := []int{1, 5, 9}
		for _,v := range s {
			println(v)
		}
	}

	{
		m := map[string]int{"http": 3000, "https": 433}
		for k,v := range m {
			println(k,v)
		}

		for k := range m {
			println(k)
		}

		for _,v := range m {
			println(v)
		}
	}
}
