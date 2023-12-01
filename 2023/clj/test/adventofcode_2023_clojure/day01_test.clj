(ns adventofcode-2023-clojure.day01-test
  (:require [clojure.test :refer :all]
            [adventofcode-2023-clojure.day01 :refer :all]))

(deftest part1-test
  (let [test-data (str "1abc2\n"
                       "pqr3stu8vwx\n"
                       "a1b2c3d4e5f\n"
                       "treb7uchet")]
  (are [expected input] (= expected (part1 input))
                         142 test-data))
)

(deftest part2-test
  (let [test-data (str "two1nine\n"
                        "eightwothree\n"
                        "abcone2threexyz\n"
                        "xtwone3four\n"
                        "4nineeightseven2\n"
                        "zoneight234\n"
                        "7pqrstsixteen")]
  (are [expected input] (= expected (part2 input))
                        281 test-data))
)