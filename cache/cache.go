package cache

/**
 * Higher-order function to cache results of another function.
 * @param fn func(int) int - The function whose results need to be cached.
 * @returns func(int) int - Cached version of the input function.
 */

func cacheFunction(fn func(int) int) func(int) int {
	cache := make(map[int]int)

	return func(n int) int {
		if val, ok := cache[n]; ok {
			// Return cached result if available
			return val
		}

		result := fn(n)   // Compute result using the original function
		cache[n] = result // Store result in cache for future use

		return result
	}
}

/**
 * Computes the factorial of a given number.
 * @param n int - The number for which factorial is to be computed.
 * @returns int - The factorial of the input number.
 */

func factorial(n int) int {
	result := 1

	for ; n > 1; n-- {
		result *= n
	}

	return result
}
