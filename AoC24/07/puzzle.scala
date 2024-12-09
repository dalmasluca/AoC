import java.math.BigInteger
import scala.io.Source

class Day7(path: String) {
  val lines = Source
    .fromFile(path)
    .getLines()
    .toList
    .map(_.split(":"))

  val operator1 = List("+", "*")
  val operator2 = List("+", "*", "|")

  def combinationsWithRepetition(s: List[String], m: Int): List[String] = {
    def combine(current: String, depth: Int): List[String] = {
      if (depth == m) {
        List(current)
      } else {
        (0 until s.size).toList.flatMap(i => combine(current + s(i), depth + 1))
      }
    }

    combine("", 0)
  }

  def applyOperator(number: List[BigInteger], operator: String): BigInteger = {
    var op = operator
    number.slice(1, number.size).foldLeft(number(0)) { (acc, el) =>
      op(0) match {
        case '*' => { op = op.slice(1, op.size); (acc.multiply(el)) }
        case '+' => { op = op.slice(1, op.size); (acc.add(el)) }
        case '|' => {
          op = op.slice(1, op.size);
          (new BigInteger(acc.toString + el.toString))
        }
        case _ => { op = op.slice(1, op.size); (acc.add(new BigInteger("0"))) }
      }
    }
  }

  def part1: BigInteger = {
    lines
      .map { case Array(x, xs) =>
        val numbers = xs.trim.split(" ").map(new BigInteger(_)).toList
        val number  = new BigInteger(x.trim)
        if (
          combinationsWithRepetition(operator1, numbers.size - 1)
            .map { operators =>
              applyOperator(numbers, operators)
            }
            .exists(_ == number)
        )
          number
        else new BigInteger("0")
      }
      .foldLeft(new BigInteger("0")) { (acc, el) => acc.add(el) }
  }
  def part2: BigInteger = {
    lines
      .map { case Array(x, xs) =>
        val numbers = xs.trim.split(" ").map(new BigInteger(_)).toList
        val number  = new BigInteger(x.trim)
        if (
          combinationsWithRepetition(operator2, numbers.size - 1)
            .map { operators =>
              applyOperator(numbers, operators)
            }
            .exists(_ == number)
        )
          number
        else new BigInteger("0")
      }
      .foldLeft(new BigInteger("0")) { (acc, el) => acc.add(el) }

  }
}
