package day

object AoC {
  def main(args: Array[String]) = {
    val test = new Day("small-input.txt")
    val real = new Day("input.txt")

    if (test.part1 == 11) {
      println(real.part1)
    }
    if (test.part2 == 0) {
      println(real.part2)
    }
  }
}
