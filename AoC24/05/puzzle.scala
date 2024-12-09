import scala.io.Source

class Day5(path: String) {
  val Array(rawRules, rawOrder) = Source
    .fromFile(path)
    .getLines()
    .toList
    .mkString(" ")
    .split("  ")

  val rules  = rawRules.split(" ").map(_.split("\\|").map(_.toInt))
  val orders = rawOrder.split(" ").map(order => order.split(",").map(_.toInt))
  def part1: Int = {
    orders.map { order =>
      val map = order.zipWithIndex.toMap;
      if (
        rules
          .map { case Array(first, second) =>
            map.get(first) match {
              case None => true
              case Some(position) =>
                position < map.getOrElse(second, Int.MaxValue)
            }
          }
          .forall(p => p)
      ) order(order.length / 2)
      else 0
    }.sum
  }
  def part2: Int = {
    orders.map { order =>
      val map = order.zipWithIndex.toMap;
      if (
        rules
          .map { case Array(first, second) =>
            map.get(first) match {
              case None => true
              case Some(position) =>
                position < map.getOrElse(second, Int.MaxValue)
            }
          }
          .forall(p => p)
      ) 0
      else {}
    }.sum
  }
}
