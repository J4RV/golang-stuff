# Genetic algorithm to finds good functions for approximating a cloud of points
*Or just GATFGFFAACOP, IDK TBH*

From an arbitrary array of points, find "the best" (almost) function that approximates them.

### Function structures implemented:

**"Cubes": (Should be renamed)**
`f(x) = <?> + <?>*x + <?>*x^2 + <?>*x^3`

**Sines2:**
`f(x) = <?> + <?>*x + <?>*sin(<?>*x + <?>) + <?>*sin(<?>*x + <?>)`

**Sines3:**
`f(x) = <?> + <?>*x + <?>*sin(<?>*x + <?>) + <?>*sin(<?>*x + <?>) + <?>*sin(<?>*x + <?>)`

### Examples:

![10 generations](https://raw.githubusercontent.com/j4rv/gostuff/master/approximations/img/10gens.png)

![100 generations](https://raw.githubusercontent.com/j4rv/gostuff/master/approximations/img/100gens.png)

![1000 generations](https://raw.githubusercontent.com/j4rv/gostuff/master/approximations/img/1000gens.png)