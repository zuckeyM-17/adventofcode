(ns adventofcode-2023-clojure.day07-test
  (:require [clojure.test :refer :all]
            [adventofcode-2023-clojure.day07 :refer :all]))

(def input-data  (slurp "resources/day07.txt"))
(def test-data  (slurp "resources/day07_test.txt"))

(deftest part1-test
  (are [expected input] (= expected (part1 input))
                        6440 test-data
                        249638405 input-data)
)

(deftest part2-test
  (are [expected input] (= expected (part2 input))
                        5905 test-data
                        249776650 input-data)
)