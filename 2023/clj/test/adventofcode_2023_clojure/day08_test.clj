(ns adventofcode-2023-clojure.day08-test
  (:require [clojure.test :refer :all]
            [adventofcode-2023-clojure.day08 :refer :all]))

(def input-data  (slurp "resources/day08.txt"))
(def test-data  (slurp "resources/day08_test.txt"))
(def test-data2  (slurp "resources/day08_test2.txt"))


(deftest part1-test
  (are [expected input] (= expected (part1 input))
                        6 test-data
                        23147 input-data)
)

;; (deftest part2-test
;;   (are [expected input] (= expected (part2 input))
;;                         6 test-data2
;;                         22289513667691 input-data)
;; )