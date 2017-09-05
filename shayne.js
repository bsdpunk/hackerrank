#!/bin/node
/**
 *  * iterate through some shit till theres no more shit starting with an
 *   * random index of an array.
 *    *
 *     * @param  {[type]} prisoners  [description]
 *      * @param  {[type]} food       [description]
 *       * @param  {[type]} startIndex [description]
 *        *
 *         * @return {[type]}            [description]
 *          */
function feed(prisoners, food, startIndex) {

        startIndex = startIndex || Math.floor(Math.random() * prisoners.length) -1;

            console.log(`\nfeeding ${prisoners.length} prisoners starting with prisoner id: ${startIndex} until ${food.length} meals are gone \n`);

                for (var i = 0; i < food.length; i++) {
                            let prisoner = prisoners[(startIndex + i) % prisoners.length];
                                    prisoner.eatCount++
                                            }

                    console.log(prisoners);
}

var prisoners = [
        { id: 1, eatCount: 0 },
            { id: 2, eatCount: 0 },
                { id: 3, eatCount: 0 },
                    { id: 4, eatCount: 0 },
                        { id: 5, eatCount: 0 },
                            { id: 6, eatCount: 0 }
];
var food = [1, 3, 4, 5, 6, 7, 8, 9, 10];

feed(prisoners, food, 3);

