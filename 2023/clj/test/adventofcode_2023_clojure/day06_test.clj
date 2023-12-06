(ns adventofcode-2023-clojure.day06-test
  (:require [clojure.test :refer :all]
            [adventofcode-2023-clojure.day06 :refer :all]))

(def input-data  (slurp "resources/day06.txt"))
(def test-data  (slurp "resources/day06_test.txt"))

(deftest part1-test
  (are [expected input] (= expected (part1 input))
                        288 test-data
                        275724 input-data)
)

(deftest part2-test
  (are [expected input] (= expected (part2 input))
                        71503 test-data
                        37286485 input-data)
)