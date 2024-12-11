import scala.io.Source

object Day {

  def main(args: Array[String]): Unit = {
    val inputs = List(
      Source
        .fromFile(args(0))
        .getLines
        .toList,
      Source
        .fromFile(args(1))
        .getLines
        .toList
    )

    if (part1(inputs(0)) == 0) println("part1: " + part1(inputs(1)))
    if (part2(inputs(0)) == 0) println("part2: " + part2(inputs(1)))
  }

  def part1(input: List[String]): Int = {
    0
  }

  def part2(input: List[String]): Int = {
    0
  }
}
