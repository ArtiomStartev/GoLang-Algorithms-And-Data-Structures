package cache

type CacheFunc func(int) int

// CacheFunction takes a function as an argument and returns a new function that caches the results of the original function
func CacheFunction(fn CacheFunc) CacheFunc {
	cache := make(map[int]int)

	return func(n int) int {
		if val, ok := cache[n]; ok {
			return val
		}

		result := fn(n)
		cache[n] = result

		return result
	}
}

// factorial is a simple example function that calculates the factorial of a number
func factorial(n int) int {
	result := 1
	for ; n > 1; n-- {
		result *= n
	}
	return result
}
