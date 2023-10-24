package main

func main() {

	group := make(map[string][]string)

	group["test"] = append(group["test"], "test1")

	println(group["test"][0])

	group["test"] = append(group["test"], "test2")

	println(group["test"][1])

}
