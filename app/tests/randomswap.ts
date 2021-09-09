import React from "react"
const randomSwap = (a: string, b: string, c: string, d: string):string[] => {
  let array = [a, b, c, d]
  for (let i = array.length - 1; i > 0; i--) {
    let j = Math.floor(Math.random() * (i + 1));
    let tmp = swap(array[i],array[j])
    array[i] = tmp[0]
    array[j] = tmp[1]
  }
  return array
}
export const swap = (a:string,b:string):string[] =>{
  return [b,a]
}
export default randomSwap