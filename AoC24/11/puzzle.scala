import scala.io.Source
import java.math.BigInteger

object Day11 {

  private def trasnform(stone: BigInt): List[BigInt] = {
    if (stone == 0) List(1)
    else {
      val len = stone.toString.length
      if (len % 2 == 0) {
        val num = stone.toString
        List(num.substring(0, len / 2).toInt, num.substring(len / 2).toInt)
      } else List(stone * 2024)
    }
  }

  // private def blink(stones: List[BigInt], times: Int): List[Int] = {
  // stones.map { stone =>
  // var afterblink = trasnform(stone)
  // (1 until times).foreach { _ =>
  // afterblink = afterblink.flatMap { stone => trasnform(stone) }
  //     }
  //   afterblink.size
  //  }
  // }

  private def blink(stones: List[BigInt], times: Int): BigInt = {
    var stonesMap =
      stones
        .groupBy(identity)
        .view
        .mapValues(l => new BigInt(new BigInteger(l.size.toString)))
        .toMap

    (0 until times).foreach { _ =>
      var copy: Map[BigInt, BigInt] = Map[BigInt, BigInt]()
      stonesMap.keys.foreach { key =>
        val keyAfter = trasnform(key)
        val value    = stonesMap(key)
        keyAfter.foreach { key =>
          copy = copy.updated(key, copy.getOrElse(key, BigInt(0)) + value)
        }
      }
      stonesMap = copy
    }

    stonesMap.values.toList.foldLeft(BigInt(0)) { (acc, el) => acc + el }

  }

  def part1(input: List[BigInt]): BigInt = {
    blink(input, 25)
    // (1 until 6).foreach(_ => afterblink = blink(afterblink))
  }
  def part2(input: List[BigInt]): BigInt = {
    blink(input, 75)
  }

  def main(args: Array[String]): Unit = {
    val lines = Source
      .fromFile("input.txt")
      .mkString
      .replace("\n", "")
      .split(" ")
      .map(maybenum => new BigInt(new BigInteger(maybenum)))
      .toList
    println("part1: " + part1(lines))
    println("part2: " + part2(lines))
  }
}
