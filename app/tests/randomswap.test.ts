import React from "react";
import { swap } from "./randomswap";
test("swap test", () => {
  expect(swap("hello", "world")).toEqual(["world", "hello"])
})