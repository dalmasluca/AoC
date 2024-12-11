import scala.util.matching.Regex
import scala.io.Source

class Day3(path: String) {
  val lines = Source
    .fromFile(path)
    .getLines()
    .toList

  private val pattern: Regex = raw"mul\(\d{1,3},\d{1,3}\)".r
  private val delimiter: Regex =
    """(?:^|do\(\))(.*?)(?:don't\(\)|$)""".r
  private val number: Regex = raw"\d+".r

  def part1: Int = {
    val result = lines.map { line =>
      val matched = pattern.findAllMatchIn(line)
      val numbers = matched.map { elem =>
        number
          .findAllMatchIn(elem.matched)
          .map(_.matched.toInt)
          .foldLeft(1) { (acc, el) => acc * el }
      }.sum
      numbers
    }
    result.sum
  }
  def part2: Int = {
    delimiter
      .findAllMatchIn(lines.mkString)
      .map { matched =>
        pattern
          .findAllMatchIn(matched.matched)
          .map { elem =>
            number
              .findAllMatchIn(elem.matched)
              .map(_.matched.toInt)
              .foldLeft(1) { (acc, el) => acc * el }
          }
          .sum
      }
      .sum
  }
}
