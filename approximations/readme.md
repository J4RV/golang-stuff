# Genetic algorithm to find good functions for approximating a cloud of points
*Or just GATFGFFAACOP, IDK TBH*

From an arbitrary array of points, find "the best" (almost) function that approximates them.

### Function structures implemented:

**"Cubes": (Should be renamed)**  
`f(x) = <?> + <?>*x + <?>*x^2 + <?>*x^3`

**Sines2:**  
`f(x) = <?> + <?>*x + <?>*sin(<?>*x + <?>) + <?>*sin(<?>*x + <?>)`

**Sines3:**  
`f(x) = <?> + <?>*x + <?>*sin(<?>*x + <?>) + <?>*sin(<?>*x + <?>) + <?>*sin(<?>*x + <?>)`

### Parameters and examples:

#### Generations

One of the most important things to configure correctly is the amount of generations.  
A higher amount means the resulting function will be closer to the "perfect" one, but it will take more time to calculate.

##### 10 generations  
![10 generations](https://raw.githubusercontent.com/j4rv/gostuff/master/approximations/img/10gens.png)

##### 100 generations  
![100 generations](https://raw.githubusercontent.com/j4rv/gostuff/master/approximations/img/100gens.png)

##### 1000 generations  
![1000 generations](https://raw.githubusercontent.com/j4rv/gostuff/master/approximations/img/1000gens.png)

#### Initial temperature

Another important parameter to configure is the 'initial temperature'.  
This parameter represents how much the <?> values are allowed to vary in the first generation.  
A higher value means they can reach higher maximum (absolute) values. **But this can also cause overfitting, specially with the sines functions.**

##### Initial temperature of 0.5  
![Initial temperature of 0.5](https://raw.githubusercontent.com/j4rv/gostuff/master/approximations/img/initTemp05.png)

##### Initial temperature of 8  
![Initial temperature of 8](https://raw.githubusercontent.com/j4rv/gostuff/master/approximations/img/initTemp8.png)

##### Initial temperature of 32  
![Initial temperature of 32](https://raw.githubusercontent.com/j4rv/gostuff/master/approximations/img/initTemp32.png)

With a temperature of 32 we can see that the function starts to overfit, so this temperature is **TOO HOT**.